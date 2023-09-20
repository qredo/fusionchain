// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/policy/policy.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
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

type Policy struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The actual policy informations. It must be one the supported policy types:
	// - BlackbirdPolicy
	Policy *types.Any `protobuf:"bytes,3,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6c43888f5f679, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return m.Size()
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Policy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Policy) GetPolicy() *types.Any {
	if m != nil {
		return m.Policy
	}
	return nil
}

type BlackbirdPolicy struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BlackbirdPolicy) Reset()         { *m = BlackbirdPolicy{} }
func (m *BlackbirdPolicy) String() string { return proto.CompactTextString(m) }
func (*BlackbirdPolicy) ProtoMessage()    {}
func (*BlackbirdPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6c43888f5f679, []int{1}
}
func (m *BlackbirdPolicy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlackbirdPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlackbirdPolicy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlackbirdPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlackbirdPolicy.Merge(m, src)
}
func (m *BlackbirdPolicy) XXX_Size() int {
	return m.Size()
}
func (m *BlackbirdPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_BlackbirdPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_BlackbirdPolicy proto.InternalMessageInfo

func (m *BlackbirdPolicy) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlackbirdPolicyPayload struct {
	Witness []byte `protobuf:"bytes,1,opt,name=witness,proto3" json:"witness,omitempty"`
}

func (m *BlackbirdPolicyPayload) Reset()         { *m = BlackbirdPolicyPayload{} }
func (m *BlackbirdPolicyPayload) String() string { return proto.CompactTextString(m) }
func (*BlackbirdPolicyPayload) ProtoMessage()    {}
func (*BlackbirdPolicyPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_50a6c43888f5f679, []int{2}
}
func (m *BlackbirdPolicyPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlackbirdPolicyPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlackbirdPolicyPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlackbirdPolicyPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlackbirdPolicyPayload.Merge(m, src)
}
func (m *BlackbirdPolicyPayload) XXX_Size() int {
	return m.Size()
}
func (m *BlackbirdPolicyPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BlackbirdPolicyPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BlackbirdPolicyPayload proto.InternalMessageInfo

func (m *BlackbirdPolicyPayload) GetWitness() []byte {
	if m != nil {
		return m.Witness
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "fusionchain.policy.Policy")
	proto.RegisterType((*BlackbirdPolicy)(nil), "fusionchain.policy.BlackbirdPolicy")
	proto.RegisterType((*BlackbirdPolicyPayload)(nil), "fusionchain.policy.BlackbirdPolicyPayload")
}

func init() { proto.RegisterFile("fusionchain/policy/policy.proto", fileDescriptor_50a6c43888f5f679) }

var fileDescriptor_50a6c43888f5f679 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x18, 0x85, 0x9b, 0x5a, 0x2a, 0x46, 0x51, 0x08, 0x22, 0xd5, 0x45, 0x2c, 0x05, 0xa1, 0xa0, 0x24,
	0x30, 0x9e, 0xc0, 0x01, 0xf7, 0x43, 0x97, 0xb3, 0x4b, 0x9b, 0x4c, 0x27, 0xd8, 0x49, 0x6a, 0x9b,
	0xa2, 0xbd, 0x85, 0xc7, 0x72, 0x39, 0x4b, 0x97, 0xd2, 0x5e, 0x44, 0x48, 0x53, 0x28, 0xae, 0xf2,
	0xc2, 0xff, 0xfd, 0xef, 0xfd, 0x3c, 0x78, 0xbf, 0xeb, 0x5a, 0xa9, 0x55, 0xb1, 0x67, 0x52, 0xd1,
	0x5a, 0x57, 0xb2, 0xe8, 0xdd, 0x43, 0xea, 0x46, 0x1b, 0x8d, 0xd0, 0x02, 0x20, 0xd3, 0xe4, 0xee,
	0xb6, 0xd4, 0xba, 0xac, 0x04, 0xb5, 0x44, 0xde, 0xed, 0x28, 0x53, 0x0e, 0x4f, 0xb6, 0x30, 0xdc,
	0x58, 0x08, 0x5d, 0x42, 0x5f, 0xf2, 0x08, 0xc4, 0x20, 0x0d, 0x32, 0x5f, 0x72, 0x84, 0x60, 0xa0,
	0xd8, 0x41, 0x44, 0x7e, 0x0c, 0xd2, 0xb3, 0xcc, 0x6a, 0xf4, 0x04, 0xc3, 0xc9, 0x32, 0x3a, 0x89,
	0x41, 0x7a, 0xbe, 0xba, 0x26, 0x93, 0x33, 0x99, 0x9d, 0xc9, 0x8b, 0xea, 0x33, 0xc7, 0x24, 0x0f,
	0xf0, 0x6a, 0x5d, 0xb1, 0xe2, 0x2d, 0x97, 0x0d, 0x77, 0x21, 0x08, 0x06, 0x9c, 0x19, 0x66, 0x63,
	0x2e, 0x32, 0xab, 0x93, 0x15, 0xbc, 0xf9, 0x87, 0x6d, 0x58, 0x5f, 0x69, 0xc6, 0x51, 0x04, 0x4f,
	0x3f, 0xa4, 0x51, 0xa2, 0x6d, 0xdd, 0xc2, 0xfc, 0x5d, 0xbf, 0x7e, 0x0f, 0x18, 0x1c, 0x07, 0x0c,
	0x7e, 0x07, 0x0c, 0xbe, 0x46, 0xec, 0x1d, 0x47, 0xec, 0xfd, 0x8c, 0xd8, 0xdb, 0x3e, 0x96, 0xd2,
	0xec, 0xbb, 0x9c, 0x14, 0xfa, 0x40, 0xdf, 0x1b, 0xc1, 0x35, 0x5d, 0x36, 0xf6, 0x39, 0x77, 0x66,
	0xfa, 0x5a, 0xb4, 0x79, 0x68, 0xef, 0x7e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x46, 0x72, 0x85,
	0x8b, 0x56, 0x01, 0x00, 0x00,
}

func (m *Policy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Policy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Policy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPolicy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPolicy(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlackbirdPolicy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlackbirdPolicy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlackbirdPolicy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlackbirdPolicyPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlackbirdPolicyPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlackbirdPolicyPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Witness) > 0 {
		i -= len(m.Witness)
		copy(dAtA[i:], m.Witness)
		i = encodeVarintPolicy(dAtA, i, uint64(len(m.Witness)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPolicy(dAtA []byte, offset int, v uint64) int {
	offset -= sovPolicy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Policy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPolicy(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func (m *BlackbirdPolicy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func (m *BlackbirdPolicyPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Witness)
	if l > 0 {
		n += 1 + l + sovPolicy(uint64(l))
	}
	return n
}

func sovPolicy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPolicy(x uint64) (n int) {
	return sovPolicy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Policy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: Policy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Policy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types.Any{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BlackbirdPolicy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BlackbirdPolicy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlackbirdPolicy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func (m *BlackbirdPolicyPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPolicy
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
			return fmt.Errorf("proto: BlackbirdPolicyPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlackbirdPolicyPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Witness", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPolicy
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
				return ErrInvalidLengthPolicy
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPolicy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Witness = append(m.Witness[:0], dAtA[iNdEx:postIndex]...)
			if m.Witness == nil {
				m.Witness = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPolicy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPolicy
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
func skipPolicy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
					return 0, ErrIntOverflowPolicy
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
				return 0, ErrInvalidLengthPolicy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPolicy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPolicy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPolicy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPolicy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPolicy = fmt.Errorf("proto: unexpected end of group")
)
