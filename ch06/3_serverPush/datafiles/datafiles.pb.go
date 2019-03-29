// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datafiles.proto

package datafiles

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TransactionRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount               float32  `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionRequest) Reset()         { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()    {}
func (*TransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef83bb5dfcecfcd1, []int{0}
}

func (m *TransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionRequest.Unmarshal(m, b)
}
func (m *TransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
}
func (m *TransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *TransactionRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionRequest.Size(m)
}
func (m *TransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo

func (m *TransactionRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransactionRequest) GetAmount() float32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransactionResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Step                 int32    `protobuf:"varint,2,opt,name=step,proto3" json:"step,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef83bb5dfcecfcd1, []int{1}
}

func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResponse.Unmarshal(m, b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionResponse.Size(m)
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func (m *TransactionResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TransactionResponse) GetStep() int32 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *TransactionResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionRequest)(nil), "datafiles.TransactionRequest")
	proto.RegisterType((*TransactionResponse)(nil), "datafiles.TransactionResponse")
}

func init() { proto.RegisterFile("datafiles.proto", fileDescriptor_ef83bb5dfcecfcd1) }

var fileDescriptor_ef83bb5dfcecfcd1 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x4d, 0xd4, 0x85, 0x8e, 0xe0, 0xca, 0x08, 0x52, 0x04, 0xa5, 0xf4, 0xb4, 0xa7, 0x45,
	0xf4, 0x39, 0x16, 0x24, 0xf4, 0x05, 0x62, 0x3b, 0xc5, 0xa2, 0xcd, 0xc4, 0xcc, 0xf4, 0xe0, 0xdb,
	0x4b, 0x63, 0xd1, 0x8a, 0xec, 0x6d, 0xe6, 0x4b, 0xf8, 0xfe, 0xfc, 0x81, 0x6d, 0xe7, 0xd5, 0xf7,
	0xc3, 0x3b, 0xc9, 0x3e, 0x26, 0x56, 0xc6, 0xe2, 0x07, 0xd4, 0xcf, 0x80, 0x4d, 0xf2, 0x41, 0x7c,
	0xab, 0x03, 0x07, 0x47, 0x1f, 0x13, 0x89, 0x22, 0xc2, 0x59, 0x9f, 0x78, 0x2c, 0x4d, 0x65, 0x76,
	0x85, 0xcb, 0x33, 0x5e, 0x82, 0x55, 0x2e, 0x6d, 0x26, 0x56, 0x19, 0x6f, 0x60, 0xe3, 0x47, 0x9e,
	0x82, 0x96, 0xa7, 0x95, 0xd9, 0x59, 0xb7, 0x6c, 0x75, 0x0b, 0xd7, 0x7f, 0x8c, 0x12, 0x39, 0x08,
	0xcd, 0xd7, 0x45, 0xbd, 0x4e, 0xb2, 0x48, 0x97, 0x6d, 0x8e, 0x12, 0xa5, 0x98, 0xc5, 0xe7, 0x2e,
	0xcf, 0x58, 0xc1, 0x45, 0x47, 0xd2, 0xa6, 0x21, 0xce, 0x8a, 0xec, 0x2f, 0xdc, 0x1a, 0x3d, 0xbe,
	0xc2, 0xd5, 0x81, 0x03, 0x7d, 0xae, 0x92, 0xb0, 0x81, 0xed, 0xc1, 0xbf, 0xd1, 0x1a, 0xdd, 0xed,
	0x7f, 0xab, 0xff, 0xaf, 0x79, 0x7b, 0x7f, 0xec, 0xf8, 0xfb, 0xcd, 0xf5, 0xc9, 0x83, 0x79, 0xd9,
	0xe4, 0x2f, 0x7b, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x02, 0x6a, 0x0e, 0xd3, 0x45, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MoneyTransactionClient is the client API for MoneyTransaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MoneyTransactionClient interface {
	MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (MoneyTransaction_MakeTransactionClient, error)
}

type moneyTransactionClient struct {
	cc *grpc.ClientConn
}

func NewMoneyTransactionClient(cc *grpc.ClientConn) MoneyTransactionClient {
	return &moneyTransactionClient{cc}
}

func (c *moneyTransactionClient) MakeTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (MoneyTransaction_MakeTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MoneyTransaction_serviceDesc.Streams[0], "/datafiles.MoneyTransaction/MakeTransaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &moneyTransactionMakeTransactionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MoneyTransaction_MakeTransactionClient interface {
	Recv() (*TransactionResponse, error)
	grpc.ClientStream
}

type moneyTransactionMakeTransactionClient struct {
	grpc.ClientStream
}

func (x *moneyTransactionMakeTransactionClient) Recv() (*TransactionResponse, error) {
	m := new(TransactionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MoneyTransactionServer is the server API for MoneyTransaction service.
type MoneyTransactionServer interface {
	MakeTransaction(*TransactionRequest, MoneyTransaction_MakeTransactionServer) error
}

// UnimplementedMoneyTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}

func (*UnimplementedMoneyTransactionServer) MakeTransaction(req *TransactionRequest, srv MoneyTransaction_MakeTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method MakeTransaction not implemented")
}

func RegisterMoneyTransactionServer(s *grpc.Server, srv MoneyTransactionServer) {
	s.RegisterService(&_MoneyTransaction_serviceDesc, srv)
}

func _MoneyTransaction_MakeTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransactionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MoneyTransactionServer).MakeTransaction(m, &moneyTransactionMakeTransactionServer{stream})
}

type MoneyTransaction_MakeTransactionServer interface {
	Send(*TransactionResponse) error
	grpc.ServerStream
}

type moneyTransactionMakeTransactionServer struct {
	grpc.ServerStream
}

func (x *moneyTransactionMakeTransactionServer) Send(m *TransactionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MoneyTransaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datafiles.MoneyTransaction",
	HandlerType: (*MoneyTransactionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MakeTransaction",
			Handler:       _MoneyTransaction_MakeTransaction_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "datafiles.proto",
}