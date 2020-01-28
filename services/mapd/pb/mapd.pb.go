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

type Entry struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	LogStore             int32    `protobuf:"varint,2,opt,name=logStore,proto3" json:"logStore,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	SumOffset            int64    `protobuf:"varint,4,opt,name=sumOffset,proto3" json:"sumOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{4}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Entry) GetLogStore() int32 {
	if m != nil {
		return m.LogStore
	}
	return 0
}

func (m *Entry) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Entry) GetSumOffset() int64 {
	if m != nil {
		return m.SumOffset
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{5}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type ProbeRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Deph                 uint64   `protobuf:"varint,2,opt,name=deph,proto3" json:"deph,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeRequest) Reset()         { *m = ProbeRequest{} }
func (m *ProbeRequest) String() string { return proto.CompactTextString(m) }
func (*ProbeRequest) ProtoMessage()    {}
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{6}
}

func (m *ProbeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeRequest.Unmarshal(m, b)
}
func (m *ProbeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeRequest.Marshal(b, m, deterministic)
}
func (m *ProbeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeRequest.Merge(m, src)
}
func (m *ProbeRequest) XXX_Size() int {
	return xxx_messageInfo_ProbeRequest.Size(m)
}
func (m *ProbeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeRequest proto.InternalMessageInfo

func (m *ProbeRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProbeRequest) GetDeph() uint64 {
	if m != nil {
		return m.Deph
	}
	return 0
}

type ProbeReply struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Value                []byte   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProbeReply) Reset()         { *m = ProbeReply{} }
func (m *ProbeReply) String() string { return proto.CompactTextString(m) }
func (*ProbeReply) ProtoMessage()    {}
func (*ProbeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_47c62a75174acd61, []int{7}
}

func (m *ProbeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProbeReply.Unmarshal(m, b)
}
func (m *ProbeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProbeReply.Marshal(b, m, deterministic)
}
func (m *ProbeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProbeReply.Merge(m, src)
}
func (m *ProbeReply) XXX_Size() int {
	return xxx_messageInfo_ProbeReply.Size(m)
}
func (m *ProbeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProbeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProbeReply proto.InternalMessageInfo

func (m *ProbeReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProbeReply) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ProbeReply) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ProbeReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "mapd.GetRequest")
	proto.RegisterType((*SetRequest)(nil), "mapd.SetRequest")
	proto.RegisterType((*GetReply)(nil), "mapd.GetReply")
	proto.RegisterType((*SetReply)(nil), "mapd.SetReply")
	proto.RegisterType((*Entry)(nil), "mapd.Entry")
	proto.RegisterType((*Void)(nil), "mapd.Void")
	proto.RegisterType((*ProbeRequest)(nil), "mapd.ProbeRequest")
	proto.RegisterType((*ProbeReply)(nil), "mapd.ProbeReply")
}

func init() { proto.RegisterFile("mapd.proto", fileDescriptor_47c62a75174acd61) }

var fileDescriptor_47c62a75174acd61 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0x6e, 0x9a, 0x6d, 0x5e, 0x3a, 0x2f, 0x3c, 0xc2, 0xf2, 0x90, 0x12, 0x8a, 0x84, 0x3d, 0x48,
	0x7b, 0xb0, 0x07, 0x2d, 0xfe, 0x00, 0x51, 0x7a, 0x12, 0x65, 0x03, 0x1e, 0x3c, 0x99, 0x9a, 0x49,
	0x5b, 0x9a, 0x36, 0x71, 0xb3, 0x15, 0xfa, 0x5b, 0xfc, 0xb3, 0x92, 0xd9, 0x6a, 0x52, 0xad, 0xe2,
	0x6d, 0x66, 0xbe, 0xef, 0x9b, 0x99, 0xfd, 0x66, 0x01, 0x56, 0x71, 0x91, 0x8c, 0x0a, 0x95, 0xeb,
	0x9c, 0xb3, 0x2a, 0x16, 0xc7, 0x00, 0x13, 0xd4, 0x12, 0x9f, 0x37, 0x58, 0x6a, 0xee, 0x83, 0xbd,
	0xc4, 0x6d, 0xcf, 0x0a, 0xad, 0x41, 0x57, 0x56, 0xa1, 0x18, 0x03, 0x44, 0x3f, 0xe0, 0xfc, 0x3f,
	0x74, 0x5e, 0xe2, 0x6c, 0x83, 0xbd, 0x76, 0x68, 0x0d, 0x3c, 0x69, 0x12, 0x11, 0x82, 0x4b, 0x5d,
	0x8b, 0xac, 0xc1, 0xb0, 0x9a, 0x8c, 0x3e, 0xb8, 0xd1, 0x3b, 0xc3, 0x07, 0x1b, 0x95, 0x22, 0xdc,
	0x95, 0x55, 0x28, 0x96, 0xd0, 0xb9, 0x5e, 0x6b, 0xb5, 0x3d, 0x30, 0x30, 0x00, 0x37, 0xcb, 0x67,
	0x91, 0xce, 0x95, 0x99, 0xd9, 0x91, 0x1f, 0x39, 0x3f, 0x02, 0x27, 0x4f, 0xd3, 0x12, 0x75, 0xcf,
	0x0e, 0xad, 0x81, 0x2d, 0x77, 0x19, 0xef, 0x43, 0xb7, 0xdc, 0xac, 0x6e, 0x0d, 0xc4, 0x08, 0xaa,
	0x0b, 0xc2, 0x01, 0x76, 0x9f, 0x2f, 0x12, 0x31, 0x06, 0xef, 0x4e, 0xe5, 0x53, 0xfc, 0xfe, 0xb1,
	0x1c, 0x58, 0x82, 0xc5, 0x9c, 0xe6, 0x32, 0x49, 0xb1, 0x78, 0x04, 0xd8, 0xa9, 0x76, 0x4f, 0xf9,
	0xa4, 0xa9, 0x77, 0x6a, 0xef, 0xed, 0xc4, 0x81, 0xcd, 0xe3, 0x72, 0x4e, 0x9b, 0x7a, 0x92, 0xe2,
	0xda, 0x2a, 0xd6, 0xb0, 0xea, 0xec, 0xb5, 0x0d, 0xec, 0x26, 0x2e, 0x12, 0x3e, 0x04, 0xbb, 0x52,
	0xfa, 0x23, 0xba, 0x62, 0x7d, 0x96, 0xe0, 0x5f, 0xa3, 0x52, 0x64, 0x5b, 0xd1, 0xe2, 0xa7, 0xf0,
	0xa7, 0x44, 0x1d, 0xc5, 0x29, 0xfe, 0x8a, 0x3e, 0x04, 0x7b, 0x56, 0x77, 0x9e, 0x7c, 0xa1, 0x4e,
	0x6a, 0xea, 0x09, 0x74, 0x15, 0x16, 0xd9, 0xe2, 0x29, 0xd6, 0xc8, 0xff, 0x1a, 0x98, 0x6e, 0x15,
	0x80, 0x49, 0xc8, 0xcb, 0x16, 0xbf, 0x00, 0x6f, 0x86, 0xfa, 0x0a, 0xd3, 0x05, 0xd9, 0xc3, 0xb9,
	0x41, 0x9b, 0x0e, 0x07, 0xfe, 0x5e, 0xcd, 0xf4, 0x1f, 0x91, 0x4e, 0xc6, 0xeb, 0xc4, 0xe8, 0x1a,
	0x5d, 0x0f, 0xf1, 0x2f, 0x9d, 0x07, 0xfa, 0xc8, 0x53, 0x87, 0x7e, 0xf5, 0xf9, 0x5b, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x26, 0x4f, 0x0b, 0xf3, 0xe3, 0x02, 0x00, 0x00,
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
	Replicate(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Void, error)
	GetDefiProbe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeReply, error)
	GetRandProbe(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ProbeReply, error)
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

func (c *mapdClient) Replicate(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/replicate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapdClient) GetDefiProbe(ctx context.Context, in *ProbeRequest, opts ...grpc.CallOption) (*ProbeReply, error) {
	out := new(ProbeReply)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/getDefiProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mapdClient) GetRandProbe(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ProbeReply, error) {
	out := new(ProbeReply)
	err := c.cc.Invoke(ctx, "/mapd.Mapd/getRandProbe", in, out, opts...)
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
	Replicate(context.Context, *Entry) (*Void, error)
	GetDefiProbe(context.Context, *ProbeRequest) (*ProbeReply, error)
	GetRandProbe(context.Context, *Void) (*ProbeReply, error)
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
func (*UnimplementedMapdServer) Replicate(ctx context.Context, req *Entry) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Replicate not implemented")
}
func (*UnimplementedMapdServer) GetDefiProbe(ctx context.Context, req *ProbeRequest) (*ProbeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefiProbe not implemented")
}
func (*UnimplementedMapdServer) GetRandProbe(ctx context.Context, req *Void) (*ProbeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandProbe not implemented")
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

func _Mapd_Replicate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).Replicate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/Replicate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).Replicate(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mapd_GetDefiProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).GetDefiProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/GetDefiProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).GetDefiProbe(ctx, req.(*ProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mapd_GetRandProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MapdServer).GetRandProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapd.Mapd/GetRandProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MapdServer).GetRandProbe(ctx, req.(*Void))
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
		{
			MethodName: "replicate",
			Handler:    _Mapd_Replicate_Handler,
		},
		{
			MethodName: "getDefiProbe",
			Handler:    _Mapd_GetDefiProbe_Handler,
		},
		{
			MethodName: "getRandProbe",
			Handler:    _Mapd_GetRandProbe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapd.proto",
}
