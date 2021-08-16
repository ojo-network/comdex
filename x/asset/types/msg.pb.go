// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/msg.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_bandprotocol_bandchain_packet_packet "github.com/bandprotocol/bandchain-packet/packet"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgFetchPriceFromBandRequest struct {
	From             string                                                                  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	SourceChannel    string                                                                  `protobuf:"bytes,2,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	TimeoutHeight    types.Height                                                            `protobuf:"bytes,3,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	TimeoutTimestamp uint64                                                                  `protobuf:"varint,4,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	PacketData       github_com_bandprotocol_bandchain_packet_packet.OracleRequestPacketData `protobuf:"bytes,5,opt,name=packet_data,json=packetData,proto3,customtype=github.com/bandprotocol/bandchain-packet/packet.OracleRequestPacketData" json:"packet_data" yaml:"packet"`
}

func (m *MsgFetchPriceFromBandRequest) Reset()         { *m = MsgFetchPriceFromBandRequest{} }
func (m *MsgFetchPriceFromBandRequest) String() string { return proto.CompactTextString(m) }
func (*MsgFetchPriceFromBandRequest) ProtoMessage()    {}
func (*MsgFetchPriceFromBandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_71dc8806a9f74763, []int{0}
}
func (m *MsgFetchPriceFromBandRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFetchPriceFromBandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFetchPriceFromBandRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFetchPriceFromBandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFetchPriceFromBandRequest.Merge(m, src)
}
func (m *MsgFetchPriceFromBandRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgFetchPriceFromBandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFetchPriceFromBandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFetchPriceFromBandRequest proto.InternalMessageInfo

type MsgFetchPriceFromBandResponse struct {
}

func (m *MsgFetchPriceFromBandResponse) Reset()         { *m = MsgFetchPriceFromBandResponse{} }
func (m *MsgFetchPriceFromBandResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFetchPriceFromBandResponse) ProtoMessage()    {}
func (*MsgFetchPriceFromBandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_71dc8806a9f74763, []int{1}
}
func (m *MsgFetchPriceFromBandResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFetchPriceFromBandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFetchPriceFromBandResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFetchPriceFromBandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFetchPriceFromBandResponse.Merge(m, src)
}
func (m *MsgFetchPriceFromBandResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFetchPriceFromBandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFetchPriceFromBandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFetchPriceFromBandResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFetchPriceFromBandRequest)(nil), "comdex.asset.v1beta1.MsgFetchPriceFromBandRequest")
	proto.RegisterType((*MsgFetchPriceFromBandResponse)(nil), "comdex.asset.v1beta1.MsgFetchPriceFromBandResponse")
}

func init() { proto.RegisterFile("comdex/asset/v1beta1/msg.proto", fileDescriptor_71dc8806a9f74763) }

var fileDescriptor_71dc8806a9f74763 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x63, 0x56, 0x90, 0x70, 0xe9, 0x80, 0x68, 0x93, 0x42, 0xb5, 0x25, 0x55, 0xb8, 0xf4,
	0x32, 0x47, 0xed, 0x6e, 0x9c, 0x50, 0x40, 0x03, 0x0e, 0x13, 0x23, 0x70, 0xe2, 0xd2, 0x39, 0xee,
	0xd7, 0xc4, 0x22, 0x89, 0x43, 0xec, 0x56, 0xec, 0x06, 0x4f, 0x00, 0x8f, 0xc1, 0x13, 0xf0, 0x0c,
	0x3d, 0xee, 0x88, 0x38, 0x44, 0xd0, 0xbe, 0x41, 0x9f, 0x00, 0x25, 0x4e, 0x07, 0x95, 0x0a, 0xd2,
	0x4e, 0xf6, 0xf7, 0xff, 0xff, 0x3e, 0xfb, 0xcb, 0x97, 0xcf, 0xd8, 0x66, 0x22, 0x1d, 0xc3, 0x07,
	0x8f, 0x4a, 0x09, 0xca, 0x9b, 0x0d, 0x42, 0x50, 0x74, 0xe0, 0xa5, 0x32, 0x22, 0x79, 0x21, 0x94,
	0x30, 0xf7, 0xb4, 0x4f, 0x6a, 0x9f, 0x34, 0x7e, 0x77, 0x2f, 0x12, 0x91, 0xa8, 0x01, 0xaf, 0xda,
	0x69, 0xb6, 0xeb, 0xf0, 0x90, 0x79, 0x4c, 0x14, 0xe0, 0xb1, 0x84, 0x43, 0x56, 0x1d, 0xd7, 0xec,
	0x34, 0xe0, 0x7e, 0xdb, 0xc1, 0x07, 0xa7, 0x32, 0x3a, 0x01, 0xc5, 0xe2, 0xb3, 0x82, 0x33, 0x38,
	0x29, 0x44, 0xea, 0xd3, 0x6c, 0x1c, 0xc0, 0xfb, 0x29, 0x48, 0x65, 0x3e, 0xc4, 0xad, 0x49, 0x21,
	0x52, 0x0b, 0xf5, 0x50, 0xff, 0xb6, 0x7f, 0x77, 0x55, 0x3a, 0xed, 0x0b, 0x9a, 0x26, 0x8f, 0xdc,
	0x4a, 0x75, 0x83, 0xda, 0x34, 0x1f, 0xe3, 0x5d, 0x29, 0xa6, 0x05, 0x83, 0x11, 0x8b, 0x69, 0x96,
	0x41, 0x62, 0xdd, 0xa8, 0xf1, 0x07, 0xab, 0xd2, 0xd9, 0xd7, 0xf8, 0xa6, 0xef, 0x06, 0x1d, 0x2d,
	0x3c, 0xd1, 0xb1, 0x79, 0x8e, 0x77, 0x15, 0x4f, 0x41, 0x4c, 0xd5, 0x28, 0x06, 0x1e, 0xc5, 0xca,
	0xda, 0xe9, 0xa1, 0x7e, 0x7b, 0xd8, 0x25, 0x3c, 0x64, 0xa4, 0xfa, 0x02, 0xd2, 0xd4, 0x3d, 0x1b,
	0x90, 0xe7, 0x35, 0xe1, 0x1f, 0xce, 0x4b, 0xc7, 0xf8, 0x73, 0xc3, 0x66, 0xbe, 0x1b, 0x74, 0x1a,
	0x41, 0xd3, 0xe6, 0x0b, 0x7c, 0x7f, 0x4d, 0x54, 0xab, 0x54, 0x34, 0xcd, 0xad, 0x56, 0x0f, 0xf5,
	0x5b, 0xfe, 0xc1, 0xaa, 0x74, 0xac, 0xcd, 0x43, 0xae, 0x10, 0x37, 0xb8, 0xd7, 0x68, 0x6f, 0xd6,
	0x92, 0xf9, 0x09, 0xe1, 0x76, 0x4e, 0xd9, 0x3b, 0x50, 0xa3, 0x31, 0x55, 0xd4, 0xba, 0xd9, 0x43,
	0xfd, 0x3b, 0xfe, 0x79, 0x55, 0xce, 0x8f, 0xd2, 0x79, 0x16, 0x71, 0x15, 0x4f, 0x43, 0xc2, 0x44,
	0xea, 0x85, 0x34, 0x1b, 0xd7, 0x8d, 0x66, 0x22, 0xa9, 0x03, 0x16, 0x53, 0x9e, 0x1d, 0xe9, 0x64,
	0x4f, 0x2f, 0xe4, 0x65, 0x41, 0x59, 0x02, 0x4d, 0xc3, 0xcf, 0x6a, 0xed, 0x29, 0x55, 0x74, 0x55,
	0x3a, 0x1d, 0x5d, 0x94, 0xe6, 0xdc, 0x00, 0xe7, 0x57, 0xa6, 0xeb, 0xe0, 0xc3, 0x7f, 0xfc, 0x37,
	0x99, 0x8b, 0x4c, 0xc2, 0xf0, 0x33, 0xc2, 0xf8, 0x54, 0x46, 0xaf, 0xa1, 0x98, 0x71, 0x06, 0xe6,
	0x47, 0x84, 0xf7, 0xb7, 0x26, 0x98, 0x43, 0xb2, 0x6d, 0xa0, 0xc8, 0xff, 0xa6, 0xa2, 0x7b, 0x7c,
	0xad, 0x1c, 0x5d, 0x91, 0xff, 0x6a, 0xfe, 0xcb, 0x36, 0xbe, 0x2e, 0x6c, 0x63, 0xbe, 0xb0, 0xd1,
	0xe5, 0xc2, 0x46, 0x3f, 0x17, 0x36, 0xfa, 0xb2, 0xb4, 0x8d, 0xcb, 0xa5, 0x6d, 0x7c, 0x5f, 0xda,
	0xc6, 0x5b, 0xef, 0xaf, 0xd6, 0xe9, 0x0b, 0x8e, 0xc4, 0x64, 0xc2, 0x19, 0xa7, 0x49, 0x13, 0x7b,
	0xeb, 0x77, 0xa1, 0x2e, 0x72, 0x90, 0xe1, 0xad, 0xba, 0xb9, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x63, 0xf5, 0x3e, 0x87, 0x34, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	MsgFetchPriceFromBand(ctx context.Context, in *MsgFetchPriceFromBandRequest, opts ...grpc.CallOption) (*MsgFetchPriceFromBandResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) MsgFetchPriceFromBand(ctx context.Context, in *MsgFetchPriceFromBandRequest, opts ...grpc.CallOption) (*MsgFetchPriceFromBandResponse, error) {
	out := new(MsgFetchPriceFromBandResponse)
	err := c.cc.Invoke(ctx, "/comdex.asset.v1beta1.MsgService/MsgFetchPriceFromBand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	MsgFetchPriceFromBand(context.Context, *MsgFetchPriceFromBandRequest) (*MsgFetchPriceFromBandResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) MsgFetchPriceFromBand(ctx context.Context, req *MsgFetchPriceFromBandRequest) (*MsgFetchPriceFromBandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MsgFetchPriceFromBand not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_MsgFetchPriceFromBand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFetchPriceFromBandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).MsgFetchPriceFromBand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.asset.v1beta1.MsgService/MsgFetchPriceFromBand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).MsgFetchPriceFromBand(ctx, req.(*MsgFetchPriceFromBandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.asset.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MsgFetchPriceFromBand",
			Handler:    _MsgService_MsgFetchPriceFromBand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/asset/v1beta1/msg.proto",
}

func (m *MsgFetchPriceFromBandRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFetchPriceFromBandRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFetchPriceFromBandRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PacketData.Size()
		i -= size
		if _, err := m.PacketData.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsg(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFetchPriceFromBandResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFetchPriceFromBandResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFetchPriceFromBandResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgFetchPriceFromBandRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsg(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsg(uint64(m.TimeoutTimestamp))
	}
	l = m.PacketData.Size()
	n += 1 + l + sovMsg(uint64(l))
	return n
}

func (m *MsgFetchPriceFromBandResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgFetchPriceFromBandRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgFetchPriceFromBandRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFetchPriceFromBandRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgFetchPriceFromBandResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgFetchPriceFromBandResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFetchPriceFromBandResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
