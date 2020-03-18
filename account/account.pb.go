// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account

import (
	fmt "fmt"
	transaction "github.com/chriswilliams1977/moneytree-protos/transaction"
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

type Account_AccountType int32

const (
	Account_UNKNOWN_ACCOUNT_TYPE Account_AccountType = 0
	Account_CurrentAccount       Account_AccountType = 1
	Account_SavingAccount        Account_AccountType = 2
	Account_MortgageAccount      Account_AccountType = 3
)

var Account_AccountType_name = map[int32]string{
	0: "UNKNOWN_ACCOUNT_TYPE",
	1: "CurrentAccount",
	2: "SavingAccount",
	3: "MortgageAccount",
}

var Account_AccountType_value = map[string]int32{
	"UNKNOWN_ACCOUNT_TYPE": 0,
	"CurrentAccount":       1,
	"SavingAccount":        2,
	"MortgageAccount":      3,
}

func (x Account_AccountType) String() string {
	return proto.EnumName(Account_AccountType_name, int32(x))
}

func (Account_AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0, 0}
}

type Account struct {
	Number               string              `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Account_AccountType `protobuf:"varint,2,opt,name=type,proto3,enum=moneytree.svc.account.Account_AccountType" json:"type,omitempty"`
	Status               string              `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Balance              float64             `protobuf:"fixed64,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Uuid                 int64               `protobuf:"varint,5,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
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

func (m *Account) GetType() Account_AccountType {
	if m != nil {
		return m.Type
	}
	return Account_UNKNOWN_ACCOUNT_TYPE
}

func (m *Account) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Account) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetUuid() int64 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

type CurrentAccount struct {
	InterestRate         float64  `protobuf:"fixed64,1,opt,name=interest_rate,json=interestRate,proto3" json:"interest_rate,omitempty"`
	MonthlyCharge        int32    `protobuf:"varint,2,opt,name=monthly_charge,json=monthlyCharge,proto3" json:"monthly_charge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentAccount) Reset()         { *m = CurrentAccount{} }
func (m *CurrentAccount) String() string { return proto.CompactTextString(m) }
func (*CurrentAccount) ProtoMessage()    {}
func (*CurrentAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *CurrentAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentAccount.Unmarshal(m, b)
}
func (m *CurrentAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentAccount.Marshal(b, m, deterministic)
}
func (m *CurrentAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentAccount.Merge(m, src)
}
func (m *CurrentAccount) XXX_Size() int {
	return xxx_messageInfo_CurrentAccount.Size(m)
}
func (m *CurrentAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentAccount.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentAccount proto.InternalMessageInfo

func (m *CurrentAccount) GetInterestRate() float64 {
	if m != nil {
		return m.InterestRate
	}
	return 0
}

func (m *CurrentAccount) GetMonthlyCharge() int32 {
	if m != nil {
		return m.MonthlyCharge
	}
	return 0
}

type SavingsAccount struct {
	InterestRate         float64  `protobuf:"fixed64,1,opt,name=interest_rate,json=interestRate,proto3" json:"interest_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SavingsAccount) Reset()         { *m = SavingsAccount{} }
func (m *SavingsAccount) String() string { return proto.CompactTextString(m) }
func (*SavingsAccount) ProtoMessage()    {}
func (*SavingsAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *SavingsAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavingsAccount.Unmarshal(m, b)
}
func (m *SavingsAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavingsAccount.Marshal(b, m, deterministic)
}
func (m *SavingsAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavingsAccount.Merge(m, src)
}
func (m *SavingsAccount) XXX_Size() int {
	return xxx_messageInfo_SavingsAccount.Size(m)
}
func (m *SavingsAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_SavingsAccount.DiscardUnknown(m)
}

var xxx_messageInfo_SavingsAccount proto.InternalMessageInfo

func (m *SavingsAccount) GetInterestRate() float64 {
	if m != nil {
		return m.InterestRate
	}
	return 0
}

type MortgageAccount struct {
	InterestRate         float64  `protobuf:"fixed64,1,opt,name=interest_rate,json=interestRate,proto3" json:"interest_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MortgageAccount) Reset()         { *m = MortgageAccount{} }
func (m *MortgageAccount) String() string { return proto.CompactTextString(m) }
func (*MortgageAccount) ProtoMessage()    {}
func (*MortgageAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *MortgageAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MortgageAccount.Unmarshal(m, b)
}
func (m *MortgageAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MortgageAccount.Marshal(b, m, deterministic)
}
func (m *MortgageAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MortgageAccount.Merge(m, src)
}
func (m *MortgageAccount) XXX_Size() int {
	return xxx_messageInfo_MortgageAccount.Size(m)
}
func (m *MortgageAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MortgageAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MortgageAccount proto.InternalMessageInfo

func (m *MortgageAccount) GetInterestRate() float64 {
	if m != nil {
		return m.InterestRate
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
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
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
	Account              *Account                   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Accounts             []*Account                 `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Transactions         []*transaction.Transaction `protobuf:"bytes,3,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
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

func (m *Response) GetTransactions() []*transaction.Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterEnum("moneytree.svc.account.Account_AccountType", Account_AccountType_name, Account_AccountType_value)
	proto.RegisterType((*Account)(nil), "moneytree.svc.account.Account")
	proto.RegisterType((*CurrentAccount)(nil), "moneytree.svc.account.CurrentAccount")
	proto.RegisterType((*SavingsAccount)(nil), "moneytree.svc.account.SavingsAccount")
	proto.RegisterType((*MortgageAccount)(nil), "moneytree.svc.account.MortgageAccount")
	proto.RegisterType((*Request)(nil), "moneytree.svc.account.Request")
	proto.RegisterType((*Response)(nil), "moneytree.svc.account.Response")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x6d, 0x6f, 0xd2, 0x50,
	0x14, 0xc7, 0x57, 0x60, 0x63, 0x1e, 0x56, 0xb6, 0x5d, 0x1f, 0xd2, 0xec, 0x85, 0x92, 0x1a, 0x0d,
	0x31, 0xb1, 0x28, 0x66, 0xa2, 0xbe, 0x30, 0x01, 0x62, 0x96, 0x68, 0x64, 0xcb, 0x1d, 0x64, 0xd1,
	0x98, 0x90, 0x4b, 0x77, 0x02, 0x4d, 0xe0, 0x16, 0xef, 0x3d, 0xc5, 0xf4, 0x3b, 0xf8, 0x85, 0xfc,
	0x0e, 0x7e, 0x28, 0x43, 0x7b, 0xcb, 0x60, 0x3a, 0xc5, 0x84, 0x57, 0x3d, 0x4f, 0xff, 0x5f, 0x4f,
	0xce, 0x39, 0x2d, 0xd8, 0xc2, 0xf7, 0xc3, 0x48, 0x92, 0x37, 0x55, 0x21, 0x85, 0xec, 0xee, 0x24,
	0x94, 0x18, 0x93, 0x42, 0xf4, 0xf4, 0xcc, 0xf7, 0x4c, 0xf2, 0xe8, 0x90, 0x94, 0x90, 0x5a, 0xf8,
	0x14, 0x84, 0x32, 0xad, 0x74, 0xbf, 0xe7, 0xa0, 0xd8, 0x4c, 0xd3, 0xec, 0x1e, 0xec, 0xc8, 0x68,
	0x32, 0x40, 0xe5, 0x58, 0x15, 0xab, 0x7a, 0x8b, 0x1b, 0x8f, 0xbd, 0x85, 0x02, 0xc5, 0x53, 0x74,
	0x72, 0x15, 0xab, 0x5a, 0xae, 0x3f, 0xf1, 0xfe, 0x08, 0xf7, 0x9a, 0xab, 0xcf, 0x6e, 0x3c, 0x45,
	0x9e, 0xe8, 0xe6, 0x5c, 0x4d, 0x82, 0x22, 0xed, 0xe4, 0x53, 0x6e, 0xea, 0x31, 0x07, 0x8a, 0x03,
	0x31, 0x16, 0xd2, 0x47, 0xa7, 0x50, 0xb1, 0xaa, 0x16, 0xcf, 0x5c, 0xc6, 0xa0, 0x10, 0x45, 0xc1,
	0xa5, 0xb3, 0x5d, 0xb1, 0xaa, 0x79, 0x9e, 0xd8, 0xae, 0x0f, 0xa5, 0x25, 0x34, 0x73, 0xe0, 0x4e,
	0xaf, 0xf3, 0xa1, 0x73, 0x7a, 0xd1, 0xe9, 0x37, 0xdb, 0xed, 0xd3, 0x5e, 0xa7, 0xdb, 0xef, 0x7e,
	0x3a, 0x7b, 0x77, 0xb0, 0xc5, 0x18, 0x94, 0xdb, 0x91, 0x52, 0x28, 0xc9, 0xd4, 0x1f, 0x58, 0xec,
	0x10, 0xec, 0x73, 0x31, 0x0b, 0xe4, 0x30, 0x0b, 0xe5, 0xd8, 0x6d, 0xd8, 0xff, 0x18, 0x2a, 0x1a,
	0x8a, 0x21, 0x66, 0xc1, 0xbc, 0xfb, 0xe5, 0xba, 0x96, 0x3d, 0x04, 0x3b, 0x90, 0x84, 0x0a, 0x35,
	0xf5, 0x95, 0x20, 0x4c, 0x66, 0x63, 0xf1, 0xbd, 0x2c, 0xc8, 0x05, 0x21, 0x7b, 0x04, 0xe5, 0x49,
	0x28, 0x69, 0x34, 0x8e, 0xfb, 0xfe, 0x48, 0xa8, 0x61, 0x3a, 0xab, 0x6d, 0x6e, 0x9b, 0x68, 0x3b,
	0x09, 0xba, 0xc7, 0x50, 0x4e, 0xbb, 0xd0, 0xff, 0x43, 0x77, 0x5f, 0xfe, 0xd6, 0xe9, 0x7a, 0xba,
	0x67, 0x50, 0xe4, 0xf8, 0x35, 0x42, 0x4d, 0xf3, 0x06, 0xcd, 0x9e, 0xfa, 0x2b, 0x2b, 0xce, 0xee,
	0xa6, 0x93, 0x04, 0xdd, 0x9f, 0x16, 0xec, 0x72, 0xd4, 0xd3, 0x50, 0x6a, 0x64, 0xaf, 0xa0, 0x68,
	0xb2, 0x49, 0x71, 0xa9, 0x7e, 0xff, 0xef, 0x9b, 0xe7, 0x59, 0x39, 0x7b, 0x03, 0xbb, 0xc6, 0xd4,
	0x4e, 0xae, 0x92, 0x5f, 0x43, 0xba, 0xa8, 0x67, 0xef, 0x61, 0x6f, 0xe9, 0x4a, 0xe7, 0x27, 0x33,
	0xd7, 0x3f, 0xbe, 0xa6, 0x5f, 0x3e, 0xe4, 0xee, 0x95, 0xcd, 0x57, 0xb4, 0xf5, 0x1f, 0x79, 0x28,
	0x9b, 0x37, 0x9c, 0xa3, 0x9a, 0x05, 0x3e, 0x32, 0x0e, 0x76, 0x5b, 0xa1, 0xa0, 0xc5, 0x24, 0xff,
	0xd1, 0xd9, 0xd1, 0x83, 0x1b, 0xf2, 0xd9, 0x98, 0xdc, 0x2d, 0x76, 0x06, 0xa5, 0x13, 0xcc, 0x0e,
	0x46, 0xdf, 0x48, 0x34, 0xbb, 0x58, 0x87, 0x78, 0x01, 0xec, 0x8a, 0xd8, 0x8a, 0xd3, 0xed, 0x6c,
	0x02, 0xcc, 0xc1, 0xee, 0x4d, 0x2f, 0x05, 0x61, 0xcb, 0x7c, 0x69, 0x1b, 0x60, 0x76, 0x61, 0xff,
	0x04, 0x69, 0x69, 0x0b, 0x9b, 0x18, 0x41, 0xab, 0xf1, 0xf9, 0x78, 0x18, 0xd0, 0x28, 0x1a, 0x78,
	0x7e, 0x38, 0xa9, 0xf9, 0x23, 0x15, 0xe8, 0x6f, 0xc1, 0x78, 0x1c, 0x88, 0x89, 0x7e, 0xfe, 0xba,
	0xd1, 0xa8, 0x2d, 0x00, 0x4f, 0x93, 0x1f, 0x99, 0xae, 0x19, 0xc6, 0x60, 0x27, 0xf1, 0x5f, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x71, 0xdf, 0x53, 0x1f, 0x13, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AccountService service

type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Response, error)
	GetAccounts(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetAccountByNumber(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateBalance(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetTransactions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
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

func (c *accountServiceClient) GetAccountByNumber(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetAccountByNumber", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateBalance(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.UpdateBalance", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetTransactions(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "AccountService.GetTransactions", in)
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
	GetAccountByNumber(context.Context, *Request, *Response) error
	UpdateBalance(context.Context, *Request, *Response) error
	GetTransactions(context.Context, *Request, *Response) error
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

func (h *AccountService) GetAccountByNumber(ctx context.Context, in *Request, out *Response) error {
	return h.AccountServiceHandler.GetAccountByNumber(ctx, in, out)
}

func (h *AccountService) UpdateBalance(ctx context.Context, in *Request, out *Response) error {
	return h.AccountServiceHandler.UpdateBalance(ctx, in, out)
}

func (h *AccountService) GetTransactions(ctx context.Context, in *Request, out *Response) error {
	return h.AccountServiceHandler.GetTransactions(ctx, in, out)
}
