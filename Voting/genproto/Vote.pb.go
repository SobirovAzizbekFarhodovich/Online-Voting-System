// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: Vote.proto

package genproto

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

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CandidateId string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vote) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type GetAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CandidateId string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
}

func (x *GetAll) Reset() {
	*x = GetAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAll) ProtoMessage() {}

func (x *GetAll) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAll.ProtoReflect.Descriptor instead.
func (*GetAll) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{1}
}

func (x *GetAll) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAll) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

type GetAllVotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes []*GetAll `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (x *GetAllVotes) Reset() {
	*x = GetAllVotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVotes) ProtoMessage() {}

func (x *GetAllVotes) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVotes.ProtoReflect.Descriptor instead.
func (*GetAllVotes) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllVotes) GetVotes() []*GetAll {
	if x != nil {
		return x.Votes
	}
	return nil
}

type Vote_Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Vote_Void) Reset() {
	*x = Vote_Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote_Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote_Void) ProtoMessage() {}

func (x *Vote_Void) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote_Void.ProtoReflect.Descriptor instead.
func (*Vote_Void) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{3}
}

// Create
type CreateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote *Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *CreateVoteRequest) Reset() {
	*x = CreateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVoteRequest) ProtoMessage() {}

func (x *CreateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVoteRequest.ProtoReflect.Descriptor instead.
func (*CreateVoteRequest) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{4}
}

func (x *CreateVoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{5}
}

func (x *CreateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Get
type GetVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVoteRequest) Reset() {
	*x = GetVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVoteRequest) ProtoMessage() {}

func (x *GetVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVoteRequest.ProtoReflect.Descriptor instead.
func (*GetVoteRequest) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{6}
}

func (x *GetVoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Update
type UpdateVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vote *Vote `protobuf:"bytes,1,opt,name=vote,proto3" json:"vote,omitempty"`
}

func (x *UpdateVoteRequest) Reset() {
	*x = UpdateVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateVoteRequest) ProtoMessage() {}

func (x *UpdateVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateVoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateVoteRequest) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateVoteRequest) GetVote() *Vote {
	if x != nil {
		return x.Vote
	}
	return nil
}

// Delete
type DeleteVoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteVoteRequest) Reset() {
	*x = DeleteVoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Vote_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVoteRequest) ProtoMessage() {}

func (x *DeleteVoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Vote_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteVoteRequest) Descriptor() ([]byte, []int) {
	return file_Vote_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteVoteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_Vote_proto protoreflect.FileDescriptor

var file_Vote_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x56, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x76, 0x6f,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0x39, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x3b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65,
	0x73, 0x22, 0x0b, 0x0a, 0x09, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52,
	0x04, 0x76, 0x6f, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x04, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76,
	0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x76, 0x6f, 0x74, 0x65,
	0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x9f, 0x02, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x1a, 0x11, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x5f,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x11, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x5f, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x13, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x76, 0x6f, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x5f, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Vote_proto_rawDescOnce sync.Once
	file_Vote_proto_rawDescData = file_Vote_proto_rawDesc
)

func file_Vote_proto_rawDescGZIP() []byte {
	file_Vote_proto_rawDescOnce.Do(func() {
		file_Vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_Vote_proto_rawDescData)
	})
	return file_Vote_proto_rawDescData
}

var file_Vote_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Vote_proto_goTypes = []interface{}{
	(*Vote)(nil),              // 0: voting.Vote
	(*GetAll)(nil),            // 1: voting.GetAll
	(*GetAllVotes)(nil),       // 2: voting.GetAllVotes
	(*Vote_Void)(nil),         // 3: voting.Vote_Void
	(*CreateVoteRequest)(nil), // 4: voting.CreateVoteRequest
	(*CreateResponse)(nil),    // 5: voting.CreateResponse
	(*GetVoteRequest)(nil),    // 6: voting.GetVoteRequest
	(*UpdateVoteRequest)(nil), // 7: voting.UpdateVoteRequest
	(*DeleteVoteRequest)(nil), // 8: voting.DeleteVoteRequest
}
var file_Vote_proto_depIdxs = []int32{
	1, // 0: voting.GetAllVotes.votes:type_name -> voting.GetAll
	0, // 1: voting.CreateVoteRequest.vote:type_name -> voting.Vote
	0, // 2: voting.UpdateVoteRequest.vote:type_name -> voting.Vote
	0, // 3: voting.VoteService.CreateVote:input_type -> voting.Vote
	3, // 4: voting.VoteService.GetAllVote:input_type -> voting.Vote_Void
	6, // 5: voting.VoteService.GetByIdVote:input_type -> voting.GetVoteRequest
	7, // 6: voting.VoteService.UpdateVote:input_type -> voting.UpdateVoteRequest
	8, // 7: voting.VoteService.DeleteVote:input_type -> voting.DeleteVoteRequest
	3, // 8: voting.VoteService.CreateVote:output_type -> voting.Vote_Void
	2, // 9: voting.VoteService.GetAllVote:output_type -> voting.GetAllVotes
	0, // 10: voting.VoteService.GetByIdVote:output_type -> voting.Vote
	3, // 11: voting.VoteService.UpdateVote:output_type -> voting.Vote_Void
	3, // 12: voting.VoteService.DeleteVote:output_type -> voting.Vote_Void
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Vote_proto_init() }
func file_Vote_proto_init() {
	if File_Vote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_Vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAll); i {
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
		file_Vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVotes); i {
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
		file_Vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote_Void); i {
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
		file_Vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVoteRequest); i {
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
		file_Vote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_Vote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVoteRequest); i {
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
		file_Vote_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateVoteRequest); i {
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
		file_Vote_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVoteRequest); i {
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
			RawDescriptor: file_Vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Vote_proto_goTypes,
		DependencyIndexes: file_Vote_proto_depIdxs,
		MessageInfos:      file_Vote_proto_msgTypes,
	}.Build()
	File_Vote_proto = out.File
	file_Vote_proto_rawDesc = nil
	file_Vote_proto_goTypes = nil
	file_Vote_proto_depIdxs = nil
}
