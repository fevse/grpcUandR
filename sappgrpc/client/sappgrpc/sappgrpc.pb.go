// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: sappgrpc.proto

package sappgrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Product) Reset() {
	*x = Product{}
	mi := &file_sappgrpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_sappgrpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_sappgrpc_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProductID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductID) Reset() {
	*x = ProductID{}
	mi := &file_sappgrpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductID) ProtoMessage() {}

func (x *ProductID) ProtoReflect() protoreflect.Message {
	mi := &file_sappgrpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductID.ProtoReflect.Descriptor instead.
func (*ProductID) Descriptor() ([]byte, []int) {
	return file_sappgrpc_proto_rawDescGZIP(), []int{1}
}

func (x *ProductID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_sappgrpc_proto protoreflect.FileDescriptor

var file_sappgrpc_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x73, 0x61, 0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x61, 0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x22, 0x4f, 0x0a, 0x07, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x79,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a,
	0x0a, 0x61, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x61,
	0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x13,
	0x2e, 0x73, 0x61, 0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x13, 0x2e, 0x73, 0x61, 0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x73, 0x61, 0x70, 0x70, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x73,
	0x61, 0x70, 0x70, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_sappgrpc_proto_rawDescOnce sync.Once
	file_sappgrpc_proto_rawDescData []byte
)

func file_sappgrpc_proto_rawDescGZIP() []byte {
	file_sappgrpc_proto_rawDescOnce.Do(func() {
		file_sappgrpc_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_sappgrpc_proto_rawDesc), len(file_sappgrpc_proto_rawDesc)))
	})
	return file_sappgrpc_proto_rawDescData
}

var file_sappgrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sappgrpc_proto_goTypes = []any{
	(*Product)(nil),   // 0: sappgrpc.Product
	(*ProductID)(nil), // 1: sappgrpc.ProductID
}
var file_sappgrpc_proto_depIdxs = []int32{
	0, // 0: sappgrpc.ProductInfo.addProduct:input_type -> sappgrpc.Product
	1, // 1: sappgrpc.ProductInfo.getProduct:input_type -> sappgrpc.ProductID
	1, // 2: sappgrpc.ProductInfo.addProduct:output_type -> sappgrpc.ProductID
	0, // 3: sappgrpc.ProductInfo.getProduct:output_type -> sappgrpc.Product
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sappgrpc_proto_init() }
func file_sappgrpc_proto_init() {
	if File_sappgrpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sappgrpc_proto_rawDesc), len(file_sappgrpc_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sappgrpc_proto_goTypes,
		DependencyIndexes: file_sappgrpc_proto_depIdxs,
		MessageInfos:      file_sappgrpc_proto_msgTypes,
	}.Build()
	File_sappgrpc_proto = out.File
	file_sappgrpc_proto_goTypes = nil
	file_sappgrpc_proto_depIdxs = nil
}
