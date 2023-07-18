// B站开放接口 - 公共参数管理 bilibili-webapp

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.0
// source: bilibili-webapp/common.proto

package bilibili_webapp

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pn    int64 `protobuf:"varint,1,opt,name=pn,proto3" json:"pn,omitempty"`
	Ps    int64 `protobuf:"varint,2,opt,name=ps,proto3" json:"ps,omitempty"`
	Total int64 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_common_proto_rawDescGZIP(), []int{0}
}

func (x *Page) GetPn() int64 {
	if x != nil {
		return x.Pn
	}
	return 0
}

func (x *Page) GetPs() int64 {
	if x != nil {
		return x.Ps
	}
	return 0
}

func (x *Page) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AnthologyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ImageUrl    string `protobuf:"bytes,3,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	UpdateTime  int64  `protobuf:"varint,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Ctime       int64  `protobuf:"varint,5,opt,name=ctime,proto3" json:"ctime,omitempty"`
	PublishTime int64  `protobuf:"varint,6,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Summary     string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	Words       int64  `protobuf:"varint,8,opt,name=words,proto3" json:"words,omitempty"`
	Read        int64  `protobuf:"varint,9,opt,name=read,proto3" json:"read,omitempty"`
	State       int64  `protobuf:"varint,10,opt,name=state,proto3" json:"state,omitempty"`
	ApplyTime   string `protobuf:"bytes,11,opt,name=apply_time,json=applyTime,proto3" json:"apply_time,omitempty"`
	CheckTime   string `protobuf:"bytes,12,opt,name=check_time,json=checkTime,proto3" json:"check_time,omitempty"`
	Total       int64  `protobuf:"varint,13,opt,name=total,proto3" json:"total,omitempty"`
	Reason      string `protobuf:"bytes,14,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *AnthologyData) Reset() {
	*x = AnthologyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnthologyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnthologyData) ProtoMessage() {}

func (x *AnthologyData) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnthologyData.ProtoReflect.Descriptor instead.
func (*AnthologyData) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_common_proto_rawDescGZIP(), []int{1}
}

func (x *AnthologyData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AnthologyData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AnthologyData) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *AnthologyData) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *AnthologyData) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *AnthologyData) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

func (x *AnthologyData) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *AnthologyData) GetWords() int64 {
	if x != nil {
		return x.Words
	}
	return 0
}

func (x *AnthologyData) GetRead() int64 {
	if x != nil {
		return x.Read
	}
	return 0
}

func (x *AnthologyData) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *AnthologyData) GetApplyTime() string {
	if x != nil {
		return x.ApplyTime
	}
	return ""
}

func (x *AnthologyData) GetCheckTime() string {
	if x != nil {
		return x.CheckTime
	}
	return ""
}

func (x *AnthologyData) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AnthologyData) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ArticleCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId int64              `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Children []*ArticleCategory `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ArticleCategory) Reset() {
	*x = ArticleCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCategory) ProtoMessage() {}

func (x *ArticleCategory) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCategory.ProtoReflect.Descriptor instead.
func (*ArticleCategory) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_common_proto_rawDescGZIP(), []int{2}
}

func (x *ArticleCategory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleCategory) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ArticleCategory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArticleCategory) GetChildren() []*ArticleCategory {
	if x != nil {
		return x.Children
	}
	return nil
}

type ArticleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category     *ArticleCategory   `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Title        string             `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Summary      string             `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	BannerUrl    string             `protobuf:"bytes,5,opt,name=banner_url,json=bannerUrl,proto3" json:"banner_url,omitempty"`
	TemplateId   int64              `protobuf:"varint,6,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	State        int64              `protobuf:"varint,7,opt,name=state,proto3" json:"state,omitempty"`
	ImageUrls    []string           `protobuf:"bytes,8,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	PublishTime  int64              `protobuf:"varint,9,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`
	Ctime        int64              `protobuf:"varint,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Stats        *ArticleData_Stats `protobuf:"bytes,11,opt,name=stats,proto3" json:"stats,omitempty"`
	Reason       string             `protobuf:"bytes,12,opt,name=reason,proto3" json:"reason,omitempty"`
	Words        int64              `protobuf:"varint,13,opt,name=words,proto3" json:"words,omitempty"`
	List         *AnthologyData     `protobuf:"bytes,14,opt,name=list,proto3" json:"list,omitempty"`
	TopVideoBvid string             `protobuf:"bytes,15,opt,name=top_video_bvid,json=topVideoBvid,proto3" json:"top_video_bvid,omitempty"`
}

func (x *ArticleData) Reset() {
	*x = ArticleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleData) ProtoMessage() {}

func (x *ArticleData) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleData.ProtoReflect.Descriptor instead.
func (*ArticleData) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_common_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ArticleData) GetCategory() *ArticleCategory {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ArticleData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ArticleData) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ArticleData) GetBannerUrl() string {
	if x != nil {
		return x.BannerUrl
	}
	return ""
}

func (x *ArticleData) GetTemplateId() int64 {
	if x != nil {
		return x.TemplateId
	}
	return 0
}

func (x *ArticleData) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *ArticleData) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *ArticleData) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

func (x *ArticleData) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *ArticleData) GetStats() *ArticleData_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *ArticleData) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *ArticleData) GetWords() int64 {
	if x != nil {
		return x.Words
	}
	return 0
}

func (x *ArticleData) GetList() *AnthologyData {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ArticleData) GetTopVideoBvid() string {
	if x != nil {
		return x.TopVideoBvid
	}
	return ""
}

type ArticleData_Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	View     int64 `protobuf:"varint,1,opt,name=view,proto3" json:"view,omitempty"`
	Favorite int64 `protobuf:"varint,2,opt,name=favorite,proto3" json:"favorite,omitempty"`
	Like     int64 `protobuf:"varint,3,opt,name=like,proto3" json:"like,omitempty"`
	Dislike  int64 `protobuf:"varint,4,opt,name=dislike,proto3" json:"dislike,omitempty"`
	Reply    int64 `protobuf:"varint,5,opt,name=reply,proto3" json:"reply,omitempty"`
	Share    int64 `protobuf:"varint,6,opt,name=share,proto3" json:"share,omitempty"`
	Coin     int64 `protobuf:"varint,7,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (x *ArticleData_Stats) Reset() {
	*x = ArticleData_Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bilibili_webapp_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArticleData_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleData_Stats) ProtoMessage() {}

func (x *ArticleData_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_bilibili_webapp_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleData_Stats.ProtoReflect.Descriptor instead.
func (*ArticleData_Stats) Descriptor() ([]byte, []int) {
	return file_bilibili_webapp_common_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ArticleData_Stats) GetView() int64 {
	if x != nil {
		return x.View
	}
	return 0
}

func (x *ArticleData_Stats) GetFavorite() int64 {
	if x != nil {
		return x.Favorite
	}
	return 0
}

func (x *ArticleData_Stats) GetLike() int64 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *ArticleData_Stats) GetDislike() int64 {
	if x != nil {
		return x.Dislike
	}
	return 0
}

func (x *ArticleData_Stats) GetReply() int64 {
	if x != nil {
		return x.Reply
	}
	return 0
}

func (x *ArticleData_Stats) GetShare() int64 {
	if x != nil {
		return x.Share
	}
	return 0
}

func (x *ArticleData_Stats) GetCoin() int64 {
	if x != nil {
		return x.Coin
	}
	return 0
}

var File_bilibili_webapp_common_proto protoreflect.FileDescriptor

var file_bilibili_webapp_common_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x77, 0x65, 0x62, 0x61, 0x70,
	0x70, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3c, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x70, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0xf0, 0x02, 0x0a, 0x0d, 0x41, 0x6e, 0x74, 0x68, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x22, 0x94, 0x01, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64,
	0x72, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xaf, 0x05, 0x0a, 0x0b, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x61, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x6c,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x41, 0x6e, 0x74, 0x68, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x0e, 0x74, 0x6f, 0x70, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x62, 0x76, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x70, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x42,
	0x76, 0x69, 0x64, 0x1a, 0xa5, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x76, 0x69, 0x65,
	0x77, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c, 0x69, 0x6b,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x64, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x69, 0x6e, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6c, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x2d, 0x77,
	0x65, 0x62, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bilibili_webapp_common_proto_rawDescOnce sync.Once
	file_bilibili_webapp_common_proto_rawDescData = file_bilibili_webapp_common_proto_rawDesc
)

func file_bilibili_webapp_common_proto_rawDescGZIP() []byte {
	file_bilibili_webapp_common_proto_rawDescOnce.Do(func() {
		file_bilibili_webapp_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_bilibili_webapp_common_proto_rawDescData)
	})
	return file_bilibili_webapp_common_proto_rawDescData
}

var file_bilibili_webapp_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bilibili_webapp_common_proto_goTypes = []interface{}{
	(*Page)(nil),              // 0: member.bilibili.com.Page
	(*AnthologyData)(nil),     // 1: member.bilibili.com.AnthologyData
	(*ArticleCategory)(nil),   // 2: member.bilibili.com.ArticleCategory
	(*ArticleData)(nil),       // 3: member.bilibili.com.ArticleData
	(*ArticleData_Stats)(nil), // 4: member.bilibili.com.ArticleData.Stats
}
var file_bilibili_webapp_common_proto_depIdxs = []int32{
	2, // 0: member.bilibili.com.ArticleCategory.children:type_name -> member.bilibili.com.ArticleCategory
	2, // 1: member.bilibili.com.ArticleData.category:type_name -> member.bilibili.com.ArticleCategory
	4, // 2: member.bilibili.com.ArticleData.stats:type_name -> member.bilibili.com.ArticleData.Stats
	1, // 3: member.bilibili.com.ArticleData.list:type_name -> member.bilibili.com.AnthologyData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bilibili_webapp_common_proto_init() }
func file_bilibili_webapp_common_proto_init() {
	if File_bilibili_webapp_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bilibili_webapp_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_bilibili_webapp_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnthologyData); i {
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
		file_bilibili_webapp_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleCategory); i {
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
		file_bilibili_webapp_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleData); i {
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
		file_bilibili_webapp_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArticleData_Stats); i {
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
			RawDescriptor: file_bilibili_webapp_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bilibili_webapp_common_proto_goTypes,
		DependencyIndexes: file_bilibili_webapp_common_proto_depIdxs,
		MessageInfos:      file_bilibili_webapp_common_proto_msgTypes,
	}.Build()
	File_bilibili_webapp_common_proto = out.File
	file_bilibili_webapp_common_proto_rawDesc = nil
	file_bilibili_webapp_common_proto_goTypes = nil
	file_bilibili_webapp_common_proto_depIdxs = nil
}
