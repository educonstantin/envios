// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel.proto

package envios_vessel

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{0}
}

func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (m *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(m, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{1}
}

func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (m *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(m, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created              bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c0e45ee319d513, []int{2}
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

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "envios.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "envios.vessel.Specification")
	proto.RegisterType((*Response)(nil), "envios.vessel.Response")
}

func init() { proto.RegisterFile("vessel.proto", fileDescriptor_e0c0e45ee319d513) }

var fileDescriptor_e0c0e45ee319d513 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0xb4, 0x4d, 0xd3, 0xd1, 0x78, 0x18, 0x10, 0xd7, 0x52, 0x21, 0xe4, 0xd4, 0x8b,
	0x11, 0xea, 0xcd, 0x9b, 0x08, 0xa2, 0x1e, 0x53, 0xd0, 0x63, 0xd9, 0x26, 0xa3, 0x0e, 0xb4, 0xbb,
	0x21, 0x1b, 0xd2, 0x7a, 0xf7, 0x2b, 0x78, 0xf7, 0xa3, 0x0a, 0x1b, 0x5b, 0x69, 0xd0, 0xde, 0xe6,
	0xcf, 0x9b, 0x97, 0x97, 0x1f, 0x0b, 0x47, 0x35, 0x59, 0x4b, 0x8b, 0xa4, 0x28, 0x4d, 0x65, 0x30,
	0x24, 0x5d, 0xb3, 0xb1, 0x49, 0x33, 0x8c, 0xbf, 0x04, 0xf8, 0x4f, 0xae, 0xc4, 0x63, 0xf0, 0x38,
	0x97, 0x22, 0x12, 0xe3, 0x41, 0xea, 0x71, 0x8e, 0x43, 0x08, 0x32, 0x55, 0xa8, 0x8c, 0xab, 0x77,
	0xe9, 0x45, 0x62, 0xdc, 0x4b, 0xb7, 0x3d, 0x9e, 0x03, 0x2c, 0xd5, 0x7a, 0xb6, 0x22, 0x7e, 0x7d,
	0xab, 0x64, 0xc7, 0x6d, 0x07, 0x4b, 0xb5, 0x7e, 0x76, 0x03, 0x44, 0xe8, 0x6a, 0xb5, 0x24, 0xd9,
	0x75, 0x66, 0xae, 0xc6, 0x11, 0x0c, 0x54, 0xad, 0x78, 0xa1, 0xe6, 0x0b, 0x92, 0xbd, 0x48, 0x8c,
	0x83, 0xf4, 0x77, 0x80, 0x67, 0x10, 0x98, 0x95, 0xa6, 0x72, 0xc6, 0xb9, 0xf4, 0xdd, 0x55, 0xdf,
	0xf5, 0x0f, 0x79, 0xfc, 0x08, 0xe1, 0xb4, 0xa0, 0x8c, 0x5f, 0x38, 0x53, 0x15, 0x1b, 0xbd, 0x13,
	0x4c, 0xec, 0x0d, 0xe6, 0xb5, 0x82, 0xc5, 0x1f, 0x02, 0x82, 0x94, 0x6c, 0x61, 0xb4, 0x25, 0xbc,
	0x00, 0xbf, 0xa1, 0xe0, 0x5c, 0x0e, 0x27, 0x27, 0xc9, 0x0e, 0x9b, 0xa4, 0xe1, 0x92, 0xfe, 0x88,
	0xf0, 0x12, 0xfa, 0x4d, 0x65, 0xa5, 0x17, 0x75, 0xfe, 0xd7, 0x6f, 0x54, 0x28, 0xa1, 0x9f, 0x95,
	0xa4, 0x2a, 0xca, 0x1d, 0xa1, 0x20, 0xdd, 0xb4, 0x93, 0x4f, 0x01, 0x61, 0xa3, 0x9e, 0x52, 0x59,
	0x73, 0x46, 0x78, 0x0f, 0xe1, 0x1d, 0xeb, 0xfc, 0x66, 0x0b, 0x64, 0xd4, 0x32, 0xdf, 0x41, 0x30,
	0x3c, 0x6d, 0x6d, 0x37, 0xff, 0x14, 0x1f, 0xe0, 0x35, 0xf8, 0xb7, 0xee, 0x33, 0xf8, 0x77, 0xbe,
	0x3d, 0xb7, 0x73, 0xdf, 0xbd, 0x91, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x3a, 0x8f,
	0xf5, 0x33, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "envios.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}
