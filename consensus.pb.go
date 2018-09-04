// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus.proto

package consensus // import "github.com/jaimemartinez88/consensus"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

//
// Invoked by candidates to gather votes (§5.2).
// Arguments:
// term candidate’s term
// candidateId candidate requesting vote
// lastLogIndex index of candidate’s last log entry (§5.4)
// lastLogTerm term of candidate’s last log entry (§5.4)
// Results:
// term currentTerm, for candidate to update itself
// voteGranted true means candidate received vote
//
type RequestVoteReq struct {
	CurrentTerm          int64    `protobuf:"varint,1,opt,name=current_term,json=currentTerm" json:"current_term,omitempty"`
	CandidateId          string   `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	LastLogIndex         int64    `protobuf:"varint,3,opt,name=last_log_index,json=lastLogIndex" json:"last_log_index,omitempty"`
	LastLogTerm          int64    `protobuf:"varint,4,opt,name=last_log_term,json=lastLogTerm" json:"last_log_term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteReq) Reset()         { *m = RequestVoteReq{} }
func (m *RequestVoteReq) String() string { return proto.CompactTextString(m) }
func (*RequestVoteReq) ProtoMessage()    {}
func (*RequestVoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_d235faf4f9fb335e, []int{0}
}
func (m *RequestVoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteReq.Unmarshal(m, b)
}
func (m *RequestVoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteReq.Marshal(b, m, deterministic)
}
func (dst *RequestVoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteReq.Merge(dst, src)
}
func (m *RequestVoteReq) XXX_Size() int {
	return xxx_messageInfo_RequestVoteReq.Size(m)
}
func (m *RequestVoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteReq proto.InternalMessageInfo

func (m *RequestVoteReq) GetCurrentTerm() int64 {
	if m != nil {
		return m.CurrentTerm
	}
	return 0
}

func (m *RequestVoteReq) GetCandidateId() string {
	if m != nil {
		return m.CandidateId
	}
	return ""
}

func (m *RequestVoteReq) GetLastLogIndex() int64 {
	if m != nil {
		return m.LastLogIndex
	}
	return 0
}

func (m *RequestVoteReq) GetLastLogTerm() int64 {
	if m != nil {
		return m.LastLogTerm
	}
	return 0
}

type RequestVoteRes struct {
	CurrentTerm          int64    `protobuf:"varint,1,opt,name=current_term,json=currentTerm" json:"current_term,omitempty"`
	VoteGranted          bool     `protobuf:"varint,2,opt,name=vote_granted,json=voteGranted" json:"vote_granted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteRes) Reset()         { *m = RequestVoteRes{} }
func (m *RequestVoteRes) String() string { return proto.CompactTextString(m) }
func (*RequestVoteRes) ProtoMessage()    {}
func (*RequestVoteRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_d235faf4f9fb335e, []int{1}
}
func (m *RequestVoteRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteRes.Unmarshal(m, b)
}
func (m *RequestVoteRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteRes.Marshal(b, m, deterministic)
}
func (dst *RequestVoteRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteRes.Merge(dst, src)
}
func (m *RequestVoteRes) XXX_Size() int {
	return xxx_messageInfo_RequestVoteRes.Size(m)
}
func (m *RequestVoteRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteRes.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteRes proto.InternalMessageInfo

func (m *RequestVoteRes) GetCurrentTerm() int64 {
	if m != nil {
		return m.CurrentTerm
	}
	return 0
}

func (m *RequestVoteRes) GetVoteGranted() bool {
	if m != nil {
		return m.VoteGranted
	}
	return false
}

type AppendEntriesReq struct {
	CurrentTerm          int64    `protobuf:"varint,1,opt,name=current_term,json=currentTerm" json:"current_term,omitempty"`
	LeaderId             string   `protobuf:"bytes,2,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	PrevLogIndex         int64    `protobuf:"varint,3,opt,name=prev_log_index,json=prevLogIndex" json:"prev_log_index,omitempty"`
	PrevLogTerm          int64    `protobuf:"varint,4,opt,name=prev_log_term,json=prevLogTerm" json:"prev_log_term,omitempty"`
	Entries              []string `protobuf:"bytes,5,rep,name=entries" json:"entries,omitempty"`
	LeaderCommit         int64    `protobuf:"varint,6,opt,name=leader_commit,json=leaderCommit" json:"leader_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesReq) Reset()         { *m = AppendEntriesReq{} }
func (m *AppendEntriesReq) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesReq) ProtoMessage()    {}
func (*AppendEntriesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_d235faf4f9fb335e, []int{2}
}
func (m *AppendEntriesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesReq.Unmarshal(m, b)
}
func (m *AppendEntriesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesReq.Marshal(b, m, deterministic)
}
func (dst *AppendEntriesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesReq.Merge(dst, src)
}
func (m *AppendEntriesReq) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesReq.Size(m)
}
func (m *AppendEntriesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesReq proto.InternalMessageInfo

func (m *AppendEntriesReq) GetCurrentTerm() int64 {
	if m != nil {
		return m.CurrentTerm
	}
	return 0
}

func (m *AppendEntriesReq) GetLeaderId() string {
	if m != nil {
		return m.LeaderId
	}
	return ""
}

func (m *AppendEntriesReq) GetPrevLogIndex() int64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesReq) GetPrevLogTerm() int64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesReq) GetEntries() []string {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AppendEntriesReq) GetLeaderCommit() int64 {
	if m != nil {
		return m.LeaderCommit
	}
	return 0
}

type AppendEntriesRes struct {
	CurrentTerm          int64    `protobuf:"varint,1,opt,name=current_term,json=currentTerm" json:"current_term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesRes) Reset()         { *m = AppendEntriesRes{} }
func (m *AppendEntriesRes) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRes) ProtoMessage()    {}
func (*AppendEntriesRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_consensus_d235faf4f9fb335e, []int{3}
}
func (m *AppendEntriesRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesRes.Unmarshal(m, b)
}
func (m *AppendEntriesRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesRes.Marshal(b, m, deterministic)
}
func (dst *AppendEntriesRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesRes.Merge(dst, src)
}
func (m *AppendEntriesRes) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesRes.Size(m)
}
func (m *AppendEntriesRes) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesRes.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesRes proto.InternalMessageInfo

func (m *AppendEntriesRes) GetCurrentTerm() int64 {
	if m != nil {
		return m.CurrentTerm
	}
	return 0
}

func (m *AppendEntriesRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*RequestVoteReq)(nil), "consensus.RequestVoteReq")
	proto.RegisterType((*RequestVoteRes)(nil), "consensus.RequestVoteRes")
	proto.RegisterType((*AppendEntriesReq)(nil), "consensus.AppendEntriesReq")
	proto.RegisterType((*AppendEntriesRes)(nil), "consensus.AppendEntriesRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RaftService service

type RaftServiceClient interface {
	// rpc RegisterServer(Server) returns (Success);
	RequestVote(ctx context.Context, in *RequestVoteReq, opts ...grpc.CallOption) (*RequestVoteRes, error)
	AppendEntries(ctx context.Context, in *AppendEntriesReq, opts ...grpc.CallOption) (*AppendEntriesRes, error)
}

type raftServiceClient struct {
	cc *grpc.ClientConn
}

func NewRaftServiceClient(cc *grpc.ClientConn) RaftServiceClient {
	return &raftServiceClient{cc}
}

func (c *raftServiceClient) RequestVote(ctx context.Context, in *RequestVoteReq, opts ...grpc.CallOption) (*RequestVoteRes, error) {
	out := new(RequestVoteRes)
	err := grpc.Invoke(ctx, "/consensus.RaftService/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftServiceClient) AppendEntries(ctx context.Context, in *AppendEntriesReq, opts ...grpc.CallOption) (*AppendEntriesRes, error) {
	out := new(AppendEntriesRes)
	err := grpc.Invoke(ctx, "/consensus.RaftService/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RaftService service

type RaftServiceServer interface {
	// rpc RegisterServer(Server) returns (Success);
	RequestVote(context.Context, *RequestVoteReq) (*RequestVoteRes, error)
	AppendEntries(context.Context, *AppendEntriesReq) (*AppendEntriesRes, error)
}

func RegisterRaftServiceServer(s *grpc.Server, srv RaftServiceServer) {
	s.RegisterService(&_RaftService_serviceDesc, srv)
}

func _RaftService_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consensus.RaftService/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).RequestVote(ctx, req.(*RequestVoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftService_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServiceServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consensus.RaftService/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServiceServer).AppendEntries(ctx, req.(*AppendEntriesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RaftService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consensus.RaftService",
	HandlerType: (*RaftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _RaftService_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _RaftService_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consensus.proto",
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_consensus_d235faf4f9fb335e) }

var fileDescriptor_consensus_d235faf4f9fb335e = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4e, 0xea, 0x40,
	0x18, 0xc5, 0xd3, 0xcb, 0xbd, 0x40, 0xa7, 0xc0, 0xbd, 0x99, 0x55, 0x2f, 0x6c, 0xb0, 0x12, 0xc3,
	0x0a, 0x12, 0xdd, 0xb0, 0x55, 0x62, 0x0c, 0x89, 0x89, 0x49, 0x35, 0x2c, 0xdc, 0x34, 0x43, 0xfb,
	0x59, 0xc7, 0xd0, 0x99, 0x32, 0xf3, 0x95, 0x18, 0xdf, 0xc5, 0xc4, 0xf7, 0xf2, 0x65, 0x4c, 0xff,
	0x50, 0x0a, 0x44, 0xc3, 0x72, 0x7e, 0x73, 0xda, 0x73, 0xce, 0x37, 0x33, 0xe4, 0xaf, 0x2f, 0x85,
	0x06, 0xa1, 0x13, 0x3d, 0x8a, 0x95, 0x44, 0x49, 0xcd, 0x12, 0x38, 0x1f, 0x06, 0xe9, 0xb8, 0xb0,
	0x4a, 0x40, 0xe3, 0x5c, 0x22, 0xb8, 0xb0, 0xa2, 0x27, 0xa4, 0xe5, 0x27, 0x4a, 0x81, 0x40, 0x0f,
	0x41, 0x45, 0xb6, 0xd1, 0x37, 0x86, 0x35, 0xd7, 0x2a, 0xd8, 0x03, 0xa8, 0x28, 0x93, 0x30, 0x11,
	0xf0, 0x80, 0x21, 0x78, 0x3c, 0xb0, 0x7f, 0xf5, 0x8d, 0xa1, 0xe9, 0x5a, 0x25, 0x9b, 0x05, 0x74,
	0x40, 0x3a, 0x4b, 0xa6, 0xd1, 0x5b, 0xca, 0xd0, 0xe3, 0x22, 0x80, 0x57, 0xbb, 0x96, 0xfd, 0xa7,
	0x95, 0xd2, 0x5b, 0x19, 0xce, 0x52, 0x46, 0x1d, 0xd2, 0x2e, 0x55, 0x99, 0xd9, 0xef, 0xdc, 0xac,
	0x10, 0xa5, 0x66, 0xce, 0x7c, 0x2f, 0xa1, 0x3e, 0x32, 0xe1, 0x5a, 0x22, 0x78, 0xa1, 0x62, 0x02,
	0x21, 0x4f, 0xd8, 0x74, 0xad, 0x94, 0xdd, 0xe4, 0xc8, 0xf9, 0x34, 0xc8, 0xbf, 0xcb, 0x38, 0x06,
	0x11, 0x5c, 0x0b, 0x54, 0x1c, 0xf4, 0x91, 0xe5, 0x7b, 0xc4, 0x5c, 0x02, 0x0b, 0x40, 0x6d, 0x9b,
	0x37, 0x73, 0x90, 0xd7, 0x8e, 0x15, 0xac, 0x0f, 0x6b, 0xa7, 0xb4, 0x5a, 0xbb, 0x54, 0x55, 0x6b,
	0x17, 0xa2, 0xcc, 0xc6, 0x26, 0x0d, 0xc8, 0x73, 0xd9, 0x7f, 0xfa, 0xb5, 0xa1, 0xe9, 0x6e, 0x96,
	0xf4, 0x94, 0xb4, 0x8b, 0x00, 0xbe, 0x8c, 0x22, 0x8e, 0x76, 0xbd, 0x98, 0x6c, 0x06, 0xa7, 0x19,
	0x73, 0xee, 0x0e, 0xca, 0x1d, 0x35, 0x37, 0x9b, 0x34, 0x74, 0xe2, 0xfb, 0xa0, 0x75, 0x31, 0xb2,
	0xcd, 0xf2, 0xfc, 0xdd, 0x20, 0x96, 0xcb, 0x9e, 0xf0, 0x1e, 0xd4, 0x9a, 0xfb, 0x40, 0xa7, 0xc4,
	0xaa, 0x1c, 0x0b, 0xfd, 0x3f, 0xda, 0xde, 0xb2, 0xdd, 0x0b, 0xd5, 0xfd, 0x76, 0x4b, 0xd3, 0x19,
	0x69, 0xef, 0xa4, 0xa4, 0xbd, 0x8a, 0x76, 0xff, 0x70, 0xba, 0x3f, 0x6c, 0xea, 0xab, 0xb3, 0xc7,
	0x41, 0xc8, 0xf1, 0x39, 0x59, 0x8c, 0x7c, 0x19, 0x8d, 0x5f, 0x18, 0x8f, 0x20, 0x62, 0x0a, 0xb9,
	0x80, 0xb7, 0xc9, 0x64, 0x5c, 0x7e, 0xb8, 0xa8, 0x67, 0x6f, 0xe0, 0xe2, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x74, 0x89, 0x08, 0x87, 0x16, 0x03, 0x00, 0x00,
}