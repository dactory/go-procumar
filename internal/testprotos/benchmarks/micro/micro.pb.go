// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/benchmarks/micro/micro.proto

package micro

import (
	protoreflect "github.com/dactory/go-procumar/reflect/protoreflect"
	protoimpl "github.com/dactory/go-procumar/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type SixteenRequired struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	F1  *int32 `protobuf:"varint,1,req,name=f1" json:"f1,omitempty"`
	F2  *int32 `protobuf:"varint,2,req,name=f2" json:"f2,omitempty"`
	F3  *int32 `protobuf:"varint,3,req,name=f3" json:"f3,omitempty"`
	F4  *int32 `protobuf:"varint,4,req,name=f4" json:"f4,omitempty"`
	F5  *int32 `protobuf:"varint,5,req,name=f5" json:"f5,omitempty"`
	F6  *int32 `protobuf:"varint,6,req,name=f6" json:"f6,omitempty"`
	F7  *int32 `protobuf:"varint,7,req,name=f7" json:"f7,omitempty"`
	F8  *int32 `protobuf:"varint,8,req,name=f8" json:"f8,omitempty"`
	F9  *int32 `protobuf:"varint,9,req,name=f9" json:"f9,omitempty"`
	F10 *int32 `protobuf:"varint,10,req,name=f10" json:"f10,omitempty"`
	F11 *int32 `protobuf:"varint,11,req,name=f11" json:"f11,omitempty"`
	F12 *int32 `protobuf:"varint,12,req,name=f12" json:"f12,omitempty"`
	F13 *int32 `protobuf:"varint,13,req,name=f13" json:"f13,omitempty"`
	F14 *int32 `protobuf:"varint,14,req,name=f14" json:"f14,omitempty"`
	F15 *int32 `protobuf:"varint,15,req,name=f15" json:"f15,omitempty"`
	F16 *int32 `protobuf:"varint,16,req,name=f16" json:"f16,omitempty"`
}

func (x *SixteenRequired) Reset() {
	*x = SixteenRequired{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_benchmarks_micro_micro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SixteenRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SixteenRequired) ProtoMessage() {}

func (x *SixteenRequired) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_benchmarks_micro_micro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SixteenRequired.ProtoReflect.Descriptor instead.
func (*SixteenRequired) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_benchmarks_micro_micro_proto_rawDescGZIP(), []int{0}
}

func (x *SixteenRequired) GetF1() int32 {
	if x != nil && x.F1 != nil {
		return *x.F1
	}
	return 0
}

func (x *SixteenRequired) GetF2() int32 {
	if x != nil && x.F2 != nil {
		return *x.F2
	}
	return 0
}

func (x *SixteenRequired) GetF3() int32 {
	if x != nil && x.F3 != nil {
		return *x.F3
	}
	return 0
}

func (x *SixteenRequired) GetF4() int32 {
	if x != nil && x.F4 != nil {
		return *x.F4
	}
	return 0
}

func (x *SixteenRequired) GetF5() int32 {
	if x != nil && x.F5 != nil {
		return *x.F5
	}
	return 0
}

func (x *SixteenRequired) GetF6() int32 {
	if x != nil && x.F6 != nil {
		return *x.F6
	}
	return 0
}

func (x *SixteenRequired) GetF7() int32 {
	if x != nil && x.F7 != nil {
		return *x.F7
	}
	return 0
}

func (x *SixteenRequired) GetF8() int32 {
	if x != nil && x.F8 != nil {
		return *x.F8
	}
	return 0
}

func (x *SixteenRequired) GetF9() int32 {
	if x != nil && x.F9 != nil {
		return *x.F9
	}
	return 0
}

func (x *SixteenRequired) GetF10() int32 {
	if x != nil && x.F10 != nil {
		return *x.F10
	}
	return 0
}

func (x *SixteenRequired) GetF11() int32 {
	if x != nil && x.F11 != nil {
		return *x.F11
	}
	return 0
}

func (x *SixteenRequired) GetF12() int32 {
	if x != nil && x.F12 != nil {
		return *x.F12
	}
	return 0
}

func (x *SixteenRequired) GetF13() int32 {
	if x != nil && x.F13 != nil {
		return *x.F13
	}
	return 0
}

func (x *SixteenRequired) GetF14() int32 {
	if x != nil && x.F14 != nil {
		return *x.F14
	}
	return 0
}

func (x *SixteenRequired) GetF15() int32 {
	if x != nil && x.F15 != nil {
		return *x.F15
	}
	return 0
}

func (x *SixteenRequired) GetF16() int32 {
	if x != nil && x.F16 != nil {
		return *x.F16
	}
	return 0
}

var File_internal_testprotos_benchmarks_micro_micro_proto protoreflect.FileDescriptor

var file_internal_testprotos_benchmarks_micro_micro_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x0f, 0x53, 0x69, 0x78, 0x74, 0x65, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x31, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x32, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x32, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x33, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x33, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x34, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x35, 0x18, 0x05, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x35, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x36, 0x18, 0x06, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x36, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x37, 0x18, 0x07, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x37, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x38, 0x18, 0x08, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x38, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x39, 0x18, 0x09, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x02, 0x66, 0x39, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x30, 0x18, 0x0a,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x66, 0x31, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x31,
	0x18, 0x0b, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x66, 0x31, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x31, 0x32, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x66, 0x31, 0x32, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x31, 0x33, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x66, 0x31, 0x33, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x31, 0x34, 0x18, 0x0e, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03, 0x66, 0x31,
	0x34, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x35, 0x18, 0x0f, 0x20, 0x02, 0x28, 0x05, 0x52, 0x03,
	0x66, 0x31, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x31, 0x36, 0x18, 0x10, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x03, 0x66, 0x31, 0x36, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
}

var (
	file_internal_testprotos_benchmarks_micro_micro_proto_rawDescOnce sync.Once
	file_internal_testprotos_benchmarks_micro_micro_proto_rawDescData = file_internal_testprotos_benchmarks_micro_micro_proto_rawDesc
)

func file_internal_testprotos_benchmarks_micro_micro_proto_rawDescGZIP() []byte {
	file_internal_testprotos_benchmarks_micro_micro_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_benchmarks_micro_micro_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_benchmarks_micro_micro_proto_rawDescData)
	})
	return file_internal_testprotos_benchmarks_micro_micro_proto_rawDescData
}

var file_internal_testprotos_benchmarks_micro_micro_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_benchmarks_micro_micro_proto_goTypes = []interface{}{
	(*SixteenRequired)(nil), // 0: goproto.proto.benchmarks.microt.SixteenRequired
}
var file_internal_testprotos_benchmarks_micro_micro_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_benchmarks_micro_micro_proto_init() }
func file_internal_testprotos_benchmarks_micro_micro_proto_init() {
	if File_internal_testprotos_benchmarks_micro_micro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_benchmarks_micro_micro_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SixteenRequired); i {
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
			RawDescriptor: file_internal_testprotos_benchmarks_micro_micro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_benchmarks_micro_micro_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_benchmarks_micro_micro_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_benchmarks_micro_micro_proto_msgTypes,
	}.Build()
	File_internal_testprotos_benchmarks_micro_micro_proto = out.File
	file_internal_testprotos_benchmarks_micro_micro_proto_rawDesc = nil
	file_internal_testprotos_benchmarks_micro_micro_proto_goTypes = nil
	file_internal_testprotos_benchmarks_micro_micro_proto_depIdxs = nil
}
