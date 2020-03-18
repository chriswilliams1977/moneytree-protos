// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package customer

import (
	fmt "fmt"
	account "github.com/chriswilliams1977/moneytree-protos/account"
	address "github.com/chriswilliams1977/moneytree-protos/address"
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
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string             `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string             `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Addresses            []*address.Address `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	PhoneNumbers         []string           `protobuf:"bytes,5,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	EmailAddress         string             `protobuf:"bytes,6,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	AccountNumber        string             `protobuf:"bytes,7,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Uuid                 string             `protobuf:"bytes,8,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *Customer) GetAddresses() []*address.Address {
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

func (m *Customer) GetUuid() string {
	if m != nil {
		return m.Uuid
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
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
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

type Response struct {
	Created  bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Customer *Customer `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	// Added a pluralised consignment to our generic response message
	// array of these types
	Customers            []*Customer      `protobuf:"bytes,3,rep,name=customers,proto3" json:"customers,omitempty"`
	Account              *account.Account `protobuf:"bytes,4,opt,name=Account,proto3" json:"Account,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
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

func (m *Response) GetAccount() *account.Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto.RegisterType((*Customer)(nil), "moneytree.svc.customer.Customer")
	proto.RegisterType((*Request)(nil), "moneytree.svc.customer.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.customer.Response")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x25, 0x9b, 0xd0, 0xec, 0x4e, 0x93, 0x54, 0x32, 0x12, 0xb2, 0x82, 0x80, 0x28, 0x08, 0x29,
	0x17, 0x36, 0x10, 0x24, 0x5a, 0xa4, 0x0a, 0xa9, 0xcd, 0x01, 0x71, 0xe9, 0x61, 0x2b, 0xf5, 0xd0,
	0x4b, 0xb4, 0xf1, 0x0e, 0xc4, 0x52, 0xbc, 0x0e, 0xb6, 0xb7, 0xa8, 0x9f, 0xc2, 0xbf, 0xf1, 0x15,
	0x7c, 0x01, 0xea, 0xac, 0xbd, 0x2d, 0x22, 0x55, 0x39, 0xe4, 0xb4, 0xf6, 0x9b, 0x37, 0xcf, 0x6f,
	0xdf, 0xd8, 0x30, 0x10, 0x95, 0x75, 0x5a, 0xa1, 0x49, 0x37, 0x46, 0x3b, 0xcd, 0x9e, 0x2a, 0x5d,
	0xe2, 0xb5, 0x33, 0x88, 0xa9, 0xbd, 0x12, 0x69, 0xa8, 0x0e, 0xfb, 0xb9, 0x10, 0xba, 0x2a, 0x5d,
	0x4d, 0x1b, 0xf6, 0xf3, 0xa2, 0x30, 0x68, 0x6d, 0xbd, 0x1d, 0xff, 0x8c, 0x20, 0x9e, 0x7b, 0x2a,
	0x1b, 0x40, 0x24, 0x0b, 0xde, 0x1a, 0xb5, 0x26, 0x49, 0x16, 0xc9, 0x82, 0x3d, 0x07, 0xf8, 0x2a,
	0x8d, 0x75, 0x8b, 0x32, 0x57, 0xc8, 0x23, 0xc2, 0x13, 0x42, 0xce, 0x72, 0x85, 0xec, 0x19, 0x24,
	0xeb, 0x3c, 0x54, 0xdb, 0x54, 0x8d, 0x6f, 0x00, 0x2a, 0x1e, 0x43, 0xe2, 0x4f, 0x42, 0xcb, 0x3b,
	0xa3, 0xf6, 0x64, 0x7f, 0xf6, 0x22, 0xfd, 0xdb, 0x62, 0x70, 0x72, 0x52, 0x7f, 0xb3, 0xdb, 0x06,
	0xf6, 0x0a, 0xfa, 0x9b, 0x95, 0x2e, 0x71, 0x51, 0x56, 0x6a, 0x89, 0xc6, 0xf2, 0xc7, 0xa3, 0xf6,
	0x24, 0xc9, 0x7a, 0x04, 0x9e, 0xd5, 0xd8, 0x0d, 0x09, 0x55, 0x2e, 0xd7, 0x0b, 0xdf, 0xc7, 0xf7,
	0xc8, 0x43, 0x8f, 0x40, 0x2f, 0xca, 0x5e, 0xc3, 0xc0, 0x07, 0xe0, 0xb5, 0x78, 0x97, 0x58, 0x21,
	0x96, 0x5a, 0x8c, 0x31, 0xe8, 0x54, 0x95, 0x2c, 0x78, 0x4c, 0x45, 0x5a, 0x8f, 0xdf, 0x42, 0x37,
	0xc3, 0xef, 0x15, 0x5a, 0xb7, 0x45, 0xa5, 0xb5, 0x45, 0x65, 0xfc, 0xab, 0x05, 0x71, 0x86, 0x76,
	0xa3, 0x4b, 0x8b, 0x8c, 0x43, 0x57, 0x18, 0xcc, 0x1d, 0xd6, 0x91, 0xc6, 0x59, 0xd8, 0xb2, 0x63,
	0x88, 0xc3, 0x78, 0x28, 0xd5, 0xfd, 0xd9, 0x28, 0xdd, 0x3e, 0xbd, 0x34, 0xcc, 0x26, 0x6b, 0x3a,
	0xd8, 0x27, 0x48, 0xc2, 0xda, 0xf2, 0x36, 0x25, 0xfb, 0x70, 0xfb, 0x6d, 0x0b, 0x3b, 0x82, 0xee,
	0x49, 0xed, 0x9a, 0x77, 0xe8, 0xf0, 0x7f, 0xe6, 0xe2, 0x2f, 0x8c, 0x67, 0x65, 0x81, 0x3e, 0xfb,
	0x1d, 0xc1, 0x41, 0x50, 0x3c, 0x47, 0x73, 0x25, 0x05, 0xb2, 0x0b, 0x18, 0xcc, 0xe9, 0xb7, 0x9a,
	0x5b, 0xf4, 0xa0, 0x99, 0xe1, 0xbd, 0x8c, 0x90, 0xdd, 0xf8, 0x11, 0x3b, 0x87, 0xde, 0x67, 0x74,
	0xf3, 0xc6, 0xf5, 0xcb, 0xfb, 0x7b, 0x68, 0x44, 0xff, 0x25, 0x7a, 0x01, 0x07, 0x77, 0x44, 0x4f,
	0xaf, 0xbf, 0x14, 0xbb, 0xd1, 0xbd, 0x84, 0x27, 0x77, 0x74, 0x7d, 0x5c, 0xbb, 0xf1, 0x7c, 0x7a,
	0x74, 0xf9, 0xe1, 0x9b, 0x74, 0xab, 0x6a, 0x99, 0x0a, 0xad, 0xa6, 0x62, 0x65, 0xa4, 0xfd, 0x21,
	0xd7, 0x6b, 0x99, 0x2b, 0xfb, 0xee, 0xe3, 0xe1, 0xe1, 0xb4, 0x51, 0x78, 0x43, 0x2f, 0xda, 0x4e,
	0x83, 0xc8, 0x72, 0x8f, 0x80, 0xf7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x19, 0xda, 0xc9, 0x42,
	0x2a, 0x04, 0x00, 0x00,
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
	GetCustomerAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
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

func (c *customerServiceClient) GetCustomerAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "CustomerService.GetCustomerAccounts", in)
	out := new(Response)
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
	GetCustomerAccounts(context.Context, *Request, *Response) error
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

func (h *CustomerService) GetCustomerAccounts(ctx context.Context, in *Request, out *Response) error {
	return h.CustomerServiceHandler.GetCustomerAccounts(ctx, in, out)
}
