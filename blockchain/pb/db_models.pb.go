// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: db_models.proto

package pb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/project-illium/ilxd/types/transactions"
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

type DBValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerId         string                   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	TotalStake     uint64                   `protobuf:"varint,2,opt,name=total_stake,json=totalStake,proto3" json:"total_stake,omitempty"`
	Nullifiers     []*DBValidator_Nullifier `protobuf:"bytes,3,rep,name=nullifiers,proto3" json:"nullifiers,omitempty"`
	UnclaimedCoins uint64                   `protobuf:"varint,4,opt,name=unclaimed_coins,json=unclaimedCoins,proto3" json:"unclaimed_coins,omitempty"`
	EpochBLocks    uint32                   `protobuf:"varint,5,opt,name=epochBLocks,proto3" json:"epochBLocks,omitempty"`
}

func (x *DBValidator) Reset() {
	*x = DBValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBValidator) ProtoMessage() {}

func (x *DBValidator) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBValidator.ProtoReflect.Descriptor instead.
func (*DBValidator) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{0}
}

func (x *DBValidator) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *DBValidator) GetTotalStake() uint64 {
	if x != nil {
		return x.TotalStake
	}
	return 0
}

func (x *DBValidator) GetNullifiers() []*DBValidator_Nullifier {
	if x != nil {
		return x.Nullifiers
	}
	return nil
}

func (x *DBValidator) GetUnclaimedCoins() uint64 {
	if x != nil {
		return x.UnclaimedCoins
	}
	return 0
}

func (x *DBValidator) GetEpochBLocks() uint32 {
	if x != nil {
		return x.EpochBLocks
	}
	return 0
}

type DBTxs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*transactions.Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *DBTxs) Reset() {
	*x = DBTxs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTxs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTxs) ProtoMessage() {}

func (x *DBTxs) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTxs.ProtoReflect.Descriptor instead.
func (*DBTxs) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{1}
}

func (x *DBTxs) GetTransactions() []*transactions.Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type DBBlockNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockID []byte `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	Height  uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *DBBlockNode) Reset() {
	*x = DBBlockNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBBlockNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBBlockNode) ProtoMessage() {}

func (x *DBBlockNode) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBBlockNode.ProtoReflect.Descriptor instead.
func (*DBBlockNode) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{2}
}

func (x *DBBlockNode) GetBlockID() []byte {
	if x != nil {
		return x.BlockID
	}
	return nil
}

func (x *DBBlockNode) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type DBAccumulator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accumulator [][]byte                        `protobuf:"bytes,1,rep,name=accumulator,proto3" json:"accumulator,omitempty"`
	NElements   uint64                          `protobuf:"varint,2,opt,name=nElements,proto3" json:"nElements,omitempty"`
	Proofs      []*DBAccumulator_InclusionProof `protobuf:"bytes,3,rep,name=proofs,proto3" json:"proofs,omitempty"`
	LookupMap   []*DBAccumulator_InclusionProof `protobuf:"bytes,4,rep,name=lookupMap,proto3" json:"lookupMap,omitempty"`
}

func (x *DBAccumulator) Reset() {
	*x = DBAccumulator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBAccumulator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBAccumulator) ProtoMessage() {}

func (x *DBAccumulator) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBAccumulator.ProtoReflect.Descriptor instead.
func (*DBAccumulator) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{3}
}

func (x *DBAccumulator) GetAccumulator() [][]byte {
	if x != nil {
		return x.Accumulator
	}
	return nil
}

func (x *DBAccumulator) GetNElements() uint64 {
	if x != nil {
		return x.NElements
	}
	return 0
}

func (x *DBAccumulator) GetProofs() []*DBAccumulator_InclusionProof {
	if x != nil {
		return x.Proofs
	}
	return nil
}

func (x *DBAccumulator) GetLookupMap() []*DBAccumulator_InclusionProof {
	if x != nil {
		return x.LookupMap
	}
	return nil
}

type DBValidator_Nullifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash       []byte               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Amount     uint64               `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Blockstamp *timestamp.Timestamp `protobuf:"bytes,3,opt,name=blockstamp,proto3" json:"blockstamp,omitempty"`
}

func (x *DBValidator_Nullifier) Reset() {
	*x = DBValidator_Nullifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBValidator_Nullifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBValidator_Nullifier) ProtoMessage() {}

func (x *DBValidator_Nullifier) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBValidator_Nullifier.ProtoReflect.Descriptor instead.
func (*DBValidator_Nullifier) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DBValidator_Nullifier) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *DBValidator_Nullifier) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DBValidator_Nullifier) GetBlockstamp() *timestamp.Timestamp {
	if x != nil {
		return x.Blockstamp
	}
	return nil
}

type DBAccumulator_InclusionProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index  uint64   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Hashes [][]byte `protobuf:"bytes,3,rep,name=hashes,proto3" json:"hashes,omitempty"`
	Flags  uint64   `protobuf:"varint,4,opt,name=flags,proto3" json:"flags,omitempty"`
	Last   []byte   `protobuf:"bytes,5,opt,name=last,proto3" json:"last,omitempty"`
}

func (x *DBAccumulator_InclusionProof) Reset() {
	*x = DBAccumulator_InclusionProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_db_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBAccumulator_InclusionProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBAccumulator_InclusionProof) ProtoMessage() {}

func (x *DBAccumulator_InclusionProof) ProtoReflect() protoreflect.Message {
	mi := &file_db_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBAccumulator_InclusionProof.ProtoReflect.Descriptor instead.
func (*DBAccumulator_InclusionProof) Descriptor() ([]byte, []int) {
	return file_db_models_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DBAccumulator_InclusionProof) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DBAccumulator_InclusionProof) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DBAccumulator_InclusionProof) GetHashes() [][]byte {
	if x != nil {
		return x.Hashes
	}
	return nil
}

func (x *DBAccumulator_InclusionProof) GetFlags() uint64 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *DBAccumulator_InclusionProof) GetLast() []byte {
	if x != nil {
		return x.Last
	}
	return nil
}

var File_db_models_proto protoreflect.FileDescriptor

var file_db_models_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x62, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x0b, 0x44, 0x42, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x12, 0x36, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x44, 0x42, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x0a, 0x6e, 0x75,
	0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0e, 0x75, 0x6e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x4c, 0x6f, 0x63, 0x6b, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x42, 0x4c, 0x6f,
	0x63, 0x6b, 0x73, 0x1a, 0x73, 0x0a, 0x09, 0x4e, 0x75, 0x6c, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x39, 0x0a, 0x05, 0x44, 0x42, 0x54, 0x78,
	0x73, 0x12, 0x30, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x44, 0x42, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x44, 0x42, 0x41, 0x63, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x45, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x44, 0x42, 0x41, 0x63, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x12, 0x3b, 0x0a,
	0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x44, 0x42, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52,
	0x09, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x4d, 0x61, 0x70, 0x1a, 0x78, 0x0a, 0x0e, 0x49, 0x6e,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6c, 0x61, 0x73, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_db_models_proto_rawDescOnce sync.Once
	file_db_models_proto_rawDescData = file_db_models_proto_rawDesc
)

func file_db_models_proto_rawDescGZIP() []byte {
	file_db_models_proto_rawDescOnce.Do(func() {
		file_db_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_db_models_proto_rawDescData)
	})
	return file_db_models_proto_rawDescData
}

var file_db_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_db_models_proto_goTypes = []interface{}{
	(*DBValidator)(nil),                  // 0: DBValidator
	(*DBTxs)(nil),                        // 1: DBTxs
	(*DBBlockNode)(nil),                  // 2: DBBlockNode
	(*DBAccumulator)(nil),                // 3: DBAccumulator
	(*DBValidator_Nullifier)(nil),        // 4: DBValidator.Nullifier
	(*DBAccumulator_InclusionProof)(nil), // 5: DBAccumulator.InclusionProof
	(*transactions.Transaction)(nil),     // 6: Transaction
	(*timestamp.Timestamp)(nil),          // 7: google.protobuf.Timestamp
}
var file_db_models_proto_depIdxs = []int32{
	4, // 0: DBValidator.nullifiers:type_name -> DBValidator.Nullifier
	6, // 1: DBTxs.transactions:type_name -> Transaction
	5, // 2: DBAccumulator.proofs:type_name -> DBAccumulator.InclusionProof
	5, // 3: DBAccumulator.lookupMap:type_name -> DBAccumulator.InclusionProof
	7, // 4: DBValidator.Nullifier.blockstamp:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_db_models_proto_init() }
func file_db_models_proto_init() {
	if File_db_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_db_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBValidator); i {
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
		file_db_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBTxs); i {
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
		file_db_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBBlockNode); i {
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
		file_db_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBAccumulator); i {
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
		file_db_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBValidator_Nullifier); i {
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
		file_db_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBAccumulator_InclusionProof); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_db_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_db_models_proto_goTypes,
		DependencyIndexes: file_db_models_proto_depIdxs,
		MessageInfos:      file_db_models_proto_msgTypes,
	}.Build()
	File_db_models_proto = out.File
	file_db_models_proto_rawDesc = nil
	file_db_models_proto_goTypes = nil
	file_db_models_proto_depIdxs = nil
}
