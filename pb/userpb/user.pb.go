// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: user.proto

package userpb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignupUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile   string `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SignupUserRequest) Reset() {
	*x = SignupUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupUserRequest) ProtoMessage() {}

func (x *SignupUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupUserRequest.ProtoReflect.Descriptor instead.
func (*SignupUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignupUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignupUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupUserRequest) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SignupUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Role string `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type AddRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *AddRoleReq) Reset() {
	*x = AddRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleReq) ProtoMessage() {}

func (x *AddRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleReq.ProtoReflect.Descriptor instead.
func (*AddRoleReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *AddRoleReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type StatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID      uint32 `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	IsAvailable bool   `protobuf:"varint,2,opt,name=IsAvailable,proto3" json:"IsAvailable,omitempty"`
	UserID      string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *StatusReq) Reset() {
	*x = StatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReq) ProtoMessage() {}

func (x *StatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReq.ProtoReflect.Descriptor instead.
func (*StatusReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *StatusReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *StatusReq) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

func (x *StatusReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type UserResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Email  string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile string `protobuf:"bytes,4,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
}

func (x *UserResponce) Reset() {
	*x = UserResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponce) ProtoMessage() {}

func (x *UserResponce) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponce.ProtoReflect.Descriptor instead.
func (*UserResponce) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponce) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponce) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserResponce) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponce) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type GetByEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *GetByEmailReq) Reset() {
	*x = GetByEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByEmailReq) ProtoMessage() {}

func (x *GetByEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByEmailReq.ProtoReflect.Descriptor instead.
func (*GetByEmailReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetByEmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetByEmailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetByEmailRes) Reset() {
	*x = GetByEmailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByEmailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByEmailRes) ProtoMessage() {}

func (x *GetByEmailRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByEmailRes.ProtoReflect.Descriptor instead.
func (*GetByEmailRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetByEmailRes) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type SearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID uint32 `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *SearchReq) Reset() {
	*x = SearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReq.ProtoReflect.Descriptor instead.
func (*SearchReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *SearchReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type SearchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *SearchRes) Reset() {
	*x = SearchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRes) ProtoMessage() {}

func (x *SearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRes.ProtoReflect.Descriptor instead.
func (*SearchRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{8}
}

func (x *SearchRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserDetailsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	RoleID uint32 `protobuf:"varint,2,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *GetUserDetailsReq) Reset() {
	*x = GetUserDetailsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsReq) ProtoMessage() {}

func (x *GetUserDetailsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsReq.ProtoReflect.Descriptor instead.
func (*GetUserDetailsReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserDetailsReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUserDetailsReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type GetUserDetailsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Mobile string `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Role   string `protobuf:"bytes,5,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *GetUserDetailsRes) Reset() {
	*x = GetUserDetailsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDetailsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailsRes) ProtoMessage() {}

func (x *GetUserDetailsRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailsRes.ProtoReflect.Descriptor instead.
func (*GetUserDetailsRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserDetailsRes) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetUserDetailsRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserDetailsRes) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *GetUserDetailsRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserDetailsRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type GetStreamofRolesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID uint32 `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *GetStreamofRolesReq) Reset() {
	*x = GetStreamofRolesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamofRolesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamofRolesReq) ProtoMessage() {}

func (x *GetStreamofRolesReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamofRolesReq.ProtoReflect.Descriptor instead.
func (*GetStreamofRolesReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{11}
}

func (x *GetStreamofRolesReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

type GetStreamofRolesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleIDsWithNames map[uint32]string `protobuf:"bytes,1,rep,name=RoleIDsWithNames,proto3" json:"RoleIDsWithNames,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetStreamofRolesRes) Reset() {
	*x = GetStreamofRolesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStreamofRolesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStreamofRolesRes) ProtoMessage() {}

func (x *GetStreamofRolesRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStreamofRolesRes.ProtoReflect.Descriptor instead.
func (*GetStreamofRolesRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{12}
}

func (x *GetStreamofRolesRes) GetRoleIDsWithNames() map[uint32]string {
	if x != nil {
		return x.RoleIDsWithNames
	}
	return nil
}

type GetRoleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetRoleReq) Reset() {
	*x = GetRoleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleReq) ProtoMessage() {}

func (x *GetRoleReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleReq.ProtoReflect.Descriptor instead.
func (*GetRoleReq) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{13}
}

func (x *GetRoleReq) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetRoleRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role string `protobuf:"bytes,1,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *GetRoleRes) Reset() {
	*x = GetRoleRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleRes) ProtoMessage() {}

func (x *GetRoleRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleRes.ProtoReflect.Descriptor instead.
func (*GetRoleRes) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{14}
}

func (x *GetRoleRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x71, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x20,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65,
	0x22, 0x5d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x73, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x60, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x23, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x43, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x6f, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0xb7, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x6f, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x5b, 0x0a,
	0x10, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6f, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44,
	0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x1a, 0x43, 0x0a, 0x15, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x44, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x22, 0x20, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x32,
	0xe5, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x10, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x66, 0x6f, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x30, 0x01, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x10,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6f, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x6f, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x6f, 0x66, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x28, 0x01, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x62, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_user_proto_goTypes = []interface{}{
	(*SignupUserRequest)(nil),   // 0: user.SignupUserRequest
	(*Role)(nil),                // 1: user.Role
	(*AddRoleReq)(nil),          // 2: user.AddRoleReq
	(*StatusReq)(nil),           // 3: user.StatusReq
	(*UserResponce)(nil),        // 4: user.UserResponce
	(*GetByEmailReq)(nil),       // 5: user.GetByEmailReq
	(*GetByEmailRes)(nil),       // 6: user.GetByEmailRes
	(*SearchReq)(nil),           // 7: user.SearchReq
	(*SearchRes)(nil),           // 8: user.SearchRes
	(*GetUserDetailsReq)(nil),   // 9: user.GetUserDetailsReq
	(*GetUserDetailsRes)(nil),   // 10: user.GetUserDetailsRes
	(*GetStreamofRolesReq)(nil), // 11: user.GetStreamofRolesReq
	(*GetStreamofRolesRes)(nil), // 12: user.GetStreamofRolesRes
	(*GetRoleReq)(nil),          // 13: user.GetRoleReq
	(*GetRoleRes)(nil),          // 14: user.GetRoleRes
	nil,                         // 15: user.GetStreamofRolesRes.RoleIDsWithNamesEntry
	(*empty.Empty)(nil),         // 16: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	15, // 0: user.GetStreamofRolesRes.RoleIDsWithNames:type_name -> user.GetStreamofRolesRes.RoleIDsWithNamesEntry
	0,  // 1: user.UserService.SignupUser:input_type -> user.SignupUserRequest
	16, // 2: user.UserService.GetRoles:input_type -> google.protobuf.Empty
	3,  // 3: user.UserService.SetStatus:input_type -> user.StatusReq
	5,  // 4: user.UserService.GetByEmail:input_type -> user.GetByEmailReq
	7,  // 5: user.UserService.SearchforMembers:input_type -> user.SearchReq
	2,  // 6: user.UserService.AddRoles:input_type -> user.AddRoleReq
	9,  // 7: user.UserService.GetUserDetails:input_type -> user.GetUserDetailsReq
	9,  // 8: user.UserService.GetStreamofUserDetails:input_type -> user.GetUserDetailsReq
	11, // 9: user.UserService.GetStreamofRoles:input_type -> user.GetStreamofRolesReq
	13, // 10: user.UserService.GetRole:input_type -> user.GetRoleReq
	4,  // 11: user.UserService.SignupUser:output_type -> user.UserResponce
	1,  // 12: user.UserService.GetRoles:output_type -> user.Role
	16, // 13: user.UserService.SetStatus:output_type -> google.protobuf.Empty
	6,  // 14: user.UserService.GetByEmail:output_type -> user.GetByEmailRes
	8,  // 15: user.UserService.SearchforMembers:output_type -> user.SearchRes
	16, // 16: user.UserService.AddRoles:output_type -> google.protobuf.Empty
	10, // 17: user.UserService.GetUserDetails:output_type -> user.GetUserDetailsRes
	10, // 18: user.UserService.GetStreamofUserDetails:output_type -> user.GetUserDetailsRes
	12, // 19: user.UserService.GetStreamofRoles:output_type -> user.GetStreamofRolesRes
	14, // 20: user.UserService.GetRole:output_type -> user.GetRoleRes
	11, // [11:21] is the sub-list for method output_type
	1,  // [1:11] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponce); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByEmailReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByEmailRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDetailsRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamofRolesReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStreamofRolesRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
