// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type AnimationInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsAnimationInfo(buf []byte, offset flatbuffers.UOffsetT) *AnimationInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &AnimationInfo{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *AnimationInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *AnimationInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *AnimationInfo) Name() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *AnimationInfo) StartIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AnimationInfo) MutateStartIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *AnimationInfo) EndIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *AnimationInfo) MutateEndIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func AnimationInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AnimationInfoAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func AnimationInfoAddStartIndex(builder *flatbuffers.Builder, startIndex int32) {
	builder.PrependInt32Slot(1, startIndex, 0)
}
func AnimationInfoAddEndIndex(builder *flatbuffers.Builder, endIndex int32) {
	builder.PrependInt32Slot(2, endIndex, 0)
}
func AnimationInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
