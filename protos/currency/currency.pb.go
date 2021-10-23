// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: currency.proto

package currency

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

type CurrencyValue int32

const (
	CurrencyValue_a CurrencyValue = 0
	CurrencyValue_b CurrencyValue = 1
	CurrencyValue_c CurrencyValue = 2
	CurrencyValue_d CurrencyValue = 3
	CurrencyValue_e CurrencyValue = 4
	CurrencyValue_f CurrencyValue = 5
	CurrencyValue_g CurrencyValue = 6
	CurrencyValue_h CurrencyValue = 7
	CurrencyValue_i CurrencyValue = 8
	CurrencyValue_j CurrencyValue = 9
)

// Enum value maps for CurrencyValue.
var (
	CurrencyValue_name = map[int32]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
		5: "f",
		6: "g",
		7: "h",
		8: "i",
		9: "j",
	}
	CurrencyValue_value = map[string]int32{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
	}
)

func (x CurrencyValue) Enum() *CurrencyValue {
	p := new(CurrencyValue)
	*p = x
	return p
}

func (x CurrencyValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CurrencyValue) Descriptor() protoreflect.EnumDescriptor {
	return file_currency_proto_enumTypes[0].Descriptor()
}

func (CurrencyValue) Type() protoreflect.EnumType {
	return &file_currency_proto_enumTypes[0]
}

func (x CurrencyValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CurrencyValue.Descriptor instead.
func (CurrencyValue) EnumDescriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base        CurrencyValue `protobuf:"varint,1,opt,name=Base,proto3,enum=currency.CurrencyValue" json:"Base,omitempty"`
	Destination CurrencyValue `protobuf:"varint,2,opt,name=Destination,proto3,enum=currency.CurrencyValue" json:"Destination,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetBase() CurrencyValue {
	if x != nil {
		return x.Base
	}
	return CurrencyValue_a
}

func (x *RateRequest) GetDestination() CurrencyValue {
	if x != nil {
		return x.Destination
	}
	return CurrencyValue_a
}

type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate float32 `protobuf:"fixed32,1,opt,name=Rate,proto3" json:"Rate,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_currency_proto protoreflect.FileDescriptor

var file_currency_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x75, 0x0a, 0x0b, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x52, 0x61, 0x74, 0x65, 0x2a, 0x55, 0x0a, 0x0d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0x0a, 0x01, 0x61, 0x10, 0x00, 0x12, 0x05, 0x0a,
	0x01, 0x62, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x63, 0x10, 0x02, 0x12, 0x05, 0x0a, 0x01, 0x64,
	0x10, 0x03, 0x12, 0x05, 0x0a, 0x01, 0x65, 0x10, 0x04, 0x12, 0x05, 0x0a, 0x01, 0x66, 0x10, 0x05,
	0x12, 0x05, 0x0a, 0x01, 0x67, 0x10, 0x06, 0x12, 0x05, 0x0a, 0x01, 0x68, 0x10, 0x07, 0x12, 0x05,
	0x0a, 0x01, 0x69, 0x10, 0x08, 0x12, 0x05, 0x0a, 0x01, 0x6a, 0x10, 0x09, 0x32, 0x44, 0x0a, 0x08,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x61, 0x75, 0x72, 0x61, 0x62, 0x68, 0x73, 0x69, 0x73, 0x6f, 0x64, 0x69, 0x61, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_currency_proto_rawDescOnce sync.Once
	file_currency_proto_rawDescData = file_currency_proto_rawDesc
)

func file_currency_proto_rawDescGZIP() []byte {
	file_currency_proto_rawDescOnce.Do(func() {
		file_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_currency_proto_rawDescData)
	})
	return file_currency_proto_rawDescData
}

var file_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_currency_proto_goTypes = []interface{}{
	(CurrencyValue)(0),   // 0: currency.CurrencyValue
	(*RateRequest)(nil),  // 1: currency.RateRequest
	(*RateResponse)(nil), // 2: currency.RateResponse
}
var file_currency_proto_depIdxs = []int32{
	0, // 0: currency.RateRequest.Base:type_name -> currency.CurrencyValue
	0, // 1: currency.RateRequest.Destination:type_name -> currency.CurrencyValue
	1, // 2: currency.Currency.GetRate:input_type -> currency.RateRequest
	2, // 3: currency.Currency.GetRate:output_type -> currency.RateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_currency_proto_init() }
func file_currency_proto_init() {
	if File_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
			RawDescriptor: file_currency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_proto_goTypes,
		DependencyIndexes: file_currency_proto_depIdxs,
		EnumInfos:         file_currency_proto_enumTypes,
		MessageInfos:      file_currency_proto_msgTypes,
	}.Build()
	File_currency_proto = out.File
	file_currency_proto_rawDesc = nil
	file_currency_proto_goTypes = nil
	file_currency_proto_depIdxs = nil
}
