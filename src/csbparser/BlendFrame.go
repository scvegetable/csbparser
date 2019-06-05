// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BlendFrame struct {
	_tab flatbuffers.Table
}

func GetRootAsBlendFrame(buf []byte, offset flatbuffers.UOffsetT) *BlendFrame {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BlendFrame{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BlendFrame) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BlendFrame) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BlendFrame) FrameIndex() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BlendFrame) MutateFrameIndex(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *BlendFrame) Tween() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *BlendFrame) MutateTween(n bool) bool {
	return rcv._tab.MutateBoolSlot(6, n)
}

func (rcv *BlendFrame) BlendFunc(obj *BlendFunc) *BlendFunc {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(BlendFunc)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *BlendFrame) EasingData(obj *EasingData) *EasingData {
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

func BlendFrameStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func BlendFrameAddFrameIndex(builder *flatbuffers.Builder, frameIndex int32) {
	builder.PrependInt32Slot(0, frameIndex, 0)
}
func BlendFrameAddTween(builder *flatbuffers.Builder, tween bool) {
	builder.PrependBoolSlot(1, tween, true)
}
func BlendFrameAddBlendFunc(builder *flatbuffers.Builder, blendFunc flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(blendFunc), 0)
}
func BlendFrameAddEasingData(builder *flatbuffers.Builder, easingData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(easingData), 0)
}
func BlendFrameEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}