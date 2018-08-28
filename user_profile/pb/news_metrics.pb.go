// Code generated by protoc-gen-go. DO NOT EDIT.
// source: news_metrics.proto

/*
Package userprofile is a generated protocol buffer package.

It is generated from these files:
	news_metrics.proto

It has these top-level messages:
	NewsCard
	NewsSwipe
	NewsNotif
	NewsEngagement
	NewsClickViewMetric
	NewsEngagementMetric
	RedisNewsMessage
	RedisNewsMetric
*/
package userprofile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Messages
type NewsCard struct {
	Ig   map[string]*NewsClickViewMetric `protobuf:"bytes,1,rep,name=ig" json:"ig,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Loc  map[string]*NewsClickViewMetric `protobuf:"bytes,2,rep,name=loc" json:"loc,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ner  map[string]*NewsClickViewMetric `protobuf:"bytes,3,rep,name=ner" json:"ner,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type map[string]*NewsClickViewMetric `protobuf:"bytes,4,rep,name=type" json:"type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Src  map[string]*NewsClickViewMetric `protobuf:"bytes,5,rep,name=src" json:"src,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pqs  map[string]*NewsClickViewMetric `protobuf:"bytes,6,rep,name=pqs" json:"pqs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nsfw map[string]*NewsClickViewMetric `protobuf:"bytes,7,rep,name=nsfw" json:"nsfw,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Lang map[string]*NewsClickViewMetric `protobuf:"bytes,8,rep,name=lang" json:"lang,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feed map[string]*NewsClickViewMetric `protobuf:"bytes,9,rep,name=feed" json:"feed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Grp  map[string]*NewsClickViewMetric `protobuf:"bytes,10,rep,name=grp" json:"grp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tab  map[string]*NewsClickViewMetric `protobuf:"bytes,11,rep,name=tab" json:"tab,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Info map[string]*NewsClickViewMetric `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NewsCard) Reset()                    { *m = NewsCard{} }
func (m *NewsCard) String() string            { return proto.CompactTextString(m) }
func (*NewsCard) ProtoMessage()               {}
func (*NewsCard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewsCard) GetIg() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ig
	}
	return nil
}

func (m *NewsCard) GetLoc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *NewsCard) GetNer() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ner
	}
	return nil
}

func (m *NewsCard) GetType() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NewsCard) GetSrc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *NewsCard) GetPqs() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Pqs
	}
	return nil
}

func (m *NewsCard) GetNsfw() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Nsfw
	}
	return nil
}

func (m *NewsCard) GetLang() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Lang
	}
	return nil
}

func (m *NewsCard) GetFeed() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *NewsCard) GetGrp() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Grp
	}
	return nil
}

func (m *NewsCard) GetTab() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Tab
	}
	return nil
}

func (m *NewsCard) GetInfo() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Info
	}
	return nil
}

type NewsSwipe struct {
	Ig   map[string]*NewsClickViewMetric `protobuf:"bytes,1,rep,name=ig" json:"ig,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Loc  map[string]*NewsClickViewMetric `protobuf:"bytes,2,rep,name=loc" json:"loc,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ner  map[string]*NewsClickViewMetric `protobuf:"bytes,3,rep,name=ner" json:"ner,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type map[string]*NewsClickViewMetric `protobuf:"bytes,4,rep,name=type" json:"type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Src  map[string]*NewsClickViewMetric `protobuf:"bytes,5,rep,name=src" json:"src,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pqs  map[string]*NewsClickViewMetric `protobuf:"bytes,6,rep,name=pqs" json:"pqs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nsfw map[string]*NewsClickViewMetric `protobuf:"bytes,7,rep,name=nsfw" json:"nsfw,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Lang map[string]*NewsClickViewMetric `protobuf:"bytes,8,rep,name=lang" json:"lang,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feed map[string]*NewsClickViewMetric `protobuf:"bytes,9,rep,name=feed" json:"feed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Grp  map[string]*NewsClickViewMetric `protobuf:"bytes,10,rep,name=grp" json:"grp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tab  map[string]*NewsClickViewMetric `protobuf:"bytes,11,rep,name=tab" json:"tab,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Info map[string]*NewsClickViewMetric `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NewsSwipe) Reset()                    { *m = NewsSwipe{} }
func (m *NewsSwipe) String() string            { return proto.CompactTextString(m) }
func (*NewsSwipe) ProtoMessage()               {}
func (*NewsSwipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewsSwipe) GetIg() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ig
	}
	return nil
}

func (m *NewsSwipe) GetLoc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *NewsSwipe) GetNer() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ner
	}
	return nil
}

func (m *NewsSwipe) GetType() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NewsSwipe) GetSrc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *NewsSwipe) GetPqs() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Pqs
	}
	return nil
}

func (m *NewsSwipe) GetNsfw() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Nsfw
	}
	return nil
}

func (m *NewsSwipe) GetLang() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Lang
	}
	return nil
}

func (m *NewsSwipe) GetFeed() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *NewsSwipe) GetGrp() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Grp
	}
	return nil
}

func (m *NewsSwipe) GetTab() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Tab
	}
	return nil
}

func (m *NewsSwipe) GetInfo() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Info
	}
	return nil
}

type NewsNotif struct {
	Ig   map[string]*NewsClickViewMetric `protobuf:"bytes,1,rep,name=ig" json:"ig,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Loc  map[string]*NewsClickViewMetric `protobuf:"bytes,2,rep,name=loc" json:"loc,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ner  map[string]*NewsClickViewMetric `protobuf:"bytes,3,rep,name=ner" json:"ner,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type map[string]*NewsClickViewMetric `protobuf:"bytes,4,rep,name=type" json:"type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Src  map[string]*NewsClickViewMetric `protobuf:"bytes,5,rep,name=src" json:"src,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pqs  map[string]*NewsClickViewMetric `protobuf:"bytes,6,rep,name=pqs" json:"pqs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nsfw map[string]*NewsClickViewMetric `protobuf:"bytes,7,rep,name=nsfw" json:"nsfw,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Lang map[string]*NewsClickViewMetric `protobuf:"bytes,8,rep,name=lang" json:"lang,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feed map[string]*NewsClickViewMetric `protobuf:"bytes,9,rep,name=feed" json:"feed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Grp  map[string]*NewsClickViewMetric `protobuf:"bytes,10,rep,name=grp" json:"grp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tab  map[string]*NewsClickViewMetric `protobuf:"bytes,11,rep,name=tab" json:"tab,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Info map[string]*NewsClickViewMetric `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NewsNotif) Reset()                    { *m = NewsNotif{} }
func (m *NewsNotif) String() string            { return proto.CompactTextString(m) }
func (*NewsNotif) ProtoMessage()               {}
func (*NewsNotif) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewsNotif) GetIg() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ig
	}
	return nil
}

func (m *NewsNotif) GetLoc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *NewsNotif) GetNer() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Ner
	}
	return nil
}

func (m *NewsNotif) GetType() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NewsNotif) GetSrc() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *NewsNotif) GetPqs() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Pqs
	}
	return nil
}

func (m *NewsNotif) GetNsfw() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Nsfw
	}
	return nil
}

func (m *NewsNotif) GetLang() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Lang
	}
	return nil
}

func (m *NewsNotif) GetFeed() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *NewsNotif) GetGrp() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Grp
	}
	return nil
}

func (m *NewsNotif) GetTab() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Tab
	}
	return nil
}

func (m *NewsNotif) GetInfo() map[string]*NewsClickViewMetric {
	if m != nil {
		return m.Info
	}
	return nil
}

type NewsEngagement struct {
	Ig   map[string]*NewsEngagementMetric `protobuf:"bytes,1,rep,name=ig" json:"ig,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Loc  map[string]*NewsEngagementMetric `protobuf:"bytes,2,rep,name=loc" json:"loc,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ner  map[string]*NewsEngagementMetric `protobuf:"bytes,3,rep,name=ner" json:"ner,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type map[string]*NewsEngagementMetric `protobuf:"bytes,4,rep,name=type" json:"type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Src  map[string]*NewsEngagementMetric `protobuf:"bytes,5,rep,name=src" json:"src,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pqs  map[string]*NewsEngagementMetric `protobuf:"bytes,6,rep,name=pqs" json:"pqs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nsfw map[string]*NewsEngagementMetric `protobuf:"bytes,7,rep,name=nsfw" json:"nsfw,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Lang map[string]*NewsEngagementMetric `protobuf:"bytes,8,rep,name=lang" json:"lang,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feed map[string]*NewsEngagementMetric `protobuf:"bytes,9,rep,name=feed" json:"feed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Grp  map[string]*NewsEngagementMetric `protobuf:"bytes,10,rep,name=grp" json:"grp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tab  map[string]*NewsEngagementMetric `protobuf:"bytes,11,rep,name=tab" json:"tab,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Info map[string]*NewsEngagementMetric `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NewsEngagement) Reset()                    { *m = NewsEngagement{} }
func (m *NewsEngagement) String() string            { return proto.CompactTextString(m) }
func (*NewsEngagement) ProtoMessage()               {}
func (*NewsEngagement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NewsEngagement) GetIg() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Ig
	}
	return nil
}

func (m *NewsEngagement) GetLoc() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *NewsEngagement) GetNer() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Ner
	}
	return nil
}

func (m *NewsEngagement) GetType() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *NewsEngagement) GetSrc() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *NewsEngagement) GetPqs() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Pqs
	}
	return nil
}

func (m *NewsEngagement) GetNsfw() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Nsfw
	}
	return nil
}

func (m *NewsEngagement) GetLang() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Lang
	}
	return nil
}

func (m *NewsEngagement) GetFeed() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *NewsEngagement) GetGrp() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Grp
	}
	return nil
}

func (m *NewsEngagement) GetTab() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Tab
	}
	return nil
}

func (m *NewsEngagement) GetInfo() map[string]*NewsEngagementMetric {
	if m != nil {
		return m.Info
	}
	return nil
}

// Metrics
type NewsClickViewMetric struct {
	ClickCnt  int64   `protobuf:"varint,1,opt,name=click_cnt,json=clickCnt" json:"click_cnt,omitempty"`
	ClickDs   float32 `protobuf:"fixed32,2,opt,name=click_ds,json=clickDs" json:"click_ds,omitempty"`
	ViewCnt   int64   `protobuf:"varint,3,opt,name=view_cnt,json=viewCnt" json:"view_cnt,omitempty"`
	ViewDs    float32 `protobuf:"fixed32,4,opt,name=view_ds,json=viewDs" json:"view_ds,omitempty"`
	LastDecay int32   `protobuf:"varint,5,opt,name=last_decay,json=lastDecay" json:"last_decay,omitempty"`
}

func (m *NewsClickViewMetric) Reset()                    { *m = NewsClickViewMetric{} }
func (m *NewsClickViewMetric) String() string            { return proto.CompactTextString(m) }
func (*NewsClickViewMetric) ProtoMessage()               {}
func (*NewsClickViewMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NewsClickViewMetric) GetClickCnt() int64 {
	if m != nil {
		return m.ClickCnt
	}
	return 0
}

func (m *NewsClickViewMetric) GetClickDs() float32 {
	if m != nil {
		return m.ClickDs
	}
	return 0
}

func (m *NewsClickViewMetric) GetViewCnt() int64 {
	if m != nil {
		return m.ViewCnt
	}
	return 0
}

func (m *NewsClickViewMetric) GetViewDs() float32 {
	if m != nil {
		return m.ViewDs
	}
	return 0
}

func (m *NewsClickViewMetric) GetLastDecay() int32 {
	if m != nil {
		return m.LastDecay
	}
	return 0
}

type NewsEngagementMetric struct {
	BounceCnt   int64   `protobuf:"varint,1,opt,name=bounce_cnt,json=bounceCnt" json:"bounce_cnt,omitempty"`
	BounceDs    float32 `protobuf:"fixed32,2,opt,name=bounce_ds,json=bounceDs" json:"bounce_ds,omitempty"`
	ShallowCnt  int64   `protobuf:"varint,3,opt,name=shallow_cnt,json=shallowCnt" json:"shallow_cnt,omitempty"`
	ShallowDs   float32 `protobuf:"fixed32,4,opt,name=shallow_ds,json=shallowDs" json:"shallow_ds,omitempty"`
	DeepCnt     int64   `protobuf:"varint,5,opt,name=deep_cnt,json=deepCnt" json:"deep_cnt,omitempty"`
	DeepDs      float32 `protobuf:"fixed32,6,opt,name=deep_ds,json=deepDs" json:"deep_ds,omitempty"`
	CompleteCnt int64   `protobuf:"varint,7,opt,name=complete_cnt,json=completeCnt" json:"complete_cnt,omitempty"`
	CompleteDs  float32 `protobuf:"fixed32,8,opt,name=complete_ds,json=completeDs" json:"complete_ds,omitempty"`
	Timespent   int64   `protobuf:"varint,9,opt,name=timespent" json:"timespent,omitempty"`
	LikeCnt     int64   `protobuf:"varint,10,opt,name=like_cnt,json=likeCnt" json:"like_cnt,omitempty"`
	LikeDs      float32 `protobuf:"fixed32,11,opt,name=like_ds,json=likeDs" json:"like_ds,omitempty"`
	ShareCnt    int64   `protobuf:"varint,12,opt,name=share_cnt,json=shareCnt" json:"share_cnt,omitempty"`
	ShareDs     float32 `protobuf:"fixed32,13,opt,name=share_ds,json=shareDs" json:"share_ds,omitempty"`
	DislikeCnt  int64   `protobuf:"varint,14,opt,name=dislike_cnt,json=dislikeCnt" json:"dislike_cnt,omitempty"`
	DislikeDs   float32 `protobuf:"fixed32,15,opt,name=dislike_ds,json=dislikeDs" json:"dislike_ds,omitempty"`
	CommentCnt  int64   `protobuf:"varint,16,opt,name=comment_cnt,json=commentCnt" json:"comment_cnt,omitempty"`
	CommentDs   float32 `protobuf:"fixed32,17,opt,name=comment_ds,json=commentDs" json:"comment_ds,omitempty"`
	LastDecay   int32   `protobuf:"varint,18,opt,name=last_decay,json=lastDecay" json:"last_decay,omitempty"`
}

func (m *NewsEngagementMetric) Reset()                    { *m = NewsEngagementMetric{} }
func (m *NewsEngagementMetric) String() string            { return proto.CompactTextString(m) }
func (*NewsEngagementMetric) ProtoMessage()               {}
func (*NewsEngagementMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *NewsEngagementMetric) GetBounceCnt() int64 {
	if m != nil {
		return m.BounceCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetBounceDs() float32 {
	if m != nil {
		return m.BounceDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetShallowCnt() int64 {
	if m != nil {
		return m.ShallowCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetShallowDs() float32 {
	if m != nil {
		return m.ShallowDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetDeepCnt() int64 {
	if m != nil {
		return m.DeepCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetDeepDs() float32 {
	if m != nil {
		return m.DeepDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetCompleteCnt() int64 {
	if m != nil {
		return m.CompleteCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetCompleteDs() float32 {
	if m != nil {
		return m.CompleteDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetTimespent() int64 {
	if m != nil {
		return m.Timespent
	}
	return 0
}

func (m *NewsEngagementMetric) GetLikeCnt() int64 {
	if m != nil {
		return m.LikeCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetLikeDs() float32 {
	if m != nil {
		return m.LikeDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetShareCnt() int64 {
	if m != nil {
		return m.ShareCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetShareDs() float32 {
	if m != nil {
		return m.ShareDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetDislikeCnt() int64 {
	if m != nil {
		return m.DislikeCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetDislikeDs() float32 {
	if m != nil {
		return m.DislikeDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetCommentCnt() int64 {
	if m != nil {
		return m.CommentCnt
	}
	return 0
}

func (m *NewsEngagementMetric) GetCommentDs() float32 {
	if m != nil {
		return m.CommentDs
	}
	return 0
}

func (m *NewsEngagementMetric) GetLastDecay() int32 {
	if m != nil {
		return m.LastDecay
	}
	return 0
}

// Redis message
type RedisNewsMessage struct {
	Ig   map[string]*RedisNewsMetric `protobuf:"bytes,1,rep,name=ig" json:"ig,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Loc  map[string]*RedisNewsMetric `protobuf:"bytes,2,rep,name=loc" json:"loc,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Ner  map[string]*RedisNewsMetric `protobuf:"bytes,3,rep,name=ner" json:"ner,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Type map[string]*RedisNewsMetric `protobuf:"bytes,4,rep,name=type" json:"type,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Src  map[string]*RedisNewsMetric `protobuf:"bytes,5,rep,name=src" json:"src,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pqs  map[string]*RedisNewsMetric `protobuf:"bytes,6,rep,name=pqs" json:"pqs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Nsfw map[string]*RedisNewsMetric `protobuf:"bytes,7,rep,name=nsfw" json:"nsfw,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Lang map[string]*RedisNewsMetric `protobuf:"bytes,8,rep,name=lang" json:"lang,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Feed map[string]*RedisNewsMetric `protobuf:"bytes,9,rep,name=feed" json:"feed,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Grp  map[string]*RedisNewsMetric `protobuf:"bytes,10,rep,name=grp" json:"grp,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Tab  map[string]*RedisNewsMetric `protobuf:"bytes,11,rep,name=tab" json:"tab,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Info map[string]*RedisNewsMetric `protobuf:"bytes,12,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RedisNewsMessage) Reset()                    { *m = RedisNewsMessage{} }
func (m *RedisNewsMessage) String() string            { return proto.CompactTextString(m) }
func (*RedisNewsMessage) ProtoMessage()               {}
func (*RedisNewsMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RedisNewsMessage) GetIg() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Ig
	}
	return nil
}

func (m *RedisNewsMessage) GetLoc() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Loc
	}
	return nil
}

func (m *RedisNewsMessage) GetNer() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Ner
	}
	return nil
}

func (m *RedisNewsMessage) GetType() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RedisNewsMessage) GetSrc() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *RedisNewsMessage) GetPqs() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Pqs
	}
	return nil
}

func (m *RedisNewsMessage) GetNsfw() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Nsfw
	}
	return nil
}

func (m *RedisNewsMessage) GetLang() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Lang
	}
	return nil
}

func (m *RedisNewsMessage) GetFeed() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *RedisNewsMessage) GetGrp() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Grp
	}
	return nil
}

func (m *RedisNewsMessage) GetTab() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Tab
	}
	return nil
}

func (m *RedisNewsMessage) GetInfo() map[string]*RedisNewsMetric {
	if m != nil {
		return m.Info
	}
	return nil
}

// Redis metric
type RedisNewsMetric struct {
	CardMetric       *NewsClickViewMetric  `protobuf:"bytes,1,opt,name=card_metric,json=cardMetric" json:"card_metric,omitempty"`
	SwipeMetric      *NewsClickViewMetric  `protobuf:"bytes,2,opt,name=swipe_metric,json=swipeMetric" json:"swipe_metric,omitempty"`
	NotifMetric      *NewsClickViewMetric  `protobuf:"bytes,3,opt,name=notif_metric,json=notifMetric" json:"notif_metric,omitempty"`
	EngagementMetric *NewsEngagementMetric `protobuf:"bytes,4,opt,name=engagement_metric,json=engagementMetric" json:"engagement_metric,omitempty"`
}

func (m *RedisNewsMetric) Reset()                    { *m = RedisNewsMetric{} }
func (m *RedisNewsMetric) String() string            { return proto.CompactTextString(m) }
func (*RedisNewsMetric) ProtoMessage()               {}
func (*RedisNewsMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RedisNewsMetric) GetCardMetric() *NewsClickViewMetric {
	if m != nil {
		return m.CardMetric
	}
	return nil
}

func (m *RedisNewsMetric) GetSwipeMetric() *NewsClickViewMetric {
	if m != nil {
		return m.SwipeMetric
	}
	return nil
}

func (m *RedisNewsMetric) GetNotifMetric() *NewsClickViewMetric {
	if m != nil {
		return m.NotifMetric
	}
	return nil
}

func (m *RedisNewsMetric) GetEngagementMetric() *NewsEngagementMetric {
	if m != nil {
		return m.EngagementMetric
	}
	return nil
}

func init() {
	proto.RegisterType((*NewsCard)(nil), "userprofile.NewsCard")
	proto.RegisterType((*NewsSwipe)(nil), "userprofile.NewsSwipe")
	proto.RegisterType((*NewsNotif)(nil), "userprofile.NewsNotif")
	proto.RegisterType((*NewsEngagement)(nil), "userprofile.NewsEngagement")
	proto.RegisterType((*NewsClickViewMetric)(nil), "userprofile.NewsClickViewMetric")
	proto.RegisterType((*NewsEngagementMetric)(nil), "userprofile.NewsEngagementMetric")
	proto.RegisterType((*RedisNewsMessage)(nil), "userprofile.RedisNewsMessage")
	proto.RegisterType((*RedisNewsMetric)(nil), "userprofile.RedisNewsMetric")
}

func init() { proto.RegisterFile("news_metrics.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x9a, 0xcf, 0x4f, 0xe3, 0x46,
	0x14, 0xc7, 0x95, 0x84, 0x1f, 0xf1, 0x04, 0x58, 0xd6, 0xad, 0xd4, 0xe9, 0xee, 0x52, 0xb2, 0xf4,
	0x17, 0x97, 0xa2, 0x06, 0x68, 0x1a, 0xca, 0xa9, 0x22, 0xdb, 0x6a, 0xa5, 0x5d, 0x54, 0x05, 0xda,
	0x6e, 0x7b, 0x41, 0xc6, 0x9e, 0x64, 0x2d, 0x8c, 0x6d, 0x3c, 0x66, 0x23, 0xee, 0xfd, 0x43, 0x7a,
	0xee, 0xa1, 0x97, 0x1e, 0xfa, 0xff, 0xf4, 0x2f, 0xa9, 0xe6, 0x87, 0xc3, 0x8c, 0xc1, 0xcf, 0x78,
	0xc7, 0xc7, 0xdc, 0x32, 0xcf, 0xef, 0xfb, 0xe6, 0xcd, 0xcb, 0x98, 0xcf, 0x57, 0x28, 0xc8, 0x0e,
	0xc9, 0x94, 0x9e, 0x5d, 0x92, 0x34, 0xf1, 0x5d, 0xba, 0x13, 0x27, 0x51, 0x1a, 0xd9, 0x9d, 0x6b,
	0x4a, 0x92, 0x38, 0x89, 0xc6, 0x7e, 0x40, 0xb6, 0xfe, 0x5a, 0x45, 0xed, 0x63, 0x32, 0xa5, 0x47,
	0x4e, 0xe2, 0xd9, 0x5f, 0xa1, 0xa6, 0x3f, 0xc1, 0x8d, 0x6e, 0x6b, 0xbb, 0xb3, 0xbb, 0xb1, 0xa3,
	0xa4, 0xed, 0x64, 0x29, 0x3b, 0x2f, 0x27, 0x2f, 0xc2, 0x34, 0xb9, 0x19, 0x35, 0xfd, 0x89, 0xfd,
	0x35, 0x6a, 0x05, 0x91, 0x8b, 0x9b, 0x3c, 0xff, 0x93, 0xfb, 0xf3, 0x5f, 0x45, 0xae, 0x10, 0xb0,
	0x54, 0xa6, 0x08, 0x49, 0x82, 0x5b, 0x90, 0xe2, 0x98, 0x24, 0x52, 0x11, 0x92, 0xc4, 0xde, 0x43,
	0x0b, 0xe9, 0x4d, 0x4c, 0xf0, 0x02, 0x97, 0x6c, 0xde, 0x2f, 0x39, 0xbd, 0x89, 0x89, 0xd0, 0xf0,
	0x64, 0xb6, 0x0d, 0x4d, 0x5c, 0xbc, 0x08, 0x6d, 0x73, 0x92, 0x64, 0x8d, 0xd1, 0x84, 0x37, 0x16,
	0x5f, 0x51, 0xbc, 0x04, 0x29, 0x7e, 0xba, 0xa2, 0x52, 0x11, 0x5f, 0x51, 0xd6, 0x58, 0x48, 0xc7,
	0x53, 0xbc, 0x0c, 0x35, 0x76, 0x4c, 0xc7, 0x53, 0xd9, 0x18, 0x4b, 0x66, 0xa2, 0xc0, 0x09, 0x27,
	0xb8, 0x0d, 0x89, 0x5e, 0x39, 0xa1, 0x1c, 0x32, 0x4f, 0x66, 0xa2, 0x31, 0x21, 0x1e, 0xb6, 0x20,
	0xd1, 0x0f, 0x84, 0x78, 0x52, 0xc4, 0x92, 0xd9, 0x81, 0x26, 0x49, 0x8c, 0x11, 0x74, 0xa0, 0x1f,
	0x93, 0x58, 0x1e, 0x68, 0x92, 0xc4, 0x4c, 0x91, 0x3a, 0xe7, 0xb8, 0x03, 0x29, 0x4e, 0x9d, 0x73,
	0xa9, 0x48, 0x9d, 0x73, 0xd6, 0x98, 0x1f, 0x8e, 0x23, 0xbc, 0x02, 0x35, 0xf6, 0x32, 0x1c, 0x47,
	0xb2, 0x31, 0x96, 0xfc, 0xe4, 0x57, 0xb4, 0x2c, 0xef, 0x90, 0xbd, 0x8e, 0x5a, 0x17, 0xe4, 0x06,
	0x37, 0xba, 0x8d, 0x6d, 0x6b, 0xc4, 0x3e, 0xda, 0x7d, 0xb4, 0xf8, 0xce, 0x09, 0xae, 0x09, 0x6e,
	0x76, 0x1b, 0xdb, 0x9d, 0xdd, 0xee, 0xdd, 0x92, 0x81, 0xef, 0x5e, 0xfc, 0xe2, 0x93, 0xe9, 0x6b,
	0x7e, 0xa5, 0x47, 0x22, 0xfd, 0xbb, 0xe6, 0xa0, 0xf1, 0xe4, 0x0d, 0x6a, 0x67, 0x97, 0xad, 0xfe,
	0xca, 0xd9, 0xa5, 0xac, 0xb9, 0xf2, 0x6f, 0xc8, 0x9a, 0xdd, 0xdd, 0xfa, 0x9b, 0xce, 0xae, 0x78,
	0xfd, 0x95, 0xb3, 0x57, 0xa1, 0xfe, 0x71, 0xcc, 0xde, 0x98, 0xfa, 0x4b, 0xcf, 0xde, 0xab, 0xfa,
	0x4b, 0xcf, 0xde, 0xbe, 0xfa, 0x47, 0x9d, 0xbd, 0xa4, 0xf5, 0x57, 0xce, 0x5e, 0xe6, 0xfa, 0xc7,
	0x31, 0x7b, 0xe7, 0xeb, 0x2d, 0xbd, 0xf5, 0xef, 0x2a, 0xb2, 0x58, 0xca, 0xc9, 0xd4, 0x8f, 0x89,
	0xbd, 0xa3, 0xd0, 0xea, 0xee, 0xdf, 0x2b, 0x9e, 0xa3, 0xe1, 0xaa, 0xa7, 0xe2, 0x6a, 0xb3, 0x40,
	0xa0, 0xf3, 0xaa, 0xa7, 0xf2, 0xaa, 0x48, 0xa2, 0x03, 0x6b, 0x5f, 0x03, 0x56, 0xb7, 0x40, 0x93,
	0x27, 0x56, 0x4f, 0x25, 0x56, 0xd1, 0x46, 0x3a, 0xb2, 0x7a, 0x2a, 0xb2, 0x8a, 0x24, 0x3a, 0xb3,
	0xf6, 0x35, 0x66, 0x15, 0xf5, 0x96, 0x87, 0xd6, 0xbe, 0x06, 0xad, 0x22, 0x55, 0x9e, 0x5a, 0xfb,
	0x1a, 0xb5, 0x8a, 0x54, 0x79, 0x6c, 0xf5, 0x54, 0x6c, 0x15, 0x1d, 0x4a, 0xe7, 0x56, 0x4f, 0xe5,
	0x56, 0x91, 0x44, 0x07, 0xd7, 0xbe, 0x06, 0xae, 0xa2, 0xde, 0xe6, 0xe4, 0x9a, 0x93, 0x6b, 0x4e,
	0xae, 0x39, 0xb9, 0xea, 0x24, 0xd7, 0x71, 0x94, 0xfa, 0xe3, 0x12, 0x72, 0xf1, 0x9c, 0x2a, 0xe4,
	0x12, 0x82, 0x4a, 0xe4, 0x12, 0x92, 0x6a, 0xe4, 0x12, 0x9a, 0x8a, 0xe4, 0x12, 0xa2, 0x4a, 0xe4,
	0x12, 0x92, 0x6a, 0xe4, 0x92, 0xe7, 0xa9, 0x48, 0x2e, 0x39, 0xb8, 0x8a, 0xe4, 0x12, 0xaa, 0x8a,
	0xe4, 0x12, 0xa2, 0x4a, 0xe4, 0x92, 0xf3, 0xae, 0x44, 0x2e, 0x79, 0x77, 0xe6, 0xe4, 0x9a, 0x93,
	0x6b, 0x4e, 0xae, 0x39, 0xb9, 0x6a, 0x20, 0xd7, 0x1f, 0x6b, 0x68, 0x8d, 0xa5, 0xbc, 0x08, 0x27,
	0xce, 0x84, 0x5c, 0x92, 0x30, 0xb5, 0xf7, 0x14, 0x7c, 0x7d, 0x7a, 0xa7, 0xd6, 0x6d, 0xa2, 0xc6,
	0xb0, 0xbe, 0xca, 0xb0, 0xcf, 0x20, 0x95, 0x0e, 0xb2, 0xbe, 0x0a, 0x32, 0x50, 0xa7, 0xd3, 0xec,
	0x40, 0xa3, 0xd9, 0xe7, 0x90, 0x30, 0x8f, 0xb4, 0xbe, 0x8a, 0x34, 0x70, 0x4b, 0x9d, 0x6b, 0x7d,
	0x95, 0x6b, 0xa0, 0x4e, 0x87, 0xdb, 0x81, 0x06, 0x37, 0xb0, 0xd5, 0x3c, 0xe1, 0x0e, 0x34, 0xc2,
	0x81, 0xd2, 0x3c, 0xe6, 0x0e, 0x34, 0xcc, 0x81, 0xd2, 0x3c, 0xeb, 0xfa, 0x2a, 0xeb, 0xc0, 0x83,
	0xea, 0xc0, 0xeb, 0xab, 0xc0, 0x03, 0x75, 0x3a, 0xf5, 0x0e, 0x34, 0xea, 0x81, 0xad, 0xe6, 0xd1,
	0xf7, 0x06, 0x42, 0xdf, 0xb7, 0xfa, 0x7b, 0xf1, 0x1c, 0x28, 0x7c, 0xdf, 0x3b, 0x07, 0xb1, 0xcf,
	0xb4, 0x34, 0x00, 0x3f, 0xa3, 0xd2, 0xbf, 0xc3, 0xf4, 0x33, 0x6d, 0x1b, 0xc0, 0x9f, 0x69, 0x69,
	0x80, 0x7f, 0xa6, 0x13, 0x81, 0x00, 0x68, 0x5a, 0x1b, 0x22, 0xa0, 0x69, 0x6d, 0x08, 0x81, 0xa6,
	0xe3, 0x06, 0x18, 0x68, 0x5a, 0x1a, 0x80, 0xa0, 0xe9, 0x44, 0x20, 0x0a, 0x9a, 0xd4, 0xde, 0xfa,
	0xb3, 0x81, 0x3e, 0xb8, 0x87, 0x94, 0xf6, 0x53, 0x64, 0xb9, 0x2c, 0x74, 0xe6, 0x86, 0x29, 0xdf,
	0xac, 0x35, 0x6a, 0xf3, 0xc0, 0x51, 0x98, 0xda, 0x1f, 0x23, 0xf1, 0xf9, 0xcc, 0xa3, 0x7c, 0xd3,
	0xe6, 0x68, 0x99, 0xaf, 0x87, 0x94, 0x3d, 0x7a, 0xe7, 0x93, 0x29, 0x97, 0xb5, 0xb8, 0x6c, 0x99,
	0xad, 0x99, 0xea, 0x23, 0xc4, 0x3f, 0x32, 0xd1, 0x02, 0x17, 0x2d, 0xb1, 0xe5, 0x90, 0xda, 0x1b,
	0x08, 0x05, 0x0e, 0x4d, 0xcf, 0x3c, 0xe2, 0x3a, 0x37, 0x78, 0xb1, 0xdb, 0xd8, 0x5e, 0x1c, 0x59,
	0x2c, 0x32, 0x64, 0x81, 0xad, 0x7f, 0x16, 0xd0, 0x87, 0xf7, 0x1d, 0x83, 0xe9, 0xce, 0xa3, 0xeb,
	0xd0, 0x25, 0x4a, 0x93, 0x96, 0x88, 0xb0, 0xfd, 0x9e, 0x22, 0xb9, 0xb8, 0x6d, 0xb3, 0x2d, 0x02,
	0x43, 0x6a, 0x6f, 0xa2, 0x0e, 0x7d, 0xeb, 0x04, 0x41, 0xa4, 0xb6, 0x8a, 0x64, 0x88, 0xa9, 0x37,
	0x50, 0xb6, 0xba, 0x6d, 0xd8, 0x92, 0x11, 0x71, 0x4e, 0x8f, 0x90, 0x98, 0x8b, 0x17, 0xc5, 0x39,
	0xd9, 0x5a, 0x9e, 0x93, 0x3f, 0xf2, 0x18, 0x32, 0xf9, 0x39, 0xd9, 0x72, 0x48, 0xed, 0xe7, 0x68,
	0xc5, 0x8d, 0x2e, 0xe3, 0x80, 0xa4, 0xa2, 0xe3, 0x65, 0xae, 0xeb, 0x64, 0x31, 0xa6, 0xdd, 0x44,
	0xb3, 0x25, 0xd3, 0xb7, 0xb9, 0x1e, 0x65, 0xa1, 0x21, 0xb5, 0x9f, 0x21, 0x2b, 0xf5, 0x2f, 0x09,
	0x8d, 0x49, 0x98, 0x62, 0x4b, 0x1c, 0x79, 0x16, 0x60, 0x5d, 0x05, 0xfe, 0x85, 0xa8, 0x8e, 0x44,
	0x57, 0x6c, 0x2d, 0xbb, 0xe2, 0x8f, 0x3c, 0x8a, 0x3b, 0xa2, 0x2b, 0xb6, 0x1c, 0x52, 0x36, 0x26,
	0xfa, 0xd6, 0x49, 0x84, 0x68, 0x45, 0x7c, 0xd3, 0x3c, 0x20, 0xbf, 0x69, 0xf1, 0xd0, 0xa3, 0x78,
	0x55, 0x7c, 0xd3, 0x7c, 0x2d, 0x26, 0xe8, 0xf9, 0x74, 0xb6, 0xdd, 0x9a, 0x98, 0xa0, 0x0c, 0xc9,
	0x09, 0x66, 0x09, 0x1e, 0xc5, 0x8f, 0xc4, 0x04, 0x65, 0x44, 0xe8, 0xdd, 0xe8, 0x92, 0x7d, 0x9d,
	0x5c, 0xbf, 0x2e, 0xf4, 0x32, 0x24, 0xf5, 0x59, 0x82, 0x47, 0xf1, 0x63, 0xa1, 0x97, 0x91, 0x3b,
	0xb7, 0xc6, 0xce, 0xdf, 0x9a, 0xff, 0x56, 0xd1, 0xfa, 0x88, 0x78, 0x3e, 0x65, 0x57, 0xe7, 0x35,
	0xa1, 0xd4, 0x99, 0x10, 0xfb, 0x1b, 0xc5, 0xe1, 0xe9, 0xb8, 0xcd, 0xa7, 0x6a, 0x1e, 0x6f, 0xa0,
	0x7a, 0xbc, 0x2f, 0x60, 0x9d, 0xee, 0xf2, 0x06, 0xaa, 0xcb, 0x2b, 0x51, 0xea, 0x3e, 0xef, 0x50,
	0xf3, 0x79, 0x5f, 0xc2, 0xd2, 0xbc, 0xd3, 0x1b, 0xa8, 0x4e, 0xaf, 0x64, 0x5b, 0xdd, 0xeb, 0x0d,
	0x54, 0xaf, 0x57, 0xa2, 0xd4, 0xdd, 0xde, 0xa1, 0xe6, 0xf6, 0x4a, 0x1a, 0xce, 0xfb, 0xbd, 0x43,
	0xcd, 0xef, 0x95, 0x88, 0xf3, 0x8e, 0xef, 0x50, 0x73, 0x7c, 0x25, 0xe2, 0xbc, 0xe7, 0x1b, 0xa8,
	0x9e, 0xaf, 0xe4, 0xc0, 0xba, 0xeb, 0x1b, 0xa8, 0xae, 0xaf, 0x44, 0xa9, 0xfb, 0xbe, 0x43, 0xcd,
	0xf7, 0x95, 0x34, 0x9c, 0x77, 0x7e, 0x27, 0x90, 0xf3, 0xdb, 0xd5, 0x59, 0xf0, 0xac, 0xa8, 0x74,
	0x1e, 0x31, 0xa7, 0xa0, 0xe9, 0x33, 0xa8, 0x0a, 0xf8, 0xbd, 0xf7, 0xad, 0xfa, 0x33, 0x6c, 0xf5,
	0x0c, 0x9a, 0x05, 0x5c, 0x9e, 0x41, 0x55, 0xc0, 0xe0, 0x19, 0x8c, 0x00, 0xf2, 0x76, 0x06, 0x65,
	0x21, 0x5b, 0x67, 0x50, 0x16, 0x72, 0x74, 0x06, 0xa3, 0x05, 0xcc, 0x9c, 0x41, 0x55, 0xc0, 0xc7,
	0x19, 0x8c, 0x00, 0xb2, 0x70, 0xef, 0x59, 0x76, 0xeb, 0xef, 0x26, 0x7a, 0x94, 0x7b, 0x6c, 0x7f,
	0x8f, 0x3a, 0xae, 0x93, 0x78, 0xf2, 0xc7, 0x51, 0x7c, 0x97, 0x87, 0xfc, 0x6b, 0x04, 0x31, 0x91,
	0x2c, 0x71, 0x84, 0x56, 0xe8, 0xd4, 0x8f, 0x49, 0x56, 0xe3, 0xa1, 0xff, 0x5e, 0xe9, 0x70, 0xd5,
	0x6d, 0x91, 0x30, 0x4a, 0xfd, 0x71, 0x56, 0xa4, 0xf5, 0xd0, 0x22, 0x5c, 0x25, 0x8b, 0x1c, 0xa3,
	0xc7, 0x64, 0x66, 0xfb, 0xb2, 0x4a, 0x0b, 0x0f, 0xf5, 0xb9, 0xeb, 0x24, 0x17, 0x39, 0x5f, 0xe2,
	0x3f, 0x15, 0xdb, 0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x2b, 0x22, 0x0b, 0x40, 0x26, 0x00,
	0x00,
}