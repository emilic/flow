// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow/entities/block.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Block struct {
	Id                       []byte                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId                 []byte                  `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Height                   uint64                  `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp                *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CollectionGuarantees     []*CollectionGuarantee  `protobuf:"bytes,5,rep,name=collection_guarantees,json=collectionGuarantees,proto3" json:"collection_guarantees,omitempty"`
	BlockSeals               []*BlockSeal            `protobuf:"bytes,6,rep,name=block_seals,json=blockSeals,proto3" json:"block_seals,omitempty"`
	Signatures               [][]byte                `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty"`
	ExecutionReceiptMetaList []*ExecutionReceiptMeta `protobuf:"bytes,8,rep,name=execution_receipt_metaList,json=executionReceiptMetaList,proto3" json:"execution_receipt_metaList,omitempty"`
	ExecutionResultList      []*ExecutionResult      `protobuf:"bytes,9,rep,name=execution_result_list,json=executionResultList,proto3" json:"execution_result_list,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                `json:"-"`
	XXX_unrecognized         []byte                  `json:"-"`
	XXX_sizecache            int32                   `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1afc3335f2172fc, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Block) GetParentId() []byte {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetTimestamp() *timestamppb.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetCollectionGuarantees() []*CollectionGuarantee {
	if m != nil {
		return m.CollectionGuarantees
	}
	return nil
}

func (m *Block) GetBlockSeals() []*BlockSeal {
	if m != nil {
		return m.BlockSeals
	}
	return nil
}

func (m *Block) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Block) GetExecutionReceiptMetaList() []*ExecutionReceiptMeta {
	if m != nil {
		return m.ExecutionReceiptMetaList
	}
	return nil
}

func (m *Block) GetExecutionResultList() []*ExecutionResult {
	if m != nil {
		return m.ExecutionResultList
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "flow.entities.Block")
}

func init() { proto.RegisterFile("flow/entities/block.proto", fileDescriptor_c1afc3335f2172fc) }

var fileDescriptor_c1afc3335f2172fc = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x8b, 0xd4, 0x30,
	0x14, 0xc6, 0xe9, 0x74, 0x77, 0xdc, 0xc9, 0xac, 0x1e, 0xa2, 0x2b, 0xb1, 0xca, 0x58, 0x56, 0x0f,
	0x3d, 0xa5, 0xb2, 0x5e, 0xf4, 0xba, 0x22, 0x22, 0x28, 0x48, 0x14, 0x04, 0x2f, 0x25, 0x6d, 0xdf,
	0x66, 0x82, 0x69, 0x53, 0x9a, 0x57, 0xf4, 0xdf, 0xf6, 0x3f, 0x90, 0xa6, 0xd3, 0xce, 0x74, 0x1c,
	0x2f, 0xa5, 0xf9, 0xbe, 0xef, 0xfd, 0x1e, 0xef, 0x25, 0xe4, 0xc9, 0x9d, 0xb1, 0xbf, 0x52, 0xa8,
	0x51, 0xa3, 0x06, 0x97, 0xe6, 0xc6, 0x16, 0x3f, 0x79, 0xd3, 0x5a, 0xb4, 0xf4, 0x7e, 0x6f, 0xf1,
	0xd1, 0x8a, 0x9e, 0x2b, 0x6b, 0x95, 0x81, 0xd4, 0x9b, 0x79, 0x77, 0x97, 0xa2, 0xae, 0xc0, 0xa1,
	0xac, 0x9a, 0x21, 0x1f, 0x6d, 0xe6, 0xa8, 0xc2, 0x1a, 0x03, 0x05, 0x6a, 0x5b, 0x9f, 0xf6, 0x7d,
	0xab, 0xcc, 0x81, 0x34, 0x3b, 0xff, 0xe5, 0xdc, 0x87, 0xdf, 0x50, 0x74, 0x7d, 0x79, 0xd6, 0x82,
	0xeb, 0x0c, 0x0e, 0xa9, 0xeb, 0x3f, 0x21, 0x39, 0xbf, 0xed, 0x4b, 0xe9, 0x03, 0xb2, 0xd0, 0x25,
	0x0b, 0xe2, 0x20, 0xb9, 0x14, 0x0b, 0x5d, 0xd2, 0xa7, 0x64, 0xd5, 0xc8, 0x16, 0x6a, 0xcc, 0x74,
	0xc9, 0x16, 0x5e, 0xbe, 0x18, 0x84, 0x8f, 0x25, 0x7d, 0x4c, 0x96, 0x5b, 0xd0, 0x6a, 0x8b, 0x2c,
	0x8c, 0x83, 0xe4, 0x4c, 0xec, 0x4e, 0xf4, 0x0d, 0x59, 0x4d, 0x73, 0xb0, 0xb3, 0x38, 0x48, 0xd6,
	0x37, 0x11, 0x1f, 0x26, 0xe5, 0xe3, 0xa4, 0xfc, 0xdb, 0x98, 0x10, 0xfb, 0x30, 0xfd, 0x4e, 0xae,
	0xf6, 0x23, 0x66, 0xaa, 0x93, 0xad, 0xac, 0x11, 0xc0, 0xb1, 0xf3, 0x38, 0x4c, 0xd6, 0x37, 0xd7,
	0x7c, 0xb6, 0x3e, 0xfe, 0x6e, 0xca, 0x7e, 0x18, 0xa3, 0xe2, 0x51, 0xf1, 0xaf, 0xe8, 0xe8, 0x5b,
	0xb2, 0xde, 0xef, 0xc6, 0xb1, 0xa5, 0xc7, 0xb1, 0x23, 0x9c, 0x5f, 0xc1, 0x57, 0x90, 0x46, 0x90,
	0x7c, 0xfc, 0x75, 0x74, 0x43, 0x88, 0xd3, 0xaa, 0x96, 0xd8, 0xb5, 0xe0, 0xd8, 0xbd, 0x38, 0x4c,
	0x2e, 0xc5, 0x81, 0x42, 0x25, 0x89, 0x0e, 0xd7, 0x5a, 0x80, 0x6e, 0x30, 0xab, 0x00, 0xe5, 0x27,
	0xed, 0x90, 0x5d, 0xf8, 0x4e, 0x2f, 0x8e, 0x3a, 0xbd, 0x1f, 0x0b, 0xc4, 0x90, 0xff, 0x0c, 0x28,
	0x05, 0x83, 0x13, 0x6a, 0x0f, 0xa1, 0x82, 0x5c, 0x1d, 0xdf, 0x5c, 0x66, 0x7a, 0xfa, 0xca, 0xd3,
	0x37, 0xff, 0xa7, 0xf7, 0x51, 0xf1, 0x10, 0xe6, 0x42, 0xcf, 0xbc, 0xfd, 0x42, 0x9e, 0xd9, 0x56,
	0x71, 0x5b, 0xfb, 0xda, 0xe9, 0x5a, 0x46, 0xc8, 0x8f, 0x57, 0x4a, 0xe3, 0xb6, 0xcb, 0x79, 0x61,
	0xab, 0x74, 0x08, 0xa5, 0xfe, 0x33, 0x3d, 0x55, 0x65, 0xd3, 0xd9, 0xe3, 0xca, 0x97, 0xde, 0x7a,
	0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xb2, 0xb4, 0x10, 0xff, 0x02, 0x00, 0x00,
}
