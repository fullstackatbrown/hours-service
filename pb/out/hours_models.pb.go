// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: hours_models.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket_TicketState int32

const (
	// The ticket was created, but has not yet been claimed.
	Ticket_UNCLAIMED Ticket_TicketState = 0
	// The ticket is currently claimed by a TA.
	Ticket_CLAIMED Ticket_TicketState = 1
	// The ticket has been marked complete.
	Ticket_COMPLETE Ticket_TicketState = 2
	// The ticket has been marked missing.
	Ticket_MISSING Ticket_TicketState = 3
	// A ticket marked missing has been marked returned.
	Ticket_RETURNED Ticket_TicketState = 4
)

// Enum value maps for Ticket_TicketState.
var (
	Ticket_TicketState_name = map[int32]string{
		0: "UNCLAIMED",
		1: "CLAIMED",
		2: "COMPLETE",
		3: "MISSING",
		4: "RETURNED",
	}
	Ticket_TicketState_value = map[string]int32{
		"UNCLAIMED": 0,
		"CLAIMED":   1,
		"COMPLETE":  2,
		"MISSING":   3,
		"RETURNED":  4,
	}
)

func (x Ticket_TicketState) Enum() *Ticket_TicketState {
	p := new(Ticket_TicketState)
	*p = x
	return p
}

func (x Ticket_TicketState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ticket_TicketState) Descriptor() protoreflect.EnumDescriptor {
	return file_hours_models_proto_enumTypes[0].Descriptor()
}

func (Ticket_TicketState) Type() protoreflect.EnumType {
	return &file_hours_models_proto_enumTypes[0]
}

func (x Ticket_TicketState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ticket_TicketState.Descriptor instead.
func (Ticket_TicketState) EnumDescriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{2, 0}
}

// Represents a course.
type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier.
	CourseId         string            `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CreationMetadata *CreationMetadata `protobuf:"bytes,2,opt,name=creation_metadata,json=creationMetadata,proto3" json:"creation_metadata,omitempty"`
	// The academic term this course belongs to.
	Term string `protobuf:"bytes,3,opt,name=term,proto3" json:"term,omitempty"`
	// Required. Contains user-configurable course settings.
	CourseOptions *Course_CourseOptions `protobuf:"bytes,4,opt,name=course_options,json=courseOptions,proto3" json:"course_options,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hours_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_hours_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *Course) GetCreationMetadata() *CreationMetadata {
	if x != nil {
		return x.CreationMetadata
	}
	return nil
}

func (x *Course) GetTerm() string {
	if x != nil {
		return x.Term
	}
	return ""
}

func (x *Course) GetCourseOptions() *Course_CourseOptions {
	if x != nil {
		return x.CourseOptions
	}
	return nil
}

// Represents a queue.
type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier.
	QueueId          string            `protobuf:"bytes,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	CreationMetadata *CreationMetadata `protobuf:"bytes,2,opt,name=creation_metadata,json=creationMetadata,proto3" json:"creation_metadata,omitempty"`
	// The course the queue belongs to.
	Course string `protobuf:"bytes,3,opt,name=course,proto3" json:"course,omitempty"`
	// Required. Contains user-configurable queue settings.
	QueueOptions *Queue_QueueOptions `protobuf:"bytes,4,opt,name=queue_options,json=queueOptions,proto3" json:"queue_options,omitempty"`
	// Tickets that are waiting to be claimed.
	PendingTickets []*Ticket `protobuf:"bytes,5,rep,name=pending_tickets,json=pendingTickets,proto3" json:"pending_tickets,omitempty"`
	// Tickets that have been completed.
	CompletedTickets []*Ticket `protobuf:"bytes,6,rep,name=completed_tickets,json=completedTickets,proto3" json:"completed_tickets,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hours_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_hours_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{1}
}

func (x *Queue) GetQueueId() string {
	if x != nil {
		return x.QueueId
	}
	return ""
}

func (x *Queue) GetCreationMetadata() *CreationMetadata {
	if x != nil {
		return x.CreationMetadata
	}
	return nil
}

func (x *Queue) GetCourse() string {
	if x != nil {
		return x.Course
	}
	return ""
}

func (x *Queue) GetQueueOptions() *Queue_QueueOptions {
	if x != nil {
		return x.QueueOptions
	}
	return nil
}

func (x *Queue) GetPendingTickets() []*Ticket {
	if x != nil {
		return x.PendingTickets
	}
	return nil
}

func (x *Queue) GetCompletedTickets() []*Ticket {
	if x != nil {
		return x.CompletedTickets
	}
	return nil
}

// Represents a ticket in a queue.
type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier.
	TicketId         string            `protobuf:"bytes,1,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	CreationMetadata *CreationMetadata `protobuf:"bytes,2,opt,name=creation_metadata,json=creationMetadata,proto3" json:"creation_metadata,omitempty"`
	// The text content of the ticket (i.e. the student's question).
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// If set, hides the student's name from other student's in the queue.
	IsAnonymous bool `protobuf:"varint,4,opt,name=is_anonymous,json=isAnonymous,proto3" json:"is_anonymous,omitempty"`
	// The current state of the ticket.
	TicketState Ticket_TicketState `protobuf:"varint,5,opt,name=ticket_state,json=ticketState,proto3,enum=hours.Ticket_TicketState" json:"ticket_state,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hours_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_hours_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{2}
}

func (x *Ticket) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *Ticket) GetCreationMetadata() *CreationMetadata {
	if x != nil {
		return x.CreationMetadata
	}
	return nil
}

func (x *Ticket) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Ticket) GetIsAnonymous() bool {
	if x != nil {
		return x.IsAnonymous
	}
	return false
}

func (x *Ticket) GetTicketState() Ticket_TicketState {
	if x != nil {
		return x.TicketState
	}
	return Ticket_UNCLAIMED
}

// User-configurable course parameters.
type Course_CourseOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human-readable title for the course.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *Course_CourseOptions) Reset() {
	*x = Course_CourseOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hours_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course_CourseOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course_CourseOptions) ProtoMessage() {}

func (x *Course_CourseOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hours_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course_CourseOptions.ProtoReflect.Descriptor instead.
func (*Course_CourseOptions) Descriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Course_CourseOptions) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// User-configurable queue parameters.
type Queue_QueueOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human-readable title for the queue.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// An optional description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The location for the queue.
	// Examples: Sun Lab, CIT 205, Zoom, etc.
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// The end time of the queue.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *Queue_QueueOptions) Reset() {
	*x = Queue_QueueOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hours_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue_QueueOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue_QueueOptions) ProtoMessage() {}

func (x *Queue_QueueOptions) ProtoReflect() protoreflect.Message {
	mi := &file_hours_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue_QueueOptions.ProtoReflect.Descriptor instead.
func (*Queue_QueueOptions) Descriptor() ([]byte, []int) {
	return file_hours_models_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Queue_QueueOptions) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Queue_QueueOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Queue_QueueOptions) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Queue_QueueOptions) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_hours_models_proto protoreflect.FileDescriptor

var file_hours_models_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x1a, 0x0e, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x42, 0x0a, 0x0e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x25, 0x0a, 0x0d, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x22, 0xd5, 0x03, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x0c, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x36, 0x0a, 0x0f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0e, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x1a, 0x99, 0x01, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a,
	0x09, 0x82, 0xb5, 0x18, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0xc0, 0x02, 0x0a, 0x06, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x3e, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x69, 0x73, 0x5f, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x12,
	0x3c, 0x0a, 0x0c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x2e, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0a,
	0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x55, 0x4e, 0x43, 0x4c, 0x41, 0x49, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4c, 0x41, 0x49, 0x4d, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e,
	0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x45, 0x44, 0x10,
	0x04, 0x3a, 0x0a, 0x82, 0xb5, 0x18, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6c, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x61, 0x74, 0x62, 0x72, 0x6f, 0x77, 0x6e, 0x2f, 0x68, 0x6f, 0x75,
	0x72, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hours_models_proto_rawDescOnce sync.Once
	file_hours_models_proto_rawDescData = file_hours_models_proto_rawDesc
)

func file_hours_models_proto_rawDescGZIP() []byte {
	file_hours_models_proto_rawDescOnce.Do(func() {
		file_hours_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_hours_models_proto_rawDescData)
	})
	return file_hours_models_proto_rawDescData
}

var file_hours_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_hours_models_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_hours_models_proto_goTypes = []interface{}{
	(Ticket_TicketState)(0),       // 0: hours.Ticket.TicketState
	(*Course)(nil),                // 1: hours.Course
	(*Queue)(nil),                 // 2: hours.Queue
	(*Ticket)(nil),                // 3: hours.Ticket
	(*Course_CourseOptions)(nil),  // 4: hours.Course.CourseOptions
	(*Queue_QueueOptions)(nil),    // 5: hours.Queue.QueueOptions
	(*CreationMetadata)(nil),      // 6: CreationMetadata
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_hours_models_proto_depIdxs = []int32{
	6, // 0: hours.Course.creation_metadata:type_name -> CreationMetadata
	4, // 1: hours.Course.course_options:type_name -> hours.Course.CourseOptions
	6, // 2: hours.Queue.creation_metadata:type_name -> CreationMetadata
	5, // 3: hours.Queue.queue_options:type_name -> hours.Queue.QueueOptions
	3, // 4: hours.Queue.pending_tickets:type_name -> hours.Ticket
	3, // 5: hours.Queue.completed_tickets:type_name -> hours.Ticket
	6, // 6: hours.Ticket.creation_metadata:type_name -> CreationMetadata
	0, // 7: hours.Ticket.ticket_state:type_name -> hours.Ticket.TicketState
	7, // 8: hours.Queue.QueueOptions.end_time:type_name -> google.protobuf.Timestamp
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_hours_models_proto_init() }
func file_hours_models_proto_init() {
	if File_hours_models_proto != nil {
		return
	}
	file_metadata_proto_init()
	file_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_hours_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_hours_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
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
		file_hours_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_hours_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course_CourseOptions); i {
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
		file_hours_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue_QueueOptions); i {
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
			RawDescriptor: file_hours_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hours_models_proto_goTypes,
		DependencyIndexes: file_hours_models_proto_depIdxs,
		EnumInfos:         file_hours_models_proto_enumTypes,
		MessageInfos:      file_hours_models_proto_msgTypes,
	}.Build()
	File_hours_models_proto = out.File
	file_hours_models_proto_rawDesc = nil
	file_hours_models_proto_goTypes = nil
	file_hours_models_proto_depIdxs = nil
}