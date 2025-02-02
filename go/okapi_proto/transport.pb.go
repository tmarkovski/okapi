// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: transport.okapi_proto

package okapi

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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload  []byte         `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Key      *JsonWebKey    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	AppendTo *SignedMessage `protobuf:"bytes,3,opt,name=append_to,json=appendTo,proto3" json:"append_to,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *SignRequest) GetKey() *JsonWebKey {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SignRequest) GetAppendTo() *SignedMessage {
	if x != nil {
		return x.AppendTo
	}
	return nil
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *SignedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetMessage() *SignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *SignedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Key     *JsonWebKey    `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetMessage() *SignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *VerifyRequest) GetKey() *JsonWebKey {
	if x != nil {
		return x.Key
	}
	return nil
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

type PackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderKey      *JsonWebKey         `protobuf:"bytes,1,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	ReceiverKey    *JsonWebKey         `protobuf:"bytes,2,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	AssociatedData []byte              `protobuf:"bytes,3,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	Plaintext      []byte              `protobuf:"bytes,4,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	Mode           EncryptionMode      `protobuf:"varint,5,opt,name=mode,proto3,enum=pbmse.EncryptionMode" json:"mode,omitempty"`
	Algorithm      EncryptionAlgorithm `protobuf:"varint,6,opt,name=algorithm,proto3,enum=pbmse.EncryptionAlgorithm" json:"algorithm,omitempty"`
}

func (x *PackRequest) Reset() {
	*x = PackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackRequest) ProtoMessage() {}

func (x *PackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackRequest.ProtoReflect.Descriptor instead.
func (*PackRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{4}
}

func (x *PackRequest) GetSenderKey() *JsonWebKey {
	if x != nil {
		return x.SenderKey
	}
	return nil
}

func (x *PackRequest) GetReceiverKey() *JsonWebKey {
	if x != nil {
		return x.ReceiverKey
	}
	return nil
}

func (x *PackRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

func (x *PackRequest) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *PackRequest) GetMode() EncryptionMode {
	if x != nil {
		return x.Mode
	}
	return EncryptionMode_direct
}

func (x *PackRequest) GetAlgorithm() EncryptionAlgorithm {
	if x != nil {
		return x.Algorithm
	}
	return EncryptionAlgorithm_xchacha20poly1305
}

type PackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *EncryptedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PackResponse) Reset() {
	*x = PackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackResponse) ProtoMessage() {}

func (x *PackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackResponse.ProtoReflect.Descriptor instead.
func (*PackResponse) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{5}
}

func (x *PackResponse) GetMessage() *EncryptedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type UnpackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderKey   *JsonWebKey       `protobuf:"bytes,1,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	ReceiverKey *JsonWebKey       `protobuf:"bytes,2,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	Message     *EncryptedMessage `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UnpackRequest) Reset() {
	*x = UnpackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnpackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnpackRequest) ProtoMessage() {}

func (x *UnpackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnpackRequest.ProtoReflect.Descriptor instead.
func (*UnpackRequest) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{6}
}

func (x *UnpackRequest) GetSenderKey() *JsonWebKey {
	if x != nil {
		return x.SenderKey
	}
	return nil
}

func (x *UnpackRequest) GetReceiverKey() *JsonWebKey {
	if x != nil {
		return x.ReceiverKey
	}
	return nil
}

func (x *UnpackRequest) GetMessage() *EncryptedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type UnpackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plaintext []byte `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
}

func (x *UnpackResponse) Reset() {
	*x = UnpackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnpackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnpackResponse) ProtoMessage() {}

func (x *UnpackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnpackResponse.ProtoReflect.Descriptor instead.
func (*UnpackResponse) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{7}
}

func (x *UnpackResponse) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

type CoreMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Body    []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	To      []string `protobuf:"bytes,4,rep,name=to,proto3" json:"to,omitempty"`
	From    string   `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Created int64    `protobuf:"varint,6,opt,name=created,json=created_time,proto3" json:"created,omitempty"`
	Expires int64    `protobuf:"varint,7,opt,name=expires,json=expires_time,proto3" json:"expires,omitempty"`
}

func (x *CoreMessage) Reset() {
	*x = CoreMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoreMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoreMessage) ProtoMessage() {}

func (x *CoreMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoreMessage.ProtoReflect.Descriptor instead.
func (*CoreMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{8}
}

func (x *CoreMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CoreMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CoreMessage) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *CoreMessage) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *CoreMessage) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CoreMessage) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *CoreMessage) GetExpires() int64 {
	if x != nil {
		return x.Expires
	}
	return 0
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x1a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x70, 0x62, 0x6d, 0x73, 0x65, 0x2f, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x28, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x69, 0x67, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x6d, 0x73,
	0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x69, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x62, 0x6d,
	0x73, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b,
	0x65, 0x79, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x2b, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x22, 0xab, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x35, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79,
	0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57,
	0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4b,
	0x65, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x22, 0x41,
	0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x70, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e,
	0x6b, 0x65, 0x79, 0x73, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52,
	0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x4a, 0x73,
	0x6f, 0x6e, 0x57, 0x65, 0x62, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x62, 0x6d, 0x73, 0x65, 0x2e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x55, 0x6e, 0x70, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x72,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x42, 0x3e, 0x0a, 0x0d, 0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2e, 0x6f, 0x6b,
	0x61, 0x70, 0x69, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x72, 0x69, 0x6e, 0x73, 0x69, 0x63, 0x2d, 0x69, 0x64, 0x2f, 0x6f, 0x6b, 0x61, 0x70, 0x69,
	0xaa, 0x02, 0x0f, 0x4f, 0x6b, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_transport_proto_goTypes = []interface{}{
	(*SignRequest)(nil),      // 0: okapi.transport.SignRequest
	(*SignResponse)(nil),     // 1: okapi.transport.SignResponse
	(*VerifyRequest)(nil),    // 2: okapi.transport.VerifyRequest
	(*VerifyResponse)(nil),   // 3: okapi.transport.VerifyResponse
	(*PackRequest)(nil),      // 4: okapi.transport.PackRequest
	(*PackResponse)(nil),     // 5: okapi.transport.PackResponse
	(*UnpackRequest)(nil),    // 6: okapi.transport.UnpackRequest
	(*UnpackResponse)(nil),   // 7: okapi.transport.UnpackResponse
	(*CoreMessage)(nil),      // 8: okapi.transport.CoreMessage
	(*JsonWebKey)(nil),       // 9: okapi.keys.JsonWebKey
	(*SignedMessage)(nil),    // 10: pbmse.SignedMessage
	(EncryptionMode)(0),      // 11: pbmse.EncryptionMode
	(EncryptionAlgorithm)(0), // 12: pbmse.EncryptionAlgorithm
	(*EncryptedMessage)(nil), // 13: pbmse.EncryptedMessage
}
var file_transport_proto_depIdxs = []int32{
	9,  // 0: okapi.transport.SignRequest.key:type_name -> okapi.keys.JsonWebKey
	10, // 1: okapi.transport.SignRequest.append_to:type_name -> pbmse.SignedMessage
	10, // 2: okapi.transport.SignResponse.message:type_name -> pbmse.SignedMessage
	10, // 3: okapi.transport.VerifyRequest.message:type_name -> pbmse.SignedMessage
	9,  // 4: okapi.transport.VerifyRequest.key:type_name -> okapi.keys.JsonWebKey
	9,  // 5: okapi.transport.PackRequest.sender_key:type_name -> okapi.keys.JsonWebKey
	9,  // 6: okapi.transport.PackRequest.receiver_key:type_name -> okapi.keys.JsonWebKey
	11, // 7: okapi.transport.PackRequest.mode:type_name -> pbmse.EncryptionMode
	12, // 8: okapi.transport.PackRequest.algorithm:type_name -> pbmse.EncryptionAlgorithm
	13, // 9: okapi.transport.PackResponse.message:type_name -> pbmse.EncryptedMessage
	9,  // 10: okapi.transport.UnpackRequest.sender_key:type_name -> okapi.keys.JsonWebKey
	9,  // 11: okapi.transport.UnpackRequest.receiver_key:type_name -> okapi.keys.JsonWebKey
	13, // 12: okapi.transport.UnpackRequest.message:type_name -> pbmse.EncryptedMessage
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	file_keys_proto_init()
	file_pbmse_pbmse_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
		file_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackRequest); i {
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
		file_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackResponse); i {
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
		file_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnpackRequest); i {
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
		file_transport_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnpackResponse); i {
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
		file_transport_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoreMessage); i {
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
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
