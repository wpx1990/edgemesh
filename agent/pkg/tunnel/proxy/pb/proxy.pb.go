// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proxy.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Proxy_Type int32

const (
	Proxy_CONNECT Proxy_Type = 0
	Proxy_SUCCESS Proxy_Type = 1
	Proxy_FAILED  Proxy_Type = 2
)

var Proxy_Type_name = map[int32]string{
	0: "CONNECT",
	1: "SUCCESS",
	2: "FAILED",
}

var Proxy_Type_value = map[string]int32{
	"CONNECT": 0,
	"SUCCESS": 1,
	"FAILED":  2,
}

func (x Proxy_Type) Enum() *Proxy_Type {
	p := new(Proxy_Type)
	*p = x
	return p
}

func (x Proxy_Type) String() string {
	return proto.EnumName(Proxy_Type_name, int32(x))
}

func (x *Proxy_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Proxy_Type_value, data, "Proxy_Type")
	if err != nil {
		return err
	}
	*x = Proxy_Type(value)
	return nil
}

func (Proxy_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0, 0}
}

type Proxy struct {
	Type                 *Proxy_Type `protobuf:"varint,1,opt,name=type,enum=pb.Proxy_Type" json:"type,omitempty"`
	Protocol             *string     `protobuf:"bytes,2,opt,name=protocol" json:"protocol,omitempty"`
	NodeName             *string     `protobuf:"bytes,3,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	Ip                   *string     `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	Port                 *int32      `protobuf:"varint,5,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Proxy) Reset()         { *m = Proxy{} }
func (m *Proxy) String() string { return proto.CompactTextString(m) }
func (*Proxy) ProtoMessage()    {}
func (*Proxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_700b50b08ed8dbaf, []int{0}
}

func (m *Proxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proxy.Unmarshal(m, b)
}
func (m *Proxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proxy.Marshal(b, m, deterministic)
}
func (m *Proxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proxy.Merge(m, src)
}
func (m *Proxy) XXX_Size() int {
	return xxx_messageInfo_Proxy.Size(m)
}
func (m *Proxy) XXX_DiscardUnknown() {
	xxx_messageInfo_Proxy.DiscardUnknown(m)
}

var xxx_messageInfo_Proxy proto.InternalMessageInfo

func (m *Proxy) GetType() Proxy_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Proxy_CONNECT
}

func (m *Proxy) GetProtocol() string {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return ""
}

func (m *Proxy) GetNodeName() string {
	if m != nil && m.NodeName != nil {
		return *m.NodeName
	}
	return ""
}

func (m *Proxy) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *Proxy) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Proxy_Type", Proxy_Type_name, Proxy_Type_value)
	proto.RegisterType((*Proxy)(nil), "pb.Proxy")
}

func init() { proto.RegisterFile("proxy.proto", fileDescriptor_700b50b08ed8dbaf) }

var fileDescriptor_700b50b08ed8dbaf = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0xaf,
	0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xda, 0xc6, 0xc8, 0xc5,
	0x1a, 0x00, 0x12, 0x13, 0x52, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x33, 0xe2, 0xd3, 0x2b, 0x48, 0xd2, 0x03, 0x4b, 0xe8, 0x85, 0x54, 0x16, 0xa4, 0x06, 0x81,
	0xe5, 0x84, 0xa4, 0xb8, 0x38, 0xc0, 0x5a, 0x93, 0xf3, 0x73, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38,
	0x83, 0xe0, 0x7c, 0x21, 0x69, 0x2e, 0xce, 0xbc, 0xfc, 0x94, 0xd4, 0xf8, 0xbc, 0xc4, 0xdc, 0x54,
	0x09, 0x66, 0x88, 0x24, 0x48, 0xc0, 0x2f, 0x31, 0x37, 0x55, 0x88, 0x8f, 0x8b, 0x29, 0xb3, 0x40,
	0x82, 0x05, 0x2c, 0xca, 0x94, 0x59, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x90, 0x5f, 0x54, 0x22, 0xc1,
	0xaa, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x2b, 0xe9, 0x70, 0xb1, 0x80, 0xac, 0x12, 0xe2, 0xe6,
	0x62, 0x77, 0xf6, 0xf7, 0xf3, 0x73, 0x75, 0x0e, 0x11, 0x60, 0x00, 0x71, 0x82, 0x43, 0x9d, 0x9d,
	0x5d, 0x83, 0x83, 0x05, 0x18, 0x85, 0xb8, 0xb8, 0xd8, 0xdc, 0x1c, 0x3d, 0x7d, 0x5c, 0x5d, 0x04,
	0x98, 0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x14, 0xb8, 0x1d, 0xca, 0x00, 0x00, 0x00,
}