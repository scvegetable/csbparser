// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LoadingBarOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsLoadingBarOptions(buf []byte, offset flatbuffers.UOffsetT) *LoadingBarOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LoadingBarOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LoadingBarOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LoadingBarOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LoadingBarOptions) WidgetOptions(obj *WidgetOptions) *WidgetOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(WidgetOptions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *LoadingBarOptions) TextureData(obj *ResourceData) *ResourceData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ResourceData)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *LoadingBarOptions) Percent() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 80
}

func (rcv *LoadingBarOptions) MutatePercent(n int32) bool {
	return rcv._tab.MutateInt32Slot(8, n)
}

func (rcv *LoadingBarOptions) Direction() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *LoadingBarOptions) MutateDirection(n int32) bool {
	return rcv._tab.MutateInt32Slot(10, n)
}

func LoadingBarOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func LoadingBarOptionsAddWidgetOptions(builder *flatbuffers.Builder, widgetOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(widgetOptions), 0)
}
func LoadingBarOptionsAddTextureData(builder *flatbuffers.Builder, textureData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(textureData), 0)
}
func LoadingBarOptionsAddPercent(builder *flatbuffers.Builder, percent int32) {
	builder.PrependInt32Slot(2, percent, 80)
}
func LoadingBarOptionsAddDirection(builder *flatbuffers.Builder, direction int32) {
	builder.PrependInt32Slot(3, direction, 0)
}
func LoadingBarOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}