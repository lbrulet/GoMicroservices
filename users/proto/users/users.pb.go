// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/users/users.proto

package go_micro_srv_users

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

type CreateRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{0}
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

func (m *CreateRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{4}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{5}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{6}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{7}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb78b70ef3ad6f73, []int{8}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.users.CreateRequest")
	proto.RegisterType((*User)(nil), "go.micro.srv.users.User")
	proto.RegisterType((*Message)(nil), "go.micro.srv.users.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.users.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.users.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.users.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.users.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.users.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.users.Pong")
}

func init() { proto.RegisterFile("proto/users/users.proto", fileDescriptor_bb78b70ef3ad6f73) }

var fileDescriptor_bb78b70ef3ad6f73 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x6e, 0x3e, 0x9a, 0xf6, 0x1d, 0x78, 0xa5, 0x0e, 0x45, 0x43, 0x5a, 0x45, 0x17, 0x85, 0x7a,
	0x89, 0x45, 0x7f, 0x42, 0x05, 0x4f, 0x82, 0x44, 0x3c, 0x78, 0xf0, 0xb0, 0x36, 0x4b, 0x08, 0x36,
	0xd9, 0xba, 0x93, 0x2a, 0xfe, 0x31, 0x7f, 0x9f, 0x64, 0x37, 0x91, 0xb6, 0x6e, 0xf1, 0x12, 0x76,
	0xf2, 0x7c, 0xcc, 0x3c, 0x9b, 0x09, 0x1c, 0x2e, 0x95, 0xac, 0xe4, 0xe5, 0x8a, 0x84, 0x22, 0xf3,
	0x8c, 0xf5, 0x1b, 0xc4, 0x4c, 0xc6, 0x45, 0x3e, 0x57, 0x32, 0x26, 0xf5, 0x1e, 0x6b, 0x84, 0x3d,
	0xc3, 0xff, 0x99, 0x12, 0xbc, 0x12, 0x89, 0x78, 0x5b, 0x09, 0xaa, 0x30, 0x82, 0x7e, 0x8d, 0x94,
	0xbc, 0x10, 0xa1, 0x73, 0xe2, 0x4c, 0xfe, 0x25, 0x3f, 0x35, 0x0e, 0xa1, 0x2b, 0x0a, 0x9e, 0x2f,
	0x42, 0x57, 0x03, 0xa6, 0xa8, 0x15, 0x4b, 0x4e, 0xf4, 0x21, 0x55, 0x1a, 0x7a, 0x46, 0xd1, 0xd6,
	0x2c, 0x05, 0xff, 0x91, 0x84, 0xc2, 0x3d, 0x70, 0xf3, 0x54, 0xfb, 0x79, 0x89, 0x9b, 0xa7, 0x1b,
	0x5d, 0xdc, 0x5d, 0x5d, 0xbc, 0x5d, 0x5d, 0xfc, 0xad, 0x2e, 0x23, 0xe8, 0xdd, 0x09, 0x22, 0x9e,
	0x09, 0x1c, 0x80, 0x47, 0xfc, 0xb3, 0x99, 0xbc, 0x3e, 0xb2, 0x23, 0xe8, 0xb5, 0xd9, 0x10, 0xfc,
	0xb5, 0x5c, 0xfa, 0xcc, 0xc6, 0xd0, 0x4f, 0x04, 0x2d, 0x65, 0x49, 0x5a, 0x5c, 0x50, 0xd6, 0x8a,
	0x0b, 0xca, 0xd8, 0x04, 0x06, 0x0f, 0x95, 0x12, 0xbc, 0xc8, 0xcb, 0xac, 0x75, 0x19, 0x42, 0x77,
	0x2e, 0x57, 0x65, 0xd5, 0xc4, 0x31, 0x05, 0xbb, 0x80, 0xfd, 0x35, 0x66, 0x63, 0x68, 0xa7, 0x1e,
	0x83, 0x7f, 0x9f, 0x97, 0x19, 0x1e, 0x40, 0x40, 0x95, 0x92, 0xaf, 0xa2, 0x81, 0x9b, 0x4a, 0xe3,
	0x72, 0x37, 0x7e, 0xf5, 0xe5, 0x42, 0xb7, 0xbe, 0x55, 0xc2, 0x19, 0xf8, 0x33, 0xbe, 0x58, 0xe0,
	0x28, 0xfe, 0xfd, 0x69, 0xe3, 0x66, 0xde, 0x68, 0x6c, 0x07, 0xcd, 0x88, 0xac, 0x83, 0x4f, 0x10,
	0x98, 0xc9, 0xf1, 0xcc, 0xc6, 0xdc, 0xce, 0x1f, 0x9d, 0xff, 0xc1, 0x6a, 0x8d, 0xa7, 0x0e, 0xde,
	0x40, 0xbf, 0x4e, 0xaa, 0xd3, 0x84, 0x36, 0x59, 0x8d, 0x46, 0x76, 0x44, 0x96, 0x19, 0xeb, 0x4c,
	0x9c, 0xa9, 0x83, 0xb7, 0x10, 0x98, 0x1d, 0xc5, 0x53, 0x1b, 0x73, 0x63, 0x7f, 0xed, 0x66, 0xf5,
	0x6d, 0xb1, 0xce, 0x4b, 0xa0, 0xff, 0x83, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x84,
	0xb1, 0x3a, 0x22, 0x03, 0x00, 0x00,
}
