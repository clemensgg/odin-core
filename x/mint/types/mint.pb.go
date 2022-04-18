// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mint/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type MintPool struct {
	// treasury pool
	TreasuryPool github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=treasury_pool,json=treasuryPool,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"treasury_pool" yaml:"treasury_pool"`
}

func (m *MintPool) Reset()         { *m = MintPool{} }
func (m *MintPool) String() string { return proto.CompactTextString(m) }
func (*MintPool) ProtoMessage()    {}
func (*MintPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b9fbb701b2a577, []int{0}
}
func (m *MintPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintPool.Merge(m, src)
}
func (m *MintPool) XXX_Size() int {
	return m.Size()
}
func (m *MintPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MintPool.DiscardUnknown(m)
}

var xxx_messageInfo_MintPool proto.InternalMessageInfo

func (m *MintPool) GetTreasuryPool() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TreasuryPool
	}
	return nil
}

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	// current annual expected provisions
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions" yaml:"annual_provisions"`
	// current mint volume
	CurrentMintVolume github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=current_mint_volume,json=currentMintVolume,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"current_mint_volume" yaml:"current_mint_volume"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1b9fbb701b2a577, []int{1}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetCurrentMintVolume() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CurrentMintVolume
	}
	return nil
}

func init() {
	proto.RegisterType((*MintPool)(nil), "mint.MintPool")
	proto.RegisterType((*Minter)(nil), "mint.Minter")
}

func init() { proto.RegisterFile("mint/mint.proto", fileDescriptor_e1b9fbb701b2a577) }

var fileDescriptor_e1b9fbb701b2a577 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x6f, 0xd4, 0x30,
	0x18, 0xc7, 0xcf, 0x6d, 0x55, 0xb5, 0x01, 0x04, 0x0d, 0x1d, 0xd2, 0x1b, 0xe2, 0x2a, 0x03, 0xba,
	0xe5, 0x62, 0x15, 0xb6, 0x1b, 0x43, 0x25, 0x5e, 0x54, 0x50, 0x75, 0x03, 0x03, 0x4b, 0xe4, 0xe4,
	0xdc, 0x60, 0x11, 0xfb, 0x89, 0x6c, 0xe7, 0x20, 0xdf, 0x80, 0x91, 0x11, 0x26, 0x3a, 0xf3, 0x49,
	0x3a, 0x76, 0x44, 0x0c, 0x01, 0xdd, 0x2d, 0xac, 0xf4, 0x13, 0x20, 0xc7, 0xe1, 0x1d, 0x89, 0xde,
	0x92, 0x38, 0xcf, 0xcb, 0xef, 0xff, 0xe4, 0xef, 0xc7, 0xbb, 0x2e, 0xb8, 0x34, 0xc4, 0x3e, 0xe2,
	0x4a, 0x81, 0x01, 0x7f, 0xc3, 0x9e, 0x87, 0xbb, 0x05, 0x14, 0xd0, 0x05, 0x88, 0x3d, 0xb9, 0xdc,
	0x10, 0x17, 0x00, 0x45, 0xc9, 0x48, 0xf7, 0x95, 0xd5, 0x27, 0xc4, 0x70, 0xc1, 0xb4, 0xa1, 0xa2,
	0xea, 0x0b, 0xf6, 0xfe, 0x2c, 0xa0, 0xb2, 0xe9, 0x53, 0x61, 0x0e, 0x5a, 0x80, 0x26, 0x19, 0xd5,
	0x8c, 0xcc, 0x0f, 0x32, 0x66, 0xe8, 0x01, 0xc9, 0x81, 0x4b, 0x97, 0x8f, 0xde, 0x21, 0x6f, 0xeb,
	0x11, 0x97, 0xe6, 0x18, 0xa0, 0xf4, 0x5f, 0x21, 0xef, 0x9a, 0x51, 0x8c, 0xea, 0x5a, 0x35, 0x69,
	0x05, 0x50, 0x06, 0x68, 0x7f, 0x7d, 0x74, 0xe5, 0xf6, 0x5e, 0xec, 0x28, 0xb1, 0xa5, 0xc4, 0x3d,
	0x25, 0xbe, 0x0b, 0x5c, 0x26, 0xf7, 0xcf, 0x5a, 0x3c, 0xb8, 0x68, 0xf1, 0x6e, 0x43, 0x45, 0x39,
	0x89, 0x7e, 0xeb, 0x8e, 0xde, 0x7f, 0xc2, 0xa3, 0x82, 0x9b, 0x67, 0x75, 0x16, 0xe7, 0x20, 0x48,
	0x3f, 0x8a, 0x7b, 0x8d, 0xf5, 0xec, 0x39, 0x31, 0x4d, 0xc5, 0x74, 0x07, 0xd2, 0xd3, 0xab, 0xdf,
	0x7b, 0xed, 0x28, 0x93, 0xad, 0x37, 0xa7, 0x18, 0x7d, 0x39, 0xc5, 0x28, 0xfa, 0xba, 0xe6, 0x6d,
	0xda, 0x09, 0x99, 0xf2, 0x8f, 0xbc, 0x6d, 0x2e, 0x4f, 0x4a, 0x6a, 0x38, 0xc8, 0x00, 0xed, 0xa3,
	0xd1, 0x76, 0x12, 0x5b, 0xfd, 0x8f, 0x2d, 0xbe, 0x75, 0x09, 0x9d, 0x43, 0x96, 0x4f, 0x7f, 0x02,
	0xfc, 0x17, 0xde, 0x0e, 0x95, 0xb2, 0xa6, 0x65, 0x5a, 0x29, 0x98, 0x73, 0xcd, 0x41, 0xea, 0x60,
	0xad, 0xa3, 0x3e, 0x5c, 0x8d, 0x7a, 0xd1, 0xe2, 0xc0, 0xfd, 0xff, 0x5f, 0xc0, 0x68, 0x7a, 0xc3,
	0xc5, 0x8e, 0x7f, 0x84, 0xfc, 0xb7, 0xc8, 0xbb, 0x99, 0xd7, 0x4a, 0x31, 0x69, 0x52, 0x7b, 0xed,
	0xe9, 0x1c, 0xca, 0x5a, 0xb0, 0x60, 0xfd, 0x7f, 0x66, 0x3f, 0xee, 0xcd, 0x1e, 0x3a, 0xb1, 0x7f,
	0x30, 0x56, 0xb3, 0x7c, 0xa7, 0x27, 0x58, 0x7b, 0x9f, 0x74, 0xfd, 0x93, 0x0d, 0xeb, 0x79, 0xf2,
	0xe0, 0x6c, 0x11, 0xa2, 0xf3, 0x45, 0x88, 0x3e, 0x2f, 0x42, 0xf4, 0x7a, 0x19, 0x0e, 0xce, 0x97,
	0xe1, 0xe0, 0xc3, 0x32, 0x1c, 0x3c, 0x25, 0xbf, 0xc0, 0xef, 0x31, 0x38, 0x4c, 0xc6, 0x47, 0x5c,
	0x70, 0xc3, 0x66, 0x04, 0x66, 0x5c, 0x8e, 0x73, 0x50, 0x8c, 0xbc, 0xec, 0x16, 0xdb, 0x29, 0x65,
	0x9b, 0xdd, 0x9e, 0xdd, 0xf9, 0x16, 0x00, 0x00, 0xff, 0xff, 0x03, 0x0e, 0x7c, 0x51, 0xf2, 0x02,
	0x00, 0x00,
}

func (this *MintPool) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MintPool)
	if !ok {
		that2, ok := that.(MintPool)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.TreasuryPool) != len(that1.TreasuryPool) {
		return false
	}
	for i := range this.TreasuryPool {
		if !this.TreasuryPool[i].Equal(&that1.TreasuryPool[i]) {
			return false
		}
	}
	return true
}
func (this *Minter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Minter)
	if !ok {
		that2, ok := that.(Minter)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Inflation.Equal(that1.Inflation) {
		return false
	}
	if !this.AnnualProvisions.Equal(that1.AnnualProvisions) {
		return false
	}
	if len(this.CurrentMintVolume) != len(that1.CurrentMintVolume) {
		return false
	}
	for i := range this.CurrentMintVolume {
		if !this.CurrentMintVolume[i].Equal(&that1.CurrentMintVolume[i]) {
			return false
		}
	}
	return true
}
func (m *MintPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TreasuryPool) > 0 {
		for iNdEx := len(m.TreasuryPool) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TreasuryPool[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CurrentMintVolume) > 0 {
		for iNdEx := len(m.CurrentMintVolume) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CurrentMintVolume[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMint(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MintPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TreasuryPool) > 0 {
		for _, e := range m.TreasuryPool {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.AnnualProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	if len(m.CurrentMintVolume) > 0 {
		for _, e := range m.CurrentMintVolume {
			l = e.Size()
			n += 1 + l + sovMint(uint64(l))
		}
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MintPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: MintPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreasuryPool", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreasuryPool = append(m.TreasuryPool, types.Coin{})
			if err := m.TreasuryPool[len(m.TreasuryPool)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentMintVolume", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentMintVolume = append(m.CurrentMintVolume, types.Coin{})
			if err := m.CurrentMintVolume[len(m.CurrentMintVolume)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
