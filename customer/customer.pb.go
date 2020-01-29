// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

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
type Request struct {
	AccountNumber        string   `protobuf:"bytes,1,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
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

func (m *Request) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

type Account struct {
	Account              *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{3}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type Response struct {
	Created  bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Customer *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// Added a pluralised consignment to our generic response message
	// array of these types
	Customers            []*Customer `protobuf:"bytes,3,rep,name=customers,proto3" json:"customers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{4}
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

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

func (m *Response) GetCustomers() []*Customer {
	if m != nil {
		return m.Customers
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "moneytree.svc.customer.Customer")
	proto.RegisterType((*Address)(nil), "moneytree.svc.customer.Address")
	proto.RegisterType((*Request)(nil), "moneytree.svc.customer.Request")
	proto.RegisterType((*Account)(nil), "moneytree.svc.customer.Account")
	proto.RegisterType((*Response)(nil), "moneytree.svc.customer.Response")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x6e, 0xb2, 0xdb, 0x26, 0x79, 0xbb, 0x4d, 0x61, 0x04, 0x19, 0x2b, 0xda, 0x25, 0x2a, 0xec,
	0x29, 0xd8, 0x78, 0xf1, 0xa0, 0x42, 0xbb, 0x07, 0x11, 0xa4, 0x87, 0x14, 0x0a, 0x7a, 0x09, 0x69,
	0xf2, 0xc4, 0x81, 0x4d, 0x26, 0xce, 0xcc, 0x16, 0xb6, 0x3f, 0xc0, 0x9f, 0xe1, 0xc1, 0x7f, 0xe8,
	0x3f, 0x90, 0xbe, 0xcc, 0x6c, 0x17, 0xdd, 0xa5, 0x3d, 0xf4, 0x36, 0xf3, 0xbd, 0xef, 0xfb, 0x78,
	0xf3, 0xde, 0x97, 0x40, 0x5c, 0x2d, 0xb4, 0x91, 0x0d, 0xaa, 0xb4, 0x53, 0xd2, 0x48, 0xf6, 0xb8,
	0x91, 0x2d, 0x2e, 0x8d, 0x42, 0x4c, 0xf5, 0x55, 0x95, 0xba, 0xea, 0xe1, 0x7e, 0x59, 0x55, 0x72,
	0xd1, 0x9a, 0x9e, 0x96, 0xfc, 0xf4, 0x21, 0x9c, 0xd9, 0x1a, 0x8b, 0xc1, 0x17, 0x35, 0xf7, 0x26,
	0xde, 0x34, 0xca, 0x7d, 0x51, 0xb3, 0x67, 0x00, 0xdf, 0x84, 0xd2, 0xa6, 0x68, 0xcb, 0x06, 0xb9,
	0x4f, 0x78, 0x44, 0xc8, 0x59, 0xd9, 0x20, 0x7b, 0x0a, 0xd1, 0xbc, 0x74, 0xd5, 0x01, 0x55, 0xc3,
	0x1b, 0x80, 0x8a, 0xef, 0x21, 0x2a, 0xeb, 0x5a, 0xa1, 0xd6, 0xa8, 0xf9, 0x70, 0x32, 0x98, 0x8e,
	0xb2, 0xa3, 0x74, 0x73, 0x4f, 0xe9, 0x49, 0x4f, 0xcc, 0x6f, 0x15, 0xec, 0x05, 0xec, 0x77, 0xdf,
	0x65, 0x8b, 0x45, 0xbb, 0x68, 0x2e, 0x51, 0x69, 0xbe, 0x3b, 0x19, 0x4c, 0xa3, 0x7c, 0x4c, 0xe0,
	0x59, 0x8f, 0xdd, 0x90, 0xb0, 0x29, 0xc5, 0xbc, 0xb0, 0x3a, 0xbe, 0x47, 0x4d, 0x8c, 0x09, 0xb4,
	0xa6, 0xec, 0x15, 0xc4, 0xf6, 0xc9, 0xd6, 0x8b, 0x07, 0xc4, 0x72, 0x83, 0xe8, 0xcd, 0x92, 0x5f,
	0x1e, 0x04, 0x4e, 0xf2, 0x12, 0x62, 0xeb, 0x58, 0xcc, 0x45, 0x8b, 0xc5, 0xb1, 0x9d, 0xc9, 0xd8,
	0xa2, 0x9f, 0x45, 0x8b, 0xc7, 0xff, 0xb1, 0x32, 0x3b, 0xa1, 0x75, 0x56, 0xc6, 0x9e, 0x40, 0x78,
	0x2d, 0xba, 0xa2, 0x92, 0xb5, 0x9b, 0x51, 0x70, 0x2d, 0xba, 0x99, 0xac, 0x91, 0x31, 0x18, 0x56,
	0xc2, 0x2c, 0xf9, 0x90, 0x60, 0x3a, 0x33, 0x0e, 0x01, 0x75, 0xa5, 0x96, 0x7c, 0xb7, 0x67, 0xdb,
	0x6b, 0xf2, 0x1a, 0x82, 0x1c, 0x7f, 0x2c, 0x50, 0x9b, 0x0d, 0x4f, 0xf2, 0x36, 0x3d, 0x69, 0x06,
	0xc1, 0x49, 0x0f, 0xb0, 0xb7, 0x10, 0xd8, 0x1a, 0x51, 0x47, 0xd9, 0xf3, 0x7f, 0x76, 0xe1, 0x52,
	0x61, 0x05, 0xb9, 0xa3, 0x27, 0xbf, 0x3d, 0x08, 0x73, 0xd4, 0x9d, 0x6c, 0x35, 0x52, 0x77, 0x0a,
	0x4b, 0x83, 0x7d, 0x4a, 0xc2, 0xdc, 0x5d, 0xd9, 0x3b, 0x08, 0xdd, 0x3a, 0x69, 0x0c, 0xa3, 0x6c,
	0xb2, 0x6d, 0xdb, 0x2e, 0x6e, 0xf9, 0x4a, 0xc1, 0x3e, 0x40, 0xe4, 0xce, 0x9a, 0x0f, 0x28, 0x2c,
	0x77, 0xcb, 0x6f, 0x25, 0xd9, 0x1f, 0x1f, 0x0e, 0x1c, 0x7e, 0x8e, 0xea, 0x4a, 0x54, 0xc8, 0x2e,
	0x20, 0x9e, 0x51, 0x73, 0xab, 0x78, 0xdf, 0x69, 0x79, 0xb8, 0x95, 0xe1, 0x26, 0x90, 0xec, 0xb0,
	0x73, 0x18, 0x7f, 0x44, 0xe3, 0x24, 0x9a, 0x1d, 0x6d, 0xd7, 0xd0, 0xb6, 0xee, 0x65, 0x7a, 0x01,
	0x07, 0x6b, 0xa6, 0xa7, 0xcb, 0x4f, 0xf5, 0xc3, 0xf8, 0x7e, 0x81, 0x47, 0x6b, 0xbe, 0x76, 0xb9,
	0xf7, 0xe8, 0x79, 0xfb, 0xa7, 0x6a, 0x63, 0xb1, 0x73, 0x0a, 0x5f, 0x57, 0xfb, 0xbb, 0xdc, 0xa3,
	0x9f, 0xc9, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04, 0x3e, 0xb4, 0x7d, 0x85, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for CustomerService service

type CustomerServiceClient interface {
	CreateCustomer(ctx context.Context, in *Customer, opts ...client.CallOption) (*Response, error)
	// Created a new method
	GetCustomers(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetCustomerById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetCustomerAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Account, error)
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

func (c *customerServiceClient) CreateCustomer(ctx context.Context, in *Customer, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.CreateCustomer", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomers(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomers", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerById(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomerById", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetCustomerAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Account, error) {
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
	CreateCustomer(context.Context, *Customer, *Response) error
	// Created a new method
	GetCustomers(context.Context, *Request, *Response) error
	GetCustomerById(context.Context, *Request, *Response) error
	GetCustomerAccounts(context.Context, *Request, *Account) error
}

func RegisterCustomerServiceHandler(s server.Server, hdlr CustomerServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&CustomerService{hdlr}, opts...))
}

type CustomerService struct {
	CustomerServiceHandler
}

func (h *CustomerService) CreateCustomer(ctx context.Context, in *Customer, out *Response) error {
	return h.CustomerServiceHandler.CreateCustomer(ctx, in, out)
}

func (h *CustomerService) GetCustomers(ctx context.Context, in *Request, out *Response) error {
	return h.CustomerServiceHandler.GetCustomers(ctx, in, out)
}

func (h *CustomerService) GetCustomerById(ctx context.Context, in *Request, out *Response) error {
	return h.CustomerServiceHandler.GetCustomerById(ctx, in, out)
}

func (h *CustomerService) GetCustomerAccounts(ctx context.Context, in *Request, out *Account) error {
	return h.CustomerServiceHandler.GetCustomerAccounts(ctx, in, out)
}
