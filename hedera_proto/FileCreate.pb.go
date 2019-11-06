// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FileCreate.proto

package hedera_proto

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

// Create a new file, containing the given contents.  It is referenced by its FileID, and does not have a filename, so it is important to get the FileID. After the file is created, the FileID for it can be found in the receipt, or retrieved with a GetByKey query, or by asking for a Record of the transaction to be created, and retrieving that.
//
// The file contains the given contents (possibly empty). The file will automatically disappear at the fileExpirationTime, unless its expiration is extended by another transaction before that time. If the file is deleted, then its contents will become empty and it will be marked as deleted until it expires, and then it will cease to exist. See FileGetInfoQuery for more information about files.
//
// The keys field is a list of keys. All the keys on the list must sign to create or modify a file, but only one of them needs to sign in order to delete the file.  Each of those "keys" may itself be threshold key containing other keys (including other threshold keys). In other words, the behavior is an AND for create/modify, OR for delete. This is useful for acting as a revocation server. If it is desired to have the behavior be AND for all 3 operations (or OR for all 3), then the list should have only a single Key, which is a threshold key, with N=1 for OR, N=M for AND.
//
// If a file is created without ANY keys in the keys field, the file is immutable ONLY the expirationTime of the file can be changed using FileUpdate API. The file contents or its keys cannot be changed.
//
// An entity (account, file, or smart contract instance) must be created in a particular realm. If the realmID is left null, then a new realm will be created with the given admin key. If a new realm has a null adminKey, then anyone can create/modify/delete entities in that realm. But if an admin key is given, then any transaction to create/modify/delete an entity in that realm must be signed by that key, though anyone can still call functions on smart contract instances that exist in that realm. A realm ceases to exist when everything within it has expired and no longer exists.
//
// The current API ignores shardID, realmID, and newRealmAdminKey, and creates everything in shard 0 and realm 0, with a null key. Future versions of the API will support multiple realms and multiple shards.
type FileCreateTransactionBody struct {
	ExpirationTime       *Timestamp `protobuf:"bytes,2,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	Keys                 *KeyList   `protobuf:"bytes,3,opt,name=keys,proto3" json:"keys,omitempty"`
	Contents             []byte     `protobuf:"bytes,4,opt,name=contents,proto3" json:"contents,omitempty"`
	ShardID              *ShardID   `protobuf:"bytes,5,opt,name=shardID,proto3" json:"shardID,omitempty"`
	RealmID              *RealmID   `protobuf:"bytes,6,opt,name=realmID,proto3" json:"realmID,omitempty"`
	NewRealmAdminKey     *Key       `protobuf:"bytes,7,opt,name=newRealmAdminKey,proto3" json:"newRealmAdminKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FileCreateTransactionBody) Reset()         { *m = FileCreateTransactionBody{} }
func (m *FileCreateTransactionBody) String() string { return proto.CompactTextString(m) }
func (*FileCreateTransactionBody) ProtoMessage()    {}
func (*FileCreateTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce7f6c7fa88f371, []int{0}
}

func (m *FileCreateTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileCreateTransactionBody.Unmarshal(m, b)
}
func (m *FileCreateTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileCreateTransactionBody.Marshal(b, m, deterministic)
}
func (m *FileCreateTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileCreateTransactionBody.Merge(m, src)
}
func (m *FileCreateTransactionBody) XXX_Size() int {
	return xxx_messageInfo_FileCreateTransactionBody.Size(m)
}
func (m *FileCreateTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_FileCreateTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_FileCreateTransactionBody proto.InternalMessageInfo

func (m *FileCreateTransactionBody) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *FileCreateTransactionBody) GetKeys() *KeyList {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *FileCreateTransactionBody) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *FileCreateTransactionBody) GetShardID() *ShardID {
	if m != nil {
		return m.ShardID
	}
	return nil
}

func (m *FileCreateTransactionBody) GetRealmID() *RealmID {
	if m != nil {
		return m.RealmID
	}
	return nil
}

func (m *FileCreateTransactionBody) GetNewRealmAdminKey() *Key {
	if m != nil {
		return m.NewRealmAdminKey
	}
	return nil
}

func init() {
	proto.RegisterType((*FileCreateTransactionBody)(nil), "proto.FileCreateTransactionBody")
}

func init() { proto.RegisterFile("FileCreate.proto", fileDescriptor_dce7f6c7fa88f371) }

var fileDescriptor_dce7f6c7fa88f371 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0xd2, 0x1f, 0x64, 0xaa, 0x12, 0x79, 0x32, 0x99, 0xaa, 0x4c, 0x99, 0x32, 0x00,
	0x42, 0xac, 0x84, 0x0a, 0xa9, 0x2a, 0x03, 0x32, 0x99, 0x58, 0xd0, 0x25, 0xb9, 0x22, 0x86, 0xc6,
	0xb6, 0x6c, 0x0b, 0xc8, 0xe3, 0xf0, 0xa6, 0x28, 0x76, 0x7f, 0x54, 0x3a, 0x5d, 0xf9, 0x3b, 0xdf,
	0x39, 0x32, 0x89, 0x1f, 0xc4, 0x1a, 0xef, 0x0d, 0x82, 0xc3, 0x5c, 0x1b, 0xe5, 0x14, 0x1d, 0xf9,
	0x93, 0xc4, 0x05, 0x58, 0x51, 0x95, 0x9d, 0x46, 0x1b, 0x82, 0xe4, 0xbc, 0x14, 0x2d, 0x5a, 0x07,
	0xad, 0x0e, 0x20, 0xfd, 0x1d, 0x90, 0x8b, 0x7d, 0xbd, 0x34, 0x20, 0x2d, 0x54, 0x4e, 0x28, 0x59,
	0xa8, 0xba, 0xa3, 0xb7, 0x64, 0x86, 0x3f, 0x5a, 0x18, 0xe8, 0x49, 0x5f, 0x65, 0x83, 0x79, 0x94,
	0x9d, 0x5d, 0xc6, 0xa1, 0x9d, 0xef, 0xd6, 0xf8, 0x3f, 0x8f, 0xa6, 0x64, 0xf8, 0x89, 0x9d, 0x65,
	0x27, 0xde, 0x9f, 0x6d, 0xfc, 0x15, 0x76, 0x8f, 0xc2, 0x3a, 0xee, 0x33, 0x9a, 0x90, 0xd3, 0x4a,
	0x49, 0x87, 0xd2, 0x59, 0x36, 0x9c, 0x47, 0xd9, 0x94, 0xef, 0xde, 0x34, 0x23, 0x13, 0xdb, 0x80,
	0xa9, 0x97, 0x0b, 0x36, 0x3a, 0x98, 0x78, 0x0e, 0x94, 0x6f, 0xe3, 0xde, 0x34, 0x08, 0xeb, 0x76,
	0xb9, 0x60, 0xe3, 0x03, 0x93, 0x07, 0xca, 0xb7, 0x31, 0xbd, 0x21, 0xb1, 0xc4, 0x6f, 0x8f, 0xef,
	0xea, 0x56, 0xc8, 0x15, 0x76, 0x6c, 0xe2, 0x2b, 0x64, 0xff, 0x3f, 0x7e, 0xe4, 0x14, 0xd7, 0x24,
	0xad, 0x54, 0x9b, 0x37, 0x58, 0xa3, 0x81, 0x06, 0x6c, 0xf3, 0x6e, 0x40, 0x37, 0x39, 0x68, 0xb1,
	0xa9, 0x7d, 0xc0, 0x17, 0x3c, 0x45, 0x2f, 0xd3, 0x60, 0xbc, 0x7a, 0xf8, 0x36, 0xf6, 0xe7, 0xea,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x95, 0xa3, 0xff, 0x9e, 0x01, 0x00, 0x00,
}