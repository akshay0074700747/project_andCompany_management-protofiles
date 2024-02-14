// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: company.proto

package companypb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street   string `protobuf:"bytes,1,opt,name=Street,proto3" json:"Street,omitempty"`
	StreetNo int32  `protobuf:"varint,2,opt,name=StreetNo,proto3" json:"StreetNo,omitempty"`
	PinNo    int32  `protobuf:"varint,3,opt,name=PinNo,proto3" json:"PinNo,omitempty"`
	District string `protobuf:"bytes,4,opt,name=District,proto3" json:"District,omitempty"`
	State    string `protobuf:"bytes,5,opt,name=State,proto3" json:"State,omitempty"`
	Nation   string `protobuf:"bytes,6,opt,name=Nation,proto3" json:"Nation,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetStreetNo() int32 {
	if x != nil {
		return x.StreetNo
	}
	return 0
}

func (x *Address) GetPinNo() int32 {
	if x != nil {
		return x.PinNo
	}
	return 0
}

func (x *Address) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetNation() string {
	if x != nil {
		return x.Nation
	}
	return ""
}

type RegisterCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address         *Address `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Emails          []string `protobuf:"bytes,3,rep,name=Emails,proto3" json:"Emails,omitempty"`
	Mobiles         []string `protobuf:"bytes,4,rep,name=Mobiles,proto3" json:"Mobiles,omitempty"`
	Companyusername string   `protobuf:"bytes,5,opt,name=Companyusername,proto3" json:"Companyusername,omitempty"`
	TypeID          uint32   `protobuf:"varint,6,opt,name=TypeID,proto3" json:"TypeID,omitempty"`
}

func (x *RegisterCompanyRequest) Reset() {
	*x = RegisterCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCompanyRequest) ProtoMessage() {}

func (x *RegisterCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCompanyRequest.ProtoReflect.Descriptor instead.
func (*RegisterCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterCompanyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterCompanyRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RegisterCompanyRequest) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *RegisterCompanyRequest) GetMobiles() []string {
	if x != nil {
		return x.Mobiles
	}
	return nil
}

func (x *RegisterCompanyRequest) GetCompanyusername() string {
	if x != nil {
		return x.Companyusername
	}
	return ""
}

func (x *RegisterCompanyRequest) GetTypeID() uint32 {
	if x != nil {
		return x.TypeID
	}
	return 0
}

type GetCompanyTypesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetCompanyTypesRes) Reset() {
	*x = GetCompanyTypesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyTypesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyTypesRes) ProtoMessage() {}

func (x *GetCompanyTypesRes) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyTypesRes.ProtoReflect.Descriptor instead.
func (*GetCompanyTypesRes) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{2}
}

func (x *GetCompanyTypesRes) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetCompanyTypesRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CompanyResponce struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address         *Address `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Emails          []string `protobuf:"bytes,3,rep,name=Emails,proto3" json:"Emails,omitempty"`
	Mobiles         []string `protobuf:"bytes,4,rep,name=Mobiles,proto3" json:"Mobiles,omitempty"`
	Companyusername string   `protobuf:"bytes,5,opt,name=Companyusername,proto3" json:"Companyusername,omitempty"`
	CompanyID       string   `protobuf:"bytes,6,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
}

func (x *CompanyResponce) Reset() {
	*x = CompanyResponce{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyResponce) ProtoMessage() {}

func (x *CompanyResponce) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyResponce.ProtoReflect.Descriptor instead.
func (*CompanyResponce) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{3}
}

func (x *CompanyResponce) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompanyResponce) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *CompanyResponce) GetEmails() []string {
	if x != nil {
		return x.Emails
	}
	return nil
}

func (x *CompanyResponce) GetMobiles() []string {
	if x != nil {
		return x.Mobiles
	}
	return nil
}

func (x *CompanyResponce) GetCompanyusername() string {
	if x != nil {
		return x.Companyusername
	}
	return ""
}

func (x *CompanyResponce) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Permission string `protobuf:"bytes,2,opt,name=Permission,proto3" json:"Permission,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{4}
}

func (x *Permission) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Permission) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

type AddEmployeeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string               `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	RoleID    uint32               `protobuf:"varint,3,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Deadline  *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Deadline,proto3" json:"Deadline,omitempty"`
	CompanyID string               `protobuf:"bytes,5,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
}

func (x *AddEmployeeReq) Reset() {
	*x = AddEmployeeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEmployeeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmployeeReq) ProtoMessage() {}

func (x *AddEmployeeReq) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmployeeReq.ProtoReflect.Descriptor instead.
func (*AddEmployeeReq) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{5}
}

func (x *AddEmployeeReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddEmployeeReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *AddEmployeeReq) GetDeadline() *timestamp.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *AddEmployeeReq) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type AttachRoleWithPermisssionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleID       uint32 `protobuf:"varint,1,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	PermissionID uint32 `protobuf:"varint,2,opt,name=PermissionID,proto3" json:"PermissionID,omitempty"`
	CompanyID    string `protobuf:"bytes,3,opt,name=CompanyID,proto3" json:"CompanyID,omitempty"`
}

func (x *AttachRoleWithPermisssionsReq) Reset() {
	*x = AttachRoleWithPermisssionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttachRoleWithPermisssionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttachRoleWithPermisssionsReq) ProtoMessage() {}

func (x *AttachRoleWithPermisssionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttachRoleWithPermisssionsReq.ProtoReflect.Descriptor instead.
func (*AttachRoleWithPermisssionsReq) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{6}
}

func (x *AttachRoleWithPermisssionsReq) GetRoleID() uint32 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *AttachRoleWithPermisssionsReq) GetPermissionID() uint32 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

func (x *AttachRoleWithPermisssionsReq) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type GetAttachedRoleswithPermissionsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID string `protobuf:"bytes,1,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *GetAttachedRoleswithPermissionsReq) Reset() {
	*x = GetAttachedRoleswithPermissionsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachedRoleswithPermissionsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachedRoleswithPermissionsReq) ProtoMessage() {}

func (x *GetAttachedRoleswithPermissionsReq) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachedRoleswithPermissionsReq.ProtoReflect.Descriptor instead.
func (*GetAttachedRoleswithPermissionsReq) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{7}
}

func (x *GetAttachedRoleswithPermissionsReq) GetCompanyID() string {
	if x != nil {
		return x.CompanyID
	}
	return ""
}

type GetAttachedRoleswithPermissionsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Role         string `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
	PermissionID uint32 `protobuf:"varint,3,opt,name=PermissionID,proto3" json:"PermissionID,omitempty"`
}

func (x *GetAttachedRoleswithPermissionsRes) Reset() {
	*x = GetAttachedRoleswithPermissionsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAttachedRoleswithPermissionsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAttachedRoleswithPermissionsRes) ProtoMessage() {}

func (x *GetAttachedRoleswithPermissionsRes) ProtoReflect() protoreflect.Message {
	mi := &file_company_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAttachedRoleswithPermissionsRes.ProtoReflect.Descriptor instead.
func (*GetAttachedRoleswithPermissionsRes) Descriptor() ([]byte, []int) {
	return file_company_proto_rawDescGZIP(), []int{8}
}

func (x *GetAttachedRoleswithPermissionsRes) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetAttachedRoleswithPermissionsRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetAttachedRoleswithPermissionsRes) GetPermissionID() uint32 {
	if x != nil {
		return x.PermissionID
	}
	return 0
}

var File_company_proto protoreflect.FileDescriptor

var file_company_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x69, 0x6e, 0x4e, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x69, 0x6e, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xcb, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x3c, 0x0a,
	0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x08,
	0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x44, 0x22, 0x79, 0x0a, 0x1d, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x42, 0x0a,
	0x22, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x44, 0x22, 0x6c, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x32,
	0x85, 0x04, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x3f, 0x0a, 0x0c, 0x41,
	0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5c, 0x0a, 0x1a,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x52, 0x6f, 0x6c, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x7b, 0x0a, 0x1f, 0x47, 0x65,
	0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x77, 0x69,
	0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x77, 0x69, 0x74, 0x68, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_company_proto_rawDescOnce sync.Once
	file_company_proto_rawDescData = file_company_proto_rawDesc
)

func file_company_proto_rawDescGZIP() []byte {
	file_company_proto_rawDescOnce.Do(func() {
		file_company_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_proto_rawDescData)
	})
	return file_company_proto_rawDescData
}

var file_company_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_company_proto_goTypes = []interface{}{
	(*Address)(nil),                            // 0: company.Address
	(*RegisterCompanyRequest)(nil),             // 1: company.RegisterCompanyRequest
	(*GetCompanyTypesRes)(nil),                 // 2: company.GetCompanyTypesRes
	(*CompanyResponce)(nil),                    // 3: company.CompanyResponce
	(*Permission)(nil),                         // 4: company.Permission
	(*AddEmployeeReq)(nil),                     // 5: company.AddEmployeeReq
	(*AttachRoleWithPermisssionsReq)(nil),      // 6: company.AttachRoleWithPermisssionsReq
	(*GetAttachedRoleswithPermissionsReq)(nil), // 7: company.GetAttachedRoleswithPermissionsReq
	(*GetAttachedRoleswithPermissionsRes)(nil), // 8: company.GetAttachedRoleswithPermissionsRes
	(*timestamp.Timestamp)(nil),                // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),                        // 10: google.protobuf.Empty
}
var file_company_proto_depIdxs = []int32{
	0,  // 0: company.RegisterCompanyRequest.Address:type_name -> company.Address
	0,  // 1: company.CompanyResponce.Address:type_name -> company.Address
	9,  // 2: company.AddEmployeeReq.Deadline:type_name -> google.protobuf.Timestamp
	1,  // 3: company.CompanyService.RegisterCompany:input_type -> company.RegisterCompanyRequest
	10, // 4: company.CompanyService.GetCompanyTypes:input_type -> google.protobuf.Empty
	10, // 5: company.CompanyService.GetPermissions:input_type -> google.protobuf.Empty
	5,  // 6: company.CompanyService.AddEmployees:input_type -> company.AddEmployeeReq
	6,  // 7: company.CompanyService.AttachRoleWithPermisssions:input_type -> company.AttachRoleWithPermisssionsReq
	7,  // 8: company.CompanyService.GetAttachedRoleswithPermissions:input_type -> company.GetAttachedRoleswithPermissionsReq
	3,  // 9: company.CompanyService.RegisterCompany:output_type -> company.CompanyResponce
	2,  // 10: company.CompanyService.GetCompanyTypes:output_type -> company.GetCompanyTypesRes
	4,  // 11: company.CompanyService.GetPermissions:output_type -> company.Permission
	10, // 12: company.CompanyService.AddEmployees:output_type -> google.protobuf.Empty
	10, // 13: company.CompanyService.AttachRoleWithPermisssions:output_type -> google.protobuf.Empty
	8,  // 14: company.CompanyService.GetAttachedRoleswithPermissions:output_type -> company.GetAttachedRoleswithPermissionsRes
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_company_proto_init() }
func file_company_proto_init() {
	if File_company_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_company_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCompanyRequest); i {
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
		file_company_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyTypesRes); i {
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
		file_company_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyResponce); i {
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
		file_company_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_company_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEmployeeReq); i {
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
		file_company_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttachRoleWithPermisssionsReq); i {
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
		file_company_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachedRoleswithPermissionsReq); i {
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
		file_company_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAttachedRoleswithPermissionsRes); i {
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
			RawDescriptor: file_company_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_proto_goTypes,
		DependencyIndexes: file_company_proto_depIdxs,
		MessageInfos:      file_company_proto_msgTypes,
	}.Build()
	File_company_proto = out.File
	file_company_proto_rawDesc = nil
	file_company_proto_goTypes = nil
	file_company_proto_depIdxs = nil
}
