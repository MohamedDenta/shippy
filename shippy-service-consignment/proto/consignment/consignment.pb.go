// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool         `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "consignment.Consignment")
	proto.RegisterType((*Container)(nil), "consignment.Container")
	proto.RegisterType((*Response)(nil), "consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4f, 0x4b, 0x3b, 0x31,
	0x10, 0xfd, 0x6d, 0xff, 0xef, 0x04, 0x7e, 0x62, 0xc0, 0x1a, 0xf4, 0xe0, 0xb2, 0x5e, 0x7a, 0xaa,
	0x50, 0xc1, 0x83, 0xd7, 0x82, 0xd0, 0x6b, 0x7a, 0xf2, 0xa4, 0x35, 0x19, 0xb6, 0x03, 0x36, 0x59,
	0x92, 0xb4, 0x7e, 0x29, 0x3f, 0xa4, 0x6c, 0xb6, 0xab, 0x51, 0xf1, 0x36, 0xef, 0xbd, 0x99, 0x79,
	0x2f, 0x19, 0xb8, 0xae, 0x9d, 0x0d, 0xf6, 0x46, 0x59, 0xe3, 0xa9, 0x32, 0x3b, 0x34, 0x21, 0xad,
	0xe7, 0x51, 0xe5, 0x2c, 0xa1, 0xca, 0xf7, 0x0c, 0xd8, 0xf2, 0x0b, 0xf3, 0xff, 0xd0, 0x23, 0x2d,
	0xb2, 0x22, 0x9b, 0xe5, 0xb2, 0x47, 0x9a, 0x17, 0xc0, 0x34, 0x7a, 0xe5, 0xa8, 0x0e, 0x64, 0x8d,
	0xe8, 0x45, 0x21, 0xa5, 0xf8, 0x14, 0x46, 0x6f, 0x48, 0xd5, 0x36, 0x88, 0x7e, 0x91, 0xcd, 0x86,
	0xf2, 0x88, 0xf8, 0x1d, 0x80, 0xb2, 0x26, 0x6c, 0xc8, 0xa0, 0xf3, 0x62, 0x50, 0xf4, 0x67, 0x6c,
	0x31, 0x9d, 0xa7, 0x71, 0x96, 0x9d, 0x2c, 0x93, 0x4e, 0x7e, 0x09, 0xf9, 0x01, 0xbd, 0xc7, 0xd7,
	0x27, 0xd2, 0x62, 0x18, 0xfd, 0x26, 0x2d, 0xb1, 0xd2, 0xe5, 0x0e, 0xf2, 0xcf, 0xa9, 0x5f, 0x59,
	0xaf, 0x80, 0xa9, 0xbd, 0x0f, 0x76, 0x87, 0xae, 0x99, 0x6d, 0xb3, 0x42, 0x47, 0xad, 0x74, 0x13,
	0xd5, 0x3a, 0xaa, 0xc8, 0xc4, 0xa8, 0xb9, 0x3c, 0x22, 0x7e, 0x0e, 0xe3, 0xbd, 0x6f, 0x87, 0x06,
	0xad, 0xd0, 0xc0, 0x95, 0x2e, 0x9f, 0x61, 0x22, 0xd1, 0xd7, 0xd6, 0x78, 0xe4, 0x02, 0xc6, 0xca,
	0xe1, 0x26, 0x60, 0x6b, 0x39, 0x91, 0x1d, 0xe4, 0xf7, 0x90, 0x7e, 0x69, 0xf4, 0x65, 0x0b, 0xf1,
	0xf3, 0xa9, 0x5d, 0x2d, 0xd3, 0xe6, 0xc5, 0x23, 0x9c, 0xac, 0xb7, 0x54, 0xd7, 0x64, 0xaa, 0x35,
	0xba, 0x03, 0x29, 0xe4, 0x0f, 0x70, 0xba, 0x8c, 0x9b, 0xd3, 0xbb, 0xfc, 0xb9, 0xee, 0xe2, 0xec,
	0x9b, 0xd2, 0xc5, 0x2d, 0xff, 0xbd, 0x8c, 0xe2, 0xb9, 0x6f, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0xef, 0xe4, 0xb7, 0x15, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/consignment.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
}

// UnimplementedShippingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShippingServiceServer struct {
}

func (*UnimplementedShippingServiceServer) CreateConsignment(ctx context.Context, req *Consignment) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsignment not implemented")
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}
