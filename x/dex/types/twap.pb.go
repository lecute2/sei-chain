// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/twap.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Twap struct {
	Pair            *Pair                                  `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	Twap            github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=twap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"twap" yaml:"twap"`
	LookbackSeconds uint64                                 `protobuf:"varint,3,opt,name=lookbackSeconds,proto3" json:"lookbackSeconds,omitempty"`
}

func (m *Twap) Reset()         { *m = Twap{} }
func (m *Twap) String() string { return proto.CompactTextString(m) }
func (*Twap) ProtoMessage()    {}
func (*Twap) Descriptor() ([]byte, []int) {
	return fileDescriptor_10aa4b136085207a, []int{0}
}
func (m *Twap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Twap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Twap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Twap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Twap.Merge(m, src)
}
func (m *Twap) XXX_Size() int {
	return m.Size()
}
func (m *Twap) XXX_DiscardUnknown() {
	xxx_messageInfo_Twap.DiscardUnknown(m)
}

var xxx_messageInfo_Twap proto.InternalMessageInfo

func (m *Twap) GetPair() *Pair {
	if m != nil {
		return m.Pair
	}
	return nil
}

func (m *Twap) GetLookbackSeconds() uint64 {
	if m != nil {
		return m.LookbackSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*Twap)(nil), "seiprotocol.seichain.dex.Twap")
}

func init() { proto.RegisterFile("dex/twap.proto", fileDescriptor_10aa4b136085207a) }

var fileDescriptor_10aa4b136085207a = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x49, 0xad, 0xd0,
	0x2f, 0x29, 0x4f, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x4e, 0xcd, 0x04,
	0xb3, 0x92, 0xf3, 0x73, 0xf4, 0x8a, 0x53, 0x33, 0x93, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0x52, 0x52,
	0x2b, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x52, 0xfa, 0x20, 0x16, 0x44, 0xbd, 0x14, 0x58,
	0x7f, 0x41, 0x62, 0x66, 0x11, 0x84, 0xaf, 0xb4, 0x9d, 0x91, 0x8b, 0x25, 0xa4, 0x3c, 0xb1, 0x40,
	0xc8, 0x88, 0x8b, 0x05, 0x24, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa7, 0x87, 0xcb,
	0x5c, 0xbd, 0x80, 0xc4, 0xcc, 0xa2, 0x20, 0xb0, 0x5a, 0xa1, 0x40, 0x2e, 0x16, 0x90, 0x53, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x9d, 0x6c, 0x4f, 0xdc, 0x93, 0x67, 0xb8, 0x75, 0x4f, 0x5e, 0x2d,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf,
	0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xb9, 0xa4, 0x26,
	0x7f, 0xba, 0x27, 0xcf, 0x5d, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x04, 0x32, 0x43, 0x29, 0x08, 0x6c,
	0x94, 0x90, 0x06, 0x17, 0x7f, 0x4e, 0x7e, 0x7e, 0x76, 0x52, 0x62, 0x72, 0x76, 0x70, 0x6a, 0x72,
	0x7e, 0x5e, 0x4a, 0xb1, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0xba, 0xb0, 0x93, 0xfb, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0x22, 0x39, 0xa0, 0x38, 0x35, 0x53,
	0x17, 0xe6, 0x0f, 0x30, 0x07, 0xec, 0x11, 0xfd, 0x0a, 0x7d, 0x70, 0x38, 0x82, 0xdc, 0x92, 0xc4,
	0x06, 0x96, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x99, 0x67, 0x7b, 0xff, 0x5b, 0x01, 0x00,
	0x00,
}

func (m *Twap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Twap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Twap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LookbackSeconds != 0 {
		i = encodeVarintTwap(dAtA, i, uint64(m.LookbackSeconds))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Twap.Size()
		i -= size
		if _, err := m.Twap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTwap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Pair != nil {
		{
			size, err := m.Pair.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTwap(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTwap(dAtA []byte, offset int, v uint64) int {
	offset -= sovTwap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Twap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pair != nil {
		l = m.Pair.Size()
		n += 1 + l + sovTwap(uint64(l))
	}
	l = m.Twap.Size()
	n += 1 + l + sovTwap(uint64(l))
	if m.LookbackSeconds != 0 {
		n += 1 + sovTwap(uint64(m.LookbackSeconds))
	}
	return n
}

func sovTwap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTwap(x uint64) (n int) {
	return sovTwap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Twap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwap
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
			return fmt.Errorf("proto: Twap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Twap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
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
				return ErrInvalidLengthTwap
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pair == nil {
				m.Pair = &Pair{}
			}
			if err := m.Pair.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Twap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
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
				return ErrInvalidLengthTwap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Twap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LookbackSeconds", wireType)
			}
			m.LookbackSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LookbackSeconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTwap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwap
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
func skipTwap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTwap
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
					return 0, ErrIntOverflowTwap
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
					return 0, ErrIntOverflowTwap
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
				return 0, ErrInvalidLengthTwap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTwap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTwap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTwap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTwap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTwap = fmt.Errorf("proto: unexpected end of group")
)
