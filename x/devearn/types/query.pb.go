// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sidechain/devearn/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryDevEarnInfosRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDevEarnInfosRequest) Reset()         { *m = QueryDevEarnInfosRequest{} }
func (m *QueryDevEarnInfosRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDevEarnInfosRequest) ProtoMessage()    {}
func (*QueryDevEarnInfosRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{2}
}
func (m *QueryDevEarnInfosRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDevEarnInfosRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDevEarnInfosRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDevEarnInfosRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDevEarnInfosRequest.Merge(m, src)
}
func (m *QueryDevEarnInfosRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDevEarnInfosRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDevEarnInfosRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDevEarnInfosRequest proto.InternalMessageInfo

func (m *QueryDevEarnInfosRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryDevEarnInfosResponse struct {
	// dev_earn_infos is a slice of dev earn infos
	DevEarnInfos []DevEarnInfo `protobuf:"bytes,1,rep,name=dev_earn_infos,json=devEarnInfos,proto3" json:"dev_earn_infos"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDevEarnInfosResponse) Reset()         { *m = QueryDevEarnInfosResponse{} }
func (m *QueryDevEarnInfosResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDevEarnInfosResponse) ProtoMessage()    {}
func (*QueryDevEarnInfosResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{3}
}
func (m *QueryDevEarnInfosResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDevEarnInfosResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDevEarnInfosResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDevEarnInfosResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDevEarnInfosResponse.Merge(m, src)
}
func (m *QueryDevEarnInfosResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDevEarnInfosResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDevEarnInfosResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDevEarnInfosResponse proto.InternalMessageInfo

func (m *QueryDevEarnInfosResponse) GetDevEarnInfos() []DevEarnInfo {
	if m != nil {
		return m.DevEarnInfos
	}
	return nil
}

func (m *QueryDevEarnInfosResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryDevEarnInfoRequest struct {
	Contract string `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (m *QueryDevEarnInfoRequest) Reset()         { *m = QueryDevEarnInfoRequest{} }
func (m *QueryDevEarnInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDevEarnInfoRequest) ProtoMessage()    {}
func (*QueryDevEarnInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{4}
}
func (m *QueryDevEarnInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDevEarnInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDevEarnInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDevEarnInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDevEarnInfoRequest.Merge(m, src)
}
func (m *QueryDevEarnInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDevEarnInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDevEarnInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDevEarnInfoRequest proto.InternalMessageInfo

func (m *QueryDevEarnInfoRequest) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type QueryDevEarnInfoResponse struct {
	DevEarnInfo DevEarnInfo `protobuf:"bytes,1,opt,name=dev_earn_info,json=devEarnInfo,proto3" json:"dev_earn_info"`
}

func (m *QueryDevEarnInfoResponse) Reset()         { *m = QueryDevEarnInfoResponse{} }
func (m *QueryDevEarnInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDevEarnInfoResponse) ProtoMessage()    {}
func (*QueryDevEarnInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e958742cdc1ac27b, []int{5}
}
func (m *QueryDevEarnInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDevEarnInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDevEarnInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDevEarnInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDevEarnInfoResponse.Merge(m, src)
}
func (m *QueryDevEarnInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDevEarnInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDevEarnInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDevEarnInfoResponse proto.InternalMessageInfo

func (m *QueryDevEarnInfoResponse) GetDevEarnInfo() DevEarnInfo {
	if m != nil {
		return m.DevEarnInfo
	}
	return DevEarnInfo{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "sidechain.devearn.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "sidechain.devearn.QueryParamsResponse")
	proto.RegisterType((*QueryDevEarnInfosRequest)(nil), "sidechain.devearn.QueryDevEarnInfosRequest")
	proto.RegisterType((*QueryDevEarnInfosResponse)(nil), "sidechain.devearn.QueryDevEarnInfosResponse")
	proto.RegisterType((*QueryDevEarnInfoRequest)(nil), "sidechain.devearn.QueryDevEarnInfoRequest")
	proto.RegisterType((*QueryDevEarnInfoResponse)(nil), "sidechain.devearn.QueryDevEarnInfoResponse")
}

func init() { proto.RegisterFile("sidechain/devearn/query.proto", fileDescriptor_e958742cdc1ac27b) }

var fileDescriptor_e958742cdc1ac27b = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x2d, 0x44, 0x30, 0x29, 0x48, 0x2c, 0x95, 0x48, 0x0c, 0x2c, 0xa9, 0x11, 0x25,
	0xb4, 0x68, 0x57, 0x4d, 0x85, 0xb8, 0x57, 0xfc, 0x3f, 0xa0, 0x92, 0x23, 0x97, 0x6a, 0x13, 0x2f,
	0xc6, 0x12, 0xdd, 0x75, 0xbd, 0xdb, 0x88, 0x72, 0xe4, 0x09, 0x2a, 0xc1, 0x6b, 0x20, 0xf1, 0x18,
	0x3d, 0x56, 0xe2, 0xc2, 0x09, 0xa1, 0x84, 0x07, 0x41, 0xde, 0xdd, 0xa4, 0x76, 0xe3, 0x2a, 0xb9,
	0x25, 0x9e, 0x6f, 0xbe, 0xf9, 0x7d, 0xe3, 0x91, 0xe1, 0xae, 0x4e, 0x22, 0x31, 0xf8, 0xc8, 0x13,
	0xc9, 0x22, 0x31, 0x14, 0x3c, 0x93, 0xec, 0xe0, 0x50, 0x64, 0x47, 0x34, 0xcd, 0x94, 0x51, 0xf8,
	0xc6, 0xb4, 0x4c, 0x7d, 0x39, 0x58, 0x8d, 0x55, 0xac, 0x6c, 0x95, 0xe5, 0xbf, 0x9c, 0x30, 0xb8,
	0x13, 0x2b, 0x15, 0x7f, 0x12, 0x8c, 0xa7, 0x09, 0xe3, 0x52, 0x2a, 0xc3, 0x4d, 0xa2, 0xa4, 0xf6,
	0xd5, 0x8d, 0x81, 0xd2, 0xfb, 0x4a, 0xb3, 0x3e, 0xd7, 0xc2, 0xf9, 0xb3, 0xe1, 0x56, 0x5f, 0x18,
	0xbe, 0xc5, 0x52, 0x1e, 0x27, 0xd2, 0x8a, 0xbd, 0x96, 0xcc, 0x12, 0xa5, 0x3c, 0xe3, 0xfb, 0xde,
	0x2b, 0x5c, 0x05, 0xfc, 0x2e, 0x77, 0xd8, 0xb5, 0x0f, 0x7b, 0xe2, 0xe0, 0x50, 0x68, 0x13, 0xbe,
	0x85, 0x9b, 0xa5, 0xa7, 0x3a, 0x55, 0x52, 0x0b, 0xfc, 0x14, 0xea, 0xae, 0xb9, 0x89, 0xda, 0xa8,
	0xd3, 0xe8, 0xb6, 0xe8, 0x4c, 0x20, 0xea, 0x5a, 0x76, 0x2e, 0x9d, 0xfc, 0xb9, 0x57, 0xeb, 0x79,
	0x79, 0xd8, 0x87, 0xa6, 0xf5, 0x7b, 0x26, 0x86, 0xcf, 0x79, 0x26, 0x5f, 0xcb, 0x0f, 0x6a, 0x32,
	0x0b, 0xbf, 0x00, 0x38, 0xa3, 0xf6, 0xc6, 0xeb, 0xd4, 0x45, 0xa4, 0x79, 0x44, 0xea, 0x56, 0xe8,
	0x23, 0xd2, 0x5d, 0x1e, 0x0b, 0xdf, 0xdb, 0x2b, 0x74, 0x86, 0x3f, 0x11, 0xb4, 0x2a, 0x86, 0x78,
	0xf4, 0x37, 0x70, 0x3d, 0x12, 0xc3, 0xbd, 0x1c, 0x71, 0x2f, 0xc9, 0x2b, 0x4d, 0xd4, 0x5e, 0xee,
	0x34, 0xba, 0xa4, 0x22, 0x42, 0xc1, 0xc0, 0xe7, 0x58, 0x89, 0x0a, 0x9e, 0xf8, 0x65, 0x89, 0x78,
	0xc9, 0x12, 0x3f, 0x9c, 0x4b, 0xec, 0x40, 0x4a, 0xc8, 0x4f, 0xe0, 0xd6, 0x79, 0xe2, 0xc9, 0x56,
	0x02, 0xb8, 0x32, 0x50, 0xd2, 0x64, 0x7c, 0x60, 0xec, 0x4e, 0xae, 0xf6, 0xa6, 0xff, 0xc3, 0x68,
	0x76, 0x9b, 0xd3, 0x9c, 0xaf, 0xe0, 0x5a, 0x29, 0xa7, 0x5f, 0xe8, 0x62, 0x31, 0x1b, 0x85, 0x98,
	0xdd, 0x1f, 0xcb, 0x70, 0xd9, 0x8e, 0xc1, 0x5f, 0xa0, 0xee, 0xde, 0x2a, 0x7e, 0x50, 0x61, 0x33,
	0x7b, 0x3e, 0xc1, 0xfa, 0x3c, 0x99, 0x83, 0x0d, 0xd7, 0xbe, 0xfe, 0xfa, 0xf7, 0x6d, 0xe9, 0x36,
	0x6e, 0xb1, 0x8b, 0xae, 0x14, 0x7f, 0x47, 0xb0, 0x52, 0x7c, 0xa1, 0x78, 0xf3, 0x22, 0xef, 0x8a,
	0xdb, 0x0a, 0x1e, 0x2f, 0x26, 0xf6, 0x38, 0x8f, 0x2c, 0xce, 0x7d, 0xbc, 0x56, 0x81, 0x53, 0x3e,
	0x1e, 0x7c, 0x8c, 0xa0, 0x51, 0xf0, 0xc0, 0x1b, 0x0b, 0x0c, 0x9a, 0x40, 0x6d, 0x2e, 0xa4, 0xf5,
	0x4c, 0x1d, 0xcb, 0x14, 0xe2, 0xf6, 0x3c, 0xa6, 0x9d, 0xed, 0x93, 0x11, 0x41, 0xa7, 0x23, 0x82,
	0xfe, 0x8e, 0x08, 0x3a, 0x1e, 0x93, 0xda, 0xe9, 0x98, 0xd4, 0x7e, 0x8f, 0x49, 0xed, 0x7d, 0xeb,
	0xac, 0xf5, 0xf3, 0xb4, 0xd9, 0x1c, 0xa5, 0x42, 0xf7, 0xeb, 0xf6, 0x2b, 0xb0, 0xfd, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0xd8, 0x45, 0x19, 0xb9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of DevEarnInfos items.
	DevEarnInfos(ctx context.Context, in *QueryDevEarnInfosRequest, opts ...grpc.CallOption) (*QueryDevEarnInfosResponse, error)
	// Queries a list of DevEarnInfo items.
	DevEarnInfo(ctx context.Context, in *QueryDevEarnInfoRequest, opts ...grpc.CallOption) (*QueryDevEarnInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sidechain.devearn.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DevEarnInfos(ctx context.Context, in *QueryDevEarnInfosRequest, opts ...grpc.CallOption) (*QueryDevEarnInfosResponse, error) {
	out := new(QueryDevEarnInfosResponse)
	err := c.cc.Invoke(ctx, "/sidechain.devearn.Query/DevEarnInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DevEarnInfo(ctx context.Context, in *QueryDevEarnInfoRequest, opts ...grpc.CallOption) (*QueryDevEarnInfoResponse, error) {
	out := new(QueryDevEarnInfoResponse)
	err := c.cc.Invoke(ctx, "/sidechain.devearn.Query/DevEarnInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of DevEarnInfos items.
	DevEarnInfos(context.Context, *QueryDevEarnInfosRequest) (*QueryDevEarnInfosResponse, error)
	// Queries a list of DevEarnInfo items.
	DevEarnInfo(context.Context, *QueryDevEarnInfoRequest) (*QueryDevEarnInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DevEarnInfos(ctx context.Context, req *QueryDevEarnInfosRequest) (*QueryDevEarnInfosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevEarnInfos not implemented")
}
func (*UnimplementedQueryServer) DevEarnInfo(ctx context.Context, req *QueryDevEarnInfoRequest) (*QueryDevEarnInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DevEarnInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidechain.devearn.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DevEarnInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDevEarnInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DevEarnInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidechain.devearn.Query/DevEarnInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DevEarnInfos(ctx, req.(*QueryDevEarnInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DevEarnInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDevEarnInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DevEarnInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sidechain.devearn.Query/DevEarnInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DevEarnInfo(ctx, req.(*QueryDevEarnInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sidechain.devearn.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DevEarnInfos",
			Handler:    _Query_DevEarnInfos_Handler,
		},
		{
			MethodName: "DevEarnInfo",
			Handler:    _Query_DevEarnInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sidechain/devearn/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDevEarnInfosRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDevEarnInfosRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDevEarnInfosRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDevEarnInfosResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDevEarnInfosResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDevEarnInfosResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DevEarnInfos) > 0 {
		for iNdEx := len(m.DevEarnInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DevEarnInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryDevEarnInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDevEarnInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDevEarnInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDevEarnInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDevEarnInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDevEarnInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DevEarnInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDevEarnInfosRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDevEarnInfosResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DevEarnInfos) > 0 {
		for _, e := range m.DevEarnInfos {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDevEarnInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDevEarnInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DevEarnInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDevEarnInfosRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDevEarnInfosRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDevEarnInfosRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDevEarnInfosResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDevEarnInfosResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDevEarnInfosResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEarnInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DevEarnInfos = append(m.DevEarnInfos, DevEarnInfo{})
			if err := m.DevEarnInfos[len(m.DevEarnInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDevEarnInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDevEarnInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDevEarnInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDevEarnInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDevEarnInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDevEarnInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DevEarnInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DevEarnInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
