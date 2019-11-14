// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recipient.proto

package grpc

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

type Recipient struct {
	// Types that are valid to be assigned to Recipient:
	//	*Recipient_Address
	//	*Recipient_Alias
	Recipient            isRecipient_Recipient `protobuf_oneof:"recipient"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Recipient) Reset()         { *m = Recipient{} }
func (m *Recipient) String() string { return proto.CompactTextString(m) }
func (*Recipient) ProtoMessage()    {}
func (*Recipient) Descriptor() ([]byte, []int) {
	return fileDescriptor_72994ab5a87b4bee, []int{0}
}

func (m *Recipient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recipient.Unmarshal(m, b)
}
func (m *Recipient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recipient.Marshal(b, m, deterministic)
}
func (m *Recipient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recipient.Merge(m, src)
}
func (m *Recipient) XXX_Size() int {
	return xxx_messageInfo_Recipient.Size(m)
}
func (m *Recipient) XXX_DiscardUnknown() {
	xxx_messageInfo_Recipient.DiscardUnknown(m)
}

var xxx_messageInfo_Recipient proto.InternalMessageInfo

type isRecipient_Recipient interface {
	isRecipient_Recipient()
}

type Recipient_Address struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type Recipient_Alias struct {
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3,oneof"`
}

func (*Recipient_Address) isRecipient_Recipient() {}

func (*Recipient_Alias) isRecipient_Recipient() {}

func (m *Recipient) GetRecipient() isRecipient_Recipient {
	if m != nil {
		return m.Recipient
	}
	return nil
}

func (m *Recipient) GetAddress() []byte {
	if x, ok := m.GetRecipient().(*Recipient_Address); ok {
		return x.Address
	}
	return nil
}

func (m *Recipient) GetAlias() string {
	if x, ok := m.GetRecipient().(*Recipient_Alias); ok {
		return x.Alias
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Recipient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Recipient_Address)(nil),
		(*Recipient_Alias)(nil),
	}
}

func init() {
	proto.RegisterType((*Recipient)(nil), "waves.Recipient")
}

func init() { proto.RegisterFile("recipient.proto", fileDescriptor_72994ab5a87b4bee) }

var fileDescriptor_72994ab5a87b4bee = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x4d, 0xce,
	0x2c, 0xc8, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x4f, 0x2c,
	0x4b, 0x2d, 0x56, 0xf2, 0xe1, 0xe2, 0x0c, 0x82, 0xc9, 0x08, 0x49, 0x71, 0xb1, 0x27, 0xa6, 0xa4,
	0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x78, 0x30, 0x04, 0xc1, 0x04, 0x84,
	0xc4, 0xb8, 0x58, 0x13, 0x73, 0x32, 0x13, 0x8b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x3d, 0x18,
	0x82, 0x20, 0x5c, 0x27, 0x6e, 0x2e, 0x4e, 0xb8, 0xd1, 0x4e, 0x66, 0x5c, 0x6a, 0xc9, 0xf9, 0xb9,
	0x7a, 0x60, 0xa3, 0x0b, 0x72, 0x12, 0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0x21, 0xf6, 0x25, 0x95, 0xa6,
	0xe9, 0x95, 0x14, 0x25, 0xe6, 0x15, 0x27, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x45, 0xb1, 0xa4, 0x17,
	0x15, 0x24, 0xaf, 0x62, 0x62, 0x0d, 0x07, 0x29, 0x4d, 0x62, 0x03, 0xab, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x19, 0xf3, 0x13, 0xad, 0xa6, 0x00, 0x00, 0x00,
}