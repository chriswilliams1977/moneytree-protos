// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deposit.proto

package deposit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Deposit struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount               int32                `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	TransactionId        int32                `protobuf:"varint,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d647de60f1ae88, []int{0}
}

func (m *Deposit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deposit.Unmarshal(m, b)
}
func (m *Deposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deposit.Marshal(b, m, deterministic)
}
func (m *Deposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deposit.Merge(m, src)
}
func (m *Deposit) XXX_Size() int {
	return xxx_messageInfo_Deposit.Size(m)
}
func (m *Deposit) XXX_DiscardUnknown() {
	xxx_messageInfo_Deposit.DiscardUnknown(m)
}

var xxx_messageInfo_Deposit proto.InternalMessageInfo

func (m *Deposit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deposit) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Deposit) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Deposit) GetTransactionId() int32 {
	if m != nil {
		return m.TransactionId
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
	return fileDescriptor_c2d647de60f1ae88, []int{1}
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
	Created              bool       `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Deposit              *Deposit   `protobuf:"bytes,2,opt,name=deposit,proto3" json:"deposit,omitempty"`
	Deposits             []*Deposit `protobuf:"bytes,3,rep,name=deposits,proto3" json:"deposits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2d647de60f1ae88, []int{2}
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

func (m *Response) GetDeposit() *Deposit {
	if m != nil {
		return m.Deposit
	}
	return nil
}

func (m *Response) GetDeposits() []*Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func init() {
	proto.RegisterType((*Deposit)(nil), "moneytree.svc.deposit.Deposit")
	proto.RegisterType((*Request)(nil), "moneytree.svc.deposit.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.deposit.Response")
}

func init() { proto.RegisterFile("deposit.proto", fileDescriptor_c2d647de60f1ae88) }

var fileDescriptor_c2d647de60f1ae88 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4a, 0xeb, 0x40,
	0x14, 0xbe, 0x69, 0x7b, 0x9b, 0xde, 0x29, 0xe9, 0x62, 0xe0, 0x4a, 0xe8, 0xc2, 0x86, 0x80, 0xd0,
	0x8d, 0x13, 0xad, 0x48, 0xd5, 0xa5, 0x0a, 0xe2, 0x46, 0x64, 0x74, 0xe5, 0xa6, 0x4c, 0x26, 0xc7,
	0x76, 0xa0, 0xc9, 0xc4, 0x99, 0x49, 0xc5, 0x37, 0xf0, 0x19, 0x7c, 0x10, 0x9f, 0x4f, 0x3a, 0x99,
	0x14, 0x17, 0x16, 0xbb, 0xcb, 0xf9, 0xe6, 0xfb, 0x39, 0xe7, 0x0b, 0x0a, 0x32, 0x28, 0xa5, 0x16,
	0x86, 0x94, 0x4a, 0x1a, 0x89, 0xff, 0xe7, 0xb2, 0x80, 0x37, 0xa3, 0x00, 0x88, 0x5e, 0x71, 0xe2,
	0x1e, 0x87, 0xa3, 0xb9, 0x94, 0xf3, 0x25, 0x24, 0x96, 0x94, 0x56, 0xcf, 0x89, 0x11, 0x39, 0x68,
	0xc3, 0xf2, 0xb2, 0xd6, 0xc5, 0xef, 0x1e, 0xf2, 0xaf, 0x6b, 0x32, 0x1e, 0xa0, 0x96, 0xc8, 0x42,
	0x2f, 0xf2, 0xc6, 0xff, 0x68, 0x4b, 0x64, 0x78, 0x0f, 0x75, 0x59, 0x2e, 0xab, 0xc2, 0x84, 0xad,
	0xc8, 0x1b, 0xff, 0xa5, 0x6e, 0xc2, 0x04, 0x75, 0x32, 0x66, 0x20, 0x6c, 0x47, 0xde, 0xb8, 0x3f,
	0x19, 0x92, 0x3a, 0x83, 0x34, 0x19, 0xe4, 0xb1, 0xc9, 0xa0, 0x96, 0x87, 0x0f, 0xd0, 0xc0, 0x28,
	0x56, 0x68, 0xc6, 0x8d, 0x90, 0xc5, 0x4c, 0x64, 0x61, 0xc7, 0xfa, 0x05, 0xdf, 0xd0, 0xdb, 0x2c,
	0x3e, 0x42, 0x3e, 0x85, 0x97, 0x0a, 0xb4, 0x59, 0x2b, 0x18, 0xe7, 0xeb, 0xb0, 0x59, 0x51, 0xe5,
	0x29, 0x28, 0xb7, 0x55, 0xe0, 0xd0, 0x3b, 0x0b, 0xc6, 0x1f, 0x1e, 0xea, 0x51, 0xd0, 0xa5, 0x2c,
	0x34, 0xe0, 0x10, 0xf9, 0x5c, 0x01, 0x33, 0x50, 0x9f, 0xd0, 0xa3, 0xcd, 0x88, 0xcf, 0x90, 0xef,
	0xfa, 0xb0, 0x87, 0xf4, 0x27, 0xfb, 0xe4, 0xc7, 0xb6, 0x88, 0x2b, 0x82, 0x36, 0x74, 0x7c, 0x81,
	0x7a, 0xee, 0x53, 0x87, 0xed, 0xa8, 0xbd, 0x83, 0x74, 0xc3, 0x9f, 0x7c, 0x7a, 0x68, 0xe0, 0xd0,
	0x07, 0x50, 0x2b, 0xc1, 0x01, 0x53, 0x14, 0x5c, 0xd9, 0x9d, 0x9a, 0xc6, 0x7f, 0x71, 0x1b, 0x8e,
	0xb6, 0xbc, 0x37, 0x47, 0xc7, 0x7f, 0xf0, 0x3d, 0xea, 0xdf, 0x80, 0x71, 0x02, 0xbd, 0xd5, 0xd1,
	0x35, 0xbb, 0x83, 0xe3, 0xe5, 0xf4, 0xe9, 0x74, 0x2e, 0xcc, 0xa2, 0x4a, 0x09, 0x97, 0x79, 0xc2,
	0x17, 0x4a, 0xe8, 0x57, 0xb1, 0x5c, 0x0a, 0x96, 0xeb, 0xe3, 0xf3, 0xe9, 0x34, 0xd9, 0x18, 0x1c,
	0xda, 0x3f, 0xae, 0x13, 0xe7, 0x91, 0x76, 0xed, 0x7c, 0xf2, 0x15, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x7b, 0xd4, 0xa7, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DepositService service

type DepositServiceClient interface {
	CreateDeposit(ctx context.Context, in *Deposit, opts ...client.CallOption) (*Response, error)
	GetDeposits(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type depositServiceClient struct {
	c           client.Client
	serviceName string
}

func NewDepositServiceClient(serviceName string, c client.Client) DepositServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "moneytree.svc.deposit"
	}
	return &depositServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *depositServiceClient) CreateDeposit(ctx context.Context, in *Deposit, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "DepositService.CreateDeposit", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *depositServiceClient) GetDeposits(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "DepositService.GetDeposits", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DepositService service

type DepositServiceHandler interface {
	CreateDeposit(context.Context, *Deposit, *Response) error
	GetDeposits(context.Context, *Request, *Response) error
}

func RegisterDepositServiceHandler(s server.Server, hdlr DepositServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DepositService{hdlr}, opts...))
}

type DepositService struct {
	DepositServiceHandler
}

func (h *DepositService) CreateDeposit(ctx context.Context, in *Deposit, out *Response) error {
	return h.DepositServiceHandler.CreateDeposit(ctx, in, out)
}

func (h *DepositService) GetDeposits(ctx context.Context, in *Request, out *Response) error {
	return h.DepositServiceHandler.GetDeposits(ctx, in, out)
}
