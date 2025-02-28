// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: synchronization/core/snapshot.proto

package core

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

// Snapshot bundles a filesystem content snapshot with associated metadata.
// Snapshot objects should be considered immutable and must not be modified.
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Content is the filesystem content at the snapshot root. It may be nil to
	// indicate an absence of content.
	Content *Entry `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// PreservesExecutability indicates whether or not the associated filesystem
	// preserves POSIX executability bits.
	PreservesExecutability bool `protobuf:"varint,2,opt,name=preservesExecutability,proto3" json:"preservesExecutability,omitempty"`
	// DecomposesUnicode indicates whether or not the associated filesystem
	// decomposes Unicode names.
	DecomposesUnicode bool `protobuf:"varint,3,opt,name=decomposesUnicode,proto3" json:"decomposesUnicode,omitempty"`
	// Directories is the number of synchronizable directory entries contained
	// in the snapshot.
	Directories uint64 `protobuf:"varint,4,opt,name=directories,proto3" json:"directories,omitempty"`
	// Files is the number of synchronizable file entries contained in the
	// snapshot.
	Files uint64 `protobuf:"varint,5,opt,name=files,proto3" json:"files,omitempty"`
	// SymbolicLinks is the number of synchronizable symbolic link entries
	// contained in the snapshot.
	SymbolicLinks uint64 `protobuf:"varint,6,opt,name=symbolicLinks,proto3" json:"symbolicLinks,omitempty"`
	// TotalFileSize is the total size of all synchronizable files referenced by
	// the snapshot.
	TotalFileSize uint64 `protobuf:"varint,7,opt,name=totalFileSize,proto3" json:"totalFileSize,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_core_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_core_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_synchronization_core_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *Snapshot) GetContent() *Entry {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Snapshot) GetPreservesExecutability() bool {
	if x != nil {
		return x.PreservesExecutability
	}
	return false
}

func (x *Snapshot) GetDecomposesUnicode() bool {
	if x != nil {
		return x.DecomposesUnicode
	}
	return false
}

func (x *Snapshot) GetDirectories() uint64 {
	if x != nil {
		return x.Directories
	}
	return 0
}

func (x *Snapshot) GetFiles() uint64 {
	if x != nil {
		return x.Files
	}
	return 0
}

func (x *Snapshot) GetSymbolicLinks() uint64 {
	if x != nil {
		return x.SymbolicLinks
	}
	return 0
}

func (x *Snapshot) GetTotalFileSize() uint64 {
	if x != nil {
		return x.TotalFileSize
	}
	return 0
}

var File_synchronization_core_snapshot_proto protoreflect.FileDescriptor

var file_synchronization_core_snapshot_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x20, 0x73, 0x79, 0x6e,
	0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02,
	0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x16, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x73, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x73, 0x55, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x73,
	0x55, 0x6e, 0x69, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x69, 0x63,
	0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65,
	0x6e, 0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_snapshot_proto_rawDescOnce sync.Once
	file_synchronization_core_snapshot_proto_rawDescData = file_synchronization_core_snapshot_proto_rawDesc
)

func file_synchronization_core_snapshot_proto_rawDescGZIP() []byte {
	file_synchronization_core_snapshot_proto_rawDescOnce.Do(func() {
		file_synchronization_core_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_snapshot_proto_rawDescData)
	})
	return file_synchronization_core_snapshot_proto_rawDescData
}

var file_synchronization_core_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_synchronization_core_snapshot_proto_goTypes = []interface{}{
	(*Snapshot)(nil), // 0: core.Snapshot
	(*Entry)(nil),    // 1: core.Entry
}
var file_synchronization_core_snapshot_proto_depIdxs = []int32{
	1, // 0: core.Snapshot.content:type_name -> core.Entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_synchronization_core_snapshot_proto_init() }
func file_synchronization_core_snapshot_proto_init() {
	if File_synchronization_core_snapshot_proto != nil {
		return
	}
	file_synchronization_core_entry_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_synchronization_core_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
			RawDescriptor: file_synchronization_core_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_snapshot_proto_goTypes,
		DependencyIndexes: file_synchronization_core_snapshot_proto_depIdxs,
		MessageInfos:      file_synchronization_core_snapshot_proto_msgTypes,
	}.Build()
	File_synchronization_core_snapshot_proto = out.File
	file_synchronization_core_snapshot_proto_rawDesc = nil
	file_synchronization_core_snapshot_proto_goTypes = nil
	file_synchronization_core_snapshot_proto_depIdxs = nil
}
