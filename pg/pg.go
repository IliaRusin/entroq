// Package pg provides a backend for a entroq.Client.
package pg

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shiblon/entroq"

	_ "github.com/lib/pq"
)

const (
	DefaultPassword = "password" // default postgres password
	DefaultUsername = "postgres" // default postgres username
	DefaultDBName   = "entroq"   // default postgres database
	DefaultSSLMode  = "disable"  // default postgres SSL mode
)

var escapePattern = regexp.MustCompile(`(['" \\])`)

func escapeOptVal(v string) string {
	return "'" + escapePattern.ReplaceAllString(v, "\\$1") + "'"
}

type PG struct {
	db *sql.DB

	dbName   string
	username string
	sslMode  string
	password string
}

// PGOption is used to pass options to New.
type PGOption func(pg *PG)

// WithDBName returns a PGOption that sets the database name for the
// connection string used to talk to Postgres.
//
// Usage:
//   New(WithDBName("mydb"))
func WithDBName(n string) PGOption {
	return func(p *PG) {
		p.dbName = n
	}
}

// WithUsername returns a CLientOption that sets the username for the
// Postgres connection string.
//
// Usage:
//   New(WithUsername("myusername"))
func WithUsername(n string) PGOption {
	return func(p *PG) {
		p.username = n
	}
}

// WithSSLMode sets the ssl mode in the connection string, e.g.,
//   New(WithSSLMode("disabled"))
func WithSSLMode(mode string) PGOption {
	return func(p *PG) {
		p.sslMode = mode
	}
}

// WithPassword sets the database password, e.g.,
//   New(WithPassword(getPassword()))
func WithPassword(pwd string) PGOption {
	return func(p *PG) {
		p.password = pwd
	}
}

// Opener creates a function that returns a new PostgreSQL backend with the given options.
// Options are created using package-level functions that produce a
// PGOption, e.g.,
//
//   openFunc, err := Open(
//   	WithDBName("postgres"),
//   	WithUsername("myuser"),
//   	WithPassword("thepassword"),
//   	WithSSLMode("disable"),
//   )
//
// Note that there are defaults for all of these fields, specified as constants.
// The default claimant value is a new random UUID, created in New.
func Opener(opts ...PGOption) entroq.Open {
	pg := &PG{
		username: DefaultUsername,
		dbName:   DefaultDBName,
		sslMode:  DefaultSSLMode,
		password: DefaultPassword,
	}
	for _, opt := range opts {
		opt(pg)
	}
	keyVals := []string{
		"user=" + escapeOptVal(pg.username),
		"dbname=" + escapeOptVal(pg.dbName),
		"sslmode=" + escapeOptVal(pg.sslMode),
		"password=" + escapeOptVal(pg.password),
	}

	return func(ctx context.Context) (entroq.Backend, error) {
		db, err := sql.Open("postgres", strings.Join(keyVals, " "))
		if err != nil {
			return nil, fmt.Errorf("failed to open database: %v", err)
		}
		pg.db = db

		if err := pg.initDB(ctx); err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to initialize database: %v", err)
		}
		return pg, nil
	}
}

// initDB sets up the database to have the appropriate tables and necessary
// extensions to work as a task queue backend.
func (p *PG) initDB(ctx context.Context) error {
	_, err := p.db.ExecContext(ctx, `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS tasks (
		  id UUID PRIMARY KEY NOT NULL DEFAULT UUID_GENERATE_V4(),
		  version INTEGER NOT NULL DEFAULT 0,
		  queue TEXT NOT NULL DEFAULT '',
		  at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		  created TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		  modified TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		  claimant UUID,
		  value BYTEA
		);
		CREATE INDEX IF NOT EXISTS byQueue ON tasks (queue);
		CREATE INDEX IF NOT EXISTS byQueueAT ON tasks (queue, at);
	`)
	if err != nil {
		return fmt.Errorf("failed to create (or reuse) tasks table: %v", err)
	}
	return nil
}

// Queues returns a slice of non-empty queue names.
func (p *PG) Queues(ctx context.Context) ([]string, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT DISTINCT queue FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("failed to get queue names: %v", err)
	}
	defer rows.Close()
	var queues []string
	for rows.Next() {
		q := ""
		if err := rows.Scan(&q); err != nil {
			return nil, fmt.Errorf("queue scan failed: %v", err)
		}
		queues = append(queues, q)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("queue iteration failed: %v", err)
	}
	return queues, nil
}

// Tasks returns a slice of all tasks in the given queue.
func (p *PG) Tasks(ctx context.Context, queue string, claimant uuid.UUID) ([]*entroq.Task, error) {
	var zeroID uuid.UUID
	values := []interface{}{queue}
	q := "SELECT id, version, queue, at, created, modified, claimant, value FROM tasks WHERE queue = $1"

	if claimant != zeroID {
		q += " AND (claimant = '00000000-0000-0000-0000-000000000000' OR claimant = $2 OR at < NOW())"
		values = append(values, claimant)
	}

	rows, err := p.db.QueryContext(ctx, q, values...)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks for queue %q: %v", queue, err)
	}
	defer rows.Close()
	var tasks []*entroq.Task
	for rows.Next() {
		t := &entroq.Task{}
		if err := rows.Scan(&t.ID, &t.Version, &t.Queue, &t.At, &t.Created, &t.Modified, &t.Claimant, &t.Value); err != nil {
			return nil, fmt.Errorf("task scan failed: %v", err)
		}
		tasks = append(tasks, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("task iteration failed for queue %q: %v", queue, err)
	}
	return tasks, nil
}

// TryClaim attempts to claim an "arrived" task from the queue.
// Returns an error if something goes wrong, a nil task if there is
// nothing to claim.
func (p *PG) TryClaim(ctx context.Context, queue string, claimant uuid.UUID, duration time.Duration) (*entroq.Task, error) {
	task := new(entroq.Task)
	if duration == 0 {
		return nil, fmt.Errorf("no duration set for claim")
	}
	err := p.db.QueryRowContext(ctx, `
		WITH top AS (
			SELECT * FROM tasks
			WHERE
				queue = $1 AND
				at <= NOW()
			ORDER BY at, version, id ASC
			LIMIT 1
		)
		UPDATE tasks
		SET
			version=version+1,
			at = $2,
			claimant = $3,
			modified = NOW()
		WHERE id IN (SELECT id FROM top)
		RETURNING id, version, queue, at, created, modified, claimant, value
	`, queue, time.Now().Add(duration), claimant).Scan(
		&task.ID,
		&task.Version,
		&task.Queue,
		&task.At,
		&task.Created,
		&task.Modified,
		&task.Claimant,
		&task.Value,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, nil
	}
	return task, nil
}

// Modify attempts to apply an atomic modification to the task
// store. Either all succeeds or all fails.
func (p *PG) Modify(ctx context.Context, claimant uuid.UUID, mod *entroq.Modification) (inserted []*entroq.Task, changed []*entroq.Task, err error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	foundDeps, err := depQuery(ctx, tx, mod)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get dependencies: %v", err)
	}

	if err := mod.DependencyError(foundDeps); err != nil {
		return nil, nil, err
	}

	// Once we get here, we know that all of the dependencies were present.
	// These should now all succeed.

	for _, td := range mod.Inserts {
		columns := []string{"queue", "claimant", "value"}
		values := []interface{}{td.Queue, claimant, td.Value}

		if !td.At.IsZero() {
			columns = append(columns, "at")
			values = append(values, td.At)
		}

		var placeholders []string
		for i := range columns {
			placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
		}

		q := "INSERT INTO tasks (" + strings.Join(columns, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ") RETURNING id, version, queue, at, claimant, value, created, modified"
		t := new(entroq.Task)
		if err := tx.QueryRowContext(ctx, q, values...).Scan(
			&t.ID,
			&t.Version,
			&t.Queue,
			&t.At,
			&t.Claimant,
			&t.Value,
			&t.Created,
			&t.Modified,
		); err != nil {
			return nil, nil, fmt.Errorf("insert failed in queue %q: %v", td.Queue, err)
		}
		inserted = append(inserted, t)
	}

	for _, tid := range mod.Deletes {
		q := "DELETE FROM tasks WHERE id = $1 AND version = $2"
		if _, err := tx.ExecContext(ctx, q, tid.ID, tid.Version); err != nil {
			return nil, nil, fmt.Errorf("delete failed for task %q: %v", tid.ID, err)
		}
	}

	for _, t := range mod.Changes {
		q := `UPDATE tasks SET
				version = version + 1,
				modified = NOW(),
				queue = $1,
				at = $2,
				value = $3
			WHERE id = $4 AND version = $5
			RETURNING id, version, queue, at, claimant, modified, created, value`
		nt := new(entroq.Task)
		row := tx.QueryRowContext(ctx, q, t.Queue, t.At, t.Value, t.ID, t.Version)
		if err := row.Scan(&nt.ID, &nt.Version, &nt.Queue, &nt.At, &nt.Claimant, &nt.Modified, &nt.Created, &nt.Value); err != nil {
			return nil, nil, fmt.Errorf("scan failed for newly-changed task %q: %v", t.ID, err)
		}
		changed = append(changed, nt)
	}

	return inserted, changed, nil
}

// depQuery sends a query to the database, within a transaction, that will lock
// all rows for all dependencies of a modification, allowing updates to be
// performed on those rows without other transactions getting in the middle of
// things. Returns a map from ID to version for every dependency that is found,
// or an error if the query itself fails.
func depQuery(ctx context.Context, tx *sql.Tx, m *entroq.Modification) (map[uuid.UUID]*entroq.Task, error) {
	// We must craft a query that ensures that changes, deletes, and depends
	// all exist with the right versions, and then insert all inserts, delete
	// all deletes, and update all changes.
	dependencies, err := m.AllDependencies()
	if err != nil {
		return nil, err
	}

	foundDeps := make(map[uuid.UUID]*entroq.Task)

	if len(dependencies) == 0 {
		return foundDeps, nil
	}

	// Form a SELECT FOR UPDATE that looks at all of these rows so that we can
	// touch them all while inserting, deleting, and changing, and so we can
	// guarantee that dependencies don't disappear while we're at it (we need
	// to know that the dependencies are there while we work).
	var placeholders []string
	var values []interface{}
	first := 1
	for id, version := range dependencies {
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d)", first, first+1))
		values = append(values, id, version)
		first++
	}
	instr := strings.Join(placeholders, ", ")
	rows, err := tx.QueryContext(ctx,
		"SELECT id, version, queue, at, claimant, value, created, modified FROM tasks WHERE (id, version) IN ("+instr+") FOR UPDATE",
		values...)
	if err != nil {
		return nil, fmt.Errorf("error in dependencies query: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		t := new(entroq.Task)
		if err := rows.Scan(&t.ID, &t.Version, &t.Queue, &t.At, &t.Claimant, &t.Value, &t.Created, &t.Modified); err != nil {
			return nil, fmt.Errorf("row scan failed: %v", err)
		}
		if foundDeps[t.ID] != nil {
			return nil, fmt.Errorf("duplicate ID %q found in database (more than one version)", t.ID)
		}
		foundDeps[t.ID] = t
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error scanning dependencies: %v", err)
	}

	return foundDeps, nil
}

// Close cleans up the database client connection.
func (p *PG) Close() error {
	return p.db.Close()
}
