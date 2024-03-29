// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TimeLine struct {
	_tab flatbuffers.Table
}

func GetRootAsTimeLine(buf []byte, offset flatbuffers.UOffsetT) *TimeLine {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TimeLine{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TimeLine) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TimeLine) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TimeLine) Property() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *TimeLine) ActionTag() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TimeLine) MutateActionTag(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *TimeLine) Frames(obj *Frame, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TimeLine) FramesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TimeLineStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TimeLineAddProperty(builder *flatbuffers.Builder, property flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(property), 0)
}
func TimeLineAddActionTag(builder *flatbuffers.Builder, actionTag int32) {
	builder.PrependInt32Slot(1, actionTag, 0)
}
func TimeLineAddFrames(builder *flatbuffers.Builder, frames flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(frames), 0)
}
func TimeLineStartFramesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TimeLineEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
