// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sysl.proto

package annotations

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type CallRule struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallRule) Reset()         { *m = CallRule{} }
func (m *CallRule) String() string { return proto.CompactTextString(m) }
func (*CallRule) ProtoMessage()    {}
func (*CallRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d185230f424090, []int{0}
}

func (m *CallRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRule.Unmarshal(m, b)
}
func (m *CallRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRule.Marshal(b, m, deterministic)
}
func (m *CallRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRule.Merge(m, src)
}
func (m *CallRule) XXX_Size() int {
	return xxx_messageInfo_CallRule.Size(m)
}
func (m *CallRule) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRule.DiscardUnknown(m)
}

var xxx_messageInfo_CallRule proto.InternalMessageInfo

func (m *CallRule) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *CallRule) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

var E_Calls = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: ([]*CallRule)(nil),
	Field:         72295728,
	Name:          "sysl.option.calls",
	Tag:           "bytes,72295728,rep,name=calls",
	Filename:      "sysl.proto",
}

func init() {
	proto.RegisterType((*CallRule)(nil), "sysl.option.CallRule")
	proto.RegisterExtension(E_Calls)
}

func init() { proto.RegisterFile("sysl.proto", fileDescriptor_a9d185230f424090) }

var fileDescriptor_a9d185230f424090 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xae, 0x2c, 0xce,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xb3, 0xf3, 0x0b, 0x4a, 0x32, 0xf3, 0xf3,
	0xa4, 0x14, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x52, 0x49, 0xa5, 0x69, 0xfa, 0x29,
	0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x25, 0xf9, 0x45, 0x10, 0xe5, 0x4a, 0x36, 0x5c, 0x1c, 0xce,
	0x89, 0x39, 0x39, 0x41, 0xa5, 0x39, 0xa9, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x18, 0x17, 0x5b, 0x6e,
	0x6a, 0x49, 0x46, 0x7e, 0x8a, 0x04, 0x13, 0x58, 0x02, 0xca, 0xb3, 0xf2, 0xe3, 0x62, 0x4d, 0x4e,
	0xcc, 0xc9, 0x29, 0x16, 0x92, 0xd3, 0x83, 0xd8, 0xa4, 0x07, 0xb3, 0x49, 0xcf, 0x17, 0xac, 0xc2,
	0x1f, 0xec, 0x8e, 0x62, 0x89, 0x0d, 0xa7, 0xf6, 0x28, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0xea,
	0x21, 0xb9, 0x4f, 0x0f, 0x66, 0x75, 0x10, 0xc4, 0x18, 0xa7, 0x3c, 0x2e, 0xbe, 0xe4, 0xfc, 0x5c,
	0x98, 0x59, 0x89, 0x05, 0x99, 0x4e, 0x02, 0x8e, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x60, 0x53, 0x02,
	0x40, 0x66, 0x07, 0x30, 0x46, 0x39, 0x42, 0xe5, 0xd3, 0xf3, 0x73, 0x12, 0xf3, 0xd2, 0xf5, 0xf2,
	0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0x36, 0xeb, 0x43, 0xa4, 0x12, 0x0b, 0x32, 0x8b, 0xf5,
	0x13, 0x0b, 0x32, 0xf5, 0x13, 0x11, 0xba, 0xad, 0x91, 0xd8, 0x8b, 0x98, 0x58, 0xdc, 0x1d, 0x03,
	0x3c, 0x93, 0xd8, 0xc0, 0x9a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0xb0, 0xd9, 0x34,
	0x41, 0x01, 0x00, 0x00,
}