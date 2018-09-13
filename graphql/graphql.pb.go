// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graphql/graphql.proto

package graphql // import "github.com/srikrsna/protoc-gen-graphql/graphql"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FieldOptions struct {
	Key                  bool     `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Ignore               bool     `protobuf:"varint,2,opt,name=ignore" json:"ignore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldOptions) Reset()         { *m = FieldOptions{} }
func (m *FieldOptions) String() string { return proto.CompactTextString(m) }
func (*FieldOptions) ProtoMessage()    {}
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_graphql_b9faf4912ffdc5eb, []int{0}
}
func (m *FieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldOptions.Unmarshal(m, b)
}
func (m *FieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldOptions.Marshal(b, m, deterministic)
}
func (dst *FieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldOptions.Merge(dst, src)
}
func (m *FieldOptions) XXX_Size() int {
	return xxx_messageInfo_FieldOptions.Size(m)
}
func (m *FieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_FieldOptions proto.InternalMessageInfo

func (m *FieldOptions) GetKey() bool {
	if m != nil {
		return m.Key
	}
	return false
}

func (m *FieldOptions) GetIgnore() bool {
	if m != nil {
		return m.Ignore
	}
	return false
}

type MethodOptions struct {
	// Types that are valid to be assigned to Type:
	//	*MethodOptions_Query
	//	*MethodOptions_Mutation
	Type                 isMethodOptions_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MethodOptions) Reset()         { *m = MethodOptions{} }
func (m *MethodOptions) String() string { return proto.CompactTextString(m) }
func (*MethodOptions) ProtoMessage()    {}
func (*MethodOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_graphql_b9faf4912ffdc5eb, []int{1}
}
func (m *MethodOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodOptions.Unmarshal(m, b)
}
func (m *MethodOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodOptions.Marshal(b, m, deterministic)
}
func (dst *MethodOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodOptions.Merge(dst, src)
}
func (m *MethodOptions) XXX_Size() int {
	return xxx_messageInfo_MethodOptions.Size(m)
}
func (m *MethodOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodOptions.DiscardUnknown(m)
}

var xxx_messageInfo_MethodOptions proto.InternalMessageInfo

type isMethodOptions_Type interface {
	isMethodOptions_Type()
}

type MethodOptions_Query struct {
	Query string `protobuf:"bytes,1,opt,name=query,oneof"`
}
type MethodOptions_Mutation struct {
	Mutation string `protobuf:"bytes,2,opt,name=mutation,oneof"`
}

func (*MethodOptions_Query) isMethodOptions_Type()    {}
func (*MethodOptions_Mutation) isMethodOptions_Type() {}

func (m *MethodOptions) GetType() isMethodOptions_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MethodOptions) GetQuery() string {
	if x, ok := m.GetType().(*MethodOptions_Query); ok {
		return x.Query
	}
	return ""
}

func (m *MethodOptions) GetMutation() string {
	if x, ok := m.GetType().(*MethodOptions_Mutation); ok {
		return x.Mutation
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MethodOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MethodOptions_OneofMarshaler, _MethodOptions_OneofUnmarshaler, _MethodOptions_OneofSizer, []interface{}{
		(*MethodOptions_Query)(nil),
		(*MethodOptions_Mutation)(nil),
	}
}

func _MethodOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MethodOptions)
	// type
	switch x := m.Type.(type) {
	case *MethodOptions_Query:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Query)
	case *MethodOptions_Mutation:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Mutation)
	case nil:
	default:
		return fmt.Errorf("MethodOptions.Type has unexpected type %T", x)
	}
	return nil
}

func _MethodOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MethodOptions)
	switch tag {
	case 1: // type.query
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &MethodOptions_Query{x}
		return true, err
	case 2: // type.mutation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &MethodOptions_Mutation{x}
		return true, err
	default:
		return false, nil
	}
}

func _MethodOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MethodOptions)
	// type
	switch x := m.Type.(type) {
	case *MethodOptions_Query:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Query)))
		n += len(x.Query)
	case *MethodOptions_Mutation:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Mutation)))
		n += len(x.Mutation)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_GraphqlFields = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldOptions)(nil),
	Field:         9892748,
	Name:          "graphql.graphql_fields",
	Tag:           "bytes,9892748,opt,name=graphql_fields,json=graphqlFields",
	Filename:      "graphql/graphql.proto",
}

var E_GraphqlMethods = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*MethodOptions)(nil),
	Field:         9892748,
	Name:          "graphql.graphql_methods",
	Tag:           "bytes,9892748,opt,name=graphql_methods,json=graphqlMethods",
	Filename:      "graphql/graphql.proto",
}

func init() {
	proto.RegisterType((*FieldOptions)(nil), "graphql.FieldOptions")
	proto.RegisterType((*MethodOptions)(nil), "graphql.MethodOptions")
	proto.RegisterExtension(E_GraphqlFields)
	proto.RegisterExtension(E_GraphqlMethods)
}

func init() { proto.RegisterFile("graphql/graphql.proto", fileDescriptor_graphql_b9faf4912ffdc5eb) }

var fileDescriptor_graphql_b9faf4912ffdc5eb = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2f, 0x4a, 0x2c,
	0xc8, 0x28, 0xcc, 0xd1, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x42, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x38, 0xa9, 0x34, 0x4d, 0x3f, 0x25,
	0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0xbf, 0x08, 0xa2, 0x54, 0xc9, 0x82, 0x8b, 0xc7, 0x2d,
	0x33, 0x35, 0x27, 0xc5, 0xbf, 0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x48, 0x80, 0x8b, 0x39, 0x3b,
	0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc4, 0x14, 0x12, 0xe3, 0x62, 0xcb, 0x4c,
	0xcf, 0xcb, 0x2f, 0x4a, 0x95, 0x60, 0x02, 0x0b, 0x42, 0x79, 0x4a, 0xbe, 0x5c, 0xbc, 0xbe, 0xa9,
	0x25, 0x19, 0xf9, 0x70, 0xad, 0x62, 0x5c, 0xac, 0x85, 0xa5, 0xa9, 0x45, 0x10, 0xcd, 0x9c, 0x1e,
	0x0c, 0x41, 0x10, 0xae, 0x90, 0x0c, 0x17, 0x47, 0x6e, 0x69, 0x49, 0x22, 0x48, 0x11, 0xd8, 0x08,
	0x90, 0x14, 0x5c, 0xc4, 0x89, 0x8d, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0xd5, 0x2a, 0x8e, 0x8b, 0x0f,
	0xea, 0xea, 0xf8, 0x34, 0x90, 0x83, 0x8a, 0x85, 0x64, 0xf5, 0x20, 0xae, 0xd7, 0x83, 0xb9, 0x5e,
	0x0f, 0xd9, 0xa5, 0x12, 0x3d, 0xcf, 0x6f, 0xb3, 0x28, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xea, 0xc1,
	0x7c, 0x8f, 0x2c, 0x1f, 0xc4, 0x0b, 0x15, 0x05, 0x0b, 0x16, 0x5b, 0x25, 0x71, 0xf1, 0xc3, 0xcc,
	0xcf, 0x05, 0x3b, 0xbb, 0x58, 0x48, 0x0e, 0xc3, 0x02, 0x14, 0x0f, 0x21, 0x6c, 0x10, 0x83, 0xdb,
	0x80, 0xa2, 0x20, 0x08, 0xe6, 0x62, 0x88, 0x68, 0xb1, 0x93, 0x41, 0x94, 0x5e, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x51, 0x66, 0x76, 0x51, 0x71, 0x5e, 0x22,
	0x24, 0xf4, 0x93, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd1, 0xe2, 0x2b, 0x89, 0x0d, 0x2c, 0x67, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x13, 0xc7, 0x51, 0x97, 0xc9, 0x01, 0x00, 0x00,
}