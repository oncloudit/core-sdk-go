// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/auth/v1beta1/auth.proto

package auth

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/core-sdk-go/common/codec/types"
	_ "github.com/regen-network/cosmos-proto"
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

// BaseAccount defines a base account type. It contains all the necessary fields
// for basic account functionality. Any custom account type should extend this
// type for additional functionality (e.g. vesting).
type BaseAccount struct {
	Address       string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PubKey        *types.Any `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"public_key,omitempty" yaml:"public_key"`
	AccountNumber uint64     `protobuf:"varint,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty" yaml:"account_number"`
	Sequence      uint64     `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *BaseAccount) Reset()      { *m = BaseAccount{} }
func (*BaseAccount) ProtoMessage() {}
func (*BaseAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{0}
}
func (m *BaseAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAccount.Merge(m, src)
}
func (m *BaseAccount) XXX_Size() int {
	return m.Size()
}
func (m *BaseAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAccount.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAccount proto.InternalMessageInfo

// ModuleAccount defines an account for modules that holds coins on a pool.
type ModuleAccount struct {
	*BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty" yaml:"base_account"`
	Name         string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Permissions  []string `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (m *ModuleAccount) Reset()      { *m = ModuleAccount{} }
func (*ModuleAccount) ProtoMessage() {}
func (*ModuleAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{1}
}
func (m *ModuleAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ModuleAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ModuleAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ModuleAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleAccount.Merge(m, src)
}
func (m *ModuleAccount) XXX_Size() int {
	return m.Size()
}
func (m *ModuleAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleAccount proto.InternalMessageInfo

// Params defines the parameters for the auth module.
type Params struct {
	MaxMemoCharacters      uint64 `protobuf:"varint,1,opt,name=max_memo_characters,json=maxMemoCharacters,proto3" json:"max_memo_characters,omitempty" yaml:"max_memo_characters"`
	TxSigLimit             uint64 `protobuf:"varint,2,opt,name=tx_sig_limit,json=txSigLimit,proto3" json:"tx_sig_limit,omitempty" yaml:"tx_sig_limit"`
	TxSizeCostPerByte      uint64 `protobuf:"varint,3,opt,name=tx_size_cost_per_byte,json=txSizeCostPerByte,proto3" json:"tx_size_cost_per_byte,omitempty" yaml:"tx_size_cost_per_byte"`
	SigVerifyCostED25519   uint64 `protobuf:"varint,4,opt,name=sig_verify_cost_ed25519,json=sigVerifyCostEd25519,proto3" json:"sig_verify_cost_ed25519,omitempty" yaml:"sig_verify_cost_ed25519"`
	SigVerifyCostSecp256k1 uint64 `protobuf:"varint,5,opt,name=sig_verify_cost_secp256k1,json=sigVerifyCostSecp256k1,proto3" json:"sig_verify_cost_secp256k1,omitempty" yaml:"sig_verify_cost_secp256k1"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1f7e915d020d2d, []int{2}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxMemoCharacters() uint64 {
	if m != nil {
		return m.MaxMemoCharacters
	}
	return 0
}

func (m *Params) GetTxSigLimit() uint64 {
	if m != nil {
		return m.TxSigLimit
	}
	return 0
}

func (m *Params) GetTxSizeCostPerByte() uint64 {
	if m != nil {
		return m.TxSizeCostPerByte
	}
	return 0
}

func (m *Params) GetSigVerifyCostED25519() uint64 {
	if m != nil {
		return m.SigVerifyCostED25519
	}
	return 0
}

func (m *Params) GetSigVerifyCostSecp256k1() uint64 {
	if m != nil {
		return m.SigVerifyCostSecp256k1
	}
	return 0
}

func init() {
	proto.RegisterType((*BaseAccount)(nil), "cosmos.auth.v1beta1.BaseAccount")
	proto.RegisterType((*ModuleAccount)(nil), "cosmos.auth.v1beta1.ModuleAccount")
	proto.RegisterType((*Params)(nil), "cosmos.auth.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/auth/v1beta1/auth.proto", fileDescriptor_7e1f7e915d020d2d) }

var fileDescriptor_7e1f7e915d020d2d = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0x8f, 0x37, 0x59, 0xfe, 0x4c, 0x00, 0x09, 0x13, 0xc0, 0xc9, 0xae, 0x6c, 0xcb, 0xa7, 0xac,
	0xb4, 0xb1, 0x95, 0xac, 0x58, 0x89, 0x1c, 0x56, 0x8b, 0xd9, 0x3d, 0xa0, 0x5d, 0x10, 0x32, 0x52,
	0x0f, 0x55, 0x25, 0x77, 0xec, 0x0c, 0xc6, 0x22, 0xe3, 0x31, 0x9e, 0x31, 0x8a, 0xf9, 0x04, 0x3d,
	0xf6, 0xd8, 0x23, 0x1f, 0x82, 0x6f, 0xd0, 0x4b, 0x8f, 0x88, 0x53, 0x4f, 0x6e, 0x15, 0x2e, 0x55,
	0x8f, 0xb9, 0x57, 0xaa, 0x32, 0xe3, 0x84, 0x04, 0xa5, 0xb7, 0xf9, 0xfd, 0x79, 0xbf, 0xf7, 0xe6,
	0x4d, 0x62, 0xa0, 0xfa, 0x84, 0x62, 0x42, 0x2d, 0x98, 0xb2, 0x0b, 0xeb, 0xba, 0xed, 0x21, 0x06,
	0xdb, 0x1c, 0x98, 0x71, 0x42, 0x18, 0x91, 0xb7, 0x84, 0x6e, 0x72, 0xaa, 0xd0, 0x1b, 0x75, 0x41,
	0xba, 0xdc, 0x62, 0x15, 0x0e, 0x0e, 0x1a, 0xb5, 0x80, 0x04, 0x44, 0xf0, 0xe3, 0x53, 0xc1, 0xd6,
	0x03, 0x42, 0x82, 0x3e, 0xb2, 0x38, 0xf2, 0xd2, 0x73, 0x0b, 0x46, 0x99, 0x90, 0x8c, 0x6f, 0x12,
	0xa8, 0xda, 0x90, 0xa2, 0x03, 0xdf, 0x27, 0x69, 0xc4, 0x64, 0x05, 0x2c, 0xc3, 0x5e, 0x2f, 0x41,
	0x94, 0x2a, 0x92, 0x2e, 0x35, 0x57, 0x9d, 0x09, 0x94, 0x5f, 0x81, 0xe5, 0x38, 0xf5, 0xdc, 0x4b,
	0x94, 0x29, 0x3f, 0xe9, 0x52, 0xb3, 0xda, 0xa9, 0x99, 0x22, 0xd6, 0x9c, 0xc4, 0x9a, 0x07, 0x51,
	0x66, 0xb7, 0xbe, 0xe6, 0x5a, 0x2d, 0x4e, 0xbd, 0x7e, 0xe8, 0x8f, 0xbd, 0xbf, 0x13, 0x1c, 0x32,
	0x84, 0x63, 0x96, 0x8d, 0x72, 0x6d, 0x33, 0x83, 0xb8, 0xdf, 0x35, 0x9e, 0x54, 0xc3, 0x59, 0x8a,
	0x53, 0xef, 0x3f, 0x94, 0xc9, 0x7f, 0x83, 0x0d, 0x28, 0x46, 0x70, 0xa3, 0x14, 0x7b, 0x28, 0x51,
	0xca, 0xba, 0xd4, 0xac, 0xd8, 0xf5, 0x51, 0xae, 0x6d, 0x8b, 0xb2, 0x79, 0xdd, 0x70, 0xd6, 0x0b,
	0xe2, 0x84, 0x63, 0xb9, 0x01, 0x56, 0x28, 0xba, 0x4a, 0x51, 0xe4, 0x23, 0xa5, 0x32, 0xae, 0x75,
	0xa6, 0xb8, 0xab, 0xbc, 0xb9, 0xd5, 0x4a, 0xef, 0x6e, 0xb5, 0xd2, 0x97, 0x5b, 0xad, 0xf4, 0x70,
	0xd7, 0x5a, 0x29, 0xae, 0x7b, 0x64, 0xbc, 0x97, 0xc0, 0xfa, 0x31, 0xe9, 0xa5, 0xfd, 0xe9, 0x06,
	0x5e, 0x83, 0x35, 0x0f, 0x52, 0xe4, 0x16, 0xe9, 0x7c, 0x0d, 0xd5, 0x8e, 0x6e, 0x2e, 0x78, 0x09,
	0x73, 0x66, 0x73, 0xf6, 0x2f, 0xf7, 0xb9, 0x26, 0x8d, 0x72, 0x6d, 0x4b, 0x4c, 0x3b, 0x9b, 0x61,
	0x38, 0x55, 0x6f, 0x66, 0xc7, 0x32, 0xa8, 0x44, 0x10, 0x23, 0xbe, 0xc6, 0x55, 0x87, 0x9f, 0x65,
	0x1d, 0x54, 0x63, 0x94, 0xe0, 0x90, 0xd2, 0x90, 0x44, 0x54, 0x29, 0xeb, 0xe5, 0xe6, 0xaa, 0x33,
	0x4b, 0x75, 0x1b, 0x93, 0x3b, 0x3c, 0xdc, 0xb5, 0x36, 0xe6, 0x46, 0x3e, 0x32, 0x3e, 0x95, 0xc1,
	0xd2, 0x29, 0x4c, 0x20, 0xa6, 0xf2, 0x09, 0xd8, 0xc2, 0x70, 0xe0, 0x62, 0x84, 0x89, 0xeb, 0x5f,
	0xc0, 0x04, 0xfa, 0x0c, 0x25, 0xe2, 0x31, 0x2b, 0xb6, 0x3a, 0xca, 0xb5, 0x86, 0x98, 0x6f, 0x81,
	0xc9, 0x70, 0x36, 0x31, 0x1c, 0x1c, 0x23, 0x4c, 0x0e, 0xa7, 0x9c, 0xbc, 0x0f, 0xd6, 0xd8, 0xc0,
	0xa5, 0x61, 0xe0, 0xf6, 0x43, 0x1c, 0x32, 0x3e, 0x74, 0xc5, 0xde, 0x7d, 0xba, 0xe8, 0xac, 0x6a,
	0x38, 0x80, 0x0d, 0xce, 0xc2, 0xe0, 0xff, 0x31, 0x90, 0x1d, 0xb0, 0xcd, 0xc5, 0x1b, 0xe4, 0xfa,
	0x84, 0x32, 0x37, 0x46, 0x89, 0xeb, 0x65, 0x0c, 0x15, 0x4f, 0xab, 0x8f, 0x72, 0xed, 0xd7, 0x99,
	0x8c, 0xe7, 0x36, 0xc3, 0xd9, 0x1c, 0x87, 0xdd, 0xa0, 0x43, 0x42, 0xd9, 0x29, 0x4a, 0xec, 0x8c,
	0x21, 0xf9, 0x0a, 0xec, 0x8e, 0xbb, 0x5d, 0xa3, 0x24, 0x3c, 0xcf, 0x84, 0x1f, 0xf5, 0x3a, 0x7b,
	0x7b, 0xed, 0x7d, 0xf1, 0xe8, 0x76, 0x77, 0x98, 0x6b, 0xb5, 0xb3, 0x30, 0x78, 0xc1, 0x1d, 0xe3,
	0xd2, 0x7f, 0xff, 0xe1, 0xfa, 0x28, 0xd7, 0x54, 0xd1, 0xed, 0x07, 0x01, 0x86, 0x53, 0xa3, 0x73,
	0x75, 0x82, 0x96, 0x33, 0x50, 0x7f, 0x5e, 0x41, 0x91, 0x1f, 0x77, 0xf6, 0xfe, 0xbc, 0x6c, 0x2b,
	0x3f, 0xf3, 0xa6, 0x7f, 0x0d, 0x73, 0x6d, 0x67, 0xae, 0xe9, 0xd9, 0xc4, 0x31, 0xca, 0x35, 0x7d,
	0x71, 0xdb, 0x69, 0x88, 0xe1, 0xec, 0xd0, 0x85, 0xb5, 0xdd, 0x95, 0xe2, 0x37, 0x2b, 0xd9, 0x87,
	0x1f, 0x86, 0xaa, 0x74, 0x3f, 0x54, 0xa5, 0xcf, 0x43, 0x55, 0x7a, 0xfb, 0xa8, 0x96, 0xee, 0x1f,
	0xd5, 0xd2, 0xc7, 0x47, 0xb5, 0xf4, 0xf2, 0xb7, 0x20, 0x64, 0x17, 0xa9, 0x67, 0xfa, 0x04, 0x5b,
	0x61, 0x12, 0xd2, 0x08, 0x31, 0xcb, 0x27, 0x09, 0x6a, 0xd1, 0xde, 0x65, 0x2b, 0x20, 0x16, 0xcb,
	0x62, 0x24, 0x3e, 0x30, 0xde, 0x12, 0xff, 0xa7, 0xfe, 0xf1, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x36,
	0xc6, 0x83, 0xdf, 0x76, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.MaxMemoCharacters != that1.MaxMemoCharacters {
		return false
	}
	if this.TxSigLimit != that1.TxSigLimit {
		return false
	}
	if this.TxSizeCostPerByte != that1.TxSizeCostPerByte {
		return false
	}
	if this.SigVerifyCostED25519 != that1.SigVerifyCostED25519 {
		return false
	}
	if this.SigVerifyCostSecp256k1 != that1.SigVerifyCostSecp256k1 {
		return false
	}
	return true
}
func (m *BaseAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if m.AccountNumber != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x18
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ModuleAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ModuleAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ModuleAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Permissions) > 0 {
		for iNdEx := len(m.Permissions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Permissions[iNdEx])
			copy(dAtA[i:], m.Permissions[iNdEx])
			i = encodeVarintAuth(dAtA, i, uint64(len(m.Permissions[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SigVerifyCostSecp256k1 != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.SigVerifyCostSecp256k1))
		i--
		dAtA[i] = 0x28
	}
	if m.SigVerifyCostED25519 != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.SigVerifyCostED25519))
		i--
		dAtA[i] = 0x20
	}
	if m.TxSizeCostPerByte != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.TxSizeCostPerByte))
		i--
		dAtA[i] = 0x18
	}
	if m.TxSigLimit != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.TxSigLimit))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxMemoCharacters != 0 {
		i = encodeVarintAuth(dAtA, i, uint64(m.MaxMemoCharacters))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovAuth(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovAuth(uint64(m.Sequence))
	}
	return n
}

func (m *ModuleAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if len(m.Permissions) > 0 {
		for _, s := range m.Permissions {
			l = len(s)
			n += 1 + l + sovAuth(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxMemoCharacters != 0 {
		n += 1 + sovAuth(uint64(m.MaxMemoCharacters))
	}
	if m.TxSigLimit != 0 {
		n += 1 + sovAuth(uint64(m.TxSigLimit))
	}
	if m.TxSizeCostPerByte != 0 {
		n += 1 + sovAuth(uint64(m.TxSizeCostPerByte))
	}
	if m.SigVerifyCostED25519 != 0 {
		n += 1 + sovAuth(uint64(m.SigVerifyCostED25519))
	}
	if m.SigVerifyCostSecp256k1 != 0 {
		n += 1 + sovAuth(uint64(m.SigVerifyCostSecp256k1))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: BaseAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
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
func (m *ModuleAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: ModuleAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ModuleAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Permissions = append(m.Permissions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMemoCharacters", wireType)
			}
			m.MaxMemoCharacters = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMemoCharacters |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSigLimit", wireType)
			}
			m.TxSigLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxSigLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxSizeCostPerByte", wireType)
			}
			m.TxSizeCostPerByte = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxSizeCostPerByte |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostED25519", wireType)
			}
			m.SigVerifyCostED25519 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigVerifyCostED25519 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigVerifyCostSecp256k1", wireType)
			}
			m.SigVerifyCostSecp256k1 = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigVerifyCostSecp256k1 |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuth = fmt.Errorf("proto: unexpected end of group")
)
