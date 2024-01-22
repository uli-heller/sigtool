// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: internal/pb/hdr.proto

package pb

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

//
// Every encrypted file starts with a header describing the
// Block Size, Salt, Recipient keys etc. Header represents a
// decoded version of this information. It is encoded in
// protobuf format before writing to disk.
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChunkSize uint32        `protobuf:"varint,1,opt,name=chunk_size,json=chunkSize,proto3" json:"chunk_size,omitempty"` // encryption block size
	Salt      []byte        `protobuf:"bytes,2,opt,name=salt,proto3" json:"salt,omitempty"`                             // master salt (nonces are derived from this)
	Pk        []byte        `protobuf:"bytes,3,opt,name=pk,proto3" json:"pk,omitempty"`                                 // ephemeral curve PK
	Sender    []byte        `protobuf:"bytes,4,opt,name=sender,proto3" json:"sender,omitempty"`                         // sender signed artifacts
	Keys      []*WrappedKey `protobuf:"bytes,5,rep,name=keys,proto3" json:"keys,omitempty"`                             // list of wrapped receiver blocks
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_hdr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_hdr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_internal_pb_hdr_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetChunkSize() uint32 {
	if x != nil {
		return x.ChunkSize
	}
	return 0
}

func (x *Header) GetSalt() []byte {
	if x != nil {
		return x.Salt
	}
	return nil
}

func (x *Header) GetPk() []byte {
	if x != nil {
		return x.Pk
	}
	return nil
}

func (x *Header) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *Header) GetKeys() []*WrappedKey {
	if x != nil {
		return x.Keys
	}
	return nil
}

//
// A file encryption key is wrapped by a recipient specific public
// key. WrappedKey describes such a wrapped key.
type WrappedKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DKey  []byte `protobuf:"bytes,1,opt,name=d_key,json=dKey,proto3" json:"d_key,omitempty"` // encrypted data key
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`           // nonce used for encryption
}

func (x *WrappedKey) Reset() {
	*x = WrappedKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_hdr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrappedKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrappedKey) ProtoMessage() {}

func (x *WrappedKey) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_hdr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrappedKey.ProtoReflect.Descriptor instead.
func (*WrappedKey) Descriptor() ([]byte, []int) {
	return file_internal_pb_hdr_proto_rawDescGZIP(), []int{1}
}

func (x *WrappedKey) GetDKey() []byte {
	if x != nil {
		return x.DKey
	}
	return nil
}

func (x *WrappedKey) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

var File_internal_pb_hdr_proto protoreflect.FileDescriptor

var file_internal_pb_hdr_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x64,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22,
	0x38, 0x0a, 0x0b, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x13,
	0x0a, 0x05, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_hdr_proto_rawDescOnce sync.Once
	file_internal_pb_hdr_proto_rawDescData = file_internal_pb_hdr_proto_rawDesc
)

func file_internal_pb_hdr_proto_rawDescGZIP() []byte {
	file_internal_pb_hdr_proto_rawDescOnce.Do(func() {
		file_internal_pb_hdr_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_hdr_proto_rawDescData)
	})
	return file_internal_pb_hdr_proto_rawDescData
}

var file_internal_pb_hdr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_pb_hdr_proto_goTypes = []interface{}{
	(*Header)(nil),     // 0: header
	(*WrappedKey)(nil), // 1: wrapped_key
}
var file_internal_pb_hdr_proto_depIdxs = []int32{
	1, // 0: header.keys:type_name -> wrapped_key
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_pb_hdr_proto_init() }
func file_internal_pb_hdr_proto_init() {
	if File_internal_pb_hdr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_hdr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_internal_pb_hdr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrappedKey); i {
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
			RawDescriptor: file_internal_pb_hdr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_hdr_proto_goTypes,
		DependencyIndexes: file_internal_pb_hdr_proto_depIdxs,
		MessageInfos:      file_internal_pb_hdr_proto_msgTypes,
	}.Build()
	File_internal_pb_hdr_proto = out.File
	file_internal_pb_hdr_proto_rawDesc = nil
	file_internal_pb_hdr_proto_goTypes = nil
	file_internal_pb_hdr_proto_depIdxs = nil
}
