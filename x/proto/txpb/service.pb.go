// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: kwil/txsvc/service.proto

package txpb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_kwil_txsvc_service_proto protoreflect.FileDescriptor

var file_kwil_txsvc_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x78, 0x73, 0x76,
	0x63, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2f, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6b, 0x77, 0x69,
	0x6c, 0x2f, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6b, 0x77, 0x69, 0x6c, 0x2f,
	0x74, 0x78, 0x73, 0x76, 0x63, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2f,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbf, 0x03, 0x0a,
	0x09, 0x54, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x62, 0x72, 0x6f, 0x61,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x69, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x17, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x78,
	0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x7b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x7b,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x7d, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x6d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x73, 0x12, 0x1b, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x7b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12,
	0x7d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x12, 0x1c, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x74, 0x78, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x7b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x7d, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x7d, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x42, 0x13,
	0x5a, 0x11, 0x6b, 0x77, 0x69, 0x6c, 0x2f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x78, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_kwil_txsvc_service_proto_goTypes = []interface{}{
	(*BroadcastRequest)(nil),       // 0: txsvc.BroadcastRequest
	(*GetSchemaRequest)(nil),       // 1: txsvc.GetSchemaRequest
	(*ListDatabasesRequest)(nil),   // 2: txsvc.ListDatabasesRequest
	(*GetExecutablesRequest)(nil),  // 3: txsvc.GetExecutablesRequest
	(*BroadcastResponse)(nil),      // 4: txsvc.BroadcastResponse
	(*GetSchemaResponse)(nil),      // 5: txsvc.GetSchemaResponse
	(*ListDatabasesResponse)(nil),  // 6: txsvc.ListDatabasesResponse
	(*GetExecutablesResponse)(nil), // 7: txsvc.GetExecutablesResponse
}
var file_kwil_txsvc_service_proto_depIdxs = []int32{
	0, // 0: txsvc.TxService.Broadcast:input_type -> txsvc.BroadcastRequest
	1, // 1: txsvc.TxService.GetSchema:input_type -> txsvc.GetSchemaRequest
	2, // 2: txsvc.TxService.ListDatabases:input_type -> txsvc.ListDatabasesRequest
	3, // 3: txsvc.TxService.GetExecutables:input_type -> txsvc.GetExecutablesRequest
	4, // 4: txsvc.TxService.Broadcast:output_type -> txsvc.BroadcastResponse
	5, // 5: txsvc.TxService.GetSchema:output_type -> txsvc.GetSchemaResponse
	6, // 6: txsvc.TxService.ListDatabases:output_type -> txsvc.ListDatabasesResponse
	7, // 7: txsvc.TxService.GetExecutables:output_type -> txsvc.GetExecutablesResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kwil_txsvc_service_proto_init() }
func file_kwil_txsvc_service_proto_init() {
	if File_kwil_txsvc_service_proto != nil {
		return
	}
	file_kwil_txsvc_broadcast_proto_init()
	file_kwil_txsvc_executables_proto_init()
	file_kwil_txsvc_list_db_proto_init()
	file_kwil_txsvc_schema_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kwil_txsvc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kwil_txsvc_service_proto_goTypes,
		DependencyIndexes: file_kwil_txsvc_service_proto_depIdxs,
	}.Build()
	File_kwil_txsvc_service_proto = out.File
	file_kwil_txsvc_service_proto_rawDesc = nil
	file_kwil_txsvc_service_proto_goTypes = nil
	file_kwil_txsvc_service_proto_depIdxs = nil
}
