// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SimWorkerClient is the client API for SimWorker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimWorkerClient interface {
	InitGnb(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
	InitUe(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GnbConfig, error)
	StartUeRegister(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartUeDeregister(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetAction(ctx context.Context, in *ActionProfile, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartAction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StopAction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type simWorkerClient struct {
	cc grpc.ClientConnInterface
}

func NewSimWorkerClient(cc grpc.ClientConnInterface) SimWorkerClient {
	return &simWorkerClient{cc}
}

func (c *simWorkerClient) InitGnb(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/InitGnb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) InitUe(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/InitUe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GnbConfig, error) {
	out := new(GnbConfig)
	err := c.cc.Invoke(ctx, "/api.sim_worker/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) StartUeRegister(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/StartUeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) StartUeDeregister(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/StartUeDeregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) SetAction(ctx context.Context, in *ActionProfile, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/SetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) StartAction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/StartAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simWorkerClient) StopAction(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_worker/StopAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimWorkerServer is the server API for SimWorker service.
// All implementations must embed UnimplementedSimWorkerServer
// for forward compatibility
type SimWorkerServer interface {
	InitGnb(context.Context, *GnbConfig) (*emptypb.Empty, error)
	InitUe(context.Context, *GnbConfig) (*emptypb.Empty, error)
	GetConfig(context.Context, *emptypb.Empty) (*GnbConfig, error)
	StartUeRegister(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	StartUeDeregister(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	SetAction(context.Context, *ActionProfile) (*emptypb.Empty, error)
	StartAction(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	StopAction(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedSimWorkerServer()
}

// UnimplementedSimWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedSimWorkerServer struct {
}

func (UnimplementedSimWorkerServer) InitGnb(context.Context, *GnbConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitGnb not implemented")
}
func (UnimplementedSimWorkerServer) InitUe(context.Context, *GnbConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitUe not implemented")
}
func (UnimplementedSimWorkerServer) GetConfig(context.Context, *emptypb.Empty) (*GnbConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedSimWorkerServer) StartUeRegister(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUeRegister not implemented")
}
func (UnimplementedSimWorkerServer) StartUeDeregister(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUeDeregister not implemented")
}
func (UnimplementedSimWorkerServer) SetAction(context.Context, *ActionProfile) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAction not implemented")
}
func (UnimplementedSimWorkerServer) StartAction(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAction not implemented")
}
func (UnimplementedSimWorkerServer) StopAction(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAction not implemented")
}
func (UnimplementedSimWorkerServer) mustEmbedUnimplementedSimWorkerServer() {}

// UnsafeSimWorkerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimWorkerServer will
// result in compilation errors.
type UnsafeSimWorkerServer interface {
	mustEmbedUnimplementedSimWorkerServer()
}

func RegisterSimWorkerServer(s grpc.ServiceRegistrar, srv SimWorkerServer) {
	s.RegisterService(&SimWorker_ServiceDesc, srv)
}

func _SimWorker_InitGnb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnbConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).InitGnb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/InitGnb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).InitGnb(ctx, req.(*GnbConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_InitUe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnbConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).InitUe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/InitUe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).InitUe(ctx, req.(*GnbConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_StartUeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).StartUeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/StartUeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).StartUeRegister(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_StartUeDeregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).StartUeDeregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/StartUeDeregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).StartUeDeregister(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_SetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).SetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/SetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).SetAction(ctx, req.(*ActionProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_StartAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).StartAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/StartAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).StartAction(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimWorker_StopAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimWorkerServer).StopAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_worker/StopAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimWorkerServer).StopAction(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SimWorker_ServiceDesc is the grpc.ServiceDesc for SimWorker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimWorker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sim_worker",
	HandlerType: (*SimWorkerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitGnb",
			Handler:    _SimWorker_InitGnb_Handler,
		},
		{
			MethodName: "InitUe",
			Handler:    _SimWorker_InitUe_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _SimWorker_GetConfig_Handler,
		},
		{
			MethodName: "StartUeRegister",
			Handler:    _SimWorker_StartUeRegister_Handler,
		},
		{
			MethodName: "StartUeDeregister",
			Handler:    _SimWorker_StartUeDeregister_Handler,
		},
		{
			MethodName: "SetAction",
			Handler:    _SimWorker_SetAction_Handler,
		},
		{
			MethodName: "StartAction",
			Handler:    _SimWorker_StartAction_Handler,
		},
		{
			MethodName: "StopAction",
			Handler:    _SimWorker_StopAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uegnbsim.proto",
}

// SimMasterClient is the client API for SimMaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimMasterClient interface {
	StreamChannel(ctx context.Context, opts ...grpc.CallOption) (SimMaster_StreamChannelClient, error)
}

type simMasterClient struct {
	cc grpc.ClientConnInterface
}

func NewSimMasterClient(cc grpc.ClientConnInterface) SimMasterClient {
	return &simMasterClient{cc}
}

func (c *simMasterClient) StreamChannel(ctx context.Context, opts ...grpc.CallOption) (SimMaster_StreamChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &SimMaster_ServiceDesc.Streams[0], "/api.sim_master/StreamChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &simMasterStreamChannelClient{stream}
	return x, nil
}

type SimMaster_StreamChannelClient interface {
	Send(*emptypb.Empty) error
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type simMasterStreamChannelClient struct {
	grpc.ClientStream
}

func (x *simMasterStreamChannelClient) Send(m *emptypb.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *simMasterStreamChannelClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimMasterServer is the server API for SimMaster service.
// All implementations must embed UnimplementedSimMasterServer
// for forward compatibility
type SimMasterServer interface {
	StreamChannel(SimMaster_StreamChannelServer) error
	mustEmbedUnimplementedSimMasterServer()
}

// UnimplementedSimMasterServer must be embedded to have forward compatible implementations.
type UnimplementedSimMasterServer struct {
}

func (UnimplementedSimMasterServer) StreamChannel(SimMaster_StreamChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamChannel not implemented")
}
func (UnimplementedSimMasterServer) mustEmbedUnimplementedSimMasterServer() {}

// UnsafeSimMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimMasterServer will
// result in compilation errors.
type UnsafeSimMasterServer interface {
	mustEmbedUnimplementedSimMasterServer()
}

func RegisterSimMasterServer(s grpc.ServiceRegistrar, srv SimMasterServer) {
	s.RegisterService(&SimMaster_ServiceDesc, srv)
}

func _SimMaster_StreamChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SimMasterServer).StreamChannel(&simMasterStreamChannelServer{stream})
}

type SimMaster_StreamChannelServer interface {
	Send(*emptypb.Empty) error
	Recv() (*emptypb.Empty, error)
	grpc.ServerStream
}

type simMasterStreamChannelServer struct {
	grpc.ServerStream
}

func (x *simMasterStreamChannelServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *simMasterStreamChannelServer) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SimMaster_ServiceDesc is the grpc.ServiceDesc for SimMaster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimMaster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sim_master",
	HandlerType: (*SimMasterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamChannel",
			Handler:       _SimMaster_StreamChannel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "uegnbsim.proto",
}

// SimCliClient is the client API for SimCli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimCliClient interface {
	CreateGnb(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateUe(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DelGnb(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DelUe(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListGnb(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GnbConfigList, error)
	ListUe(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*UeConfigList, error)
	StartUeRegister(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartUeDeregister(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetAction(ctx context.Context, in *ActionProfile, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartAction(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StopAction(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type simCliClient struct {
	cc grpc.ClientConnInterface
}

func NewSimCliClient(cc grpc.ClientConnInterface) SimCliClient {
	return &simCliClient{cc}
}

func (c *simCliClient) CreateGnb(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/CreateGnb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) CreateUe(ctx context.Context, in *GnbConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/CreateUe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) DelGnb(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/DelGnb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) DelUe(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/DelUe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) ListGnb(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GnbConfigList, error) {
	out := new(GnbConfigList)
	err := c.cc.Invoke(ctx, "/api.sim_cli/ListGnb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) ListUe(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*UeConfigList, error) {
	out := new(UeConfigList)
	err := c.cc.Invoke(ctx, "/api.sim_cli/ListUe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) StartUeRegister(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/StartUeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) StartUeDeregister(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/StartUeDeregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) SetAction(ctx context.Context, in *ActionProfile, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/SetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) StartAction(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/StartAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simCliClient) StopAction(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.sim_cli/StopAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimCliServer is the server API for SimCli service.
// All implementations must embed UnimplementedSimCliServer
// for forward compatibility
type SimCliServer interface {
	CreateGnb(context.Context, *GnbConfig) (*emptypb.Empty, error)
	CreateUe(context.Context, *GnbConfig) (*emptypb.Empty, error)
	DelGnb(context.Context, *IdMessage) (*emptypb.Empty, error)
	DelUe(context.Context, *IdMessage) (*emptypb.Empty, error)
	ListGnb(context.Context, *emptypb.Empty) (*GnbConfigList, error)
	ListUe(context.Context, *IdMessage) (*UeConfigList, error)
	StartUeRegister(context.Context, *IdMessage) (*emptypb.Empty, error)
	StartUeDeregister(context.Context, *IdMessage) (*emptypb.Empty, error)
	SetAction(context.Context, *ActionProfile) (*emptypb.Empty, error)
	StartAction(context.Context, *IdMessage) (*emptypb.Empty, error)
	StopAction(context.Context, *IdMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedSimCliServer()
}

// UnimplementedSimCliServer must be embedded to have forward compatible implementations.
type UnimplementedSimCliServer struct {
}

func (UnimplementedSimCliServer) CreateGnb(context.Context, *GnbConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGnb not implemented")
}
func (UnimplementedSimCliServer) CreateUe(context.Context, *GnbConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUe not implemented")
}
func (UnimplementedSimCliServer) DelGnb(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelGnb not implemented")
}
func (UnimplementedSimCliServer) DelUe(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUe not implemented")
}
func (UnimplementedSimCliServer) ListGnb(context.Context, *emptypb.Empty) (*GnbConfigList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGnb not implemented")
}
func (UnimplementedSimCliServer) ListUe(context.Context, *IdMessage) (*UeConfigList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUe not implemented")
}
func (UnimplementedSimCliServer) StartUeRegister(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUeRegister not implemented")
}
func (UnimplementedSimCliServer) StartUeDeregister(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUeDeregister not implemented")
}
func (UnimplementedSimCliServer) SetAction(context.Context, *ActionProfile) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAction not implemented")
}
func (UnimplementedSimCliServer) StartAction(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAction not implemented")
}
func (UnimplementedSimCliServer) StopAction(context.Context, *IdMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopAction not implemented")
}
func (UnimplementedSimCliServer) mustEmbedUnimplementedSimCliServer() {}

// UnsafeSimCliServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimCliServer will
// result in compilation errors.
type UnsafeSimCliServer interface {
	mustEmbedUnimplementedSimCliServer()
}

func RegisterSimCliServer(s grpc.ServiceRegistrar, srv SimCliServer) {
	s.RegisterService(&SimCli_ServiceDesc, srv)
}

func _SimCli_CreateGnb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnbConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).CreateGnb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/CreateGnb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).CreateGnb(ctx, req.(*GnbConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_CreateUe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GnbConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).CreateUe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/CreateUe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).CreateUe(ctx, req.(*GnbConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_DelGnb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).DelGnb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/DelGnb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).DelGnb(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_DelUe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).DelUe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/DelUe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).DelUe(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_ListGnb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).ListGnb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/ListGnb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).ListGnb(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_ListUe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).ListUe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/ListUe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).ListUe(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_StartUeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).StartUeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/StartUeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).StartUeRegister(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_StartUeDeregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).StartUeDeregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/StartUeDeregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).StartUeDeregister(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_SetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).SetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/SetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).SetAction(ctx, req.(*ActionProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_StartAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).StartAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/StartAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).StartAction(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimCli_StopAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimCliServer).StopAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.sim_cli/StopAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimCliServer).StopAction(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SimCli_ServiceDesc is the grpc.ServiceDesc for SimCli service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimCli_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.sim_cli",
	HandlerType: (*SimCliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGnb",
			Handler:    _SimCli_CreateGnb_Handler,
		},
		{
			MethodName: "CreateUe",
			Handler:    _SimCli_CreateUe_Handler,
		},
		{
			MethodName: "DelGnb",
			Handler:    _SimCli_DelGnb_Handler,
		},
		{
			MethodName: "DelUe",
			Handler:    _SimCli_DelUe_Handler,
		},
		{
			MethodName: "ListGnb",
			Handler:    _SimCli_ListGnb_Handler,
		},
		{
			MethodName: "ListUe",
			Handler:    _SimCli_ListUe_Handler,
		},
		{
			MethodName: "StartUeRegister",
			Handler:    _SimCli_StartUeRegister_Handler,
		},
		{
			MethodName: "StartUeDeregister",
			Handler:    _SimCli_StartUeDeregister_Handler,
		},
		{
			MethodName: "SetAction",
			Handler:    _SimCli_SetAction_Handler,
		},
		{
			MethodName: "StartAction",
			Handler:    _SimCli_StartAction_Handler,
		},
		{
			MethodName: "StopAction",
			Handler:    _SimCli_StopAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uegnbsim.proto",
}
