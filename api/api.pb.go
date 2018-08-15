// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	VersionRequest
	VersionResponse
	ParseRequest
	ParseResponse
*/
package api

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

type ParseRequest_ParseKind int32

const (
	ParseRequest_INVALID ParseRequest_ParseKind = 0
	ParseRequest_LANG    ParseRequest_ParseKind = 1
)

var ParseRequest_ParseKind_name = map[int32]string{
	0: "INVALID",
	1: "LANG",
}
var ParseRequest_ParseKind_value = map[string]int32{
	"INVALID": 0,
	"LANG":    1,
}

func (x ParseRequest_ParseKind) String() string {
	return proto.EnumName(ParseRequest_ParseKind_name, int32(x))
}
func (ParseRequest_ParseKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type VersionResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ParseRequest struct {
	Kind    ParseRequest_ParseKind `protobuf:"varint,1,opt,name=kind,enum=ParseRequest_ParseKind" json:"kind,omitempty"`
	Name    string                 `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Content []byte                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *ParseRequest) Reset()                    { *m = ParseRequest{} }
func (m *ParseRequest) String() string            { return proto.CompactTextString(m) }
func (*ParseRequest) ProtoMessage()               {}
func (*ParseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ParseRequest) GetKind() ParseRequest_ParseKind {
	if m != nil {
		return m.Kind
	}
	return ParseRequest_INVALID
}

func (m *ParseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ParseRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ParseResponse struct {
	Lang string `protobuf:"bytes,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *ParseResponse) Reset()                    { *m = ParseResponse{} }
func (m *ParseResponse) String() string            { return proto.CompactTextString(m) }
func (*ParseResponse) ProtoMessage()               {}
func (*ParseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ParseResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "VersionResponse")
	proto.RegisterType((*ParseRequest)(nil), "ParseRequest")
	proto.RegisterType((*ParseResponse)(nil), "ParseResponse")
	proto.RegisterEnum("ParseRequest_ParseKind", ParseRequest_ParseKind_name, ParseRequest_ParseKind_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/Engine/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineClient) Parse(ctx context.Context, in *ParseRequest, opts ...grpc.CallOption) (*ParseResponse, error) {
	out := new(ParseResponse)
	err := grpc.Invoke(ctx, "/Engine/Parse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	Parse(context.Context, *ParseRequest) (*ParseResponse, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Engine_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Engine/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Parse(ctx, req.(*ParseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Engine_Version_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _Engine_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xee, 0x6a, 0xa5, 0x76, 0x84, 0xd2, 0xcc, 0xc5, 0x86, 0x13, 0x19, 0x2f, 0x4d, 0x48, 0xf6,
	0x80, 0x4f, 0x40, 0xa2, 0x31, 0x44, 0x42, 0x4c, 0x0f, 0xdc, 0x8b, 0x4c, 0xc8, 0x46, 0x9d, 0xad,
	0xdd, 0xea, 0x5b, 0xf8, 0xce, 0xc6, 0x65, 0x21, 0x96, 0xdb, 0x7c, 0xb3, 0xb3, 0xdf, 0x1f, 0xa4,
	0x75, 0x63, 0x74, 0xd3, 0xda, 0xce, 0x52, 0x0e, 0xd9, 0x86, 0x5b, 0x67, 0xac, 0x54, 0xfc, 0xf9,
	0xc5, 0xae, 0xa3, 0x19, 0x8c, 0x4f, 0x1b, 0xd7, 0x58, 0x71, 0x8c, 0x05, 0x24, 0xdf, 0x87, 0x55,
	0xa1, 0xa6, 0xaa, 0x4c, 0xab, 0x23, 0xa4, 0x1f, 0x05, 0xc3, 0x97, 0xba, 0x75, 0x1c, 0x7e, 0xe3,
	0x0c, 0xe2, 0x37, 0x23, 0x3b, 0x7f, 0x97, 0xcd, 0x6f, 0xf5, 0xff, 0xc7, 0x03, 0x78, 0x36, 0xb2,
	0xab, 0xfc, 0x11, 0x22, 0xc4, 0x52, 0x7f, 0x70, 0x71, 0xe1, 0x49, 0xfd, 0xfc, 0xa7, 0xf5, 0x6a,
	0xa5, 0x63, 0xe9, 0x8a, 0xcb, 0xa9, 0x2a, 0x87, 0xd5, 0x11, 0x12, 0x41, 0x7a, 0x22, 0xc0, 0x1b,
	0x48, 0x96, 0xeb, 0xcd, 0x62, 0xb5, 0x7c, 0xc8, 0x23, 0xbc, 0x86, 0x78, 0xb5, 0x58, 0x3f, 0xe5,
	0x8a, 0xee, 0x60, 0x14, 0x14, 0x83, 0x75, 0x84, 0xf8, 0xbd, 0x96, 0x7d, 0xf0, 0xed, 0xe7, 0xf9,
	0x16, 0x06, 0x8f, 0xb2, 0x37, 0xc2, 0xa8, 0x21, 0x09, 0x59, 0x71, 0xac, 0xfb, 0x3d, 0x4c, 0x72,
	0x7d, 0x56, 0x03, 0x45, 0x58, 0xc2, 0x95, 0xa7, 0xc7, 0x51, 0x2f, 0xd8, 0x24, 0xd3, 0x3d, 0x55,
	0x8a, 0xb6, 0x03, 0x5f, 0xef, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x7d, 0x22, 0xe9,
	0x6b, 0x01, 0x00, 0x00,
}
