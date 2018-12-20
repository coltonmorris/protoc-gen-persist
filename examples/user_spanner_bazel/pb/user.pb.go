// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/tcncloud/protoc-gen-persist/persist"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends              *Friends             `protobuf:"bytes,3,opt,name=friends,proto3" json:"friends,omitempty"`
	CreatedOn            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetFriends() *Friends {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *User) GetCreatedOn() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

type InsertUserReq struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Friends              *Friends             `protobuf:"bytes,3,opt,name=friends,proto3" json:"friends,omitempty"`
	CreatedOn            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *InsertUserReq) Reset()         { *m = InsertUserReq{} }
func (m *InsertUserReq) String() string { return proto.CompactTextString(m) }
func (*InsertUserReq) ProtoMessage()    {}
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{1}
}
func (m *InsertUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserReq.Unmarshal(m, b)
}
func (m *InsertUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserReq.Marshal(b, m, deterministic)
}
func (dst *InsertUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserReq.Merge(dst, src)
}
func (m *InsertUserReq) XXX_Size() int {
	return xxx_messageInfo_InsertUserReq.Size(m)
}
func (m *InsertUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserReq proto.InternalMessageInfo

func (m *InsertUserReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertUserReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InsertUserReq) GetFriends() *Friends {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *InsertUserReq) GetCreatedOn() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

type Friends struct {
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Friends) Reset()         { *m = Friends{} }
func (m *Friends) String() string { return proto.CompactTextString(m) }
func (*Friends) ProtoMessage()    {}
func (*Friends) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{2}
}
func (m *Friends) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Friends.Unmarshal(m, b)
}
func (m *Friends) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Friends.Marshal(b, m, deterministic)
}
func (dst *Friends) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Friends.Merge(dst, src)
}
func (m *Friends) XXX_Size() int {
	return xxx_messageInfo_Friends.Size(m)
}
func (m *Friends) XXX_DiscardUnknown() {
	xxx_messageInfo_Friends.DiscardUnknown(m)
}

var xxx_messageInfo_Friends proto.InternalMessageInfo

func (m *Friends) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type SliceStringParam struct {
	Slice                []string `protobuf:"bytes,1,rep,name=slice,proto3" json:"slice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SliceStringParam) Reset()         { *m = SliceStringParam{} }
func (m *SliceStringParam) String() string { return proto.CompactTextString(m) }
func (*SliceStringParam) ProtoMessage()    {}
func (*SliceStringParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{3}
}
func (m *SliceStringParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceStringParam.Unmarshal(m, b)
}
func (m *SliceStringParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceStringParam.Marshal(b, m, deterministic)
}
func (dst *SliceStringParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceStringParam.Merge(dst, src)
}
func (m *SliceStringParam) XXX_Size() int {
	return xxx_messageInfo_SliceStringParam.Size(m)
}
func (m *SliceStringParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceStringParam.DiscardUnknown(m)
}

var xxx_messageInfo_SliceStringParam proto.InternalMessageInfo

func (m *SliceStringParam) GetSlice() []string {
	if m != nil {
		return m.Slice
	}
	return nil
}

type FriendsReq struct {
	Names                *SliceStringParam `protobuf:"bytes,2,opt,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FriendsReq) Reset()         { *m = FriendsReq{} }
func (m *FriendsReq) String() string { return proto.CompactTextString(m) }
func (*FriendsReq) ProtoMessage()    {}
func (*FriendsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{4}
}
func (m *FriendsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendsReq.Unmarshal(m, b)
}
func (m *FriendsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendsReq.Marshal(b, m, deterministic)
}
func (dst *FriendsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendsReq.Merge(dst, src)
}
func (m *FriendsReq) XXX_Size() int {
	return xxx_messageInfo_FriendsReq.Size(m)
}
func (m *FriendsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendsReq.DiscardUnknown(m)
}

var xxx_messageInfo_FriendsReq proto.InternalMessageInfo

func (m *FriendsReq) GetNames() *SliceStringParam {
	if m != nil {
		return m.Names
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_111827ed6e76fcb7, []int{5}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*InsertUserReq)(nil), "pb.InsertUserReq")
	proto.RegisterType((*Friends)(nil), "pb.Friends")
	proto.RegisterType((*SliceStringParam)(nil), "pb.SliceStringParam")
	proto.RegisterType((*FriendsReq)(nil), "pb.FriendsReq")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UServClient is the client API for UServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UServClient interface {
	// rpc CreateTable(Empty) returns (Empty) {
	//   option (persist.opts) = {
	//     query: "create_users_table"
	//   };
	// };
	InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error)
	GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error)
	// rpc SelectUserById(User) returns (User) {
	//   option (persist.opts) = {
	//     query: "select_user_by_id"
	//   };
	// };
	// rpc UpdateUserNames(stream User) returns (stream User) {
	//   option (persist.opts) = {
	//     query: "update_user_name"
	//   };
	// };
	// rpc UpdateNameToFoo(User) returns (Empty) {
	//   option (persist.opts) = {
	//     query: "update_name_to_foo"
	//   };
	// };
	UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error)
}

type uServClient struct {
	cc *grpc.ClientConn
}

func NewUServClient(cc *grpc.ClientConn) UServClient {
	return &uServClient{cc}
}

func (c *uServClient) InsertUsers(ctx context.Context, opts ...grpc.CallOption) (UServ_InsertUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[0], "/pb.UServ/InsertUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServInsertUsersClient{stream}
	return x, nil
}

type UServ_InsertUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type uServInsertUsersClient struct {
	grpc.ClientStream
}

func (x *uServInsertUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uServInsertUsersClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) GetAllUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_GetAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[1], "/pb.UServ/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_GetAllUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *uServGetAllUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uServClient) UpdateAllNames(ctx context.Context, in *Empty, opts ...grpc.CallOption) (UServ_UpdateAllNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UServ_serviceDesc.Streams[2], "/pb.UServ/UpdateAllNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &uServUpdateAllNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UServ_UpdateAllNamesClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type uServUpdateAllNamesClient struct {
	grpc.ClientStream
}

func (x *uServUpdateAllNamesClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UServServer is the server API for UServ service.
type UServServer interface {
	// rpc CreateTable(Empty) returns (Empty) {
	//   option (persist.opts) = {
	//     query: "create_users_table"
	//   };
	// };
	InsertUsers(UServ_InsertUsersServer) error
	GetAllUsers(*Empty, UServ_GetAllUsersServer) error
	// rpc SelectUserById(User) returns (User) {
	//   option (persist.opts) = {
	//     query: "select_user_by_id"
	//   };
	// };
	// rpc UpdateUserNames(stream User) returns (stream User) {
	//   option (persist.opts) = {
	//     query: "update_user_name"
	//   };
	// };
	// rpc UpdateNameToFoo(User) returns (Empty) {
	//   option (persist.opts) = {
	//     query: "update_name_to_foo"
	//   };
	// };
	UpdateAllNames(*Empty, UServ_UpdateAllNamesServer) error
}

func RegisterUServServer(s *grpc.Server, srv UServServer) {
	s.RegisterService(&_UServ_serviceDesc, srv)
}

func _UServ_InsertUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UServServer).InsertUsers(&uServInsertUsersServer{stream})
}

type UServ_InsertUsersServer interface {
	SendAndClose(*Empty) error
	Recv() (*User, error)
	grpc.ServerStream
}

type uServInsertUsersServer struct {
	grpc.ServerStream
}

func (x *uServInsertUsersServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uServInsertUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _UServ_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).GetAllUsers(m, &uServGetAllUsersServer{stream})
}

type UServ_GetAllUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *uServGetAllUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UServ_UpdateAllNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UServServer).UpdateAllNames(m, &uServUpdateAllNamesServer{stream})
}

type UServ_UpdateAllNamesServer interface {
	Send(*User) error
	grpc.ServerStream
}

type uServUpdateAllNamesServer struct {
	grpc.ServerStream
}

func (x *uServUpdateAllNamesServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _UServ_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UServ",
	HandlerType: (*UServServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InsertUsers",
			Handler:       _UServ_InsertUsers_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAllUsers",
			Handler:       _UServ_GetAllUsers_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateAllNames",
			Handler:       _UServ_UpdateAllNames_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/user.proto",
}

func init() { proto.RegisterFile("pb/user.proto", fileDescriptor_user_111827ed6e76fcb7) }

var fileDescriptor_user_111827ed6e76fcb7 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xdf, 0x6b, 0xd3, 0x50,
	0x14, 0xc7, 0x4d, 0xd6, 0xfd, 0xe8, 0xa9, 0x1b, 0x35, 0x4e, 0x56, 0xf2, 0xe0, 0x2e, 0x81, 0xb1,
	0x74, 0x96, 0x76, 0x6c, 0x08, 0x2a, 0x6c, 0x34, 0x1b, 0x99, 0x2b, 0xdd, 0xba, 0x92, 0xa4, 0xd3,
	0x82, 0x18, 0xd2, 0xe5, 0xb6, 0x04, 0xd2, 0x24, 0x4b, 0xee, 0x94, 0x3d, 0xea, 0xa3, 0xf8, 0xb0,
	0x3e, 0xf8, 0x0f, 0xf8, 0x2f, 0x88, 0xd0, 0xff, 0x4e, 0xb9, 0xb9, 0x4d, 0x93, 0x69, 0x1d, 0xfa,
	0xe6, 0xdb, 0xfd, 0xf1, 0x3d, 0x9f, 0xf3, 0xbd, 0xe7, 0x1c, 0x2e, 0x2c, 0x07, 0xbd, 0xda, 0x55,
	0x84, 0xc3, 0x6a, 0x10, 0xfa, 0xc4, 0x17, 0xf8, 0xa0, 0x27, 0x3e, 0x0a, 0x70, 0x18, 0x39, 0x11,
	0xa9, 0xf9, 0x01, 0x71, 0x7c, 0x2f, 0x62, 0x57, 0xe2, 0xfa, 0xc0, 0xf7, 0x07, 0x2e, 0xae, 0xc5,
	0xbb, 0xde, 0x55, 0xbf, 0x46, 0x9c, 0x21, 0x8e, 0x88, 0x35, 0x0c, 0x98, 0x40, 0xfa, 0xcc, 0x41,
	0xae, 0x13, 0xe1, 0x50, 0x58, 0x01, 0xde, 0xb1, 0x4b, 0x1c, 0xe2, 0xe4, 0x39, 0x8d, 0x77, 0x6c,
	0x41, 0x80, 0x9c, 0x67, 0x0d, 0x71, 0x89, 0x47, 0x9c, 0x9c, 0xd7, 0xe2, 0xb5, 0xb0, 0x01, 0x8b,
	0xfd, 0xd0, 0xc1, 0x9e, 0x1d, 0x95, 0xe6, 0x10, 0x27, 0x17, 0x76, 0x0a, 0xd5, 0xa0, 0x57, 0x3d,
	0x62, 0x47, 0x5a, 0x72, 0x27, 0x3c, 0x07, 0xb8, 0x08, 0xb1, 0x45, 0xb0, 0x6d, 0xfa, 0x5e, 0x29,
	0x17, 0x2b, 0xc5, 0x2a, 0x73, 0x52, 0x4d, 0x9c, 0x54, 0x8d, 0xc4, 0x89, 0x96, 0x9f, 0xa8, 0xcf,
	0x3c, 0xe9, 0x0b, 0x07, 0xcb, 0x0d, 0x2f, 0xc2, 0x21, 0xa1, 0xa6, 0x34, 0x7c, 0xf9, 0x9f, 0xf8,
	0x5a, 0x87, 0xc5, 0x09, 0x4e, 0x58, 0x85, 0x79, 0x9a, 0x34, 0x2a, 0x71, 0x68, 0x4e, 0xce, 0x6b,
	0x6c, 0x23, 0xc9, 0x50, 0xd4, 0x5d, 0xe7, 0x02, 0xeb, 0x24, 0x74, 0xbc, 0x41, 0xdb, 0x0a, 0xad,
	0x21, 0x55, 0x46, 0xf4, 0x2c, 0x51, 0xc6, 0x1b, 0xe9, 0x19, 0x40, 0xe2, 0x0c, 0x5f, 0x0a, 0x5b,
	0x09, 0x8d, 0x8f, 0xed, 0xac, 0x52, 0xe3, 0xbf, 0x82, 0x92, 0x1c, 0x8b, 0x30, 0xaf, 0x0e, 0x03,
	0x72, 0xbd, 0xf3, 0x6d, 0x09, 0xe6, 0x3b, 0x3a, 0x0e, 0xdf, 0x09, 0x7b, 0x50, 0x48, 0xcb, 0x15,
	0x09, 0x4b, 0x34, 0x9c, 0x2e, 0xc5, 0x3c, 0x5d, 0xc5, 0x6a, 0x69, 0xed, 0xeb, 0x78, 0xc4, 0x0b,
	0x70, 0xdf, 0x89, 0x85, 0x26, 0x1d, 0x9c, 0xa8, 0xcd, 0xbd, 0xe6, 0x64, 0x4e, 0xd8, 0x87, 0xc2,
	0x4b, 0x4c, 0x14, 0xd7, 0x65, 0xe1, 0x69, 0x90, 0x38, 0x25, 0x49, 0x25, 0x1a, 0xfe, 0x10, 0x96,
	0x07, 0x98, 0x98, 0x96, 0xeb, 0xa6, 0xf1, 0xdb, 0x9c, 0x50, 0x86, 0x95, 0x4e, 0x60, 0x5b, 0x04,
	0x2b, 0xae, 0xdb, 0xa2, 0x1e, 0x67, 0x23, 0xee, 0x6d, 0x73, 0xe2, 0x8f, 0x85, 0x8f, 0xe3, 0x11,
	0xff, 0x7d, 0x01, 0x6e, 0x38, 0x68, 0x1e, 0x6a, 0xaa, 0x62, 0xa8, 0xc8, 0x50, 0x0e, 0x4e, 0x54,
	0x14, 0x03, 0x65, 0xc7, 0x46, 0x8e, 0x47, 0xf0, 0x00, 0x87, 0xa8, 0xad, 0x35, 0x4e, 0x15, 0xad,
	0x8b, 0x9a, 0x6a, 0xb7, 0x82, 0xe8, 0xd3, 0xd1, 0xb9, 0xa2, 0x1d, 0x1e, 0x2b, 0x9a, 0xfc, 0x74,
	0xbb, 0x5c, 0x41, 0x93, 0x76, 0xa2, 0x83, 0xae, 0xa1, 0x2a, 0x15, 0x58, 0x4b, 0x7b, 0x9a, 0xd5,
	0x95, 0x45, 0x81, 0x5d, 0x30, 0xc3, 0x26, 0xb1, 0x7a, 0x2e, 0x96, 0x52, 0x7b, 0x5b, 0xe9, 0x12,
	0x3e, 0x70, 0xf0, 0xa6, 0xd1, 0xd2, 0x55, 0xcd, 0x40, 0x8d, 0x96, 0x71, 0xc6, 0x1c, 0x21, 0xd9,
	0xb1, 0x59, 0xfa, 0x69, 0xca, 0x0a, 0x4a, 0x73, 0x95, 0xd1, 0xb9, 0x72, 0xd2, 0x51, 0x75, 0x24,
	0xd7, 0xa9, 0xae, 0xce, 0x84, 0xf5, 0xa9, 0xb2, 0x9e, 0x91, 0x8a, 0xb7, 0x6a, 0x9f, 0x35, 0x31,
	0xad, 0x11, 0x74, 0xa0, 0xa6, 0xab, 0x27, 0xea, 0xa1, 0x81, 0xee, 0x4c, 0x8b, 0x8e, 0xb4, 0xb3,
	0x53, 0x66, 0x50, 0xbc, 0xdd, 0x92, 0xd9, 0xd8, 0x01, 0xec, 0xff, 0x23, 0x16, 0xbd, 0x3a, 0x56,
	0x35, 0x15, 0x39, 0x36, 0xda, 0x43, 0x75, 0xc7, 0x16, 0x1f, 0x44, 0xd8, 0xc5, 0x17, 0xcc, 0xbb,
	0xd9, 0xbb, 0x36, 0x1d, 0x5b, 0x9a, 0xe2, 0x33, 0x89, 0x74, 0x78, 0xc2, 0x46, 0x61, 0x42, 0x89,
	0x30, 0x61, 0xbd, 0xdb, 0x63, 0xb5, 0x41, 0xed, 0xa6, 0x9c, 0x30, 0xcb, 0x48, 0x2c, 0x5e, 0xc5,
	0x62, 0x46, 0xa5, 0xf7, 0x33, 0xa1, 0x6f, 0x61, 0xf7, 0x4f, 0xd0, 0xcd, 0xbe, 0xef, 0x6f, 0x26,
	0xa3, 0x62, 0x36, 0xd5, 0x6e, 0x4a, 0x17, 0x85, 0x09, 0x9c, 0x4a, 0x4d, 0xe2, 0x9b, 0x7d, 0xdf,
	0xcf, 0xe0, 0x33, 0x8d, 0x7f, 0x0f, 0x1b, 0x7f, 0x55, 0x1d, 0xd8, 0xfc, 0xad, 0x3e, 0xb1, 0x8d,
	0x46, 0x0b, 0x75, 0x5a, 0x2d, 0x55, 0x37, 0xe4, 0xf8, 0x8d, 0x51, 0x59, 0x2c, 0xd0, 0x9e, 0x4c,
	0x10, 0xd2, 0x4a, 0xf6, 0x1b, 0xc2, 0x97, 0x99, 0x87, 0x55, 0xa1, 0x68, 0x87, 0x7e, 0x80, 0xe2,
	0xa9, 0x9c, 0xb4, 0x33, 0x47, 0x4f, 0x66, 0x4f, 0xe8, 0xa7, 0xf1, 0x88, 0x7f, 0x01, 0x8f, 0xe1,
	0x8e, 0xdf, 0xaa, 0x58, 0x80, 0x12, 0xcc, 0xfc, 0x3e, 0x8a, 0x85, 0x9b, 0xf1, 0x88, 0xe7, 0x7a,
	0x0b, 0x71, 0xd0, 0xee, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x31, 0x91, 0x77, 0x3e, 0x06,
	0x00, 0x00,
}
