// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmwasm/wasm/v1/genesis.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState - genesis state of x/wasm
type GenesisState struct {
	Params    Params     `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Codes     []Code     `protobuf:"bytes,2,rep,name=codes,proto3" json:"codes,omitempty"`
	Contracts []Contract `protobuf:"bytes,3,rep,name=contracts,proto3" json:"contracts,omitempty"`
	Sequences []Sequence `protobuf:"bytes,4,rep,name=sequences,proto3" json:"sequences,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ab3f539b23472a6, []int{0}
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

func (m *GenesisState) GetCodes() []Code {
	if m != nil {
		return m.Codes
	}
	return nil
}

func (m *GenesisState) GetContracts() []Contract {
	if m != nil {
		return m.Contracts
	}
	return nil
}

func (m *GenesisState) GetSequences() []Sequence {
	if m != nil {
		return m.Sequences
	}
	return nil
}

// Code struct encompasses CodeInfo and CodeBytes
type Code struct {
	CodeID    uint64   `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	CodeInfo  CodeInfo `protobuf:"bytes,2,opt,name=code_info,json=codeInfo,proto3" json:"code_info"`
	CodeBytes []byte   `protobuf:"bytes,3,opt,name=code_bytes,json=codeBytes,proto3" json:"code_bytes,omitempty"`
	// Pinned to wasmvm cache
	Pinned bool `protobuf:"varint,4,opt,name=pinned,proto3" json:"pinned,omitempty"`
}

func (m *Code) Reset()         { *m = Code{} }
func (m *Code) String() string { return proto.CompactTextString(m) }
func (*Code) ProtoMessage()    {}
func (*Code) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ab3f539b23472a6, []int{1}
}

func (m *Code) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Code) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Code.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Code) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Code.Merge(m, src)
}

func (m *Code) XXX_Size() int {
	return m.Size()
}

func (m *Code) XXX_DiscardUnknown() {
	xxx_messageInfo_Code.DiscardUnknown(m)
}

var xxx_messageInfo_Code proto.InternalMessageInfo

func (m *Code) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *Code) GetCodeInfo() CodeInfo {
	if m != nil {
		return m.CodeInfo
	}
	return CodeInfo{}
}

func (m *Code) GetCodeBytes() []byte {
	if m != nil {
		return m.CodeBytes
	}
	return nil
}

func (m *Code) GetPinned() bool {
	if m != nil {
		return m.Pinned
	}
	return false
}

// Contract struct encompasses ContractAddress, ContractInfo, and ContractState
type Contract struct {
	ContractAddress     string                     `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ContractInfo        ContractInfo               `protobuf:"bytes,2,opt,name=contract_info,json=contractInfo,proto3" json:"contract_info"`
	ContractState       []Model                    `protobuf:"bytes,3,rep,name=contract_state,json=contractState,proto3" json:"contract_state"`
	ContractCodeHistory []ContractCodeHistoryEntry `protobuf:"bytes,4,rep,name=contract_code_history,json=contractCodeHistory,proto3" json:"contract_code_history"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ab3f539b23472a6, []int{2}
}

func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}

func (m *Contract) XXX_Size() int {
	return m.Size()
}

func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Contract) GetContractInfo() ContractInfo {
	if m != nil {
		return m.ContractInfo
	}
	return ContractInfo{}
}

func (m *Contract) GetContractState() []Model {
	if m != nil {
		return m.ContractState
	}
	return nil
}

func (m *Contract) GetContractCodeHistory() []ContractCodeHistoryEntry {
	if m != nil {
		return m.ContractCodeHistory
	}
	return nil
}

// Sequence key and value of an id generation counter
type Sequence struct {
	IDKey []byte `protobuf:"bytes,1,opt,name=id_key,json=idKey,proto3" json:"id_key,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Sequence) Reset()         { *m = Sequence{} }
func (m *Sequence) String() string { return proto.CompactTextString(m) }
func (*Sequence) ProtoMessage()    {}
func (*Sequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ab3f539b23472a6, []int{3}
}

func (m *Sequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Sequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Sequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequence.Merge(m, src)
}

func (m *Sequence) XXX_Size() int {
	return m.Size()
}

func (m *Sequence) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequence.DiscardUnknown(m)
}

var xxx_messageInfo_Sequence proto.InternalMessageInfo

func (m *Sequence) GetIDKey() []byte {
	if m != nil {
		return m.IDKey
	}
	return nil
}

func (m *Sequence) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmwasm.wasm.v1.GenesisState")
	proto.RegisterType((*Code)(nil), "cosmwasm.wasm.v1.Code")
	proto.RegisterType((*Contract)(nil), "cosmwasm.wasm.v1.Contract")
	proto.RegisterType((*Sequence)(nil), "cosmwasm.wasm.v1.Sequence")
}

func init() { proto.RegisterFile("cosmwasm/wasm/v1/genesis.proto", fileDescriptor_2ab3f539b23472a6) }

var fileDescriptor_2ab3f539b23472a6 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0xd4, 0x31, 0xe9, 0x36, 0xd0, 0x6a, 0x29, 0xad, 0x15, 0xc0, 0x89, 0x82, 0x84,
	0x02, 0x42, 0xb6, 0x5a, 0x24, 0x6e, 0x48, 0xe0, 0xa6, 0x02, 0xab, 0x42, 0x42, 0xae, 0x10, 0x12,
	0x97, 0xc8, 0xf1, 0x6e, 0x53, 0x8b, 0xda, 0x6b, 0xbc, 0x9b, 0x80, 0xdf, 0x82, 0x57, 0xe0, 0xce,
	0x1b, 0xf0, 0x02, 0x3d, 0xf6, 0xc8, 0x29, 0x42, 0xce, 0x8d, 0xa7, 0x40, 0xfb, 0xc7, 0xae, 0x45,
	0x9a, 0x8b, 0xe5, 0x9d, 0xf9, 0xe6, 0xe7, 0xf5, 0x37, 0x33, 0xc0, 0x0a, 0x09, 0x8d, 0xbf, 0x06,
	0x34, 0x76, 0xc4, 0x63, 0x7e, 0xe0, 0x4c, 0x71, 0x82, 0x69, 0x44, 0xed, 0x34, 0x23, 0x8c, 0xc0,
	0x9d, 0x32, 0x6f, 0x8b, 0xc7, 0xfc, 0xa0, 0xbb, 0x3b, 0x25, 0x53, 0x22, 0x92, 0x0e, 0x7f, 0x93,
	0xba, 0xee, 0x83, 0x15, 0x0e, 0xcb, 0x53, 0xac, 0x28, 0x83, 0x5f, 0x4d, 0xd0, 0x79, 0x23, 0xb9,
	0xa7, 0x2c, 0x60, 0x18, 0xbe, 0x00, 0x46, 0x1a, 0x64, 0x41, 0x4c, 0x4d, 0xad, 0xaf, 0x0d, 0xb7,
	0x0e, 0x4d, 0xfb, 0xff, 0xef, 0xd8, 0xef, 0x45, 0xde, 0xd5, 0x2f, 0x17, 0xbd, 0x86, 0xaf, 0xd4,
	0xf0, 0x18, 0xb4, 0x42, 0x82, 0x30, 0x35, 0x9b, 0xfd, 0x8d, 0xe1, 0xd6, 0xe1, 0xde, 0x6a, 0xd9,
	0x11, 0x41, 0xd8, 0xdd, 0xe7, 0x45, 0x7f, 0x17, 0xbd, 0x6d, 0x21, 0x7e, 0x46, 0xe2, 0x88, 0xe1,
	0x38, 0x65, 0xb9, 0x2f, 0xab, 0xe1, 0x07, 0xb0, 0x19, 0x92, 0x84, 0x65, 0x41, 0xc8, 0xa8, 0xb9,
	0x21, 0x50, 0xdd, 0x9b, 0x50, 0x52, 0xe2, 0xde, 0x57, 0xb8, 0xbb, 0x55, 0x51, 0x0d, 0x79, 0x4d,
	0xe2, 0x58, 0x8a, 0xbf, 0xcc, 0x70, 0x12, 0x62, 0x6a, 0xea, 0xeb, 0xb0, 0xa7, 0x4a, 0x72, 0x8d,
	0xad, 0x8a, 0xea, 0xd8, 0x2a, 0x38, 0xf8, 0xa1, 0x01, 0x9d, 0xff, 0x16, 0x7c, 0x04, 0x6e, 0xf1,
	0xfb, 0x8f, 0x23, 0x24, 0x6c, 0xd3, 0x5d, 0x50, 0x2c, 0x7a, 0x06, 0x4f, 0x79, 0x23, 0xdf, 0xe0,
	0x29, 0x0f, 0xc1, 0x97, 0xfc, 0xdf, 0xb8, 0x28, 0x39, 0x23, 0x66, 0x53, 0xb8, 0xdb, 0xbd, 0xd9,
	0x26, 0x2f, 0x39, 0x23, 0xca, 0xdf, 0x76, 0xa8, 0xce, 0xf0, 0x21, 0x00, 0xa2, 0x7c, 0x92, 0x33,
	0xcc, 0xbd, 0xd1, 0x86, 0x1d, 0x5f, 0x00, 0x5d, 0x1e, 0x80, 0x7b, 0xc0, 0x48, 0xa3, 0x24, 0xc1,
	0xc8, 0xd4, 0xfb, 0xda, 0xb0, 0xed, 0xab, 0xd3, 0xe0, 0x67, 0x13, 0xb4, 0x4b, 0xbf, 0xe0, 0x13,
	0xb0, 0x53, 0x9a, 0x32, 0x0e, 0x10, 0xca, 0x30, 0x95, 0x7d, 0xde, 0xf4, 0xb7, 0xcb, 0xf8, 0x6b,
	0x19, 0x86, 0x1e, 0xb8, 0x5d, 0x49, 0x6b, 0x37, 0xb6, 0xd6, 0x77, 0xa3, 0x76, 0xeb, 0x4e, 0x58,
	0x8b, 0xc1, 0x11, 0xb8, 0x53, 0xa1, 0x28, 0x9f, 0x32, 0xd5, 0xd9, 0xfd, 0x55, 0xd6, 0x3b, 0x82,
	0xf0, 0x85, 0x82, 0x54, 0xdf, 0x97, 0x93, 0x89, 0xc0, 0xbd, 0x8a, 0x22, 0x8c, 0x38, 0x8f, 0x28,
	0x23, 0x59, 0xae, 0xfa, 0xf9, 0x74, 0xfd, 0xc5, 0xb8, 0xa5, 0x6f, 0xa5, 0xf8, 0x38, 0x61, 0x59,
	0xae, 0xf8, 0xd5, 0xd0, 0xd4, 0xf2, 0x03, 0x17, 0xb4, 0xcb, 0x31, 0x80, 0x7d, 0x60, 0x44, 0x68,
	0xfc, 0x19, 0xe7, 0xc2, 0xa3, 0x8e, 0xbb, 0x59, 0x2c, 0x7a, 0x2d, 0x6f, 0x74, 0x82, 0x73, 0xbf,
	0x15, 0xa1, 0x13, 0x9c, 0xc3, 0x5d, 0xd0, 0x9a, 0x07, 0x17, 0x33, 0x2c, 0xcc, 0xd1, 0x7d, 0x79,
	0x70, 0x5f, 0x5d, 0x16, 0x96, 0x76, 0x55, 0x58, 0xda, 0x9f, 0xc2, 0xd2, 0xbe, 0x2f, 0xad, 0xc6,
	0xd5, 0xd2, 0x6a, 0xfc, 0x5e, 0x5a, 0x8d, 0x4f, 0x8f, 0xa7, 0x11, 0x3b, 0x9f, 0x4d, 0xec, 0x90,
	0xc4, 0xce, 0x11, 0xa1, 0xf1, 0xc7, 0x72, 0x2f, 0x91, 0xf3, 0x4d, 0xee, 0xa7, 0x58, 0xce, 0x89,
	0x21, 0xb6, 0xf3, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0x33, 0x7b, 0x9f, 0x05, 0x04,
	0x00, 0x00,
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
	if len(m.Sequences) > 0 {
		for iNdEx := len(m.Sequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Contracts) > 0 {
		for iNdEx := len(m.Contracts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Contracts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Codes) > 0 {
		for iNdEx := len(m.Codes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Codes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *Code) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Code) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Code) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pinned {
		i--
		if m.Pinned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.CodeBytes) > 0 {
		i -= len(m.CodeBytes)
		copy(dAtA[i:], m.CodeBytes)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CodeBytes)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.CodeInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.CodeID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractCodeHistory) > 0 {
		for iNdEx := len(m.ContractCodeHistory) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractCodeHistory[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ContractState) > 0 {
		for iNdEx := len(m.ContractState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
		size, err := m.ContractInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Sequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.IDKey) > 0 {
		i -= len(m.IDKey)
		copy(dAtA[i:], m.IDKey)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IDKey)))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.Codes) > 0 {
		for _, e := range m.Codes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Contracts) > 0 {
		for _, e := range m.Contracts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Sequences) > 0 {
		for _, e := range m.Sequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Code) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovGenesis(uint64(m.CodeID))
	}
	l = m.CodeInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.CodeBytes)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Pinned {
		n += 2
	}
	return n
}

func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.ContractInfo.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ContractState) > 0 {
		for _, e := range m.ContractState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ContractCodeHistory) > 0 {
		for _, e := range m.ContractCodeHistory {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Sequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IDKey)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovGenesis(uint64(m.Value))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Codes", wireType)
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
			m.Codes = append(m.Codes, Code{})
			if err := m.Codes[len(m.Codes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contracts", wireType)
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
			m.Contracts = append(m.Contracts, Contract{})
			if err := m.Contracts[len(m.Contracts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequences", wireType)
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
			m.Sequences = append(m.Sequences, Sequence{})
			if err := m.Sequences[len(m.Sequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func (m *Code) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Code: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Code: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeInfo", wireType)
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
			if err := m.CodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeBytes = append(m.CodeBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeBytes == nil {
				m.CodeBytes = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pinned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Pinned = bool(v != 0)
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

func (m *Contract) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractInfo", wireType)
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
			if err := m.ContractInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractState", wireType)
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
			m.ContractState = append(m.ContractState, Model{})
			if err := m.ContractState[len(m.ContractState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractCodeHistory", wireType)
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
			m.ContractCodeHistory = append(m.ContractCodeHistory, ContractCodeHistoryEntry{})
			if err := m.ContractCodeHistory[len(m.ContractCodeHistory)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

func (m *Sequence) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Sequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDKey = append(m.IDKey[:0], dAtA[iNdEx:postIndex]...)
			if m.IDKey == nil {
				m.IDKey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
