// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: protos/common.proto

package common

import (
	any1 "github.com/golang/protobuf/ptypes/any"
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

type ProductType int32

const (
	ProductType_PJC ProductType = 0
)

// Enum value maps for ProductType.
var (
	ProductType_name = map[int32]string{
		0: "PJC",
	}
	ProductType_value = map[string]int32{
		"PJC": 0,
	}
)

func (x ProductType) Enum() *ProductType {
	p := new(ProductType)
	*p = x
	return p
}

func (x ProductType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_common_proto_enumTypes[0].Descriptor()
}

func (ProductType) Type() protoreflect.EnumType {
	return &file_protos_common_proto_enumTypes[0]
}

func (x ProductType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductType.Descriptor instead.
func (ProductType) EnumDescriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{0}
}

type Common struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId             string               `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	TenantUserUid         uint64               `protobuf:"varint,2,opt,name=tenant_user_uid,json=tenantUserUid,proto3" json:"tenant_user_uid,omitempty"`
	TenantUserDisplayName string               `protobuf:"bytes,3,opt,name=tenant_user_display_name,json=tenantUserDisplayName,proto3" json:"tenant_user_display_name,omitempty"`
	ProductType           ProductType          `protobuf:"varint,4,opt,name=product_type,json=productType,proto3,enum=ProductType" json:"product_type,omitempty"`
	CustomFields          map[string]*any1.Any `protobuf:"bytes,5,rep,name=custom_fields,json=customFields,proto3" json:"custom_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Common) Reset() {
	*x = Common{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Common) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Common) ProtoMessage() {}

func (x *Common) ProtoReflect() protoreflect.Message {
	mi := &file_protos_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Common.ProtoReflect.Descriptor instead.
func (*Common) Descriptor() ([]byte, []int) {
	return file_protos_common_proto_rawDescGZIP(), []int{0}
}

func (x *Common) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Common) GetTenantUserUid() uint64 {
	if x != nil {
		return x.TenantUserUid
	}
	return 0
}

func (x *Common) GetTenantUserDisplayName() string {
	if x != nil {
		return x.TenantUserDisplayName
	}
	return ""
}

func (x *Common) GetProductType() ProductType {
	if x != nil {
		return x.ProductType
	}
	return ProductType_PJC
}

func (x *Common) GetCustomFields() map[string]*any1.Any {
	if x != nil {
		return x.CustomFields
	}
	return nil
}

var File_protos_common_proto protoreflect.FileDescriptor

var file_protos_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd0, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0d,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x55, 0x0a, 0x11,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x16, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4a, 0x43, 0x10, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_common_proto_rawDescOnce sync.Once
	file_protos_common_proto_rawDescData = file_protos_common_proto_rawDesc
)

func file_protos_common_proto_rawDescGZIP() []byte {
	file_protos_common_proto_rawDescOnce.Do(func() {
		file_protos_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_common_proto_rawDescData)
	})
	return file_protos_common_proto_rawDescData
}

var file_protos_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_common_proto_goTypes = []interface{}{
	(ProductType)(0), // 0: ProductType
	(*Common)(nil),   // 1: Common
	nil,              // 2: Common.CustomFieldsEntry
	(*any1.Any)(nil), // 3: google.protobuf.Any
}
var file_protos_common_proto_depIdxs = []int32{
	0, // 0: Common.product_type:type_name -> ProductType
	2, // 1: Common.custom_fields:type_name -> Common.CustomFieldsEntry
	3, // 2: Common.CustomFieldsEntry.value:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_common_proto_init() }
func file_protos_common_proto_init() {
	if File_protos_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Common); i {
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
			RawDescriptor: file_protos_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_common_proto_goTypes,
		DependencyIndexes: file_protos_common_proto_depIdxs,
		EnumInfos:         file_protos_common_proto_enumTypes,
		MessageInfos:      file_protos_common_proto_msgTypes,
	}.Build()
	File_protos_common_proto = out.File
	file_protos_common_proto_rawDesc = nil
	file_protos_common_proto_goTypes = nil
	file_protos_common_proto_depIdxs = nil
}
