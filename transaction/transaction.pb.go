// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package transaction

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

type Transaction_TransactionType int32

const (
	Transaction_UNKNOWN_TRANSACTION_TYPE Transaction_TransactionType = 0
	Transaction_Withdrawal               Transaction_TransactionType = 1
	Transaction_Deposit                  Transaction_TransactionType = 2
	Transaction_Transfer                 Transaction_TransactionType = 3
)

var Transaction_TransactionType_name = map[int32]string{
	0: "UNKNOWN_TRANSACTION_TYPE",
	1: "Withdrawal",
	2: "Deposit",
	3: "Transfer",
}

var Transaction_TransactionType_value = map[string]int32{
	"UNKNOWN_TRANSACTION_TYPE": 0,
	"Withdrawal":               1,
	"Deposit":                  2,
	"Transfer":                 3,
}

func (x Transaction_TransactionType) String() string {
	return proto.EnumName(Transaction_TransactionType_name, int32(x))
}

func (Transaction_TransactionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0, 0}
}

type Transaction struct {
	Id                   string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Transaction_TransactionType `protobuf:"varint,2,opt,name=type,proto3,enum=moneytree.svc.transaction.Transaction_TransactionType" json:"type,omitempty"`
	Date                 *timestamp.Timestamp        `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Uuid                 string                      `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Transaction) GetType() Transaction_TransactionType {
	if m != nil {
		return m.Type
	}
	return Transaction_UNKNOWN_TRANSACTION_TYPE
}

func (m *Transaction) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Transaction) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Withdrawal struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Withdrawal) Reset()         { *m = Withdrawal{} }
func (m *Withdrawal) String() string { return proto.CompactTextString(m) }
func (*Withdrawal) ProtoMessage()    {}
func (*Withdrawal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{1}
}

func (m *Withdrawal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Withdrawal.Unmarshal(m, b)
}
func (m *Withdrawal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Withdrawal.Marshal(b, m, deterministic)
}
func (m *Withdrawal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Withdrawal.Merge(m, src)
}
func (m *Withdrawal) XXX_Size() int {
	return xxx_messageInfo_Withdrawal.Size(m)
}
func (m *Withdrawal) XXX_DiscardUnknown() {
	xxx_messageInfo_Withdrawal.DiscardUnknown(m)
}

var xxx_messageInfo_Withdrawal proto.InternalMessageInfo

func (m *Withdrawal) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Withdrawal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Withdrawal) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type Deposit struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deposit) Reset()         { *m = Deposit{} }
func (m *Deposit) String() string { return proto.CompactTextString(m) }
func (*Deposit) ProtoMessage()    {}
func (*Deposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{2}
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

func (m *Deposit) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Deposit) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Deposit) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

type Transfer struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Reference            string   `protobuf:"bytes,3,opt,name=reference,proto3" json:"reference,omitempty"`
	Payee                string   `protobuf:"bytes,4,opt,name=payee,proto3" json:"payee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{3}
}

func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transfer.Unmarshal(m, b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return xxx_messageInfo_Transfer.Size(m)
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Transfer) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Transfer) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Transfer) GetPayee() string {
	if m != nil {
		return m.Payee
	}
	return ""
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
	return fileDescriptor_2cc4e03d2c28c490, []int{4}
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
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Transaction          *Transaction   `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cc4e03d2c28c490, []int{5}
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

func (m *Response) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *Response) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterEnum("moneytree.svc.transaction.Transaction_TransactionType", Transaction_TransactionType_name, Transaction_TransactionType_value)
	proto.RegisterType((*Transaction)(nil), "moneytree.svc.transaction.Transaction")
	proto.RegisterType((*Withdrawal)(nil), "moneytree.svc.transaction.Withdrawal")
	proto.RegisterType((*Deposit)(nil), "moneytree.svc.transaction.Deposit")
	proto.RegisterType((*Transfer)(nil), "moneytree.svc.transaction.Transfer")
	proto.RegisterType((*Request)(nil), "moneytree.svc.transaction.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.transaction.Response")
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_2cc4e03d2c28c490) }

var fileDescriptor_2cc4e03d2c28c490 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xef, 0x6a, 0x13, 0x41,
	0x10, 0xef, 0x25, 0xb1, 0x49, 0xe6, 0x6a, 0x9a, 0x0e, 0x22, 0x67, 0x28, 0x18, 0x4e, 0x94, 0x7c,
	0xf1, 0xa2, 0x11, 0x2c, 0xc5, 0x4f, 0xb5, 0x8a, 0x5a, 0xe1, 0x2a, 0xdb, 0x93, 0x62, 0x11, 0xc2,
	0xe6, 0x6e, 0x92, 0x2c, 0xe4, 0xfe, 0xb8, 0xbb, 0xd7, 0x9a, 0x47, 0xf0, 0x95, 0x7c, 0x10, 0x9f,
	0x47, 0xba, 0x97, 0x34, 0xa7, 0x68, 0x49, 0x3f, 0xf4, 0xdb, 0xcd, 0xdc, 0xef, 0xdf, 0xce, 0xce,
	0xc2, 0x8e, 0x96, 0x3c, 0x51, 0x3c, 0xd4, 0x22, 0x4d, 0xbc, 0x4c, 0xa6, 0x3a, 0xc5, 0x07, 0x71,
	0x9a, 0xd0, 0x5c, 0x4b, 0x22, 0x4f, 0x9d, 0x87, 0x5e, 0x09, 0xd0, 0x79, 0x38, 0x49, 0xd3, 0xc9,
	0x8c, 0xfa, 0x06, 0x38, 0xca, 0xc7, 0x7d, 0x2d, 0x62, 0x52, 0x9a, 0xc7, 0x59, 0xc1, 0x75, 0x7f,
	0x54, 0xc0, 0x0e, 0x56, 0x04, 0x6c, 0x41, 0x45, 0x44, 0x8e, 0xd5, 0xb5, 0x7a, 0x4d, 0x56, 0x11,
	0x11, 0x1e, 0x41, 0x4d, 0xcf, 0x33, 0x72, 0x2a, 0x5d, 0xab, 0xd7, 0x1a, 0xbc, 0xf4, 0xfe, 0x6b,
	0xe5, 0x05, 0xff, 0xfe, 0x0e, 0xe6, 0x19, 0x31, 0xa3, 0x81, 0x1e, 0xd4, 0x22, 0xae, 0xc9, 0xa9,
	0x76, 0xad, 0x9e, 0x3d, 0xe8, 0x78, 0x45, 0x36, 0x6f, 0x99, 0xcd, 0x0b, 0x96, 0xd9, 0x98, 0xc1,
	0x21, 0x42, 0x2d, 0xcf, 0x45, 0xe4, 0xd4, 0x4c, 0x1a, 0xf3, 0xed, 0x9e, 0xc1, 0xf6, 0x5f, 0xe2,
	0xb8, 0x0b, 0xce, 0x67, 0xff, 0xa3, 0x7f, 0x7c, 0xea, 0x0f, 0x03, 0x76, 0xe0, 0x9f, 0x1c, 0x1c,
	0x06, 0x1f, 0x8e, 0xfd, 0x61, 0xf0, 0xe5, 0xd3, 0xdb, 0xf6, 0x06, 0xb6, 0x00, 0x4e, 0x85, 0x9e,
	0x46, 0x92, 0x5f, 0xf0, 0x59, 0xdb, 0x42, 0x1b, 0xea, 0x6f, 0x28, 0x4b, 0x95, 0xd0, 0xed, 0x0a,
	0x6e, 0x41, 0xc3, 0xa8, 0x8d, 0x49, 0xb6, 0xab, 0x6e, 0x54, 0x86, 0xe2, 0x7d, 0xd8, 0xe4, 0x71,
	0x9a, 0x27, 0xda, 0x4c, 0xc3, 0x62, 0x8b, 0x0a, 0xbb, 0x60, 0x47, 0xa4, 0x42, 0x29, 0xb2, 0xcb,
	0x04, 0x66, 0x30, 0x4d, 0x56, 0x6e, 0xe1, 0x2e, 0x34, 0x25, 0x8d, 0x49, 0x52, 0x12, 0x16, 0x87,
	0x6d, 0xb2, 0x55, 0xc3, 0xe5, 0x57, 0x01, 0x6e, 0xcd, 0xe2, 0xfb, 0xea, 0x58, 0xb7, 0xe5, 0x81,
	0xf7, 0xe0, 0x4e, 0xc6, 0xe7, 0x44, 0x8b, 0xdb, 0x29, 0x0a, 0xf7, 0x19, 0xd4, 0x19, 0x7d, 0xcb,
	0x49, 0x69, 0x7c, 0x0c, 0x2d, 0x1e, 0x86, 0x97, 0x5e, 0xc3, 0x24, 0x8f, 0x47, 0x24, 0x17, 0x5b,
	0x75, 0x77, 0xd1, 0xf5, 0x4d, 0xd3, 0xfd, 0x69, 0x41, 0x83, 0x91, 0xca, 0xd2, 0x44, 0x11, 0x3a,
	0x50, 0x0f, 0x25, 0x71, 0x4d, 0xc5, 0x0a, 0x36, 0xd8, 0xb2, 0xc4, 0xf7, 0x60, 0x97, 0x96, 0xcd,
	0xc4, 0xb5, 0x07, 0x4f, 0xd6, 0x5b, 0x47, 0x56, 0xa6, 0xe2, 0x11, 0x6c, 0x95, 0x4a, 0xe5, 0x54,
	0xbb, 0xd5, 0x1b, 0x48, 0xfd, 0xc1, 0x1d, 0xfc, 0xb2, 0x00, 0x4b, 0x7f, 0x4f, 0x48, 0x9e, 0x8b,
	0x90, 0x70, 0x04, 0x3b, 0x87, 0x26, 0x77, 0xf9, 0x65, 0xad, 0xe9, 0xd0, 0x79, 0x74, 0x0d, 0x6e,
	0x39, 0x28, 0x77, 0x03, 0xbf, 0xc2, 0xf6, 0x3b, 0xd2, 0x25, 0xa2, 0x42, 0xf7, 0x5a, 0xa6, 0xb9,
	0x95, 0x35, 0xd5, 0x5f, 0xbf, 0x3a, 0xdb, 0x9f, 0x08, 0x3d, 0xcd, 0x47, 0x5e, 0x98, 0xc6, 0xfd,
	0x70, 0x2a, 0x85, 0xba, 0x10, 0xb3, 0x99, 0xe0, 0xb1, 0x7a, 0xbe, 0xbf, 0xb7, 0xd7, 0xbf, 0x12,
	0x79, 0x6a, 0x5e, 0xaf, 0xea, 0x97, 0x74, 0x46, 0x9b, 0xa6, 0xf7, 0xe2, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0b, 0xe7, 0x07, 0xbc, 0xab, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TransactionService service

type TransactionServiceClient interface {
	CreateTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*Response, error)
	GetTransactions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type transactionServiceClient struct {
	c           client.Client
	serviceName string
}

func NewTransactionServiceClient(serviceName string, c client.Client) TransactionServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "moneytree.svc.transaction"
	}
	return &transactionServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *transactionServiceClient) CreateTransaction(ctx context.Context, in *Transaction, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "TransactionService.CreateTransaction", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) GetTransactions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "TransactionService.GetTransactions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TransactionService service

type TransactionServiceHandler interface {
	CreateTransaction(context.Context, *Transaction, *Response) error
	GetTransactions(context.Context, *Request, *Response) error
}

func RegisterTransactionServiceHandler(s server.Server, hdlr TransactionServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&TransactionService{hdlr}, opts...))
}

type TransactionService struct {
	TransactionServiceHandler
}

func (h *TransactionService) CreateTransaction(ctx context.Context, in *Transaction, out *Response) error {
	return h.TransactionServiceHandler.CreateTransaction(ctx, in, out)
}

func (h *TransactionService) GetTransactions(ctx context.Context, in *Request, out *Response) error {
	return h.TransactionServiceHandler.GetTransactions(ctx, in, out)
}
