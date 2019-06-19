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

type CinemaData struct {
	// Name of the cinema.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Cinema Id
	Id int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Data for occupation of seats
	Seats                []*SeatData `protobuf:"bytes,3,rep,name=seats,proto3" json:"seats,omitempty"`
	RowCount             int64       `protobuf:"varint,4,opt,name=rowCount,proto3" json:"rowCount,omitempty"`
	SeatCount            int64       `protobuf:"varint,5,opt,name=seatCount,proto3" json:"seatCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CinemaData) Reset()         { *m = CinemaData{} }
func (m *CinemaData) String() string { return proto.CompactTextString(m) }
func (*CinemaData) ProtoMessage()    {}
func (*CinemaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *CinemaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CinemaData.Unmarshal(m, b)
}
func (m *CinemaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CinemaData.Marshal(b, m, deterministic)
}
func (m *CinemaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CinemaData.Merge(m, src)
}
func (m *CinemaData) XXX_Size() int {
	return xxx_messageInfo_CinemaData.Size(m)
}
func (m *CinemaData) XXX_DiscardUnknown() {
	xxx_messageInfo_CinemaData.DiscardUnknown(m)
}

var xxx_messageInfo_CinemaData proto.InternalMessageInfo

func (m *CinemaData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CinemaData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CinemaData) GetSeats() []*SeatData {
	if m != nil {
		return m.Seats
	}
	return nil
}

func (m *CinemaData) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *CinemaData) GetSeatCount() int64 {
	if m != nil {
		return m.SeatCount
	}
	return 0
}

type SeatData struct {
	Row                  int64    `protobuf:"varint,4,opt,name=row,proto3" json:"row,omitempty"`
	Seat                 int64    `protobuf:"varint,5,opt,name=seat,proto3" json:"seat,omitempty"`
	Occupied             bool     `protobuf:"varint,6,opt,name=occupied,proto3" json:"occupied,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SeatData) Reset()         { *m = SeatData{} }
func (m *SeatData) String() string { return proto.CompactTextString(m) }
func (*SeatData) ProtoMessage()    {}
func (*SeatData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *SeatData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SeatData.Unmarshal(m, b)
}
func (m *SeatData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SeatData.Marshal(b, m, deterministic)
}
func (m *SeatData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeatData.Merge(m, src)
}
func (m *SeatData) XXX_Size() int {
	return xxx_messageInfo_SeatData.Size(m)
}
func (m *SeatData) XXX_DiscardUnknown() {
	xxx_messageInfo_SeatData.DiscardUnknown(m)
}

var xxx_messageInfo_SeatData proto.InternalMessageInfo

func (m *SeatData) GetRow() int64 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *SeatData) GetSeat() int64 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func (m *SeatData) GetOccupied() bool {
	if m != nil {
		return m.Occupied
	}
	return false
}

type CreateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Row                  int64    `protobuf:"varint,4,opt,name=row,proto3" json:"row,omitempty"`
	Seats                int64    `protobuf:"varint,5,opt,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
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

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetRow() int64 {
	if m != nil {
		return m.Row
	}
	return 0
}

func (m *CreateRequest) GetSeats() int64 {
	if m != nil {
		return m.Seats
	}
	return 0
}

type CreateResponse struct {
	Data                 *CinemaData `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
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

func (m *CreateResponse) GetData() *CinemaData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteRequest struct {
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
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
	Success              bool     `protobuf:"varint,8,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
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

func (m *DeleteResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ReadRequest struct {
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
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
	Success              bool        `protobuf:"varint,8,opt,name=success,proto3" json:"success,omitempty"`
	Data                 *CinemaData `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
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

func (m *ReadResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ReadResponse) GetData() *CinemaData {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Data                 []*CinemaData `protobuf:"bytes,9,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetData() []*CinemaData {
	if m != nil {
		return m.Data
	}
	return nil
}

type OccupiedRequest struct {
	Id                   int64       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Seats                []*SeatData `protobuf:"bytes,10,rep,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OccupiedRequest) Reset()         { *m = OccupiedRequest{} }
func (m *OccupiedRequest) String() string { return proto.CompactTextString(m) }
func (*OccupiedRequest) ProtoMessage()    {}
func (*OccupiedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{10}
}

func (m *OccupiedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OccupiedRequest.Unmarshal(m, b)
}
func (m *OccupiedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OccupiedRequest.Marshal(b, m, deterministic)
}
func (m *OccupiedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OccupiedRequest.Merge(m, src)
}
func (m *OccupiedRequest) XXX_Size() int {
	return xxx_messageInfo_OccupiedRequest.Size(m)
}
func (m *OccupiedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OccupiedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OccupiedRequest proto.InternalMessageInfo

func (m *OccupiedRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OccupiedRequest) GetSeats() []*SeatData {
	if m != nil {
		return m.Seats
	}
	return nil
}

type OccupiedResponse struct {
	Seats                []*SeatData `protobuf:"bytes,10,rep,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *OccupiedResponse) Reset()         { *m = OccupiedResponse{} }
func (m *OccupiedResponse) String() string { return proto.CompactTextString(m) }
func (*OccupiedResponse) ProtoMessage()    {}
func (*OccupiedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{11}
}

func (m *OccupiedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OccupiedResponse.Unmarshal(m, b)
}
func (m *OccupiedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OccupiedResponse.Marshal(b, m, deterministic)
}
func (m *OccupiedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OccupiedResponse.Merge(m, src)
}
func (m *OccupiedResponse) XXX_Size() int {
	return xxx_messageInfo_OccupiedResponse.Size(m)
}
func (m *OccupiedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OccupiedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OccupiedResponse proto.InternalMessageInfo

func (m *OccupiedResponse) GetSeats() []*SeatData {
	if m != nil {
		return m.Seats
	}
	return nil
}

type AvailableRequest struct {
	Id                   int64       `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Seats                []*SeatData `protobuf:"bytes,10,rep,name=seats,proto3" json:"seats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AvailableRequest) Reset()         { *m = AvailableRequest{} }
func (m *AvailableRequest) String() string { return proto.CompactTextString(m) }
func (*AvailableRequest) ProtoMessage()    {}
func (*AvailableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{12}
}

func (m *AvailableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableRequest.Unmarshal(m, b)
}
func (m *AvailableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableRequest.Marshal(b, m, deterministic)
}
func (m *AvailableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableRequest.Merge(m, src)
}
func (m *AvailableRequest) XXX_Size() int {
	return xxx_messageInfo_AvailableRequest.Size(m)
}
func (m *AvailableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableRequest proto.InternalMessageInfo

func (m *AvailableRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AvailableRequest) GetSeats() []*SeatData {
	if m != nil {
		return m.Seats
	}
	return nil
}

type AvailableResponse struct {
	Available            bool     `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvailableResponse) Reset()         { *m = AvailableResponse{} }
func (m *AvailableResponse) String() string { return proto.CompactTextString(m) }
func (*AvailableResponse) ProtoMessage()    {}
func (*AvailableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{13}
}

func (m *AvailableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvailableResponse.Unmarshal(m, b)
}
func (m *AvailableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvailableResponse.Marshal(b, m, deterministic)
}
func (m *AvailableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvailableResponse.Merge(m, src)
}
func (m *AvailableResponse) XXX_Size() int {
	return xxx_messageInfo_AvailableResponse.Size(m)
}
func (m *AvailableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvailableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvailableResponse proto.InternalMessageInfo

func (m *AvailableResponse) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

type ClearRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearRequest) Reset()         { *m = ClearRequest{} }
func (m *ClearRequest) String() string { return proto.CompactTextString(m) }
func (*ClearRequest) ProtoMessage()    {}
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{14}
}

func (m *ClearRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearRequest.Unmarshal(m, b)
}
func (m *ClearRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearRequest.Marshal(b, m, deterministic)
}
func (m *ClearRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearRequest.Merge(m, src)
}
func (m *ClearRequest) XXX_Size() int {
	return xxx_messageInfo_ClearRequest.Size(m)
}
func (m *ClearRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearRequest proto.InternalMessageInfo

type ClearResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearResponse) Reset()         { *m = ClearResponse{} }
func (m *ClearResponse) String() string { return proto.CompactTextString(m) }
func (*ClearResponse) ProtoMessage()    {}
func (*ClearResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{15}
}

func (m *ClearResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearResponse.Unmarshal(m, b)
}
func (m *ClearResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearResponse.Marshal(b, m, deterministic)
}
func (m *ClearResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearResponse.Merge(m, src)
}
func (m *ClearResponse) XXX_Size() int {
	return xxx_messageInfo_ClearResponse.Size(m)
}
func (m *ClearResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CinemaData)(nil), "proto.CinemaData")
	proto.RegisterType((*SeatData)(nil), "proto.SeatData")
	proto.RegisterType((*CreateRequest)(nil), "proto.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "proto.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "proto.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "proto.DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "proto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "proto.ReadResponse")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*OccupiedRequest)(nil), "proto.OccupiedRequest")
	proto.RegisterType((*OccupiedResponse)(nil), "proto.OccupiedResponse")
	proto.RegisterType((*AvailableRequest)(nil), "proto.AvailableRequest")
	proto.RegisterType((*AvailableResponse)(nil), "proto.AvailableResponse")
	proto.RegisterType((*ClearRequest)(nil), "proto.ClearRequest")
	proto.RegisterType((*ClearResponse)(nil), "proto.ClearResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x6d, 0x6b, 0xd3, 0x50,
	0x14, 0x6e, 0xda, 0x34, 0x4b, 0x4f, 0x9b, 0xb6, 0xbb, 0x56, 0x0d, 0x61, 0x62, 0xb8, 0x30, 0x08,
	0x7e, 0x18, 0xac, 0x2a, 0x63, 0xf8, 0x69, 0x64, 0x88, 0xe2, 0x60, 0x10, 0x7f, 0xc1, 0x5d, 0x72,
	0x3e, 0x04, 0xba, 0xa6, 0xe6, 0xa6, 0x1b, 0xfe, 0x0c, 0x7f, 0x9b, 0x7f, 0x48, 0x72, 0xdf, 0x92,
	0xa6, 0x56, 0x87, 0x7e, 0x4a, 0xce, 0x73, 0xce, 0xf3, 0xdc, 0xf3, 0xf2, 0x80, 0xc7, 0xb1, 0x7c,
	0xc8, 0x53, 0x3c, 0xdb, 0x94, 0x45, 0x55, 0x90, 0xa1, 0xf8, 0xd0, 0x1f, 0x16, 0x40, 0x9c, 0xaf,
	0xf1, 0x9e, 0x5d, 0xb3, 0x8a, 0x11, 0x02, 0xf6, 0x9a, 0xdd, 0xa3, 0x6f, 0x85, 0x56, 0x34, 0x4a,
	0xc4, 0x3f, 0x99, 0x42, 0x3f, 0xcf, 0xfc, 0x7e, 0x68, 0x45, 0x83, 0xa4, 0x9f, 0x67, 0xe4, 0x14,
	0x86, 0x1c, 0x59, 0xc5, 0xfd, 0x41, 0x38, 0x88, 0xc6, 0xcb, 0x99, 0x14, 0x3c, 0xfb, 0x8a, 0xac,
	0xaa, 0x35, 0x12, 0x99, 0x25, 0x01, 0xb8, 0x65, 0xf1, 0x18, 0x17, 0xdb, 0x75, 0xe5, 0xdb, 0x82,
	0x6c, 0x62, 0x72, 0x02, 0xa3, 0xba, 0x48, 0x26, 0x87, 0x22, 0xd9, 0x00, 0xf4, 0x06, 0x5c, 0x2d,
	0x46, 0xe6, 0x30, 0x28, 0x8b, 0x47, 0x25, 0x50, 0xff, 0xd6, 0x2d, 0xd6, 0xa5, 0x8a, 0x26, 0xfe,
	0xeb, 0xb7, 0x8a, 0x34, 0xdd, 0x6e, 0x72, 0xcc, 0x7c, 0x27, 0xb4, 0x22, 0x37, 0x31, 0x31, 0xfd,
	0x02, 0x5e, 0x5c, 0x22, 0xab, 0x30, 0xc1, 0x6f, 0x5b, 0xe4, 0xd5, 0x6f, 0x67, 0xdc, 0x7f, 0x66,
	0xa1, 0xa7, 0x94, 0xef, 0xc8, 0x80, 0x5e, 0xc0, 0x54, 0x8b, 0xf1, 0x4d, 0xb1, 0xe6, 0x48, 0x4e,
	0xc1, 0xce, 0x58, 0xc5, 0xfc, 0xa3, 0xd0, 0x8a, 0xc6, 0xcb, 0x63, 0xb5, 0x8c, 0x66, 0xa5, 0x89,
	0x48, 0xd3, 0xd7, 0xe0, 0x5d, 0xe3, 0x0a, 0x9b, 0x2e, 0x3a, 0x5b, 0xa5, 0x6f, 0x60, 0xaa, 0x0b,
	0x94, 0xb2, 0x0f, 0x47, 0x7c, 0x9b, 0xa6, 0xc8, 0xb9, 0xef, 0x8a, 0x99, 0x74, 0x48, 0x5f, 0xc1,
	0x38, 0x41, 0x96, 0x1d, 0x92, 0xba, 0x85, 0x89, 0x4c, 0xff, 0x4d, 0xe8, 0xa9, 0xcd, 0x7b, 0x30,
	0xbe, 0xc9, 0x79, 0xa5, 0xde, 0xa3, 0xef, 0x61, 0x22, 0xc3, 0xce, 0x0a, 0x46, 0xc2, 0x0f, 0x07,
	0x55, 0x3e, 0xc1, 0xec, 0x56, 0x1d, 0xe5, 0x40, 0xe7, 0x8d, 0xb5, 0xe0, 0x4f, 0xd6, 0xa2, 0x97,
	0x30, 0x6f, 0x94, 0x4c, 0x13, 0x4f, 0xa2, 0x7e, 0x86, 0xf9, 0xd5, 0x03, 0xcb, 0x57, 0xec, 0x6e,
	0x85, 0xff, 0xd9, 0xc5, 0x39, 0x1c, 0xb7, 0xa4, 0x54, 0x1b, 0x27, 0x30, 0x62, 0x1a, 0x14, 0x0e,
	0x73, 0x93, 0x06, 0xa0, 0x53, 0x98, 0xc4, 0x2b, 0x64, 0xa5, 0xde, 0xe4, 0x0c, 0x3c, 0x15, 0x4b,
	0xfa, 0xf2, 0xe7, 0x00, 0x1c, 0xb9, 0x38, 0x72, 0x01, 0x8e, 0xb4, 0x1a, 0x59, 0xe8, 0x8d, 0xb6,
	0x6d, 0x1c, 0x3c, 0xef, 0xa0, 0x52, 0x81, 0xf6, 0x6a, 0xa2, 0x74, 0x92, 0x21, 0xee, 0x38, 0xcf,
	0x10, 0x77, 0xed, 0x46, 0x7b, 0xe4, 0x1c, 0xec, 0xda, 0x37, 0x84, 0xa8, 0x82, 0x96, 0xc7, 0x82,
	0x67, 0x3b, 0x58, 0x9b, 0x52, 0x5b, 0xc1, 0x50, 0x5a, 0x36, 0x31, 0x94, 0xb6, 0x57, 0x68, 0x8f,
	0x7c, 0x00, 0x47, 0x1c, 0xef, 0x3b, 0x79, 0xa1, 0x0a, 0x3a, 0xae, 0x08, 0x5e, 0xee, 0xe1, 0x86,
	0x7c, 0x09, 0xf6, 0xc7, 0x12, 0xf1, 0x5f, 0xa8, 0x31, 0x4c, 0xae, 0x4a, 0x34, 0x17, 0x23, 0xba,
	0xb4, 0x6b, 0x87, 0xc0, 0xdf, 0x4f, 0x18, 0x91, 0x77, 0x30, 0x14, 0x07, 0x23, 0x7a, 0xb8, 0xf6,
	0x39, 0x83, 0xc5, 0x2e, 0xa8, 0x59, 0x77, 0x8e, 0x80, 0xdf, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x85, 0x16, 0xaa, 0x9f, 0x83, 0x05, 0x00, 0x00,
}
