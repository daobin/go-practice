// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package main

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

type SearchReq_State int32

const (
	SearchReq_Unknown SearchReq_State = 0
	SearchReq_Enable  SearchReq_State = 1
	SearchReq_Disable SearchReq_State = 2
	SearchReq_Deleted SearchReq_State = 3
)

var SearchReq_State_name = map[int32]string{
	0: "Unknown",
	1: "Enable",
	2: "Disable",
	3: "Deleted",
}

var SearchReq_State_value = map[string]int32{
	"Unknown": 0,
	"Enable":  1,
	"Disable": 2,
	"Deleted": 3,
}

func (x SearchReq_State) String() string {
	return proto.EnumName(SearchReq_State_name, int32(x))
}

func (SearchReq_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0, 0}
}

// 查询请求
type SearchReq struct {
	Query                string          `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageIndex            int32           `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize             int32           `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	State                SearchReq_State `protobuf:"varint,4,opt,name=state,proto3,enum=main.SearchReq_State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SearchReq) Reset()         { *m = SearchReq{} }
func (m *SearchReq) String() string { return proto.CompactTextString(m) }
func (*SearchReq) ProtoMessage()    {}
func (*SearchReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *SearchReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchReq.Unmarshal(m, b)
}
func (m *SearchReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchReq.Marshal(b, m, deterministic)
}
func (m *SearchReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchReq.Merge(m, src)
}
func (m *SearchReq) XXX_Size() int {
	return xxx_messageInfo_SearchReq.Size(m)
}
func (m *SearchReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchReq.DiscardUnknown(m)
}

var xxx_messageInfo_SearchReq proto.InternalMessageInfo

func (m *SearchReq) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchReq) GetPageIndex() int32 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *SearchReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *SearchReq) GetState() SearchReq_State {
	if m != nil {
		return m.State
	}
	return SearchReq_Unknown
}

// 查询响应
type SearchRes struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRes) Reset()         { *m = SearchRes{} }
func (m *SearchRes) String() string { return proto.CompactTextString(m) }
func (*SearchRes) ProtoMessage()    {}
func (*SearchRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *SearchRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRes.Unmarshal(m, b)
}
func (m *SearchRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRes.Marshal(b, m, deterministic)
}
func (m *SearchRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRes.Merge(m, src)
}
func (m *SearchRes) XXX_Size() int {
	return xxx_messageInfo_SearchRes.Size(m)
}
func (m *SearchRes) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRes.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRes proto.InternalMessageInfo

func (m *SearchRes) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func init() {
	proto.RegisterEnum("main.SearchReq_State", SearchReq_State_name, SearchReq_State_value)
	proto.RegisterType((*SearchReq)(nil), "main.SearchReq")
	proto.RegisterType((*SearchRes)(nil), "main.SearchRes")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2c,
	0x4a, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x3a,
	0xc2, 0xc8, 0xc5, 0x19, 0x0c, 0x16, 0x0e, 0x4a, 0x2d, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x2c, 0x4d,
	0x2d, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x64, 0xb8, 0x38, 0x0b,
	0x12, 0xd3, 0x53, 0x3d, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0x10,
	0x02, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0x4e, 0x70, 0x66, 0x55, 0xaa, 0x04, 0x33, 0x58, 0x12, 0xce,
	0x17, 0xd2, 0xe6, 0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x49, 0x95, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x33,
	0x12, 0xd5, 0x03, 0xd9, 0xa9, 0x07, 0xb7, 0x4f, 0x2f, 0x18, 0x24, 0x19, 0x04, 0x51, 0xa3, 0x64,
	0xc5, 0xc5, 0x0a, 0xe6, 0x0b, 0x71, 0x73, 0xb1, 0x87, 0xe6, 0x65, 0xe7, 0xe5, 0x97, 0xe7, 0x09,
	0x30, 0x08, 0x71, 0x71, 0xb1, 0xb9, 0xe6, 0x25, 0x26, 0xe5, 0xa4, 0x0a, 0x30, 0x82, 0x24, 0x5c,
	0x32, 0x8b, 0xc1, 0x1c, 0x26, 0x30, 0x27, 0x35, 0x27, 0xb5, 0x24, 0x35, 0x45, 0x80, 0x59, 0x49,
	0x11, 0xe1, 0x8b, 0x62, 0xec, 0xbe, 0x48, 0x62, 0x03, 0x7b, 0xdb, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x47, 0x18, 0xd5, 0x4d, 0x06, 0x01, 0x00, 0x00,
}