// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/x/fixationstore/types"
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

// GenesisState defines the subscription module's genesis state.
type GenesisState struct {
	Params      Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	SubsFS      types.GenesisState `protobuf:"bytes,2,opt,name=subsFS,proto3" json:"subsFS"`
	SubsTS      []types.RawMessage `protobuf:"bytes,3,rep,name=subsTS,proto3" json:"subsTS"`
	CuTrackerFS types.GenesisState `protobuf:"bytes,4,opt,name=cuTrackerFS,proto3" json:"cuTrackerFS"`
	CuTrackerTS []types.RawMessage `protobuf:"bytes,5,rep,name=cuTrackerTS,proto3" json:"cuTrackerTS"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc6c60f9c112fe52, []int{0}
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

func (m *GenesisState) GetSubsFS() types.GenesisState {
	if m != nil {
		return m.SubsFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetSubsTS() []types.RawMessage {
	if m != nil {
		return m.SubsTS
	}
	return nil
}

func (m *GenesisState) GetCuTrackerFS() types.GenesisState {
	if m != nil {
		return m.CuTrackerFS
	}
	return types.GenesisState{}
}

func (m *GenesisState) GetCuTrackerTS() []types.RawMessage {
	if m != nil {
		return m.CuTrackerTS
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.subscription.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/genesis.proto", fileDescriptor_dc6c60f9c112fe52)
}

var fileDescriptor_dc6c60f9c112fe52 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x93, 0xb6, 0x64, 0x48, 0x98, 0x2c, 0x86, 0x90, 0xc1, 0x14, 0x86, 0x92, 0x01, 0x39,
	0x52, 0x39, 0x00, 0x52, 0x05, 0x61, 0x02, 0x55, 0x75, 0x26, 0x36, 0x27, 0x32, 0x21, 0x82, 0xc6,
	0x91, 0xed, 0x40, 0x7b, 0x0b, 0x8e, 0xd5, 0xb1, 0x23, 0x13, 0x42, 0xc9, 0x41, 0x40, 0x89, 0x03,
	0x8a, 0x87, 0x32, 0x74, 0x7a, 0xb6, 0xf4, 0xfd, 0x9f, 0x7f, 0xe9, 0xd9, 0x3e, 0x7f, 0x21, 0xaf,
	0x24, 0xa7, 0x32, 0x68, 0x66, 0x20, 0xca, 0x58, 0x24, 0x3c, 0x2b, 0x64, 0xc6, 0xf2, 0x20, 0xa5,
	0x39, 0x15, 0x99, 0x40, 0x05, 0x67, 0x92, 0x81, 0xe3, 0x0e, 0x44, 0xcd, 0x44, 0x7d, 0xd0, 0x3b,
	0x4a, 0x59, 0xca, 0x5a, 0x2a, 0x68, 0x4e, 0x2a, 0xe0, 0x4d, 0x76, 0x9b, 0x0b, 0xc2, 0xc9, 0xb2,
	0x13, 0x7b, 0x48, 0xe3, 0x1e, 0xb3, 0x15, 0x69, 0x18, 0x21, 0x19, 0xa7, 0x7f, 0xb7, 0x9b, 0x5c,
	0xf2, 0xb5, 0xe2, 0xcf, 0xbe, 0x07, 0xf6, 0xe1, 0xad, 0xaa, 0x86, 0x25, 0x91, 0x14, 0x5c, 0xd9,
	0x96, 0x12, 0xba, 0xe6, 0xd8, 0xf4, 0x9d, 0xe9, 0x29, 0xda, 0x59, 0x15, 0xcd, 0x5b, 0x70, 0x36,
	0xda, 0x7c, 0x9e, 0x18, 0x8b, 0x2e, 0x06, 0x42, 0xdb, 0x6a, 0xa0, 0x10, 0xbb, 0x83, 0x56, 0xe0,
	0xeb, 0x02, 0xad, 0x12, 0xea, 0x3f, 0xfd, 0xeb, 0x51, 0x69, 0x70, 0xad, 0x3c, 0x11, 0x76, 0x87,
	0xe3, 0xa1, 0xef, 0x4c, 0x27, 0xff, 0x79, 0x16, 0xe4, 0xed, 0x8e, 0x0a, 0x41, 0x52, 0xcd, 0x12,
	0x61, 0x30, 0xb7, 0x9d, 0xa4, 0x8c, 0x38, 0x49, 0x9e, 0x29, 0x0f, 0xb1, 0x3b, 0xda, 0xab, 0x52,
	0x5f, 0x01, 0xee, 0x7b, 0xc6, 0x08, 0xbb, 0x07, 0x7b, 0x94, 0xeb, 0x0b, 0x66, 0xe1, 0xa6, 0x82,
	0xe6, 0xb6, 0x82, 0xe6, 0x57, 0x05, 0xcd, 0xf7, 0x1a, 0x1a, 0xdb, 0x1a, 0x1a, 0x1f, 0x35, 0x34,
	0x1e, 0x2e, 0xd2, 0x4c, 0x3e, 0x95, 0x31, 0x4a, 0xd8, 0x32, 0xd0, 0xd6, 0xba, 0xd2, 0x3f, 0x80,
	0x5c, 0x17, 0x54, 0xc4, 0x56, 0xbb, 0xd0, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x7d,
	0xab, 0xd9, 0x84, 0x02, 0x00, 0x00,
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
	if len(m.CuTrackerTS) > 0 {
		for iNdEx := len(m.CuTrackerTS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CuTrackerTS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.CuTrackerFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.SubsTS) > 0 {
		for iNdEx := len(m.SubsTS) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubsTS[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.SubsFS.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
	l = m.SubsFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SubsTS) > 0 {
		for _, e := range m.SubsTS {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CuTrackerFS.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CuTrackerTS) > 0 {
		for _, e := range m.CuTrackerTS {
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
				return fmt.Errorf("proto: wrong wireType = %d for field SubsFS", wireType)
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
			if err := m.SubsFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubsTS", wireType)
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
			m.SubsTS = append(m.SubsTS, types.RawMessage{})
			if err := m.SubsTS[len(m.SubsTS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuTrackerFS", wireType)
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
			if err := m.CuTrackerFS.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuTrackerTS", wireType)
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
			m.CuTrackerTS = append(m.CuTrackerTS, types.RawMessage{})
			if err := m.CuTrackerTS[len(m.CuTrackerTS)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
