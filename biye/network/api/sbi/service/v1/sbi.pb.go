// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: sbi.proto

package v1

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

type GetOrdersByUserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetOrdersByUserIdReq) Reset() {
	*x = GetOrdersByUserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersByUserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByUserIdReq) ProtoMessage() {}

func (x *GetOrdersByUserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByUserIdReq.ProtoReflect.Descriptor instead.
func (*GetOrdersByUserIdReq) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrdersByUserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetOrdersByUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*GetOrdersByUserReply_Orders `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrdersByUserReply) Reset() {
	*x = GetOrdersByUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersByUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByUserReply) ProtoMessage() {}

func (x *GetOrdersByUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByUserReply.ProtoReflect.Descriptor instead.
func (*GetOrdersByUserReply) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrdersByUserReply) GetOrders() []*GetOrdersByUserReply_Orders {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateOrdersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders *CreateOrdersReq_Orders `protobuf:"bytes,1,opt,name=orders,proto3" json:"orders,omitempty"`
}

func (x *CreateOrdersReq) Reset() {
	*x = CreateOrdersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrdersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrdersReq) ProtoMessage() {}

func (x *CreateOrdersReq) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrdersReq.ProtoReflect.Descriptor instead.
func (*CreateOrdersReq) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrdersReq) GetOrders() *CreateOrdersReq_Orders {
	if x != nil {
		return x.Orders
	}
	return nil
}

type CreateOrdersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOrdersReply) Reset() {
	*x = CreateOrdersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrdersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrdersReply) ProtoMessage() {}

func (x *CreateOrdersReply) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrdersReply.ProtoReflect.Descriptor instead.
func (*CreateOrdersReply) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{3}
}

type GetSbiByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSbiByIdReq) Reset() {
	*x = GetSbiByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSbiByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSbiByIdReq) ProtoMessage() {}

func (x *GetSbiByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSbiByIdReq.ProtoReflect.Descriptor instead.
func (*GetSbiByIdReq) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{4}
}

func (x *GetSbiByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSbiByIdReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sbi *GetSbiByIdReply_Sbi `protobuf:"bytes,1,opt,name=sbi,proto3" json:"sbi,omitempty"`
}

func (x *GetSbiByIdReply) Reset() {
	*x = GetSbiByIdReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSbiByIdReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSbiByIdReply) ProtoMessage() {}

func (x *GetSbiByIdReply) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSbiByIdReply.ProtoReflect.Descriptor instead.
func (*GetSbiByIdReply) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{5}
}

func (x *GetSbiByIdReply) GetSbi() *GetSbiByIdReply_Sbi {
	if x != nil {
		return x.Sbi
	}
	return nil
}

type ListSbiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  int64 `protobuf:"varint,1,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListSbiReq) Reset() {
	*x = ListSbiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSbiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSbiReq) ProtoMessage() {}

func (x *ListSbiReq) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSbiReq.ProtoReflect.Descriptor instead.
func (*ListSbiReq) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{6}
}

func (x *ListSbiReq) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListSbiReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSbiReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64               `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Sbis  []*ListSbiReply_Sbi `protobuf:"bytes,2,rep,name=sbis,proto3" json:"sbis,omitempty"`
}

func (x *ListSbiReply) Reset() {
	*x = ListSbiReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSbiReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSbiReply) ProtoMessage() {}

func (x *ListSbiReply) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSbiReply.ProtoReflect.Descriptor instead.
func (*ListSbiReply) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{7}
}

func (x *ListSbiReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListSbiReply) GetSbis() []*ListSbiReply_Sbi {
	if x != nil {
		return x.Sbis
	}
	return nil
}

type GetOrdersByUserReply_Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SbiId    int64   `protobuf:"varint,2,opt,name=sbi_id,json=sbiId,proto3" json:"sbi_id,omitempty"`
	Price    float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Receiver string  `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Address  string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Mobile   string  `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *GetOrdersByUserReply_Orders) Reset() {
	*x = GetOrdersByUserReply_Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersByUserReply_Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByUserReply_Orders) ProtoMessage() {}

func (x *GetOrdersByUserReply_Orders) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByUserReply_Orders.ProtoReflect.Descriptor instead.
func (*GetOrdersByUserReply_Orders) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetOrdersByUserReply_Orders) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetOrdersByUserReply_Orders) GetSbiId() int64 {
	if x != nil {
		return x.SbiId
	}
	return 0
}

func (x *GetOrdersByUserReply_Orders) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetOrdersByUserReply_Orders) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *GetOrdersByUserReply_Orders) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetOrdersByUserReply_Orders) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type CreateOrdersReq_Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SbiId    int64   `protobuf:"varint,2,opt,name=sbi_id,json=sbiId,proto3" json:"sbi_id,omitempty"`
	Price    float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	Receiver string  `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Address  string  `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Mobile   string  `protobuf:"bytes,6,opt,name=mobile,proto3" json:"mobile,omitempty"`
}

func (x *CreateOrdersReq_Orders) Reset() {
	*x = CreateOrdersReq_Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrdersReq_Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrdersReq_Orders) ProtoMessage() {}

func (x *CreateOrdersReq_Orders) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrdersReq_Orders.ProtoReflect.Descriptor instead.
func (*CreateOrdersReq_Orders) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CreateOrdersReq_Orders) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrdersReq_Orders) GetSbiId() int64 {
	if x != nil {
		return x.SbiId
	}
	return 0
}

func (x *CreateOrdersReq_Orders) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateOrdersReq_Orders) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *CreateOrdersReq_Orders) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateOrdersReq_Orders) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

type GetSbiByIdReply_Sbi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist   string  `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	CreateAt string  `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *GetSbiByIdReply_Sbi) Reset() {
	*x = GetSbiByIdReply_Sbi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSbiByIdReply_Sbi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSbiByIdReply_Sbi) ProtoMessage() {}

func (x *GetSbiByIdReply_Sbi) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSbiByIdReply_Sbi.ProtoReflect.Descriptor instead.
func (*GetSbiByIdReply_Sbi) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetSbiByIdReply_Sbi) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetSbiByIdReply_Sbi) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetSbiByIdReply_Sbi) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *GetSbiByIdReply_Sbi) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetSbiByIdReply_Sbi) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type ListSbiReply_Sbi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist   string  `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Price    float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	CreateAt string  `protobuf:"bytes,5,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
}

func (x *ListSbiReply_Sbi) Reset() {
	*x = ListSbiReply_Sbi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sbi_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSbiReply_Sbi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSbiReply_Sbi) ProtoMessage() {}

func (x *ListSbiReply_Sbi) ProtoReflect() protoreflect.Message {
	mi := &file_sbi_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSbiReply_Sbi.ProtoReflect.Descriptor instead.
func (*ListSbiReply_Sbi) Descriptor() ([]byte, []int) {
	return file_sbi_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListSbiReply_Sbi) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListSbiReply_Sbi) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListSbiReply_Sbi) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *ListSbiReply_Sbi) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ListSbiReply_Sbi) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

var File_sbi_proto protoreflect.FileDescriptor

var file_sbi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xef,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0xa0, 0x01,
	0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x22, 0xe5, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x06, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0xa0, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xbd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x78, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2d, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73,
	0x1a, 0x78, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x32, 0xeb, 0x01, 0x0a, 0x05, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x12, 0x0d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x72, 0x61, 0x64,
	0x75, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sbi_proto_rawDescOnce sync.Once
	file_sbi_proto_rawDescData = file_sbi_proto_rawDesc
)

func file_sbi_proto_rawDescGZIP() []byte {
	file_sbi_proto_rawDescOnce.Do(func() {
		file_sbi_proto_rawDescData = protoimpl.X.CompressGZIP(file_sbi_proto_rawDescData)
	})
	return file_sbi_proto_rawDescData
}

var file_sbi_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_sbi_proto_goTypes = []interface{}{
	(*GetOrdersByUserIdReq)(nil),        // 0: GetOrdersByUserIdReq
	(*GetOrdersByUserReply)(nil),        // 1: GetOrdersByUserReply
	(*CreateOrdersReq)(nil),             // 2: CreateOrdersReq
	(*CreateOrdersReply)(nil),           // 3: CreateOrdersReply
	(*GetSbiByIdReq)(nil),               // 4: GetSbiByIdReq
	(*GetSbiByIdReply)(nil),             // 5: GetSbiByIdReply
	(*ListSbiReq)(nil),                  // 6: ListSbiReq
	(*ListSbiReply)(nil),                // 7: ListSbiReply
	(*GetOrdersByUserReply_Orders)(nil), // 8: GetOrdersByUserReply.Orders
	(*CreateOrdersReq_Orders)(nil),      // 9: CreateOrdersReq.Orders
	(*GetSbiByIdReply_Sbi)(nil),         // 10: GetSbiByIdReply.Sbi
	(*ListSbiReply_Sbi)(nil),            // 11: ListSbiReply.Sbi
}
var file_sbi_proto_depIdxs = []int32{
	8,  // 0: GetOrdersByUserReply.orders:type_name -> GetOrdersByUserReply.Orders
	9,  // 1: CreateOrdersReq.orders:type_name -> CreateOrdersReq.Orders
	10, // 2: GetSbiByIdReply.sbi:type_name -> GetSbiByIdReply.Sbi
	11, // 3: ListSbiReply.sbis:type_name -> ListSbiReply.Sbi
	6,  // 4: Sbi.ListSbi:input_type -> ListSbiReq
	4,  // 5: Sbi.GetSbiById:input_type -> GetSbiByIdReq
	2,  // 6: Sbi.CreateOrders:input_type -> CreateOrdersReq
	0,  // 7: Sbi.GetOrdersByUserId:input_type -> GetOrdersByUserIdReq
	7,  // 8: Sbi.ListSbi:output_type -> ListSbiReply
	5,  // 9: Sbi.GetSbiById:output_type -> GetSbiByIdReply
	3,  // 10: Sbi.CreateOrders:output_type -> CreateOrdersReply
	1,  // 11: Sbi.GetOrdersByUserId:output_type -> GetOrdersByUserReply
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_sbi_proto_init() }
func file_sbi_proto_init() {
	if File_sbi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sbi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersByUserIdReq); i {
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
		file_sbi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersByUserReply); i {
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
		file_sbi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrdersReq); i {
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
		file_sbi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrdersReply); i {
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
		file_sbi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSbiByIdReq); i {
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
		file_sbi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSbiByIdReply); i {
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
		file_sbi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSbiReq); i {
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
		file_sbi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSbiReply); i {
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
		file_sbi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersByUserReply_Orders); i {
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
		file_sbi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrdersReq_Orders); i {
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
		file_sbi_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSbiByIdReply_Sbi); i {
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
		file_sbi_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSbiReply_Sbi); i {
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
			RawDescriptor: file_sbi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sbi_proto_goTypes,
		DependencyIndexes: file_sbi_proto_depIdxs,
		MessageInfos:      file_sbi_proto_msgTypes,
	}.Build()
	File_sbi_proto = out.File
	file_sbi_proto_rawDesc = nil
	file_sbi_proto_goTypes = nil
	file_sbi_proto_depIdxs = nil
}