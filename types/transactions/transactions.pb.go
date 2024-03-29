// Copyright (c) 2024 Project Illium
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: transactions.proto

package transactions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MintTransaction_AssetType int32

const (
	MintTransaction_FIXED_SUPPLY    MintTransaction_AssetType = 0
	MintTransaction_VARIABLE_SUPPLY MintTransaction_AssetType = 1
)

// Enum value maps for MintTransaction_AssetType.
var (
	MintTransaction_AssetType_name = map[int32]string{
		0: "FIXED_SUPPLY",
		1: "VARIABLE_SUPPLY",
	}
	MintTransaction_AssetType_value = map[string]int32{
		"FIXED_SUPPLY":    0,
		"VARIABLE_SUPPLY": 1,
	}
)

func (x MintTransaction_AssetType) Enum() *MintTransaction_AssetType {
	p := new(MintTransaction_AssetType)
	*p = x
	return p
}

func (x MintTransaction_AssetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MintTransaction_AssetType) Descriptor() protoreflect.EnumDescriptor {
	return file_transactions_proto_enumTypes[0].Descriptor()
}

func (MintTransaction_AssetType) Type() protoreflect.EnumType {
	return &file_transactions_proto_enumTypes[0]
}

func (x MintTransaction_AssetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MintTransaction_AssetType.Descriptor instead.
func (MintTransaction_AssetType) EnumDescriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{6, 0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Tx:
	//	*Transaction_StandardTransaction
	//	*Transaction_CoinbaseTransaction
	//	*Transaction_StakeTransaction
	//	*Transaction_TreasuryTransaction
	//	*Transaction_MintTransaction
	Tx         isTransaction_Tx `protobuf_oneof:"Tx"`
	cachedTxid []byte
	cachedWid  []byte
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{0}
}

func (m *Transaction) GetTx() isTransaction_Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (x *Transaction) GetStandardTransaction() *StandardTransaction {
	if x, ok := x.GetTx().(*Transaction_StandardTransaction); ok {
		return x.StandardTransaction
	}
	return nil
}

func (x *Transaction) GetCoinbaseTransaction() *CoinbaseTransaction {
	if x, ok := x.GetTx().(*Transaction_CoinbaseTransaction); ok {
		return x.CoinbaseTransaction
	}
	return nil
}

func (x *Transaction) GetStakeTransaction() *StakeTransaction {
	if x, ok := x.GetTx().(*Transaction_StakeTransaction); ok {
		return x.StakeTransaction
	}
	return nil
}

func (x *Transaction) GetTreasuryTransaction() *TreasuryTransaction {
	if x, ok := x.GetTx().(*Transaction_TreasuryTransaction); ok {
		return x.TreasuryTransaction
	}
	return nil
}

func (x *Transaction) GetMintTransaction() *MintTransaction {
	if x, ok := x.GetTx().(*Transaction_MintTransaction); ok {
		return x.MintTransaction
	}
	return nil
}

type isTransaction_Tx interface {
	isTransaction_Tx()
}

type Transaction_StandardTransaction struct {
	StandardTransaction *StandardTransaction `protobuf:"bytes,1,opt,name=standard_transaction,json=standardTransaction,proto3,oneof"`
}

type Transaction_CoinbaseTransaction struct {
	CoinbaseTransaction *CoinbaseTransaction `protobuf:"bytes,2,opt,name=coinbase_transaction,json=coinbaseTransaction,proto3,oneof"`
}

type Transaction_StakeTransaction struct {
	StakeTransaction *StakeTransaction `protobuf:"bytes,3,opt,name=stake_transaction,json=stakeTransaction,proto3,oneof"`
}

type Transaction_TreasuryTransaction struct {
	TreasuryTransaction *TreasuryTransaction `protobuf:"bytes,4,opt,name=treasury_transaction,json=treasuryTransaction,proto3,oneof"`
}

type Transaction_MintTransaction struct {
	MintTransaction *MintTransaction `protobuf:"bytes,5,opt,name=mint_transaction,json=mintTransaction,proto3,oneof"`
}

func (*Transaction_StandardTransaction) isTransaction_Tx() {}

func (*Transaction_CoinbaseTransaction) isTransaction_Tx() {}

func (*Transaction_StakeTransaction) isTransaction_Tx() {}

func (*Transaction_TreasuryTransaction) isTransaction_Tx() {}

func (*Transaction_MintTransaction) isTransaction_Tx() {}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	Ciphertext []byte `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *Output) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

type StandardTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outputs    []*Output `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Nullifiers [][]byte  `protobuf:"bytes,2,rep,name=nullifiers,proto3" json:"nullifiers,omitempty"`
	TxoRoot    []byte    `protobuf:"bytes,3,opt,name=txo_root,json=txoRoot,proto3" json:"txo_root,omitempty"`
	Locktime   *Locktime `protobuf:"bytes,4,opt,name=locktime,proto3" json:"locktime,omitempty"`
	Fee        uint64    `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Proof      []byte    `protobuf:"bytes,6,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *StandardTransaction) Reset() {
	*x = StandardTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StandardTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardTransaction) ProtoMessage() {}

func (x *StandardTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardTransaction.ProtoReflect.Descriptor instead.
func (*StandardTransaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{2}
}

func (x *StandardTransaction) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *StandardTransaction) GetNullifiers() [][]byte {
	if x != nil {
		return x.Nullifiers
	}
	return nil
}

func (x *StandardTransaction) GetTxoRoot() []byte {
	if x != nil {
		return x.TxoRoot
	}
	return nil
}

func (x *StandardTransaction) GetLocktime() *Locktime {
	if x != nil {
		return x.Locktime
	}
	return nil
}

func (x *StandardTransaction) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *StandardTransaction) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type CoinbaseTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validator_ID []byte    `protobuf:"bytes,1,opt,name=validator_ID,json=validatorID,proto3" json:"validator_ID,omitempty"`
	NewCoins     uint64    `protobuf:"varint,2,opt,name=new_coins,json=newCoins,proto3" json:"new_coins,omitempty"`
	Outputs      []*Output `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Signature    []byte    `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Proof        []byte    `protobuf:"bytes,5,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *CoinbaseTransaction) Reset() {
	*x = CoinbaseTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinbaseTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinbaseTransaction) ProtoMessage() {}

func (x *CoinbaseTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinbaseTransaction.ProtoReflect.Descriptor instead.
func (*CoinbaseTransaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{3}
}

func (x *CoinbaseTransaction) GetValidator_ID() []byte {
	if x != nil {
		return x.Validator_ID
	}
	return nil
}

func (x *CoinbaseTransaction) GetNewCoins() uint64 {
	if x != nil {
		return x.NewCoins
	}
	return 0
}

func (x *CoinbaseTransaction) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *CoinbaseTransaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CoinbaseTransaction) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type StakeTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Validator_ID []byte `protobuf:"bytes,1,opt,name=validator_ID,json=validatorID,proto3" json:"validator_ID,omitempty"`
	Amount       uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Nullifier    []byte `protobuf:"bytes,3,opt,name=nullifier,proto3" json:"nullifier,omitempty"`
	TxoRoot      []byte `protobuf:"bytes,4,opt,name=txo_root,json=txoRoot,proto3" json:"txo_root,omitempty"`
	LockedUntil  int64  `protobuf:"varint,5,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	Signature    []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Proof        []byte `protobuf:"bytes,7,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *StakeTransaction) Reset() {
	*x = StakeTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StakeTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StakeTransaction) ProtoMessage() {}

func (x *StakeTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StakeTransaction.ProtoReflect.Descriptor instead.
func (*StakeTransaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{4}
}

func (x *StakeTransaction) GetValidator_ID() []byte {
	if x != nil {
		return x.Validator_ID
	}
	return nil
}

func (x *StakeTransaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *StakeTransaction) GetNullifier() []byte {
	if x != nil {
		return x.Nullifier
	}
	return nil
}

func (x *StakeTransaction) GetTxoRoot() []byte {
	if x != nil {
		return x.TxoRoot
	}
	return nil
}

func (x *StakeTransaction) GetLockedUntil() int64 {
	if x != nil {
		return x.LockedUntil
	}
	return 0
}

func (x *StakeTransaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *StakeTransaction) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type TreasuryTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount       uint64    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Outputs      []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
	ProposalHash []byte    `protobuf:"bytes,3,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	Proof        []byte    `protobuf:"bytes,4,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *TreasuryTransaction) Reset() {
	*x = TreasuryTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasuryTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasuryTransaction) ProtoMessage() {}

func (x *TreasuryTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasuryTransaction.ProtoReflect.Descriptor instead.
func (*TreasuryTransaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{5}
}

func (x *TreasuryTransaction) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TreasuryTransaction) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *TreasuryTransaction) GetProposalHash() []byte {
	if x != nil {
		return x.ProposalHash
	}
	return nil
}

func (x *TreasuryTransaction) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type MintTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         MintTransaction_AssetType `protobuf:"varint,1,opt,name=type,proto3,enum=MintTransaction_AssetType" json:"type,omitempty"`
	Asset_ID     []byte                    `protobuf:"bytes,2,opt,name=asset_ID,json=assetID,proto3" json:"asset_ID,omitempty"`
	DocumentHash []byte                    `protobuf:"bytes,3,opt,name=document_hash,json=documentHash,proto3" json:"document_hash,omitempty"`
	NewTokens    uint64                    `protobuf:"varint,4,opt,name=new_tokens,json=newTokens,proto3" json:"new_tokens,omitempty"`
	Outputs      []*Output                 `protobuf:"bytes,5,rep,name=outputs,proto3" json:"outputs,omitempty"`
	Fee          uint64                    `protobuf:"varint,6,opt,name=fee,proto3" json:"fee,omitempty"`
	Nullifiers   [][]byte                  `protobuf:"bytes,7,rep,name=nullifiers,proto3" json:"nullifiers,omitempty"`
	TxoRoot      []byte                    `protobuf:"bytes,8,opt,name=txo_root,json=txoRoot,proto3" json:"txo_root,omitempty"`
	MintKey      []byte                    `protobuf:"bytes,9,opt,name=mint_key,json=mintKey,proto3" json:"mint_key,omitempty"`
	Locktime     *Locktime                 `protobuf:"bytes,10,opt,name=locktime,proto3" json:"locktime,omitempty"`
	Signature    []byte                    `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	Proof        []byte                    `protobuf:"bytes,12,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (x *MintTransaction) Reset() {
	*x = MintTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintTransaction) ProtoMessage() {}

func (x *MintTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintTransaction.ProtoReflect.Descriptor instead.
func (*MintTransaction) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{6}
}

func (x *MintTransaction) GetType() MintTransaction_AssetType {
	if x != nil {
		return x.Type
	}
	return MintTransaction_FIXED_SUPPLY
}

func (x *MintTransaction) GetAsset_ID() []byte {
	if x != nil {
		return x.Asset_ID
	}
	return nil
}

func (x *MintTransaction) GetDocumentHash() []byte {
	if x != nil {
		return x.DocumentHash
	}
	return nil
}

func (x *MintTransaction) GetNewTokens() uint64 {
	if x != nil {
		return x.NewTokens
	}
	return 0
}

func (x *MintTransaction) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *MintTransaction) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *MintTransaction) GetNullifiers() [][]byte {
	if x != nil {
		return x.Nullifiers
	}
	return nil
}

func (x *MintTransaction) GetTxoRoot() []byte {
	if x != nil {
		return x.TxoRoot
	}
	return nil
}

func (x *MintTransaction) GetMintKey() []byte {
	if x != nil {
		return x.MintKey
	}
	return nil
}

func (x *MintTransaction) GetLocktime() *Locktime {
	if x != nil {
		return x.Locktime
	}
	return nil
}

func (x *MintTransaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *MintTransaction) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

type Locktime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Precision int64 `protobuf:"varint,2,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *Locktime) Reset() {
	*x = Locktime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transactions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locktime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locktime) ProtoMessage() {}

func (x *Locktime) ProtoReflect() protoreflect.Message {
	mi := &file_transactions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locktime.ProtoReflect.Descriptor instead.
func (*Locktime) Descriptor() ([]byte, []int) {
	return file_transactions_proto_rawDescGZIP(), []int{7}
}

func (x *Locktime) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Locktime) GetPrecision() int64 {
	if x != nil {
		return x.Precision
	}
	return 0
}

var File_transactions_proto protoreflect.FileDescriptor

var file_transactions_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x73, 0x74, 0x61, 0x6e,
	0x64, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x49, 0x0a, 0x14, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x13, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x11, 0x73, 0x74,
	0x61, 0x6b, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x73, 0x74, 0x61, 0x6b,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x14,
	0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x54, 0x72, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x13, 0x74, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x10, 0x6d, 0x69, 0x6e, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x4d, 0x69, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0x0a, 0x02, 0x54, 0x78, 0x22, 0x48, 0x0a, 0x06,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x6e, 0x64,
	0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6f, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6f, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xac, 0x01, 0x0a, 0x13,
	0x43, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f,
	0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xdd, 0x01, 0x0a, 0x10, 0x53,
	0x74, 0x61, 0x6b, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x21, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x75,
	0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6e,
	0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6f, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6f, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x75, 0x6e,
	0x74, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x54,
	0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0xba, 0x03, 0x0a, 0x0f, 0x4d, 0x69, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x4d, 0x69, 0x6e,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x65, 0x77, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x07, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x6f, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x6f, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d,
	0x69, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d,
	0x69, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x74,
	0x69, 0x6d, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x22, 0x32, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10,
	0x0a, 0x0c, 0x46, 0x49, 0x58, 0x45, 0x44, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4c, 0x59, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x50,
	0x50, 0x4c, 0x59, 0x10, 0x01, 0x22, 0x46, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transactions_proto_rawDescOnce sync.Once
	file_transactions_proto_rawDescData = file_transactions_proto_rawDesc
)

func file_transactions_proto_rawDescGZIP() []byte {
	file_transactions_proto_rawDescOnce.Do(func() {
		file_transactions_proto_rawDescData = protoimpl.X.CompressGZIP(file_transactions_proto_rawDescData)
	})
	return file_transactions_proto_rawDescData
}

var file_transactions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transactions_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_transactions_proto_goTypes = []interface{}{
	(MintTransaction_AssetType)(0), // 0: MintTransaction.AssetType
	(*Transaction)(nil),            // 1: Transaction
	(*Output)(nil),                 // 2: Output
	(*StandardTransaction)(nil),    // 3: StandardTransaction
	(*CoinbaseTransaction)(nil),    // 4: CoinbaseTransaction
	(*StakeTransaction)(nil),       // 5: StakeTransaction
	(*TreasuryTransaction)(nil),    // 6: TreasuryTransaction
	(*MintTransaction)(nil),        // 7: MintTransaction
	(*Locktime)(nil),               // 8: Locktime
}
var file_transactions_proto_depIdxs = []int32{
	3,  // 0: Transaction.standard_transaction:type_name -> StandardTransaction
	4,  // 1: Transaction.coinbase_transaction:type_name -> CoinbaseTransaction
	5,  // 2: Transaction.stake_transaction:type_name -> StakeTransaction
	6,  // 3: Transaction.treasury_transaction:type_name -> TreasuryTransaction
	7,  // 4: Transaction.mint_transaction:type_name -> MintTransaction
	2,  // 5: StandardTransaction.outputs:type_name -> Output
	8,  // 6: StandardTransaction.locktime:type_name -> Locktime
	2,  // 7: CoinbaseTransaction.outputs:type_name -> Output
	2,  // 8: TreasuryTransaction.outputs:type_name -> Output
	0,  // 9: MintTransaction.type:type_name -> MintTransaction.AssetType
	2,  // 10: MintTransaction.outputs:type_name -> Output
	8,  // 11: MintTransaction.locktime:type_name -> Locktime
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_transactions_proto_init() }
func file_transactions_proto_init() {
	if File_transactions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transactions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StandardTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinbaseTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StakeTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasuryTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transactions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locktime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_transactions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Transaction_StandardTransaction)(nil),
		(*Transaction_CoinbaseTransaction)(nil),
		(*Transaction_StakeTransaction)(nil),
		(*Transaction_TreasuryTransaction)(nil),
		(*Transaction_MintTransaction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transactions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transactions_proto_goTypes,
		DependencyIndexes: file_transactions_proto_depIdxs,
		EnumInfos:         file_transactions_proto_enumTypes,
		MessageInfos:      file_transactions_proto_msgTypes,
	}.Build()
	File_transactions_proto = out.File
	file_transactions_proto_rawDesc = nil
	file_transactions_proto_goTypes = nil
	file_transactions_proto_depIdxs = nil
}
