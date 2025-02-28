// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: synchronization/core/entry.proto

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

// EntryKind encodes the type of entry represented by an Entry object.
type EntryKind int32

const (
	// EntryKind_Directory indicates a directory.
	EntryKind_Directory EntryKind = 0
	// EntryKind_File indicates a regular file.
	EntryKind_File EntryKind = 1
	// EntryKind_SymbolicLink indicates a symbolic link.
	EntryKind_SymbolicLink EntryKind = 2
	// EntryKind_Untracked indicates content (or the root of content) that is
	// intentionally excluded from synchronization by Mutagen. This includes
	// explicitly ignored content, content that is ignored due to settings (such
	// as symbolic links in the "ignore" symbolic link mode), as well as content
	// types that Mutagen doesn't understand and/or have a way to propagate
	// (such as FIFOs and Unix domain sockets). This type of entry is not
	// synchronizable.
	EntryKind_Untracked EntryKind = 100
	// EntryKind_Problematic indicates content (or the root of content) that
	// would normally be synchronized, but which is currently inaccessible to
	// scanning. This includes (but is not limited to) content that is modified
	// concurrently with scanning, content that is inaccessible due to
	// permissions, content that can't be read due to filesystem errors, content
	// that cannot be properly encoded given the current settings (such as
	// absolute symbolic links found when using the "portable" symbolic link
	// mode), and content that Mutagen cannot scan or watch reliably (such as
	// directories that are also mount points). This type of entry is not
	// synchronizable.
	EntryKind_Problematic EntryKind = 101
)

// Enum value maps for EntryKind.
var (
	EntryKind_name = map[int32]string{
		0:   "Directory",
		1:   "File",
		2:   "SymbolicLink",
		100: "Untracked",
		101: "Problematic",
	}
	EntryKind_value = map[string]int32{
		"Directory":    0,
		"File":         1,
		"SymbolicLink": 2,
		"Untracked":    100,
		"Problematic":  101,
	}
)

func (x EntryKind) Enum() *EntryKind {
	p := new(EntryKind)
	*p = x
	return p
}

func (x EntryKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryKind) Descriptor() protoreflect.EnumDescriptor {
	return file_synchronization_core_entry_proto_enumTypes[0].Descriptor()
}

func (EntryKind) Type() protoreflect.EnumType {
	return &file_synchronization_core_entry_proto_enumTypes[0]
}

func (x EntryKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryKind.Descriptor instead.
func (EntryKind) EnumDescriptor() ([]byte, []int) {
	return file_synchronization_core_entry_proto_rawDescGZIP(), []int{0}
}

// Entry encodes a filesystem entry (e.g. a directory, a file, or a symbolic
// link). A nil Entry represents an absence of content. An zero-value Entry
// represents an empty Directory. Entry objects should be considered immutable
// and must not be modified.
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Kind encodes the type of filesystem entry being represented.
	Kind EntryKind `protobuf:"varint,1,opt,name=kind,proto3,enum=core.EntryKind" json:"kind,omitempty"`
	// Contents represents a directory entry's contents. It must only be non-nil
	// for directory entries.
	Contents map[string]*Entry `protobuf:"bytes,5,rep,name=contents,proto3" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Digest represents the hash of a file entry's contents. It must only be
	// non-nil for file entries.
	Digest []byte `protobuf:"bytes,8,opt,name=digest,proto3" json:"digest,omitempty"`
	// Executable indicates whether or not a file entry is marked as executable.
	// It must only be set (if appropriate) for file entries.
	Executable bool `protobuf:"varint,9,opt,name=executable,proto3" json:"executable,omitempty"`
	// Target is the symbolic link target for symbolic link entries. It must be
	// non-empty if and only if the entry is a symbolic link.
	Target string `protobuf:"bytes,12,opt,name=target,proto3" json:"target,omitempty"`
	// Problem indicates the relevant error for problematic content. It must be
	// non-empty if and only if the entry represents problematic content.
	Problem string `protobuf:"bytes,15,opt,name=problem,proto3" json:"problem,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synchronization_core_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_synchronization_core_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_synchronization_core_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetKind() EntryKind {
	if x != nil {
		return x.Kind
	}
	return EntryKind_Directory
}

func (x *Entry) GetContents() map[string]*Entry {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *Entry) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

func (x *Entry) GetExecutable() bool {
	if x != nil {
		return x.Executable
	}
	return false
}

func (x *Entry) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Entry) GetProblem() string {
	if x != nil {
		return x.Problem
	}
	return ""
}

var File_synchronization_core_entry_proto protoreflect.FileDescriptor

var file_synchronization_core_entry_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x97, 0x02, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b, 0x69, 0x6e,
	0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x1a, 0x48, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x2a, 0x56, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x0d, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x62, 0x6c, 0x65, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x10, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e,
	0x2d, 0x69, 0x6f, 0x2f, 0x6d, 0x75, 0x74, 0x61, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synchronization_core_entry_proto_rawDescOnce sync.Once
	file_synchronization_core_entry_proto_rawDescData = file_synchronization_core_entry_proto_rawDesc
)

func file_synchronization_core_entry_proto_rawDescGZIP() []byte {
	file_synchronization_core_entry_proto_rawDescOnce.Do(func() {
		file_synchronization_core_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_synchronization_core_entry_proto_rawDescData)
	})
	return file_synchronization_core_entry_proto_rawDescData
}

var file_synchronization_core_entry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_synchronization_core_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_synchronization_core_entry_proto_goTypes = []interface{}{
	(EntryKind)(0), // 0: core.EntryKind
	(*Entry)(nil),  // 1: core.Entry
	nil,            // 2: core.Entry.ContentsEntry
}
var file_synchronization_core_entry_proto_depIdxs = []int32{
	0, // 0: core.Entry.kind:type_name -> core.EntryKind
	2, // 1: core.Entry.contents:type_name -> core.Entry.ContentsEntry
	1, // 2: core.Entry.ContentsEntry.value:type_name -> core.Entry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_synchronization_core_entry_proto_init() }
func file_synchronization_core_entry_proto_init() {
	if File_synchronization_core_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synchronization_core_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_synchronization_core_entry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_synchronization_core_entry_proto_goTypes,
		DependencyIndexes: file_synchronization_core_entry_proto_depIdxs,
		EnumInfos:         file_synchronization_core_entry_proto_enumTypes,
		MessageInfos:      file_synchronization_core_entry_proto_msgTypes,
	}.Build()
	File_synchronization_core_entry_proto = out.File
	file_synchronization_core_entry_proto_rawDesc = nil
	file_synchronization_core_entry_proto_goTypes = nil
	file_synchronization_core_entry_proto_depIdxs = nil
}
