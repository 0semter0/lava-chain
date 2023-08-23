// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/dualstaking/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/common/types"
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

// GenesisState defines the dualstaking module's genesis state.
type GenesisState struct {
	Params        Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	DelegationsFS []types.RawMessage `protobuf:"bytes,2,rep,name=delegationsFS,proto3" json:"delegationsFS"`
	DelegatorsFS  []types.RawMessage `protobuf:"bytes,3,rep,name=delegatorsFS,proto3" json:"delegatorsFS"`
	UnbondingsTS  []types.RawMessage `protobuf:"bytes,4,rep,name=unbondingsTS,proto3" json:"unbondingsTS"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5bca863c53f218f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetDelegationsFS() []types.RawMessage {
	if m != nil {
		return m.DelegationsFS
	}
	return nil
}

func (m *GenesisState) GetDelegatorsFS() []types.RawMessage {
	if m != nil {
		return m.DelegatorsFS
	}
	return nil
}

func (m *GenesisState) GetUnbondingsTS() []types.RawMessage {
	if m != nil {
		return m.UnbondingsTS
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.dualstaking.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/dualstaking/genesis.proto", fileDescriptor_d5bca863c53f218f)
}

var fileDescriptor_d5bca863c53f218f = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0xc3, 0x40,
	0x1c, 0xc6, 0x93, 0xb6, 0x74, 0x48, 0xeb, 0x12, 0x1c, 0x42, 0x86, 0x6b, 0x10, 0xd4, 0x82, 0x70,
	0x07, 0x75, 0x77, 0x28, 0x54, 0x11, 0x11, 0xa4, 0x71, 0x72, 0xbb, 0x98, 0xf3, 0x3c, 0x4c, 0xee,
	0x42, 0xee, 0xa2, 0xed, 0x5b, 0xf8, 0x2a, 0xbe, 0x45, 0xc7, 0x8e, 0x4e, 0x22, 0xc9, 0x8b, 0x94,
	0x4b, 0x6e, 0xe8, 0x0d, 0x1d, 0x3a, 0xfd, 0x43, 0xf8, 0x7d, 0xbf, 0xfb, 0xe0, 0xf3, 0x2e, 0x32,
	0xfc, 0x89, 0x39, 0x51, 0x48, 0x5f, 0x94, 0x56, 0x38, 0x93, 0x0a, 0x7f, 0x30, 0x4e, 0x11, 0x25,
	0x9c, 0x48, 0x26, 0x61, 0x51, 0x0a, 0x25, 0xfc, 0xc0, 0x70, 0x50, 0x5f, 0xb8, 0xc7, 0x85, 0xa7,
	0x54, 0x50, 0xd1, 0x42, 0x48, 0x7f, 0x75, 0x7c, 0x78, 0x7e, 0xd0, 0x5b, 0xe0, 0x12, 0xe7, 0x46,
	0x1b, 0x5e, 0x5a, 0xd8, 0xab, 0xc8, 0x73, 0xc1, 0xd1, 0x1b, 0x5b, 0x61, 0xc5, 0x04, 0x5f, 0x70,
	0x55, 0xae, 0x3b, 0xf0, 0xec, 0xa7, 0xe7, 0x8d, 0xef, 0xba, 0x46, 0xb1, 0xc2, 0x8a, 0xf8, 0x37,
	0xde, 0xb0, 0x33, 0x05, 0x6e, 0xe4, 0x4e, 0x47, 0xb3, 0x08, 0x1e, 0x6a, 0x08, 0x9f, 0x5a, 0x6e,
	0x3e, 0xd8, 0xfc, 0x4d, 0x9c, 0xa5, 0x49, 0xf9, 0x0f, 0xde, 0x49, 0x4a, 0x32, 0x42, 0xdb, 0x97,
	0xe4, 0x6d, 0x1c, 0xf4, 0xa2, 0xfe, 0x74, 0x34, 0x9b, 0xd8, 0x9a, 0xae, 0x11, 0x5c, 0xe2, 0xaf,
	0x47, 0x22, 0x25, 0xa6, 0xc4, 0x58, 0xec, 0xac, 0x7f, 0xef, 0x8d, 0xcd, 0x0f, 0x51, 0x6a, 0x57,
	0xff, 0x18, 0x97, 0x15, 0xd5, 0xaa, 0x8a, 0x27, 0x82, 0xa7, 0x8c, 0x53, 0xf9, 0x1c, 0x07, 0x83,
	0xa3, 0x54, 0xfb, 0xd1, 0xf9, 0x62, 0x53, 0x03, 0x77, 0x5b, 0x03, 0xf7, 0xbf, 0x06, 0xee, 0x77,
	0x03, 0x9c, 0x6d, 0x03, 0x9c, 0xdf, 0x06, 0x38, 0x2f, 0x57, 0x94, 0xa9, 0xf7, 0x2a, 0xd1, 0x1e,
	0x64, 0x2d, 0xb0, 0xb2, 0xa6, 0x52, 0xeb, 0x82, 0xc8, 0x64, 0xd8, 0x2e, 0x70, 0xbd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xca, 0x6f, 0x3f, 0xe5, 0x2b, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnbondingsTS) > 0 {
		for iNdEx := len(m.UnbondingsTS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingsTS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.DelegatorsFS) > 0 {
		for iNdEx := len(m.DelegatorsFS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorsFS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DelegationsFS) > 0 {
		for iNdEx := len(m.DelegationsFS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationsFS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.DelegationsFS) > 0 {
		for _, e := range m.DelegationsFS {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegatorsFS) > 0 {
		for _, e := range m.DelegatorsFS {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UnbondingsTS) > 0 {
		for _, e := range m.UnbondingsTS {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationsFS = append(m.DelegationsFS, types.RawMessage{})
			if err := m.DelegationsFS[len(m.DelegationsFS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorsFS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorsFS = append(m.DelegatorsFS, types.RawMessage{})
			if err := m.DelegatorsFS[len(m.DelegatorsFS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingsTS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingsTS = append(m.UnbondingsTS, types.RawMessage{})
			if err := m.UnbondingsTS[len(m.UnbondingsTS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
