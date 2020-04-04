// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/main.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Move_Direction int32

const (
	Move_UP    Move_Direction = 0
	Move_DOWN  Move_Direction = 1
	Move_LEFT  Move_Direction = 2
	Move_RIGHT Move_Direction = 3
	Move_STOP  Move_Direction = 4
)

var Move_Direction_name = map[int32]string{
	0: "UP",
	1: "DOWN",
	2: "LEFT",
	3: "RIGHT",
	4: "STOP",
}

var Move_Direction_value = map[string]int32{
	"UP":    0,
	"DOWN":  1,
	"LEFT":  2,
	"RIGHT": 3,
	"STOP":  4,
}

func (x Move_Direction) String() string {
	return proto.EnumName(Move_Direction_name, int32(x))
}

func (Move_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{3, 0}
}

type Coordinate struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{0}
}

func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (m *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(m, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinate) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Connect struct {
	Player               string   `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{1}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type Initialize struct {
	Position             *Coordinate `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Initialize) Reset()         { *m = Initialize{} }
func (m *Initialize) String() string { return proto.CompactTextString(m) }
func (*Initialize) ProtoMessage()    {}
func (*Initialize) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{2}
}

func (m *Initialize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Initialize.Unmarshal(m, b)
}
func (m *Initialize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Initialize.Marshal(b, m, deterministic)
}
func (m *Initialize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Initialize.Merge(m, src)
}
func (m *Initialize) XXX_Size() int {
	return xxx_messageInfo_Initialize.Size(m)
}
func (m *Initialize) XXX_DiscardUnknown() {
	xxx_messageInfo_Initialize.DiscardUnknown(m)
}

var xxx_messageInfo_Initialize proto.InternalMessageInfo

func (m *Initialize) GetPosition() *Coordinate {
	if m != nil {
		return m.Position
	}
	return nil
}

type Move struct {
	Direction            Move_Direction `protobuf:"varint,1,opt,name=direction,proto3,enum=proto.Move_Direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Move) Reset()         { *m = Move{} }
func (m *Move) String() string { return proto.CompactTextString(m) }
func (*Move) ProtoMessage()    {}
func (*Move) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{3}
}

func (m *Move) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Move.Unmarshal(m, b)
}
func (m *Move) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Move.Marshal(b, m, deterministic)
}
func (m *Move) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Move.Merge(m, src)
}
func (m *Move) XXX_Size() int {
	return xxx_messageInfo_Move.Size(m)
}
func (m *Move) XXX_DiscardUnknown() {
	xxx_messageInfo_Move.DiscardUnknown(m)
}

var xxx_messageInfo_Move proto.InternalMessageInfo

func (m *Move) GetDirection() Move_Direction {
	if m != nil {
		return m.Direction
	}
	return Move_UP
}

type AddPlayer struct {
	Position             *Coordinate `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddPlayer) Reset()         { *m = AddPlayer{} }
func (m *AddPlayer) String() string { return proto.CompactTextString(m) }
func (*AddPlayer) ProtoMessage()    {}
func (*AddPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{4}
}

func (m *AddPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPlayer.Unmarshal(m, b)
}
func (m *AddPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPlayer.Marshal(b, m, deterministic)
}
func (m *AddPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPlayer.Merge(m, src)
}
func (m *AddPlayer) XXX_Size() int {
	return xxx_messageInfo_AddPlayer.Size(m)
}
func (m *AddPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_AddPlayer proto.InternalMessageInfo

func (m *AddPlayer) GetPosition() *Coordinate {
	if m != nil {
		return m.Position
	}
	return nil
}

type UpdatePlayer struct {
	Position             *Coordinate `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdatePlayer) Reset()         { *m = UpdatePlayer{} }
func (m *UpdatePlayer) String() string { return proto.CompactTextString(m) }
func (*UpdatePlayer) ProtoMessage()    {}
func (*UpdatePlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{5}
}

func (m *UpdatePlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePlayer.Unmarshal(m, b)
}
func (m *UpdatePlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePlayer.Marshal(b, m, deterministic)
}
func (m *UpdatePlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePlayer.Merge(m, src)
}
func (m *UpdatePlayer) XXX_Size() int {
	return xxx_messageInfo_UpdatePlayer.Size(m)
}
func (m *UpdatePlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePlayer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePlayer proto.InternalMessageInfo

func (m *UpdatePlayer) GetPosition() *Coordinate {
	if m != nil {
		return m.Position
	}
	return nil
}

type RemovePlayer struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePlayer) Reset()         { *m = RemovePlayer{} }
func (m *RemovePlayer) String() string { return proto.CompactTextString(m) }
func (*RemovePlayer) ProtoMessage()    {}
func (*RemovePlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{6}
}

func (m *RemovePlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePlayer.Unmarshal(m, b)
}
func (m *RemovePlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePlayer.Marshal(b, m, deterministic)
}
func (m *RemovePlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePlayer.Merge(m, src)
}
func (m *RemovePlayer) XXX_Size() int {
	return xxx_messageInfo_RemovePlayer.Size(m)
}
func (m *RemovePlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePlayer.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePlayer proto.InternalMessageInfo

type Request struct {
	// Types that are valid to be assigned to Action:
	//	*Request_Connect
	//	*Request_Move
	Action               isRequest_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{7}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type isRequest_Action interface {
	isRequest_Action()
}

type Request_Connect struct {
	Connect *Connect `protobuf:"bytes,1,opt,name=connect,proto3,oneof"`
}

type Request_Move struct {
	Move *Move `protobuf:"bytes,2,opt,name=move,proto3,oneof"`
}

func (*Request_Connect) isRequest_Action() {}

func (*Request_Move) isRequest_Action() {}

func (m *Request) GetAction() isRequest_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Request) GetConnect() *Connect {
	if x, ok := m.GetAction().(*Request_Connect); ok {
		return x.Connect
	}
	return nil
}

func (m *Request) GetMove() *Move {
	if x, ok := m.GetAction().(*Request_Move); ok {
		return x.Move
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Request) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Request_Connect)(nil),
		(*Request_Move)(nil),
	}
}

type Response struct {
	Player string `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*Response_Initialize
	//	*Response_Addplayer
	//	*Response_Updateplayer
	//	*Response_Removeplayer
	Action               isResponse_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_098391ad7281b52b, []int{8}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetPlayer() string {
	if m != nil {
		return m.Player
	}
	return ""
}

type isResponse_Action interface {
	isResponse_Action()
}

type Response_Initialize struct {
	Initialize *Initialize `protobuf:"bytes,3,opt,name=initialize,proto3,oneof"`
}

type Response_Addplayer struct {
	Addplayer *AddPlayer `protobuf:"bytes,4,opt,name=addplayer,proto3,oneof"`
}

type Response_Updateplayer struct {
	Updateplayer *UpdatePlayer `protobuf:"bytes,5,opt,name=updateplayer,proto3,oneof"`
}

type Response_Removeplayer struct {
	Removeplayer *RemovePlayer `protobuf:"bytes,6,opt,name=removeplayer,proto3,oneof"`
}

func (*Response_Initialize) isResponse_Action() {}

func (*Response_Addplayer) isResponse_Action() {}

func (*Response_Updateplayer) isResponse_Action() {}

func (*Response_Removeplayer) isResponse_Action() {}

func (m *Response) GetAction() isResponse_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Response) GetInitialize() *Initialize {
	if x, ok := m.GetAction().(*Response_Initialize); ok {
		return x.Initialize
	}
	return nil
}

func (m *Response) GetAddplayer() *AddPlayer {
	if x, ok := m.GetAction().(*Response_Addplayer); ok {
		return x.Addplayer
	}
	return nil
}

func (m *Response) GetUpdateplayer() *UpdatePlayer {
	if x, ok := m.GetAction().(*Response_Updateplayer); ok {
		return x.Updateplayer
	}
	return nil
}

func (m *Response) GetRemoveplayer() *RemovePlayer {
	if x, ok := m.GetAction().(*Response_Removeplayer); ok {
		return x.Removeplayer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Response) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Response_Initialize)(nil),
		(*Response_Addplayer)(nil),
		(*Response_Updateplayer)(nil),
		(*Response_Removeplayer)(nil),
	}
}

func init() {
	proto.RegisterEnum("proto.Move_Direction", Move_Direction_name, Move_Direction_value)
	proto.RegisterType((*Coordinate)(nil), "proto.Coordinate")
	proto.RegisterType((*Connect)(nil), "proto.Connect")
	proto.RegisterType((*Initialize)(nil), "proto.Initialize")
	proto.RegisterType((*Move)(nil), "proto.Move")
	proto.RegisterType((*AddPlayer)(nil), "proto.AddPlayer")
	proto.RegisterType((*UpdatePlayer)(nil), "proto.UpdatePlayer")
	proto.RegisterType((*RemovePlayer)(nil), "proto.RemovePlayer")
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*Response)(nil), "proto.Response")
}

func init() {
	proto.RegisterFile("proto/main.proto", fileDescriptor_098391ad7281b52b)
}

var fileDescriptor_098391ad7281b52b = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x93, 0x2c, 0x49, 0x93, 0xd7, 0x28, 0x04, 0x23, 0x50, 0xc5, 0x89, 0xf9, 0x54, 0x21,
	0xd1, 0x4d, 0xe9, 0x01, 0xf1, 0xeb, 0x00, 0x1b, 0x2c, 0x93, 0x80, 0x55, 0x5e, 0x27, 0xce, 0x5e,
	0xe3, 0x83, 0xa5, 0x25, 0x0e, 0x89, 0x5b, 0x35, 0xfc, 0x07, 0xfc, 0xd7, 0x28, 0x8e, 0xeb, 0xa4,
	0x07, 0x0e, 0x70, 0x8a, 0xdf, 0x7b, 0xdf, 0xcf, 0x73, 0xfc, 0xfd, 0x42, 0x52, 0xd5, 0x42, 0x8a,
	0xb3, 0x82, 0xf2, 0x72, 0xa1, 0x8e, 0xc8, 0x53, 0x1f, 0x3c, 0x07, 0xb8, 0x10, 0xa2, 0xce, 0x79,
	0x49, 0x25, 0x43, 0x11, 0xd8, 0xfb, 0x99, 0xfd, 0xc2, 0x9e, 0x7b, 0xc4, 0xde, 0x77, 0x55, 0x3b,
	0x73, 0xfa, 0xaa, 0xc5, 0xa7, 0x30, 0xb9, 0x10, 0x65, 0xc9, 0x36, 0x12, 0x3d, 0x03, 0xbf, 0x7a,
	0xa0, 0x2d, 0xab, 0x95, 0x36, 0x24, 0xba, 0xc2, 0xef, 0x00, 0xae, 0x4b, 0x2e, 0x39, 0x7d, 0xe0,
	0xbf, 0x18, 0x7a, 0x05, 0x41, 0x25, 0x1a, 0x2e, 0xb9, 0x28, 0x95, 0x6e, 0x9a, 0x3e, 0xee, 0xef,
	0x5e, 0x0c, 0x37, 0x12, 0x23, 0xc1, 0x2d, 0xb8, 0xdf, 0xc4, 0x8e, 0xa1, 0x25, 0x84, 0x39, 0xaf,
	0xd9, 0xc6, 0x70, 0x71, 0xfa, 0x54, 0x73, 0xdd, 0x7c, 0x71, 0x79, 0x18, 0x92, 0x41, 0x87, 0xdf,
	0x43, 0x68, 0xfa, 0xc8, 0x07, 0xe7, 0x6e, 0x95, 0x58, 0x28, 0x00, 0xf7, 0xf2, 0xe6, 0xc7, 0xf7,
	0xc4, 0xee, 0x4e, 0x5f, 0x3f, 0x7f, 0x59, 0x27, 0x0e, 0x0a, 0xc1, 0x23, 0xd7, 0x57, 0xd9, 0x3a,
	0x39, 0xe9, 0x9a, 0xb7, 0xeb, 0x9b, 0x55, 0xe2, 0xe2, 0xb7, 0x10, 0x7e, 0xcc, 0xf3, 0x95, 0x7a,
	0xc4, 0xbf, 0xfe, 0xf6, 0x07, 0x88, 0xee, 0xaa, 0x9c, 0x4a, 0xf6, 0x7f, 0x78, 0x0c, 0x11, 0x61,
	0x85, 0xd8, 0x69, 0x1c, 0xdf, 0xc3, 0x84, 0xb0, 0x9f, 0x5b, 0xd6, 0x48, 0xf4, 0x12, 0x26, 0x9b,
	0xde, 0x70, 0xbd, 0x28, 0x36, 0x8b, 0x54, 0x37, 0xb3, 0xc8, 0x41, 0x80, 0x4e, 0xc1, 0xed, 0x96,
	0xa8, 0xb4, 0xa6, 0xe9, 0x74, 0xe4, 0x57, 0x66, 0x11, 0x35, 0xfa, 0x14, 0x80, 0x4f, 0x7b, 0xb3,
	0x7e, 0x3b, 0x10, 0x10, 0xd6, 0x54, 0xa2, 0x6c, 0xd8, 0xdf, 0xb2, 0x44, 0x4b, 0x00, 0x6e, 0xb2,
	0x9c, 0x9d, 0x1c, 0xbd, 0x64, 0x08, 0x39, 0xb3, 0xc8, 0x48, 0x86, 0xce, 0x21, 0xa4, 0x79, 0xae,
	0xf7, 0xb9, 0x8a, 0x49, 0x34, 0x63, 0x0c, 0xce, 0x2c, 0x32, 0x88, 0xd0, 0x1b, 0x88, 0xb6, 0xca,
	0x3e, 0x0d, 0x79, 0x0a, 0x7a, 0xa2, 0xa1, 0xb1, 0xb3, 0x99, 0x45, 0x8e, 0xa4, 0x1d, 0x5a, 0x2b,
	0xeb, 0x34, 0xea, 0x1f, 0xa1, 0x63, 0x57, 0x3b, 0x74, 0x2c, 0x1d, 0xbc, 0x48, 0x5f, 0x83, 0x7b,
	0x45, 0x0b, 0x86, 0xce, 0xc0, 0xbf, 0x95, 0x35, 0xa3, 0x05, 0x8a, 0xcd, 0x02, 0x15, 0xc3, 0xf3,
	0x47, 0xa6, 0xee, 0x1d, 0xc3, 0xd6, 0xdc, 0x3e, 0xb7, 0xef, 0x7d, 0xd5, 0x5d, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xf5, 0x23, 0x6d, 0x5a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error)
}

type gameClient struct {
	cc grpc.ClientConnInterface
}

func NewGameClient(cc grpc.ClientConnInterface) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Game_serviceDesc.Streams[0], "/proto.Game/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameStreamClient{stream}
	return x, nil
}

type Game_StreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type gameStreamClient struct {
	grpc.ClientStream
}

func (x *gameStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	Stream(Game_StreamServer) error
}

// UnimplementedGameServer can be embedded to have forward compatible implementations.
type UnimplementedGameServer struct {
}

func (*UnimplementedGameServer) Stream(srv Game_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Stream(&gameStreamServer{stream})
}

type Game_StreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type gameStreamServer struct {
	grpc.ServerStream
}

func (x *gameStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Game",
	HandlerType: (*GameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Game_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/main.proto",
}
