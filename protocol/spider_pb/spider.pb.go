// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: spider.proto

package spider_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The request message containing the user's name.
type SpiderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsMobile bool  `protobuf:"varint,1,opt,name=is_mobile,json=isMobile,proto3" json:"is_mobile,omitempty"`
	Size     int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *SpiderReq) Reset() {
	*x = SpiderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpiderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpiderReq) ProtoMessage() {}

func (x *SpiderReq) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpiderReq.ProtoReflect.Descriptor instead.
func (*SpiderReq) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{0}
}

func (x *SpiderReq) GetIsMobile() bool {
	if x != nil {
		return x.IsMobile
	}
	return false
}

func (x *SpiderReq) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

// The response message containing the greetings
type SpiderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url            string           `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error          string           `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
	WeiboHotList   []*WeiboHot      `protobuf:"bytes,2,rep,name=weiboHotList,proto3" json:"weiboHotList,omitempty"`
	D36KrHotList   []*D36KrHot      `protobuf:"bytes,3,rep,name=d36KrHotList,proto3" json:"d36KrHotList,omitempty"`
	WallStreetNews []*WallStreetNew `protobuf:"bytes,4,rep,name=wallStreetNews,proto3" json:"wallStreetNews,omitempty"`
	ZhihuHotList   []*ZhihuHot      `protobuf:"bytes,5,rep,name=zhihuHotList,proto3" json:"zhihuHotList,omitempty"`
}

func (x *SpiderResp) Reset() {
	*x = SpiderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpiderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpiderResp) ProtoMessage() {}

func (x *SpiderResp) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpiderResp.ProtoReflect.Descriptor instead.
func (*SpiderResp) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{1}
}

func (x *SpiderResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SpiderResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SpiderResp) GetWeiboHotList() []*WeiboHot {
	if x != nil {
		return x.WeiboHotList
	}
	return nil
}

func (x *SpiderResp) GetD36KrHotList() []*D36KrHot {
	if x != nil {
		return x.D36KrHotList
	}
	return nil
}

func (x *SpiderResp) GetWallStreetNews() []*WallStreetNew {
	if x != nil {
		return x.WallStreetNews
	}
	return nil
}

func (x *SpiderResp) GetZhihuHotList() []*ZhihuHot {
	if x != nil {
		return x.ZhihuHotList
	}
	return nil
}

type WeiboHot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Hot   int64  `protobuf:"varint,3,opt,name=hot,proto3" json:"hot,omitempty"`
	Rank  int64  `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *WeiboHot) Reset() {
	*x = WeiboHot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeiboHot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeiboHot) ProtoMessage() {}

func (x *WeiboHot) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeiboHot.ProtoReflect.Descriptor instead.
func (*WeiboHot) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{2}
}

func (x *WeiboHot) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WeiboHot) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WeiboHot) GetHot() int64 {
	if x != nil {
		return x.Hot
	}
	return 0
}

func (x *WeiboHot) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type D36KrHot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url   string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Rank  int64  `protobuf:"varint,3,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *D36KrHot) Reset() {
	*x = D36KrHot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *D36KrHot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*D36KrHot) ProtoMessage() {}

func (x *D36KrHot) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use D36KrHot.ProtoReflect.Descriptor instead.
func (*D36KrHot) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{3}
}

func (x *D36KrHot) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *D36KrHot) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *D36KrHot) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

type WallStreetNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *WallStreetNew) Reset() {
	*x = WallStreetNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WallStreetNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WallStreetNew) ProtoMessage() {}

func (x *WallStreetNew) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WallStreetNew.ProtoReflect.Descriptor instead.
func (*WallStreetNew) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{4}
}

func (x *WallStreetNew) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WallStreetNew) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *WallStreetNew) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ZhihuHot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Excerpt string `protobuf:"bytes,3,opt,name=excerpt,proto3" json:"excerpt,omitempty"`
	Rank    int64  `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	Created int64  `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *ZhihuHot) Reset() {
	*x = ZhihuHot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZhihuHot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZhihuHot) ProtoMessage() {}

func (x *ZhihuHot) ProtoReflect() protoreflect.Message {
	mi := &file_spider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZhihuHot.ProtoReflect.Descriptor instead.
func (*ZhihuHot) Descriptor() ([]byte, []int) {
	return file_spider_proto_rawDescGZIP(), []int{5}
}

func (x *ZhihuHot) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ZhihuHot) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ZhihuHot) GetExcerpt() string {
	if x != nil {
		return x.Excerpt
	}
	return ""
}

func (x *ZhihuHot) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *ZhihuHot) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

var File_spider_proto protoreflect.FileDescriptor

var file_spider_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x09, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x0a, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x0c, 0x77, 0x65,
	0x69, 0x62, 0x6f, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x69, 0x62, 0x6f, 0x48, 0x6f,
	0x74, 0x52, 0x0c, 0x77, 0x65, 0x69, 0x62, 0x6f, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x33, 0x0a, 0x0c, 0x64, 0x33, 0x36, 0x4b, 0x72, 0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x33,
	0x36, 0x4b, 0x72, 0x48, 0x6f, 0x74, 0x52, 0x0c, 0x64, 0x33, 0x36, 0x4b, 0x72, 0x48, 0x6f, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x52, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x65,
	0x77, 0x73, 0x12, 0x33, 0x0a, 0x0c, 0x7a, 0x68, 0x69, 0x68, 0x75, 0x48, 0x6f, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x5a, 0x68, 0x69, 0x68, 0x75, 0x48, 0x6f, 0x74, 0x52, 0x0c, 0x7a, 0x68, 0x69, 0x68, 0x75,
	0x48, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x08, 0x57, 0x65, 0x69, 0x62, 0x6f,
	0x48, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x68,
	0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x68, 0x6f, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x22, 0x46, 0x0a, 0x08, 0x44, 0x33, 0x36, 0x4b, 0x72, 0x48, 0x6f, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x22, 0x51, 0x0a, 0x0d, 0x57, 0x61, 0x6c,
	0x6c, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x08,
	0x5a, 0x68, 0x69, 0x68, 0x75, 0x48, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x65, 0x72, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x78, 0x63, 0x65, 0x72, 0x70, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61,
	0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0xdb, 0x02, 0x0a, 0x0d, 0x53, 0x70, 0x69,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x43, 0x4e, 0x42, 0x69, 0x6e, 0x67, 0x57, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x53,
	0x42, 0x69, 0x6e, 0x67, 0x57, 0x61, 0x6c, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x57, 0x65, 0x69, 0x62, 0x6f, 0x48, 0x6f, 0x74,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x44, 0x33, 0x36, 0x4b, 0x72,
	0x48, 0x6f, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0e, 0x57, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x5a, 0x68, 0x69, 0x68, 0x75, 0x48, 0x6f, 0x74, 0x12,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x70, 0x69, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x70, 0x62, 0x2f, 0x3b, 0x73, 0x70, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spider_proto_rawDescOnce sync.Once
	file_spider_proto_rawDescData = file_spider_proto_rawDesc
)

func file_spider_proto_rawDescGZIP() []byte {
	file_spider_proto_rawDescOnce.Do(func() {
		file_spider_proto_rawDescData = protoimpl.X.CompressGZIP(file_spider_proto_rawDescData)
	})
	return file_spider_proto_rawDescData
}

var file_spider_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spider_proto_goTypes = []interface{}{
	(*SpiderReq)(nil),     // 0: proto.SpiderReq
	(*SpiderResp)(nil),    // 1: proto.SpiderResp
	(*WeiboHot)(nil),      // 2: proto.WeiboHot
	(*D36KrHot)(nil),      // 3: proto.D36KrHot
	(*WallStreetNew)(nil), // 4: proto.WallStreetNew
	(*ZhihuHot)(nil),      // 5: proto.ZhihuHot
}
var file_spider_proto_depIdxs = []int32{
	2,  // 0: proto.SpiderResp.weiboHotList:type_name -> proto.WeiboHot
	3,  // 1: proto.SpiderResp.d36KrHotList:type_name -> proto.D36KrHot
	4,  // 2: proto.SpiderResp.wallStreetNews:type_name -> proto.WallStreetNew
	5,  // 3: proto.SpiderResp.zhihuHotList:type_name -> proto.ZhihuHot
	0,  // 4: proto.SpiderService.GetCNBingWallPaper:input_type -> proto.SpiderReq
	0,  // 5: proto.SpiderService.GetUSBingWallPaper:input_type -> proto.SpiderReq
	0,  // 6: proto.SpiderService.WeiboHot:input_type -> proto.SpiderReq
	0,  // 7: proto.SpiderService.D36KrHot:input_type -> proto.SpiderReq
	0,  // 8: proto.SpiderService.WallStreetNews:input_type -> proto.SpiderReq
	0,  // 9: proto.SpiderService.ZhihuHot:input_type -> proto.SpiderReq
	1,  // 10: proto.SpiderService.GetCNBingWallPaper:output_type -> proto.SpiderResp
	1,  // 11: proto.SpiderService.GetUSBingWallPaper:output_type -> proto.SpiderResp
	1,  // 12: proto.SpiderService.WeiboHot:output_type -> proto.SpiderResp
	1,  // 13: proto.SpiderService.D36KrHot:output_type -> proto.SpiderResp
	1,  // 14: proto.SpiderService.WallStreetNews:output_type -> proto.SpiderResp
	1,  // 15: proto.SpiderService.ZhihuHot:output_type -> proto.SpiderResp
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_spider_proto_init() }
func file_spider_proto_init() {
	if File_spider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpiderReq); i {
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
		file_spider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpiderResp); i {
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
		file_spider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeiboHot); i {
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
		file_spider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*D36KrHot); i {
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
		file_spider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WallStreetNew); i {
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
		file_spider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZhihuHot); i {
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
			RawDescriptor: file_spider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spider_proto_goTypes,
		DependencyIndexes: file_spider_proto_depIdxs,
		MessageInfos:      file_spider_proto_msgTypes,
	}.Build()
	File_spider_proto = out.File
	file_spider_proto_rawDesc = nil
	file_spider_proto_goTypes = nil
	file_spider_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpiderServiceClient is the client API for SpiderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpiderServiceClient interface {
	// Sends a greeting
	GetCNBingWallPaper(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
	GetUSBingWallPaper(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
	WeiboHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
	D36KrHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
	WallStreetNews(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
	ZhihuHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error)
}

type spiderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpiderServiceClient(cc grpc.ClientConnInterface) SpiderServiceClient {
	return &spiderServiceClient{cc}
}

func (c *spiderServiceClient) GetCNBingWallPaper(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/GetCNBingWallPaper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderServiceClient) GetUSBingWallPaper(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/GetUSBingWallPaper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderServiceClient) WeiboHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/WeiboHot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderServiceClient) D36KrHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/D36KrHot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderServiceClient) WallStreetNews(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/WallStreetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderServiceClient) ZhihuHot(ctx context.Context, in *SpiderReq, opts ...grpc.CallOption) (*SpiderResp, error) {
	out := new(SpiderResp)
	err := c.cc.Invoke(ctx, "/proto.SpiderService/ZhihuHot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpiderServiceServer is the server API for SpiderService service.
type SpiderServiceServer interface {
	// Sends a greeting
	GetCNBingWallPaper(context.Context, *SpiderReq) (*SpiderResp, error)
	GetUSBingWallPaper(context.Context, *SpiderReq) (*SpiderResp, error)
	WeiboHot(context.Context, *SpiderReq) (*SpiderResp, error)
	D36KrHot(context.Context, *SpiderReq) (*SpiderResp, error)
	WallStreetNews(context.Context, *SpiderReq) (*SpiderResp, error)
	ZhihuHot(context.Context, *SpiderReq) (*SpiderResp, error)
}

// UnimplementedSpiderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSpiderServiceServer struct {
}

func (*UnimplementedSpiderServiceServer) GetCNBingWallPaper(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCNBingWallPaper not implemented")
}
func (*UnimplementedSpiderServiceServer) GetUSBingWallPaper(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUSBingWallPaper not implemented")
}
func (*UnimplementedSpiderServiceServer) WeiboHot(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WeiboHot not implemented")
}
func (*UnimplementedSpiderServiceServer) D36KrHot(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method D36KrHot not implemented")
}
func (*UnimplementedSpiderServiceServer) WallStreetNews(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WallStreetNews not implemented")
}
func (*UnimplementedSpiderServiceServer) ZhihuHot(context.Context, *SpiderReq) (*SpiderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ZhihuHot not implemented")
}

func RegisterSpiderServiceServer(s *grpc.Server, srv SpiderServiceServer) {
	s.RegisterService(&_SpiderService_serviceDesc, srv)
}

func _SpiderService_GetCNBingWallPaper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).GetCNBingWallPaper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/GetCNBingWallPaper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).GetCNBingWallPaper(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiderService_GetUSBingWallPaper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).GetUSBingWallPaper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/GetUSBingWallPaper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).GetUSBingWallPaper(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiderService_WeiboHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).WeiboHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/WeiboHot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).WeiboHot(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiderService_D36KrHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).D36KrHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/D36KrHot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).D36KrHot(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiderService_WallStreetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).WallStreetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/WallStreetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).WallStreetNews(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpiderService_ZhihuHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpiderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpiderServiceServer).ZhihuHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SpiderService/ZhihuHot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpiderServiceServer).ZhihuHot(ctx, req.(*SpiderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpiderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SpiderService",
	HandlerType: (*SpiderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCNBingWallPaper",
			Handler:    _SpiderService_GetCNBingWallPaper_Handler,
		},
		{
			MethodName: "GetUSBingWallPaper",
			Handler:    _SpiderService_GetUSBingWallPaper_Handler,
		},
		{
			MethodName: "WeiboHot",
			Handler:    _SpiderService_WeiboHot_Handler,
		},
		{
			MethodName: "D36KrHot",
			Handler:    _SpiderService_D36KrHot_Handler,
		},
		{
			MethodName: "WallStreetNews",
			Handler:    _SpiderService_WallStreetNews_Handler,
		},
		{
			MethodName: "ZhihuHot",
			Handler:    _SpiderService_ZhihuHot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spider.proto",
}
