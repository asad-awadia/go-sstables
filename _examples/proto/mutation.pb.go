// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: _examples/proto/mutation.proto

package proto

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

type UpdateMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName  string `protobuf:"bytes,1,opt,name=columnName,proto3" json:"columnName,omitempty"`
	ColumnValue string `protobuf:"bytes,2,opt,name=columnValue,proto3" json:"columnValue,omitempty"`
}

func (x *UpdateMutation) Reset() {
	*x = UpdateMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file___examples_proto_mutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMutation) ProtoMessage() {}

func (x *UpdateMutation) ProtoReflect() protoreflect.Message {
	mi := &file___examples_proto_mutation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMutation.ProtoReflect.Descriptor instead.
func (*UpdateMutation) Descriptor() ([]byte, []int) {
	return file___examples_proto_mutation_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateMutation) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

func (x *UpdateMutation) GetColumnValue() string {
	if x != nil {
		return x.ColumnValue
	}
	return ""
}

type DeleteMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColumnName string `protobuf:"bytes,1,opt,name=columnName,proto3" json:"columnName,omitempty"`
}

func (x *DeleteMutation) Reset() {
	*x = DeleteMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file___examples_proto_mutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMutation) ProtoMessage() {}

func (x *DeleteMutation) ProtoReflect() protoreflect.Message {
	mi := &file___examples_proto_mutation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMutation.ProtoReflect.Descriptor instead.
func (*DeleteMutation) Descriptor() ([]byte, []int) {
	return file___examples_proto_mutation_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteMutation) GetColumnName() string {
	if x != nil {
		return x.ColumnName
	}
	return ""
}

type Mutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNumber uint64 `protobuf:"varint,1,opt,name=seqNumber,proto3" json:"seqNumber,omitempty"`
	// Types that are assignable to Mutation:
	//	*Mutation_Update
	//	*Mutation_Delete
	Mutation isMutation_Mutation `protobuf_oneof:"mutation"`
}

func (x *Mutation) Reset() {
	*x = Mutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file___examples_proto_mutation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation) ProtoMessage() {}

func (x *Mutation) ProtoReflect() protoreflect.Message {
	mi := &file___examples_proto_mutation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation.ProtoReflect.Descriptor instead.
func (*Mutation) Descriptor() ([]byte, []int) {
	return file___examples_proto_mutation_proto_rawDescGZIP(), []int{2}
}

func (x *Mutation) GetSeqNumber() uint64 {
	if x != nil {
		return x.SeqNumber
	}
	return 0
}

func (m *Mutation) GetMutation() isMutation_Mutation {
	if m != nil {
		return m.Mutation
	}
	return nil
}

func (x *Mutation) GetUpdate() *UpdateMutation {
	if x, ok := x.GetMutation().(*Mutation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *Mutation) GetDelete() *DeleteMutation {
	if x, ok := x.GetMutation().(*Mutation_Delete); ok {
		return x.Delete
	}
	return nil
}

type isMutation_Mutation interface {
	isMutation_Mutation()
}

type Mutation_Update struct {
	Update *UpdateMutation `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type Mutation_Delete struct {
	Delete *DeleteMutation `protobuf:"bytes,3,opt,name=delete,proto3,oneof"`
}

func (*Mutation_Update) isMutation_Mutation() {}

func (*Mutation_Delete) isMutation_Mutation() {}

var File___examples_proto_mutation_proto protoreflect.FileDescriptor

var file___examples_proto_mutation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x96, 0x01,
	0x0a, 0x08, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x65, 0x71, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x00, 0x52, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x6d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x6f, 0x6d, 0x61, 0x73, 0x6a, 0x75, 0x6e, 0x67, 0x62,
	0x6c, 0x75, 0x74, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file___examples_proto_mutation_proto_rawDescOnce sync.Once
	file___examples_proto_mutation_proto_rawDescData = file___examples_proto_mutation_proto_rawDesc
)

func file___examples_proto_mutation_proto_rawDescGZIP() []byte {
	file___examples_proto_mutation_proto_rawDescOnce.Do(func() {
		file___examples_proto_mutation_proto_rawDescData = protoimpl.X.CompressGZIP(file___examples_proto_mutation_proto_rawDescData)
	})
	return file___examples_proto_mutation_proto_rawDescData
}

var file___examples_proto_mutation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file___examples_proto_mutation_proto_goTypes = []interface{}{
	(*UpdateMutation)(nil), // 0: proto.UpdateMutation
	(*DeleteMutation)(nil), // 1: proto.DeleteMutation
	(*Mutation)(nil),       // 2: proto.Mutation
}
var file___examples_proto_mutation_proto_depIdxs = []int32{
	0, // 0: proto.Mutation.update:type_name -> proto.UpdateMutation
	1, // 1: proto.Mutation.delete:type_name -> proto.DeleteMutation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file___examples_proto_mutation_proto_init() }
func file___examples_proto_mutation_proto_init() {
	if File___examples_proto_mutation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___examples_proto_mutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMutation); i {
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
		file___examples_proto_mutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMutation); i {
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
		file___examples_proto_mutation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutation); i {
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
	file___examples_proto_mutation_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Mutation_Update)(nil),
		(*Mutation_Delete)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file___examples_proto_mutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file___examples_proto_mutation_proto_goTypes,
		DependencyIndexes: file___examples_proto_mutation_proto_depIdxs,
		MessageInfos:      file___examples_proto_mutation_proto_msgTypes,
	}.Build()
	File___examples_proto_mutation_proto = out.File
	file___examples_proto_mutation_proto_rawDesc = nil
	file___examples_proto_mutation_proto_goTypes = nil
	file___examples_proto_mutation_proto_depIdxs = nil
}