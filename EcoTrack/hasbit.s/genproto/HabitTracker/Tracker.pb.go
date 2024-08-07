// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: Tracker.proto

package HabitTraker

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

type Frequency int32

const (
	Frequency_DAILY   Frequency = 0
	Frequency_WEEKLY  Frequency = 1
	Frequency_MONTHLY Frequency = 2
)

// Enum value maps for Frequency.
var (
	Frequency_name = map[int32]string{
		0: "DAILY",
		1: "WEEKLY",
		2: "MONTHLY",
	}
	Frequency_value = map[string]int32{
		"DAILY":   0,
		"WEEKLY":  1,
		"MONTHLY": 2,
	}
)

func (x Frequency) Enum() *Frequency {
	p := new(Frequency)
	*p = x
	return p
}

func (x Frequency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Frequency) Descriptor() protoreflect.EnumDescriptor {
	return file_Tracker_proto_enumTypes[0].Descriptor()
}

func (Frequency) Type() protoreflect.EnumType {
	return &file_Tracker_proto_enumTypes[0]
}

func (x Frequency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Frequency.Descriptor instead.
func (Frequency) EnumDescriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{0}
}

type Habit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      int32     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Frequency   Frequency `protobuf:"varint,5,opt,name=frequency,proto3,enum=habittracker.Frequency" json:"frequency,omitempty"`
	CreatedAt   string    `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Habit) Reset() {
	*x = Habit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Habit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Habit) ProtoMessage() {}

func (x *Habit) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Habit.ProtoReflect.Descriptor instead.
func (*Habit) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{0}
}

func (x *Habit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Habit) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Habit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Habit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Habit) GetFrequency() Frequency {
	if x != nil {
		return x.Frequency
	}
	return Frequency_DAILY
}

func (x *Habit) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type HabitLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HabitId  string `protobuf:"bytes,2,opt,name=habit_id,json=habitId,proto3" json:"habit_id,omitempty"`
	LoggedAt string `protobuf:"bytes,3,opt,name=logged_at,json=loggedAt,proto3" json:"logged_at,omitempty"`
	Notes    string `protobuf:"bytes,4,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *HabitLog) Reset() {
	*x = HabitLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HabitLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HabitLog) ProtoMessage() {}

func (x *HabitLog) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HabitLog.ProtoReflect.Descriptor instead.
func (*HabitLog) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{1}
}

func (x *HabitLog) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HabitLog) GetHabitId() string {
	if x != nil {
		return x.HabitId
	}
	return ""
}

func (x *HabitLog) GetLoggedAt() string {
	if x != nil {
		return x.LoggedAt
	}
	return ""
}

func (x *HabitLog) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type CreateHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32     `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string    `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Frequency   Frequency `protobuf:"varint,4,opt,name=frequency,proto3,enum=habittracker.Frequency" json:"frequency,omitempty"`
}

func (x *CreateHabitRequest) Reset() {
	*x = CreateHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHabitRequest) ProtoMessage() {}

func (x *CreateHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHabitRequest.ProtoReflect.Descriptor instead.
func (*CreateHabitRequest) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{2}
}

func (x *CreateHabitRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateHabitRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHabitRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateHabitRequest) GetFrequency() Frequency {
	if x != nil {
		return x.Frequency
	}
	return Frequency_DAILY
}

type CreateHabitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habit *Habit `protobuf:"bytes,1,opt,name=habit,proto3" json:"habit,omitempty"`
}

func (x *CreateHabitResponse) Reset() {
	*x = CreateHabitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHabitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHabitResponse) ProtoMessage() {}

func (x *CreateHabitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHabitResponse.ProtoReflect.Descriptor instead.
func (*CreateHabitResponse) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHabitResponse) GetHabit() *Habit {
	if x != nil {
		return x.Habit
	}
	return nil
}

type GetHabitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetHabitsRequest) Reset() {
	*x = GetHabitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitsRequest) ProtoMessage() {}

func (x *GetHabitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitsRequest.ProtoReflect.Descriptor instead.
func (*GetHabitsRequest) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{4}
}

func (x *GetHabitsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetHabitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Habits []*Habit `protobuf:"bytes,1,rep,name=habits,proto3" json:"habits,omitempty"`
}

func (x *GetHabitsResponse) Reset() {
	*x = GetHabitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitsResponse) ProtoMessage() {}

func (x *GetHabitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitsResponse.ProtoReflect.Descriptor instead.
func (*GetHabitsResponse) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{5}
}

func (x *GetHabitsResponse) GetHabits() []*Habit {
	if x != nil {
		return x.Habits
	}
	return nil
}

type LogHabitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitId string `protobuf:"bytes,1,opt,name=habit_id,json=habitId,proto3" json:"habit_id,omitempty"`
	Notes   string `protobuf:"bytes,2,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *LogHabitRequest) Reset() {
	*x = LogHabitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogHabitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogHabitRequest) ProtoMessage() {}

func (x *LogHabitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogHabitRequest.ProtoReflect.Descriptor instead.
func (*LogHabitRequest) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{6}
}

func (x *LogHabitRequest) GetHabitId() string {
	if x != nil {
		return x.HabitId
	}
	return ""
}

func (x *LogHabitRequest) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type LogHabitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitLog *HabitLog `protobuf:"bytes,1,opt,name=habit_log,json=habitLog,proto3" json:"habit_log,omitempty"`
}

func (x *LogHabitResponse) Reset() {
	*x = LogHabitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogHabitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogHabitResponse) ProtoMessage() {}

func (x *LogHabitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogHabitResponse.ProtoReflect.Descriptor instead.
func (*LogHabitResponse) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{7}
}

func (x *LogHabitResponse) GetHabitLog() *HabitLog {
	if x != nil {
		return x.HabitLog
	}
	return nil
}

type GetHabitLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitId string `protobuf:"bytes,1,opt,name=habit_id,json=habitId,proto3" json:"habit_id,omitempty"`
}

func (x *GetHabitLogsRequest) Reset() {
	*x = GetHabitLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitLogsRequest) ProtoMessage() {}

func (x *GetHabitLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitLogsRequest.ProtoReflect.Descriptor instead.
func (*GetHabitLogsRequest) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{8}
}

func (x *GetHabitLogsRequest) GetHabitId() string {
	if x != nil {
		return x.HabitId
	}
	return ""
}

type GetHabitLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HabitLogs []*HabitLog `protobuf:"bytes,1,rep,name=habit_logs,json=habitLogs,proto3" json:"habit_logs,omitempty"`
}

func (x *GetHabitLogsResponse) Reset() {
	*x = GetHabitLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Tracker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHabitLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHabitLogsResponse) ProtoMessage() {}

func (x *GetHabitLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Tracker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHabitLogsResponse.ProtoReflect.Descriptor instead.
func (*GetHabitLogsResponse) Descriptor() ([]byte, []int) {
	return file_Tracker_proto_rawDescGZIP(), []int{9}
}

func (x *GetHabitLogsResponse) GetHabitLogs() []*HabitLog {
	if x != nil {
		return x.HabitLogs
	}
	return nil
}

var File_Tracker_proto protoreflect.FileDescriptor

var file_Tracker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x68, 0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x22, 0xbc, 0x01,
	0x0a, 0x05, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x68, 0x0a, 0x08,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x17, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x46,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x09, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x79, 0x22, 0x40, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x05,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x06, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x73, 0x22, 0x42, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x68, 0x61, 0x62, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x68, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x68, 0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f,
	0x67, 0x73, 0x2a, 0x2f, 0x0a, 0x09, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45,
	0x45, 0x4b, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x4c,
	0x59, 0x10, 0x02, 0x32, 0xd9, 0x02, 0x0a, 0x13, 0x48, 0x61, 0x62, 0x69, 0x74, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x20, 0x2e, 0x68, 0x61, 0x62,
	0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x61, 0x62, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x68,
	0x61, 0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x62, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x68, 0x61, 0x62, 0x69,
	0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48,
	0x61, 0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x21, 0x2e, 0x68, 0x61, 0x62, 0x69, 0x74,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x62, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x68, 0x61,
	0x62, 0x69, 0x74, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61,
	0x62, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x61, 0x62, 0x69,
	0x74, 0x54, 0x72, 0x61, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Tracker_proto_rawDescOnce sync.Once
	file_Tracker_proto_rawDescData = file_Tracker_proto_rawDesc
)

func file_Tracker_proto_rawDescGZIP() []byte {
	file_Tracker_proto_rawDescOnce.Do(func() {
		file_Tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_Tracker_proto_rawDescData)
	})
	return file_Tracker_proto_rawDescData
}

var file_Tracker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Tracker_proto_goTypes = []interface{}{
	(Frequency)(0),               // 0: habittracker.Frequency
	(*Habit)(nil),                // 1: habittracker.Habit
	(*HabitLog)(nil),             // 2: habittracker.HabitLog
	(*CreateHabitRequest)(nil),   // 3: habittracker.CreateHabitRequest
	(*CreateHabitResponse)(nil),  // 4: habittracker.CreateHabitResponse
	(*GetHabitsRequest)(nil),     // 5: habittracker.GetHabitsRequest
	(*GetHabitsResponse)(nil),    // 6: habittracker.GetHabitsResponse
	(*LogHabitRequest)(nil),      // 7: habittracker.LogHabitRequest
	(*LogHabitResponse)(nil),     // 8: habittracker.LogHabitResponse
	(*GetHabitLogsRequest)(nil),  // 9: habittracker.GetHabitLogsRequest
	(*GetHabitLogsResponse)(nil), // 10: habittracker.GetHabitLogsResponse
}
var file_Tracker_proto_depIdxs = []int32{
	0,  // 0: habittracker.Habit.frequency:type_name -> habittracker.Frequency
	0,  // 1: habittracker.CreateHabitRequest.frequency:type_name -> habittracker.Frequency
	1,  // 2: habittracker.CreateHabitResponse.habit:type_name -> habittracker.Habit
	1,  // 3: habittracker.GetHabitsResponse.habits:type_name -> habittracker.Habit
	2,  // 4: habittracker.LogHabitResponse.habit_log:type_name -> habittracker.HabitLog
	2,  // 5: habittracker.GetHabitLogsResponse.habit_logs:type_name -> habittracker.HabitLog
	3,  // 6: habittracker.HabitTrackerService.CreateHabit:input_type -> habittracker.CreateHabitRequest
	5,  // 7: habittracker.HabitTrackerService.GetHabits:input_type -> habittracker.GetHabitsRequest
	7,  // 8: habittracker.HabitTrackerService.LogHabit:input_type -> habittracker.LogHabitRequest
	9,  // 9: habittracker.HabitTrackerService.GetHabitLogs:input_type -> habittracker.GetHabitLogsRequest
	4,  // 10: habittracker.HabitTrackerService.CreateHabit:output_type -> habittracker.CreateHabitResponse
	6,  // 11: habittracker.HabitTrackerService.GetHabits:output_type -> habittracker.GetHabitsResponse
	8,  // 12: habittracker.HabitTrackerService.LogHabit:output_type -> habittracker.LogHabitResponse
	10, // 13: habittracker.HabitTrackerService.GetHabitLogs:output_type -> habittracker.GetHabitLogsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_Tracker_proto_init() }
func file_Tracker_proto_init() {
	if File_Tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Habit); i {
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
		file_Tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HabitLog); i {
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
		file_Tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHabitRequest); i {
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
		file_Tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHabitResponse); i {
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
		file_Tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitsRequest); i {
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
		file_Tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitsResponse); i {
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
		file_Tracker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogHabitRequest); i {
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
		file_Tracker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogHabitResponse); i {
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
		file_Tracker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitLogsRequest); i {
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
		file_Tracker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHabitLogsResponse); i {
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
			RawDescriptor: file_Tracker_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Tracker_proto_goTypes,
		DependencyIndexes: file_Tracker_proto_depIdxs,
		EnumInfos:         file_Tracker_proto_enumTypes,
		MessageInfos:      file_Tracker_proto_msgTypes,
	}.Build()
	File_Tracker_proto = out.File
	file_Tracker_proto_rawDesc = nil
	file_Tracker_proto_goTypes = nil
	file_Tracker_proto_depIdxs = nil
}
