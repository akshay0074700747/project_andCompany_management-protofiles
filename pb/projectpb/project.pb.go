// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: project.proto

package projectpb

import (
	companypb "github.com/akshay0074700747/projectandCompany_management_protofiles/pb/companypb"
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

type CreateProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Description     string `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	ProjectUsername string `protobuf:"bytes,3,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	Aim             string `protobuf:"bytes,4,opt,name=Aim,proto3" json:"Aim,omitempty"`
	IsCompanybased  bool   `protobuf:"varint,5,opt,name=IsCompanybased,proto3" json:"IsCompanybased,omitempty"`
	ComapanyId      string `protobuf:"bytes,6,opt,name=ComapanyId,proto3" json:"ComapanyId,omitempty"`
}

func (x *CreateProjectReq) Reset() {
	*x = CreateProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectReq) ProtoMessage() {}

func (x *CreateProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectReq.ProtoReflect.Descriptor instead.
func (*CreateProjectReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProjectReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProjectReq) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *CreateProjectReq) GetAim() string {
	if x != nil {
		return x.Aim
	}
	return ""
}

func (x *CreateProjectReq) GetIsCompanybased() bool {
	if x != nil {
		return x.IsCompanybased
	}
	return false
}

func (x *CreateProjectReq) GetComapanyId() string {
	if x != nil {
		return x.ComapanyId
	}
	return ""
}

type CreateProjectRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID       string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description     string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	ProjectUsername string `protobuf:"bytes,4,opt,name=ProjectUsername,proto3" json:"ProjectUsername,omitempty"`
	Aim             string `protobuf:"bytes,5,opt,name=Aim,proto3" json:"Aim,omitempty"`
	IsCompanybased  bool   `protobuf:"varint,6,opt,name=IsCompanybased,proto3" json:"IsCompanybased,omitempty"`
	ComapanyID      string `protobuf:"bytes,7,opt,name=ComapanyID,proto3" json:"ComapanyID,omitempty"`
}

func (x *CreateProjectRes) Reset() {
	*x = CreateProjectRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProjectRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProjectRes) ProtoMessage() {}

func (x *CreateProjectRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProjectRes.ProtoReflect.Descriptor instead.
func (*CreateProjectRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProjectRes) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *CreateProjectRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProjectRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateProjectRes) GetProjectUsername() string {
	if x != nil {
		return x.ProjectUsername
	}
	return ""
}

func (x *CreateProjectRes) GetAim() string {
	if x != nil {
		return x.Aim
	}
	return ""
}

func (x *CreateProjectRes) GetIsCompanybased() bool {
	if x != nil {
		return x.IsCompanybased
	}
	return false
}

func (x *CreateProjectRes) GetComapanyID() string {
	if x != nil {
		return x.ComapanyID
	}
	return ""
}

type AddMemberReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email        string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	PermissionID uint32 `protobuf:"varint,2,opt,name=PermissionID,proto3" json:"PermissionID,omitempty"`
	RoleID       uint32 `protobuf:"varint,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
}

func (x *AddMemberReq) Reset() {
	*x = AddMemberReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMemberReq) ProtoMessage() {}

func (x *AddMemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMemberReq.ProtoReflect.Descriptor instead.
func (*AddMemberReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{2}
}

func (x *AddMemberReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddMemberReq) GetPermissionID() uint32 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

func (x *AddMemberReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
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
		mi := &file_project_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReq) ProtoMessage() {}

func (x *SearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[3]
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
	return file_project_proto_rawDescGZIP(), []int{3}
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
		mi := &file_project_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRes) ProtoMessage() {}

func (x *SearchRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[4]
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
	return file_project_proto_rawDescGZIP(), []int{4}
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

type ProjectInvitesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID  string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	ProjectAim string `protobuf:"bytes,2,opt,name=ProjectAim,proto3" json:"ProjectAim,omitempty"`
	Members    uint32 `protobuf:"varint,3,opt,name=Members,proto3" json:"Members,omitempty"`
}

func (x *ProjectInvitesRes) Reset() {
	*x = ProjectInvitesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectInvitesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectInvitesRes) ProtoMessage() {}

func (x *ProjectInvitesRes) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectInvitesRes.ProtoReflect.Descriptor instead.
func (*ProjectInvitesRes) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectInvitesRes) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *ProjectInvitesRes) GetProjectAim() string {
	if x != nil {
		return x.ProjectAim
	}
	return ""
}

func (x *ProjectInvitesRes) GetMembers() uint32 {
	if x != nil {
		return x.Members
	}
	return 0
}

type AcceptProjectInviteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectID string `protobuf:"bytes,1,opt,name=ProjectID,proto3" json:"ProjectID,omitempty"`
	Accept    bool   `protobuf:"varint,2,opt,name=Accept,proto3" json:"Accept,omitempty"`
	UserID    string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *AcceptProjectInviteReq) Reset() {
	*x = AcceptProjectInviteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptProjectInviteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptProjectInviteReq) ProtoMessage() {}

func (x *AcceptProjectInviteReq) ProtoReflect() protoreflect.Message {
	mi := &file_project_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptProjectInviteReq.ProtoReflect.Descriptor instead.
func (*AcceptProjectInviteReq) Descriptor() ([]byte, []int) {
	return file_project_proto_rawDescGZIP(), []int{6}
}

func (x *AcceptProjectInviteReq) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *AcceptProjectInviteReq) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

func (x *AcceptProjectInviteReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

var File_project_proto protoreflect.FileDescriptor

var file_project_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x28, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x69, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x69, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x49,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61,
	0x73, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x69, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x69, 0x6d, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x62, 0x61, 0x73, 0x65, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6d, 0x61, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x22, 0x60, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f,
	0x6c, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x22, 0x23, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x35, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6b,
	0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x69, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x69,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x16, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x32, 0xab, 0x03, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x3b,
	0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a, 0x10, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x66, 0x6f, 0x72, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x30,
	0x01, 0x12, 0x4e, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_proto_rawDescOnce sync.Once
	file_project_proto_rawDescData = file_project_proto_rawDesc
)

func file_project_proto_rawDescGZIP() []byte {
	file_project_proto_rawDescOnce.Do(func() {
		file_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_proto_rawDescData)
	})
	return file_project_proto_rawDescData
}

var file_project_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_project_proto_goTypes = []interface{}{
	(*CreateProjectReq)(nil),       // 0: project.CreateProjectReq
	(*CreateProjectRes)(nil),       // 1: project.CreateProjectRes
	(*AddMemberReq)(nil),           // 2: project.AddMemberReq
	(*SearchReq)(nil),              // 3: project.SearchReq
	(*SearchRes)(nil),              // 4: project.SearchRes
	(*ProjectInvitesRes)(nil),      // 5: project.ProjectInvitesRes
	(*AcceptProjectInviteReq)(nil), // 6: project.AcceptProjectInviteReq
	(*empty.Empty)(nil),            // 7: google.protobuf.Empty
	(*companypb.Permission)(nil),   // 8: company.Permission
}
var file_project_proto_depIdxs = []int32{
	0, // 0: project.ProjectService.CreateProject:input_type -> project.CreateProjectReq
	7, // 1: project.ProjectService.GetPermissions:input_type -> google.protobuf.Empty
	2, // 2: project.ProjectService.AddMembers:input_type -> project.AddMemberReq
	3, // 3: project.ProjectService.SearchforMembers:input_type -> project.SearchReq
	7, // 4: project.ProjectService.ProjectInvites:input_type -> google.protobuf.Empty
	6, // 5: project.ProjectService.AcceptProjectInvite:input_type -> project.AcceptProjectInviteReq
	1, // 6: project.ProjectService.CreateProject:output_type -> project.CreateProjectRes
	8, // 7: project.ProjectService.GetPermissions:output_type -> company.Permission
	7, // 8: project.ProjectService.AddMembers:output_type -> google.protobuf.Empty
	4, // 9: project.ProjectService.SearchforMembers:output_type -> project.SearchRes
	5, // 10: project.ProjectService.ProjectInvites:output_type -> project.ProjectInvitesRes
	7, // 11: project.ProjectService.AcceptProjectInvite:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_proto_init() }
func file_project_proto_init() {
	if File_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectReq); i {
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
		file_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProjectRes); i {
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
		file_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMemberReq); i {
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
		file_project_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_project_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_project_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectInvitesRes); i {
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
		file_project_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptProjectInviteReq); i {
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
			RawDescriptor: file_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_project_proto_goTypes,
		DependencyIndexes: file_project_proto_depIdxs,
		MessageInfos:      file_project_proto_msgTypes,
	}.Build()
	File_project_proto = out.File
	file_project_proto_rawDesc = nil
	file_project_proto_goTypes = nil
	file_project_proto_depIdxs = nil
}
