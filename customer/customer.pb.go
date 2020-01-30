// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Customer struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	//repeated fields are used for arrays or list
	//multiple addresses
	Addresses            []*Address `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	PhoneNumbers         []string   `protobuf:"bytes,5,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	EmailAddress         string     `protobuf:"bytes,6,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	AccountNumber        string     `protobuf:"bytes,7,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Customer) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Customer) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Customer) GetAddresses() []*Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *Customer) GetPhoneNumbers() []string {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func (m *Customer) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Customer) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

//TODO these can be imported
//Address type
type Address struct {
	AddressLine_1        string   `protobuf:"bytes,1,opt,name=address_line_1,json=addressLine1,proto3" json:"address_line_1,omitempty"`
	AddressLine_2        string   `protobuf:"bytes,2,opt,name=address_line_2,json=addressLine2,proto3" json:"address_line_2,omitempty"`
	ZipCode              string   `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAddressLine_1() string {
	if m != nil {
		return m.AddressLine_1
	}
	return ""
}

func (m *Address) GetAddressLine_2() string {
	if m != nil {
		return m.AddressLine_2
	}
	return ""
}

func (m *Address) GetZipCode() string {
	if m != nil {
		return m.ZipCode
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

// Created a blank get request
type CustomerRequest struct {
	AccountNumber        string   `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerRequest) Reset()         { *m = CustomerRequest{} }
func (m *CustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerRequest) ProtoMessage()    {}
func (*CustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *CustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerRequest.Unmarshal(m, b)
}
func (m *CustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerRequest.Marshal(b, m, deterministic)
}
func (m *CustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerRequest.Merge(m, src)
}
func (m *CustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CustomerRequest.Size(m)
}
func (m *CustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerRequest proto.InternalMessageInfo

func (m *CustomerRequest) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

type CustomerResponse struct {
	Created  bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Customer *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// Added a pluralised consignment to our generic response message
	// array of these types
	Customers            []*Customer `protobuf:"bytes,3,rep,name=customers,proto3" json:"customers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CustomerResponse) Reset()         { *m = CustomerResponse{} }
func (m *CustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CustomerResponse) ProtoMessage()    {}
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *CustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerResponse.Unmarshal(m, b)
}
func (m *CustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerResponse.Marshal(b, m, deterministic)
}
func (m *CustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerResponse.Merge(m, src)
}
func (m *CustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CustomerResponse.Size(m)
}
func (m *CustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerResponse proto.InternalMessageInfo

func (m *CustomerResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *CustomerResponse) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *CustomerResponse) GetCustomers() []*Customer {
	if m != nil {
		return m.Customers
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "moneytree.svc.customer.Customer")
	proto.RegisterType((*Address)(nil), "moneytree.svc.customer.Address")
	proto.RegisterType((*CustomerRequest)(nil), "moneytree.svc.customer.CustomerRequest")
	proto.RegisterType((*CustomerResponse)(nil), "moneytree.svc.customer.CustomerResponse")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xa4, 0xb6, 0xa7, 0xa9, 0x8b, 0x16, 0x09, 0x2d, 0x45, 0x40, 0x64, 0x40, 0xe4,
	0x64, 0xa9, 0xe6, 0xc2, 0x01, 0x90, 0xda, 0x1c, 0x10, 0x12, 0xea, 0xc1, 0xdc, 0x38, 0x60, 0x5c,
	0x7b, 0x2a, 0x56, 0x8a, 0xbd, 0x66, 0x77, 0x53, 0x29, 0xfd, 0x01, 0xfc, 0x0c, 0xc4, 0x99, 0x5f,
	0x89, 0x3a, 0xde, 0x4d, 0x22, 0x08, 0x8a, 0x7a, 0xe8, 0xc9, 0xde, 0x37, 0xef, 0xbd, 0xf9, 0x5a,
	0x2d, 0xc4, 0xd5, 0x42, 0x1b, 0xd9, 0xa0, 0x4a, 0x3b, 0x25, 0x8d, 0x64, 0x0f, 0x1a, 0xd9, 0xe2,
	0xd2, 0x28, 0xc4, 0x54, 0x5f, 0x55, 0xa9, 0x8b, 0x1e, 0x1f, 0x96, 0x55, 0x25, 0x17, 0xad, 0xe9,
	0x69, 0xc9, 0x8f, 0x01, 0x84, 0x33, 0x1b, 0x63, 0x31, 0x0c, 0x44, 0xcd, 0xbd, 0x89, 0x37, 0x8d,
	0xf2, 0x81, 0xa8, 0xd9, 0x63, 0x80, 0x4b, 0xa1, 0xb4, 0x29, 0xda, 0xb2, 0x41, 0x3e, 0x20, 0x3c,
	0x22, 0xe4, 0xbc, 0x6c, 0x90, 0x3d, 0x82, 0x68, 0x5e, 0xba, 0xa8, 0x4f, 0xd1, 0xf0, 0x06, 0xa0,
	0xe0, 0x5b, 0x88, 0xca, 0xba, 0x56, 0xa8, 0x35, 0x6a, 0x3e, 0x9c, 0xf8, 0xd3, 0x83, 0xec, 0x69,
	0xba, 0xbd, 0xa6, 0xf4, 0xb4, 0x27, 0xe6, 0x6b, 0x05, 0x7b, 0x06, 0x87, 0xdd, 0x37, 0xd9, 0x62,
	0xd1, 0x2e, 0x9a, 0x0b, 0x54, 0x9a, 0x8f, 0x26, 0xfe, 0x34, 0xca, 0xc7, 0x04, 0x9e, 0xf7, 0xd8,
	0x0d, 0x09, 0x9b, 0x52, 0xcc, 0x0b, 0xab, 0xe3, 0xfb, 0x54, 0xc4, 0x98, 0x40, 0x6b, 0xca, 0x5e,
	0x40, 0x6c, 0x5b, 0xb6, 0x5e, 0x3c, 0x20, 0x96, 0x1b, 0x44, 0x6f, 0x96, 0xfc, 0xf4, 0x20, 0x70,
	0x92, 0xe7, 0x10, 0x5b, 0xc7, 0x62, 0x2e, 0x5a, 0x2c, 0x4e, 0xec, 0x4c, 0xc6, 0x16, 0xfd, 0x28,
	0x5a, 0x3c, 0xf9, 0x87, 0x95, 0xd9, 0x09, 0x6d, 0xb2, 0x32, 0xf6, 0x10, 0xc2, 0x6b, 0xd1, 0x15,
	0x95, 0xac, 0xdd, 0x8c, 0x82, 0x6b, 0xd1, 0xcd, 0x64, 0x8d, 0x8c, 0xc1, 0xb0, 0x12, 0x66, 0xc9,
	0x87, 0x04, 0xd3, 0x3f, 0xe3, 0x10, 0x50, 0x55, 0x6a, 0xc9, 0x47, 0x3d, 0xdb, 0x1e, 0x93, 0xd7,
	0x70, 0xe4, 0x16, 0x95, 0xe3, 0xf7, 0x05, 0x6a, 0xb3, 0xa5, 0x35, 0x6f, 0x5b, 0x6b, 0xbf, 0x3d,
	0xb8, 0xb7, 0x96, 0xea, 0x4e, 0xb6, 0x1a, 0x29, 0x91, 0xc2, 0xd2, 0x60, 0xbf, 0xf0, 0x30, 0x77,
	0x47, 0xf6, 0x06, 0x42, 0xb7, 0x19, 0xea, 0xe8, 0x20, 0x9b, 0xfc, 0x6f, 0x71, 0x2b, 0xd7, 0x95,
	0x82, 0xbd, 0x83, 0xc8, 0xfd, 0x6b, 0xee, 0xd3, 0xde, 0x77, 0xcb, 0xd7, 0x92, 0xec, 0x97, 0xbf,
	0xee, 0xf3, 0x13, 0xaa, 0x2b, 0x51, 0x21, 0xfb, 0x02, 0xf1, 0x8c, 0x8a, 0x5b, 0xdd, 0xd4, 0x9d,
	0x96, 0xc7, 0xd3, 0x9d, 0x49, 0xed, 0x24, 0x92, 0x3d, 0x56, 0xc1, 0xf8, 0x3d, 0x1a, 0x17, 0xd0,
	0xec, 0xe5, 0x6e, 0x2d, 0x2d, 0xe0, 0x56, 0x49, 0x2e, 0xe1, 0x68, 0x23, 0xc9, 0xd9, 0xf2, 0x43,
	0x7d, 0x37, 0x79, 0xbe, 0xc2, 0xfd, 0x8d, 0x3c, 0xa7, 0xfd, 0x4d, 0xb8, 0x45, 0x4f, 0x4f, 0xfe,
	0x22, 0xba, 0xf7, 0xc2, 0x3a, 0x25, 0x7b, 0x67, 0xc1, 0xe7, 0x11, 0x3d, 0x1e, 0x17, 0xfb, 0xf4,
	0x79, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x30, 0x88, 0x28, 0xc0, 0x7c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CustomerService service

type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error)
	// Created a new method
	GetCustomers(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	GetCustomerById(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error)
	GetCustomerAccounts(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*Account, error)
}

type customerServiceClient struct {
	c           client.Client
	serviceName string
}

func NewCustomerServiceClient(serviceName string, c client.Client) CustomerServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "moneytree.svc.customer"
	}
	return &customerServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.CreateCustomer", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomers", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerById(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*CustomerResponse, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomerById", in)
	out := new(CustomerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerAccounts(ctx context.Context, in *CustomerRequest, opts ...client.CallOption) (*Account, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomerAccounts", in)
	out := new(Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CustomerService service

type CustomerServiceHandler interface {
	CreateCustomer(context.Context, *Customer, *CustomerResponse) error
	// Created a new method
	GetCustomers(context.Context, *CustomerRequest, *CustomerResponse) error
	GetCustomerById(context.Context, *CustomerRequest, *CustomerResponse) error
	GetCustomerAccounts(context.Context, *CustomerRequest, *Account) error
}

func RegisterCustomerServiceHandler(s server.Server, hdlr CustomerServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CustomerService{hdlr}, opts...))
}

type CustomerService struct {
	CustomerServiceHandler
}

func (h *CustomerService) CreateCustomer(ctx context.Context, in *Customer, out *CustomerResponse) error {
	return h.CustomerServiceHandler.CreateCustomer(ctx, in, out)
}

func (h *CustomerService) GetCustomers(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.GetCustomers(ctx, in, out)
}

func (h *CustomerService) GetCustomerById(ctx context.Context, in *CustomerRequest, out *CustomerResponse) error {
	return h.CustomerServiceHandler.GetCustomerById(ctx, in, out)
}

func (h *CustomerService) GetCustomerAccounts(ctx context.Context, in *CustomerRequest, out *Account) error {
	return h.CustomerServiceHandler.GetCustomerAccounts(ctx, in, out)
}
