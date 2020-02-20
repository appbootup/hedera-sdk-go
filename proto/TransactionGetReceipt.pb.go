// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.19.0-devel
// 	protoc        v3.11.3
// source: proto/TransactionGetReceipt.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(19 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 19)
)

// Get the receipt of a transaction, given its transaction ID. Once a transaction reaches consensus, then information about whether it succeeded or failed will be available until the end of the receipt period.  Before and after the receipt period, and for a transaction that was never submitted, the receipt is unknown.  This query is free (the payment field is left empty). No State proof is available for this response
type TransactionGetReceiptQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header        *QueryHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`               // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	TransactionID *TransactionID `protobuf:"bytes,2,opt,name=transactionID,proto3" json:"transactionID,omitempty"` // The ID of the transaction for which the receipt is requested.
}

func (x *TransactionGetReceiptQuery) Reset() {
	*x = TransactionGetReceiptQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TransactionGetReceipt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetReceiptQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetReceiptQuery) ProtoMessage() {}

func (x *TransactionGetReceiptQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TransactionGetReceipt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetReceiptQuery.ProtoReflect.Descriptor instead.
func (*TransactionGetReceiptQuery) Descriptor() ([]byte, []int) {
	return file_proto_TransactionGetReceipt_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionGetReceiptQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetReceiptQuery) GetTransactionID() *TransactionID {
	if x != nil {
		return x.TransactionID
	}
	return nil
}

// Response when the client sends the node TransactionGetReceiptQuery. If it created a new entity (account, file, or smart contract instance) then one of the three ID fields will be filled in with the ID of the new entity. Sometimes a single transaction will create more than one new entity, such as when a new contract instance is created, and this also creates the new account that it owned by that instance. No State proof is available for this response
type TransactionGetReceiptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *ResponseHeader     `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`   //Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	Receipt *TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"` // The receipt, indicating it reached consensus (and whether it succeeded or failed) or is currently unknown (because it hasn't reached consensus yet, or the transaction has expired already), and including the ID of any new account/file/instance created by that transaction.
}

func (x *TransactionGetReceiptResponse) Reset() {
	*x = TransactionGetReceiptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_TransactionGetReceipt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionGetReceiptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionGetReceiptResponse) ProtoMessage() {}

func (x *TransactionGetReceiptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_TransactionGetReceipt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionGetReceiptResponse.ProtoReflect.Descriptor instead.
func (*TransactionGetReceiptResponse) Descriptor() ([]byte, []int) {
	return file_proto_TransactionGetReceipt_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionGetReceiptResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TransactionGetReceiptResponse) GetReceipt() *TransactionReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

var File_proto_TransactionGetReceipt_proto protoreflect.FileDescriptor

var file_proto_TransactionGetReceipt_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x52,
	0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x83,
	0x01, 0x0a, 0x1d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x33, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x42, 0x48, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x68, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_TransactionGetReceipt_proto_rawDescOnce sync.Once
	file_proto_TransactionGetReceipt_proto_rawDescData = file_proto_TransactionGetReceipt_proto_rawDesc
)

func file_proto_TransactionGetReceipt_proto_rawDescGZIP() []byte {
	file_proto_TransactionGetReceipt_proto_rawDescOnce.Do(func() {
		file_proto_TransactionGetReceipt_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_TransactionGetReceipt_proto_rawDescData)
	})
	return file_proto_TransactionGetReceipt_proto_rawDescData
}

var file_proto_TransactionGetReceipt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_TransactionGetReceipt_proto_goTypes = []interface{}{
	(*TransactionGetReceiptQuery)(nil),    // 0: proto.TransactionGetReceiptQuery
	(*TransactionGetReceiptResponse)(nil), // 1: proto.TransactionGetReceiptResponse
	(*QueryHeader)(nil),                   // 2: proto.QueryHeader
	(*TransactionID)(nil),                 // 3: proto.TransactionID
	(*ResponseHeader)(nil),                // 4: proto.ResponseHeader
	(*TransactionReceipt)(nil),            // 5: proto.TransactionReceipt
}
var file_proto_TransactionGetReceipt_proto_depIdxs = []int32{
	2, // 0: proto.TransactionGetReceiptQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.TransactionGetReceiptQuery.transactionID:type_name -> proto.TransactionID
	4, // 2: proto.TransactionGetReceiptResponse.header:type_name -> proto.ResponseHeader
	5, // 3: proto.TransactionGetReceiptResponse.receipt:type_name -> proto.TransactionReceipt
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_TransactionGetReceipt_proto_init() }
func file_proto_TransactionGetReceipt_proto_init() {
	if File_proto_TransactionGetReceipt_proto != nil {
		return
	}
	file_proto_TransactionReceipt_proto_init()
	file_proto_BasicTypes_proto_init()
	file_proto_QueryHeader_proto_init()
	file_proto_ResponseHeader_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_TransactionGetReceipt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetReceiptQuery); i {
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
		file_proto_TransactionGetReceipt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionGetReceiptResponse); i {
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
			RawDescriptor: file_proto_TransactionGetReceipt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_TransactionGetReceipt_proto_goTypes,
		DependencyIndexes: file_proto_TransactionGetReceipt_proto_depIdxs,
		MessageInfos:      file_proto_TransactionGetReceipt_proto_msgTypes,
	}.Build()
	File_proto_TransactionGetReceipt_proto = out.File
	file_proto_TransactionGetReceipt_proto_rawDesc = nil
	file_proto_TransactionGetReceipt_proto_goTypes = nil
	file_proto_TransactionGetReceipt_proto_depIdxs = nil
}
