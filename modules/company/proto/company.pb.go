// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/company/proto/company.proto

package companypb

import (
	context "context"
	fmt "fmt"
	_ "github.com/fzerorubigd/protobuf/extra"
	types "github.com/fzerorubigd/protobuf/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CompanyStatus int32

const (
	CompanyStatus_COMPANY_STATUS_INVALID CompanyStatus = 0
	CompanyStatus_COMPANY_STATUS_PENDING CompanyStatus = 1
	CompanyStatus_COMPANY_STATUS_ACTIVE  CompanyStatus = 2
	CompanyStatus_COMPANY_STATUS_BANNED  CompanyStatus = 3
)

var CompanyStatus_name = map[int32]string{
	0: "COMPANY_STATUS_INVALID",
	1: "COMPANY_STATUS_PENDING",
	2: "COMPANY_STATUS_ACTIVE",
	3: "COMPANY_STATUS_BANNED",
}

var CompanyStatus_value = map[string]int32{
	"COMPANY_STATUS_INVALID": 0,
	"COMPANY_STATUS_PENDING": 1,
	"COMPANY_STATUS_ACTIVE":  2,
	"COMPANY_STATUS_BANNED":  3,
}

func (x CompanyStatus) String() string {
	return proto.EnumName(CompanyStatus_name, int32(x))
}

func (CompanyStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{0}
}

type Company struct {
	Id                   int64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" db:"id" `
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" db:"name" `
	Status               CompanyStatus    `protobuf:"varint,3,opt,name=status,proto3,enum=company.CompanyStatus" json:"status,omitempty" db:"status" `
	CreatedAt            *types.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" db:"created_at" `
	UpdatedAt            *types.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" db:"updated_at" `
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Company) Reset()         { *m = Company{} }
func (m *Company) String() string { return proto.CompactTextString(m) }
func (*Company) ProtoMessage()    {}
func (*Company) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{0}
}
func (m *Company) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Company.Unmarshal(m, b)
}
func (m *Company) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Company.Marshal(b, m, deterministic)
}
func (m *Company) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Company.Merge(m, src)
}
func (m *Company) XXX_Size() int {
	return xxx_messageInfo_Company.Size(m)
}
func (m *Company) XXX_DiscardUnknown() {
	xxx_messageInfo_Company.DiscardUnknown(m)
}

var xxx_messageInfo_Company proto.InternalMessageInfo

func (m *Company) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Company) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Company) GetStatus() CompanyStatus {
	if m != nil {
		return m.Status
	}
	return CompanyStatus_COMPANY_STATUS_INVALID
}

func (m *Company) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Company) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GetCompanyRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCompanyRequest) Reset()         { *m = GetCompanyRequest{} }
func (m *GetCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCompanyRequest) ProtoMessage()    {}
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{1}
}
func (m *GetCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompanyRequest.Unmarshal(m, b)
}
func (m *GetCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompanyRequest.Marshal(b, m, deterministic)
}
func (m *GetCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompanyRequest.Merge(m, src)
}
func (m *GetCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCompanyRequest.Size(m)
}
func (m *GetCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompanyRequest proto.InternalMessageInfo

func (m *GetCompanyRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetCompanyResponse struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               CompanyStatus `protobuf:"varint,3,opt,name=status,proto3,enum=company.CompanyStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetCompanyResponse) Reset()         { *m = GetCompanyResponse{} }
func (m *GetCompanyResponse) String() string { return proto.CompactTextString(m) }
func (*GetCompanyResponse) ProtoMessage()    {}
func (*GetCompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{2}
}
func (m *GetCompanyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCompanyResponse.Unmarshal(m, b)
}
func (m *GetCompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCompanyResponse.Marshal(b, m, deterministic)
}
func (m *GetCompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCompanyResponse.Merge(m, src)
}
func (m *GetCompanyResponse) XXX_Size() int {
	return xxx_messageInfo_GetCompanyResponse.Size(m)
}
func (m *GetCompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCompanyResponse proto.InternalMessageInfo

func (m *GetCompanyResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetCompanyResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetCompanyResponse) GetStatus() CompanyStatus {
	if m != nil {
		return m.Status
	}
	return CompanyStatus_COMPANY_STATUS_INVALID
}

type CreateCompanyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCompanyRequest) Reset()         { *m = CreateCompanyRequest{} }
func (m *CreateCompanyRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyRequest) ProtoMessage()    {}
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{3}
}
func (m *CreateCompanyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyRequest.Unmarshal(m, b)
}
func (m *CreateCompanyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyRequest.Marshal(b, m, deterministic)
}
func (m *CreateCompanyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyRequest.Merge(m, src)
}
func (m *CreateCompanyRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyRequest.Size(m)
}
func (m *CreateCompanyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyRequest proto.InternalMessageInfo

func (m *CreateCompanyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateCompanyResponse struct {
	Id                   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               CompanyStatus `protobuf:"varint,3,opt,name=status,proto3,enum=company.CompanyStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateCompanyResponse) Reset()         { *m = CreateCompanyResponse{} }
func (m *CreateCompanyResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCompanyResponse) ProtoMessage()    {}
func (*CreateCompanyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_24f2033e12fc8969, []int{4}
}
func (m *CreateCompanyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCompanyResponse.Unmarshal(m, b)
}
func (m *CreateCompanyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCompanyResponse.Marshal(b, m, deterministic)
}
func (m *CreateCompanyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCompanyResponse.Merge(m, src)
}
func (m *CreateCompanyResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCompanyResponse.Size(m)
}
func (m *CreateCompanyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCompanyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCompanyResponse proto.InternalMessageInfo

func (m *CreateCompanyResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CreateCompanyResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCompanyResponse) GetStatus() CompanyStatus {
	if m != nil {
		return m.Status
	}
	return CompanyStatus_COMPANY_STATUS_INVALID
}

func init() {
	proto.RegisterEnum("company.CompanyStatus", CompanyStatus_name, CompanyStatus_value)
	proto.RegisterType((*Company)(nil), "company.Company")
	proto.RegisterType((*GetCompanyRequest)(nil), "company.GetCompanyRequest")
	proto.RegisterType((*GetCompanyResponse)(nil), "company.GetCompanyResponse")
	proto.RegisterType((*CreateCompanyRequest)(nil), "company.CreateCompanyRequest")
	proto.RegisterType((*CreateCompanyResponse)(nil), "company.CreateCompanyResponse")
}

func init() {
	proto.RegisterFile("modules/company/proto/company.proto", fileDescriptor_24f2033e12fc8969)
}

var fileDescriptor_24f2033e12fc8969 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x41, 0x4f, 0xd4, 0x40,
	0x18, 0xa5, 0x5d, 0x04, 0x19, 0x60, 0x53, 0x26, 0x0b, 0xd9, 0xad, 0xe8, 0x36, 0xc3, 0x65, 0x43,
	0x64, 0x2b, 0x18, 0x63, 0xc2, 0xc9, 0x76, 0xd9, 0x90, 0x4d, 0xb4, 0x6e, 0x96, 0x95, 0x44, 0x2f,
	0x64, 0xda, 0x0e, 0xa5, 0x91, 0x76, 0x6a, 0x3b, 0x45, 0xd1, 0xe0, 0xc1, 0xc4, 0x0b, 0x47, 0xbd,
	0xf8, 0x2b, 0x34, 0xf1, 0x57, 0x78, 0xf7, 0xe6, 0x61, 0x4f, 0xfe, 0x00, 0xb2, 0xbf, 0xc0, 0xb4,
	0x9d, 0x52, 0x58, 0xd6, 0xe0, 0xc5, 0xcb, 0xec, 0x7e, 0xfb, 0xde, 0xfb, 0xe6, 0x7d, 0x6f, 0x66,
	0x07, 0xac, 0x78, 0xd4, 0x8e, 0x0f, 0x49, 0xa4, 0x5a, 0xd4, 0x0b, 0xb0, 0x7f, 0xac, 0x06, 0x21,
	0x65, 0x34, 0xaf, 0x9a, 0x69, 0x05, 0xa7, 0x79, 0x29, 0xaf, 0x3b, 0x2e, 0x3b, 0x88, 0xcd, 0xa6,
	0x45, 0x3d, 0x75, 0xff, 0x2d, 0x09, 0x69, 0x18, 0x9b, 0xae, 0x63, 0x67, 0x22, 0x33, 0xde, 0x57,
	0xc9, 0x1b, 0x16, 0xe2, 0x6c, 0xcd, 0xb4, 0xf2, 0xbd, 0xeb, 0x24, 0xec, 0x38, 0x20, 0x51, 0xba,
	0x72, 0xc5, 0xda, 0x05, 0x85, 0x43, 0x1d, 0x5a, 0x50, 0x93, 0x2a, 0xf3, 0x97, 0x7c, 0xe3, 0xf4,
	0x65, 0x87, 0x52, 0xe7, 0x90, 0xa8, 0x38, 0x70, 0x55, 0xec, 0xfb, 0x94, 0x61, 0xe6, 0x52, 0x3f,
	0xe2, 0xe8, 0xdd, 0xf4, 0xc3, 0x5a, 0x73, 0x88, 0xbf, 0x16, 0xbd, 0xc6, 0x8e, 0x43, 0x42, 0x95,
	0x06, 0x29, 0xe3, 0x2a, 0x1b, 0x7d, 0x15, 0xc1, 0x74, 0x2b, 0x9b, 0x15, 0x2e, 0x03, 0xd1, 0xb5,
	0xab, 0x82, 0x22, 0x34, 0x4a, 0xfa, 0xdc, 0x70, 0x50, 0xbf, 0x69, 0x9b, 0x9b, 0xc8, 0xb5, 0x91,
	0xd2, 0x13, 0x5d, 0x1b, 0x22, 0x30, 0xe9, 0x63, 0x8f, 0x54, 0x45, 0x45, 0x68, 0xcc, 0xe8, 0xe5,
	0xe1, 0xa0, 0x0e, 0x12, 0x3c, 0xf9, 0x0d, 0x29, 0xbd, 0x14, 0x83, 0x8f, 0xc0, 0x54, 0xc4, 0x30,
	0x8b, 0xa3, 0x6a, 0x49, 0x11, 0x1a, 0xe5, 0x8d, 0xa5, 0x66, 0x1e, 0x2b, 0xdf, 0x63, 0x27, 0x45,
	0x75, 0x69, 0x38, 0xa8, 0xcf, 0x25, 0xea, 0x8c, 0x8d, 0x94, 0x1e, 0xd7, 0xc1, 0x36, 0x00, 0x56,
	0x48, 0x30, 0x23, 0xf6, 0x1e, 0x66, 0xd5, 0x49, 0x45, 0x68, 0xcc, 0x6e, 0x48, 0xcd, 0x34, 0xb1,
	0x66, 0xdf, 0xf5, 0x48, 0xc4, 0xb0, 0x17, 0xe8, 0x95, 0xe1, 0xa0, 0x2e, 0x25, 0xfa, 0x82, 0x8b,
	0x94, 0xde, 0x0c, 0xaf, 0x34, 0x96, 0xb4, 0x89, 0x03, 0x3b, 0x6f, 0x73, 0xe3, 0xba, 0x36, 0x05,
	0x37, 0x69, 0xc3, 0x2b, 0x8d, 0x6d, 0x2e, 0xfc, 0x38, 0xab, 0x09, 0xbf, 0xce, 0x6a, 0x33, 0xd9,
	0x18, 0x2e, 0x89, 0xd0, 0x0a, 0x58, 0xd8, 0x26, 0x8c, 0x8f, 0xd3, 0x23, 0xaf, 0x62, 0x12, 0x31,
	0x58, 0x2e, 0x92, 0x4b, 0xb2, 0x42, 0x07, 0x00, 0x5e, 0x24, 0x45, 0x01, 0xf5, 0x23, 0x32, 0xca,
	0x82, 0xf0, 0x62, 0xa2, 0x3c, 0xc1, 0xe6, 0xbf, 0x25, 0x98, 0xe7, 0x85, 0x56, 0x41, 0xa5, 0x95,
	0x4e, 0x3d, 0xe2, 0x28, 0xef, 0x2d, 0x14, 0xbd, 0xd1, 0x4b, 0xb0, 0x38, 0xc2, 0xfd, 0x7f, 0xc6,
	0x56, 0x4f, 0xc0, 0xfc, 0x25, 0x00, 0xca, 0x60, 0xa9, 0xf5, 0xf4, 0x49, 0x57, 0x33, 0x9e, 0xef,
	0xed, 0xf4, 0xb5, 0xfe, 0xb3, 0x9d, 0xbd, 0x8e, 0xb1, 0xab, 0x3d, 0xee, 0x6c, 0x49, 0x13, 0x63,
	0xb0, 0x6e, 0xdb, 0xd8, 0xea, 0x18, 0xdb, 0x92, 0x00, 0x6b, 0x60, 0x71, 0x04, 0xd3, 0x5a, 0xfd,
	0xce, 0x6e, 0x5b, 0x12, 0xc7, 0x40, 0xba, 0x66, 0x18, 0xed, 0x2d, 0xa9, 0xb4, 0xf1, 0x45, 0x04,
	0xe5, 0x7c, 0x7f, 0x12, 0x1e, 0xb9, 0x16, 0x81, 0xef, 0x01, 0x28, 0x0e, 0x05, 0xca, 0xe7, 0xfe,
	0xaf, 0x1c, 0xa7, 0x7c, 0x6b, 0x2c, 0x96, 0x85, 0x85, 0x1e, 0x7e, 0xd2, 0x96, 0xcc, 0x0a, 0x80,
	0xa0, 0xac, 0xc5, 0xec, 0x80, 0xf8, 0xcc, 0xb5, 0xd2, 0x7f, 0x14, 0x9c, 0x38, 0xfd, 0x26, 0x4f,
	0x7c, 0xf8, 0xf9, 0xfb, 0xb3, 0x58, 0x81, 0x50, 0x3d, 0x5a, 0x57, 0xcf, 0x2f, 0x8d, 0xfa, 0xce,
	0xb5, 0x4f, 0xe0, 0x47, 0x01, 0xcc, 0x5f, 0xca, 0x1f, 0xde, 0x2e, 0x32, 0x1c, 0x73, 0x86, 0xf2,
	0x9d, 0xbf, 0xc1, 0xdc, 0xc9, 0x83, 0xeb, 0x9d, 0x40, 0x34, 0x7f, 0xc9, 0xc9, 0xa6, 0xb0, 0xaa,
	0xab, 0xa7, 0xdf, 0x6b, 0x25, 0x8c, 0x31, 0x98, 0xb5, 0xa8, 0x97, 0x6f, 0xa1, 0xcf, 0xf1, 0xee,
	0xdd, 0xe4, 0x59, 0xe8, 0x0a, 0x2f, 0xf8, 0x95, 0x3f, 0x0e, 0x4c, 0x73, 0x2a, 0x7d, 0x2a, 0xee,
	0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x27, 0x95, 0xf3, 0x18, 0x3a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompanyServiceClient interface {
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
}

type companyServiceClient struct {
	cc *grpc.ClientConn
}

func NewCompanyServiceClient(cc *grpc.ClientConn) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/GetCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, "/company.CompanyService/CreateCompany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
type CompanyServiceServer interface {
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
}

// UnimplementedCompanyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompanyServiceServer struct {
}

func (*UnimplementedCompanyServiceServer) GetCompany(ctx context.Context, req *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (*UnimplementedCompanyServiceServer) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}

func RegisterCompanyServiceServer(s *grpc.Server, srv CompanyServiceServer) {
	s.RegisterService(&_CompanyService_serviceDesc, srv)
}

func _CompanyService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/GetCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.CompanyService/CreateCompany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompanyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompany",
			Handler:    _CompanyService_GetCompany_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/company/proto/company.proto",
}
