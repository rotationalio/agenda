// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgendaClient is the client API for Agenda service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgendaClient interface {
	Schedule(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Daily(ctx context.Context, in *Day, opts ...grpc.CallOption) (*Docket, error)
}

type agendaClient struct {
	cc grpc.ClientConnInterface
}

func NewAgendaClient(cc grpc.ClientConnInterface) AgendaClient {
	return &agendaClient{cc}
}

func (c *agendaClient) Schedule(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, "/agenda.v1.Agenda/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agendaClient) Daily(ctx context.Context, in *Day, opts ...grpc.CallOption) (*Docket, error) {
	out := new(Docket)
	err := c.cc.Invoke(ctx, "/agenda.v1.Agenda/Daily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgendaServer is the server API for Agenda service.
// All implementations must embed UnimplementedAgendaServer
// for forward compatibility
type AgendaServer interface {
	Schedule(context.Context, *Item) (*Item, error)
	Daily(context.Context, *Day) (*Docket, error)
	mustEmbedUnimplementedAgendaServer()
}

// UnimplementedAgendaServer must be embedded to have forward compatible implementations.
type UnimplementedAgendaServer struct {
}

func (UnimplementedAgendaServer) Schedule(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (UnimplementedAgendaServer) Daily(context.Context, *Day) (*Docket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Daily not implemented")
}
func (UnimplementedAgendaServer) mustEmbedUnimplementedAgendaServer() {}

// UnsafeAgendaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgendaServer will
// result in compilation errors.
type UnsafeAgendaServer interface {
	mustEmbedUnimplementedAgendaServer()
}

func RegisterAgendaServer(s grpc.ServiceRegistrar, srv AgendaServer) {
	s.RegisterService(&Agenda_ServiceDesc, srv)
}

func _Agenda_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agenda.v1.Agenda/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServer).Schedule(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agenda_Daily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Day)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgendaServer).Daily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agenda.v1.Agenda/Daily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgendaServer).Daily(ctx, req.(*Day))
	}
	return interceptor(ctx, in, info, handler)
}

// Agenda_ServiceDesc is the grpc.ServiceDesc for Agenda service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agenda_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agenda.v1.Agenda",
	HandlerType: (*AgendaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _Agenda_Schedule_Handler,
		},
		{
			MethodName: "Daily",
			Handler:    _Agenda_Daily_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/agenda.proto",
}
