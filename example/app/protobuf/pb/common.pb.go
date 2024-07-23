// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: common.proto

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

type OrderStatus int32

const (
	OrderStatus_Order_Status_None          OrderStatus = 0
	OrderStatus_Order_Status_Unpaid        OrderStatus = 1 // 未付款,
	OrderStatus_Order_Status_Paid          OrderStatus = 2 // 已付款
	OrderStatus_Order_Status_Shipped       OrderStatus = 3 // 已发货
	OrderStatus_Order_Status_Received      OrderStatus = 4 // 已签收
	OrderStatus_Order_Status_ReturnRequest OrderStatus = 5 // 退货申请
	OrderStatus_Order_Status_ReturnDuring  OrderStatus = 6 // 退货中
	OrderStatus_Order_Status_ReturnGoods   OrderStatus = 7 // 已退货
	OrderStatus_Order_Status_Cancel        OrderStatus = 8 // 取消交易
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "Order_Status_None",
		1: "Order_Status_Unpaid",
		2: "Order_Status_Paid",
		3: "Order_Status_Shipped",
		4: "Order_Status_Received",
		5: "Order_Status_ReturnRequest",
		6: "Order_Status_ReturnDuring",
		7: "Order_Status_ReturnGoods",
		8: "Order_Status_Cancel",
	}
	OrderStatus_value = map[string]int32{
		"Order_Status_None":          0,
		"Order_Status_Unpaid":        1,
		"Order_Status_Paid":          2,
		"Order_Status_Shipped":       3,
		"Order_Status_Received":      4,
		"Order_Status_ReturnRequest": 5,
		"Order_Status_ReturnDuring":  6,
		"Order_Status_ReturnGoods":   7,
		"Order_Status_Cancel":        8,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x2a, 0xff, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x55, 0x6e, 0x70, 0x61, 0x69, 0x64,
	0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x50, 0x61, 0x69, 0x64, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65,
	0x64, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x10, 0x04, 0x12, 0x1e,
	0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x05, 0x12, 0x1d,
	0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x44, 0x75, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x06, 0x12, 0x1c, 0x0a,
	0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x10, 0x07, 0x12, 0x17, 0x0a, 0x13, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x10, 0x08, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_goTypes = []any{
	(OrderStatus)(0), // 0: pb.OrderStatus
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}