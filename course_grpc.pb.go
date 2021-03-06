// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

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

// CourseServiceClient is the client API for CourseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourseServiceClient interface {
	PostCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Courses, error)
	GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Courses, error)
	GetCourse(ctx context.Context, in *Int, opts ...grpc.CallOption) (*Course, error)
	PutCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Courses, error)
	DeletCourse(ctx context.Context, in *Int, opts ...grpc.CallOption) (*Empty, error)
}

type courseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourseServiceClient(cc grpc.ClientConnInterface) CourseServiceClient {
	return &courseServiceClient{cc}
}

func (c *courseServiceClient) PostCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/main.CourseService/PostCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/main.CourseService/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) GetCourse(ctx context.Context, in *Int, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/main.CourseService/GetCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) PutCourse(ctx context.Context, in *Course, opts ...grpc.CallOption) (*Courses, error) {
	out := new(Courses)
	err := c.cc.Invoke(ctx, "/main.CourseService/PutCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseServiceClient) DeletCourse(ctx context.Context, in *Int, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/main.CourseService/DeletCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourseServiceServer is the server API for CourseService service.
// All implementations must embed UnimplementedCourseServiceServer
// for forward compatibility
type CourseServiceServer interface {
	PostCourse(context.Context, *Course) (*Courses, error)
	GetCourses(context.Context, *Empty) (*Courses, error)
	GetCourse(context.Context, *Int) (*Course, error)
	PutCourse(context.Context, *Course) (*Courses, error)
	DeletCourse(context.Context, *Int) (*Empty, error)
	mustEmbedUnimplementedCourseServiceServer()
}

// UnimplementedCourseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourseServiceServer struct {
}

func (UnimplementedCourseServiceServer) PostCourse(context.Context, *Course) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCourse not implemented")
}
func (UnimplementedCourseServiceServer) GetCourses(context.Context, *Empty) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourseServiceServer) GetCourse(context.Context, *Int) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourse not implemented")
}
func (UnimplementedCourseServiceServer) PutCourse(context.Context, *Course) (*Courses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCourse not implemented")
}
func (UnimplementedCourseServiceServer) DeletCourse(context.Context, *Int) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletCourse not implemented")
}
func (UnimplementedCourseServiceServer) mustEmbedUnimplementedCourseServiceServer() {}

// UnsafeCourseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourseServiceServer will
// result in compilation errors.
type UnsafeCourseServiceServer interface {
	mustEmbedUnimplementedCourseServiceServer()
}

func RegisterCourseServiceServer(s grpc.ServiceRegistrar, srv CourseServiceServer) {
	s.RegisterService(&CourseService_ServiceDesc, srv)
}

func _CourseService_PostCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).PostCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CourseService/PostCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).PostCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CourseService/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_GetCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).GetCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CourseService/GetCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).GetCourse(ctx, req.(*Int))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_PutCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Course)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).PutCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CourseService/PutCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).PutCourse(ctx, req.(*Course))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourseService_DeletCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Int)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourseServiceServer).DeletCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.CourseService/DeletCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourseServiceServer).DeletCourse(ctx, req.(*Int))
	}
	return interceptor(ctx, in, info, handler)
}

// CourseService_ServiceDesc is the grpc.ServiceDesc for CourseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CourseService",
	HandlerType: (*CourseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostCourse",
			Handler:    _CourseService_PostCourse_Handler,
		},
		{
			MethodName: "GetCourses",
			Handler:    _CourseService_GetCourses_Handler,
		},
		{
			MethodName: "GetCourse",
			Handler:    _CourseService_GetCourse_Handler,
		},
		{
			MethodName: "PutCourse",
			Handler:    _CourseService_PutCourse_Handler,
		},
		{
			MethodName: "DeletCourse",
			Handler:    _CourseService_DeletCourse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "course.proto",
}
