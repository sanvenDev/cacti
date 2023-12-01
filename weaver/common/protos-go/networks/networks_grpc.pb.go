// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: networks/networks.proto

package networks

import (
	context "context"
	common "github.com/hyperledger/cacti/weaver/common/protos-go/v2/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Network_RequestState_FullMethodName              = "/networks.networks.Network/RequestState"
	Network_GetState_FullMethodName                  = "/networks.networks.Network/GetState"
	Network_RequestDatabase_FullMethodName           = "/networks.networks.Network/RequestDatabase"
	Network_SubscribeEvent_FullMethodName            = "/networks.networks.Network/SubscribeEvent"
	Network_GetEventSubscriptionState_FullMethodName = "/networks.networks.Network/GetEventSubscriptionState"
	Network_UnsubscribeEvent_FullMethodName          = "/networks.networks.Network/UnsubscribeEvent"
	Network_GetEventStates_FullMethodName            = "/networks.networks.Network/GetEventStates"
)

// NetworkClient is the client API for Network service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkClient interface {
	// Data Sharing endpoints
	// endpoint for a network to request remote relay state via local relay
	RequestState(ctx context.Context, in *NetworkQuery, opts ...grpc.CallOption) (*common.Ack, error)
	// This rpc endpoint is for polling the local relay for request state.
	GetState(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.RequestState, error)
	// NOTE: This rpc is just for debugging.
	RequestDatabase(ctx context.Context, in *DbName, opts ...grpc.CallOption) (*RelayDatabase, error)
	// Event endpoints
	// endpoint for a client to subscribe to event via local relay initiating subscription flow.
	SubscribeEvent(ctx context.Context, in *NetworkEventSubscription, opts ...grpc.CallOption) (*common.Ack, error)
	// This rpc endpoint is for polling the local relay for subscription state.
	GetEventSubscriptionState(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.EventSubscriptionState, error)
	// endpoint for a client to subscribe to event via local relay initiating subscription flow.
	UnsubscribeEvent(ctx context.Context, in *NetworkEventUnsubscription, opts ...grpc.CallOption) (*common.Ack, error)
	// endpoint for a client to fetch received events.
	// Note: events are marked as deleted from relay database as soon as client fetches them.
	GetEventStates(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.EventStates, error)
}

type networkClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkClient(cc grpc.ClientConnInterface) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) RequestState(ctx context.Context, in *NetworkQuery, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, Network_RequestState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetState(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.RequestState, error) {
	out := new(common.RequestState)
	err := c.cc.Invoke(ctx, Network_GetState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) RequestDatabase(ctx context.Context, in *DbName, opts ...grpc.CallOption) (*RelayDatabase, error) {
	out := new(RelayDatabase)
	err := c.cc.Invoke(ctx, Network_RequestDatabase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) SubscribeEvent(ctx context.Context, in *NetworkEventSubscription, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, Network_SubscribeEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetEventSubscriptionState(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.EventSubscriptionState, error) {
	out := new(common.EventSubscriptionState)
	err := c.cc.Invoke(ctx, Network_GetEventSubscriptionState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) UnsubscribeEvent(ctx context.Context, in *NetworkEventUnsubscription, opts ...grpc.CallOption) (*common.Ack, error) {
	out := new(common.Ack)
	err := c.cc.Invoke(ctx, Network_UnsubscribeEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetEventStates(ctx context.Context, in *GetStateMessage, opts ...grpc.CallOption) (*common.EventStates, error) {
	out := new(common.EventStates)
	err := c.cc.Invoke(ctx, Network_GetEventStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServer is the server API for Network service.
// All implementations must embed UnimplementedNetworkServer
// for forward compatibility
type NetworkServer interface {
	// Data Sharing endpoints
	// endpoint for a network to request remote relay state via local relay
	RequestState(context.Context, *NetworkQuery) (*common.Ack, error)
	// This rpc endpoint is for polling the local relay for request state.
	GetState(context.Context, *GetStateMessage) (*common.RequestState, error)
	// NOTE: This rpc is just for debugging.
	RequestDatabase(context.Context, *DbName) (*RelayDatabase, error)
	// Event endpoints
	// endpoint for a client to subscribe to event via local relay initiating subscription flow.
	SubscribeEvent(context.Context, *NetworkEventSubscription) (*common.Ack, error)
	// This rpc endpoint is for polling the local relay for subscription state.
	GetEventSubscriptionState(context.Context, *GetStateMessage) (*common.EventSubscriptionState, error)
	// endpoint for a client to subscribe to event via local relay initiating subscription flow.
	UnsubscribeEvent(context.Context, *NetworkEventUnsubscription) (*common.Ack, error)
	// endpoint for a client to fetch received events.
	// Note: events are marked as deleted from relay database as soon as client fetches them.
	GetEventStates(context.Context, *GetStateMessage) (*common.EventStates, error)
	mustEmbedUnimplementedNetworkServer()
}

// UnimplementedNetworkServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkServer struct {
}

func (UnimplementedNetworkServer) RequestState(context.Context, *NetworkQuery) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestState not implemented")
}
func (UnimplementedNetworkServer) GetState(context.Context, *GetStateMessage) (*common.RequestState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}
func (UnimplementedNetworkServer) RequestDatabase(context.Context, *DbName) (*RelayDatabase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDatabase not implemented")
}
func (UnimplementedNetworkServer) SubscribeEvent(context.Context, *NetworkEventSubscription) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeEvent not implemented")
}
func (UnimplementedNetworkServer) GetEventSubscriptionState(context.Context, *GetStateMessage) (*common.EventSubscriptionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventSubscriptionState not implemented")
}
func (UnimplementedNetworkServer) UnsubscribeEvent(context.Context, *NetworkEventUnsubscription) (*common.Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribeEvent not implemented")
}
func (UnimplementedNetworkServer) GetEventStates(context.Context, *GetStateMessage) (*common.EventStates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStates not implemented")
}
func (UnimplementedNetworkServer) mustEmbedUnimplementedNetworkServer() {}

// UnsafeNetworkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServer will
// result in compilation errors.
type UnsafeNetworkServer interface {
	mustEmbedUnimplementedNetworkServer()
}

func RegisterNetworkServer(s grpc.ServiceRegistrar, srv NetworkServer) {
	s.RegisterService(&Network_ServiceDesc, srv)
}

func _Network_RequestState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).RequestState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_RequestState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).RequestState(ctx, req.(*NetworkQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_GetState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetState(ctx, req.(*GetStateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_RequestDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DbName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).RequestDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_RequestDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).RequestDatabase(ctx, req.(*DbName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_SubscribeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkEventSubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).SubscribeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_SubscribeEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).SubscribeEvent(ctx, req.(*NetworkEventSubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetEventSubscriptionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetEventSubscriptionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_GetEventSubscriptionState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetEventSubscriptionState(ctx, req.(*GetStateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_UnsubscribeEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkEventUnsubscription)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).UnsubscribeEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_UnsubscribeEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).UnsubscribeEvent(ctx, req.(*NetworkEventUnsubscription))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetEventStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetEventStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Network_GetEventStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetEventStates(ctx, req.(*GetStateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Network_ServiceDesc is the grpc.ServiceDesc for Network service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Network_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networks.networks.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestState",
			Handler:    _Network_RequestState_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Network_GetState_Handler,
		},
		{
			MethodName: "RequestDatabase",
			Handler:    _Network_RequestDatabase_Handler,
		},
		{
			MethodName: "SubscribeEvent",
			Handler:    _Network_SubscribeEvent_Handler,
		},
		{
			MethodName: "GetEventSubscriptionState",
			Handler:    _Network_GetEventSubscriptionState_Handler,
		},
		{
			MethodName: "UnsubscribeEvent",
			Handler:    _Network_UnsubscribeEvent_Handler,
		},
		{
			MethodName: "GetEventStates",
			Handler:    _Network_GetEventStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networks/networks.proto",
}
