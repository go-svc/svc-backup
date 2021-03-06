// Code generated by protoc-gen-go.
// source: todo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	todo.proto

It has these top-level messages:
	Void
	Task
	Tasks
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Void 呈現一個什麼都沒有的資料。
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Task 是單個工作記事資料。
type Task struct {
	Title   string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// Tasks 會回傳多個工作記事資料。
type Tasks struct {
	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *Tasks) Reset()                    { *m = Tasks{} }
func (m *Tasks) String() string            { return proto.CompactTextString(m) }
func (*Tasks) ProtoMessage()               {}
func (*Tasks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tasks) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "pb.Void")
	proto.RegisterType((*Task)(nil), "pb.Task")
	proto.RegisterType((*Tasks)(nil), "pb.Tasks")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Todo service

type TodoClient interface {
	Add(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Tasks, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Add(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/pb.Todo/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Tasks, error) {
	out := new(Tasks)
	err := grpc.Invoke(ctx, "/pb.Todo/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Todo service

type TodoServer interface {
	Add(context.Context, *Task) (*Task, error)
	List(context.Context, *Void) (*Tasks, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Todo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Add(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Todo/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Todo_Add_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todo_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0xbf, 0x0b, 0xc2, 0x30,
	0x10, 0x85, 0xfb, 0x23, 0xad, 0xf6, 0xdc, 0x0e, 0x87, 0x50, 0x51, 0x4a, 0x16, 0x3b, 0x75, 0xa8,
	0xe0, 0xae, 0xb3, 0x53, 0x29, 0xee, 0xd6, 0x74, 0x08, 0x4a, 0x2f, 0x98, 0xfb, 0xff, 0x91, 0x34,
	0xb4, 0xdb, 0x7b, 0xef, 0xe3, 0xc1, 0x07, 0xc0, 0xa4, 0xa9, 0xb1, 0x3f, 0x62, 0xc2, 0xc4, 0x0e,
	0x2a, 0x07, 0xf1, 0x24, 0xa3, 0xd5, 0x15, 0x44, 0xff, 0x72, 0x1f, 0xdc, 0x43, 0xc6, 0x86, 0xbf,
	0xa3, 0x8c, 0xab, 0xb8, 0x2e, 0xba, 0x50, 0x50, 0xc2, 0xe6, 0x4d, 0x13, 0x8f, 0x13, 0xcb, 0x64,
	0xde, 0x97, 0xaa, 0xce, 0x90, 0xf9, 0x9f, 0xc3, 0x13, 0x64, 0xec, 0x83, 0x8c, 0xab, 0xb4, 0xde,
	0xb5, 0xdb, 0xc6, 0x0e, 0x8d, 0x27, 0x5d, 0x98, 0xdb, 0x3b, 0x88, 0x9e, 0x34, 0xe1, 0x01, 0xd2,
	0x9b, 0xd6, 0xb8, 0xf2, 0x72, 0x4d, 0x2a, 0xc2, 0x23, 0x88, 0x87, 0x71, 0x1c, 0xa8, 0xf7, 0x2a,
	0x8b, 0x85, 0x3a, 0x15, 0x0d, 0xf9, 0xec, 0x7d, 0xf9, 0x07, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x75,
	0x07, 0xbe, 0xc5, 0x00, 0x00, 0x00,
}
