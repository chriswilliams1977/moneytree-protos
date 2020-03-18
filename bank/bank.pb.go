// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bank.proto

package bank

import (
	fmt "fmt"
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

type Bank struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RoutingNumber        int32            `protobuf:"varint,3,opt,name=routing_number,json=routingNumber,proto3" json:"routing_number,omitempty"`
	Address              *address.Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumbers         []string         `protobuf:"bytes,5,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	EmailAddress         string           `protobuf:"bytes,6,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	Uuid                 string           `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Bank) Reset()         { *m = Bank{} }
func (m *Bank) String() string { return proto.CompactTextString(m) }
func (*Bank) ProtoMessage()    {}
func (*Bank) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6371916d5cb63b4, []int{0}
}

func (m *Bank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bank.Unmarshal(m, b)
}
func (m *Bank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bank.Marshal(b, m, deterministic)
}
func (m *Bank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bank.Merge(m, src)
}
func (m *Bank) XXX_Size() int {
	return xxx_messageInfo_Bank.Size(m)
}
func (m *Bank) XXX_DiscardUnknown() {
	xxx_messageInfo_Bank.DiscardUnknown(m)
}

var xxx_messageInfo_Bank proto.InternalMessageInfo

func (m *Bank) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bank) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bank) GetRoutingNumber() int32 {
	if m != nil {
		return m.RoutingNumber
	}
	return 0
}

func (m *Bank) GetAddress() *address.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Bank) GetPhoneNumbers() []string {
	if m != nil {
		return m.PhoneNumbers
	}
	return nil
}

func (m *Bank) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *Bank) GetUuid() string {
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
	return fileDescriptor_a6371916d5cb63b4, []int{1}
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
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Bank                 *Bank    `protobuf:"bytes,2,opt,name=bank,proto3" json:"bank,omitempty"`
	Banks                []*Bank  `protobuf:"bytes,3,rep,name=banks,proto3" json:"banks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6371916d5cb63b4, []int{2}
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

func (m *Response) GetBank() *Bank {
	if m != nil {
		return m.Bank
	}
	return nil
}

func (m *Response) GetBanks() []*Bank {
	if m != nil {
		return m.Banks
	}
	return nil
}

func init() {
	proto.RegisterType((*Bank)(nil), "moneytree.svc.bank.Bank")
	proto.RegisterType((*Request)(nil), "moneytree.svc.bank.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.bank.Response")
}

func init() { proto.RegisterFile("bank.proto", fileDescriptor_a6371916d5cb63b4) }

var fileDescriptor_a6371916d5cb63b4 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x31, 0xcf, 0xd3, 0x30,
	0x10, 0x25, 0x4d, 0xda, 0xb4, 0x6e, 0xd3, 0xc1, 0x93, 0x55, 0x10, 0x8a, 0x82, 0x90, 0x32, 0x80,
	0x0b, 0x01, 0xa9, 0x30, 0xd2, 0x01, 0x36, 0x06, 0x23, 0x16, 0x96, 0xca, 0x49, 0x4e, 0xad, 0xd5,
	0xc6, 0x2e, 0xb6, 0x53, 0xc4, 0xca, 0xbf, 0xe5, 0x5f, 0x20, 0xdb, 0x09, 0x12, 0xa2, 0xfa, 0x26,
	0xfb, 0x9e, 0xdf, 0xdd, 0xbd, 0xf7, 0x12, 0x84, 0x6a, 0x2e, 0xcf, 0xf4, 0xaa, 0x95, 0x55, 0x18,
	0x77, 0x4a, 0xc2, 0x4f, 0xab, 0x01, 0xa8, 0xb9, 0x35, 0xd4, 0xbd, 0x6c, 0x32, 0xde, 0xb6, 0x1a,
	0x8c, 0x09, 0x94, 0xe2, 0x77, 0x84, 0x92, 0x3d, 0x97, 0x67, 0xbc, 0x46, 0x13, 0xd1, 0x92, 0x28,
	0x8f, 0xca, 0x05, 0x9b, 0x88, 0x16, 0x63, 0x94, 0x48, 0xde, 0x01, 0x99, 0x78, 0xc4, 0xdf, 0xf1,
	0x73, 0xb4, 0xd6, 0xaa, 0xb7, 0x42, 0x1e, 0x0f, 0xb2, 0xef, 0x6a, 0xd0, 0x24, 0xce, 0xa3, 0x72,
	0xca, 0xb2, 0x01, 0xfd, 0xec, 0x41, 0xfc, 0x0e, 0xa5, 0xc3, 0x12, 0x92, 0xe4, 0x51, 0xb9, 0xac,
	0x9e, 0xd2, 0x7f, 0x85, 0x8c, 0x12, 0x3e, 0x84, 0x93, 0x8d, 0x74, 0xfc, 0x0c, 0x65, 0xd7, 0x93,
	0x92, 0x30, 0x8c, 0x37, 0x64, 0x9a, 0xc7, 0xe5, 0x82, 0xad, 0x3c, 0x18, 0xa6, 0x7b, 0x12, 0x74,
	0x5c, 0x5c, 0x0e, 0xe3, 0x92, 0x99, 0x97, 0xb8, 0xf2, 0xe0, 0x30, 0xd2, 0xc9, 0xef, 0x7b, 0xd1,
	0x92, 0x34, 0xc8, 0x77, 0xf7, 0xe2, 0x15, 0x4a, 0x19, 0x7c, 0xef, 0xc1, 0x58, 0xe7, 0x84, 0x37,
	0x8d, 0xea, 0xa5, 0x1d, 0x9d, 0x04, 0xe7, 0xd9, 0x80, 0x86, 0x5d, 0xc5, 0xaf, 0x08, 0xcd, 0x19,
	0x98, 0xab, 0x92, 0x06, 0x30, 0x41, 0x69, 0xa3, 0x81, 0x5b, 0x08, 0x31, 0xcd, 0xd9, 0x58, 0xe2,
	0x17, 0x28, 0x71, 0xd9, 0xfa, 0xac, 0x96, 0x15, 0xa1, 0xff, 0xc7, 0x4e, 0x5d, 0xc6, 0xcc, 0xb3,
	0x30, 0x45, 0x53, 0x77, 0x1a, 0x12, 0xe7, 0xf1, 0x83, 0xf4, 0x40, 0xab, 0xbe, 0xa2, 0xa5, 0x2b,
	0xbf, 0x80, 0xbe, 0x89, 0x06, 0xf0, 0x47, 0x94, 0x7e, 0x02, 0xeb, 0xbf, 0xd9, 0xe3, 0x7b, 0xad,
	0x83, 0xc5, 0xcd, 0x93, 0xfb, 0x8f, 0xc1, 0x4c, 0xf1, 0x68, 0xff, 0xf6, 0x5b, 0x75, 0x14, 0xf6,
	0xd4, 0xd7, 0xb4, 0x51, 0xdd, 0xb6, 0x39, 0x69, 0x61, 0x7e, 0x88, 0xcb, 0x45, 0xf0, 0xce, 0xbc,
	0x7e, 0xbf, 0xdb, 0x6d, 0xff, 0x76, 0xbf, 0xf4, 0x7f, 0x8a, 0xd9, 0xba, 0x01, 0xf5, 0xcc, 0x17,
	0x6f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x00, 0xe8, 0xf4, 0x67, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for BankService service

type BankServiceClient interface {
	GetBank(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type bankServiceClient struct {
	c           client.Client
	serviceName string
}

func NewBankServiceClient(serviceName string, c client.Client) BankServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "moneytree.svc.bank"
	}
	return &bankServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *bankServiceClient) GetBank(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "BankService.GetBank", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BankService service

type BankServiceHandler interface {
	GetBank(context.Context, *Request, *Response) error
}

func RegisterBankServiceHandler(s server.Server, hdlr BankServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&BankService{hdlr}, opts...))
}

type BankService struct {
	BankServiceHandler
}

func (h *BankService) GetBank(ctx context.Context, in *Request, out *Response) error {
	return h.BankServiceHandler.GetBank(ctx, in, out)
}
