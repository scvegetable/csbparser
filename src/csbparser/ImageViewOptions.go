// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ImageViewOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsImageViewOptions(buf []byte, offset flatbuffers.UOffsetT) *ImageViewOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ImageViewOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ImageViewOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ImageViewOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ImageViewOptions) WidgetOptions(obj *WidgetOptions) *WidgetOptions {
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

func (rcv *ImageViewOptions) FileNameData(obj *ResourceData) *ResourceData {
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

func (rcv *ImageViewOptions) CapInsets(obj *CapInsets) *CapInsets {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(CapInsets)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ImageViewOptions) Scale9Size(obj *FlatSize) *FlatSize {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(FlatSize)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ImageViewOptions) Scale9Enabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ImageViewOptions) MutateScale9Enabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func ImageViewOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func ImageViewOptionsAddWidgetOptions(builder *flatbuffers.Builder, widgetOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(widgetOptions), 0)
}
func ImageViewOptionsAddFileNameData(builder *flatbuffers.Builder, fileNameData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fileNameData), 0)
}
func ImageViewOptionsAddCapInsets(builder *flatbuffers.Builder, capInsets flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(capInsets), 0)
}
func ImageViewOptionsAddScale9Size(builder *flatbuffers.Builder, scale9Size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(scale9Size), 0)
}
func ImageViewOptionsAddScale9Enabled(builder *flatbuffers.Builder, scale9Enabled bool) {
	builder.PrependBoolSlot(4, scale9Enabled, false)
}
func ImageViewOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
