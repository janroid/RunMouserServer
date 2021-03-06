// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamepb.proto

package gamepb

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

type ReqLogin struct {
	Name                 *string  `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLogin) Reset()         { *m = ReqLogin{} }
func (m *ReqLogin) String() string { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()    {}
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_86f4eab4e7afbe0c, []int{0}
}

func (m *ReqLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLogin.Unmarshal(m, b)
}
func (m *ReqLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLogin.Marshal(b, m, deterministic)
}
func (m *ReqLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLogin.Merge(m, src)
}
func (m *ReqLogin) XXX_Size() int {
	return xxx_messageInfo_ReqLogin.Size(m)
}
func (m *ReqLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLogin proto.InternalMessageInfo

func (m *ReqLogin) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ReqLogin) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type ReqRegister struct {
	Name                 *string  `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRegister) Reset()         { *m = ReqRegister{} }
func (m *ReqRegister) String() string { return proto.CompactTextString(m) }
func (*ReqRegister) ProtoMessage()    {}
func (*ReqRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_86f4eab4e7afbe0c, []int{1}
}

func (m *ReqRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRegister.Unmarshal(m, b)
}
func (m *ReqRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRegister.Marshal(b, m, deterministic)
}
func (m *ReqRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRegister.Merge(m, src)
}
func (m *ReqRegister) XXX_Size() int {
	return xxx_messageInfo_ReqRegister.Size(m)
}
func (m *ReqRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRegister.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRegister proto.InternalMessageInfo

func (m *ReqRegister) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ReqRegister) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type RpsAuthor struct {
	Result               *int32   `protobuf:"varint,1,req,name=Result" json:"Result,omitempty"`
	Type                 *int32   `protobuf:"varint,2,req,name=Type" json:"Type,omitempty"`
	Uid                  *int64   `protobuf:"varint,3,opt,name=Uid" json:"Uid,omitempty"`
	Money                *int64   `protobuf:"varint,4,opt,name=Money" json:"Money,omitempty"`
	Icon                 *int32   `protobuf:"varint,5,opt,name=Icon" json:"Icon,omitempty"`
	Exp                  *int64   `protobuf:"varint,6,opt,name=Exp" json:"Exp,omitempty"`
	Name                 *string  `protobuf:"bytes,7,opt,name=Name" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpsAuthor) Reset()         { *m = RpsAuthor{} }
func (m *RpsAuthor) String() string { return proto.CompactTextString(m) }
func (*RpsAuthor) ProtoMessage()    {}
func (*RpsAuthor) Descriptor() ([]byte, []int) {
	return fileDescriptor_86f4eab4e7afbe0c, []int{2}
}

func (m *RpsAuthor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpsAuthor.Unmarshal(m, b)
}
func (m *RpsAuthor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpsAuthor.Marshal(b, m, deterministic)
}
func (m *RpsAuthor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpsAuthor.Merge(m, src)
}
func (m *RpsAuthor) XXX_Size() int {
	return xxx_messageInfo_RpsAuthor.Size(m)
}
func (m *RpsAuthor) XXX_DiscardUnknown() {
	xxx_messageInfo_RpsAuthor.DiscardUnknown(m)
}

var xxx_messageInfo_RpsAuthor proto.InternalMessageInfo

func (m *RpsAuthor) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *RpsAuthor) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *RpsAuthor) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *RpsAuthor) GetMoney() int64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *RpsAuthor) GetIcon() int32 {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return 0
}

func (m *RpsAuthor) GetExp() int64 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *RpsAuthor) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReqUserInfo struct {
	Uid                  *int64   `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Fields               *int32   `protobuf:"varint,2,req,name=Fields" json:"Fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserInfo) Reset()         { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string { return proto.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()    {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_86f4eab4e7afbe0c, []int{3}
}

func (m *ReqUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserInfo.Unmarshal(m, b)
}
func (m *ReqUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserInfo.Marshal(b, m, deterministic)
}
func (m *ReqUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserInfo.Merge(m, src)
}
func (m *ReqUserInfo) XXX_Size() int {
	return xxx_messageInfo_ReqUserInfo.Size(m)
}
func (m *ReqUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserInfo proto.InternalMessageInfo

func (m *ReqUserInfo) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *ReqUserInfo) GetFields() int32 {
	if m != nil && m.Fields != nil {
		return *m.Fields
	}
	return 0
}

type RpsUserInfo struct {
	Uid                  *int64   `protobuf:"varint,1,req,name=Uid" json:"Uid,omitempty"`
	Money                *int64   `protobuf:"varint,2,req,name=Money" json:"Money,omitempty"`
	Exp                  *int64   `protobuf:"varint,3,opt,name=Exp" json:"Exp,omitempty"`
	Icon                 *int32   `protobuf:"varint,4,opt,name=Icon" json:"Icon,omitempty"`
	AcName               *string  `protobuf:"bytes,5,opt,name=AcName" json:"AcName,omitempty"`
	AcPwd                *string  `protobuf:"bytes,6,opt,name=AcPwd" json:"AcPwd,omitempty"`
	PlayCount            *int64   `protobuf:"varint,7,opt,name=PlayCount" json:"PlayCount,omitempty"`
	PlayWin              *int64   `protobuf:"varint,8,opt,name=PlayWin" json:"PlayWin,omitempty"`
	PlayOut              *int64   `protobuf:"varint,9,opt,name=PlayOut" json:"PlayOut,omitempty"`
	PlayCreate           *int64   `protobuf:"varint,10,opt,name=PlayCreate" json:"PlayCreate,omitempty"`
	Honor                *int64   `protobuf:"varint,11,opt,name=Honor" json:"Honor,omitempty"`
	Gold                 *int64   `protobuf:"varint,12,opt,name=Gold" json:"Gold,omitempty"`
	Title                *int32   `protobuf:"varint,13,opt,name=Title" json:"Title,omitempty"`
	Status               *int32   `protobuf:"varint,14,opt,name=Status" json:"Status,omitempty"`
	Name                 *string  `protobuf:"bytes,15,opt,name=Name" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpsUserInfo) Reset()         { *m = RpsUserInfo{} }
func (m *RpsUserInfo) String() string { return proto.CompactTextString(m) }
func (*RpsUserInfo) ProtoMessage()    {}
func (*RpsUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_86f4eab4e7afbe0c, []int{4}
}

func (m *RpsUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpsUserInfo.Unmarshal(m, b)
}
func (m *RpsUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpsUserInfo.Marshal(b, m, deterministic)
}
func (m *RpsUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpsUserInfo.Merge(m, src)
}
func (m *RpsUserInfo) XXX_Size() int {
	return xxx_messageInfo_RpsUserInfo.Size(m)
}
func (m *RpsUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RpsUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RpsUserInfo proto.InternalMessageInfo

func (m *RpsUserInfo) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *RpsUserInfo) GetMoney() int64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *RpsUserInfo) GetExp() int64 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *RpsUserInfo) GetIcon() int32 {
	if m != nil && m.Icon != nil {
		return *m.Icon
	}
	return 0
}

func (m *RpsUserInfo) GetAcName() string {
	if m != nil && m.AcName != nil {
		return *m.AcName
	}
	return ""
}

func (m *RpsUserInfo) GetAcPwd() string {
	if m != nil && m.AcPwd != nil {
		return *m.AcPwd
	}
	return ""
}

func (m *RpsUserInfo) GetPlayCount() int64 {
	if m != nil && m.PlayCount != nil {
		return *m.PlayCount
	}
	return 0
}

func (m *RpsUserInfo) GetPlayWin() int64 {
	if m != nil && m.PlayWin != nil {
		return *m.PlayWin
	}
	return 0
}

func (m *RpsUserInfo) GetPlayOut() int64 {
	if m != nil && m.PlayOut != nil {
		return *m.PlayOut
	}
	return 0
}

func (m *RpsUserInfo) GetPlayCreate() int64 {
	if m != nil && m.PlayCreate != nil {
		return *m.PlayCreate
	}
	return 0
}

func (m *RpsUserInfo) GetHonor() int64 {
	if m != nil && m.Honor != nil {
		return *m.Honor
	}
	return 0
}

func (m *RpsUserInfo) GetGold() int64 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *RpsUserInfo) GetTitle() int32 {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return 0
}

func (m *RpsUserInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *RpsUserInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqLogin)(nil), "gamepb.ReqLogin")
	proto.RegisterType((*ReqRegister)(nil), "gamepb.ReqRegister")
	proto.RegisterType((*RpsAuthor)(nil), "gamepb.RpsAuthor")
	proto.RegisterType((*ReqUserInfo)(nil), "gamepb.ReqUserInfo")
	proto.RegisterType((*RpsUserInfo)(nil), "gamepb.RpsUserInfo")
}

func init() { proto.RegisterFile("gamepb.proto", fileDescriptor_86f4eab4e7afbe0c) }

var fileDescriptor_86f4eab4e7afbe0c = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0xd2, 0xa4, 0xcd, 0xb4, 0xfe, 0x61, 0x91, 0xb2, 0x88, 0x48, 0xc8, 0xa9, 0x27,
	0xaf, 0x82, 0xe0, 0xa1, 0x48, 0xd5, 0x82, 0xda, 0xb2, 0xb6, 0x78, 0x8e, 0xcd, 0x5a, 0x03, 0x69,
	0x36, 0xcd, 0x6e, 0x68, 0xfb, 0x51, 0xfc, 0x84, 0x7e, 0x0d, 0xd9, 0xc9, 0x26, 0xad, 0x17, 0xc1,
	0xdb, 0xbc, 0x37, 0x79, 0x9b, 0xdf, 0x63, 0xa0, 0xb7, 0x8c, 0x56, 0x3c, 0x7f, 0xbf, 0xca, 0x0b,
	0xa1, 0x04, 0xf1, 0x2a, 0x15, 0xde, 0x40, 0x87, 0xf1, 0xf5, 0x93, 0x58, 0x26, 0x19, 0x21, 0xd0,
	0x7a, 0x89, 0x56, 0x9c, 0x5a, 0x81, 0x3d, 0xf0, 0x19, 0xce, 0xe4, 0x1c, 0x3a, 0xd3, 0x48, 0xca,
	0x8d, 0x28, 0x62, 0x6a, 0xa3, 0xdf, 0xe8, 0xf0, 0x16, 0xba, 0x8c, 0xaf, 0x19, 0x5f, 0x26, 0x52,
	0xf1, 0xe2, 0xdf, 0xf1, 0x2f, 0x0b, 0x7c, 0x96, 0xcb, 0x61, 0xa9, 0x3e, 0x45, 0x41, 0xfa, 0xe0,
	0x31, 0x2e, 0xcb, 0x54, 0x61, 0xde, 0x65, 0x46, 0xe9, 0x57, 0x67, 0xbb, 0x9c, 0x63, 0xda, 0x65,
	0x38, 0x93, 0x53, 0x70, 0xe6, 0x49, 0x4c, 0x9d, 0xc0, 0x1a, 0x38, 0x4c, 0x8f, 0xe4, 0x0c, 0xdc,
	0x67, 0x91, 0xf1, 0x1d, 0x6d, 0xa1, 0x57, 0x09, 0x9d, 0x1d, 0x2f, 0x44, 0x46, 0xdd, 0xc0, 0xd2,
	0x59, 0x3d, 0xeb, 0xec, 0x68, 0x9b, 0x53, 0xaf, 0xca, 0x8e, 0xb6, 0x79, 0xc3, 0xdd, 0x0e, 0xac,
	0x9a, 0x3b, 0xbc, 0xc6, 0x6a, 0x73, 0xc9, 0x8b, 0x71, 0xf6, 0x21, 0xea, 0x1f, 0x6a, 0x32, 0xf3,
	0xc3, 0x3e, 0x78, 0xf7, 0x09, 0x4f, 0x63, 0x69, 0xc0, 0x8c, 0x0a, 0xbf, 0x6d, 0xe8, 0xb2, 0x5c,
	0xfe, 0x91, 0x6c, 0x50, 0x6d, 0xf4, 0x0c, 0xaa, 0xc1, 0x72, 0x7e, 0x61, 0x21, 0x7c, 0xeb, 0x00,
	0xbe, 0x0f, 0xde, 0x70, 0x81, 0xb0, 0x2e, 0xc2, 0x1a, 0xa5, 0xdf, 0x1c, 0x2e, 0xa6, 0x9b, 0x18,
	0x6b, 0xf9, 0xac, 0x12, 0xe4, 0x02, 0xfc, 0x69, 0x1a, 0xed, 0xee, 0x44, 0x99, 0x29, 0x6c, 0xe7,
	0xb0, 0xbd, 0x41, 0x28, 0xb4, 0xb5, 0x78, 0x4b, 0x32, 0xda, 0xc1, 0x5d, 0x2d, 0xeb, 0xcd, 0xa4,
	0x54, 0xd4, 0xdf, 0x6f, 0x26, 0xa5, 0x22, 0x97, 0x00, 0xf8, 0x40, 0xc1, 0x23, 0xc5, 0x29, 0xe0,
	0xf2, 0xc0, 0xd1, 0x1c, 0x8f, 0x22, 0x13, 0x05, 0xed, 0x56, 0x67, 0x40, 0xa1, 0x9b, 0x3c, 0x88,
	0x34, 0xa6, 0x3d, 0x34, 0x71, 0xd6, 0x5f, 0xce, 0x12, 0x95, 0x72, 0x7a, 0x84, 0xf5, 0x2a, 0xa1,
	0xfb, 0xbd, 0xaa, 0x48, 0x95, 0x92, 0x1e, 0xa3, 0x6d, 0x54, 0x73, 0xa2, 0x93, 0xfd, 0x89, 0x7e,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xd6, 0x52, 0xf1, 0xd0, 0x02, 0x00, 0x00,
}
