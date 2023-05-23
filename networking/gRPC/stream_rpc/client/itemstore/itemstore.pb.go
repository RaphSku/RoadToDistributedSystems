// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: itemstore.proto

package itemstore

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

type StorageCenter int32

const (
	StorageCenter_K   StorageCenter = 0
	StorageCenter_FRA StorageCenter = 1
	StorageCenter_HH  StorageCenter = 2
)

// Enum value maps for StorageCenter.
var (
	StorageCenter_name = map[int32]string{
		0: "K",
		1: "FRA",
		2: "HH",
	}
	StorageCenter_value = map[string]int32{
		"K":   0,
		"FRA": 1,
		"HH":  2,
	}
)

func (x StorageCenter) Enum() *StorageCenter {
	p := new(StorageCenter)
	*p = x
	return p
}

func (x StorageCenter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageCenter) Descriptor() protoreflect.EnumDescriptor {
	return file_itemstore_proto_enumTypes[0].Descriptor()
}

func (StorageCenter) Type() protoreflect.EnumType {
	return &file_itemstore_proto_enumTypes[0]
}

func (x StorageCenter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageCenter.Descriptor instead.
func (StorageCenter) EnumDescriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{0}
}

type ItemStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StorageCenter StorageCenter `protobuf:"varint,4,opt,name=storage_center,json=storageCenter,proto3,enum=main.StorageCenter" json:"storage_center,omitempty"`
	StorageLayer  int32         `protobuf:"varint,5,opt,name=storage_layer,json=storageLayer,proto3" json:"storage_layer,omitempty"`
	StorageBlock  int32         `protobuf:"varint,6,opt,name=storage_block,json=storageBlock,proto3" json:"storage_block,omitempty"`
}

func (x *ItemStoreRequest) Reset() {
	*x = ItemStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStoreRequest) ProtoMessage() {}

func (x *ItemStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_itemstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStoreRequest.ProtoReflect.Descriptor instead.
func (*ItemStoreRequest) Descriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{0}
}

func (x *ItemStoreRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ItemStoreRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemStoreRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ItemStoreRequest) GetStorageCenter() StorageCenter {
	if x != nil {
		return x.StorageCenter
	}
	return StorageCenter_K
}

func (x *ItemStoreRequest) GetStorageLayer() int32 {
	if x != nil {
		return x.StorageLayer
	}
	return 0
}

func (x *ItemStoreRequest) GetStorageBlock() int32 {
	if x != nil {
		return x.StorageBlock
	}
	return 0
}

type ItemStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ItemStoreResponse) Reset() {
	*x = ItemStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStoreResponse) ProtoMessage() {}

func (x *ItemStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStoreResponse.ProtoReflect.Descriptor instead.
func (*ItemStoreResponse) Descriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{1}
}

func (x *ItemStoreResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ItemStoreAvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageCenter StorageCenter `protobuf:"varint,1,opt,name=storage_center,json=storageCenter,proto3,enum=main.StorageCenter" json:"storage_center,omitempty"`
	StorageLayer  int32         `protobuf:"varint,2,opt,name=storage_layer,json=storageLayer,proto3" json:"storage_layer,omitempty"`
	StorageBlock  int32         `protobuf:"varint,3,opt,name=storage_block,json=storageBlock,proto3" json:"storage_block,omitempty"`
}

func (x *ItemStoreAvailabilityRequest) Reset() {
	*x = ItemStoreAvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStoreAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStoreAvailabilityRequest) ProtoMessage() {}

func (x *ItemStoreAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_itemstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStoreAvailabilityRequest.ProtoReflect.Descriptor instead.
func (*ItemStoreAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{2}
}

func (x *ItemStoreAvailabilityRequest) GetStorageCenter() StorageCenter {
	if x != nil {
		return x.StorageCenter
	}
	return StorageCenter_K
}

func (x *ItemStoreAvailabilityRequest) GetStorageLayer() int32 {
	if x != nil {
		return x.StorageLayer
	}
	return 0
}

func (x *ItemStoreAvailabilityRequest) GetStorageBlock() int32 {
	if x != nil {
		return x.StorageBlock
	}
	return 0
}

type ItemStoreAvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Availability int32 `protobuf:"varint,1,opt,name=availability,proto3" json:"availability,omitempty"`
	StorageLayer int32 `protobuf:"varint,2,opt,name=storage_layer,json=storageLayer,proto3" json:"storage_layer,omitempty"`
	StorageBlock int32 `protobuf:"varint,3,opt,name=storage_block,json=storageBlock,proto3" json:"storage_block,omitempty"`
}

func (x *ItemStoreAvailabilityResponse) Reset() {
	*x = ItemStoreAvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemStoreAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemStoreAvailabilityResponse) ProtoMessage() {}

func (x *ItemStoreAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_itemstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemStoreAvailabilityResponse.ProtoReflect.Descriptor instead.
func (*ItemStoreAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{3}
}

func (x *ItemStoreAvailabilityResponse) GetAvailability() int32 {
	if x != nil {
		return x.Availability
	}
	return 0
}

func (x *ItemStoreAvailabilityResponse) GetStorageLayer() int32 {
	if x != nil {
		return x.StorageLayer
	}
	return 0
}

func (x *ItemStoreAvailabilityResponse) GetStorageBlock() int32 {
	if x != nil {
		return x.StorageBlock
	}
	return 0
}

type StorageCenterLayoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageCenter StorageCenter `protobuf:"varint,1,opt,name=storage_center,json=storageCenter,proto3,enum=main.StorageCenter" json:"storage_center,omitempty"`
}

func (x *StorageCenterLayoutRequest) Reset() {
	*x = StorageCenterLayoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itemstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageCenterLayoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageCenterLayoutRequest) ProtoMessage() {}

func (x *StorageCenterLayoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_itemstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageCenterLayoutRequest.ProtoReflect.Descriptor instead.
func (*StorageCenterLayoutRequest) Descriptor() ([]byte, []int) {
	return file_itemstore_proto_rawDescGZIP(), []int{4}
}

func (x *StorageCenterLayoutRequest) GetStorageCenter() StorageCenter {
	if x != nil {
		return x.StorageCenter
	}
	return StorageCenter_K
}

var File_itemstore_proto protoreflect.FileDescriptor

var file_itemstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x10, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x23, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa4, 0x01,
	0x0a, 0x1c, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a,
	0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x8d, 0x01, 0x0a, 0x1d, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x58, 0x0a, 0x1a, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2a, 0x27,
	0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x05, 0x0a, 0x01, 0x4b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x52, 0x41, 0x10, 0x01, 0x12,
	0x06, 0x0a, 0x02, 0x48, 0x48, 0x10, 0x02, 0x32, 0xa6, 0x02, 0x0a, 0x10, 0x49, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x18,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x54, 0x6f, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x19, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x4f,
	0x66, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01,
	0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itemstore_proto_rawDescOnce sync.Once
	file_itemstore_proto_rawDescData = file_itemstore_proto_rawDesc
)

func file_itemstore_proto_rawDescGZIP() []byte {
	file_itemstore_proto_rawDescOnce.Do(func() {
		file_itemstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_itemstore_proto_rawDescData)
	})
	return file_itemstore_proto_rawDescData
}

var file_itemstore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_itemstore_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_itemstore_proto_goTypes = []interface{}{
	(StorageCenter)(0),                    // 0: main.StorageCenter
	(*ItemStoreRequest)(nil),              // 1: main.ItemStoreRequest
	(*ItemStoreResponse)(nil),             // 2: main.ItemStoreResponse
	(*ItemStoreAvailabilityRequest)(nil),  // 3: main.ItemStoreAvailabilityRequest
	(*ItemStoreAvailabilityResponse)(nil), // 4: main.ItemStoreAvailabilityResponse
	(*StorageCenterLayoutRequest)(nil),    // 5: main.StorageCenterLayoutRequest
}
var file_itemstore_proto_depIdxs = []int32{
	0, // 0: main.ItemStoreRequest.storage_center:type_name -> main.StorageCenter
	0, // 1: main.ItemStoreAvailabilityRequest.storage_center:type_name -> main.StorageCenter
	0, // 2: main.StorageCenterLayoutRequest.storage_center:type_name -> main.StorageCenter
	3, // 3: main.ItemStoreService.CheckStorageAvailability:input_type -> main.ItemStoreAvailabilityRequest
	1, // 4: main.ItemStoreService.DeliverItemToStorage:input_type -> main.ItemStoreRequest
	5, // 5: main.ItemStoreService.ListStorageLayoutOfCenter:input_type -> main.StorageCenterLayoutRequest
	4, // 6: main.ItemStoreService.CheckStorageAvailability:output_type -> main.ItemStoreAvailabilityResponse
	2, // 7: main.ItemStoreService.DeliverItemToStorage:output_type -> main.ItemStoreResponse
	4, // 8: main.ItemStoreService.ListStorageLayoutOfCenter:output_type -> main.ItemStoreAvailabilityResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_itemstore_proto_init() }
func file_itemstore_proto_init() {
	if File_itemstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itemstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStoreRequest); i {
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
		file_itemstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStoreResponse); i {
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
		file_itemstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStoreAvailabilityRequest); i {
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
		file_itemstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemStoreAvailabilityResponse); i {
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
		file_itemstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageCenterLayoutRequest); i {
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
			RawDescriptor: file_itemstore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_itemstore_proto_goTypes,
		DependencyIndexes: file_itemstore_proto_depIdxs,
		EnumInfos:         file_itemstore_proto_enumTypes,
		MessageInfos:      file_itemstore_proto_msgTypes,
	}.Build()
	File_itemstore_proto = out.File
	file_itemstore_proto_rawDesc = nil
	file_itemstore_proto_goTypes = nil
	file_itemstore_proto_depIdxs = nil
}