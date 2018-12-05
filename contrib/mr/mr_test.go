package mr

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/shiblon/entroq"
	"github.com/shiblon/entroq/mem"
)

func TestMemoryMR(t *testing.T) {
	ctx := context.Background()

	eq, err := entroq.New(ctx, mem.Opener())
	if err != nil {
		t.Fatal(err)
	}
	defer eq.Close()

	mr := NewMapReduce(eq, "/mrtest",
		WithNumMappers(2),
		WithNumReducers(1),
		WithMap(WordCountMapper),
		WithReduce(NilReducer),
		AddInput(NewKV(nil, []byte("word1 word2 word3 word4"))),
		AddInput(NewKV(nil, []byte("word1 word3 word5 word7"))),
		AddInput(NewKV(nil, []byte("word1 word4 word7 wordA"))),
		AddInput(NewKV(nil, []byte("word1 word5 word9 wordE"))))

	outQ, err := mr.Run(ctx)
	if err != nil {
		t.Fatal(err)
	}

	tasks, err := eq.Tasks(ctx, outQ)
	if err != nil {
		t.Fatal(err)
	}

	expected := []*KV{
		NewKV([]byte("word1"), nil),
		NewKV([]byte("word2"), nil),
		NewKV([]byte("word3"), nil),
		NewKV([]byte("word4"), nil),
		NewKV([]byte("word5"), nil),
		NewKV([]byte("word7"), nil),
		NewKV([]byte("word9"), nil),
		NewKV([]byte("wordA"), nil),
		NewKV([]byte("wordE"), nil),
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 final reduced output task, got %d", len(tasks))
	}

	task := tasks[0]

	var kvs []*KV
	if err := json.Unmarshal(task.Value, &kvs); err != nil {
		t.Fatal(err)
	}

	for i, kv := range kvs {
		if kv.String() != expected[i].String() {
			t.Errorf("Expected %s, got %s", expected[i], kv)
		}
	}
}