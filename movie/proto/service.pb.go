// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

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

type MovieData struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieData) Reset()         { *m = MovieData{} }
func (m *MovieData) String() string { return proto.CompactTextString(m) }
func (*MovieData) ProtoMessage()    {}
func (*MovieData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *MovieData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieData.Unmarshal(m, b)
}
func (m *MovieData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieData.Marshal(b, m, deterministic)
}
func (m *MovieData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieData.Merge(m, src)
}
func (m *MovieData) XXX_Size() int {
	return xxx_messageInfo_MovieData.Size(m)
}
func (m *MovieData) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieData.DiscardUnknown(m)
}

var xxx_messageInfo_MovieData proto.InternalMessageInfo

func (m *MovieData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CreateRequest struct {
	Data                 *MovieData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetData() *MovieData {
	if m != nil {
		return m.Data
	}
	return nil
}

type CreateResponse struct {
	CreatedId            int64    `protobuf:"varint,2,opt,name=createdId,proto3" json:"createdId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetCreatedId() int64 {
	if m != nil {
		return m.CreatedId
	}
	return 0
}

type ReadRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadResponse struct {
	Id                   int64      `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *MovieData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadResponse) GetData() *MovieData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReadAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

type ReadAllResponse struct {
	Ids                  []int64      `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Dates                []*MovieData `protobuf:"bytes,3,rep,name=dates,proto3" json:"dates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ReadAllResponse) GetDates() []*MovieData {
	if m != nil {
		return m.Dates
	}
	return nil
}

type UpdateRequest struct {
	Id                   int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *MovieData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRequest) GetData() *MovieData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MovieData)(nil), "proto.MovieData")
	proto.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "proto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "proto.ReadResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "proto.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "proto.ReadAllResponse")
	proto.RegisterType((*UpdateRequest)(nil), "proto.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "proto.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0xdd, 0x92, 0x75, 0xb2, 0x3b, 0x3b, 0x47, 0xdc, 0xa4, 0x0c, 0xc5, 0x1a, 0x44, 0xf6, 0x34,
	0x70, 0x22, 0x03, 0xdf, 0xc4, 0xf9, 0x20, 0xe2, 0x4b, 0xc0, 0x1f, 0x10, 0x97, 0xfb, 0x50, 0x28,
	0x76, 0x2e, 0x71, 0xff, 0xc7, 0x7f, 0x2a, 0x4d, 0xd2, 0x6c, 0x2d, 0x7e, 0x3c, 0xb5, 0x39, 0xf7,
	0x9e, 0x73, 0xcf, 0x3d, 0x17, 0x62, 0x8d, 0x9b, 0x6d, 0xb6, 0xc2, 0xd9, 0x7a, 0x53, 0x98, 0x82,
	0x45, 0xf6, 0xc3, 0x2f, 0xa0, 0xf7, 0x52, 0x6c, 0x33, 0x5c, 0x4a, 0x23, 0xd9, 0x08, 0x22, 0x93,
	0x99, 0x1c, 0x93, 0x76, 0xda, 0x9e, 0xf6, 0x84, 0x7b, 0xf0, 0x5b, 0x88, 0x1f, 0x36, 0x28, 0x0d,
	0x0a, 0xfc, 0xf8, 0x44, 0x6d, 0xd8, 0x25, 0x74, 0x94, 0x34, 0xd2, 0x76, 0xf5, 0xe7, 0x43, 0x27,
	0x38, 0x0b, 0x32, 0xc2, 0x56, 0xf9, 0x0c, 0x06, 0x15, 0x4d, 0xaf, 0x8b, 0x77, 0x8d, 0xec, 0x14,
	0x7a, 0x2b, 0x8b, 0xa8, 0x27, 0x95, 0x90, 0xb4, 0x3d, 0xa5, 0x62, 0x07, 0xf0, 0x33, 0xe8, 0x0b,
	0x94, 0xaa, 0x1a, 0x32, 0x00, 0x92, 0x29, 0x3b, 0x82, 0x0a, 0x92, 0x29, 0xbe, 0x84, 0x43, 0x57,
	0xf6, 0x62, 0xae, 0x4e, 0xaa, 0x7a, 0x30, 0x45, 0xff, 0x34, 0x35, 0x84, 0x41, 0xa9, 0x72, 0x9f,
	0xe7, 0x7e, 0x0e, 0x7f, 0x86, 0xa3, 0x80, 0x78, 0xe9, 0x21, 0xd0, 0x4c, 0xe9, 0x84, 0xa4, 0x74,
	0x4a, 0x45, 0xf9, 0xcb, 0xae, 0x20, 0x52, 0xd2, 0xa0, 0x4e, 0x68, 0x4a, 0x7f, 0x54, 0x77, 0x65,
	0xfe, 0x08, 0xf1, 0xeb, 0x5a, 0xed, 0x45, 0xd5, 0xd8, 0x22, 0xb8, 0x24, 0xff, 0xb9, 0xac, 0x64,
	0x9c, 0x25, 0x7e, 0x0e, 0xf1, 0x12, 0x73, 0xfc, 0x55, 0xb8, 0xa4, 0x54, 0x0d, 0x8e, 0x32, 0xff,
	0x22, 0x10, 0x59, 0x61, 0xb6, 0x80, 0xae, 0xbb, 0x04, 0x1b, 0xf9, 0x81, 0xb5, 0x7b, 0x4e, 0xc6,
	0x0d, 0xd4, 0xcf, 0x6c, 0xb1, 0x6b, 0xe8, 0x94, 0xd9, 0x30, 0xe6, 0x1b, 0xf6, 0xee, 0x33, 0x39,
	0xae, 0x61, 0x81, 0x72, 0x07, 0x07, 0x3e, 0x4e, 0x36, 0xde, 0xeb, 0xd8, 0x05, 0x3e, 0x39, 0x69,
	0xc2, 0x81, 0xbb, 0x80, 0xae, 0x5b, 0x3b, 0xf8, 0xac, 0x85, 0x19, 0x7c, 0x36, 0xb2, 0xb1, 0x44,
	0xb7, 0x7c, 0x20, 0xd6, 0xc2, 0x0a, 0xc4, 0x7a, 0x42, 0xbc, 0xf5, 0xd6, 0xb5, 0xf8, 0xcd, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xd2, 0x2b, 0x8f, 0x1c, 0x03, 0x00, 0x00,
}