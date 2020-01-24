// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

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

type Account struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Balance              int32    `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
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

func (m *Account) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Account) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Account) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Account) GetBalance() int32 {
	if m != nil {
		return m.Balance
	}
	return 0
}

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
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
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
	Account              *Account   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Accounts             []*Account `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
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

func (m *Response) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Response) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "moneytree.svc.account.Account")
	proto.RegisterType((*Request)(nil), "moneytree.svc.account.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.account.Response")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x35, 0x6d, 0xed, 0xb6, 0x53, 0xda, 0xc3, 0x80, 0x12, 0x3c, 0xe8, 0xb2, 0x20, 0xf4, 0x14,
	0xa4, 0x5e, 0xc4, 0x9b, 0x7a, 0xf0, 0x26, 0x12, 0x6f, 0x5e, 0x24, 0x1b, 0x06, 0x11, 0x6c, 0xb2,
	0x26, 0xd9, 0x42, 0x6f, 0x7e, 0x95, 0xdf, 0x27, 0x66, 0x13, 0x41, 0xb0, 0x74, 0x6f, 0x33, 0xef,
	0xbd, 0x79, 0x93, 0x37, 0x81, 0xb9, 0xd2, 0xda, 0xb6, 0x26, 0x88, 0xc6, 0xd9, 0x60, 0xf1, 0x68,
	0x6d, 0x0d, 0x6d, 0x83, 0x23, 0x12, 0x7e, 0xa3, 0x45, 0x22, 0xab, 0x57, 0x28, 0x6e, 0xba, 0x12,
	0x8f, 0x61, 0x6c, 0xda, 0x75, 0x4d, 0x8e, 0xb3, 0x92, 0x2d, 0xa7, 0x32, 0x75, 0x88, 0x30, 0x0a,
	0xdb, 0x86, 0xf8, 0x20, 0xa2, 0xb1, 0xfe, 0xd1, 0xfa, 0xa0, 0x42, 0xeb, 0xf9, 0xb0, 0xd3, 0x76,
	0x1d, 0x72, 0x28, 0x6a, 0xf5, 0xae, 0x8c, 0x26, 0x3e, 0x2a, 0xd9, 0xf2, 0x50, 0xe6, 0xb6, 0xba,
	0x80, 0x42, 0xd2, 0x47, 0x4b, 0x3e, 0xe0, 0x39, 0x2c, 0xd2, 0xfa, 0x97, 0x3f, 0x0b, 0xf3, 0x8b,
	0x1f, 0x22, 0x58, 0x7d, 0x32, 0x98, 0x48, 0xf2, 0x8d, 0x35, 0x9e, 0xf0, 0x0a, 0x8a, 0xc4, 0x46,
	0xf1, 0x6c, 0x75, 0x2a, 0xfe, 0x0d, 0x24, 0x52, 0x1a, 0x99, 0xe5, 0x78, 0x0d, 0x93, 0x54, 0x7a,
	0x3e, 0x28, 0x87, 0x3d, 0x46, 0x7f, 0xf5, 0xab, 0x2f, 0x06, 0x8b, 0x84, 0x3e, 0x91, 0xdb, 0xbc,
	0x69, 0x42, 0x09, 0xf3, 0x3b, 0x47, 0x2a, 0x50, 0x3e, 0xdb, 0x1e, 0xb7, 0x93, 0xb3, 0x1d, 0x7c,
	0x8e, 0x56, 0x1d, 0xe0, 0x23, 0xcc, 0xee, 0x29, 0xa4, 0x01, 0xbf, 0xd3, 0x31, 0xdd, 0xaf, 0x87,
	0xe3, 0xed, 0xf4, 0x39, 0xe7, 0xaf, 0xc7, 0xf1, 0xff, 0x2f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xed, 0xe6, 0xfb, 0xe2, 0x10, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountService service

type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Response, error)
	GetAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type accountServiceClient struct {
	c           client.Client
	serviceName string
}

func NewAccountServiceClient(serviceName string, c client.Client) AccountServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "moneytree.svc.account"
	}
	return &accountServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.CreateAccount", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccounts", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceHandler interface {
	CreateAccount(context.Context, *Account, *Response) error
	GetAccounts(context.Context, *Request, *Response) error
}

func RegisterAccountServiceHandler(s server.Server, hdlr AccountServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AccountService{hdlr}, opts...))
}

type AccountService struct {
	AccountServiceHandler
}

func (h *AccountService) CreateAccount(ctx context.Context, in *Account, out *Response) error {
	return h.AccountServiceHandler.CreateAccount(ctx, in, out)
}

func (h *AccountService) GetAccounts(ctx context.Context, in *Request, out *Response) error {
	return h.AccountServiceHandler.GetAccounts(ctx, in, out)
}
