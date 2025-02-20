// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/market/v1beta1/market.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type TimeWeightedAverage struct {
	AssetID             uint64   `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	ScriptID            uint64   `protobuf:"varint,2,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
	Twa                 uint64   `protobuf:"varint,3,opt,name=twa,proto3" json:"twa,omitempty" yaml:"twa"`
	CurrentIndex        uint64   `protobuf:"varint,4,opt,name=current_index,json=currentIndex,proto3" json:"current_index,omitempty" yaml:"current_index"`
	IsPriceActive       bool     `protobuf:"varint,5,opt,name=is_price_active,json=isPriceActive,proto3" json:"is_price_active,omitempty" yaml:"is_price_active"`
	PriceValue          []uint64 `protobuf:"varint,6,rep,packed,name=price_value,json=priceValue,proto3" json:"price_value,omitempty" yaml:"price_value"`
	DiscardedHeightDiff int64    `protobuf:"varint,7,opt,name=discarded_height_diff,json=discardedHeightDiff,proto3" json:"discarded_height_diff,omitempty" yaml:"discarded_height_diff"`
}

func (m *TimeWeightedAverage) Reset()         { *m = TimeWeightedAverage{} }
func (m *TimeWeightedAverage) String() string { return proto.CompactTextString(m) }
func (*TimeWeightedAverage) ProtoMessage()    {}
func (*TimeWeightedAverage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c52e410514c538b6, []int{0}
}
func (m *TimeWeightedAverage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeWeightedAverage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeWeightedAverage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeWeightedAverage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeWeightedAverage.Merge(m, src)
}
func (m *TimeWeightedAverage) XXX_Size() int {
	return m.Size()
}
func (m *TimeWeightedAverage) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeWeightedAverage.DiscardUnknown(m)
}

var xxx_messageInfo_TimeWeightedAverage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TimeWeightedAverage)(nil), "comdex.market.v1beta1.TimeWeightedAverage")
}

func init() {
	proto.RegisterFile("comdex/market/v1beta1/market.proto", fileDescriptor_c52e410514c538b6)
}

var fileDescriptor_c52e410514c538b6 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x13, 0x52, 0xd6, 0x62, 0x36, 0x0d, 0x65, 0x2b, 0x8a, 0x26, 0x14, 0x47, 0x16, 0x12,
	0xb9, 0xd0, 0x30, 0x71, 0x02, 0x89, 0x43, 0xc3, 0x90, 0xe8, 0x01, 0x09, 0x85, 0x0a, 0x24, 0x2e,
	0x91, 0x6b, 0x3b, 0xa9, 0x45, 0xb3, 0x44, 0x89, 0xdb, 0x6e, 0x6f, 0xc1, 0x63, 0xf0, 0x28, 0x3b,
	0xee, 0xc8, 0xc9, 0x02, 0xf7, 0xc2, 0x39, 0x4f, 0x80, 0x62, 0x97, 0xaa, 0x20, 0x6e, 0xdf, 0xf7,
	0xf7, 0xef, 0xf7, 0x5d, 0xfc, 0x07, 0x88, 0x94, 0x05, 0x65, 0x57, 0x51, 0x81, 0xeb, 0x2f, 0x4c,
	0x44, 0xab, 0xf3, 0x19, 0x13, 0xf8, 0x7c, 0xbb, 0x8e, 0xaa, 0xba, 0x14, 0xa5, 0x3b, 0x34, 0xcc,
	0x68, 0x1b, 0x6e, 0x99, 0xb3, 0xd3, 0xbc, 0xcc, 0x4b, 0x4d, 0x44, 0xdd, 0x64, 0x60, 0xf4, 0xcb,
	0x01, 0x27, 0x53, 0x5e, 0xb0, 0x4f, 0x8c, 0xe7, 0x73, 0xc1, 0xe8, 0x78, 0xc5, 0x6a, 0x9c, 0x33,
	0xf7, 0x05, 0x18, 0xe0, 0xa6, 0x61, 0x22, 0xe5, 0xd4, 0xb3, 0x03, 0x3b, 0xec, 0xc5, 0xbe, 0x92,
	0xb0, 0x3f, 0xee, 0xb2, 0xc9, 0x45, 0x2b, 0xe1, 0xf1, 0x35, 0x2e, 0x16, 0x2f, 0xd1, 0x1f, 0x08,
	0x25, 0x7d, 0x3d, 0x4e, 0xa8, 0xfb, 0x0a, 0xdc, 0x6b, 0x48, 0xcd, 0x2b, 0xed, 0xde, 0xd1, 0x6e,
	0xa0, 0x24, 0x1c, 0x7c, 0xd0, 0xa1, 0x96, 0x1f, 0x18, 0x79, 0x87, 0xa1, 0x64, 0x60, 0xe6, 0x09,
	0x75, 0x9f, 0x00, 0x47, 0xac, 0xb1, 0xe7, 0x68, 0x71, 0xa8, 0x24, 0x74, 0xa6, 0x6b, 0xdc, 0x4a,
	0x08, 0x8c, 0x23, 0xd6, 0x18, 0x25, 0x1d, 0xe1, 0xbe, 0x03, 0x47, 0x64, 0x59, 0xd7, 0xec, 0x52,
	0xa4, 0xfc, 0x92, 0xb2, 0x2b, 0xaf, 0xa7, 0x95, 0x50, 0x49, 0x78, 0xf8, 0xda, 0x3c, 0x4c, 0xba,
	0xbc, 0x95, 0xf0, 0xd4, 0xb8, 0x7f, 0xe1, 0x28, 0x39, 0x24, 0x7b, 0x94, 0x1b, 0x83, 0x63, 0xde,
	0xa4, 0x55, 0xcd, 0x09, 0x4b, 0x31, 0x11, 0x7c, 0xc5, 0xbc, 0xbb, 0x81, 0x1d, 0x0e, 0xe2, 0xb3,
	0x56, 0xc2, 0x87, 0xe6, 0xc0, 0x3f, 0x00, 0x4a, 0x8e, 0x78, 0xf3, 0xbe, 0x0b, 0xc6, 0x7a, 0x77,
	0xdf, 0x80, 0xfb, 0xe6, 0x7d, 0x85, 0x17, 0x4b, 0xe6, 0x1d, 0x04, 0x4e, 0xd8, 0x8b, 0x1f, 0x2b,
	0x09, 0x81, 0xa6, 0x3e, 0x76, 0x69, 0x2b, 0xa1, 0x6b, 0xae, 0xed, 0xa1, 0x28, 0x01, 0xd5, 0x8e,
	0x70, 0xa7, 0x60, 0x48, 0x79, 0x43, 0x70, 0x4d, 0x19, 0x4d, 0xe7, 0xfa, 0x67, 0x52, 0xca, 0xb3,
	0xcc, 0xeb, 0x07, 0x76, 0xe8, 0xc4, 0x41, 0x2b, 0xe1, 0x23, 0x73, 0xe2, 0xbf, 0x18, 0x4a, 0x4e,
	0x76, 0xf9, 0x5b, 0x1d, 0x5f, 0xf0, 0x2c, 0x8b, 0x93, 0x9b, 0x9f, 0xbe, 0xf5, 0x4d, 0xf9, 0xd6,
	0x8d, 0xf2, 0xed, 0x5b, 0xe5, 0xdb, 0x3f, 0x94, 0x6f, 0x7f, 0xdd, 0xf8, 0xd6, 0xed, 0xc6, 0xb7,
	0xbe, 0x6f, 0x7c, 0xeb, 0xf3, 0xb3, 0x9c, 0x8b, 0xf9, 0x72, 0x36, 0x22, 0x65, 0x11, 0x99, 0x12,
	0x3d, 0x2d, 0xb3, 0x8c, 0x13, 0x8e, 0x17, 0xdb, 0x3d, 0xda, 0x55, 0x4f, 0x5c, 0x57, 0xac, 0x99,
	0x1d, 0xe8, 0x16, 0x3d, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x50, 0x3a, 0xfc, 0x98, 0x02,
	0x00, 0x00,
}

func (m *TimeWeightedAverage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeWeightedAverage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeWeightedAverage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DiscardedHeightDiff != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.DiscardedHeightDiff))
		i--
		dAtA[i] = 0x38
	}
	if len(m.PriceValue) > 0 {
		dAtA2 := make([]byte, len(m.PriceValue)*10)
		var j1 int
		for _, num := range m.PriceValue {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintMarket(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if m.IsPriceActive {
		i--
		if m.IsPriceActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.CurrentIndex != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.CurrentIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.Twa != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.Twa))
		i--
		dAtA[i] = 0x18
	}
	if m.ScriptID != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x10
	}
	if m.AssetID != 0 {
		i = encodeVarintMarket(dAtA, i, uint64(m.AssetID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMarket(dAtA []byte, offset int, v uint64) int {
	offset -= sovMarket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TimeWeightedAverage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetID != 0 {
		n += 1 + sovMarket(uint64(m.AssetID))
	}
	if m.ScriptID != 0 {
		n += 1 + sovMarket(uint64(m.ScriptID))
	}
	if m.Twa != 0 {
		n += 1 + sovMarket(uint64(m.Twa))
	}
	if m.CurrentIndex != 0 {
		n += 1 + sovMarket(uint64(m.CurrentIndex))
	}
	if m.IsPriceActive {
		n += 2
	}
	if len(m.PriceValue) > 0 {
		l = 0
		for _, e := range m.PriceValue {
			l += sovMarket(uint64(e))
		}
		n += 1 + sovMarket(uint64(l)) + l
	}
	if m.DiscardedHeightDiff != 0 {
		n += 1 + sovMarket(uint64(m.DiscardedHeightDiff))
	}
	return n
}

func sovMarket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMarket(x uint64) (n int) {
	return sovMarket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeWeightedAverage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMarket
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
			return fmt.Errorf("proto: TimeWeightedAverage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeWeightedAverage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			m.AssetID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Twa", wireType)
			}
			m.Twa = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Twa |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentIndex", wireType)
			}
			m.CurrentIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPriceActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsPriceActive = bool(v != 0)
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMarket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.PriceValue = append(m.PriceValue, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMarket
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthMarket
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMarket
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.PriceValue) == 0 {
					m.PriceValue = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowMarket
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.PriceValue = append(m.PriceValue, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceValue", wireType)
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiscardedHeightDiff", wireType)
			}
			m.DiscardedHeightDiff = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMarket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DiscardedHeightDiff |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMarket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMarket
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
func skipMarket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
					return 0, ErrIntOverflowMarket
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
				return 0, ErrInvalidLengthMarket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMarket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMarket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMarket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMarket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMarket = fmt.Errorf("proto: unexpected end of group")
)
