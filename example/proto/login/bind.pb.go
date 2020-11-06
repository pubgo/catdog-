// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example/proto/login/bind.proto

// 绑定手机号

package login

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CheckRequest struct {
	// 区号
	NationCode string `protobuf:"bytes,1,opt,name=nationCode,proto3" json:"nationCode,omitempty"`
	// 手机号
	Telephone string `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	// uid
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// 前缀,通常为空,抖音必须为DY-
	Origin               string   `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetNationCode() string {
	if m != nil {
		return m.NationCode
	}
	return ""
}

func (m *CheckRequest) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *CheckRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CheckRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type CheckResponse struct {
	// code,不为0为错误
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 错误信息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{1}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CheckResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CheckResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *CheckResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindVerifyRequest struct {
	// 区号
	NationCode string `protobuf:"bytes,1,opt,name=nationCode,proto3" json:"nationCode,omitempty"`
	// 手机号
	Telephone string `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	// uid
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// 验证码
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	// 前缀,通常为空,抖音必须为DY-
	Origin               string   `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindVerifyRequest) Reset()         { *m = BindVerifyRequest{} }
func (m *BindVerifyRequest) String() string { return proto.CompactTextString(m) }
func (*BindVerifyRequest) ProtoMessage()    {}
func (*BindVerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{2}
}

func (m *BindVerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindVerifyRequest.Unmarshal(m, b)
}
func (m *BindVerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindVerifyRequest.Marshal(b, m, deterministic)
}
func (m *BindVerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindVerifyRequest.Merge(m, src)
}
func (m *BindVerifyRequest) XXX_Size() int {
	return xxx_messageInfo_BindVerifyRequest.Size(m)
}
func (m *BindVerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindVerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindVerifyRequest proto.InternalMessageInfo

func (m *BindVerifyRequest) GetNationCode() string {
	if m != nil {
		return m.NationCode
	}
	return ""
}

func (m *BindVerifyRequest) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *BindVerifyRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *BindVerifyRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *BindVerifyRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type BindVerifyResponse struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BindVerifyResponse) Reset()         { *m = BindVerifyResponse{} }
func (m *BindVerifyResponse) String() string { return proto.CompactTextString(m) }
func (*BindVerifyResponse) ProtoMessage()    {}
func (*BindVerifyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{3}
}

func (m *BindVerifyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindVerifyResponse.Unmarshal(m, b)
}
func (m *BindVerifyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindVerifyResponse.Marshal(b, m, deterministic)
}
func (m *BindVerifyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindVerifyResponse.Merge(m, src)
}
func (m *BindVerifyResponse) XXX_Size() int {
	return xxx_messageInfo_BindVerifyResponse.Size(m)
}
func (m *BindVerifyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindVerifyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindVerifyResponse proto.InternalMessageInfo

func (m *BindVerifyResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BindVerifyResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BindVerifyResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *BindVerifyResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindData struct {
	// uid
	BindId               int64    `protobuf:"varint,1,opt,name=bindId,proto3" json:"bindId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindData) Reset()         { *m = BindData{} }
func (m *BindData) String() string { return proto.CompactTextString(m) }
func (*BindData) ProtoMessage()    {}
func (*BindData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{4}
}

func (m *BindData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindData.Unmarshal(m, b)
}
func (m *BindData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindData.Marshal(b, m, deterministic)
}
func (m *BindData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindData.Merge(m, src)
}
func (m *BindData) XXX_Size() int {
	return xxx_messageInfo_BindData.Size(m)
}
func (m *BindData) XXX_DiscardUnknown() {
	xxx_messageInfo_BindData.DiscardUnknown(m)
}

var xxx_messageInfo_BindData proto.InternalMessageInfo

func (m *BindData) GetBindId() int64 {
	if m != nil {
		return m.BindId
	}
	return 0
}

type BindChangeRequest struct {
	// 区号
	NationCode string `protobuf:"bytes,1,opt,name=nationCode,proto3" json:"nationCode,omitempty"`
	// 手机号
	Telephone string `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	// uid
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// 验证码
	Code string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	// 前缀,通常为空,抖音必须为DY-
	Origin               string   `protobuf:"bytes,5,opt,name=origin,proto3" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindChangeRequest) Reset()         { *m = BindChangeRequest{} }
func (m *BindChangeRequest) String() string { return proto.CompactTextString(m) }
func (*BindChangeRequest) ProtoMessage()    {}
func (*BindChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{5}
}

func (m *BindChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindChangeRequest.Unmarshal(m, b)
}
func (m *BindChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindChangeRequest.Marshal(b, m, deterministic)
}
func (m *BindChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindChangeRequest.Merge(m, src)
}
func (m *BindChangeRequest) XXX_Size() int {
	return xxx_messageInfo_BindChangeRequest.Size(m)
}
func (m *BindChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindChangeRequest proto.InternalMessageInfo

func (m *BindChangeRequest) GetNationCode() string {
	if m != nil {
		return m.NationCode
	}
	return ""
}

func (m *BindChangeRequest) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *BindChangeRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *BindChangeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *BindChangeRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type BindChangeResponse struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// msg
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 *BindData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BindChangeResponse) Reset()         { *m = BindChangeResponse{} }
func (m *BindChangeResponse) String() string { return proto.CompactTextString(m) }
func (*BindChangeResponse) ProtoMessage()    {}
func (*BindChangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{6}
}

func (m *BindChangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindChangeResponse.Unmarshal(m, b)
}
func (m *BindChangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindChangeResponse.Marshal(b, m, deterministic)
}
func (m *BindChangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindChangeResponse.Merge(m, src)
}
func (m *BindChangeResponse) XXX_Size() int {
	return xxx_messageInfo_BindChangeResponse.Size(m)
}
func (m *BindChangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindChangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindChangeResponse proto.InternalMessageInfo

func (m *BindChangeResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BindChangeResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BindChangeResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *BindChangeResponse) GetData() *BindData {
	if m != nil {
		return m.Data
	}
	return nil
}

type AutomaticBindRequest struct {
	// 区号
	NationCode string `protobuf:"bytes,1,opt,name=nationCode,proto3" json:"nationCode,omitempty"`
	// 手机号
	Telephone string `protobuf:"bytes,2,opt,name=telephone,proto3" json:"telephone,omitempty"`
	// uid
	Uid int64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// 前缀,通常为空,抖音必须为DY-
	Origin               string   `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutomaticBindRequest) Reset()         { *m = AutomaticBindRequest{} }
func (m *AutomaticBindRequest) String() string { return proto.CompactTextString(m) }
func (*AutomaticBindRequest) ProtoMessage()    {}
func (*AutomaticBindRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{7}
}

func (m *AutomaticBindRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutomaticBindRequest.Unmarshal(m, b)
}
func (m *AutomaticBindRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutomaticBindRequest.Marshal(b, m, deterministic)
}
func (m *AutomaticBindRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutomaticBindRequest.Merge(m, src)
}
func (m *AutomaticBindRequest) XXX_Size() int {
	return xxx_messageInfo_AutomaticBindRequest.Size(m)
}
func (m *AutomaticBindRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AutomaticBindRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AutomaticBindRequest proto.InternalMessageInfo

func (m *AutomaticBindRequest) GetNationCode() string {
	if m != nil {
		return m.NationCode
	}
	return ""
}

func (m *AutomaticBindRequest) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *AutomaticBindRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AutomaticBindRequest) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

type AutomaticBindResponse struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 *BindData `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AutomaticBindResponse) Reset()         { *m = AutomaticBindResponse{} }
func (m *AutomaticBindResponse) String() string { return proto.CompactTextString(m) }
func (*AutomaticBindResponse) ProtoMessage()    {}
func (*AutomaticBindResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{8}
}

func (m *AutomaticBindResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutomaticBindResponse.Unmarshal(m, b)
}
func (m *AutomaticBindResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutomaticBindResponse.Marshal(b, m, deterministic)
}
func (m *AutomaticBindResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutomaticBindResponse.Merge(m, src)
}
func (m *AutomaticBindResponse) XXX_Size() int {
	return xxx_messageInfo_AutomaticBindResponse.Size(m)
}
func (m *AutomaticBindResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AutomaticBindResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AutomaticBindResponse proto.InternalMessageInfo

func (m *AutomaticBindResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AutomaticBindResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *AutomaticBindResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *AutomaticBindResponse) GetData() *BindData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindPhoneParseRequest struct {
	// 用于解析手机号加密数据
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// 用于解析手机号加密数据
	EncryptedData string `protobuf:"bytes,2,opt,name=encryptedData,proto3" json:"encryptedData,omitempty"`
	// 用于解析手机号加密数据
	Iv string `protobuf:"bytes,3,opt,name=iv,proto3" json:"iv,omitempty"`
	// platformId
	PlatformId int64 `protobuf:"varint,4,opt,name=platformId,proto3" json:"platformId,omitempty"`
	// uid，有uid的情况下不使用code
	Uid                  int64    `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindPhoneParseRequest) Reset()         { *m = BindPhoneParseRequest{} }
func (m *BindPhoneParseRequest) String() string { return proto.CompactTextString(m) }
func (*BindPhoneParseRequest) ProtoMessage()    {}
func (*BindPhoneParseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{9}
}

func (m *BindPhoneParseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindPhoneParseRequest.Unmarshal(m, b)
}
func (m *BindPhoneParseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindPhoneParseRequest.Marshal(b, m, deterministic)
}
func (m *BindPhoneParseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindPhoneParseRequest.Merge(m, src)
}
func (m *BindPhoneParseRequest) XXX_Size() int {
	return xxx_messageInfo_BindPhoneParseRequest.Size(m)
}
func (m *BindPhoneParseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindPhoneParseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindPhoneParseRequest proto.InternalMessageInfo

func (m *BindPhoneParseRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *BindPhoneParseRequest) GetEncryptedData() string {
	if m != nil {
		return m.EncryptedData
	}
	return ""
}

func (m *BindPhoneParseRequest) GetIv() string {
	if m != nil {
		return m.Iv
	}
	return ""
}

func (m *BindPhoneParseRequest) GetPlatformId() int64 {
	if m != nil {
		return m.PlatformId
	}
	return 0
}

func (m *BindPhoneParseRequest) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type BindPhoneParseResponse struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BindPhoneParseResponse) Reset()         { *m = BindPhoneParseResponse{} }
func (m *BindPhoneParseResponse) String() string { return proto.CompactTextString(m) }
func (*BindPhoneParseResponse) ProtoMessage()    {}
func (*BindPhoneParseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{10}
}

func (m *BindPhoneParseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindPhoneParseResponse.Unmarshal(m, b)
}
func (m *BindPhoneParseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindPhoneParseResponse.Marshal(b, m, deterministic)
}
func (m *BindPhoneParseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindPhoneParseResponse.Merge(m, src)
}
func (m *BindPhoneParseResponse) XXX_Size() int {
	return xxx_messageInfo_BindPhoneParseResponse.Size(m)
}
func (m *BindPhoneParseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindPhoneParseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindPhoneParseResponse proto.InternalMessageInfo

func (m *BindPhoneParseResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BindPhoneParseResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BindPhoneParseResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *BindPhoneParseResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type BindPhoneParseByOneClickRequest struct {
	// 用于解析手机号加密数据
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// platformId
	PlatformId int64 `protobuf:"varint,2,opt,name=platformId,proto3" json:"platformId,omitempty"`
	// telephone 有手机号即验证手机号
	Telephone            string   `protobuf:"bytes,3,opt,name=telephone,proto3" json:"telephone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindPhoneParseByOneClickRequest) Reset()         { *m = BindPhoneParseByOneClickRequest{} }
func (m *BindPhoneParseByOneClickRequest) String() string { return proto.CompactTextString(m) }
func (*BindPhoneParseByOneClickRequest) ProtoMessage()    {}
func (*BindPhoneParseByOneClickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{11}
}

func (m *BindPhoneParseByOneClickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindPhoneParseByOneClickRequest.Unmarshal(m, b)
}
func (m *BindPhoneParseByOneClickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindPhoneParseByOneClickRequest.Marshal(b, m, deterministic)
}
func (m *BindPhoneParseByOneClickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindPhoneParseByOneClickRequest.Merge(m, src)
}
func (m *BindPhoneParseByOneClickRequest) XXX_Size() int {
	return xxx_messageInfo_BindPhoneParseByOneClickRequest.Size(m)
}
func (m *BindPhoneParseByOneClickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BindPhoneParseByOneClickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BindPhoneParseByOneClickRequest proto.InternalMessageInfo

func (m *BindPhoneParseByOneClickRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *BindPhoneParseByOneClickRequest) GetPlatformId() int64 {
	if m != nil {
		return m.PlatformId
	}
	return 0
}

func (m *BindPhoneParseByOneClickRequest) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

type BindPhoneParseByOneClickResponse struct {
	// code
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 消息
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// 时间戳
	NowTime int64 `protobuf:"varint,3,opt,name=nowTime,proto3" json:"nowTime,omitempty"`
	// 数据
	Data                 map[string]string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BindPhoneParseByOneClickResponse) Reset()         { *m = BindPhoneParseByOneClickResponse{} }
func (m *BindPhoneParseByOneClickResponse) String() string { return proto.CompactTextString(m) }
func (*BindPhoneParseByOneClickResponse) ProtoMessage()    {}
func (*BindPhoneParseByOneClickResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8117a4177c777f9c, []int{12}
}

func (m *BindPhoneParseByOneClickResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindPhoneParseByOneClickResponse.Unmarshal(m, b)
}
func (m *BindPhoneParseByOneClickResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindPhoneParseByOneClickResponse.Marshal(b, m, deterministic)
}
func (m *BindPhoneParseByOneClickResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindPhoneParseByOneClickResponse.Merge(m, src)
}
func (m *BindPhoneParseByOneClickResponse) XXX_Size() int {
	return xxx_messageInfo_BindPhoneParseByOneClickResponse.Size(m)
}
func (m *BindPhoneParseByOneClickResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BindPhoneParseByOneClickResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BindPhoneParseByOneClickResponse proto.InternalMessageInfo

func (m *BindPhoneParseByOneClickResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BindPhoneParseByOneClickResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *BindPhoneParseByOneClickResponse) GetNowTime() int64 {
	if m != nil {
		return m.NowTime
	}
	return 0
}

func (m *BindPhoneParseByOneClickResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "login.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "login.CheckResponse")
	proto.RegisterMapType((map[string]string)(nil), "login.CheckResponse.DataEntry")
	proto.RegisterType((*BindVerifyRequest)(nil), "login.BindVerifyRequest")
	proto.RegisterType((*BindVerifyResponse)(nil), "login.BindVerifyResponse")
	proto.RegisterMapType((map[string]string)(nil), "login.BindVerifyResponse.DataEntry")
	proto.RegisterType((*BindData)(nil), "login.BindData")
	proto.RegisterType((*BindChangeRequest)(nil), "login.BindChangeRequest")
	proto.RegisterType((*BindChangeResponse)(nil), "login.BindChangeResponse")
	proto.RegisterType((*AutomaticBindRequest)(nil), "login.AutomaticBindRequest")
	proto.RegisterType((*AutomaticBindResponse)(nil), "login.AutomaticBindResponse")
	proto.RegisterType((*BindPhoneParseRequest)(nil), "login.BindPhoneParseRequest")
	proto.RegisterType((*BindPhoneParseResponse)(nil), "login.BindPhoneParseResponse")
	proto.RegisterMapType((map[string]string)(nil), "login.BindPhoneParseResponse.DataEntry")
	proto.RegisterType((*BindPhoneParseByOneClickRequest)(nil), "login.BindPhoneParseByOneClickRequest")
	proto.RegisterType((*BindPhoneParseByOneClickResponse)(nil), "login.BindPhoneParseByOneClickResponse")
	proto.RegisterMapType((map[string]string)(nil), "login.BindPhoneParseByOneClickResponse.DataEntry")
}

func init() {
	proto.RegisterFile("example/proto/login/bind.proto", fileDescriptor_8117a4177c777f9c)
}

var fileDescriptor_8117a4177c777f9c = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x97, 0xbd, 0x8f, 0x1b, 0x45,
	0x18, 0xc6, 0x35, 0xbb, 0xe7, 0x4b, 0xee, 0x3d, 0x1c, 0xc8, 0x70, 0x89, 0x1c, 0x73, 0x64, 0xad,
	0x89, 0x4c, 0xac, 0x8b, 0xbc, 0x0b, 0xa6, 0x38, 0x74, 0x82, 0x22, 0x36, 0x29, 0x52, 0x25, 0x5a,
	0x45, 0x48, 0x94, 0xeb, 0xf5, 0xdc, 0x7a, 0x14, 0x7b, 0x66, 0xd9, 0x5d, 0x1b, 0x4c, 0x83, 0x44,
	0x4d, 0x47, 0x4b, 0x87, 0x44, 0x91, 0x92, 0xff, 0x83, 0x02, 0xd1, 0x6f, 0x85, 0x28, 0xb6, 0xa2,
	0xa0, 0xa0, 0x44, 0x33, 0xbb, 0xb6, 0x77, 0xfc, 0x71, 0xb9, 0xe6, 0x4c, 0xae, 0x59, 0xcf, 0xf7,
	0xf3, 0xcc, 0xfb, 0x9b, 0x2f, 0xc3, 0x7d, 0xfa, 0x8d, 0x37, 0x0e, 0x47, 0xd4, 0x09, 0x23, 0x91,
	0x08, 0x67, 0x24, 0x02, 0xc6, 0x9d, 0x3e, 0xe3, 0x03, 0x5b, 0x15, 0xe0, 0x8a, 0x2a, 0xa9, 0xb7,
	0x03, 0x96, 0x0c, 0x27, 0x7d, 0xdb, 0x17, 0x63, 0x27, 0x10, 0x81, 0xc8, 0x9b, 0xf7, 0x27, 0xe7,
	0x2a, 0x97, 0xf7, 0x95, 0xa9, 0xbc, 0x57, 0xfd, 0x38, 0x10, 0x22, 0x18, 0x51, 0xc7, 0x0b, 0x99,
	0xe3, 0x71, 0x2e, 0x12, 0x2f, 0x61, 0x82, 0xc7, 0x79, 0x2d, 0x79, 0x85, 0xe0, 0xad, 0xde, 0x90,
	0xfa, 0x2f, 0x5d, 0xfa, 0xd5, 0x84, 0xc6, 0x09, 0xb6, 0x01, 0xb8, 0x6a, 0xd1, 0x13, 0x03, 0x5a,
	0x43, 0x0d, 0xd4, 0x3a, 0xe8, 0xde, 0xca, 0x52, 0xab, 0x54, 0xea, 0x96, 0xd2, 0xf8, 0x11, 0x1c,
	0x24, 0x74, 0x44, 0xc3, 0xa1, 0xe0, 0xb4, 0x66, 0xa8, 0xe6, 0xd5, 0x2c, 0xb5, 0x96, 0x85, 0xee,
	0x32, 0x89, 0xef, 0x81, 0x39, 0x61, 0x83, 0x9a, 0xd9, 0x40, 0x2d, 0xb3, 0x7b, 0x23, 0x4b, 0x2d,
	0x99, 0x75, 0xe5, 0x07, 0x13, 0xd8, 0x17, 0x11, 0x0b, 0x18, 0xaf, 0xed, 0xa9, 0x41, 0x20, 0x4b,
	0xad, 0xa2, 0xc4, 0x2d, 0x7e, 0xc9, 0x5f, 0x08, 0xaa, 0x85, 0xd9, 0x38, 0x14, 0x3c, 0xa6, 0xf8,
	0x18, 0xf6, 0xfc, 0xb9, 0x4f, 0xb3, 0x7b, 0x33, 0x4b, 0x2d, 0x95, 0x77, 0xd5, 0x57, 0xca, 0x8d,
	0xe3, 0xa0, 0x70, 0xa5, 0xe4, 0xc6, 0x71, 0xe0, 0xca, 0x0f, 0x6e, 0xc2, 0x0d, 0x2e, 0xbe, 0x7e,
	0xc1, 0xc6, 0xb4, 0x70, 0x73, 0x98, 0xa5, 0xd6, 0xbc, 0xc8, 0x9d, 0x27, 0xf0, 0xa7, 0xb0, 0x37,
	0xf0, 0x12, 0xaf, 0xb6, 0xd7, 0x30, 0x5b, 0x87, 0x9d, 0xfb, 0xb6, 0x22, 0x60, 0x6b, 0x1e, 0xec,
	0xcf, 0xbd, 0xc4, 0x7b, 0xc2, 0x93, 0x68, 0x96, 0xeb, 0xcb, 0xf6, 0xae, 0xfa, 0xd6, 0x4f, 0xe1,
	0x60, 0x51, 0x89, 0xdf, 0x01, 0xf3, 0x25, 0x9d, 0xe5, 0x11, 0x75, 0x65, 0x12, 0x1f, 0x41, 0x65,
	0xea, 0x8d, 0x26, 0x45, 0xd8, 0xdc, 0x3c, 0x73, 0x66, 0x7c, 0x82, 0xc8, 0x6f, 0x08, 0x6e, 0x77,
	0x19, 0x1f, 0x7c, 0x41, 0x23, 0x76, 0x3e, 0xfb, 0xbf, 0xd1, 0xcc, 0x83, 0x9c, 0x83, 0x59, 0x0d,
	0xf2, 0x12, 0x5c, 0x65, 0x2b, 0xb8, 0xbf, 0x11, 0xe0, 0xf2, 0x7c, 0x76, 0x44, 0xef, 0xb1, 0x46,
	0xef, 0x41, 0x41, 0x6f, 0xdd, 0xc8, 0x55, 0x20, 0xb4, 0xe1, 0xa6, 0x14, 0x92, 0x9d, 0x65, 0x88,
	0xe4, 0x36, 0x7e, 0x3a, 0x28, 0x66, 0xaa, 0x42, 0x94, 0x97, 0xb8, 0xc5, 0xef, 0x02, 0x79, 0x6f,
	0xe8, 0xf1, 0x80, 0x5e, 0x7f, 0xe4, 0x3f, 0x17, 0xc8, 0xe7, 0xf3, 0xd9, 0x11, 0xf2, 0xf6, 0x02,
	0x39, 0x6a, 0x1d, 0x76, 0xde, 0x2e, 0x21, 0x97, 0x24, 0x56, 0xf1, 0x92, 0x5f, 0x11, 0x1c, 0x3d,
	0x9e, 0x24, 0x62, 0xec, 0x25, 0xcc, 0x97, 0xad, 0xae, 0xc3, 0x31, 0xf8, 0x0b, 0x82, 0x3b, 0x2b,
	0xa6, 0xdf, 0xcc, 0xe8, 0xfe, 0x8e, 0xe0, 0x8e, 0xac, 0x7c, 0x2e, 0x67, 0xfd, 0xdc, 0x8b, 0xe2,
	0xc5, 0xba, 0x2e, 0x1b, 0x5d, 0x5f, 0x5f, 0xa7, 0x50, 0xa5, 0xdc, 0x8f, 0x66, 0x61, 0x42, 0xd5,
	0xc0, 0x85, 0xe5, 0xdb, 0x59, 0x6a, 0xe9, 0x15, 0xae, 0x9e, 0xc5, 0x77, 0xc1, 0x60, 0x53, 0x35,
	0x83, 0x83, 0xee, 0x7e, 0x96, 0x5a, 0x06, 0x9b, 0xba, 0x06, 0x9b, 0x4a, 0x9a, 0xe1, 0xc8, 0x4b,
	0xce, 0x45, 0x34, 0x7e, 0x3a, 0x50, 0xee, 0xcd, 0x9c, 0xe6, 0xb2, 0xd4, 0x2d, 0xa5, 0xe7, 0x80,
	0x2a, 0xeb, 0x80, 0xc8, 0xbf, 0x08, 0xee, 0xae, 0xce, 0x69, 0x47, 0xd1, 0x7f, 0xa2, 0x1d, 0x67,
	0x0f, 0x4b, 0xd1, 0x5f, 0x37, 0x73, 0x15, 0x47, 0xda, 0x4f, 0x08, 0x2c, 0x5d, 0xad, 0x3b, 0x7b,
	0xc6, 0x69, 0x6f, 0xc4, 0x96, 0xcf, 0x87, 0x8b, 0xc1, 0xea, 0x1c, 0x8c, 0xd7, 0x72, 0xd0, 0x76,
	0x95, 0x79, 0xf1, 0xae, 0x22, 0x3f, 0x18, 0xd0, 0xd8, 0x6e, 0x6f, 0x47, 0x8c, 0x9e, 0x69, 0x8c,
	0x3e, 0xda, 0xc8, 0x68, 0xdd, 0xd6, 0x15, 0xd0, 0xea, 0xfc, 0x53, 0x81, 0xaa, 0xd4, 0x7d, 0xb1,
	0x38, 0x76, 0xbe, 0x84, 0x8a, 0x7a, 0xb9, 0xe0, 0x77, 0xf5, 0x77, 0x8c, 0x22, 0x57, 0x3f, 0xda,
	0xf4, 0xb8, 0x21, 0xcd, 0xef, 0xff, 0xf8, 0xf3, 0x47, 0xc3, 0x22, 0x75, 0x67, 0x12, 0xd3, 0x48,
	0xbd, 0x46, 0xdb, 0x8b, 0x90, 0x3b, 0xbe, 0x6c, 0x7b, 0x86, 0x4e, 0x30, 0x07, 0x58, 0x5e, 0xab,
	0xb8, 0xb6, 0xe1, 0xa6, 0xcd, 0x45, 0xee, 0x6d, 0xbd, 0x83, 0xc9, 0x23, 0xa5, 0xd4, 0x24, 0x8d,
	0x8d, 0x4a, 0x2a, 0x3b, 0x55, 0x3d, 0x4a, 0x7a, 0xf9, 0xe5, 0xa2, 0xe9, 0x69, 0xf7, 0xa7, 0xa6,
	0xa7, 0xdf, 0x44, 0x97, 0xd1, 0xf3, 0x55, 0x0f, 0xa9, 0xf7, 0x2d, 0x54, 0xb5, 0x13, 0x17, 0xbf,
	0x57, 0x0c, 0xbc, 0xe9, 0xf2, 0xa8, 0x1f, 0x6f, 0xae, 0x2c, 0x84, 0x6d, 0x25, 0xdc, 0x22, 0x0f,
	0x36, 0x0a, 0x7b, 0xf3, 0x3e, 0x6d, 0x59, 0x21, 0xb5, 0xbf, 0x83, 0x5b, 0xfa, 0xfa, 0xc1, 0xc7,
	0x5b, 0xb6, 0x7e, 0xae, 0xfe, 0xfe, 0x85, 0x07, 0x03, 0xf9, 0x50, 0xc9, 0x9f, 0x90, 0xe6, 0xf6,
	0x79, 0xab, 0x64, 0x3b, 0x94, 0xdd, 0xa4, 0x81, 0x57, 0x08, 0x6a, 0xdb, 0x56, 0x30, 0xfe, 0xe0,
	0xb5, 0x4b, 0x3c, 0x77, 0xf5, 0xf0, 0x92, 0x5b, 0x81, 0x7c, 0xa6, 0xfc, 0x9d, 0x92, 0xce, 0xa5,
	0xfc, 0xb5, 0xfb, 0xb3, 0xb6, 0xcc, 0xf8, 0x72, 0x8c, 0x33, 0x74, 0xd2, 0xdf, 0x57, 0xff, 0x6b,
	0x3e, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x72, 0xf0, 0x54, 0x4d, 0x0d, 0x00, 0x00,
}
