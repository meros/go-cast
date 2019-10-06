// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cast_channel.proto

package cast

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Always pass a version of the protocol for future compatibility
// requirements.
type CastMessage_ProtocolVersion int32

const (
	CastMessage_CASTV2_1_0 CastMessage_ProtocolVersion = 0
)

var CastMessage_ProtocolVersion_name = map[int32]string{
	0: "CASTV2_1_0",
}

var CastMessage_ProtocolVersion_value = map[string]int32{
	"CASTV2_1_0": 0,
}

func (x CastMessage_ProtocolVersion) Enum() *CastMessage_ProtocolVersion {
	p := new(CastMessage_ProtocolVersion)
	*p = x
	return p
}

func (x CastMessage_ProtocolVersion) String() string {
	return proto.EnumName(CastMessage_ProtocolVersion_name, int32(x))
}

func (x *CastMessage_ProtocolVersion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CastMessage_ProtocolVersion_value, data, "CastMessage_ProtocolVersion")
	if err != nil {
		return err
	}
	*x = CastMessage_ProtocolVersion(value)
	return nil
}

func (CastMessage_ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{0, 0}
}

// What type of data do we have in this message.
type CastMessage_PayloadType int32

const (
	CastMessage_STRING CastMessage_PayloadType = 0
	CastMessage_BINARY CastMessage_PayloadType = 1
)

var CastMessage_PayloadType_name = map[int32]string{
	0: "STRING",
	1: "BINARY",
}

var CastMessage_PayloadType_value = map[string]int32{
	"STRING": 0,
	"BINARY": 1,
}

func (x CastMessage_PayloadType) Enum() *CastMessage_PayloadType {
	p := new(CastMessage_PayloadType)
	*p = x
	return p
}

func (x CastMessage_PayloadType) String() string {
	return proto.EnumName(CastMessage_PayloadType_name, int32(x))
}

func (x *CastMessage_PayloadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CastMessage_PayloadType_value, data, "CastMessage_PayloadType")
	if err != nil {
		return err
	}
	*x = CastMessage_PayloadType(value)
	return nil
}

func (CastMessage_PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{0, 1}
}

type AuthError_ErrorType int32

const (
	AuthError_INTERNAL_ERROR AuthError_ErrorType = 0
	AuthError_NO_TLS         AuthError_ErrorType = 1
)

var AuthError_ErrorType_name = map[int32]string{
	0: "INTERNAL_ERROR",
	1: "NO_TLS",
}

var AuthError_ErrorType_value = map[string]int32{
	"INTERNAL_ERROR": 0,
	"NO_TLS":         1,
}

func (x AuthError_ErrorType) Enum() *AuthError_ErrorType {
	p := new(AuthError_ErrorType)
	*p = x
	return p
}

func (x AuthError_ErrorType) String() string {
	return proto.EnumName(AuthError_ErrorType_name, int32(x))
}

func (x *AuthError_ErrorType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AuthError_ErrorType_value, data, "AuthError_ErrorType")
	if err != nil {
		return err
	}
	*x = AuthError_ErrorType(value)
	return nil
}

func (AuthError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{3, 0}
}

type CastMessage struct {
	ProtocolVersion *CastMessage_ProtocolVersion `protobuf:"varint,1,req,name=protocol_version,json=protocolVersion,enum=cast.CastMessage_ProtocolVersion" json:"protocol_version,omitempty"`
	// source and destination ids identify the origin and destination of the
	// message.  They are used to route messages between endpoints that share a
	// device-to-device channel.
	//
	// For messages between applications:
	//   - The sender application id is a unique identifier generated on behalf of
	//     the sender application.
	//   - The receiver id is always the the session id for the application.
	//
	// For messages to or from the sender or receiver platform, the special ids
	// 'sender-0' and 'receiver-0' can be used.
	//
	// For messages intended for all endpoints using a given channel, the
	// wildcard destination_id '*' can be used.
	SourceId      *string `protobuf:"bytes,2,req,name=source_id,json=sourceId" json:"source_id,omitempty"`
	DestinationId *string `protobuf:"bytes,3,req,name=destination_id,json=destinationId" json:"destination_id,omitempty"`
	// This is the core multiplexing key.  All messages are sent on a namespace
	// and endpoints sharing a channel listen on one or more namespaces.  The
	// namespace defines the protocol and semantics of the message.
	Namespace   *string                  `protobuf:"bytes,4,req,name=namespace" json:"namespace,omitempty"`
	PayloadType *CastMessage_PayloadType `protobuf:"varint,5,req,name=payload_type,json=payloadType,enum=cast.CastMessage_PayloadType" json:"payload_type,omitempty"`
	// Depending on payload_type, exactly one of the following optional fields
	// will always be set.
	PayloadUtf8          *string  `protobuf:"bytes,6,opt,name=payload_utf8,json=payloadUtf8" json:"payload_utf8,omitempty"`
	PayloadBinary        []byte   `protobuf:"bytes,7,opt,name=payload_binary,json=payloadBinary" json:"payload_binary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CastMessage) Reset()         { *m = CastMessage{} }
func (m *CastMessage) String() string { return proto.CompactTextString(m) }
func (*CastMessage) ProtoMessage()    {}
func (*CastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{0}
}

func (m *CastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CastMessage.Unmarshal(m, b)
}
func (m *CastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CastMessage.Marshal(b, m, deterministic)
}
func (m *CastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CastMessage.Merge(m, src)
}
func (m *CastMessage) XXX_Size() int {
	return xxx_messageInfo_CastMessage.Size(m)
}
func (m *CastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CastMessage proto.InternalMessageInfo

func (m *CastMessage) GetProtocolVersion() CastMessage_ProtocolVersion {
	if m != nil && m.ProtocolVersion != nil {
		return *m.ProtocolVersion
	}
	return CastMessage_CASTV2_1_0
}

func (m *CastMessage) GetSourceId() string {
	if m != nil && m.SourceId != nil {
		return *m.SourceId
	}
	return ""
}

func (m *CastMessage) GetDestinationId() string {
	if m != nil && m.DestinationId != nil {
		return *m.DestinationId
	}
	return ""
}

func (m *CastMessage) GetNamespace() string {
	if m != nil && m.Namespace != nil {
		return *m.Namespace
	}
	return ""
}

func (m *CastMessage) GetPayloadType() CastMessage_PayloadType {
	if m != nil && m.PayloadType != nil {
		return *m.PayloadType
	}
	return CastMessage_STRING
}

func (m *CastMessage) GetPayloadUtf8() string {
	if m != nil && m.PayloadUtf8 != nil {
		return *m.PayloadUtf8
	}
	return ""
}

func (m *CastMessage) GetPayloadBinary() []byte {
	if m != nil {
		return m.PayloadBinary
	}
	return nil
}

// Messages for authentication protocol between a sender and a receiver.
type AuthChallenge struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthChallenge) Reset()         { *m = AuthChallenge{} }
func (m *AuthChallenge) String() string { return proto.CompactTextString(m) }
func (*AuthChallenge) ProtoMessage()    {}
func (*AuthChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{1}
}

func (m *AuthChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthChallenge.Unmarshal(m, b)
}
func (m *AuthChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthChallenge.Marshal(b, m, deterministic)
}
func (m *AuthChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthChallenge.Merge(m, src)
}
func (m *AuthChallenge) XXX_Size() int {
	return xxx_messageInfo_AuthChallenge.Size(m)
}
func (m *AuthChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_AuthChallenge proto.InternalMessageInfo

type AuthResponse struct {
	Signature             []byte   `protobuf:"bytes,1,req,name=signature" json:"signature,omitempty"`
	ClientAuthCertificate []byte   `protobuf:"bytes,2,req,name=client_auth_certificate,json=clientAuthCertificate" json:"client_auth_certificate,omitempty"`
	ClientCa              [][]byte `protobuf:"bytes,3,rep,name=client_ca,json=clientCa" json:"client_ca,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{2}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AuthResponse) GetClientAuthCertificate() []byte {
	if m != nil {
		return m.ClientAuthCertificate
	}
	return nil
}

func (m *AuthResponse) GetClientCa() [][]byte {
	if m != nil {
		return m.ClientCa
	}
	return nil
}

type AuthError struct {
	ErrorType            *AuthError_ErrorType `protobuf:"varint,1,req,name=error_type,json=errorType,enum=cast.AuthError_ErrorType" json:"error_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AuthError) Reset()         { *m = AuthError{} }
func (m *AuthError) String() string { return proto.CompactTextString(m) }
func (*AuthError) ProtoMessage()    {}
func (*AuthError) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{3}
}

func (m *AuthError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthError.Unmarshal(m, b)
}
func (m *AuthError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthError.Marshal(b, m, deterministic)
}
func (m *AuthError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthError.Merge(m, src)
}
func (m *AuthError) XXX_Size() int {
	return xxx_messageInfo_AuthError.Size(m)
}
func (m *AuthError) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthError.DiscardUnknown(m)
}

var xxx_messageInfo_AuthError proto.InternalMessageInfo

func (m *AuthError) GetErrorType() AuthError_ErrorType {
	if m != nil && m.ErrorType != nil {
		return *m.ErrorType
	}
	return AuthError_INTERNAL_ERROR
}

type DeviceAuthMessage struct {
	// Request fields
	Challenge *AuthChallenge `protobuf:"bytes,1,opt,name=challenge" json:"challenge,omitempty"`
	// Response fields
	Response             *AuthResponse `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Error                *AuthError    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DeviceAuthMessage) Reset()         { *m = DeviceAuthMessage{} }
func (m *DeviceAuthMessage) String() string { return proto.CompactTextString(m) }
func (*DeviceAuthMessage) ProtoMessage()    {}
func (*DeviceAuthMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ec15de9f2e74a1, []int{4}
}

func (m *DeviceAuthMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAuthMessage.Unmarshal(m, b)
}
func (m *DeviceAuthMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAuthMessage.Marshal(b, m, deterministic)
}
func (m *DeviceAuthMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAuthMessage.Merge(m, src)
}
func (m *DeviceAuthMessage) XXX_Size() int {
	return xxx_messageInfo_DeviceAuthMessage.Size(m)
}
func (m *DeviceAuthMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAuthMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAuthMessage proto.InternalMessageInfo

func (m *DeviceAuthMessage) GetChallenge() *AuthChallenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *DeviceAuthMessage) GetResponse() *AuthResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DeviceAuthMessage) GetError() *AuthError {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterEnum("cast.CastMessage_ProtocolVersion", CastMessage_ProtocolVersion_name, CastMessage_ProtocolVersion_value)
	proto.RegisterEnum("cast.CastMessage_PayloadType", CastMessage_PayloadType_name, CastMessage_PayloadType_value)
	proto.RegisterEnum("cast.AuthError_ErrorType", AuthError_ErrorType_name, AuthError_ErrorType_value)
	proto.RegisterType((*CastMessage)(nil), "cast.CastMessage")
	proto.RegisterType((*AuthChallenge)(nil), "cast.AuthChallenge")
	proto.RegisterType((*AuthResponse)(nil), "cast.AuthResponse")
	proto.RegisterType((*AuthError)(nil), "cast.AuthError")
	proto.RegisterType((*DeviceAuthMessage)(nil), "cast.DeviceAuthMessage")
}

func init() { proto.RegisterFile("cast_channel.proto", fileDescriptor_a9ec15de9f2e74a1) }

var fileDescriptor_a9ec15de9f2e74a1 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0xda, 0x40,
	0x14, 0xc7, 0x37, 0x66, 0x77, 0xbb, 0x79, 0x66, 0xd5, 0x4e, 0x29, 0x4d, 0x69, 0x0b, 0xd9, 0x80,
	0x10, 0x28, 0x84, 0xae, 0x87, 0xe2, 0xb1, 0x6a, 0xa5, 0x15, 0xac, 0x5b, 0x46, 0xbb, 0xd0, 0x53,
	0x98, 0x8e, 0x4f, 0x0d, 0xa4, 0x93, 0x30, 0x33, 0x2e, 0x78, 0xec, 0x67, 0xe8, 0xb9, 0xdf, 0xb5,
	0xcc, 0x44, 0x8d, 0x5d, 0x7a, 0x91, 0xf1, 0xf7, 0x7e, 0x6f, 0xe6, 0xe5, 0x3f, 0x03, 0x84, 0x33,
	0xa5, 0x53, 0xbe, 0x61, 0x42, 0x60, 0x9e, 0x94, 0xb2, 0xd0, 0x05, 0x39, 0x37, 0x2c, 0xfa, 0xed,
	0x42, 0x73, 0xc4, 0x94, 0xfe, 0x82, 0x4a, 0xb1, 0x35, 0x92, 0x29, 0x74, 0x6c, 0x99, 0x17, 0x79,
	0xfa, 0x80, 0x52, 0x65, 0x85, 0x08, 0x9c, 0xb0, 0x11, 0xb7, 0x7a, 0x37, 0x89, 0x69, 0x48, 0x4e,
	0xe4, 0xe4, 0xeb, 0xde, 0xbc, 0xaf, 0x44, 0xda, 0x2e, 0xff, 0x05, 0xe4, 0x15, 0x78, 0xaa, 0xd8,
	0x4a, 0x8e, 0x69, 0xb6, 0x0c, 0x1a, 0x61, 0x23, 0xf6, 0xe8, 0x55, 0x05, 0x26, 0x4b, 0xd2, 0x85,
	0xd6, 0x12, 0x95, 0xce, 0x04, 0xd3, 0x59, 0x21, 0x8c, 0xe1, 0x5a, 0xe3, 0xfa, 0x84, 0x4e, 0x96,
	0xe4, 0x35, 0x78, 0x82, 0xfd, 0x44, 0x55, 0x32, 0x8e, 0xc1, 0xb9, 0x35, 0x6a, 0x40, 0x3e, 0x80,
	0x5f, 0xb2, 0x5d, 0x5e, 0xb0, 0x65, 0xaa, 0x77, 0x25, 0x06, 0x17, 0x76, 0xd6, 0x37, 0xff, 0x99,
	0xb5, 0xb2, 0x16, 0xbb, 0x12, 0x69, 0xb3, 0xac, 0xff, 0x90, 0x9b, 0x7a, 0x87, 0xad, 0x5e, 0xf5,
	0x83, 0xcb, 0xd0, 0x89, 0xbd, 0xa3, 0xf2, 0x4d, 0xaf, 0xfa, 0x66, 0xd2, 0x83, 0xf2, 0x23, 0x13,
	0x4c, 0xee, 0x82, 0x27, 0xa1, 0x13, 0xfb, 0xf4, 0x7a, 0x4f, 0x87, 0x16, 0x46, 0x37, 0xd0, 0x7e,
	0x94, 0x08, 0x69, 0x01, 0x8c, 0x06, 0xf3, 0xc5, 0x7d, 0x2f, 0xbd, 0x4d, 0xdf, 0x75, 0xce, 0xa2,
	0x2e, 0x34, 0x4f, 0x06, 0x21, 0x00, 0x97, 0xf3, 0x05, 0x9d, 0xcc, 0x3e, 0x75, 0xce, 0xcc, 0x7a,
	0x38, 0x99, 0x0d, 0xe8, 0xf7, 0x8e, 0x13, 0xb5, 0xe1, 0x7a, 0xb0, 0xd5, 0x9b, 0xd1, 0x86, 0xe5,
	0x39, 0x8a, 0x35, 0x46, 0xbf, 0x1c, 0xf0, 0x0d, 0xa1, 0xa8, 0xca, 0x42, 0x28, 0x34, 0xa9, 0xa8,
	0x6c, 0x2d, 0x98, 0xde, 0x4a, 0xb4, 0x17, 0xe4, 0xd3, 0x1a, 0x90, 0xf7, 0xf0, 0x82, 0xe7, 0x19,
	0x0a, 0x9d, 0xb2, 0xad, 0xde, 0xa4, 0x1c, 0xa5, 0xce, 0x56, 0x19, 0x67, 0x1a, 0xed, 0x2d, 0xf8,
	0xf4, 0x79, 0x55, 0xb6, 0x87, 0xd4, 0x45, 0x73, 0x5f, 0xfb, 0x3e, 0xce, 0x02, 0x37, 0x74, 0x63,
	0x9f, 0x5e, 0x55, 0x60, 0xc4, 0x22, 0x09, 0x9e, 0xf1, 0xc7, 0x52, 0x16, 0x92, 0xf4, 0x01, 0xd0,
	0x2c, 0xaa, 0xd4, 0xab, 0x17, 0xf2, 0xb2, 0x4a, 0xfd, 0x28, 0x25, 0xf6, 0xd7, 0x26, 0xee, 0xe1,
	0x61, 0x19, 0xbd, 0x05, 0xef, 0xc8, 0x09, 0x81, 0xd6, 0x64, 0xb6, 0x18, 0xd3, 0xd9, 0x60, 0x9a,
	0x8e, 0x29, 0xbd, 0xa3, 0x55, 0x10, 0xb3, 0xbb, 0x74, 0x31, 0x9d, 0x77, 0x9c, 0xe8, 0x8f, 0x03,
	0x4f, 0x3f, 0xe2, 0x43, 0xc6, 0xd1, 0xec, 0x7a, 0x78, 0xa4, 0xb7, 0xe0, 0xf1, 0x43, 0x34, 0x81,
	0x13, 0x3a, 0x71, 0xb3, 0xf7, 0xac, 0x3e, 0xfb, 0x98, 0x1a, 0xad, 0x2d, 0x92, 0xc0, 0x95, 0xdc,
	0x67, 0x17, 0x34, 0x6c, 0x07, 0xa9, 0x3b, 0x0e, 0xa9, 0xd2, 0xa3, 0x43, 0xba, 0x70, 0x61, 0x47,
	0x0e, 0x5c, 0x2b, 0xb7, 0x1f, 0x7d, 0x1a, 0xad, 0xaa, 0xc3, 0xc6, 0x67, 0xf7, 0x6f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x24, 0x5c, 0x8e, 0x3e, 0x5d, 0x03, 0x00, 0x00,
}