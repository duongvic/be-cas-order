// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: mail.proto

package mails

import (
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

var File_mail_proto protoreflect.FileDescriptor

var file_mail_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x61,
	0x69, 0x6c, 0x73, 0x1a, 0x10, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x40, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0c, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x4d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mail_proto_goTypes = []interface{}{
	(*Order)(nil),     // 0: mails.Order
	(*MailReply)(nil), // 1: mails.MailReply
}
var file_mail_proto_depIdxs = []int32{
	0, // 0: mails.MailService.send_order_info:input_type -> mails.Order
	1, // 1: mails.MailService.send_order_info:output_type -> mails.MailReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mail_proto_init() }
func file_mail_proto_init() {
	if File_mail_proto != nil {
		return
	}
	file_mail_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mail_proto_goTypes,
		DependencyIndexes: file_mail_proto_depIdxs,
	}.Build()
	File_mail_proto = out.File
	file_mail_proto_rawDesc = nil
	file_mail_proto_goTypes = nil
	file_mail_proto_depIdxs = nil
}
