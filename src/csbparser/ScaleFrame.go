// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScaleFrame struct {
	_tab flatbuffers.Table
}

func GetRootAsScaleFrame(buf []byte, offset flatbuffers.UOffsetT) *ScaleFrame {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScaleFrame{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ScaleFrame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScaleFrame) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScaleFrame) FrameIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScaleFrame) MutateFrameIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *ScaleFrame) Tween() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *ScaleFrame) MutateTween(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *ScaleFrame) Scale(obj *Scale) *Scale {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Scale)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ScaleFrame) EasingData(obj *EasingData) *EasingData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(EasingData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ScaleFrameStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ScaleFrameAddFrameIndex(builder *flatbuffers.Builder, frameIndex int32) {
	builder.PrependInt32Slot(0, frameIndex, 0)
}
func ScaleFrameAddTween(builder *flatbuffers.Builder, tween bool) {
	builder.PrependBoolSlot(1, tween, true)
}
func ScaleFrameAddScale(builder *flatbuffers.Builder, scale flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(scale), 0)
}
func ScaleFrameAddEasingData(builder *flatbuffers.Builder, easingData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(easingData), 0)
}
func ScaleFrameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}