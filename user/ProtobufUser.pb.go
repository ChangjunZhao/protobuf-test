// Code generated by protoc-gen-go.
// source: ProtobufUser.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	ProtobufUser.proto

It has these top-level messages:
	ProtobufUser
	AddPhoneToUserRequest
	AddPhoneToUserResponse
*/
package user

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

type PhoneType int32

const (
	PhoneType_HOME  PhoneType = 0
	PhoneType_WORK  PhoneType = 1
	PhoneType_OTHER PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "HOME",
	1: "WORK",
	2: "OTHER",
}
var PhoneType_value = map[string]int32{
	"HOME":  0,
	"WORK":  1,
	"OTHER": 2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProtobufUser struct {
	Id     int32                 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Phones []*ProtobufUser_Phone `protobuf:"bytes,3,rep,name=phones" json:"phones,omitempty"`
}

func (m *ProtobufUser) Reset()                    { *m = ProtobufUser{} }
func (m *ProtobufUser) String() string            { return proto.CompactTextString(m) }
func (*ProtobufUser) ProtoMessage()               {}
func (*ProtobufUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtobufUser) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProtobufUser) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProtobufUser) GetPhones() []*ProtobufUser_Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

type ProtobufUser_Phone struct {
	PhoneType   PhoneType `protobuf:"varint,1,opt,name=phoneType,enum=PhoneType" json:"phoneType,omitempty"`
	PhoneNumber string    `protobuf:"bytes,2,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
}

func (m *ProtobufUser_Phone) Reset()                    { *m = ProtobufUser_Phone{} }
func (m *ProtobufUser_Phone) String() string            { return proto.CompactTextString(m) }
func (*ProtobufUser_Phone) ProtoMessage()               {}
func (*ProtobufUser_Phone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ProtobufUser_Phone) GetPhoneType() PhoneType {
	if m != nil {
		return m.PhoneType
	}
	return PhoneType_HOME
}

func (m *ProtobufUser_Phone) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type AddPhoneToUserRequest struct {
	Uid         int32     `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	PhoneType   PhoneType `protobuf:"varint,2,opt,name=phoneType,enum=PhoneType" json:"phoneType,omitempty"`
	PhoneNumber string    `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
}

func (m *AddPhoneToUserRequest) Reset()                    { *m = AddPhoneToUserRequest{} }
func (m *AddPhoneToUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPhoneToUserRequest) ProtoMessage()               {}
func (*AddPhoneToUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddPhoneToUserRequest) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AddPhoneToUserRequest) GetPhoneType() PhoneType {
	if m != nil {
		return m.PhoneType
	}
	return PhoneType_HOME
}

func (m *AddPhoneToUserRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type AddPhoneToUserResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *AddPhoneToUserResponse) Reset()                    { *m = AddPhoneToUserResponse{} }
func (m *AddPhoneToUserResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPhoneToUserResponse) ProtoMessage()               {}
func (*AddPhoneToUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddPhoneToUserResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*ProtobufUser)(nil), "ProtobufUser")
	proto.RegisterType((*ProtobufUser_Phone)(nil), "ProtobufUser.Phone")
	proto.RegisterType((*AddPhoneToUserRequest)(nil), "AddPhoneToUserRequest")
	proto.RegisterType((*AddPhoneToUserResponse)(nil), "AddPhoneToUserResponse")
	proto.RegisterEnum("PhoneType", PhoneType_name, PhoneType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PhoneService service

type PhoneServiceClient interface {
	AddPhoneToUser(ctx context.Context, in *AddPhoneToUserRequest, opts ...grpc.CallOption) (*AddPhoneToUserResponse, error)
}

type phoneServiceClient struct {
	cc *grpc.ClientConn
}

func NewPhoneServiceClient(cc *grpc.ClientConn) PhoneServiceClient {
	return &phoneServiceClient{cc}
}

func (c *phoneServiceClient) AddPhoneToUser(ctx context.Context, in *AddPhoneToUserRequest, opts ...grpc.CallOption) (*AddPhoneToUserResponse, error) {
	out := new(AddPhoneToUserResponse)
	err := grpc.Invoke(ctx, "/PhoneService/addPhoneToUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhoneService service

type PhoneServiceServer interface {
	AddPhoneToUser(context.Context, *AddPhoneToUserRequest) (*AddPhoneToUserResponse, error)
}

func RegisterPhoneServiceServer(s *grpc.Server, srv PhoneServiceServer) {
	s.RegisterService(&_PhoneService_serviceDesc, srv)
}

func _PhoneService_AddPhoneToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhoneToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneServiceServer).AddPhoneToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PhoneService/AddPhoneToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneServiceServer).AddPhoneToUser(ctx, req.(*AddPhoneToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PhoneService",
	HandlerType: (*PhoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addPhoneToUser",
			Handler:    _PhoneService_AddPhoneToUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ProtobufUser.proto",
}

func init() { proto.RegisterFile("ProtobufUser.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x91, 0xd1, 0x6a, 0xf2, 0x40,
	0x10, 0x85, 0xff, 0x24, 0x2a, 0x66, 0xfc, 0x91, 0x30, 0xa5, 0x2a, 0xde, 0x54, 0xbc, 0x12, 0x0b,
	0xa1, 0xd8, 0x27, 0xb0, 0x20, 0x08, 0xa5, 0x8d, 0x5d, 0x2d, 0x85, 0xde, 0x99, 0x64, 0x4a, 0x85,
	0x26, 0x9b, 0xee, 0x66, 0x85, 0xbe, 0x57, 0x1f, 0xb0, 0xd9, 0x4d, 0xb0, 0xa6, 0xf5, 0xa6, 0x77,
	0xe7, 0x9c, 0x39, 0xc9, 0x37, 0xcb, 0x00, 0xae, 0x04, 0xcf, 0x79, 0xa8, 0x5e, 0x1e, 0x25, 0x09,
	0x3f, 0xd3, 0x66, 0xfc, 0x69, 0xc1, 0xff, 0xe3, 0x18, 0xbb, 0x60, 0xef, 0xe2, 0x81, 0x35, 0xb2,
	0x26, 0x4d, 0x56, 0x28, 0x44, 0x68, 0xa4, 0xdb, 0x84, 0x06, 0x76, 0x91, 0xb8, 0xcc, 0x68, 0xbc,
	0x84, 0x56, 0xf6, 0xca, 0x53, 0x92, 0x03, 0x67, 0xe4, 0x4c, 0x3a, 0xb3, 0x33, 0xbf, 0xf6, 0xe7,
	0x95, 0x9e, 0xb1, 0xaa, 0x32, 0x5c, 0x43, 0xd3, 0x04, 0x38, 0x01, 0xd7, 0x44, 0x9b, 0x8f, 0x8c,
	0x0c, 0xa0, 0x3b, 0x83, 0xb2, 0xab, 0x13, 0xf6, 0x3d, 0xc4, 0x11, 0x74, 0x8c, 0xb9, 0x57, 0x49,
	0x48, 0xa2, 0x42, 0x1f, 0x47, 0x63, 0x05, 0xe7, 0xf3, 0x38, 0x2e, 0x3f, 0xe6, 0x1a, 0xca, 0xe8,
	0x5d, 0x91, 0xcc, 0xd1, 0x03, 0x47, 0x1d, 0xf6, 0xd7, 0xb2, 0x8e, 0xb5, 0xff, 0x80, 0x75, 0x7e,
	0x63, 0xaf, 0xa0, 0xf7, 0x13, 0x2b, 0x33, 0x9e, 0x4a, 0xc2, 0x1e, 0xb4, 0x04, 0x49, 0xf5, 0x96,
	0x1b, 0x74, 0x9b, 0x55, 0x6e, 0x3a, 0x05, 0xf7, 0xc0, 0xc2, 0x36, 0x34, 0x96, 0xc1, 0xdd, 0xc2,
	0xfb, 0xa7, 0xd5, 0x53, 0xc0, 0x6e, 0x3d, 0x0b, 0x5d, 0x68, 0x06, 0x9b, 0xe5, 0x82, 0x79, 0xf6,
	0xec, 0xa1, 0x38, 0x85, 0xee, 0xae, 0x49, 0xec, 0x77, 0x11, 0xe1, 0x1c, 0xba, 0xdb, 0x1a, 0x0d,
	0x7b, 0xfe, 0xc9, 0x57, 0x0f, 0xfb, 0xfe, 0xe9, 0xb5, 0x6e, 0x2e, 0xa0, 0x1f, 0xf1, 0xc4, 0xdf,
	0x53, 0xaa, 0x24, 0x57, 0x22, 0xa2, 0xf2, 0xec, 0xc5, 0xa5, 0x9e, 0x1b, 0xaa, 0x28, 0x86, 0x2d,
	0xe3, 0xaf, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x49, 0x90, 0xfc, 0xac, 0x1c, 0x02, 0x00, 0x00,
}
