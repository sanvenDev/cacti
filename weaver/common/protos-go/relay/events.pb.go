// Copyright IBM Corp. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: relay/events.proto

package relay

import (
	common "github.com/hyperledger-cacti/cacti/weaver/common/protos-go/v2/common"
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

var File_relay_events_proto protoreflect.FileDescriptor

var file_relay_events_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xd9, 0x01,
	0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x12, 0x45, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63,
	0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41,
	0x63, 0x6b, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e,
	0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61,
	0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x32, 0x8a, 0x01, 0x0a, 0x0c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x3f, 0x0a, 0x0f, 0x53, 0x65,
	0x6e, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x61, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x09, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x63, 0x6b,
	0x2e, 0x41, 0x63, 0x6b, 0x22, 0x00, 0x42, 0x77, 0x0a, 0x30, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79,
	0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2e,
	0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2d, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2f, 0x63, 0x61, 0x63, 0x74, 0x69, 0x2f, 0x77,
	0x65, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_relay_events_proto_goTypes = []interface{}{
	(*common.EventSubscription)(nil), // 0: common.events.EventSubscription
	(*common.Ack)(nil),               // 1: common.ack.Ack
	(*common.ViewPayload)(nil),       // 2: common.state.ViewPayload
}
var file_relay_events_proto_depIdxs = []int32{
	0, // 0: relay.events.EventSubscribe.SubscribeEvent:input_type -> common.events.EventSubscription
	1, // 1: relay.events.EventSubscribe.SendSubscriptionStatus:input_type -> common.ack.Ack
	1, // 2: relay.events.EventSubscribe.SendDriverSubscriptionStatus:input_type -> common.ack.Ack
	2, // 3: relay.events.EventPublish.SendDriverState:input_type -> common.state.ViewPayload
	2, // 4: relay.events.EventPublish.SendState:input_type -> common.state.ViewPayload
	1, // 5: relay.events.EventSubscribe.SubscribeEvent:output_type -> common.ack.Ack
	1, // 6: relay.events.EventSubscribe.SendSubscriptionStatus:output_type -> common.ack.Ack
	1, // 7: relay.events.EventSubscribe.SendDriverSubscriptionStatus:output_type -> common.ack.Ack
	1, // 8: relay.events.EventPublish.SendDriverState:output_type -> common.ack.Ack
	1, // 9: relay.events.EventPublish.SendState:output_type -> common.ack.Ack
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relay_events_proto_init() }
func file_relay_events_proto_init() {
	if File_relay_events_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relay_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_relay_events_proto_goTypes,
		DependencyIndexes: file_relay_events_proto_depIdxs,
	}.Build()
	File_relay_events_proto = out.File
	file_relay_events_proto_rawDesc = nil
	file_relay_events_proto_goTypes = nil
	file_relay_events_proto_depIdxs = nil
}
