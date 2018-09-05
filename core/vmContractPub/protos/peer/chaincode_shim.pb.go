// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/chaincode_shim.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED           ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER            ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED          ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                ChaincodeMessage_Type = 3
	ChaincodeMessage_READY               ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION         ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED           ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR               ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE           ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE           ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE           ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE    ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE            ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE  ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT    ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT    ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE   ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE           ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY ChaincodeMessage_Type = 19
	ChaincodeMessage_OUTCHAIN_ADDRESS     ChaincodeMessage_Type = 20
	ChaincodeMessage_OUTCHAIN_TRANSACTION ChaincodeMessage_Type = 21
	ChaincodeMessage_OUTCHAIN_QUERY       ChaincodeMessage_Type = 22
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
	19: "GET_HISTORY_FOR_KEY",
	20: "OUTCHAIN_ADDRESS",
	21: "OUTCHAIN_TRANSACTION",
	22: "OUTCHAIN_QUERY",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"REGISTER":            1,
	"REGISTERED":          2,
	"INIT":                3,
	"READY":               4,
	"TRANSACTION":         5,
	"COMPLETED":           6,
	"ERROR":               7,
	"GET_STATE":           8,
	"PUT_STATE":           9,
	"DEL_STATE":           10,
	"INVOKE_CHAINCODE":    11,
	"RESPONSE":            13,
	"GET_STATE_BY_RANGE":  14,
	"GET_QUERY_RESULT":    15,
	"QUERY_STATE_NEXT":    16,
	"QUERY_STATE_CLOSE":   17,
	"KEEPALIVE":           18,
	"GET_HISTORY_FOR_KEY": 19,
	"OUTCHAIN_ADDRESS":     20,
	"OUTCHAIN_TRANSACTION": 21,
	"OUTCHAIN_QUERY":       22,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type ChaincodeMessage struct {
	Type      ChaincodeMessage_Type      `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                     `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid      string                     `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	Proposal  *SignedProposal             `protobuf:"bytes,5,opt,name=proposal" json:"proposal,omitempty"`
	// event emitted by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
	// channel id
	ChannelId string `protobuf:"bytes,7,opt,name=channel_id,json=channelId" json:"channel_id,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if m != nil {
		return m.Type
	}
	return ChaincodeMessage_UNDEFINED
}

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ChaincodeMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ChaincodeMessage) GetProposal() *SignedProposal {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

func (m *ChaincodeMessage) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

type GetState struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
}

func (m *GetState) Reset()                    { *m = GetState{} }
func (m *GetState) String() string            { return proto.CompactTextString(m) }
func (*GetState) ProtoMessage()               {}
func (*GetState) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *GetState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type PutState struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Collection string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *PutState) Reset()                    { *m = PutState{} }
func (m *PutState) String() string            { return proto.CompactTextString(m) }
func (*PutState) ProtoMessage()               {}
func (*PutState) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *PutState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutState) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *PutState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type DelState struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
}

func (m *DelState) Reset()                    { *m = DelState{} }
func (m *DelState) String() string            { return proto.CompactTextString(m) }
func (*DelState) ProtoMessage()               {}
func (*DelState) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *DelState) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelState) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetStateByRange struct {
	StartKey   string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey     string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
	Collection string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *GetStateByRange) GetStartKey() string {
	if m != nil {
		return m.StartKey
	}
	return ""
}

func (m *GetStateByRange) GetEndKey() string {
	if m != nil {
		return m.EndKey
	}
	return ""
}

func (m *GetStateByRange) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetQueryResult struct {
	Query      string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Collection string `protobuf:"bytes,2,opt,name=collection" json:"collection,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *GetQueryResult) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *GetQueryResult) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type GetHistoryForKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetHistoryForKey) Reset()                    { *m = GetHistoryForKey{} }
func (m *GetHistoryForKey) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryForKey) ProtoMessage()               {}
func (*GetHistoryForKey) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *GetHistoryForKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *QueryStateNext) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *QueryStateClose) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryResultBytes struct {
	ResultBytes []byte `protobuf:"bytes,1,opt,name=resultBytes,proto3" json:"resultBytes,omitempty"`
}

func (m *QueryResultBytes) Reset()                    { *m = QueryResultBytes{} }
func (m *QueryResultBytes) String() string            { return proto.CompactTextString(m) }
func (*QueryResultBytes) ProtoMessage()               {}
func (*QueryResultBytes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *QueryResultBytes) GetResultBytes() []byte {
	if m != nil {
		return m.ResultBytes
	}
	return nil
}

type QueryResponse struct {
	Results []*QueryResultBytes `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
	HasMore bool                `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id      string              `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func (m *QueryResponse) GetResults() []*QueryResultBytes {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *QueryResponse) GetHasMore() bool {
	if m != nil {
		return m.HasMore
	}
	return false
}

func (m *QueryResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OutChainAddress struct {
	OutChainName string `protobuf:"bytes,1,opt,name=OutChainName" json:"OutChainName,omitempty"`
	Params       []byte `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Collection   string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *OutChainAddress) Reset()                    { *m = OutChainAddress{} }
func (m *OutChainAddress) String() string            { return proto.CompactTextString(m) }
func (*OutChainAddress) ProtoMessage()               {}
func (*OutChainAddress) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{11} }

func (m *OutChainAddress) GetOutChainName() string {
	if m != nil {
		return m.OutChainName
	}
	return ""
}

func (m *OutChainAddress) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *OutChainAddress) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type OutChainTransaction struct {
	OutChainName string `protobuf:"bytes,1,opt,name=OutChainName" json:"OutChainName,omitempty"`
	Params       []byte `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Collection   string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *OutChainTransaction) Reset()                    { *m = OutChainTransaction{} }
func (m *OutChainTransaction) String() string            { return proto.CompactTextString(m) }
func (*OutChainTransaction) ProtoMessage()               {}
func (*OutChainTransaction) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{12} }

func (m *OutChainTransaction) GetOutChainName() string {
	if m != nil {
		return m.OutChainName
	}
	return ""
}

func (m *OutChainTransaction) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *OutChainTransaction) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type OutChainQuery struct {
	OutChainName string `protobuf:"bytes,1,opt,name=OutChainName" json:"OutChainName,omitempty"`
	Params       []byte `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	Collection   string `protobuf:"bytes,3,opt,name=collection" json:"collection,omitempty"`
}

func (m *OutChainQuery) Reset()                    { *m = OutChainQuery{} }
func (m *OutChainQuery) String() string            { return proto.CompactTextString(m) }
func (*OutChainQuery) ProtoMessage()               {}
func (*OutChainQuery) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{13} }

func (m *OutChainQuery) GetOutChainName() string {
	if m != nil {
		return m.OutChainName
	}
	return ""
}

func (m *OutChainQuery) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *OutChainQuery) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func init() {
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*GetState)(nil), "protos.GetState")
	proto.RegisterType((*PutState)(nil), "protos.PutState")
	proto.RegisterType((*DelState)(nil), "protos.DelState")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*GetHistoryForKey)(nil), "protos.GetHistoryForKey")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryResultBytes)(nil), "protos.QueryResultBytes")
	proto.RegisterType((*QueryResponse)(nil), "protos.QueryResponse")
	proto.RegisterType((*OutChainAddress)(nil), "protos.OutChainAddress")
	proto.RegisterType((*OutChainTransaction)(nil), "protos.OutChainTransaction")
	proto.RegisterType((*OutChainQuery)(nil), "protos.OutChainQuery")
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer/chaincode_shim.proto",
}

func init() { proto.RegisterFile("peer/chaincode_shim.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5d, 0x4f, 0xe3, 0x46,
	0x14, 0xdd, 0x90, 0x00, 0xce, 0x0d, 0x24, 0xb3, 0x03, 0x4b, 0x5d, 0xa4, 0x6d, 0xd3, 0xa8, 0x0f,
	0x79, 0x69, 0xd2, 0xa6, 0x7d, 0xe8, 0xc3, 0x4a, 0x6d, 0x88, 0x87, 0x60, 0x01, 0x76, 0x76, 0x6c,
	0xd0, 0xd2, 0x17, 0xcb, 0xc4, 0xd3, 0xc4, 0x5a, 0xc7, 0xe3, 0xf5, 0x4c, 0xd0, 0xe6, 0x37, 0xf4,
	0xc7, 0xf5, 0x2f, 0x55, 0xe3, 0x8f, 0x10, 0x40, 0x08, 0xa9, 0xab, 0x7d, 0x8a, 0xcf, 0xb9, 0x67,
	0xce, 0x3d, 0x77, 0x3c, 0xf2, 0x04, 0x0e, 0xa7, 0x73, 0x3f, 0x8c, 0xa7, 0x3c, 0x60, 0x9e, 0x98,
	0x87, 0x8b, 0x5e, 0x92, 0x72, 0xc9, 0xf1, 0x4e, 0xf6, 0x23, 0x8e, 0xaf, 0x67, 0xa1, 0x9c, 0x2f,
	0x6f, 0x7b, 0x53, 0xbe, 0xe8, 0x27, 0x7e, 0x14, 0x31, 0xc9, 0x63, 0xd6, 0x9f, 0xf1, 0x9f, 0xee,
	0xc1, 0x94, 0xa7, 0xac, 0x7f, 0xb7, 0x18, 0xf1, 0x58, 0xa6, 0xfe, 0x54, 0x4e, 0x96, 0xb7, 0xfd,
	0x7c, 0x71, 0x3f, 0x61, 0x2c, 0xed, 0xdf, 0xdb, 0xb3, 0x3b, 0x16, 0xcb, 0xdc, 0xff, 0xd8, 0xfe,
	0x62, 0xdf, 0x24, 0xe5, 0x09, 0x17, 0x7e, 0x54, 0x18, 0x7e, 0x3f, 0xe3, 0x7c, 0x16, 0xb1, 0x5c,
	0x72, 0xbb, 0xfc, 0xbb, 0x2f, 0xc3, 0x05, 0x13, 0xd2, 0x5f, 0x24, 0xb9, 0xa0, 0xf3, 0xef, 0x36,
	0xa0, 0x51, 0x99, 0xe5, 0x92, 0x09, 0xe1, 0xcf, 0x18, 0xfe, 0x05, 0x6a, 0x72, 0x95, 0x30, 0xbd,
	0xd2, 0xae, 0x74, 0x9b, 0x83, 0xb7, 0xb9, 0x54, 0xf4, 0x1e, 0xeb, 0x7a, 0xee, 0x2a, 0x61, 0x34,
	0x93, 0xe2, 0xdf, 0xa1, 0xbe, 0xb6, 0xd6, 0xb7, 0xda, 0x95, 0x6e, 0x63, 0x70, 0xdc, 0xcb, 0x9b,
	0xf7, 0xca, 0xe6, 0x3d, 0xb7, 0x54, 0xd0, 0x7b, 0x31, 0xd6, 0x61, 0x37, 0xf1, 0x57, 0x11, 0xf7,
	0x03, 0xbd, 0xda, 0xae, 0x74, 0xf7, 0x68, 0x09, 0x31, 0x86, 0x9a, 0xfc, 0x1c, 0x06, 0x7a, 0xad,
	0x5d, 0xe9, 0xd6, 0x69, 0xf6, 0x8c, 0x07, 0xa0, 0x95, 0x23, 0xea, 0xdb, 0x59, 0x9b, 0xa3, 0x32,
	0x9e, 0x13, 0xce, 0x62, 0x16, 0x4c, 0x8a, 0x2a, 0x5d, 0xeb, 0xf0, 0x1f, 0xd0, 0x7a, 0xb4, 0xdd,
	0xfa, 0xce, 0xc3, 0xa5, 0xeb, 0xc9, 0x88, 0xaa, 0xd2, 0xe6, 0xf4, 0x01, 0xc6, 0x6f, 0x01, 0xa6,
	0x73, 0x3f, 0x8e, 0x59, 0xe4, 0x85, 0x81, 0xbe, 0x9b, 0xc5, 0xa9, 0x17, 0x8c, 0x19, 0x74, 0xfe,
	0xa9, 0x42, 0x4d, 0x6d, 0x05, 0xde, 0x87, 0xfa, 0x95, 0x65, 0x90, 0x53, 0xd3, 0x22, 0x06, 0x7a,
	0x85, 0xf7, 0x40, 0xa3, 0x64, 0x6c, 0x3a, 0x2e, 0xa1, 0xa8, 0x82, 0x9b, 0x00, 0x25, 0x22, 0x06,
	0xda, 0xc2, 0x1a, 0xd4, 0x4c, 0xcb, 0x74, 0x51, 0x15, 0xd7, 0x61, 0x9b, 0x92, 0xa1, 0x71, 0x83,
	0x6a, 0xb8, 0x05, 0x0d, 0x97, 0x0e, 0x2d, 0x67, 0x38, 0x72, 0x4d, 0xdb, 0x42, 0xdb, 0xca, 0x72,
	0x64, 0x5f, 0x4e, 0x2e, 0x88, 0x4b, 0x0c, 0xb4, 0xa3, 0xa4, 0x84, 0x52, 0x9b, 0xa2, 0x5d, 0x55,
	0x19, 0x13, 0xd7, 0x73, 0xdc, 0xa1, 0x4b, 0x90, 0xa6, 0xe0, 0xe4, 0xaa, 0x84, 0x75, 0x05, 0x0d,
	0x72, 0x51, 0x40, 0xc0, 0x87, 0x80, 0x4c, 0xeb, 0xda, 0x3e, 0x27, 0xde, 0xe8, 0x6c, 0x68, 0x5a,
	0x23, 0xdb, 0x20, 0xa8, 0x91, 0x07, 0x74, 0x26, 0xb6, 0xe5, 0x10, 0xb4, 0x8f, 0x8f, 0x00, 0xaf,
	0x0d, 0xbd, 0x93, 0x1b, 0x8f, 0x0e, 0xad, 0x31, 0x41, 0x4d, 0xb5, 0x56, 0xf1, 0xef, 0xaf, 0x08,
	0xbd, 0xf1, 0x28, 0x71, 0xae, 0x2e, 0x5c, 0xd4, 0x52, 0x6c, 0xce, 0xe4, 0x7a, 0x8b, 0x7c, 0x70,
	0x11, 0xc2, 0x6f, 0xe0, 0xf5, 0x26, 0x3b, 0xba, 0xb0, 0x1d, 0x82, 0x5e, 0xab, 0x34, 0xe7, 0x84,
	0x4c, 0x86, 0x17, 0xe6, 0x35, 0x41, 0x18, 0x7f, 0x03, 0x07, 0xca, 0xf1, 0xcc, 0x74, 0x5c, 0x9b,
	0xde, 0x78, 0xa7, 0x36, 0xf5, 0xce, 0xc9, 0x0d, 0x3a, 0x50, 0xa6, 0xf6, 0x95, 0x9b, 0x45, 0xf4,
	0x86, 0x86, 0x41, 0x89, 0xe3, 0xa0, 0x43, 0xac, 0xc3, 0xe1, 0x9a, 0xdd, 0xdc, 0x9d, 0x37, 0x18,
	0x43, 0x73, 0x5d, 0xc9, 0xfa, 0xa2, 0xa3, 0xce, 0x3b, 0xd0, 0xc6, 0x4c, 0x3a, 0xd2, 0x97, 0x0c,
	0x23, 0xa8, 0x7e, 0x64, 0xab, 0xec, 0x1c, 0xd7, 0xa9, 0x7a, 0xc4, 0xdf, 0x01, 0x4c, 0x79, 0x14,
	0xb1, 0xa9, 0x0c, 0x79, 0x9c, 0x1d, 0xd4, 0x3a, 0xdd, 0x60, 0x3a, 0x14, 0xb4, 0xc9, 0xf2, 0xd9,
	0xd5, 0x87, 0xb0, 0x7d, 0xe7, 0x47, 0x4b, 0x96, 0x2d, 0xdc, 0xa3, 0x39, 0x78, 0xe4, 0x59, 0x7d,
	0xe2, 0xf9, 0x0e, 0x34, 0x83, 0x45, 0xff, 0x37, 0x11, 0x83, 0x56, 0x39, 0xcf, 0xc9, 0x8a, 0xfa,
	0xf1, 0x8c, 0xe1, 0x63, 0xd0, 0x84, 0xf4, 0x53, 0x79, 0xbe, 0x76, 0x5a, 0x63, 0x7c, 0x04, 0x3b,
	0x2c, 0x0e, 0x54, 0x25, 0xb7, 0x2a, 0xd0, 0x8b, 0x21, 0x4f, 0xa1, 0x39, 0x66, 0xf2, 0xfd, 0x92,
	0xa5, 0x2b, 0xca, 0xc4, 0x32, 0x92, 0x6a, 0xd8, 0x4f, 0x0a, 0x16, 0x2d, 0x72, 0xf0, 0x62, 0xdc,
	0x1f, 0x01, 0x8d, 0x99, 0x3c, 0x0b, 0x85, 0xe4, 0xe9, 0xea, 0x94, 0xa7, 0xaa, 0xf7, 0x93, 0xa1,
	0x3b, 0x6d, 0x68, 0x66, 0xad, 0xb2, 0xb1, 0x2c, 0xf6, 0x59, 0xe2, 0x26, 0x6c, 0x85, 0x41, 0x21,
	0xd9, 0x0a, 0x83, 0xce, 0x0f, 0xd0, 0xba, 0x57, 0x8c, 0x22, 0x2e, 0xd8, 0x13, 0xc9, 0x6f, 0x80,
	0x36, 0xf2, 0x9e, 0xac, 0x24, 0x13, 0xb8, 0x0d, 0x8d, 0xf4, 0x1e, 0x66, 0xe2, 0x3d, 0xba, 0x49,
	0x75, 0x62, 0xd8, 0x2f, 0x57, 0x25, 0x3c, 0x16, 0x0c, 0x0f, 0x60, 0x37, 0xaf, 0x2b, 0x79, 0xb5,
	0xdb, 0x18, 0xe8, 0xe5, 0x67, 0xe1, 0xb1, 0x3b, 0x2d, 0x85, 0xf8, 0x5b, 0xd0, 0xe6, 0xbe, 0xf0,
	0x16, 0x3c, 0xcd, 0xcf, 0x82, 0x46, 0x77, 0xe7, 0xbe, 0xb8, 0xe4, 0x69, 0x99, 0xb2, 0xba, 0x4e,
	0xb9, 0x80, 0x96, 0xbd, 0x94, 0xd9, 0x17, 0x66, 0x18, 0x04, 0x29, 0x13, 0x02, 0x77, 0x60, 0xaf,
	0xa4, 0x2c, 0x7f, 0xc1, 0x8a, 0x91, 0x1e, 0x70, 0xea, 0x3d, 0x26, 0x7e, 0xea, 0x2f, 0x44, 0x71,
	0xd6, 0x0a, 0xf4, 0xe2, 0x7b, 0xfc, 0x04, 0x07, 0xa5, 0x8f, 0x9b, 0xfa, 0xb1, 0xf0, 0x33, 0xfa,
	0xab, 0xb6, 0xfc, 0x08, 0xfb, 0xa5, 0x4f, 0xb6, 0x63, 0x5f, 0xb3, 0xd9, 0xe0, 0xc3, 0xc6, 0x7d,
	0xe5, 0x2c, 0x93, 0x84, 0xa7, 0x12, 0x1b, 0xa0, 0x51, 0x36, 0x0b, 0x85, 0x64, 0x29, 0xd6, 0x9f,
	0xbb, 0xad, 0x8e, 0x9f, 0xad, 0x74, 0x5e, 0x75, 0x2b, 0x3f, 0x57, 0x4e, 0x6c, 0x68, 0x14, 0x02,
	0x75, 0x93, 0xfe, 0xf5, 0xe7, 0x97, 0xde, 0xc5, 0xb7, 0xf9, 0xbf, 0x85, 0x5f, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x31, 0x4e, 0x10, 0x11, 0x4c, 0x08, 0x00, 0x00,
}
