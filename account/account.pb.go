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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number               string   `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
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

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

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
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
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
	proto.RegisterType((*Response)(nil), "moneytree.svc.account.Response")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x2e, 0xea, 0xcf, 0x55, 0xcd, 0x70, 0x12, 0xc8, 0xea, 0x00, 0x51, 0x26, 0x26, 0x0f,
	0x65, 0x41, 0x6c, 0x50, 0x21, 0x76, 0xb3, 0xb1, 0xb5, 0xee, 0x0d, 0x19, 0x62, 0x47, 0xb6, 0x53,
	0xa9, 0x1b, 0x8f, 0xc5, 0xe3, 0x21, 0xb9, 0xe7, 0x0e, 0x08, 0xd4, 0x6e, 0x5f, 0xbe, 0xbf, 0xcb,
	0x9d, 0x61, 0xb1, 0xb1, 0xd6, 0x0f, 0x2e, 0xe9, 0x3e, 0xf8, 0xe4, 0xf1, 0xa6, 0xf3, 0x8e, 0x0e,
	0x29, 0x10, 0xe9, 0xb8, 0xb7, 0x9a, 0xc5, 0x65, 0x65, 0x87, 0x98, 0x7c, 0x47, 0xe1, 0x68, 0x6b,
	0xde, 0x60, 0xf2, 0x72, 0x94, 0xb0, 0x02, 0xd9, 0xee, 0x94, 0xa8, 0xc5, 0xc3, 0xcc, 0xc8, 0x76,
	0x87, 0xb7, 0x30, 0x76, 0x43, 0xb7, 0xa5, 0xa0, 0x64, 0xe6, 0xf8, 0x0b, 0x11, 0xae, 0xd3, 0xa1,
	0x27, 0x35, 0xca, 0x6c, 0xc6, 0xcd, 0x97, 0x80, 0xa9, 0xa1, 0xd8, 0x7b, 0x17, 0x09, 0x9f, 0x60,
	0xc2, 0xe3, 0x72, 0xdb, 0x7c, 0x75, 0xa7, 0xff, 0xfc, 0x19, 0xcd, 0x93, 0x4d, 0xb1, 0xe3, 0x33,
	0x4c, 0x19, 0x46, 0x25, 0xeb, 0xd1, 0x05, 0xd1, 0x93, 0x7f, 0xf5, 0x2d, 0xa0, 0x62, 0xf6, 0x83,
	0xc2, 0xbe, 0xb5, 0x84, 0x06, 0x16, 0xeb, 0x40, 0x9b, 0x44, 0x65, 0xc5, 0x33, 0x6d, 0xcb, 0xfb,
	0x7f, 0xf4, 0xb2, 0x5a, 0x73, 0x85, 0x06, 0xe6, 0xef, 0x94, 0x38, 0x10, 0xb1, 0xfe, 0x95, 0x38,
	0x9d, 0x77, 0xcd, 0xe0, 0x82, 0xce, 0xd7, 0xd9, 0x67, 0xb9, 0xc0, 0x76, 0x9c, 0x9f, 0xe5, 0xf1,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x39, 0x4b, 0x5f, 0x22, 0xce, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountService service

type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Response, error)
	GetAccounts(ctx context.Context, in *Customer, opts ...client.CallOption) (*Response, error)
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

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *Customer, opts ...client.CallOption) (*Response, error) {
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
	GetAccounts(context.Context, *Customer, *Response) error
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

func (h *AccountService) GetAccounts(ctx context.Context, in *Customer, out *Response) error {
	return h.AccountServiceHandler.GetAccounts(ctx, in, out)
}
