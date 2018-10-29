// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/network/v1/network.proto

/*
Package network is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/network/v1/network.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	AllocateSubnetRequest
	AllocateSubnetResponse
	GetSubnetRequest
	GetSubnetResponse
	DeallocateSubnetRequest
	AllocateIPRequest
	AllocateIPResponse
	GetIPRequest
	GetIPResponse
	ReleaseIPRequest
	SubnetsResponse
	Subnet
	ConfigureRequest
	AddRouteRequest
	DeleteRouteRequest
	Route
	RoutesResponse
*/
package network

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"

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

type InfoRequest struct {
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{0} }

type InfoResponse struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{1} }

func (m *InfoResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type AllocateSubnetRequest struct {
	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *AllocateSubnetRequest) Reset()                    { *m = AllocateSubnetRequest{} }
func (m *AllocateSubnetRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocateSubnetRequest) ProtoMessage()               {}
func (*AllocateSubnetRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{2} }

func (m *AllocateSubnetRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type AllocateSubnetResponse struct {
	SubnetCIDR string `protobuf:"bytes,1,opt,name=subnet_cidr,json=subnetCidr,proto3" json:"subnet_cidr,omitempty"`
	Node       string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *AllocateSubnetResponse) Reset()                    { *m = AllocateSubnetResponse{} }
func (m *AllocateSubnetResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocateSubnetResponse) ProtoMessage()               {}
func (*AllocateSubnetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{3} }

func (m *AllocateSubnetResponse) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

func (m *AllocateSubnetResponse) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type GetSubnetRequest struct {
	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *GetSubnetRequest) Reset()                    { *m = GetSubnetRequest{} }
func (m *GetSubnetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubnetRequest) ProtoMessage()               {}
func (*GetSubnetRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{4} }

func (m *GetSubnetRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type GetSubnetResponse struct {
	SubnetCIDR string `protobuf:"bytes,1,opt,name=subnet_cidr,json=subnetCidr,proto3" json:"subnet_cidr,omitempty"`
}

func (m *GetSubnetResponse) Reset()                    { *m = GetSubnetResponse{} }
func (m *GetSubnetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetSubnetResponse) ProtoMessage()               {}
func (*GetSubnetResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{5} }

func (m *GetSubnetResponse) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

type DeallocateSubnetRequest struct {
	SubnetCIDR string `protobuf:"bytes,1,opt,name=subnet_cidr,json=subnetCidr,proto3" json:"subnet_cidr,omitempty"`
	Node       string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *DeallocateSubnetRequest) Reset()                    { *m = DeallocateSubnetRequest{} }
func (m *DeallocateSubnetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeallocateSubnetRequest) ProtoMessage()               {}
func (*DeallocateSubnetRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{6} }

func (m *DeallocateSubnetRequest) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

func (m *DeallocateSubnetRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type AllocateIPRequest struct {
	ID         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubnetCIDR string `protobuf:"bytes,2,opt,name=subnet_cidr,json=subnetCidr,proto3" json:"subnet_cidr,omitempty"`
	Node       string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *AllocateIPRequest) Reset()                    { *m = AllocateIPRequest{} }
func (m *AllocateIPRequest) String() string            { return proto.CompactTextString(m) }
func (*AllocateIPRequest) ProtoMessage()               {}
func (*AllocateIPRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{7} }

func (m *AllocateIPRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AllocateIPRequest) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

func (m *AllocateIPRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type AllocateIPResponse struct {
	IP   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *AllocateIPResponse) Reset()                    { *m = AllocateIPResponse{} }
func (m *AllocateIPResponse) String() string            { return proto.CompactTextString(m) }
func (*AllocateIPResponse) ProtoMessage()               {}
func (*AllocateIPResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{8} }

func (m *AllocateIPResponse) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *AllocateIPResponse) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type GetIPRequest struct {
	ID   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Node string `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *GetIPRequest) Reset()                    { *m = GetIPRequest{} }
func (m *GetIPRequest) String() string            { return proto.CompactTextString(m) }
func (*GetIPRequest) ProtoMessage()               {}
func (*GetIPRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{9} }

func (m *GetIPRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *GetIPRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type GetIPResponse struct {
	IP string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (m *GetIPResponse) Reset()                    { *m = GetIPResponse{} }
func (m *GetIPResponse) String() string            { return proto.CompactTextString(m) }
func (*GetIPResponse) ProtoMessage()               {}
func (*GetIPResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{10} }

func (m *GetIPResponse) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

type ReleaseIPRequest struct {
	ID   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Node string `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
}

func (m *ReleaseIPRequest) Reset()                    { *m = ReleaseIPRequest{} }
func (m *ReleaseIPRequest) String() string            { return proto.CompactTextString(m) }
func (*ReleaseIPRequest) ProtoMessage()               {}
func (*ReleaseIPRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{11} }

func (m *ReleaseIPRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ReleaseIPRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ReleaseIPRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type SubnetsResponse struct {
	Subnets []*Subnet `protobuf:"bytes,1,rep,name=subnets" json:"subnets,omitempty"`
}

func (m *SubnetsResponse) Reset()                    { *m = SubnetsResponse{} }
func (m *SubnetsResponse) String() string            { return proto.CompactTextString(m) }
func (*SubnetsResponse) ProtoMessage()               {}
func (*SubnetsResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{12} }

func (m *SubnetsResponse) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

type Subnet struct {
	CIDR    string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (m *Subnet) Reset()                    { *m = Subnet{} }
func (m *Subnet) String() string            { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()               {}
func (*Subnet) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{13} }

func (m *Subnet) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *Subnet) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type ConfigureRequest struct {
	Subnet *Subnet `protobuf:"bytes,1,opt,name=subnet" json:"subnet,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{14} }

func (m *ConfigureRequest) GetSubnet() *Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

type AddRouteRequest struct {
	CIDR   string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *AddRouteRequest) Reset()                    { *m = AddRouteRequest{} }
func (m *AddRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRouteRequest) ProtoMessage()               {}
func (*AddRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{15} }

func (m *AddRouteRequest) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *AddRouteRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type DeleteRouteRequest struct {
	CIDR   string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *DeleteRouteRequest) Reset()                    { *m = DeleteRouteRequest{} }
func (m *DeleteRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRouteRequest) ProtoMessage()               {}
func (*DeleteRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{16} }

func (m *DeleteRouteRequest) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *DeleteRouteRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type Route struct {
	CIDR   string `protobuf:"bytes,1,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{17} }

func (m *Route) GetCIDR() string {
	if m != nil {
		return m.CIDR
	}
	return ""
}

func (m *Route) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type RoutesResponse struct {
	Routes []*Route `protobuf:"bytes,1,rep,name=routes" json:"routes,omitempty"`
}

func (m *RoutesResponse) Reset()                    { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string            { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()               {}
func (*RoutesResponse) Descriptor() ([]byte, []int) { return fileDescriptorNetwork, []int{18} }

func (m *RoutesResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "stellar.services.network.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "stellar.services.network.v1.InfoResponse")
	proto.RegisterType((*AllocateSubnetRequest)(nil), "stellar.services.network.v1.AllocateSubnetRequest")
	proto.RegisterType((*AllocateSubnetResponse)(nil), "stellar.services.network.v1.AllocateSubnetResponse")
	proto.RegisterType((*GetSubnetRequest)(nil), "stellar.services.network.v1.GetSubnetRequest")
	proto.RegisterType((*GetSubnetResponse)(nil), "stellar.services.network.v1.GetSubnetResponse")
	proto.RegisterType((*DeallocateSubnetRequest)(nil), "stellar.services.network.v1.DeallocateSubnetRequest")
	proto.RegisterType((*AllocateIPRequest)(nil), "stellar.services.network.v1.AllocateIPRequest")
	proto.RegisterType((*AllocateIPResponse)(nil), "stellar.services.network.v1.AllocateIPResponse")
	proto.RegisterType((*GetIPRequest)(nil), "stellar.services.network.v1.GetIPRequest")
	proto.RegisterType((*GetIPResponse)(nil), "stellar.services.network.v1.GetIPResponse")
	proto.RegisterType((*ReleaseIPRequest)(nil), "stellar.services.network.v1.ReleaseIPRequest")
	proto.RegisterType((*SubnetsResponse)(nil), "stellar.services.network.v1.SubnetsResponse")
	proto.RegisterType((*Subnet)(nil), "stellar.services.network.v1.Subnet")
	proto.RegisterType((*ConfigureRequest)(nil), "stellar.services.network.v1.ConfigureRequest")
	proto.RegisterType((*AddRouteRequest)(nil), "stellar.services.network.v1.AddRouteRequest")
	proto.RegisterType((*DeleteRouteRequest)(nil), "stellar.services.network.v1.DeleteRouteRequest")
	proto.RegisterType((*Route)(nil), "stellar.services.network.v1.Route")
	proto.RegisterType((*RoutesResponse)(nil), "stellar.services.network.v1.RoutesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Network service

type NetworkClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	AllocateSubnet(ctx context.Context, in *AllocateSubnetRequest, opts ...grpc.CallOption) (*AllocateSubnetResponse, error)
	GetSubnet(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*GetSubnetResponse, error)
	DeallocateSubnet(ctx context.Context, in *DeallocateSubnetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Subnets(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*SubnetsResponse, error)
	AllocateIP(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocateIPResponse, error)
	GetIP(ctx context.Context, in *GetIPRequest, opts ...grpc.CallOption) (*GetIPResponse, error)
	ReleaseIP(ctx context.Context, in *ReleaseIPRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Routes(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*RoutesResponse, error)
}

type networkClient struct {
	cc *grpc.ClientConn
}

func NewNetworkClient(cc *grpc.ClientConn) NetworkClient {
	return &networkClient{cc}
}

func (c *networkClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) AllocateSubnet(ctx context.Context, in *AllocateSubnetRequest, opts ...grpc.CallOption) (*AllocateSubnetResponse, error) {
	out := new(AllocateSubnetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/AllocateSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetSubnet(ctx context.Context, in *GetSubnetRequest, opts ...grpc.CallOption) (*GetSubnetResponse, error) {
	out := new(GetSubnetResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/GetSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) DeallocateSubnet(ctx context.Context, in *DeallocateSubnetRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/DeallocateSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) Subnets(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*SubnetsResponse, error) {
	out := new(SubnetsResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/Subnets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) AllocateIP(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocateIPResponse, error) {
	out := new(AllocateIPResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/AllocateIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) GetIP(ctx context.Context, in *GetIPRequest, opts ...grpc.CallOption) (*GetIPResponse, error) {
	out := new(GetIPResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/GetIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) ReleaseIP(ctx context.Context, in *ReleaseIPRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/ReleaseIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) AddRoute(ctx context.Context, in *AddRouteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/AddRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/DeleteRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkClient) Routes(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*RoutesResponse, error) {
	out := new(RoutesResponse)
	err := grpc.Invoke(ctx, "/stellar.services.network.v1.Network/Routes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Network service

type NetworkServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	AllocateSubnet(context.Context, *AllocateSubnetRequest) (*AllocateSubnetResponse, error)
	GetSubnet(context.Context, *GetSubnetRequest) (*GetSubnetResponse, error)
	DeallocateSubnet(context.Context, *DeallocateSubnetRequest) (*google_protobuf.Empty, error)
	Subnets(context.Context, *google_protobuf.Empty) (*SubnetsResponse, error)
	AllocateIP(context.Context, *AllocateIPRequest) (*AllocateIPResponse, error)
	GetIP(context.Context, *GetIPRequest) (*GetIPResponse, error)
	ReleaseIP(context.Context, *ReleaseIPRequest) (*google_protobuf.Empty, error)
	Configure(context.Context, *ConfigureRequest) (*google_protobuf.Empty, error)
	AddRoute(context.Context, *AddRouteRequest) (*google_protobuf.Empty, error)
	DeleteRoute(context.Context, *DeleteRouteRequest) (*google_protobuf.Empty, error)
	Routes(context.Context, *google_protobuf.Empty) (*RoutesResponse, error)
}

func RegisterNetworkServer(s *grpc.Server, srv NetworkServer) {
	s.RegisterService(&_Network_serviceDesc, srv)
}

func _Network_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_AllocateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).AllocateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/AllocateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).AllocateSubnet(ctx, req.(*AllocateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/GetSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetSubnet(ctx, req.(*GetSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_DeallocateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeallocateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).DeallocateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/DeallocateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).DeallocateSubnet(ctx, req.(*DeallocateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_Subnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).Subnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/Subnets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).Subnets(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_AllocateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).AllocateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/AllocateIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).AllocateIP(ctx, req.(*AllocateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_GetIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).GetIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/GetIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).GetIP(ctx, req.(*GetIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_ReleaseIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).ReleaseIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/ReleaseIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).ReleaseIP(ctx, req.(*ReleaseIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_AddRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).AddRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/AddRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).AddRoute(ctx, req.(*AddRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/DeleteRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).DeleteRoute(ctx, req.(*DeleteRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Network_Routes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServer).Routes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.network.v1.Network/Routes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServer).Routes(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Network_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.network.v1.Network",
	HandlerType: (*NetworkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Network_Info_Handler,
		},
		{
			MethodName: "AllocateSubnet",
			Handler:    _Network_AllocateSubnet_Handler,
		},
		{
			MethodName: "GetSubnet",
			Handler:    _Network_GetSubnet_Handler,
		},
		{
			MethodName: "DeallocateSubnet",
			Handler:    _Network_DeallocateSubnet_Handler,
		},
		{
			MethodName: "Subnets",
			Handler:    _Network_Subnets_Handler,
		},
		{
			MethodName: "AllocateIP",
			Handler:    _Network_AllocateIP_Handler,
		},
		{
			MethodName: "GetIP",
			Handler:    _Network_GetIP_Handler,
		},
		{
			MethodName: "ReleaseIP",
			Handler:    _Network_ReleaseIP_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Network_Configure_Handler,
		},
		{
			MethodName: "AddRoute",
			Handler:    _Network_AddRoute_Handler,
		},
		{
			MethodName: "DeleteRoute",
			Handler:    _Network_DeleteRoute_Handler,
		},
		{
			MethodName: "Routes",
			Handler:    _Network_Routes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/network/v1/network.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/network/v1/network.proto", fileDescriptorNetwork)
}

var fileDescriptorNetwork = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x4e, 0xdb, 0x4c,
	0x10, 0x56, 0x4c, 0x70, 0xc8, 0x84, 0x43, 0x58, 0xfd, 0x7f, 0x8a, 0x4c, 0x25, 0x90, 0x2b, 0x51,
	0x28, 0xad, 0x2d, 0xa0, 0x57, 0x20, 0x24, 0x0e, 0xa9, 0x50, 0xaa, 0xb6, 0x44, 0xae, 0x84, 0xaa,
	0x56, 0xa5, 0x75, 0xe2, 0xc1, 0xb8, 0x35, 0x59, 0xd7, 0xde, 0x80, 0xe8, 0xc3, 0x72, 0xc1, 0x13,
	0xf4, 0x11, 0x2a, 0xbc, 0x6b, 0xc7, 0x58, 0xf1, 0xa1, 0x88, 0xbb, 0xf5, 0xe6, 0x3b, 0xcc, 0xce,
	0xac, 0x3f, 0x07, 0xf6, 0x6d, 0x87, 0x9d, 0x0f, 0x7b, 0x5a, 0x9f, 0x5e, 0xe8, 0x78, 0x6e, 0xfe,
	0x76, 0x91, 0x31, 0x3d, 0x60, 0xe8, 0xba, 0xa6, 0xaf, 0x9b, 0x9e, 0xa3, 0x07, 0xe8, 0x5f, 0x3a,
	0x7d, 0x0c, 0xf4, 0x01, 0xb2, 0x2b, 0xea, 0xff, 0xd4, 0x2f, 0x37, 0xa2, 0xa5, 0xe6, 0xf9, 0x94,
	0x51, 0xb2, 0x28, 0xe0, 0x5a, 0x04, 0xd5, 0xa2, 0xdf, 0x2f, 0x37, 0x94, 0x45, 0x9b, 0x52, 0xdb,
	0x45, 0x3d, 0x84, 0xf6, 0x86, 0x67, 0x3a, 0x5e, 0x78, 0xec, 0x9a, 0x33, 0x95, 0xff, 0x6c, 0x6a,
	0xd3, 0x70, 0xa9, 0xdf, 0xad, 0xf8, 0xae, 0x3a, 0x03, 0x8d, 0xce, 0xe0, 0x8c, 0x1a, 0xf8, 0x6b,
	0x88, 0x01, 0x53, 0x57, 0x60, 0x9a, 0x3f, 0x06, 0x1e, 0x1d, 0x04, 0x48, 0x5a, 0x20, 0x39, 0xd6,
	0x42, 0x65, 0xb9, 0xb2, 0x5a, 0x3f, 0x90, 0x6f, 0x6f, 0x96, 0xa4, 0x4e, 0xdb, 0x90, 0x1c, 0x4b,
	0x5d, 0x87, 0xff, 0xf7, 0x5d, 0x97, 0xf6, 0x4d, 0x86, 0x1f, 0x87, 0xbd, 0x01, 0x32, 0x21, 0x40,
	0x08, 0x54, 0x07, 0xd4, 0x42, 0x4e, 0x31, 0xc2, 0xb5, 0xfa, 0x15, 0x5a, 0x69, 0xb0, 0x90, 0xd7,
	0xa1, 0x11, 0x84, 0x3b, 0xdf, 0xfa, 0x8e, 0xe5, 0x0b, 0x9f, 0xd9, 0xdb, 0x9b, 0x25, 0xe0, 0xc0,
	0xc3, 0x4e, 0xdb, 0x30, 0x80, 0x43, 0x0e, 0x1d, 0xcb, 0x8f, 0xe5, 0xa5, 0x84, 0xfc, 0x0a, 0x34,
	0x8f, 0x90, 0x15, 0x97, 0xd1, 0x86, 0xf9, 0x04, 0xee, 0x81, 0x15, 0xa8, 0xa7, 0xf0, 0xa4, 0x8d,
	0xe6, 0xd8, 0xb3, 0x3f, 0xca, 0x69, 0x3c, 0x98, 0x8f, 0x9a, 0xd5, 0xe9, 0x46, 0xca, 0x19, 0x63,
	0x48, 0x3b, 0x4a, 0xa5, 0x1d, 0x27, 0x12, 0x8e, 0x7b, 0x40, 0x92, 0x8e, 0x89, 0xc9, 0x7b, 0xf7,
	0x2c, 0xbb, 0x86, 0xe4, 0x78, 0x63, 0x6b, 0xde, 0x86, 0xe9, 0x23, 0x64, 0xc5, 0xe5, 0x8e, 0xe3,
	0x3e, 0x87, 0x19, 0xc1, 0xcd, 0x37, 0x56, 0x4f, 0xa0, 0x69, 0xa0, 0x8b, 0x66, 0x50, 0xa2, 0x2f,
	0x5c, 0x43, 0xca, 0x2c, 0x3e, 0x79, 0xfc, 0x2e, 0xcc, 0xf1, 0x66, 0x05, 0x71, 0x09, 0xbb, 0x50,
	0xe3, 0x3d, 0x0b, 0x16, 0x2a, 0xcb, 0x13, 0xab, 0x8d, 0xcd, 0x67, 0x5a, 0xce, 0x6b, 0xa7, 0x89,
	0x5b, 0x10, 0x71, 0xd4, 0x3d, 0x90, 0xf9, 0x16, 0x79, 0x0a, 0xd5, 0xc4, 0x55, 0x98, 0xba, 0xbd,
	0x59, 0xaa, 0x86, 0x23, 0x09, 0x77, 0xc9, 0x02, 0xd4, 0x6c, 0x93, 0xe1, 0x95, 0x79, 0x2d, 0x3a,
	0x12, 0x3d, 0xaa, 0xc7, 0xd0, 0x3c, 0xa4, 0x83, 0x33, 0xc7, 0x1e, 0xfa, 0x18, 0x9d, 0x75, 0x07,
	0x64, 0x6e, 0x10, 0xaa, 0x95, 0xac, 0x49, 0x50, 0xd4, 0x23, 0x98, 0xdb, 0xb7, 0x2c, 0x83, 0x0e,
	0x59, 0xac, 0x97, 0x5f, 0x5b, 0x0b, 0x64, 0x66, 0xfa, 0x36, 0x32, 0x51, 0x9a, 0x78, 0x52, 0xdf,
	0x02, 0x69, 0xa3, 0x8b, 0x0c, 0x1f, 0x41, 0x6b, 0x17, 0x26, 0x43, 0x95, 0x07, 0xd2, 0xdf, 0xc1,
	0x6c, 0x48, 0x1f, 0xcd, 0x6d, 0x1b, 0x64, 0x3f, 0xdc, 0x11, 0x63, 0x53, 0x73, 0x5b, 0xc4, 0x4f,
	0x20, 0x18, 0x9b, 0x7f, 0xa6, 0xa0, 0xf6, 0x81, 0xff, 0x48, 0xbe, 0x40, 0xf5, 0x2e, 0x05, 0xc9,
	0x6a, 0x2e, 0x3f, 0x91, 0x9b, 0xca, 0x5a, 0x09, 0xa4, 0x28, 0xf2, 0x1a, 0x66, 0xef, 0xa7, 0x21,
	0xd9, 0xcc, 0x25, 0x8f, 0xcd, 0x59, 0x65, 0xeb, 0x9f, 0x38, 0xc2, 0xfa, 0x07, 0xd4, 0xe3, 0x04,
	0x24, 0xaf, 0x72, 0x15, 0xd2, 0x89, 0xaa, 0x68, 0x65, 0xe1, 0xc2, 0xeb, 0x3b, 0x34, 0xd3, 0x39,
	0x49, 0x5e, 0xe7, 0x6a, 0x64, 0xc4, 0xaa, 0xd2, 0xd2, 0xf8, 0x67, 0x4d, 0x8b, 0x3e, 0x6b, 0xda,
	0x9b, 0xbb, 0xcf, 0x1a, 0x39, 0x86, 0x9a, 0x78, 0x71, 0x49, 0x06, 0x44, 0x79, 0x59, 0xe2, 0x1d,
	0x19, 0x5d, 0x9f, 0x0b, 0x80, 0x51, 0x10, 0x12, 0xad, 0x54, 0x87, 0xe3, 0x2c, 0x52, 0xf4, 0xd2,
	0x78, 0x61, 0x77, 0x0a, 0x93, 0x61, 0xf2, 0x91, 0xb5, 0xa2, 0xd6, 0x8e, 0x4c, 0x5e, 0x94, 0x81,
	0x0a, 0x7d, 0x03, 0xea, 0x71, 0x60, 0x16, 0x4c, 0x3b, 0x1d, 0xac, 0x99, 0x3d, 0x37, 0xa0, 0x1e,
	0x07, 0x53, 0x81, 0x66, 0x3a, 0xc0, 0x32, 0x35, 0xbb, 0x30, 0x15, 0x65, 0x13, 0xc9, 0x1f, 0x58,
	0x2a, 0xc2, 0x32, 0x15, 0x4f, 0xa0, 0x91, 0x08, 0x29, 0xa2, 0x17, 0x5c, 0xbb, 0x74, 0x9c, 0x65,
	0xea, 0xbe, 0x07, 0x99, 0x27, 0x4e, 0xe6, 0x85, 0x5b, 0x2f, 0x4e, 0x9c, 0xf8, 0xbe, 0x1d, 0xec,
	0x7e, 0xde, 0x79, 0xc0, 0x1f, 0xc2, 0x1d, 0xb1, 0xfc, 0x54, 0xe9, 0xc9, 0xa1, 0xfb, 0xd6, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x7c, 0xce, 0x1c, 0x58, 0x0a, 0x00, 0x00,
}
