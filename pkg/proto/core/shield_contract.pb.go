// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/contract/shield_contract.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AuthenticationPath struct {
	Value                []bool   `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationPath) Reset()         { *m = AuthenticationPath{} }
func (m *AuthenticationPath) String() string { return proto.CompactTextString(m) }
func (*AuthenticationPath) ProtoMessage()    {}
func (*AuthenticationPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{0}
}

func (m *AuthenticationPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationPath.Unmarshal(m, b)
}
func (m *AuthenticationPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationPath.Marshal(b, m, deterministic)
}
func (m *AuthenticationPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationPath.Merge(m, src)
}
func (m *AuthenticationPath) XXX_Size() int {
	return xxx_messageInfo_AuthenticationPath.Size(m)
}
func (m *AuthenticationPath) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationPath.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationPath proto.InternalMessageInfo

func (m *AuthenticationPath) GetValue() []bool {
	if m != nil {
		return m.Value
	}
	return nil
}

type MerklePath struct {
	AuthenticationPaths  []*AuthenticationPath `protobuf:"bytes,1,rep,name=authentication_paths,json=authenticationPaths,proto3" json:"authentication_paths,omitempty"`
	Index                []bool                `protobuf:"varint,2,rep,packed,name=index,proto3" json:"index,omitempty"`
	Rt                   []byte                `protobuf:"bytes,3,opt,name=rt,proto3" json:"rt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MerklePath) Reset()         { *m = MerklePath{} }
func (m *MerklePath) String() string { return proto.CompactTextString(m) }
func (*MerklePath) ProtoMessage()    {}
func (*MerklePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{1}
}

func (m *MerklePath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerklePath.Unmarshal(m, b)
}
func (m *MerklePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerklePath.Marshal(b, m, deterministic)
}
func (m *MerklePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerklePath.Merge(m, src)
}
func (m *MerklePath) XXX_Size() int {
	return xxx_messageInfo_MerklePath.Size(m)
}
func (m *MerklePath) XXX_DiscardUnknown() {
	xxx_messageInfo_MerklePath.DiscardUnknown(m)
}

var xxx_messageInfo_MerklePath proto.InternalMessageInfo

func (m *MerklePath) GetAuthenticationPaths() []*AuthenticationPath {
	if m != nil {
		return m.AuthenticationPaths
	}
	return nil
}

func (m *MerklePath) GetIndex() []bool {
	if m != nil {
		return m.Index
	}
	return nil
}

func (m *MerklePath) GetRt() []byte {
	if m != nil {
		return m.Rt
	}
	return nil
}

type OutputPoint struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputPoint) Reset()         { *m = OutputPoint{} }
func (m *OutputPoint) String() string { return proto.CompactTextString(m) }
func (*OutputPoint) ProtoMessage()    {}
func (*OutputPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{2}
}

func (m *OutputPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputPoint.Unmarshal(m, b)
}
func (m *OutputPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputPoint.Marshal(b, m, deterministic)
}
func (m *OutputPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputPoint.Merge(m, src)
}
func (m *OutputPoint) XXX_Size() int {
	return xxx_messageInfo_OutputPoint.Size(m)
}
func (m *OutputPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputPoint.DiscardUnknown(m)
}

var xxx_messageInfo_OutputPoint proto.InternalMessageInfo

func (m *OutputPoint) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *OutputPoint) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type OutputPointInfo struct {
	OutPoints            []*OutputPoint `protobuf:"bytes,1,rep,name=out_points,json=outPoints,proto3" json:"out_points,omitempty"`
	BlockNum             int32          `protobuf:"varint,2,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OutputPointInfo) Reset()         { *m = OutputPointInfo{} }
func (m *OutputPointInfo) String() string { return proto.CompactTextString(m) }
func (*OutputPointInfo) ProtoMessage()    {}
func (*OutputPointInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{3}
}

func (m *OutputPointInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputPointInfo.Unmarshal(m, b)
}
func (m *OutputPointInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputPointInfo.Marshal(b, m, deterministic)
}
func (m *OutputPointInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputPointInfo.Merge(m, src)
}
func (m *OutputPointInfo) XXX_Size() int {
	return xxx_messageInfo_OutputPointInfo.Size(m)
}
func (m *OutputPointInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputPointInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OutputPointInfo proto.InternalMessageInfo

func (m *OutputPointInfo) GetOutPoints() []*OutputPoint {
	if m != nil {
		return m.OutPoints
	}
	return nil
}

func (m *OutputPointInfo) GetBlockNum() int32 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

type PedersenHash struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PedersenHash) Reset()         { *m = PedersenHash{} }
func (m *PedersenHash) String() string { return proto.CompactTextString(m) }
func (*PedersenHash) ProtoMessage()    {}
func (*PedersenHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{4}
}

func (m *PedersenHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PedersenHash.Unmarshal(m, b)
}
func (m *PedersenHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PedersenHash.Marshal(b, m, deterministic)
}
func (m *PedersenHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PedersenHash.Merge(m, src)
}
func (m *PedersenHash) XXX_Size() int {
	return xxx_messageInfo_PedersenHash.Size(m)
}
func (m *PedersenHash) XXX_DiscardUnknown() {
	xxx_messageInfo_PedersenHash.DiscardUnknown(m)
}

var xxx_messageInfo_PedersenHash proto.InternalMessageInfo

func (m *PedersenHash) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type IncrementalMerkleTree struct {
	Left                 *PedersenHash   `protobuf:"bytes,1,opt,name=left,proto3" json:"left,omitempty"`
	Right                *PedersenHash   `protobuf:"bytes,2,opt,name=right,proto3" json:"right,omitempty"`
	Parents              []*PedersenHash `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IncrementalMerkleTree) Reset()         { *m = IncrementalMerkleTree{} }
func (m *IncrementalMerkleTree) String() string { return proto.CompactTextString(m) }
func (*IncrementalMerkleTree) ProtoMessage()    {}
func (*IncrementalMerkleTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{5}
}

func (m *IncrementalMerkleTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementalMerkleTree.Unmarshal(m, b)
}
func (m *IncrementalMerkleTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementalMerkleTree.Marshal(b, m, deterministic)
}
func (m *IncrementalMerkleTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementalMerkleTree.Merge(m, src)
}
func (m *IncrementalMerkleTree) XXX_Size() int {
	return xxx_messageInfo_IncrementalMerkleTree.Size(m)
}
func (m *IncrementalMerkleTree) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementalMerkleTree.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementalMerkleTree proto.InternalMessageInfo

func (m *IncrementalMerkleTree) GetLeft() *PedersenHash {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *IncrementalMerkleTree) GetRight() *PedersenHash {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *IncrementalMerkleTree) GetParents() []*PedersenHash {
	if m != nil {
		return m.Parents
	}
	return nil
}

type IncrementalMerkleVoucher struct {
	Tree                 *IncrementalMerkleTree `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
	Filled               []*PedersenHash        `protobuf:"bytes,2,rep,name=filled,proto3" json:"filled,omitempty"`
	Cursor               *IncrementalMerkleTree `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	CursorDepth          int64                  `protobuf:"varint,4,opt,name=cursor_depth,json=cursorDepth,proto3" json:"cursor_depth,omitempty"`
	Rt                   []byte                 `protobuf:"bytes,5,opt,name=rt,proto3" json:"rt,omitempty"`
	OutputPoint          *OutputPoint           `protobuf:"bytes,10,opt,name=output_point,json=outputPoint,proto3" json:"output_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IncrementalMerkleVoucher) Reset()         { *m = IncrementalMerkleVoucher{} }
func (m *IncrementalMerkleVoucher) String() string { return proto.CompactTextString(m) }
func (*IncrementalMerkleVoucher) ProtoMessage()    {}
func (*IncrementalMerkleVoucher) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{6}
}

func (m *IncrementalMerkleVoucher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementalMerkleVoucher.Unmarshal(m, b)
}
func (m *IncrementalMerkleVoucher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementalMerkleVoucher.Marshal(b, m, deterministic)
}
func (m *IncrementalMerkleVoucher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementalMerkleVoucher.Merge(m, src)
}
func (m *IncrementalMerkleVoucher) XXX_Size() int {
	return xxx_messageInfo_IncrementalMerkleVoucher.Size(m)
}
func (m *IncrementalMerkleVoucher) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementalMerkleVoucher.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementalMerkleVoucher proto.InternalMessageInfo

func (m *IncrementalMerkleVoucher) GetTree() *IncrementalMerkleTree {
	if m != nil {
		return m.Tree
	}
	return nil
}

func (m *IncrementalMerkleVoucher) GetFilled() []*PedersenHash {
	if m != nil {
		return m.Filled
	}
	return nil
}

func (m *IncrementalMerkleVoucher) GetCursor() *IncrementalMerkleTree {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *IncrementalMerkleVoucher) GetCursorDepth() int64 {
	if m != nil {
		return m.CursorDepth
	}
	return 0
}

func (m *IncrementalMerkleVoucher) GetRt() []byte {
	if m != nil {
		return m.Rt
	}
	return nil
}

func (m *IncrementalMerkleVoucher) GetOutputPoint() *OutputPoint {
	if m != nil {
		return m.OutputPoint
	}
	return nil
}

type IncrementalMerkleVoucherInfo struct {
	Vouchers             []*IncrementalMerkleVoucher `protobuf:"bytes,1,rep,name=vouchers,proto3" json:"vouchers,omitempty"`
	Paths                [][]byte                    `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *IncrementalMerkleVoucherInfo) Reset()         { *m = IncrementalMerkleVoucherInfo{} }
func (m *IncrementalMerkleVoucherInfo) String() string { return proto.CompactTextString(m) }
func (*IncrementalMerkleVoucherInfo) ProtoMessage()    {}
func (*IncrementalMerkleVoucherInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{7}
}

func (m *IncrementalMerkleVoucherInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrementalMerkleVoucherInfo.Unmarshal(m, b)
}
func (m *IncrementalMerkleVoucherInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrementalMerkleVoucherInfo.Marshal(b, m, deterministic)
}
func (m *IncrementalMerkleVoucherInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrementalMerkleVoucherInfo.Merge(m, src)
}
func (m *IncrementalMerkleVoucherInfo) XXX_Size() int {
	return xxx_messageInfo_IncrementalMerkleVoucherInfo.Size(m)
}
func (m *IncrementalMerkleVoucherInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrementalMerkleVoucherInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IncrementalMerkleVoucherInfo proto.InternalMessageInfo

func (m *IncrementalMerkleVoucherInfo) GetVouchers() []*IncrementalMerkleVoucher {
	if m != nil {
		return m.Vouchers
	}
	return nil
}

func (m *IncrementalMerkleVoucherInfo) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

type SpendDescription struct {
	ValueCommitment         []byte   `protobuf:"bytes,1,opt,name=value_commitment,json=valueCommitment,proto3" json:"value_commitment,omitempty"`
	Anchor                  []byte   `protobuf:"bytes,2,opt,name=anchor,proto3" json:"anchor,omitempty"`
	Nullifier               []byte   `protobuf:"bytes,3,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	Rk                      []byte   `protobuf:"bytes,4,opt,name=rk,proto3" json:"rk,omitempty"`
	Zkproof                 []byte   `protobuf:"bytes,5,opt,name=zkproof,proto3" json:"zkproof,omitempty"`
	SpendAuthoritySignature []byte   `protobuf:"bytes,6,opt,name=spend_authority_signature,json=spendAuthoritySignature,proto3" json:"spend_authority_signature,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *SpendDescription) Reset()         { *m = SpendDescription{} }
func (m *SpendDescription) String() string { return proto.CompactTextString(m) }
func (*SpendDescription) ProtoMessage()    {}
func (*SpendDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{8}
}

func (m *SpendDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpendDescription.Unmarshal(m, b)
}
func (m *SpendDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpendDescription.Marshal(b, m, deterministic)
}
func (m *SpendDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpendDescription.Merge(m, src)
}
func (m *SpendDescription) XXX_Size() int {
	return xxx_messageInfo_SpendDescription.Size(m)
}
func (m *SpendDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SpendDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SpendDescription proto.InternalMessageInfo

func (m *SpendDescription) GetValueCommitment() []byte {
	if m != nil {
		return m.ValueCommitment
	}
	return nil
}

func (m *SpendDescription) GetAnchor() []byte {
	if m != nil {
		return m.Anchor
	}
	return nil
}

func (m *SpendDescription) GetNullifier() []byte {
	if m != nil {
		return m.Nullifier
	}
	return nil
}

func (m *SpendDescription) GetRk() []byte {
	if m != nil {
		return m.Rk
	}
	return nil
}

func (m *SpendDescription) GetZkproof() []byte {
	if m != nil {
		return m.Zkproof
	}
	return nil
}

func (m *SpendDescription) GetSpendAuthoritySignature() []byte {
	if m != nil {
		return m.SpendAuthoritySignature
	}
	return nil
}

type ReceiveDescription struct {
	ValueCommitment      []byte   `protobuf:"bytes,1,opt,name=value_commitment,json=valueCommitment,proto3" json:"value_commitment,omitempty"`
	NoteCommitment       []byte   `protobuf:"bytes,2,opt,name=note_commitment,json=noteCommitment,proto3" json:"note_commitment,omitempty"`
	Epk                  []byte   `protobuf:"bytes,3,opt,name=epk,proto3" json:"epk,omitempty"`
	CEnc                 []byte   `protobuf:"bytes,4,opt,name=c_enc,json=cEnc,proto3" json:"c_enc,omitempty"`
	COut                 []byte   `protobuf:"bytes,5,opt,name=c_out,json=cOut,proto3" json:"c_out,omitempty"`
	Zkproof              []byte   `protobuf:"bytes,6,opt,name=zkproof,proto3" json:"zkproof,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveDescription) Reset()         { *m = ReceiveDescription{} }
func (m *ReceiveDescription) String() string { return proto.CompactTextString(m) }
func (*ReceiveDescription) ProtoMessage()    {}
func (*ReceiveDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{9}
}

func (m *ReceiveDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveDescription.Unmarshal(m, b)
}
func (m *ReceiveDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveDescription.Marshal(b, m, deterministic)
}
func (m *ReceiveDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveDescription.Merge(m, src)
}
func (m *ReceiveDescription) XXX_Size() int {
	return xxx_messageInfo_ReceiveDescription.Size(m)
}
func (m *ReceiveDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveDescription proto.InternalMessageInfo

func (m *ReceiveDescription) GetValueCommitment() []byte {
	if m != nil {
		return m.ValueCommitment
	}
	return nil
}

func (m *ReceiveDescription) GetNoteCommitment() []byte {
	if m != nil {
		return m.NoteCommitment
	}
	return nil
}

func (m *ReceiveDescription) GetEpk() []byte {
	if m != nil {
		return m.Epk
	}
	return nil
}

func (m *ReceiveDescription) GetCEnc() []byte {
	if m != nil {
		return m.CEnc
	}
	return nil
}

func (m *ReceiveDescription) GetCOut() []byte {
	if m != nil {
		return m.COut
	}
	return nil
}

func (m *ReceiveDescription) GetZkproof() []byte {
	if m != nil {
		return m.Zkproof
	}
	return nil
}

type ShieldedTransferContract struct {
	TransparentFromAddress []byte                `protobuf:"bytes,1,opt,name=transparent_from_address,json=transparentFromAddress,proto3" json:"transparent_from_address,omitempty"`
	FromAmount             int64                 `protobuf:"varint,2,opt,name=from_amount,json=fromAmount,proto3" json:"from_amount,omitempty"`
	SpendDescription       []*SpendDescription   `protobuf:"bytes,3,rep,name=spend_description,json=spendDescription,proto3" json:"spend_description,omitempty"`
	ReceiveDescription     []*ReceiveDescription `protobuf:"bytes,4,rep,name=receive_description,json=receiveDescription,proto3" json:"receive_description,omitempty"`
	BindingSignature       []byte                `protobuf:"bytes,5,opt,name=binding_signature,json=bindingSignature,proto3" json:"binding_signature,omitempty"`
	TransparentToAddress   []byte                `protobuf:"bytes,6,opt,name=transparent_to_address,json=transparentToAddress,proto3" json:"transparent_to_address,omitempty"`
	ToAmount               int64                 `protobuf:"varint,7,opt,name=to_amount,json=toAmount,proto3" json:"to_amount,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *ShieldedTransferContract) Reset()         { *m = ShieldedTransferContract{} }
func (m *ShieldedTransferContract) String() string { return proto.CompactTextString(m) }
func (*ShieldedTransferContract) ProtoMessage()    {}
func (*ShieldedTransferContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ff2a20a8b4c1fd, []int{10}
}

func (m *ShieldedTransferContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShieldedTransferContract.Unmarshal(m, b)
}
func (m *ShieldedTransferContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShieldedTransferContract.Marshal(b, m, deterministic)
}
func (m *ShieldedTransferContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShieldedTransferContract.Merge(m, src)
}
func (m *ShieldedTransferContract) XXX_Size() int {
	return xxx_messageInfo_ShieldedTransferContract.Size(m)
}
func (m *ShieldedTransferContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ShieldedTransferContract.DiscardUnknown(m)
}

var xxx_messageInfo_ShieldedTransferContract proto.InternalMessageInfo

func (m *ShieldedTransferContract) GetTransparentFromAddress() []byte {
	if m != nil {
		return m.TransparentFromAddress
	}
	return nil
}

func (m *ShieldedTransferContract) GetFromAmount() int64 {
	if m != nil {
		return m.FromAmount
	}
	return 0
}

func (m *ShieldedTransferContract) GetSpendDescription() []*SpendDescription {
	if m != nil {
		return m.SpendDescription
	}
	return nil
}

func (m *ShieldedTransferContract) GetReceiveDescription() []*ReceiveDescription {
	if m != nil {
		return m.ReceiveDescription
	}
	return nil
}

func (m *ShieldedTransferContract) GetBindingSignature() []byte {
	if m != nil {
		return m.BindingSignature
	}
	return nil
}

func (m *ShieldedTransferContract) GetTransparentToAddress() []byte {
	if m != nil {
		return m.TransparentToAddress
	}
	return nil
}

func (m *ShieldedTransferContract) GetToAmount() int64 {
	if m != nil {
		return m.ToAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthenticationPath)(nil), "protocol.AuthenticationPath")
	proto.RegisterType((*MerklePath)(nil), "protocol.MerklePath")
	proto.RegisterType((*OutputPoint)(nil), "protocol.OutputPoint")
	proto.RegisterType((*OutputPointInfo)(nil), "protocol.OutputPointInfo")
	proto.RegisterType((*PedersenHash)(nil), "protocol.PedersenHash")
	proto.RegisterType((*IncrementalMerkleTree)(nil), "protocol.IncrementalMerkleTree")
	proto.RegisterType((*IncrementalMerkleVoucher)(nil), "protocol.IncrementalMerkleVoucher")
	proto.RegisterType((*IncrementalMerkleVoucherInfo)(nil), "protocol.IncrementalMerkleVoucherInfo")
	proto.RegisterType((*SpendDescription)(nil), "protocol.SpendDescription")
	proto.RegisterType((*ReceiveDescription)(nil), "protocol.ReceiveDescription")
	proto.RegisterType((*ShieldedTransferContract)(nil), "protocol.ShieldedTransferContract")
}

func init() {
	proto.RegisterFile("core/contract/shield_contract.proto", fileDescriptor_53ff2a20a8b4c1fd)
}

var fileDescriptor_53ff2a20a8b4c1fd = []byte{
	// 851 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xc6, 0x3f, 0x71, 0x8e, 0xad, 0xc6, 0x9d, 0xa4, 0x61, 0xa1, 0x91, 0x9a, 0x2e, 0x17,
	0x98, 0x82, 0x6c, 0xd4, 0x56, 0x6a, 0x85, 0x04, 0x52, 0x68, 0xf9, 0xe9, 0x45, 0x48, 0xb4, 0x89,
	0xb8, 0xe0, 0x66, 0xb5, 0x9e, 0x3d, 0xf6, 0x8e, 0xbc, 0x3b, 0xb3, 0x9a, 0x99, 0x8d, 0x80, 0x5b,
	0xae, 0x79, 0x0a, 0xc4, 0x6b, 0xf0, 0x1e, 0xbc, 0x0d, 0x9a, 0x9f, 0x8d, 0xd7, 0x31, 0x46, 0x88,
	0x2b, 0xcf, 0x39, 0xf3, 0xcd, 0x9c, 0xf3, 0x7d, 0xf3, 0x9d, 0x35, 0x7c, 0x48, 0x85, 0xc4, 0x19,
	0x15, 0x5c, 0xcb, 0x94, 0xea, 0x99, 0xca, 0x19, 0x16, 0x59, 0xd2, 0xc4, 0xd3, 0x4a, 0x0a, 0x2d,
	0xc8, 0xc0, 0xfe, 0x50, 0x51, 0x44, 0xcf, 0x80, 0x9c, 0xd7, 0x3a, 0x47, 0xae, 0x19, 0x4d, 0x35,
	0x13, 0xfc, 0x2a, 0xd5, 0x39, 0x39, 0x86, 0xde, 0x6d, 0x5a, 0xd4, 0x18, 0x06, 0x67, 0x9d, 0xc9,
	0x20, 0x76, 0x41, 0xf4, 0x6b, 0x00, 0x70, 0x81, 0x72, 0x55, 0xa0, 0x05, 0x5d, 0xc2, 0x71, 0xba,
	0x71, 0x34, 0xa9, 0x52, 0x9d, 0x2b, 0x7b, 0x66, 0xf8, 0xfc, 0x74, 0xda, 0xd4, 0x98, 0x6e, 0x17,
	0x88, 0x8f, 0xd2, 0xad, 0x9c, 0x32, 0x55, 0x19, 0xcf, 0xf0, 0xa7, 0x70, 0xcf, 0x55, 0xb5, 0x01,
	0x79, 0x00, 0x7b, 0x52, 0x87, 0x9d, 0xb3, 0x60, 0x32, 0x8a, 0xf7, 0xa4, 0x8e, 0x5e, 0xc1, 0xf0,
	0xb2, 0xd6, 0x55, 0xad, 0xaf, 0x04, 0xe3, 0x9a, 0x10, 0xe8, 0xe6, 0xa9, 0xca, 0xc3, 0xc0, 0x02,
	0xec, 0xba, 0x7d, 0x51, 0x30, 0xe9, 0xf9, 0x8b, 0xa2, 0x0c, 0x0e, 0x5b, 0x07, 0xdf, 0xf1, 0x85,
	0x20, 0x2f, 0x01, 0x44, 0xad, 0x93, 0xca, 0x24, 0x9a, 0xc6, 0x1f, 0xad, 0x1b, 0x6f, 0xc1, 0xe3,
	0x03, 0xe1, 0x57, 0x8a, 0x3c, 0x86, 0x83, 0x79, 0x21, 0xe8, 0x2a, 0xe1, 0x75, 0xe9, 0x4b, 0x0c,
	0x6c, 0xe2, 0xfb, 0xba, 0x8c, 0x26, 0x30, 0xba, 0xc2, 0x0c, 0xa5, 0x42, 0xfe, 0x9d, 0xe9, 0x25,
	0x84, 0x7d, 0x23, 0x3e, 0x72, 0xed, 0x5b, 0x6c, 0xc2, 0xe8, 0xf7, 0x00, 0x1e, 0xbd, 0xe3, 0x54,
	0x62, 0x89, 0x5c, 0xa7, 0x85, 0x53, 0xf6, 0x46, 0x22, 0x92, 0x67, 0xd0, 0x2d, 0x70, 0xe1, 0x0e,
	0x0c, 0x9f, 0x9f, 0xac, 0x1b, 0x6a, 0xdf, 0x1c, 0x5b, 0x0c, 0xf9, 0x14, 0x7a, 0x92, 0x2d, 0x73,
	0x6d, 0x1b, 0xd9, 0x0d, 0x76, 0x20, 0xf2, 0x19, 0xec, 0x57, 0xa9, 0x44, 0xc3, 0xb6, 0x63, 0xd9,
	0xee, 0xc2, 0x37, 0xb0, 0xe8, 0x8f, 0x3d, 0x08, 0xb7, 0xba, 0xfc, 0x41, 0xd4, 0x34, 0x47, 0x49,
	0x5e, 0x40, 0x57, 0x4b, 0x44, 0xdf, 0xe8, 0x93, 0xf5, 0x5d, 0xff, 0xc8, 0x2b, 0xb6, 0x60, 0x32,
	0x85, 0xfe, 0x82, 0x15, 0x05, 0x66, 0xf6, 0x9d, 0x77, 0xb7, 0xe0, 0x51, 0xe4, 0x15, 0xf4, 0x69,
	0x2d, 0x95, 0x90, 0xd6, 0x04, 0xff, 0xa1, 0x8c, 0x87, 0x93, 0xa7, 0x30, 0x72, 0xab, 0x24, 0xc3,
	0x4a, 0xe7, 0x61, 0xf7, 0x2c, 0x98, 0x74, 0xe2, 0xa1, 0xcb, 0xbd, 0x35, 0x29, 0x6f, 0xae, 0x5e,
	0x63, 0x2e, 0xf2, 0x1a, 0x46, 0xc2, 0x3e, 0xba, 0xf3, 0x44, 0x08, 0xb6, 0xe2, 0x0e, 0x4b, 0x0c,
	0xc5, 0x3a, 0x88, 0x34, 0x9c, 0xee, 0x92, 0xc9, 0x5a, 0xed, 0x4b, 0x18, 0xdc, 0xba, 0xb0, 0x31,
	0x5a, 0xf4, 0x2f, 0x3c, 0xfc, 0xc9, 0xf8, 0xee, 0x8c, 0xf1, 0xb4, 0x1b, 0x2f, 0x23, 0xda, 0x28,
	0x76, 0x41, 0xf4, 0x57, 0x00, 0xe3, 0xeb, 0x0a, 0x79, 0xf6, 0x16, 0x15, 0x95, 0xac, 0x32, 0xc3,
	0x44, 0x3e, 0x86, 0xb1, 0x1d, 0xd8, 0x84, 0x8a, 0xb2, 0x64, 0xba, 0x5c, 0x7b, 0xef, 0xd0, 0xe6,
	0xdf, 0xdc, 0xa5, 0xc9, 0x09, 0xf4, 0x53, 0x4e, 0x73, 0x21, 0xad, 0x7d, 0x46, 0xb1, 0x8f, 0xc8,
	0x29, 0x1c, 0xf0, 0xba, 0x28, 0xd8, 0x82, 0xa1, 0xf4, 0xb3, 0xb7, 0x4e, 0x58, 0xd5, 0x56, 0x56,
	0x4e, 0xa3, 0xda, 0xca, 0x78, 0xfc, 0x97, 0x55, 0x25, 0x85, 0x58, 0x78, 0x29, 0x9b, 0x90, 0x7c,
	0x0e, 0xef, 0x2b, 0xd3, 0x5e, 0x62, 0xe6, 0x5d, 0x48, 0xa6, 0x7f, 0x4e, 0x14, 0x5b, 0xf2, 0x54,
	0xd7, 0x12, 0xc3, 0xbe, 0xc5, 0xbe, 0x67, 0x01, 0xe7, 0xcd, 0xfe, 0x75, 0xb3, 0x1d, 0xfd, 0x19,
	0x00, 0x89, 0x91, 0x22, 0xbb, 0xc5, 0xff, 0xc9, 0xee, 0x23, 0x38, 0xe4, 0x42, 0x6f, 0x20, 0x1d,
	0xcd, 0x07, 0x26, 0xdd, 0x02, 0x8e, 0xa1, 0x83, 0xd5, 0xca, 0x13, 0x35, 0x4b, 0x72, 0x04, 0x3d,
	0x9a, 0x20, 0xa7, 0x9e, 0x65, 0x97, 0x7e, 0xcd, 0xa9, 0x4b, 0x8a, 0xba, 0x31, 0x4c, 0x97, 0x5e,
	0xd6, 0xba, 0x4d, 0xbe, 0xbf, 0x41, 0x3e, 0xfa, 0xad, 0x03, 0xe1, 0xb5, 0xfd, 0xfe, 0x62, 0x76,
	0x23, 0x53, 0xae, 0x16, 0x28, 0xdf, 0xf8, 0x0f, 0x31, 0x79, 0x0d, 0xa1, 0x36, 0x39, 0x37, 0x67,
	0xc9, 0x42, 0x8a, 0x32, 0x49, 0xb3, 0x4c, 0xa2, 0x52, 0x9e, 0xce, 0x49, 0x6b, 0xff, 0x1b, 0x29,
	0xca, 0x73, 0xb7, 0x4b, 0x9e, 0xc0, 0xd0, 0xa1, 0x4b, 0x51, 0x7b, 0x46, 0x9d, 0x18, 0x4c, 0xea,
	0xdc, 0x66, 0xc8, 0xb7, 0xf0, 0xd0, 0x89, 0x9e, 0xad, 0x65, 0xf3, 0xe3, 0xfe, 0xc1, 0xda, 0x73,
	0xf7, 0x6d, 0x13, 0x8f, 0xd5, 0x7d, 0x23, 0x5d, 0xc0, 0x91, 0x74, 0x0f, 0xb0, 0x71, 0x55, 0xf7,
	0xfe, 0x07, 0x7e, 0xfb, 0x95, 0x62, 0x22, 0xb7, 0x5f, 0xee, 0x13, 0x78, 0x38, 0x67, 0x3c, 0x63,
	0x7c, 0xd9, 0x32, 0x81, 0x93, 0x72, 0xec, 0x37, 0xee, 0x5e, 0x9f, 0xbc, 0x84, 0x36, 0xff, 0x44,
	0x8b, 0x3b, 0x75, 0x9c, 0xca, 0xc7, 0xad, 0xdd, 0x1b, 0xd1, 0x68, 0xf3, 0x18, 0x0e, 0x0c, 0xd2,
	0x29, 0xb3, 0x6f, 0x95, 0x19, 0x68, 0xe1, 0x74, 0xf9, 0xea, 0x0b, 0x08, 0x85, 0x5c, 0x4e, 0xb5,
	0x14, 0xdc, 0xf5, 0xae, 0xa6, 0xcd, 0xff, 0xe2, 0x8f, 0x4f, 0x97, 0x4c, 0xe7, 0xf5, 0x7c, 0x4a,
	0x45, 0x39, 0x5b, 0xcc, 0x95, 0x98, 0x4b, 0x64, 0x32, 0x9d, 0x2d, 0x85, 0x41, 0xcf, 0xcc, 0x7f,
	0xea, 0xbc, 0x6f, 0xcf, 0xbc, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x99, 0x31, 0x54, 0x62,
	0x07, 0x00, 0x00,
}
