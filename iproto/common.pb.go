// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package iproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "k8s.io/api/core/v1"
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

type ListPodsInNamespace struct {
	Namespace            string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPodsInNamespace) Reset()         { *m = ListPodsInNamespace{} }
func (m *ListPodsInNamespace) String() string { return proto.CompactTextString(m) }
func (*ListPodsInNamespace) ProtoMessage()    {}
func (*ListPodsInNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *ListPodsInNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPodsInNamespace.Unmarshal(m, b)
}
func (m *ListPodsInNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPodsInNamespace.Marshal(b, m, deterministic)
}
func (m *ListPodsInNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPodsInNamespace.Merge(m, src)
}
func (m *ListPodsInNamespace) XXX_Size() int {
	return xxx_messageInfo_ListPodsInNamespace.Size(m)
}
func (m *ListPodsInNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPodsInNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_ListPodsInNamespace proto.InternalMessageInfo

func (m *ListPodsInNamespace) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ListCurrentReleases struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCurrentReleases) Reset()         { *m = ListCurrentReleases{} }
func (m *ListCurrentReleases) String() string { return proto.CompactTextString(m) }
func (*ListCurrentReleases) ProtoMessage()    {}
func (*ListCurrentReleases) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *ListCurrentReleases) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCurrentReleases.Unmarshal(m, b)
}
func (m *ListCurrentReleases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCurrentReleases.Marshal(b, m, deterministic)
}
func (m *ListCurrentReleases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCurrentReleases.Merge(m, src)
}
func (m *ListCurrentReleases) XXX_Size() int {
	return xxx_messageInfo_ListCurrentReleases.Size(m)
}
func (m *ListCurrentReleases) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCurrentReleases.DiscardUnknown(m)
}

var xxx_messageInfo_ListCurrentReleases proto.InternalMessageInfo

type ServerToClientRequest struct {
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*ServerToClientRequest_ListPodsInNamespace
	//	*ServerToClientRequest_ListCurrentHelmReleases
	Request              isServerToClientRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ServerToClientRequest) Reset()         { *m = ServerToClientRequest{} }
func (m *ServerToClientRequest) String() string { return proto.CompactTextString(m) }
func (*ServerToClientRequest) ProtoMessage()    {}
func (*ServerToClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *ServerToClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerToClientRequest.Unmarshal(m, b)
}
func (m *ServerToClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerToClientRequest.Marshal(b, m, deterministic)
}
func (m *ServerToClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerToClientRequest.Merge(m, src)
}
func (m *ServerToClientRequest) XXX_Size() int {
	return xxx_messageInfo_ServerToClientRequest.Size(m)
}
func (m *ServerToClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerToClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerToClientRequest proto.InternalMessageInfo

func (m *ServerToClientRequest) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type isServerToClientRequest_Request interface {
	isServerToClientRequest_Request()
}

type ServerToClientRequest_ListPodsInNamespace struct {
	ListPodsInNamespace *ListPodsInNamespace `protobuf:"bytes,51,opt,name=list_pods_in_namespace,json=listPodsInNamespace,proto3,oneof"`
}

type ServerToClientRequest_ListCurrentHelmReleases struct {
	ListCurrentHelmReleases *ListCurrentReleases `protobuf:"bytes,52,opt,name=list_current_helm_releases,json=listCurrentHelmReleases,proto3,oneof"`
}

func (*ServerToClientRequest_ListPodsInNamespace) isServerToClientRequest_Request() {}

func (*ServerToClientRequest_ListCurrentHelmReleases) isServerToClientRequest_Request() {}

func (m *ServerToClientRequest) GetRequest() isServerToClientRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ServerToClientRequest) GetListPodsInNamespace() *ListPodsInNamespace {
	if x, ok := m.GetRequest().(*ServerToClientRequest_ListPodsInNamespace); ok {
		return x.ListPodsInNamespace
	}
	return nil
}

func (m *ServerToClientRequest) GetListCurrentHelmReleases() *ListCurrentReleases {
	if x, ok := m.GetRequest().(*ServerToClientRequest_ListCurrentHelmReleases); ok {
		return x.ListCurrentHelmReleases
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServerToClientRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServerToClientRequest_ListPodsInNamespace)(nil),
		(*ServerToClientRequest_ListCurrentHelmReleases)(nil),
	}
}

func init() {
	proto.RegisterType((*ListPodsInNamespace)(nil), "iproto.ListPodsInNamespace")
	proto.RegisterType((*ListCurrentReleases)(nil), "iproto.ListCurrentReleases")
	proto.RegisterType((*ServerToClientRequest)(nil), "iproto.ServerToClientRequest")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x14, 0xc4, 0xdb, 0x0e, 0x45, 0x31, 0x4c, 0xa9, 0x0a, 0x55, 0x40, 0xa2, 0xca, 0xd4, 0xc9, 0x16,
	0x84, 0x01, 0x46, 0xe8, 0x12, 0x24, 0x84, 0x90, 0xe9, 0xc4, 0x62, 0xb9, 0xc9, 0x23, 0x58, 0xf8,
	0x4f, 0xb0, 0x9d, 0x7e, 0x0b, 0xbe, 0x33, 0xc2, 0x6e, 0x14, 0x09, 0x3a, 0x59, 0xef, 0xdd, 0xf9,
	0x7c, 0x3f, 0xa3, 0x93, 0xca, 0x28, 0x65, 0x34, 0x6e, 0xad, 0xf1, 0x26, 0x9d, 0x8a, 0x70, 0x66,
	0x97, 0x8d, 0x31, 0x8d, 0x04, 0x12, 0xa6, 0x6d, 0xf7, 0x4e, 0xbc, 0x50, 0xe0, 0x3c, 0x57, 0x6d,
	0x34, 0x66, 0xf9, 0xe7, 0xad, 0xc3, 0xc2, 0x10, 0xde, 0x0a, 0x52, 0x19, 0x0b, 0x64, 0x77, 0x45,
	0x1a, 0xd0, 0x60, 0xb9, 0x87, 0x3a, 0x7a, 0xf2, 0x02, 0xcd, 0x9e, 0x84, 0xf3, 0x2f, 0xa6, 0x76,
	0x8f, 0xfa, 0x99, 0x2b, 0x70, 0x2d, 0xaf, 0x20, 0xbd, 0x40, 0x89, 0xee, 0x87, 0xc5, 0x78, 0x39,
	0x5e, 0x25, 0x74, 0x58, 0xe4, 0xf3, 0x78, 0x69, 0xdd, 0x59, 0x0b, 0xda, 0x53, 0x90, 0xc0, 0x1d,
	0xb8, 0xfc, 0x7b, 0x82, 0xe6, 0xaf, 0x60, 0x77, 0x60, 0x37, 0x66, 0x2d, 0x45, 0x90, 0xbe, 0x3a,
	0x70, 0x3e, 0xbd, 0x43, 0xa8, 0xb2, 0xf0, 0xfb, 0x2c, 0xe3, 0x7e, 0x31, 0x59, 0x8e, 0x57, 0xc7,
	0xd7, 0x19, 0x8e, 0xfd, 0x71, 0xdf, 0x1f, 0x6f, 0xfa, 0xfe, 0x34, 0xd9, 0xbb, 0xef, 0x7d, 0x4a,
	0xd1, 0xa9, 0x14, 0xce, 0xb3, 0xd6, 0xd4, 0x8e, 0x09, 0xcd, 0x86, 0x5a, 0x45, 0x88, 0x39, 0xc7,
	0xf1, 0x3b, 0xf0, 0x01, 0x8c, 0x72, 0x44, 0x67, 0xf2, 0x00, 0xdd, 0x1b, 0xca, 0x42, 0x66, 0x15,
	0x01, 0xd8, 0x07, 0x48, 0xc5, 0xec, 0x1e, 0x63, 0x71, 0xf3, 0x3f, 0xf7, 0x0f, 0x69, 0x39, 0xa2,
	0x67, 0x72, 0x58, 0x97, 0x20, 0x55, 0x2f, 0x3d, 0x24, 0xe8, 0xc8, 0x46, 0xea, 0xed, 0x34, 0x04,
	0x14, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x92, 0xc2, 0xa8, 0x91, 0xbf, 0x01, 0x00, 0x00,
}
