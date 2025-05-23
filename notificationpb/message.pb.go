// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: notificationpb/message.proto

package notificationpb

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

type EmailStatus int32

const (
	EmailStatus_STATUS_PENDING EmailStatus = 0
	EmailStatus_STATUS_SENT    EmailStatus = 1
	EmailStatus_STATUS_FAILED  EmailStatus = 2
)

// Enum value maps for EmailStatus.
var (
	EmailStatus_name = map[int32]string{
		0: "STATUS_PENDING",
		1: "STATUS_SENT",
		2: "STATUS_FAILED",
	}
	EmailStatus_value = map[string]int32{
		"STATUS_PENDING": 0,
		"STATUS_SENT":    1,
		"STATUS_FAILED":  2,
	}
)

func (x EmailStatus) Enum() *EmailStatus {
	p := new(EmailStatus)
	*p = x
	return p
}

func (x EmailStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_notificationpb_message_proto_enumTypes[0].Descriptor()
}

func (EmailStatus) Type() protoreflect.EnumType {
	return &file_notificationpb_message_proto_enumTypes[0]
}

func (x EmailStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailStatus.Descriptor instead.
func (EmailStatus) EnumDescriptor() ([]byte, []int) {
	return file_notificationpb_message_proto_rawDescGZIP(), []int{0}
}

type NotificationType int32

const (
	NotificationType_NOTIFICATION_TYPE_EMAIL   NotificationType = 0
	NotificationType_NOTIFICATION_TYPE_SMS     NotificationType = 1
	NotificationType_NOTIFICATION_TYPE_PUSH    NotificationType = 2
	NotificationType_NOTIFICATION_TYPE_UNKNOWN NotificationType = 3
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "NOTIFICATION_TYPE_EMAIL",
		1: "NOTIFICATION_TYPE_SMS",
		2: "NOTIFICATION_TYPE_PUSH",
		3: "NOTIFICATION_TYPE_UNKNOWN",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_EMAIL":   0,
		"NOTIFICATION_TYPE_SMS":     1,
		"NOTIFICATION_TYPE_PUSH":    2,
		"NOTIFICATION_TYPE_UNKNOWN": 3,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notificationpb_message_proto_enumTypes[1].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_notificationpb_message_proto_enumTypes[1]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_notificationpb_message_proto_rawDescGZIP(), []int{1}
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject       string           `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Content       string           `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	SenderName    string           `protobuf:"bytes,4,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
	Sender        string           `protobuf:"bytes,5,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient     string           `protobuf:"bytes,6,opt,name=recipient,proto3" json:"recipient,omitempty"`
	RecipientName string           `protobuf:"bytes,7,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
	Status        EmailStatus      `protobuf:"varint,8,opt,name=status,proto3,enum=notificationpb.EmailStatus" json:"status,omitempty"`
	Type          NotificationType `protobuf:"varint,9,opt,name=type,proto3,enum=notificationpb.NotificationType" json:"type,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notificationpb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notificationpb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notificationpb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notification) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notification) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

func (x *Notification) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *Notification) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *Notification) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

func (x *Notification) GetStatus() EmailStatus {
	if x != nil {
		return x.Status
	}
	return EmailStatus_STATUS_PENDING
}

func (x *Notification) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_NOTIFICATION_TYPE_EMAIL
}

var File_notificationpb_message_proto protoreflect.FileDescriptor

var file_notificationpb_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x22, 0xbb,
	0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x62, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x45, 0x0a, 0x0b,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x02, 0x2a, 0x85, 0x01, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x4f, 0x54, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4d,
	0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4d, 0x53, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19,
	0x4e, 0x4f, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x42, 0xd5, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x70, 0x62, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b,
	0x69, 0x66, 0x69, 0x79, 0x61, 0x2d, 0x46, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x69, 0x61, 0x6c, 0x2d,
	0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0xa2, 0x02, 0x03,
	0x4e, 0x58, 0x58, 0xaa, 0x02, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x70, 0x62, 0xca, 0x02, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x70, 0x62, 0xe2, 0x02, 0x1a, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notificationpb_message_proto_rawDescOnce sync.Once
	file_notificationpb_message_proto_rawDescData = file_notificationpb_message_proto_rawDesc
)

func file_notificationpb_message_proto_rawDescGZIP() []byte {
	file_notificationpb_message_proto_rawDescOnce.Do(func() {
		file_notificationpb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_notificationpb_message_proto_rawDescData)
	})
	return file_notificationpb_message_proto_rawDescData
}

var file_notificationpb_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_notificationpb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notificationpb_message_proto_goTypes = []interface{}{
	(EmailStatus)(0),      // 0: notificationpb.EmailStatus
	(NotificationType)(0), // 1: notificationpb.NotificationType
	(*Notification)(nil),  // 2: notificationpb.Notification
}
var file_notificationpb_message_proto_depIdxs = []int32{
	0, // 0: notificationpb.Notification.status:type_name -> notificationpb.EmailStatus
	1, // 1: notificationpb.Notification.type:type_name -> notificationpb.NotificationType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notificationpb_message_proto_init() }
func file_notificationpb_message_proto_init() {
	if File_notificationpb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notificationpb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_notificationpb_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notificationpb_message_proto_goTypes,
		DependencyIndexes: file_notificationpb_message_proto_depIdxs,
		EnumInfos:         file_notificationpb_message_proto_enumTypes,
		MessageInfos:      file_notificationpb_message_proto_msgTypes,
	}.Build()
	File_notificationpb_message_proto = out.File
	file_notificationpb_message_proto_rawDesc = nil
	file_notificationpb_message_proto_goTypes = nil
	file_notificationpb_message_proto_depIdxs = nil
}
