// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: youtube-clone/database/pbs/media.proto

package media

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	helper "youtube-clone/database/pbs/helper"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MediaServiceClient is the client API for MediaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MediaServiceClient interface {
	GetMediaByUrl(ctx context.Context, in *MediaUrl, opts ...grpc.CallOption) (*Response, error)
	SearchMedias(ctx context.Context, in *MediaReq, opts ...grpc.CallOption) (*Response, error)
	GetMedias(ctx context.Context, in *MediaReq, opts ...grpc.CallOption) (*Response, error)
	CreateMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error)
	EditMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error)
	SetMediaThumbnail(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error)
	DeleteMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error)
	AddTagToMedia(ctx context.Context, in *TagMedia, opts ...grpc.CallOption) (*Response, error)
	RemoveTagFromMedia(ctx context.Context, in *TagMedia, opts ...grpc.CallOption) (*Response, error)
	CreateLikeMedia(ctx context.Context, in *helper.LikeReq, opts ...grpc.CallOption) (*Response, error)
	DeleteLikeMedia(ctx context.Context, in *helper.LikeReq, opts ...grpc.CallOption) (*Response, error)
}

type mediaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMediaServiceClient(cc grpc.ClientConnInterface) MediaServiceClient {
	return &mediaServiceClient{cc}
}

func (c *mediaServiceClient) GetMediaByUrl(ctx context.Context, in *MediaUrl, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetMediaByUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) SearchMedias(ctx context.Context, in *MediaReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/SearchMedias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) GetMedias(ctx context.Context, in *MediaReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/GetMedias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) CreateMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/CreateMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) EditMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/EditMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) SetMediaThumbnail(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/SetMediaThumbnail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) DeleteMedia(ctx context.Context, in *EidtMediaData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/DeleteMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) AddTagToMedia(ctx context.Context, in *TagMedia, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/AddTagToMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) RemoveTagFromMedia(ctx context.Context, in *TagMedia, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/RemoveTagFromMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) CreateLikeMedia(ctx context.Context, in *helper.LikeReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/CreateLikeMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaServiceClient) DeleteLikeMedia(ctx context.Context, in *helper.LikeReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/media.MediaService/DeleteLikeMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaServiceServer is the server API for MediaService service.
// All implementations must embed UnimplementedMediaServiceServer
// for forward compatibility
type MediaServiceServer interface {
	GetMediaByUrl(context.Context, *MediaUrl) (*Response, error)
	SearchMedias(context.Context, *MediaReq) (*Response, error)
	GetMedias(context.Context, *MediaReq) (*Response, error)
	CreateMedia(context.Context, *EidtMediaData) (*Response, error)
	EditMedia(context.Context, *EidtMediaData) (*Response, error)
	SetMediaThumbnail(context.Context, *EidtMediaData) (*Response, error)
	DeleteMedia(context.Context, *EidtMediaData) (*Response, error)
	AddTagToMedia(context.Context, *TagMedia) (*Response, error)
	RemoveTagFromMedia(context.Context, *TagMedia) (*Response, error)
	CreateLikeMedia(context.Context, *helper.LikeReq) (*Response, error)
	DeleteLikeMedia(context.Context, *helper.LikeReq) (*Response, error)
	mustEmbedUnimplementedMediaServiceServer()
}

// UnimplementedMediaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMediaServiceServer struct {
}

func (UnimplementedMediaServiceServer) GetMediaByUrl(context.Context, *MediaUrl) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMediaByUrl not implemented")
}
func (UnimplementedMediaServiceServer) SearchMedias(context.Context, *MediaReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMedias not implemented")
}
func (UnimplementedMediaServiceServer) GetMedias(context.Context, *MediaReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedias not implemented")
}
func (UnimplementedMediaServiceServer) CreateMedia(context.Context, *EidtMediaData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMedia not implemented")
}
func (UnimplementedMediaServiceServer) EditMedia(context.Context, *EidtMediaData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditMedia not implemented")
}
func (UnimplementedMediaServiceServer) SetMediaThumbnail(context.Context, *EidtMediaData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMediaThumbnail not implemented")
}
func (UnimplementedMediaServiceServer) DeleteMedia(context.Context, *EidtMediaData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMedia not implemented")
}
func (UnimplementedMediaServiceServer) AddTagToMedia(context.Context, *TagMedia) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTagToMedia not implemented")
}
func (UnimplementedMediaServiceServer) RemoveTagFromMedia(context.Context, *TagMedia) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTagFromMedia not implemented")
}
func (UnimplementedMediaServiceServer) CreateLikeMedia(context.Context, *helper.LikeReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLikeMedia not implemented")
}
func (UnimplementedMediaServiceServer) DeleteLikeMedia(context.Context, *helper.LikeReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLikeMedia not implemented")
}
func (UnimplementedMediaServiceServer) mustEmbedUnimplementedMediaServiceServer() {}

// UnsafeMediaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MediaServiceServer will
// result in compilation errors.
type UnsafeMediaServiceServer interface {
	mustEmbedUnimplementedMediaServiceServer()
}

func RegisterMediaServiceServer(s grpc.ServiceRegistrar, srv MediaServiceServer) {
	s.RegisterService(&MediaService_ServiceDesc, srv)
}

func _MediaService_GetMediaByUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaUrl)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetMediaByUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetMediaByUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetMediaByUrl(ctx, req.(*MediaUrl))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_SearchMedias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).SearchMedias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/SearchMedias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).SearchMedias(ctx, req.(*MediaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_GetMedias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).GetMedias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/GetMedias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).GetMedias(ctx, req.(*MediaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_CreateMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EidtMediaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).CreateMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/CreateMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).CreateMedia(ctx, req.(*EidtMediaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_EditMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EidtMediaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).EditMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/EditMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).EditMedia(ctx, req.(*EidtMediaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_SetMediaThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EidtMediaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).SetMediaThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/SetMediaThumbnail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).SetMediaThumbnail(ctx, req.(*EidtMediaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_DeleteMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EidtMediaData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).DeleteMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/DeleteMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).DeleteMedia(ctx, req.(*EidtMediaData))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_AddTagToMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagMedia)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).AddTagToMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/AddTagToMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).AddTagToMedia(ctx, req.(*TagMedia))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_RemoveTagFromMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TagMedia)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).RemoveTagFromMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/RemoveTagFromMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).RemoveTagFromMedia(ctx, req.(*TagMedia))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_CreateLikeMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helper.LikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).CreateLikeMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/CreateLikeMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).CreateLikeMedia(ctx, req.(*helper.LikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaService_DeleteLikeMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(helper.LikeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaServiceServer).DeleteLikeMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/media.MediaService/DeleteLikeMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaServiceServer).DeleteLikeMedia(ctx, req.(*helper.LikeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MediaService_ServiceDesc is the grpc.ServiceDesc for MediaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MediaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "media.MediaService",
	HandlerType: (*MediaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMediaByUrl",
			Handler:    _MediaService_GetMediaByUrl_Handler,
		},
		{
			MethodName: "SearchMedias",
			Handler:    _MediaService_SearchMedias_Handler,
		},
		{
			MethodName: "GetMedias",
			Handler:    _MediaService_GetMedias_Handler,
		},
		{
			MethodName: "CreateMedia",
			Handler:    _MediaService_CreateMedia_Handler,
		},
		{
			MethodName: "EditMedia",
			Handler:    _MediaService_EditMedia_Handler,
		},
		{
			MethodName: "SetMediaThumbnail",
			Handler:    _MediaService_SetMediaThumbnail_Handler,
		},
		{
			MethodName: "DeleteMedia",
			Handler:    _MediaService_DeleteMedia_Handler,
		},
		{
			MethodName: "AddTagToMedia",
			Handler:    _MediaService_AddTagToMedia_Handler,
		},
		{
			MethodName: "RemoveTagFromMedia",
			Handler:    _MediaService_RemoveTagFromMedia_Handler,
		},
		{
			MethodName: "CreateLikeMedia",
			Handler:    _MediaService_CreateLikeMedia_Handler,
		},
		{
			MethodName: "DeleteLikeMedia",
			Handler:    _MediaService_DeleteLikeMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "youtube-clone/database/pbs/media.proto",
}
