// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mapd.proto

package mapd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type SetRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{1}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetReply struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetReply) Reset()         { *m = GetReply{} }
func (m *GetReply) String() string { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()    {}
func (*GetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{2}
}

func (m *GetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetReply.Unmarshal(m, b)
}
func (m *GetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetReply.Marshal(b, m, deterministic)
}
func (m *GetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetReply.Merge(m, src)
}
func (m *GetReply) XXX_Size() int {
	return xxx_messageInfo_GetReply.Size(m)
}
func (m *GetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetReply proto.InternalMessageInfo

func (m *GetReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SetReply struct {
	Err                  bool     `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetReply) Reset()         { *m = SetReply{} }
func (m *SetReply) String() string { return proto.CompactTextString(m) }
func (*SetReply) ProtoMessage()    {}
func (*SetReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{3}
}

func (m *SetReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetReply.Unmarshal(m, b)
}
func (m *SetReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetReply.Marshal(b, m, deterministic)
}
func (m *SetReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReply.Merge(m, src)
}
func (m *SetReply) XXX_Size() int {
	return xxx_messageInfo_SetReply.Size(m)
}
func (m *SetReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetReply proto.InternalMessageInfo

func (m *SetReply) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "mapd.GetRequest")
	proto.RegisterType((*SetRequest)(nil), "mapd.SetRequest")
	proto.RegisterType((*GetReply)(nil), "mapd.GetReply")
	proto.RegisterType((*SetReply)(nil), "mapd.SetReply")
}

func init() { proto.RegisterFile("mapd.proto", fileDescriptor_47c62a75174acd61) }

var fileDescriptor_47c62a75174acd61 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2c, 0x48,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xe4, 0xb8, 0xb8, 0xdc, 0x53,
	0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x13, 0x2e, 0xae, 0x60, 0x3c, 0xf2, 0x42,
	0x22, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x10,
	0x8e, 0x92, 0x02, 0x17, 0x07, 0xd8, 0xd4, 0x82, 0x1c, 0x24, 0x15, 0x8c, 0xc8, 0x2a, 0x64, 0xb8,
	0x38, 0x82, 0x61, 0x2a, 0x04, 0xb8, 0x98, 0x53, 0x8b, 0x8a, 0xc0, 0xf2, 0x1c, 0x41, 0x20, 0xa6,
	0x51, 0x37, 0x23, 0x17, 0x8b, 0x6f, 0x62, 0x41, 0x8a, 0x90, 0x26, 0x17, 0x73, 0x71, 0x6a, 0x89,
	0x90, 0x80, 0x1e, 0xd8, 0xe1, 0x08, 0x97, 0x48, 0xf1, 0x21, 0x89, 0x14, 0xe4, 0x54, 0x2a, 0x31,
	0x08, 0xe9, 0x72, 0xb1, 0x17, 0xa7, 0x96, 0x04, 0x27, 0xa6, 0xa5, 0x12, 0xa5, 0x5c, 0x93, 0x8b,
	0x39, 0x1d, 0x61, 0xb2, 0x3b, 0x86, 0x52, 0x77, 0xb8, 0x52, 0x27, 0xb6, 0x28, 0x70, 0x58, 0x25,
	0xb1, 0x81, 0x03, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x69, 0x18, 0x20, 0x93, 0x46, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MapdClient is the client API for Mapd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MapdClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error)
	SetSafe(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
}

type mapdClient struct {
	cc *grpc.ClientConn
}

func NewMapdClient(cc *grpc.ClientConn) MapdClient {
	return &mapdClient{cc}
}

func (c *mapdClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error) {
	out := new(SetReply)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapdClient) SetSafe(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetReply, error) {
	out := new(SetReply)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/setSafe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapdClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MapdServer is the server API for Mapd service.
type MapdServer interface {
	Set(context.Context, *SetRequest) (*SetReply, error)
	SetSafe(context.Context, *SetRequest) (*SetReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
}

// UnimplementedMapdServer can be embedded to have forward compatible implementations.
type UnimplementedMapdServer struct {
}

func (*UnimplementedMapdServer) Set(ctx context.Context, req *SetRequest) (*SetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedMapdServer) SetSafe(ctx context.Context, req *SetRequest) (*SetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSafe not implemented")
}
func (*UnimplementedMapdServer) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterMapdServer(s *grpc.Server, srv MapdServer) {
	s.RegisterService(&_Mapd_serviceDesc, srv)
}

func _Mapd_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mapd_SetSafe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).SetSafe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/SetSafe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).SetSafe(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mapd_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mapd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mapd.Mapd",
	HandlerType: (*MapdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "set",
			Handler:    _Mapd_Set_Handler,
		},
		{
			MethodName: "setSafe",
			Handler:    _Mapd_SetSafe_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Mapd_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapd.proto",
}
