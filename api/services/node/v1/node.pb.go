// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/ehazlett/stellar/api/services/node/v1/node.proto

/*
Package node is a generated protocol buffer package.

It is generated from these files:
	github.com/ehazlett/stellar/api/services/node/v1/node.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	ContainersRequest
	Container
	ContainersResponse
	ContainerRequest
	ContainerResponse
	ImagesRequest
	Image
	ImagesResponse
	ContainerNetworkRequest
	Process
	Mount
	Endpoint
	Service
	CreateContainerRequest
	DeleteContainerRequest
	RestartContainerRequest
*/
package node

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
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

type Protocol int32

const (
	Protocol_UNKNOWN Protocol = 0
	Protocol_TCP     Protocol = 1
	Protocol_UDP     Protocol = 2
	Protocol_HTTP    Protocol = 3
)

var Protocol_name = map[int32]string{
	0: "UNKNOWN",
	1: "TCP",
	2: "UDP",
	3: "HTTP",
}
var Protocol_value = map[string]int32{
	"UNKNOWN": 0,
	"TCP":     1,
	"UDP":     2,
	"HTTP":    3,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptorNode, []int{0} }

type InfoRequest struct {
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{0} }

type InfoResponse struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{1} }

func (m *InfoResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ContainersRequest struct {
	Filters []string `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *ContainersRequest) Reset()                    { *m = ContainersRequest{} }
func (m *ContainersRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainersRequest) ProtoMessage()               {}
func (*ContainersRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{2} }

func (m *ContainersRequest) GetFilters() []string {
	if m != nil {
		return m.Filters
	}
	return nil
}

type Container struct {
	ID          string                           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image       string                           `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Labels      map[string]string                `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Spec        *google_protobuf1.Any            `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	Snapshotter string                           `protobuf:"bytes,5,opt,name=snapshotter,proto3" json:"snapshotter,omitempty"`
	Task        *Container_Task                  `protobuf:"bytes,6,opt,name=task" json:"task,omitempty"`
	Runtime     string                           `protobuf:"bytes,7,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Extensions  map[string]*google_protobuf1.Any `protobuf:"bytes,8,rep,name=extensions" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{3} }

func (m *Container) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Container) GetSpec() *google_protobuf1.Any {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Container) GetSnapshotter() string {
	if m != nil {
		return m.Snapshotter
	}
	return ""
}

func (m *Container) GetTask() *Container_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *Container) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *Container) GetExtensions() map[string]*google_protobuf1.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

type Container_Task struct {
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (m *Container_Task) Reset()                    { *m = Container_Task{} }
func (m *Container_Task) String() string            { return proto.CompactTextString(m) }
func (*Container_Task) ProtoMessage()               {}
func (*Container_Task) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{3, 1} }

func (m *Container_Task) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type ContainersResponse struct {
	Containers []*Container `protobuf:"bytes,1,rep,name=containers" json:"containers,omitempty"`
}

func (m *ContainersResponse) Reset()                    { *m = ContainersResponse{} }
func (m *ContainersResponse) String() string            { return proto.CompactTextString(m) }
func (*ContainersResponse) ProtoMessage()               {}
func (*ContainersResponse) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{4} }

func (m *ContainersResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

type ContainerRequest struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ContainerRequest) Reset()                    { *m = ContainerRequest{} }
func (m *ContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainerRequest) ProtoMessage()               {}
func (*ContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{5} }

func (m *ContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ContainerResponse struct {
	Container *Container `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
}

func (m *ContainerResponse) Reset()                    { *m = ContainerResponse{} }
func (m *ContainerResponse) String() string            { return proto.CompactTextString(m) }
func (*ContainerResponse) ProtoMessage()               {}
func (*ContainerResponse) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{6} }

func (m *ContainerResponse) GetContainer() *Container {
	if m != nil {
		return m.Container
	}
	return nil
}

type ImagesRequest struct {
}

func (m *ImagesRequest) Reset()                    { *m = ImagesRequest{} }
func (m *ImagesRequest) String() string            { return proto.CompactTextString(m) }
func (*ImagesRequest) ProtoMessage()               {}
func (*ImagesRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{7} }

type Image struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{8} }

func (m *Image) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ImagesResponse struct {
	Images []*Image `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
}

func (m *ImagesResponse) Reset()                    { *m = ImagesResponse{} }
func (m *ImagesResponse) String() string            { return proto.CompactTextString(m) }
func (*ImagesResponse) ProtoMessage()               {}
func (*ImagesResponse) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{9} }

func (m *ImagesResponse) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type ContainerNetworkRequest struct {
	ID      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IP      string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Gateway string `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (m *ContainerNetworkRequest) Reset()                    { *m = ContainerNetworkRequest{} }
func (m *ContainerNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainerNetworkRequest) ProtoMessage()               {}
func (*ContainerNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{10} }

func (m *ContainerNetworkRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *ContainerNetworkRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ContainerNetworkRequest) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ContainerNetworkRequest) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type Process struct {
	Uid  uint32   `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Gid  uint32   `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	Args []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	Env  []string `protobuf:"bytes,4,rep,name=env" json:"env,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{11} }

func (m *Process) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Process) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *Process) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *Process) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

type Mount struct {
	Type        string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Source      string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination string   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Options     []string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty"`
}

func (m *Mount) Reset()                    { *m = Mount{} }
func (m *Mount) String() string            { return proto.CompactTextString(m) }
func (*Mount) ProtoMessage()               {}
func (*Mount) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{12} }

func (m *Mount) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mount) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Mount) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *Mount) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

// NOTE: this uses a custom json unmarshaller (utils.go); update if fields are changed
type Endpoint struct {
	Service  string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Protocol Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=stellar.services.node.v1.Protocol" json:"protocol,omitempty"`
	Host     string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port     uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	TLS      bool     `protobuf:"varint,5,opt,name=tls,proto3" json:"tls,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{13} }

func (m *Endpoint) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Endpoint) GetProtocol() Protocol {
	if m != nil {
		return m.Protocol
	}
	return Protocol_UNKNOWN
}

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetTLS() bool {
	if m != nil {
		return m.TLS
	}
	return false
}

type Service struct {
	Name        string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Image       string      `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Runtime     string      `protobuf:"bytes,3,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Snapshotter string      `protobuf:"bytes,4,opt,name=snapshotter,proto3" json:"snapshotter,omitempty"`
	Node        string      `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	Process     *Process    `protobuf:"bytes,6,opt,name=process" json:"process,omitempty"`
	Labels      []string    `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty"`
	Network     bool        `protobuf:"varint,8,opt,name=network,proto3" json:"network,omitempty"`
	Mounts      []*Mount    `protobuf:"bytes,9,rep,name=mounts" json:"mounts,omitempty"`
	Endpoints   []*Endpoint `protobuf:"bytes,10,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{14} }

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Service) GetRuntime() string {
	if m != nil {
		return m.Runtime
	}
	return ""
}

func (m *Service) GetSnapshotter() string {
	if m != nil {
		return m.Snapshotter
	}
	return ""
}

func (m *Service) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Service) GetProcess() *Process {
	if m != nil {
		return m.Process
	}
	return nil
}

func (m *Service) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Service) GetNetwork() bool {
	if m != nil {
		return m.Network
	}
	return false
}

func (m *Service) GetMounts() []*Mount {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *Service) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type CreateContainerRequest struct {
	Application string   `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Service     *Service `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
}

func (m *CreateContainerRequest) Reset()                    { *m = CreateContainerRequest{} }
func (m *CreateContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateContainerRequest) ProtoMessage()               {}
func (*CreateContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{15} }

func (m *CreateContainerRequest) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *CreateContainerRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteContainerRequest struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *DeleteContainerRequest) Reset()                    { *m = DeleteContainerRequest{} }
func (m *DeleteContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteContainerRequest) ProtoMessage()               {}
func (*DeleteContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{16} }

func (m *DeleteContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RestartContainerRequest struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *RestartContainerRequest) Reset()                    { *m = RestartContainerRequest{} }
func (m *RestartContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*RestartContainerRequest) ProtoMessage()               {}
func (*RestartContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptorNode, []int{17} }

func (m *RestartContainerRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "stellar.services.node.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "stellar.services.node.v1.InfoResponse")
	proto.RegisterType((*ContainersRequest)(nil), "stellar.services.node.v1.ContainersRequest")
	proto.RegisterType((*Container)(nil), "stellar.services.node.v1.Container")
	proto.RegisterType((*Container_Task)(nil), "stellar.services.node.v1.Container.Task")
	proto.RegisterType((*ContainersResponse)(nil), "stellar.services.node.v1.ContainersResponse")
	proto.RegisterType((*ContainerRequest)(nil), "stellar.services.node.v1.ContainerRequest")
	proto.RegisterType((*ContainerResponse)(nil), "stellar.services.node.v1.ContainerResponse")
	proto.RegisterType((*ImagesRequest)(nil), "stellar.services.node.v1.ImagesRequest")
	proto.RegisterType((*Image)(nil), "stellar.services.node.v1.Image")
	proto.RegisterType((*ImagesResponse)(nil), "stellar.services.node.v1.ImagesResponse")
	proto.RegisterType((*ContainerNetworkRequest)(nil), "stellar.services.node.v1.ContainerNetworkRequest")
	proto.RegisterType((*Process)(nil), "stellar.services.node.v1.Process")
	proto.RegisterType((*Mount)(nil), "stellar.services.node.v1.Mount")
	proto.RegisterType((*Endpoint)(nil), "stellar.services.node.v1.Endpoint")
	proto.RegisterType((*Service)(nil), "stellar.services.node.v1.Service")
	proto.RegisterType((*CreateContainerRequest)(nil), "stellar.services.node.v1.CreateContainerRequest")
	proto.RegisterType((*DeleteContainerRequest)(nil), "stellar.services.node.v1.DeleteContainerRequest")
	proto.RegisterType((*RestartContainerRequest)(nil), "stellar.services.node.v1.RestartContainerRequest")
	proto.RegisterEnum("stellar.services.node.v1.Protocol", Protocol_name, Protocol_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersResponse, error)
	Container(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error)
	Images(ctx context.Context, in *ImagesRequest, opts ...grpc.CallOption) (*ImagesResponse, error)
	SetupContainerNetwork(ctx context.Context, in *ContainerNetworkRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	RestartContainer(ctx context.Context, in *RestartContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Containers(ctx context.Context, in *ContainersRequest, opts ...grpc.CallOption) (*ContainersResponse, error) {
	out := new(ContainersResponse)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/Containers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Container(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error) {
	out := new(ContainerResponse)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/Container", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Images(ctx context.Context, in *ImagesRequest, opts ...grpc.CallOption) (*ImagesResponse, error) {
	out := new(ImagesResponse)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/Images", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) SetupContainerNetwork(ctx context.Context, in *ContainerNetworkRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/SetupContainerNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) CreateContainer(ctx context.Context, in *CreateContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/CreateContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) DeleteContainer(ctx context.Context, in *DeleteContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/DeleteContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) RestartContainer(ctx context.Context, in *RestartContainerRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/stellar.services.node.v1.Node/RestartContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Containers(context.Context, *ContainersRequest) (*ContainersResponse, error)
	Container(context.Context, *ContainerRequest) (*ContainerResponse, error)
	Images(context.Context, *ImagesRequest) (*ImagesResponse, error)
	SetupContainerNetwork(context.Context, *ContainerNetworkRequest) (*google_protobuf2.Empty, error)
	CreateContainer(context.Context, *CreateContainerRequest) (*google_protobuf2.Empty, error)
	DeleteContainer(context.Context, *DeleteContainerRequest) (*google_protobuf2.Empty, error)
	RestartContainer(context.Context, *RestartContainerRequest) (*google_protobuf2.Empty, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Containers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Containers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/Containers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Containers(ctx, req.(*ContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Container_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Container(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/Container",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Container(ctx, req.(*ContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Images_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Images(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/Images",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Images(ctx, req.(*ImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_SetupContainerNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).SetupContainerNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/SetupContainerNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).SetupContainerNetwork(ctx, req.(*ContainerNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_CreateContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).CreateContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/CreateContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).CreateContainer(ctx, req.(*CreateContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_DeleteContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).DeleteContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/DeleteContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).DeleteContainer(ctx, req.(*DeleteContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_RestartContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).RestartContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stellar.services.node.v1.Node/RestartContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).RestartContainer(ctx, req.(*RestartContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stellar.services.node.v1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Node_Info_Handler,
		},
		{
			MethodName: "Containers",
			Handler:    _Node_Containers_Handler,
		},
		{
			MethodName: "Container",
			Handler:    _Node_Container_Handler,
		},
		{
			MethodName: "Images",
			Handler:    _Node_Images_Handler,
		},
		{
			MethodName: "SetupContainerNetwork",
			Handler:    _Node_SetupContainerNetwork_Handler,
		},
		{
			MethodName: "CreateContainer",
			Handler:    _Node_CreateContainer_Handler,
		},
		{
			MethodName: "DeleteContainer",
			Handler:    _Node_DeleteContainer_Handler,
		},
		{
			MethodName: "RestartContainer",
			Handler:    _Node_RestartContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/ehazlett/stellar/api/services/node/v1/node.proto",
}

func init() {
	proto.RegisterFile("github.com/ehazlett/stellar/api/services/node/v1/node.proto", fileDescriptorNode)
}

var fileDescriptorNode = []byte{
	// 1098 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x57, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x26, 0x89, 0x9b, 0x9f, 0x13, 0xba, 0x1b, 0x46, 0xa5, 0xeb, 0x0d, 0x17, 0x0d, 0x46, 0x2c,
	0x51, 0x17, 0xec, 0x6d, 0x2b, 0x54, 0xa0, 0x08, 0xb1, 0xfd, 0x11, 0x54, 0x2c, 0x25, 0x9a, 0xb4,
	0x20, 0x58, 0x71, 0xe1, 0x26, 0x53, 0xd7, 0xaa, 0x33, 0x63, 0x3c, 0xe3, 0x94, 0x20, 0x71, 0xcb,
	0x9b, 0x20, 0xde, 0x6a, 0x2f, 0xf6, 0x9e, 0x77, 0x40, 0x33, 0x1e, 0x3b, 0xde, 0xb4, 0x4e, 0xcc,
	0x55, 0xce, 0x39, 0x73, 0xfe, 0x7c, 0xe6, 0x3b, 0x9f, 0x1d, 0x38, 0xf0, 0x7c, 0x71, 0x1d, 0x5f,
	0xda, 0x23, 0x36, 0x71, 0xc8, 0xb5, 0xfb, 0x47, 0x40, 0x84, 0x70, 0xb8, 0x20, 0x41, 0xe0, 0x46,
	0x8e, 0x1b, 0xfa, 0x0e, 0x27, 0xd1, 0xd4, 0x1f, 0x11, 0xee, 0x50, 0x36, 0x26, 0xce, 0x74, 0x47,
	0xfd, 0xda, 0x61, 0xc4, 0x04, 0x43, 0xa6, 0x76, 0xb4, 0x53, 0x27, 0x5b, 0x1d, 0x4e, 0x77, 0xba,
	0x1b, 0x1e, 0xf3, 0x98, 0x72, 0x72, 0xa4, 0x94, 0xf8, 0x77, 0x1f, 0x7b, 0x8c, 0x79, 0x01, 0x71,
	0x94, 0x76, 0x19, 0x5f, 0x39, 0x2e, 0x9d, 0xe9, 0xa3, 0xf7, 0x16, 0x8f, 0xc8, 0x24, 0x14, 0xfa,
	0xd0, 0x5a, 0x87, 0xf6, 0x29, 0xbd, 0x62, 0x98, 0xfc, 0x16, 0x13, 0x2e, 0xac, 0x27, 0xf0, 0x76,
	0xa2, 0xf2, 0x90, 0x51, 0x4e, 0xd0, 0x26, 0x54, 0xfd, 0xb1, 0x59, 0xe9, 0x55, 0xfa, 0xad, 0xc3,
	0xfa, 0xeb, 0x57, 0x5b, 0xd5, 0xd3, 0x63, 0x5c, 0xf5, 0xc7, 0xd6, 0x27, 0xf0, 0xce, 0x11, 0xa3,
	0xc2, 0xf5, 0x29, 0x89, 0xb8, 0x0e, 0x46, 0x26, 0x34, 0xae, 0xfc, 0x40, 0x90, 0x88, 0x9b, 0x95,
	0x5e, 0xad, 0xdf, 0xc2, 0xa9, 0x6a, 0xfd, 0x6d, 0x40, 0x2b, 0xf3, 0x2f, 0x4a, 0x8a, 0x36, 0x60,
	0xcd, 0x9f, 0xb8, 0x1e, 0x31, 0xab, 0xf2, 0x08, 0x27, 0x0a, 0xfa, 0x06, 0xea, 0x81, 0x7b, 0x49,
	0x02, 0x6e, 0xd6, 0x7a, 0xb5, 0x7e, 0x7b, 0xd7, 0xb1, 0x8b, 0x46, 0x63, 0x67, 0x25, 0xec, 0x17,
	0x2a, 0xe2, 0x84, 0x8a, 0x68, 0x86, 0x75, 0x38, 0xea, 0x83, 0xc1, 0x43, 0x32, 0x32, 0x8d, 0x5e,
	0xa5, 0xdf, 0xde, 0xdd, 0xb0, 0x93, 0xb1, 0xd8, 0xe9, 0x58, 0xec, 0xe7, 0x74, 0x86, 0x95, 0x07,
	0xea, 0x41, 0x9b, 0x53, 0x37, 0xe4, 0xd7, 0x4c, 0x08, 0x12, 0x99, 0x6b, 0xaa, 0x9d, 0xbc, 0x09,
	0x7d, 0x09, 0x86, 0x70, 0xf9, 0x8d, 0x59, 0x57, 0xb9, 0xfa, 0x65, 0x5a, 0x3a, 0x77, 0xf9, 0x0d,
	0x56, 0x51, 0x72, 0x50, 0x51, 0x4c, 0x85, 0x3f, 0x21, 0x66, 0x43, 0xe5, 0x4e, 0x55, 0x34, 0x04,
	0x20, 0xbf, 0x0b, 0x42, 0xb9, 0xcf, 0x28, 0x37, 0x9b, 0xea, 0x81, 0xf7, 0xca, 0x64, 0x3f, 0xc9,
	0xa2, 0x92, 0x87, 0xce, 0xa5, 0xe9, 0x7e, 0x0e, 0xed, 0xdc, 0x3c, 0x50, 0x07, 0x6a, 0x37, 0x64,
	0x96, 0xcc, 0x1f, 0x4b, 0x51, 0x0e, 0x7e, 0xea, 0x06, 0x71, 0x36, 0x78, 0xa5, 0x7c, 0x51, 0xfd,
	0xac, 0xd2, 0x35, 0xc1, 0x90, 0x7d, 0xcb, 0x98, 0x50, 0xdf, 0xd9, 0x3a, 0x96, 0x62, 0x77, 0x08,
	0x0f, 0x17, 0x6a, 0xde, 0x93, 0x78, 0x3b, 0x9f, 0xb8, 0x68, 0xe6, 0xf3, 0x72, 0xd6, 0xcf, 0x80,
	0xf2, 0xb0, 0xd2, 0x20, 0x3c, 0x02, 0x18, 0x65, 0x56, 0x05, 0xad, 0xf6, 0xee, 0x07, 0x25, 0x86,
	0x82, 0x73, 0x61, 0xd6, 0x36, 0x74, 0xe6, 0x07, 0x1a, 0xb0, 0x45, 0xe8, 0xfe, 0x31, 0x87, 0xee,
	0xac, 0x8b, 0xe7, 0xd0, 0xca, 0xd2, 0xa9, 0x98, 0x92, 0x4d, 0xcc, 0xa3, 0xac, 0x87, 0xb0, 0x7e,
	0x2a, 0x31, 0x9d, 0x6e, 0x8c, 0xb5, 0x05, 0x6b, 0xca, 0x50, 0xd8, 0xc9, 0x29, 0x3c, 0x48, 0x23,
	0x74, 0x1b, 0xfb, 0x50, 0x57, 0x7b, 0x91, 0x0e, 0x62, 0xab, 0xb8, 0x07, 0x15, 0x89, 0xb5, 0xbb,
	0xf5, 0x27, 0x3c, 0xca, 0x9a, 0x3a, 0x23, 0xe2, 0x96, 0x45, 0x37, 0x2b, 0xe6, 0xa0, 0xec, 0x61,
	0x02, 0x0a, 0x6d, 0x1f, 0xe0, 0xaa, 0x1f, 0x4a, 0xfc, 0xd2, 0x24, 0x83, 0x59, 0x4b, 0xf0, 0xab,
	0x55, 0x79, 0xe2, 0xb9, 0x82, 0xdc, 0xba, 0x33, 0xb5, 0x66, 0x2d, 0x9c, 0xaa, 0xd6, 0x10, 0x1a,
	0x83, 0x88, 0x8d, 0x08, 0xe7, 0x12, 0x27, 0xf1, 0x1c, 0x4c, 0xb1, 0x3f, 0x96, 0x16, 0xcf, 0x1f,
	0xab, 0x4a, 0xeb, 0x58, 0x8a, 0x08, 0x81, 0xe1, 0x46, 0x5e, 0xb2, 0xf3, 0x2d, 0xac, 0x64, 0xe9,
	0x45, 0xe8, 0xd4, 0x34, 0x94, 0x49, 0x8a, 0x16, 0x83, 0xb5, 0xef, 0x59, 0x4c, 0x85, 0x74, 0x17,
	0xb3, 0x90, 0x68, 0xec, 0x29, 0x19, 0x6d, 0x42, 0x9d, 0xb3, 0x38, 0x1a, 0xa5, 0xb0, 0xd6, 0x9a,
	0xdc, 0xee, 0x31, 0xe1, 0xc2, 0xa7, 0xae, 0xf0, 0x19, 0xd5, 0x7d, 0xe6, 0x4d, 0xf2, 0x29, 0x58,
	0x28, 0xd4, 0x0a, 0xae, 0x25, 0x44, 0xa6, 0x55, 0xeb, 0x9f, 0x0a, 0x34, 0x4f, 0xe8, 0x38, 0x64,
	0x3e, 0x55, 0x7c, 0xa7, 0x67, 0xae, 0xeb, 0xa6, 0x2a, 0xfa, 0x0a, 0x9a, 0x0a, 0xe2, 0x23, 0x16,
	0xa8, 0xe2, 0x0f, 0x76, 0xad, 0xe2, 0x6b, 0x1a, 0x68, 0x4f, 0x9c, 0xc5, 0xc8, 0xc7, 0xb9, 0x66,
	0x5c, 0xe8, 0xe9, 0x2a, 0x59, 0xda, 0x42, 0x16, 0x09, 0xd5, 0xef, 0x3a, 0x56, 0x32, 0x7a, 0x0c,
	0x35, 0x11, 0x70, 0x45, 0x50, 0xcd, 0xc3, 0xc6, 0xeb, 0x57, 0x5b, 0xb5, 0xf3, 0x17, 0x43, 0x2c,
	0x6d, 0xd6, 0xbf, 0x55, 0x68, 0x0c, 0x75, 0x3b, 0x08, 0x0c, 0xea, 0x4e, 0xb2, 0xe9, 0x48, 0xb9,
	0x80, 0x6c, 0x73, 0xcc, 0x54, 0x7b, 0x93, 0x99, 0x16, 0x38, 0xd1, 0xb8, 0xcb, 0x89, 0xb2, 0x0a,
	0x1b, 0x13, 0x4d, 0x97, 0x4a, 0x46, 0x07, 0xd0, 0x08, 0x93, 0x5b, 0xd7, 0x54, 0xf9, 0xfe, 0xd2,
	0x39, 0x48, 0x47, 0x9c, 0x46, 0xc8, 0x0b, 0xd4, 0xcc, 0xdf, 0x50, 0xb7, 0x90, 0x12, 0x79, 0x0e,
	0x7e, 0x4d, 0xf9, 0xe4, 0x73, 0xf8, 0xed, 0x43, 0x7d, 0x22, 0xf1, 0xc0, 0xcd, 0xd6, 0xaa, 0xe5,
	0x50, 0xb8, 0xc1, 0xda, 0x1d, 0x7d, 0x0d, 0x2d, 0xa2, 0xaf, 0x95, 0x9b, 0xa0, 0x62, 0x97, 0xdc,
	0x58, 0x8a, 0x00, 0x3c, 0x0f, 0xb2, 0x6e, 0x61, 0xf3, 0x28, 0x22, 0xae, 0x20, 0x77, 0x58, 0xa6,
	0x07, 0x6d, 0x37, 0x0c, 0x03, 0x7f, 0x94, 0xe0, 0x2d, 0xb9, 0x84, 0xbc, 0x49, 0x4e, 0x29, 0x05,
	0x52, 0x75, 0xd5, 0x94, 0xf4, 0x9d, 0x66, 0x58, 0xb3, 0x9e, 0xc1, 0xe6, 0x31, 0x09, 0xc8, 0x3d,
	0x85, 0x8b, 0x48, 0x65, 0x07, 0x1e, 0x61, 0xc2, 0x85, 0x1b, 0x89, 0xb2, 0x21, 0xdb, 0x7b, 0xd0,
	0x4c, 0x61, 0x8a, 0xda, 0xd0, 0xb8, 0x38, 0xfb, 0xee, 0xec, 0x87, 0x9f, 0xce, 0x3a, 0x6f, 0xa1,
	0x06, 0xd4, 0xce, 0x8f, 0x06, 0x9d, 0x8a, 0x14, 0x2e, 0x8e, 0x07, 0x9d, 0x2a, 0x6a, 0x82, 0xf1,
	0xed, 0xf9, 0xf9, 0xa0, 0x53, 0xdb, 0xfd, 0xab, 0x0e, 0xc6, 0x99, 0x44, 0xc1, 0x05, 0x18, 0xf2,
	0xab, 0x02, 0x7d, 0xb8, 0x84, 0xab, 0xe6, 0x1f, 0x21, 0xdd, 0x27, 0xab, 0xdc, 0x34, 0x15, 0x7a,
	0x00, 0xf3, 0xb7, 0x05, 0x7a, 0x5a, 0x82, 0x8c, 0x53, 0xe2, 0xed, 0x7e, 0x5c, 0xce, 0x59, 0x17,
	0x1a, 0xe7, 0xbf, 0x5e, 0xb6, 0xcb, 0x90, 0xbe, 0x2e, 0xf3, 0xb4, 0x94, 0xaf, 0xae, 0xf2, 0x12,
	0xea, 0x09, 0xd7, 0xa3, 0x8f, 0x56, 0x70, 0x7a, 0xf6, 0x18, 0xfd, 0xd5, 0x8e, 0x3a, 0xf9, 0x25,
	0xbc, 0x3b, 0x24, 0x22, 0x0e, 0x17, 0x5f, 0x01, 0x68, 0xa7, 0x44, 0x8b, 0x6f, 0xbe, 0x2e, 0xba,
	0x9b, 0x77, 0x5e, 0xe3, 0x27, 0xf2, 0x8b, 0x12, 0xbd, 0x84, 0x87, 0x0b, 0x2b, 0x80, 0x9e, 0x2d,
	0xc9, 0x7e, 0xef, 0xb6, 0x2c, 0x4b, 0xbe, 0x00, 0xf3, 0x65, 0xc9, 0xef, 0xdf, 0x88, 0xc2, 0xe4,
	0xbf, 0x42, 0x67, 0x71, 0x23, 0x96, 0x0d, 0xa6, 0x60, 0x7b, 0x8a, 0xd2, 0x1f, 0xee, 0xff, 0xf2,
	0xe9, 0xff, 0xfd, 0x2f, 0x70, 0x20, 0x7f, 0x2f, 0xeb, 0x2a, 0xd1, 0xde, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0x56, 0xec, 0x69, 0x4b, 0x0c, 0x00, 0x00,
}
