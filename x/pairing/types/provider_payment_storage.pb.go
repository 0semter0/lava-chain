// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pairing/provider_payment_storage.proto

package types

import (
	fmt "fmt"
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

type ProviderPaymentStorage struct {
	Index                                  string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Epoch                                  uint64   `protobuf:"varint,3,opt,name=epoch,proto3" json:"epoch,omitempty"`
	UniquePaymentStorageClientProviderKeys []string `protobuf:"bytes,6,rep,name=uniquePaymentStorageClientProviderKeys,proto3" json:"uniquePaymentStorageClientProviderKeys,omitempty"`
	ComplainersTotalCu                     uint64   `protobuf:"varint,7,opt,name=complainersTotalCu,proto3" json:"complainersTotalCu,omitempty"`
}

func (m *ProviderPaymentStorage) Reset()         { *m = ProviderPaymentStorage{} }
func (m *ProviderPaymentStorage) String() string { return proto.CompactTextString(m) }
func (*ProviderPaymentStorage) ProtoMessage()    {}
func (*ProviderPaymentStorage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f1d2e8d774659ae, []int{0}
}
func (m *ProviderPaymentStorage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderPaymentStorage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderPaymentStorage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderPaymentStorage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderPaymentStorage.Merge(m, src)
}
func (m *ProviderPaymentStorage) XXX_Size() int {
	return m.Size()
}
func (m *ProviderPaymentStorage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderPaymentStorage.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderPaymentStorage proto.InternalMessageInfo

func (m *ProviderPaymentStorage) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ProviderPaymentStorage) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *ProviderPaymentStorage) GetUniquePaymentStorageClientProviderKeys() []string {
	if m != nil {
		return m.UniquePaymentStorageClientProviderKeys
	}
	return nil
}

func (m *ProviderPaymentStorage) GetComplainersTotalCu() uint64 {
	if m != nil {
		return m.ComplainersTotalCu
	}
	return 0
}

func init() {
	proto.RegisterType((*ProviderPaymentStorage)(nil), "lavanet.lava.pairing.ProviderPaymentStorage")
}

func init() {
	proto.RegisterFile("pairing/provider_payment_storage.proto", fileDescriptor_4f1d2e8d774659ae)
}

var fileDescriptor_4f1d2e8d774659ae = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x48, 0xcc, 0x2c,
	0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xcb, 0x4c, 0x49, 0x2d, 0x8a, 0x2f, 0x48, 0xac,
	0xcc, 0x4d, 0xcd, 0x2b, 0x89, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xc9, 0x49, 0x2c, 0x4b, 0xcc, 0x4b, 0x2d, 0xd1, 0x03, 0xd1, 0x7a, 0x50,
	0x4d, 0x52, 0x26, 0x30, 0xdd, 0xa5, 0x79, 0x99, 0x85, 0xa5, 0xa9, 0xe8, 0x7a, 0xe3, 0x93, 0x73,
	0x32, 0x41, 0x5c, 0x98, 0xd9, 0x10, 0xb3, 0x94, 0xee, 0x31, 0x72, 0x89, 0x05, 0x40, 0x85, 0x02,
	0x20, 0x3a, 0x82, 0x21, 0x1a, 0x84, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x90, 0x68, 0x6a, 0x41, 0x7e, 0x72, 0x86, 0x04, 0xb3,
	0x02, 0xa3, 0x06, 0x4b, 0x10, 0x84, 0x23, 0x14, 0xc6, 0xa5, 0x06, 0xb1, 0x16, 0xd5, 0x0c, 0x67,
	0xb0, 0x9d, 0x30, 0xf3, 0xbd, 0x53, 0x2b, 0x8b, 0x25, 0xd8, 0x14, 0x98, 0x35, 0x38, 0x83, 0x88,
	0x54, 0x2d, 0xa4, 0xc7, 0x25, 0x94, 0x9c, 0x9f, 0x5b, 0x90, 0x93, 0x98, 0x99, 0x97, 0x5a, 0x54,
	0x1c, 0x92, 0x5f, 0x92, 0x98, 0xe3, 0x5c, 0x2a, 0xc1, 0x0e, 0xb6, 0x1a, 0x8b, 0x8c, 0x17, 0x0b,
	0x07, 0x93, 0x00, 0xb3, 0x17, 0x0b, 0x07, 0x8b, 0x00, 0xab, 0x17, 0x0b, 0x07, 0xab, 0x00, 0x9b,
	0x93, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa7, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0x43, 0x14, 0x4c, 0xeb, 0x57, 0xe8, 0xc3,
	0x82, 0xb2, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x54, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x32, 0xb0, 0x09, 0xa0, 0x01, 0x00, 0x00,
}

func (m *ProviderPaymentStorage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderPaymentStorage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderPaymentStorage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ComplainersTotalCu != 0 {
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(m.ComplainersTotalCu))
		i--
		dAtA[i] = 0x38
	}
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for iNdEx := len(m.UniquePaymentStorageClientProviderKeys) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.UniquePaymentStorageClientProviderKeys[iNdEx])
			copy(dAtA[i:], m.UniquePaymentStorageClientProviderKeys[iNdEx])
			i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.UniquePaymentStorageClientProviderKeys[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Epoch != 0 {
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProviderPaymentStorage(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProviderPaymentStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovProviderPaymentStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderPaymentStorage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProviderPaymentStorage(uint64(l))
	}
	if m.Epoch != 0 {
		n += 1 + sovProviderPaymentStorage(uint64(m.Epoch))
	}
	if len(m.UniquePaymentStorageClientProviderKeys) > 0 {
		for _, s := range m.UniquePaymentStorageClientProviderKeys {
			l = len(s)
			n += 1 + l + sovProviderPaymentStorage(uint64(l))
		}
	}
	if m.ComplainersTotalCu != 0 {
		n += 1 + sovProviderPaymentStorage(uint64(m.ComplainersTotalCu))
	}
	return n
}

func sovProviderPaymentStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProviderPaymentStorage(x uint64) (n int) {
	return sovProviderPaymentStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderPaymentStorage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProviderPaymentStorage
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
			return fmt.Errorf("proto: ProviderPaymentStorage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderPaymentStorage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniquePaymentStorageClientProviderKeys", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
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
				return ErrInvalidLengthProviderPaymentStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderPaymentStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UniquePaymentStorageClientProviderKeys = append(m.UniquePaymentStorageClientProviderKeys, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComplainersTotalCu", wireType)
			}
			m.ComplainersTotalCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderPaymentStorage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComplainersTotalCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProviderPaymentStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProviderPaymentStorage
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
func skipProviderPaymentStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
					return 0, ErrIntOverflowProviderPaymentStorage
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
				return 0, ErrInvalidLengthProviderPaymentStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProviderPaymentStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProviderPaymentStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProviderPaymentStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProviderPaymentStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProviderPaymentStorage = fmt.Errorf("proto: unexpected end of group")
)
