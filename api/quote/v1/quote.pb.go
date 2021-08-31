// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/quote/v1/quote.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTicksSymbolReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *GetTicksSymbolReq) Reset() {
	*x = GetTicksSymbolReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_quote_v1_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicksSymbolReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicksSymbolReq) ProtoMessage() {}

func (x *GetTicksSymbolReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_quote_v1_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicksSymbolReq.ProtoReflect.Descriptor instead.
func (*GetTicksSymbolReq) Descriptor() ([]byte, []int) {
	return file_api_quote_v1_quote_proto_rawDescGZIP(), []int{0}
}

func (x *GetTicksSymbolReq) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *GetTicksSymbolReq) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type GetTicksReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	All      bool   `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	Exchange string `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *GetTicksReq) Reset() {
	*x = GetTicksReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_quote_v1_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTicksReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicksReq) ProtoMessage() {}

func (x *GetTicksReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_quote_v1_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicksReq.ProtoReflect.Descriptor instead.
func (*GetTicksReq) Descriptor() ([]byte, []int) {
	return file_api_quote_v1_quote_proto_rawDescGZIP(), []int{1}
}

func (x *GetTicksReq) GetAll() bool {
	if x != nil {
		return x.All
	}
	return false
}

func (x *GetTicksReq) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

type TickResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticks []byte `protobuf:"bytes,1,opt,name=ticks,proto3" json:"ticks,omitempty"`
}

func (x *TickResp) Reset() {
	*x = TickResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_quote_v1_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickResp) ProtoMessage() {}

func (x *TickResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_quote_v1_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickResp.ProtoReflect.Descriptor instead.
func (*TickResp) Descriptor() ([]byte, []int) {
	return file_api_quote_v1_quote_proto_rawDescGZIP(), []int{2}
}

func (x *TickResp) GetTicks() []byte {
	if x != nil {
		return x.Ticks
	}
	return nil
}

type RateUsdRmb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstrumentId string `protobuf:"bytes,1,opt,name=instrument_id,json=instrumentId,proto3" json:"instrument_id,omitempty"`
	Rate         string `protobuf:"bytes,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Timestamp    string `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RateUsdRmb) Reset() {
	*x = RateUsdRmb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_quote_v1_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateUsdRmb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateUsdRmb) ProtoMessage() {}

func (x *RateUsdRmb) ProtoReflect() protoreflect.Message {
	mi := &file_api_quote_v1_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateUsdRmb.ProtoReflect.Descriptor instead.
func (*RateUsdRmb) Descriptor() ([]byte, []int) {
	return file_api_quote_v1_quote_proto_rawDescGZIP(), []int{3}
}

func (x *RateUsdRmb) GetInstrumentId() string {
	if x != nil {
		return x.InstrumentId
	}
	return ""
}

func (x *RateUsdRmb) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *RateUsdRmb) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

var File_api_quote_v1_quote_proto protoreflect.FileDescriptor

var file_api_quote_v1_quote_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x73, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x3b,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x54,
	0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x63, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x69, 0x63, 0x6b, 0x73, 0x22, 0x63, 0x0a,
	0x0a, 0x52, 0x61, 0x74, 0x65, 0x55, 0x73, 0x64, 0x52, 0x6d, 0x62, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x32, 0xb2, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x57, 0x69, 0x74, 0x68, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x1a, 0x47, 0x65, 0x74,
	0x54, 0x69, 0x63, 0x6b, 0x73, 0x57, 0x69, 0x74, 0x68, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x69, 0x63, 0x6b,
	0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x64, 0x52, 0x6d, 0x62, 0x22, 0x00, 0x42, 0x2c, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1a, 0x66, 0x6f, 0x72, 0x74, 0x75,
	0x6e, 0x65, 0x2d, 0x62, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_quote_v1_quote_proto_rawDescOnce sync.Once
	file_api_quote_v1_quote_proto_rawDescData = file_api_quote_v1_quote_proto_rawDesc
)

func file_api_quote_v1_quote_proto_rawDescGZIP() []byte {
	file_api_quote_v1_quote_proto_rawDescOnce.Do(func() {
		file_api_quote_v1_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_quote_v1_quote_proto_rawDescData)
	})
	return file_api_quote_v1_quote_proto_rawDescData
}

var file_api_quote_v1_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_quote_v1_quote_proto_goTypes = []interface{}{
	(*GetTicksSymbolReq)(nil), // 0: api.quote.v1.GetTicksSymbolReq
	(*GetTicksReq)(nil),       // 1: api.quote.v1.GetTicksReq
	(*TickResp)(nil),          // 2: api.quote.v1.TickResp
	(*RateUsdRmb)(nil),        // 3: api.quote.v1.RateUsdRmb
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_api_quote_v1_quote_proto_depIdxs = []int32{
	1, // 0: api.quote.v1.Quote.GetTicksWithExchange:input_type -> api.quote.v1.GetTicksReq
	0, // 1: api.quote.v1.Quote.GetTicksWithExchangeSymbol:input_type -> api.quote.v1.GetTicksSymbolReq
	1, // 2: api.quote.v1.Quote.StreamTicks:input_type -> api.quote.v1.GetTicksReq
	4, // 3: api.quote.v1.Quote.GetRate:input_type -> google.protobuf.Empty
	2, // 4: api.quote.v1.Quote.GetTicksWithExchange:output_type -> api.quote.v1.TickResp
	2, // 5: api.quote.v1.Quote.GetTicksWithExchangeSymbol:output_type -> api.quote.v1.TickResp
	2, // 6: api.quote.v1.Quote.StreamTicks:output_type -> api.quote.v1.TickResp
	3, // 7: api.quote.v1.Quote.GetRate:output_type -> api.quote.v1.RateUsdRmb
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_quote_v1_quote_proto_init() }
func file_api_quote_v1_quote_proto_init() {
	if File_api_quote_v1_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_quote_v1_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicksSymbolReq); i {
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
		file_api_quote_v1_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTicksReq); i {
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
		file_api_quote_v1_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickResp); i {
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
		file_api_quote_v1_quote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateUsdRmb); i {
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
			RawDescriptor: file_api_quote_v1_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_quote_v1_quote_proto_goTypes,
		DependencyIndexes: file_api_quote_v1_quote_proto_depIdxs,
		MessageInfos:      file_api_quote_v1_quote_proto_msgTypes,
	}.Build()
	File_api_quote_v1_quote_proto = out.File
	file_api_quote_v1_quote_proto_rawDesc = nil
	file_api_quote_v1_quote_proto_goTypes = nil
	file_api_quote_v1_quote_proto_depIdxs = nil
}
