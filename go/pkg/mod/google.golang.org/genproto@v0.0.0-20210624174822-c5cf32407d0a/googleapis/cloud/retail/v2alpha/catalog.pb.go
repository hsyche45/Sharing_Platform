// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/retail/v2alpha/catalog.proto

package retail

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configures what level the product should be uploaded with regards to
// how users will be send events and how predictions will be made.
type ProductLevelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of [Product][google.cloud.retail.v2alpha.Product]s allowed to be
	// ingested into the catalog. Acceptable values are:
	//
	// * `primary` (default): You can only ingest
	// [Product.Type.PRIMARY][google.cloud.retail.v2alpha.Product.Type.PRIMARY]
	//   [Product][google.cloud.retail.v2alpha.Product]s. This means
	//   [Product.primary_product_id][google.cloud.retail.v2alpha.Product.primary_product_id]
	//   can only be empty or set to the same value as
	//   [Product.id][google.cloud.retail.v2alpha.Product.id].
	// * `variant`: You can only ingest
	// [Product.Type.VARIANT][google.cloud.retail.v2alpha.Product.Type.VARIANT]
	// [Product][google.cloud.retail.v2alpha.Product]s.
	//   This means
	//   [Product.primary_product_id][google.cloud.retail.v2alpha.Product.primary_product_id]
	//   cannot be empty.
	//
	// If this field is set to an invalid value other than these, an
	// INVALID_ARGUMENT error is returned.
	//
	// If this field is `variant` and
	// [merchant_center_product_id_field][google.cloud.retail.v2alpha.ProductLevelConfig.merchant_center_product_id_field]
	// is `itemGroupId`, an INVALID_ARGUMENT error is returned.
	//
	// See [Using catalog
	// levels](/retail/recommendations-ai/docs/catalog#catalog-levels) for more
	// details.
	IngestionProductType string `protobuf:"bytes,1,opt,name=ingestion_product_type,json=ingestionProductType,proto3" json:"ingestion_product_type,omitempty"`
	// Which field of [Merchant Center
	// Product](/bigquery-transfer/docs/merchant-center-products-schema) should be
	// imported as [Product.id][google.cloud.retail.v2alpha.Product.id].
	// Acceptable values are:
	//
	// * `offerId` (default): Import `offerId` as the product ID.
	// * `itemGroupId`: Import `itemGroupId` as the product ID. Notice that Retail
	//   API will choose one item from the ones with the same `itemGroupId`, and
	//   use it to represent the item group.
	//
	// If this field is set to an invalid value other than these, an
	// INVALID_ARGUMENT error is returned.
	//
	// If this field is `itemGroupId` and
	// [ingestion_product_type][google.cloud.retail.v2alpha.ProductLevelConfig.ingestion_product_type]
	// is `variant`, an INVALID_ARGUMENT error is returned.
	//
	// See [Using catalog
	// levels](/retail/recommendations-ai/docs/catalog#catalog-levels) for more
	// details.
	MerchantCenterProductIdField string `protobuf:"bytes,2,opt,name=merchant_center_product_id_field,json=merchantCenterProductIdField,proto3" json:"merchant_center_product_id_field,omitempty"`
}

func (x *ProductLevelConfig) Reset() {
	*x = ProductLevelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductLevelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductLevelConfig) ProtoMessage() {}

func (x *ProductLevelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductLevelConfig.ProtoReflect.Descriptor instead.
func (*ProductLevelConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_catalog_proto_rawDescGZIP(), []int{0}
}

func (x *ProductLevelConfig) GetIngestionProductType() string {
	if x != nil {
		return x.IngestionProductType
	}
	return ""
}

func (x *ProductLevelConfig) GetMerchantCenterProductIdField() string {
	if x != nil {
		return x.MerchantCenterProductIdField
	}
	return ""
}

// The catalog configuration.
type Catalog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The fully qualified resource name of the catalog.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Immutable. The catalog display name.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The product level configuration.
	ProductLevelConfig *ProductLevelConfig `protobuf:"bytes,4,opt,name=product_level_config,json=productLevelConfig,proto3" json:"product_level_config,omitempty"`
}

func (x *Catalog) Reset() {
	*x = Catalog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Catalog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Catalog) ProtoMessage() {}

func (x *Catalog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Catalog.ProtoReflect.Descriptor instead.
func (*Catalog) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2alpha_catalog_proto_rawDescGZIP(), []int{1}
}

func (x *Catalog) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Catalog) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Catalog) GetProductLevelConfig() *ProductLevelConfig {
	if x != nil {
		return x.ProductLevelConfig
	}
	return nil
}

var File_google_cloud_retail_v2alpha_catalog_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2alpha_catalog_proto_rawDesc = []byte{
	0x0a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x6e, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x46, 0x0a, 0x20, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x6d, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x6e, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x07, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x66, 0x0a, 0x14, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x12,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x3a, 0x5e, 0xea, 0x41, 0x5b, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x7d, 0x42, 0xda, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x3b, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41,
	0x49, 0x4c, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02,
	0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2alpha_catalog_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2alpha_catalog_proto_rawDescData = file_google_cloud_retail_v2alpha_catalog_proto_rawDesc
)

func file_google_cloud_retail_v2alpha_catalog_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2alpha_catalog_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2alpha_catalog_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2alpha_catalog_proto_rawDescData)
	})
	return file_google_cloud_retail_v2alpha_catalog_proto_rawDescData
}

var file_google_cloud_retail_v2alpha_catalog_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_retail_v2alpha_catalog_proto_goTypes = []interface{}{
	(*ProductLevelConfig)(nil), // 0: google.cloud.retail.v2alpha.ProductLevelConfig
	(*Catalog)(nil),            // 1: google.cloud.retail.v2alpha.Catalog
}
var file_google_cloud_retail_v2alpha_catalog_proto_depIdxs = []int32{
	0, // 0: google.cloud.retail.v2alpha.Catalog.product_level_config:type_name -> google.cloud.retail.v2alpha.ProductLevelConfig
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2alpha_catalog_proto_init() }
func file_google_cloud_retail_v2alpha_catalog_proto_init() {
	if File_google_cloud_retail_v2alpha_catalog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductLevelConfig); i {
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
		file_google_cloud_retail_v2alpha_catalog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Catalog); i {
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
			RawDescriptor: file_google_cloud_retail_v2alpha_catalog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2alpha_catalog_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2alpha_catalog_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2alpha_catalog_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2alpha_catalog_proto = out.File
	file_google_cloud_retail_v2alpha_catalog_proto_rawDesc = nil
	file_google_cloud_retail_v2alpha_catalog_proto_goTypes = nil
	file_google_cloud_retail_v2alpha_catalog_proto_depIdxs = nil
}
