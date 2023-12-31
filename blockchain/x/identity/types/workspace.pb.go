// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/identity/workspace.proto

package types

import (
	fmt "fmt"
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

type Workspace struct {
	Address         string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator         string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Owners          []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	ChildWorkspaces []string `protobuf:"bytes,4,rep,name=child_workspaces,json=childWorkspaces,proto3" json:"child_workspaces,omitempty"`
	// Optional ID of the policy to be applied to every *admin* operation.
	// If not specified, the default policy is used.
	//
	// Admin operations are:
	// - fusionchain.identity.Msg.AddWorkspaceOwner
	// - fusionchain.identity.Msg.RemoveWorkspaceOwner
	// - fusionchain.identity.Msg.AppendChildWorkspace
	// - fusionchain.identity.Msg.NewChildWorkspace
	//
	// The default policy is to allow any operation when at least one of its
	// owner approves it.
	AdminPolicyId uint64 `protobuf:"varint,5,opt,name=admin_policy_id,json=adminPolicyId,proto3" json:"admin_policy_id,omitempty"`
	// Optional ID of the policy to be applied to every *sign* operation.
	// If not specified, the default policy is used.
	//
	// Sign operations are:
	// - fusionchain.treasury.Msg.NewKeyRequest
	// - fusionchain.treasury.Msg.NewSignatureRequest
	// - fusionchain.treasury.Msg.NewWalletRequest
	// - fusionchain.treasury.Msg.NewSignTransactionRequest
	//
	// The default policy is to allow any operation when at least one of its
	// owner approves it.
	SignPolicyId uint64 `protobuf:"varint,6,opt,name=sign_policy_id,json=signPolicyId,proto3" json:"sign_policy_id,omitempty"`
}

func (m *Workspace) Reset()         { *m = Workspace{} }
func (m *Workspace) String() string { return proto.CompactTextString(m) }
func (*Workspace) ProtoMessage()    {}
func (*Workspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd72da9924714aca, []int{0}
}
func (m *Workspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Workspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Workspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Workspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workspace.Merge(m, src)
}
func (m *Workspace) XXX_Size() int {
	return m.Size()
}
func (m *Workspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Workspace.DiscardUnknown(m)
}

var xxx_messageInfo_Workspace proto.InternalMessageInfo

func (m *Workspace) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Workspace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Workspace) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *Workspace) GetChildWorkspaces() []string {
	if m != nil {
		return m.ChildWorkspaces
	}
	return nil
}

func (m *Workspace) GetAdminPolicyId() uint64 {
	if m != nil {
		return m.AdminPolicyId
	}
	return 0
}

func (m *Workspace) GetSignPolicyId() uint64 {
	if m != nil {
		return m.SignPolicyId
	}
	return 0
}

func init() {
	proto.RegisterType((*Workspace)(nil), "fusionchain.identity.Workspace")
}

func init() {
	proto.RegisterFile("fusionchain/identity/workspace.proto", fileDescriptor_bd72da9924714aca)
}

var fileDescriptor_bd72da9924714aca = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xbf, 0xfd, 0x83, 0x62, 0x01, 0x45, 0x16, 0x42, 0x9e, 0xac, 0x08, 0x55, 0x28,
	0x0c, 0x24, 0x03, 0x6f, 0xc0, 0x82, 0xd8, 0x50, 0x16, 0x24, 0x96, 0x28, 0xb5, 0x4d, 0x73, 0x45,
	0x6b, 0x07, 0xdb, 0x55, 0xc9, 0x5b, 0xf0, 0x58, 0x8c, 0x19, 0x19, 0x51, 0xf2, 0x22, 0xa8, 0xa6,
	0x29, 0x19, 0xcf, 0x77, 0x3e, 0xe9, 0x5e, 0x1d, 0x3c, 0x7f, 0xd9, 0x58, 0xd0, 0x8a, 0x57, 0x25,
	0xa8, 0x0c, 0x84, 0x54, 0x0e, 0x5c, 0x93, 0x6d, 0xb5, 0x79, 0xb5, 0x75, 0xc9, 0x65, 0x5a, 0x1b,
	0xed, 0x34, 0x39, 0x1f, 0x59, 0xe9, 0x60, 0x5d, 0xb6, 0x08, 0x47, 0x4f, 0x83, 0x49, 0x28, 0x3e,
	0x2a, 0x85, 0x30, 0xd2, 0x5a, 0x8a, 0x62, 0x94, 0x44, 0xf9, 0x10, 0x77, 0x0d, 0x37, 0xb2, 0x74,
	0xda, 0xd0, 0x7f, 0xbf, 0xcd, 0x3e, 0x92, 0x0b, 0x1c, 0xea, 0xad, 0x92, 0xc6, 0xd2, 0x49, 0x3c,
	0x49, 0xa2, 0x7c, 0x9f, 0xc8, 0x35, 0x3e, 0xe3, 0x15, 0xac, 0x44, 0x71, 0x78, 0xc4, 0xd2, 0xa9,
	0x37, 0x66, 0x9e, 0x1f, 0xae, 0x5a, 0x72, 0x85, 0x67, 0xa5, 0x58, 0x83, 0x2a, 0x6a, 0xbd, 0x02,
	0xde, 0x14, 0x20, 0xe8, 0xff, 0x18, 0x25, 0xd3, 0xfc, 0xc4, 0xe3, 0x47, 0x4f, 0x1f, 0x04, 0x99,
	0xe3, 0x53, 0x0b, 0xcb, 0xb1, 0x16, 0x7a, 0xed, 0x78, 0x47, 0x07, 0xeb, 0xee, 0xfe, 0xb3, 0x63,
	0xa8, 0xed, 0x18, 0xfa, 0xee, 0x18, 0xfa, 0xe8, 0x59, 0xd0, 0xf6, 0x2c, 0xf8, 0xea, 0x59, 0xf0,
	0x7c, 0xb3, 0x04, 0x57, 0x6d, 0x16, 0x29, 0xd7, 0xeb, 0xec, 0xcd, 0x48, 0xa1, 0xb3, 0xf1, 0x72,
	0xef, 0x7f, 0xdb, 0xb9, 0xa6, 0x96, 0x76, 0x11, 0xfa, 0xe1, 0x6e, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xc0, 0x25, 0xbe, 0x95, 0x60, 0x01, 0x00, 0x00,
}

func (m *Workspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Workspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Workspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignPolicyId != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.SignPolicyId))
		i--
		dAtA[i] = 0x30
	}
	if m.AdminPolicyId != 0 {
		i = encodeVarintWorkspace(dAtA, i, uint64(m.AdminPolicyId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.ChildWorkspaces) > 0 {
		for iNdEx := len(m.ChildWorkspaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ChildWorkspaces[iNdEx])
			copy(dAtA[i:], m.ChildWorkspaces[iNdEx])
			i = encodeVarintWorkspace(dAtA, i, uint64(len(m.ChildWorkspaces[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWorkspace(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Workspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWorkspace(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovWorkspace(uint64(l))
		}
	}
	if len(m.ChildWorkspaces) > 0 {
		for _, s := range m.ChildWorkspaces {
			l = len(s)
			n += 1 + l + sovWorkspace(uint64(l))
		}
	}
	if m.AdminPolicyId != 0 {
		n += 1 + sovWorkspace(uint64(m.AdminPolicyId))
	}
	if m.SignPolicyId != 0 {
		n += 1 + sovWorkspace(uint64(m.SignPolicyId))
	}
	return n
}

func sovWorkspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkspace(x uint64) (n int) {
	return sovWorkspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Workspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkspace
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
			return fmt.Errorf("proto: Workspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Workspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildWorkspaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
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
				return ErrInvalidLengthWorkspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildWorkspaces = append(m.ChildWorkspaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminPolicyId", wireType)
			}
			m.AdminPolicyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdminPolicyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignPolicyId", wireType)
			}
			m.SignPolicyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignPolicyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWorkspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWorkspace
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
func skipWorkspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
					return 0, ErrIntOverflowWorkspace
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
				return 0, ErrInvalidLengthWorkspace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWorkspace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWorkspace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWorkspace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkspace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWorkspace = fmt.Errorf("proto: unexpected end of group")
)
