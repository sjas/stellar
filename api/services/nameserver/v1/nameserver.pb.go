// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto

/*
Package nameserver is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	LookupRequest
	Record
	LookupResponse
	ListRequest
	ListResponse
	CreateRequest
	DeleteRequest
*/
package nameserver

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RecordType int32

const (
	RecordType_UNKNOWN RecordType = 0
	RecordType_A       RecordType = 1
	RecordType_CNAME   RecordType = 2
	RecordType_MX      RecordType = 3
	RecordType_TXT     RecordType = 4
	RecordType_SRV     RecordType = 5
)

var RecordType_name = map[int32]string{
	0: "UNKNOWN",
	1: "A",
	2: "CNAME",
	3: "MX",
	4: "TXT",
	5: "SRV",
}
var RecordType_value = map[string]int32{
	"UNKNOWN": 0,
	"A":       1,
	"CNAME":   2,
	"MX":      3,
	"TXT":     4,
	"SRV":     5,
}

func (x RecordType) String() string {
	return proto.EnumName(RecordType_name, int32(x))
}
func (RecordType) EnumDescriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{0} }

type InfoRequest struct {
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{0} }

type InfoResponse struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{1} }

func (m *InfoResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type LookupRequest struct {
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (m *LookupRequest) Reset()                    { *m = LookupRequest{} }
func (m *LookupRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()               {}
func (*LookupRequest) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{2} }

func (m *LookupRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Record struct {
	Type    RecordType            `protobuf:"varint,1,opt,name=type,proto3,enum=stellar.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name    string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value   string                `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Options *google_protobuf2.Any `protobuf:"bytes,4,opt,name=options" json:"options,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{3} }

func (m *Record) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *Record) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Record) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Record) GetOptions() *google_protobuf2.Any {
	if m != nil {
		return m.Options
	}
	return nil
}

type LookupResponse struct {
	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
}

func (m *LookupResponse) Reset()                    { *m = LookupResponse{} }
func (m *LookupResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()               {}
func (*LookupResponse) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{4} }

func (m *LookupResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LookupResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{5} }

type ListResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{6} }

func (m *ListResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type CreateRequest struct {
	Name    string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Records []*Record `protobuf:"bytes,2,rep,name=records" json:"records,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{7} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type DeleteRequest struct {
	Type RecordType `protobuf:"varint,1,opt,name=type,proto3,enum=stellar.services.nameserver.v1.RecordType" json:"type,omitempty"`
	Name string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorNameserver, []int{8} }

func (m *DeleteRequest) GetType() RecordType {
	if m != nil {
		return m.Type
	}
	return RecordType_UNKNOWN
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "stellar.services.nameserver.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "stellar.services.nameserver.v1.InfoResponse")
	proto.RegisterType((*LookupRequest)(nil), "stellar.services.nameserver.v1.LookupRequest")
	proto.RegisterType((*Record)(nil), "stellar.services.nameserver.v1.Record")
	proto.RegisterType((*LookupResponse)(nil), "stellar.services.nameserver.v1.LookupResponse")
	proto.RegisterType((*ListRequest)(nil), "stellar.services.nameserver.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "stellar.services.nameserver.v1.ListResponse")
	proto.RegisterType((*CreateRequest)(nil), "stellar.services.nameserver.v1.CreateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "stellar.services.nameserver.v1.DeleteRequest")
	proto.RegisterEnum("stellar.services.nameserver.v1.RecordType", RecordType_name, RecordType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Nameserver service

type NameserverClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type nameserverClient struct {
	cc *grpc.ClientConn
}

func NewNameserverClient(cc *grpc.ClientConn) NameserverClient {
	return &nameserverClient{cc}
}

func (c *nameserverClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Lookup(ctx context.Context, in *LookupRequest, opts ...grpc.CallOption) (*LookupResponse, error) {
	out := new(LookupResponse)
	err := grpc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nameserverClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.nameserver.v1.Nameserver/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nameserver service

type NameserverServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Lookup(context.Context, *LookupRequest) (*LookupResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Create(context.Context, *CreateRequest) (*google_protobuf1.Empty, error)
	Delete(context.Context, *DeleteRequest) (*google_protobuf1.Empty, error)
}

func RegisterNameserverServer(s *grpc.Server, srv NameserverServer) {
	s.RegisterService(&_Nameserver_serviceDesc, srv)
}

func _Nameserver_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Lookup(ctx, req.(*LookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nameserver_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameserverServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.nameserver.v1.Nameserver/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameserverServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nameserver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.nameserver.v1.Nameserver",
	HandlerType: (*NameserverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Nameserver_Info_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _Nameserver_Lookup_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Nameserver_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Nameserver_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Nameserver_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/nameserver/v1/nameserver.proto", fileDescriptorNameserver)
}

var fileDescriptorNameserver = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xe1, 0x6b, 0xd3, 0x40,
	0x1c, 0x5d, 0xd2, 0x34, 0x65, 0xbf, 0xae, 0x23, 0x1c, 0x65, 0xd4, 0x0a, 0x5a, 0x0a, 0x8e, 0x32,
	0xf5, 0xc2, 0xea, 0x47, 0x41, 0xec, 0xba, 0x22, 0xc3, 0xad, 0x93, 0x58, 0xb5, 0xf8, 0x2d, 0x6d,
	0x7f, 0xcd, 0x82, 0x69, 0x2e, 0x4b, 0x2e, 0x85, 0xf8, 0xe7, 0xf8, 0x7f, 0xe9, 0x07, 0xff, 0x12,
	0xb9, 0x5c, 0x62, 0x53, 0x45, 0x53, 0x06, 0xfb, 0x76, 0x8f, 0x7b, 0xef, 0xf7, 0xf2, 0x72, 0xef,
	0x0e, 0xde, 0x38, 0x2e, 0xbf, 0x89, 0x67, 0x74, 0xce, 0x56, 0x26, 0xde, 0xd8, 0x5f, 0x3d, 0xe4,
	0xdc, 0x8c, 0x38, 0x7a, 0x9e, 0x1d, 0x9a, 0x76, 0xe0, 0x9a, 0x11, 0x86, 0x6b, 0x77, 0x8e, 0x91,
	0xe9, 0xdb, 0x2b, 0x14, 0x00, 0x43, 0x73, 0x7d, 0x5a, 0x40, 0x34, 0x08, 0x19, 0x67, 0xe4, 0x51,
	0x26, 0xa2, 0xb9, 0x80, 0x16, 0x28, 0xeb, 0xd3, 0x76, 0xd3, 0x61, 0x0e, 0x4b, 0xa9, 0xa6, 0x58,
	0x49, 0x55, 0xfb, 0xa1, 0xc3, 0x98, 0xe3, 0xa1, 0x99, 0xa2, 0x59, 0xbc, 0x34, 0x71, 0x15, 0xf0,
	0x24, 0xdb, 0x7c, 0xf0, 0xe7, 0xa6, 0xed, 0x67, 0x5b, 0xdd, 0x06, 0xd4, 0x2f, 0xfc, 0x25, 0xb3,
	0xf0, 0x36, 0xc6, 0x88, 0x77, 0x8f, 0xe1, 0x40, 0xc2, 0x28, 0x60, 0x7e, 0x84, 0xe4, 0x08, 0x54,
	0x77, 0xd1, 0x52, 0x3a, 0x4a, 0x6f, 0xff, 0x4c, 0xff, 0xf9, 0xe3, 0xb1, 0x7a, 0x71, 0x6e, 0xa9,
	0xee, 0xa2, 0xfb, 0x04, 0x1a, 0x97, 0x8c, 0x7d, 0x89, 0x83, 0x4c, 0x48, 0x9a, 0x50, 0xbd, 0x8d,
	0x31, 0x4c, 0x24, 0xd7, 0x92, 0xa0, 0xfb, 0x4d, 0x01, 0xdd, 0xc2, 0x39, 0x0b, 0x17, 0xe4, 0x15,
	0x68, 0x3c, 0x09, 0x30, 0xdd, 0x3f, 0xec, 0x9f, 0xd0, 0xff, 0xa7, 0xa4, 0x52, 0x35, 0x49, 0x02,
	0xb4, 0x52, 0x1d, 0x21, 0xa0, 0x09, 0x46, 0x4b, 0x4d, 0xe7, 0xa7, 0x6b, 0x61, 0xba, 0xb6, 0xbd,
	0x18, 0x5b, 0x15, 0x69, 0x9a, 0x02, 0x42, 0xa1, 0xc6, 0x02, 0xee, 0x32, 0x3f, 0x6a, 0x69, 0x1d,
	0xa5, 0x57, 0xef, 0x37, 0xa9, 0xcc, 0x4f, 0xf3, 0xfc, 0x74, 0xe0, 0x27, 0x56, 0x4e, 0xea, 0x2e,
	0xe1, 0x30, 0xcf, 0x92, 0xa5, 0xce, 0xbd, 0x94, 0x82, 0xd7, 0x6b, 0xa8, 0x85, 0xe9, 0x37, 0x45,
	0x2d, 0xb5, 0x53, 0xe9, 0xd5, 0xfb, 0xc7, 0xbb, 0x45, 0xb0, 0x72, 0x99, 0xf8, 0xd5, 0x97, 0x6e,
	0xc4, 0xf3, 0x5f, 0xfd, 0x0e, 0x0e, 0x24, 0xcc, 0x4c, 0x0b, 0x06, 0xca, 0xdd, 0x0c, 0x10, 0x1a,
	0xc3, 0x10, 0x6d, 0x8e, 0xf9, 0xa1, 0xdc, 0x4f, 0x8e, 0x39, 0x34, 0xce, 0xd1, 0xc3, 0x8d, 0xcd,
	0x3d, 0x1c, 0xed, 0xc9, 0x08, 0x60, 0xc3, 0x23, 0x75, 0xa8, 0x7d, 0x18, 0xbf, 0x1d, 0x5f, 0x7f,
	0x1a, 0x1b, 0x7b, 0xa4, 0x0a, 0xca, 0xc0, 0x50, 0xc8, 0x3e, 0x54, 0x87, 0xe3, 0xc1, 0xd5, 0xc8,
	0x50, 0x89, 0x0e, 0xea, 0xd5, 0xd4, 0xa8, 0x90, 0x1a, 0x54, 0x26, 0xd3, 0x89, 0xa1, 0x89, 0xc5,
	0x7b, 0xeb, 0xa3, 0x51, 0xed, 0x7f, 0xaf, 0x00, 0x8c, 0x7f, 0xbb, 0x13, 0x1b, 0x34, 0x51, 0x6f,
	0xf2, 0xb4, 0xec, 0x1b, 0x0b, 0x77, 0xa2, 0xfd, 0x6c, 0x37, 0x72, 0x76, 0x8c, 0x0e, 0xe8, 0xb2,
	0x4d, 0xe4, 0x79, 0x99, 0x6e, 0xeb, 0x06, 0xb5, 0xe9, 0xae, 0xf4, 0xcc, 0xc8, 0x06, 0x4d, 0xf4,
	0xa7, 0x3c, 0x4b, 0xa1, 0x74, 0xe5, 0x59, 0xb6, 0x2a, 0x79, 0x0d, 0xba, 0x2c, 0x54, 0x79, 0x96,
	0xad, 0xe2, 0xb5, 0x8f, 0xfe, 0xba, 0x71, 0x23, 0xf1, 0x1c, 0x89, 0x81, 0xb2, 0x3a, 0xe5, 0x03,
	0xb7, 0x2a, 0xf6, 0xaf, 0x81, 0x67, 0xc3, 0xcf, 0x83, 0xbb, 0xbd, 0xbb, 0x2f, 0x37, 0x68, 0xba,
	0x37, 0xd3, 0xd3, 0xb1, 0x2f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x39, 0xba, 0x2b, 0x5d, 0xc5,
	0x05, 0x00, 0x00,
}
