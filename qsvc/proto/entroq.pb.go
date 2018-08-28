// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entroq.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	entroq.proto

It has these top-level messages:
	TaskID
	TaskData
	TaskChange
	Task
	QueueStats
	ClaimRequest
	ClaimResponse
	ModifyRequest
	ModifyResponse
	ModifyDep
	TasksRequest
	TasksResponse
	QueuesRequest
	QueuesResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type DepType int32

const (
	DepType_CLAIM  DepType = 0
	DepType_DELETE DepType = 1
	DepType_CHANGE DepType = 2
	DepType_DEPEND DepType = 3
)

var DepType_name = map[int32]string{
	0: "CLAIM",
	1: "DELETE",
	2: "CHANGE",
	3: "DEPEND",
}
var DepType_value = map[string]int32{
	"CLAIM":  0,
	"DELETE": 1,
	"CHANGE": 2,
	"DEPEND": 3,
}

func (x DepType) String() string {
	return proto1.EnumName(DepType_name, int32(x))
}
func (DepType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TaskID contains the ID and version of a task. Together these make a unique
// identifier for that task.
type TaskID struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *TaskID) Reset()                    { *m = TaskID{} }
func (m *TaskID) String() string            { return proto1.CompactTextString(m) }
func (*TaskID) ProtoMessage()               {}
func (*TaskID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TaskID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TaskID) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

// TaskData contains only the data portion of a task. Useful for insertion.
type TaskData struct {
	// The name of the queue for this task.
	Queue string `protobuf:"bytes,1,opt,name=queue" json:"queue,omitempty"`
	// The epoch time in millis when this task becomes available.
	AtMs int64 `protobuf:"varint,2,opt,name=at_ms,json=atMs" json:"at_ms,omitempty"`
	// The task's opaque payload.
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TaskData) Reset()                    { *m = TaskData{} }
func (m *TaskData) String() string            { return proto1.CompactTextString(m) }
func (*TaskData) ProtoMessage()               {}
func (*TaskData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskData) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *TaskData) GetAtMs() int64 {
	if m != nil {
		return m.AtMs
	}
	return 0
}

func (m *TaskData) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// TaskChange identifies a task by ID and specifies the new data it should contain.
// All fields should be filled in. Empty fields result in deleting data from that field.
type TaskChange struct {
	OldId   *TaskID   `protobuf:"bytes,1,opt,name=old_id,json=oldId" json:"old_id,omitempty"`
	NewData *TaskData `protobuf:"bytes,2,opt,name=new_data,json=newData" json:"new_data,omitempty"`
}

func (m *TaskChange) Reset()                    { *m = TaskChange{} }
func (m *TaskChange) String() string            { return proto1.CompactTextString(m) }
func (*TaskChange) ProtoMessage()               {}
func (*TaskChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskChange) GetOldId() *TaskID {
	if m != nil {
		return m.OldId
	}
	return nil
}

func (m *TaskChange) GetNewData() *TaskData {
	if m != nil {
		return m.NewData
	}
	return nil
}

// Task is a complete task object, containing IDs, data, and metadata.
type Task struct {
	// The name of the queue for this task.
	Queue   string `protobuf:"bytes,1,opt,name=queue" json:"queue,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Version int32  `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	// The epoch time in millis when this task becomes available.
	AtMs int64 `protobuf:"varint,4,opt,name=at_ms,json=atMs" json:"at_ms,omitempty"`
	// The UUID representing the claimant (owner) for this task.
	ClaimantId string `protobuf:"bytes,5,opt,name=claimant_id,json=claimantId" json:"claimant_id,omitempty"`
	// The task's opaque payload.
	Value []byte `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	// Epoch times in millis for creation and update of this task.
	CreatedMs  int64 `protobuf:"varint,7,opt,name=created_ms,json=createdMs" json:"created_ms,omitempty"`
	ModifiedMs int64 `protobuf:"varint,8,opt,name=modified_ms,json=modifiedMs" json:"modified_ms,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto1.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Task) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *Task) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Task) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Task) GetAtMs() int64 {
	if m != nil {
		return m.AtMs
	}
	return 0
}

func (m *Task) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *Task) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Task) GetCreatedMs() int64 {
	if m != nil {
		return m.CreatedMs
	}
	return 0
}

func (m *Task) GetModifiedMs() int64 {
	if m != nil {
		return m.ModifiedMs
	}
	return 0
}

// QueueStats contains the name of the queue and the number of tasks within it.
type QueueStats struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NumTasks int32  `protobuf:"varint,2,opt,name=num_tasks,json=numTasks" json:"num_tasks,omitempty"`
}

func (m *QueueStats) Reset()                    { *m = QueueStats{} }
func (m *QueueStats) String() string            { return proto1.CompactTextString(m) }
func (*QueueStats) ProtoMessage()               {}
func (*QueueStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueueStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueueStats) GetNumTasks() int32 {
	if m != nil {
		return m.NumTasks
	}
	return 0
}

// ClaimRequest is sent to attempt to claim a task from a queue. The claimant ID
// should be unique to the requesting worker (e.g., if multiple workers are in
// the same process, they should all have different claimant IDs assigned).
type ClaimRequest struct {
	ClaimantId string `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId" json:"claimant_id,omitempty"`
	Queue      string `protobuf:"bytes,2,opt,name=queue" json:"queue,omitempty"`
	DurationMs int64  `protobuf:"varint,3,opt,name=duration_ms,json=durationMs" json:"duration_ms,omitempty"`
}

func (m *ClaimRequest) Reset()                    { *m = ClaimRequest{} }
func (m *ClaimRequest) String() string            { return proto1.CompactTextString(m) }
func (*ClaimRequest) ProtoMessage()               {}
func (*ClaimRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClaimRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *ClaimRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *ClaimRequest) GetDurationMs() int64 {
	if m != nil {
		return m.DurationMs
	}
	return 0
}

// ClaimResponse is returned when a claim is fulfilled or becomes obviously impossible.
type ClaimResponse struct {
	Task *Task `protobuf:"bytes,1,opt,name=task" json:"task,omitempty"`
}

func (m *ClaimResponse) Reset()                    { *m = ClaimResponse{} }
func (m *ClaimResponse) String() string            { return proto1.CompactTextString(m) }
func (*ClaimResponse) ProtoMessage()               {}
func (*ClaimResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ClaimResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// ModifyRequest sends a request to modify a set of tasks with given
// dependencies. It is performed in a transaction, in which either all
// suggested modifications succeed and all dependencies are satisfied, or
// nothing is committed at all. A failure due to dependencies (in any
// of changes, deletes, or inserts) will be permanent.
//
// All successful changes will cause the requester to become the claimant.
type ModifyRequest struct {
	ClaimantId string        `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId" json:"claimant_id,omitempty"`
	Inserts    []*TaskData   `protobuf:"bytes,2,rep,name=inserts" json:"inserts,omitempty"`
	Changes    []*TaskChange `protobuf:"bytes,3,rep,name=changes" json:"changes,omitempty"`
	Deletes    []*TaskID     `protobuf:"bytes,4,rep,name=deletes" json:"deletes,omitempty"`
	Depends    []*TaskID     `protobuf:"bytes,5,rep,name=depends" json:"depends,omitempty"`
}

func (m *ModifyRequest) Reset()                    { *m = ModifyRequest{} }
func (m *ModifyRequest) String() string            { return proto1.CompactTextString(m) }
func (*ModifyRequest) ProtoMessage()               {}
func (*ModifyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ModifyRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *ModifyRequest) GetInserts() []*TaskData {
	if m != nil {
		return m.Inserts
	}
	return nil
}

func (m *ModifyRequest) GetChanges() []*TaskChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *ModifyRequest) GetDeletes() []*TaskID {
	if m != nil {
		return m.Deletes
	}
	return nil
}

func (m *ModifyRequest) GetDepends() []*TaskID {
	if m != nil {
		return m.Depends
	}
	return nil
}

// ModifyResponse returns inserted and updated tasks when successful.
// A dependency error (which is permanent) comes through as gRPC's NOT_FOUND code.
type ModifyResponse struct {
	Inserted []*Task `protobuf:"bytes,1,rep,name=inserted" json:"inserted,omitempty"`
	Changed  []*Task `protobuf:"bytes,2,rep,name=changed" json:"changed,omitempty"`
}

func (m *ModifyResponse) Reset()                    { *m = ModifyResponse{} }
func (m *ModifyResponse) String() string            { return proto1.CompactTextString(m) }
func (*ModifyResponse) ProtoMessage()               {}
func (*ModifyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ModifyResponse) GetInserted() []*Task {
	if m != nil {
		return m.Inserted
	}
	return nil
}

func (m *ModifyResponse) GetChanged() []*Task {
	if m != nil {
		return m.Changed
	}
	return nil
}

// ModifyDep can be returned with a gRPC NotFound status indicating which
// dependencies failed. This is done via the gRPC error return, not directly
// in the response proto.
type ModifyDep struct {
	Type DepType `protobuf:"varint,1,opt,name=type,enum=proto.DepType" json:"type,omitempty"`
	Id   *TaskID `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (m *ModifyDep) Reset()                    { *m = ModifyDep{} }
func (m *ModifyDep) String() string            { return proto1.CompactTextString(m) }
func (*ModifyDep) ProtoMessage()               {}
func (*ModifyDep) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ModifyDep) GetType() DepType {
	if m != nil {
		return m.Type
	}
	return DepType_CLAIM
}

func (m *ModifyDep) GetId() *TaskID {
	if m != nil {
		return m.Id
	}
	return nil
}

// TasksRequest is sent to request a complete listing of tasks for the
// given queue. If claimant_id is empty, all tasks (not just expired
// or owned tasks) are returned.
type TasksRequest struct {
	ClaimantId string `protobuf:"bytes,1,opt,name=claimant_id,json=claimantId" json:"claimant_id,omitempty"`
	Queue      string `protobuf:"bytes,2,opt,name=queue" json:"queue,omitempty"`
}

func (m *TasksRequest) Reset()                    { *m = TasksRequest{} }
func (m *TasksRequest) String() string            { return proto1.CompactTextString(m) }
func (*TasksRequest) ProtoMessage()               {}
func (*TasksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TasksRequest) GetClaimantId() string {
	if m != nil {
		return m.ClaimantId
	}
	return ""
}

func (m *TasksRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

// TasksReqponse contains the tasks requested.
type TasksResponse struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *TasksResponse) Reset()                    { *m = TasksResponse{} }
func (m *TasksResponse) String() string            { return proto1.CompactTextString(m) }
func (*TasksResponse) ProtoMessage()               {}
func (*TasksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TasksResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// QueuesRequest is sent to request a listing of all known queues.
type QueuesRequest struct {
}

func (m *QueuesRequest) Reset()                    { *m = QueuesRequest{} }
func (m *QueuesRequest) String() string            { return proto1.CompactTextString(m) }
func (*QueuesRequest) ProtoMessage()               {}
func (*QueuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

// QueuesResponse contains the requested list of queue statistics.
type QueuesResponse struct {
	Queues []*QueueStats `protobuf:"bytes,1,rep,name=queues" json:"queues,omitempty"`
}

func (m *QueuesResponse) Reset()                    { *m = QueuesResponse{} }
func (m *QueuesResponse) String() string            { return proto1.CompactTextString(m) }
func (*QueuesResponse) ProtoMessage()               {}
func (*QueuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *QueuesResponse) GetQueues() []*QueueStats {
	if m != nil {
		return m.Queues
	}
	return nil
}

func init() {
	proto1.RegisterType((*TaskID)(nil), "proto.TaskID")
	proto1.RegisterType((*TaskData)(nil), "proto.TaskData")
	proto1.RegisterType((*TaskChange)(nil), "proto.TaskChange")
	proto1.RegisterType((*Task)(nil), "proto.Task")
	proto1.RegisterType((*QueueStats)(nil), "proto.QueueStats")
	proto1.RegisterType((*ClaimRequest)(nil), "proto.ClaimRequest")
	proto1.RegisterType((*ClaimResponse)(nil), "proto.ClaimResponse")
	proto1.RegisterType((*ModifyRequest)(nil), "proto.ModifyRequest")
	proto1.RegisterType((*ModifyResponse)(nil), "proto.ModifyResponse")
	proto1.RegisterType((*ModifyDep)(nil), "proto.ModifyDep")
	proto1.RegisterType((*TasksRequest)(nil), "proto.TasksRequest")
	proto1.RegisterType((*TasksResponse)(nil), "proto.TasksResponse")
	proto1.RegisterType((*QueuesRequest)(nil), "proto.QueuesRequest")
	proto1.RegisterType((*QueuesResponse)(nil), "proto.QueuesResponse")
	proto1.RegisterEnum("proto.DepType", DepType_name, DepType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EntroQ service

type EntroQClient interface {
	TryClaim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
	Claim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
	Modify(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error)
	Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error)
	Queues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*QueuesResponse, error)
}

type entroQClient struct {
	cc *grpc.ClientConn
}

func NewEntroQClient(cc *grpc.ClientConn) EntroQClient {
	return &entroQClient{cc}
}

func (c *entroQClient) TryClaim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := grpc.Invoke(ctx, "/proto.EntroQ/TryClaim", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Claim(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := grpc.Invoke(ctx, "/proto.EntroQ/Claim", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Modify(ctx context.Context, in *ModifyRequest, opts ...grpc.CallOption) (*ModifyResponse, error) {
	out := new(ModifyResponse)
	err := grpc.Invoke(ctx, "/proto.EntroQ/Modify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Tasks(ctx context.Context, in *TasksRequest, opts ...grpc.CallOption) (*TasksResponse, error) {
	out := new(TasksResponse)
	err := grpc.Invoke(ctx, "/proto.EntroQ/Tasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entroQClient) Queues(ctx context.Context, in *QueuesRequest, opts ...grpc.CallOption) (*QueuesResponse, error) {
	out := new(QueuesResponse)
	err := grpc.Invoke(ctx, "/proto.EntroQ/Queues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EntroQ service

type EntroQServer interface {
	TryClaim(context.Context, *ClaimRequest) (*ClaimResponse, error)
	Claim(context.Context, *ClaimRequest) (*ClaimResponse, error)
	Modify(context.Context, *ModifyRequest) (*ModifyResponse, error)
	Tasks(context.Context, *TasksRequest) (*TasksResponse, error)
	Queues(context.Context, *QueuesRequest) (*QueuesResponse, error)
}

func RegisterEntroQServer(s *grpc.Server, srv EntroQServer) {
	s.RegisterService(&_EntroQ_serviceDesc, srv)
}

func _EntroQ_TryClaim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).TryClaim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/TryClaim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).TryClaim(ctx, req.(*ClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Claim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Claim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Claim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Claim(ctx, req.(*ClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Modify(ctx, req.(*ModifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Tasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Tasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Tasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Tasks(ctx, req.(*TasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntroQ_Queues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntroQServer).Queues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EntroQ/Queues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntroQServer).Queues(ctx, req.(*QueuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntroQ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EntroQ",
	HandlerType: (*EntroQServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TryClaim",
			Handler:    _EntroQ_TryClaim_Handler,
		},
		{
			MethodName: "Claim",
			Handler:    _EntroQ_Claim_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _EntroQ_Modify_Handler,
		},
		{
			MethodName: "Tasks",
			Handler:    _EntroQ_Tasks_Handler,
		},
		{
			MethodName: "Queues",
			Handler:    _EntroQ_Queues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "entroq.proto",
}

func init() { proto1.RegisterFile("entroq.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdb, 0x6a, 0xdb, 0x4a,
	0x14, 0x3d, 0x92, 0x2c, 0xd9, 0xde, 0xbe, 0xc4, 0x67, 0x92, 0x03, 0x26, 0x87, 0x90, 0x74, 0x68,
	0xc9, 0xa5, 0x10, 0x8a, 0x4b, 0xa0, 0x50, 0xfa, 0x10, 0x6c, 0xd3, 0x9a, 0xc6, 0xa1, 0x51, 0xfd,
	0x5c, 0x77, 0x9a, 0x99, 0xb4, 0x22, 0xb6, 0xa4, 0x68, 0x46, 0x09, 0xfe, 0x90, 0xfe, 0x58, 0x1f,
	0xfb, 0x35, 0x65, 0x6e, 0xb6, 0x94, 0x38, 0xd0, 0xd2, 0x27, 0x4b, 0x7b, 0xed, 0xcb, 0x5a, 0x7b,
	0x2f, 0x19, 0x9a, 0x2c, 0x16, 0x59, 0x72, 0x73, 0x9c, 0x66, 0x89, 0x48, 0x90, 0xaf, 0x7e, 0x70,
	0x0f, 0x82, 0x09, 0xe1, 0xd7, 0xa3, 0x01, 0x6a, 0x83, 0x1b, 0xd1, 0xae, 0xb3, 0xe7, 0x1c, 0xd4,
	0x43, 0x37, 0xa2, 0xa8, 0x0b, 0xd5, 0x5b, 0x96, 0xf1, 0x28, 0x89, 0xbb, 0xee, 0x9e, 0x73, 0xe0,
	0x87, 0xf6, 0x15, 0xbf, 0x87, 0x9a, 0xac, 0x19, 0x10, 0x41, 0xd0, 0x16, 0xf8, 0x37, 0x39, 0xcb,
	0x99, 0x29, 0xd4, 0x2f, 0x68, 0x13, 0x7c, 0x22, 0xa6, 0x73, 0xae, 0x2a, 0xbd, 0xb0, 0x42, 0xc4,
	0x98, 0xcb, 0xd4, 0x5b, 0x32, 0xcb, 0x59, 0xd7, 0xdb, 0x73, 0x0e, 0x9a, 0xa1, 0x7e, 0xc1, 0x9f,
	0x00, 0x64, 0xb3, 0xfe, 0x37, 0x12, 0x7f, 0x65, 0xe8, 0x29, 0x04, 0xc9, 0x8c, 0x4e, 0x0d, 0x91,
	0x46, 0xaf, 0xa5, 0xd9, 0x1e, 0x6b, 0x8e, 0xa1, 0x9f, 0xcc, 0xe8, 0x88, 0xa2, 0x23, 0xa8, 0xc5,
	0xec, 0x6e, 0x4a, 0x89, 0x20, 0x6a, 0x42, 0xa3, 0xb7, 0x51, 0xc8, 0x93, 0xbc, 0xc2, 0x6a, 0xcc,
	0xee, 0xe4, 0x03, 0xfe, 0xe1, 0x40, 0x45, 0x46, 0x1f, 0x61, 0xaa, 0x55, 0xbb, 0xeb, 0x54, 0x7b,
	0x25, 0xd5, 0x2b, 0x4d, 0x95, 0x82, 0xa6, 0x5d, 0x68, 0x5c, 0xce, 0x48, 0x34, 0x27, 0xb1, 0x90,
	0xa4, 0x7d, 0xd5, 0x07, 0x6c, 0x68, 0x44, 0x57, 0xa2, 0x83, 0x82, 0x68, 0xb4, 0x03, 0x70, 0x99,
	0x31, 0x22, 0x18, 0x95, 0x0d, 0xab, 0xaa, 0x61, 0xdd, 0x44, 0x74, 0xd7, 0x79, 0x42, 0xa3, 0xab,
	0x48, 0xe3, 0x35, 0x85, 0x83, 0x0d, 0x8d, 0x39, 0x7e, 0x03, 0x70, 0x21, 0xe9, 0x7f, 0x14, 0x44,
	0x70, 0x84, 0xa0, 0x12, 0x93, 0xb9, 0x15, 0xa6, 0x9e, 0xd1, 0xff, 0x50, 0x8f, 0xf3, 0xf9, 0x54,
	0x10, 0x7e, 0xcd, 0xcd, 0xfd, 0x6a, 0x71, 0x3e, 0x97, 0x9b, 0xe0, 0xf8, 0x0a, 0x9a, 0x7d, 0x49,
	0x31, 0x64, 0x37, 0x39, 0xe3, 0xe2, 0xbe, 0x0a, 0x67, 0x9d, 0x0a, 0xbd, 0x3b, 0xb7, 0xb8, 0xbb,
	0x5d, 0x68, 0xd0, 0x3c, 0x23, 0x22, 0x4a, 0x62, 0x49, 0xd3, 0xd3, 0x34, 0x6d, 0x68, 0xcc, 0xf1,
	0x0b, 0x68, 0x99, 0x39, 0x3c, 0x4d, 0x62, 0x2e, 0x2b, 0x2a, 0x92, 0x91, 0x39, 0x6e, 0xa3, 0x70,
	0xb4, 0x50, 0x01, 0xf8, 0xa7, 0x03, 0xad, 0xb1, 0xd4, 0xb9, 0xf8, 0x6d, 0x6e, 0x87, 0x50, 0x8d,
	0x62, 0xce, 0x32, 0x21, 0x75, 0x7a, 0x6b, 0xbd, 0x60, 0x70, 0xf4, 0x1c, 0xaa, 0x97, 0xca, 0x67,
	0x92, 0xac, 0x4c, 0xfd, 0xb7, 0x90, 0xaa, 0x1d, 0x18, 0xda, 0x0c, 0xb4, 0x0f, 0x55, 0xca, 0x66,
	0x4c, 0x30, 0x79, 0x71, 0xef, 0xa1, 0x17, 0x2d, 0xaa, 0x13, 0x53, 0x16, 0x53, 0xde, 0xf5, 0x1f,
	0x49, 0x54, 0x28, 0xfe, 0x0c, 0x6d, 0xab, 0xcd, 0xec, 0x63, 0x1f, 0x6a, 0x9a, 0x1b, 0x93, 0xca,
	0xbc, 0xfb, 0x3b, 0x59, 0x82, 0xe8, 0x99, 0x65, 0x4e, 0x8d, 0xc8, 0x52, 0x9e, 0xc5, 0xf0, 0x39,
	0xd4, 0xf5, 0x84, 0x01, 0x4b, 0x11, 0x86, 0x8a, 0x58, 0xa4, 0xda, 0x16, 0xed, 0x5e, 0xdb, 0x14,
	0x0c, 0x58, 0x3a, 0x59, 0xa4, 0x2c, 0x54, 0x18, 0xda, 0x59, 0xda, 0xff, 0x01, 0x6d, 0x37, 0xa2,
	0x78, 0x08, 0x4d, 0xe5, 0x98, 0xbf, 0x33, 0x0a, 0xee, 0x41, 0xcb, 0xb4, 0x31, 0xba, 0x9f, 0x80,
	0xaf, 0x9d, 0xb9, 0x46, 0xb4, 0x46, 0xf0, 0x06, 0xb4, 0x94, 0xc5, 0xed, 0x6c, 0xfc, 0x1a, 0xda,
	0x36, 0x60, 0xba, 0x1c, 0x42, 0xa0, 0xfa, 0xdb, 0x36, 0xf6, 0x9a, 0xab, 0x4f, 0x23, 0x34, 0x09,
	0x47, 0xaf, 0xa0, 0x6a, 0x84, 0xa3, 0x3a, 0xf8, 0xfd, 0xb3, 0xd3, 0xd1, 0xb8, 0xf3, 0x0f, 0x02,
	0x08, 0x06, 0xc3, 0xb3, 0xe1, 0x64, 0xd8, 0x71, 0xe4, 0x73, 0xff, 0xdd, 0xe9, 0xf9, 0xdb, 0x61,
	0xc7, 0xd5, 0xf1, 0x0f, 0xc3, 0xf3, 0x41, 0xc7, 0xeb, 0x7d, 0x77, 0x21, 0x18, 0xca, 0x3f, 0xce,
	0x0b, 0x74, 0x02, 0xb5, 0x49, 0xb6, 0x50, 0x8e, 0x46, 0x9b, 0x66, 0x56, 0xf1, 0x3b, 0xda, 0xde,
	0x2a, 0x07, 0x0d, 0xcd, 0x1e, 0xf8, 0x7f, 0x5c, 0x73, 0x02, 0x81, 0x3e, 0x24, 0xb2, 0x78, 0xe9,
	0xab, 0xd8, 0xfe, 0xef, 0x5e, 0x74, 0x35, 0x4a, 0x2d, 0x7a, 0x39, 0xaa, 0x78, 0xbd, 0xe5, 0xa8,
	0xf2, 0x2d, 0x4e, 0x20, 0xd0, 0x7b, 0x5d, 0x8e, 0x2a, 0xed, 0x7d, 0x39, 0xaa, 0xbc, 0xfc, 0x2f,
	0x81, 0x8a, 0xbe, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x05, 0x01, 0x7f, 0xb6, 0x56, 0x06, 0x00,
	0x00,
}
