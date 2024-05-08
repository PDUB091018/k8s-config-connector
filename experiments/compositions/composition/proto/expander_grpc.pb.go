// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The definition of a Expander service.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/expander.proto

package expander

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

const (
	Expander_Validate_FullMethodName = "/expander_grpc.Expander/Validate"
	Expander_Evaluate_FullMethodName = "/expander_grpc.Expander/Evaluate"
)

// ExpanderClient is the client API for Expander service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpanderClient interface {
	// Verify the expander config/template
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResult, error)
	// Evaluate the expander config in context of inputs and return manifests
	Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResult, error)
}

type expanderClient struct {
	cc grpc.ClientConnInterface
}

func NewExpanderClient(cc grpc.ClientConnInterface) ExpanderClient {
	return &expanderClient{cc}
}

func (c *expanderClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResult, error) {
	out := new(ValidateResult)
	err := c.cc.Invoke(ctx, Expander_Validate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expanderClient) Evaluate(ctx context.Context, in *EvaluateRequest, opts ...grpc.CallOption) (*EvaluateResult, error) {
	out := new(EvaluateResult)
	err := c.cc.Invoke(ctx, Expander_Evaluate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpanderServer is the server API for Expander service.
// All implementations must embed UnimplementedExpanderServer
// for forward compatibility
type ExpanderServer interface {
	// Verify the expander config/template
	Validate(context.Context, *ValidateRequest) (*ValidateResult, error)
	// Evaluate the expander config in context of inputs and return manifests
	Evaluate(context.Context, *EvaluateRequest) (*EvaluateResult, error)
	mustEmbedUnimplementedExpanderServer()
}

// UnimplementedExpanderServer must be embedded to have forward compatible implementations.
type UnimplementedExpanderServer struct {
}

func (UnimplementedExpanderServer) Validate(context.Context, *ValidateRequest) (*ValidateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedExpanderServer) Evaluate(context.Context, *EvaluateRequest) (*EvaluateResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evaluate not implemented")
}
func (UnimplementedExpanderServer) mustEmbedUnimplementedExpanderServer() {}

// UnsafeExpanderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpanderServer will
// result in compilation errors.
type UnsafeExpanderServer interface {
	mustEmbedUnimplementedExpanderServer()
}

func RegisterExpanderServer(s grpc.ServiceRegistrar, srv ExpanderServer) {
	s.RegisterService(&Expander_ServiceDesc, srv)
}

func _Expander_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpanderServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Expander_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpanderServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Expander_Evaluate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvaluateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpanderServer).Evaluate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Expander_Evaluate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpanderServer).Evaluate(ctx, req.(*EvaluateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Expander_ServiceDesc is the grpc.ServiceDesc for Expander service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Expander_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "expander_grpc.Expander",
	HandlerType: (*ExpanderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _Expander_Validate_Handler,
		},
		{
			MethodName: "Evaluate",
			Handler:    _Expander_Evaluate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/expander.proto",
}
