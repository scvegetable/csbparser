// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScrollViewOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsScrollViewOptions(buf []byte, offset flatbuffers.UOffsetT) *ScrollViewOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScrollViewOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ScrollViewOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScrollViewOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScrollViewOptions) WidgetOptions(obj *WidgetOptions) *WidgetOptions {
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

func (rcv *ScrollViewOptions) BackGroundImageData(obj *ResourceData) *ResourceData {
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

func (rcv *ScrollViewOptions) ClipEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ScrollViewOptions) MutateClipEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(8, n)
}

func (rcv *ScrollViewOptions) BgColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ScrollViewOptions) BgStartColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ScrollViewOptions) BgEndColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Color)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ScrollViewOptions) ColorType() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScrollViewOptions) MutateColorType(n int32) bool {
	return rcv._tab.MutateInt32Slot(16, n)
}

func (rcv *ScrollViewOptions) BgColorOpacity() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 255
}

func (rcv *ScrollViewOptions) MutateBgColorOpacity(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

func (rcv *ScrollViewOptions) ColorVector(obj *ColorVector) *ColorVector {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(ColorVector)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *ScrollViewOptions) CapInsets(obj *CapInsets) *CapInsets {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
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

func (rcv *ScrollViewOptions) Scale9Size(obj *FlatSize) *FlatSize {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
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

func (rcv *ScrollViewOptions) BackGroundScale9Enabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ScrollViewOptions) MutateBackGroundScale9Enabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func (rcv *ScrollViewOptions) InnerSize(obj *FlatSize) *FlatSize {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
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

func (rcv *ScrollViewOptions) Direction() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScrollViewOptions) MutateDirection(n int32) bool {
	return rcv._tab.MutateInt32Slot(30, n)
}

func (rcv *ScrollViewOptions) BounceEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ScrollViewOptions) MutateBounceEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(32, n)
}

func (rcv *ScrollViewOptions) ScrollbarEnabeld() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *ScrollViewOptions) MutateScrollbarEnabeld(n bool) bool {
	return rcv._tab.MutateBoolSlot(34, n)
}

func (rcv *ScrollViewOptions) ScrollbarAutoHide() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *ScrollViewOptions) MutateScrollbarAutoHide(n bool) bool {
	return rcv._tab.MutateBoolSlot(36, n)
}

func (rcv *ScrollViewOptions) ScrollbarAutoHideTime() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.2
}

func (rcv *ScrollViewOptions) MutateScrollbarAutoHideTime(n float32) bool {
	return rcv._tab.MutateFloat32Slot(38, n)
}

func ScrollViewOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(18)
}
func ScrollViewOptionsAddWidgetOptions(builder *flatbuffers.Builder, widgetOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(widgetOptions), 0)
}
func ScrollViewOptionsAddBackGroundImageData(builder *flatbuffers.Builder, backGroundImageData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(backGroundImageData), 0)
}
func ScrollViewOptionsAddClipEnabled(builder *flatbuffers.Builder, clipEnabled bool) {
	builder.PrependBoolSlot(2, clipEnabled, false)
}
func ScrollViewOptionsAddBgColor(builder *flatbuffers.Builder, bgColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(3, flatbuffers.UOffsetT(bgColor), 0)
}
func ScrollViewOptionsAddBgStartColor(builder *flatbuffers.Builder, bgStartColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(4, flatbuffers.UOffsetT(bgStartColor), 0)
}
func ScrollViewOptionsAddBgEndColor(builder *flatbuffers.Builder, bgEndColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(5, flatbuffers.UOffsetT(bgEndColor), 0)
}
func ScrollViewOptionsAddColorType(builder *flatbuffers.Builder, colorType int32) {
	builder.PrependInt32Slot(6, colorType, 0)
}
func ScrollViewOptionsAddBgColorOpacity(builder *flatbuffers.Builder, bgColorOpacity byte) {
	builder.PrependByteSlot(7, bgColorOpacity, 255)
}
func ScrollViewOptionsAddColorVector(builder *flatbuffers.Builder, colorVector flatbuffers.UOffsetT) {
	builder.PrependStructSlot(8, flatbuffers.UOffsetT(colorVector), 0)
}
func ScrollViewOptionsAddCapInsets(builder *flatbuffers.Builder, capInsets flatbuffers.UOffsetT) {
	builder.PrependStructSlot(9, flatbuffers.UOffsetT(capInsets), 0)
}
func ScrollViewOptionsAddScale9Size(builder *flatbuffers.Builder, scale9Size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(10, flatbuffers.UOffsetT(scale9Size), 0)
}
func ScrollViewOptionsAddBackGroundScale9Enabled(builder *flatbuffers.Builder, backGroundScale9Enabled bool) {
	builder.PrependBoolSlot(11, backGroundScale9Enabled, false)
}
func ScrollViewOptionsAddInnerSize(builder *flatbuffers.Builder, innerSize flatbuffers.UOffsetT) {
	builder.PrependStructSlot(12, flatbuffers.UOffsetT(innerSize), 0)
}
func ScrollViewOptionsAddDirection(builder *flatbuffers.Builder, direction int32) {
	builder.PrependInt32Slot(13, direction, 0)
}
func ScrollViewOptionsAddBounceEnabled(builder *flatbuffers.Builder, bounceEnabled bool) {
	builder.PrependBoolSlot(14, bounceEnabled, false)
}
func ScrollViewOptionsAddScrollbarEnabeld(builder *flatbuffers.Builder, scrollbarEnabeld bool) {
	builder.PrependBoolSlot(15, scrollbarEnabeld, true)
}
func ScrollViewOptionsAddScrollbarAutoHide(builder *flatbuffers.Builder, scrollbarAutoHide bool) {
	builder.PrependBoolSlot(16, scrollbarAutoHide, true)
}
func ScrollViewOptionsAddScrollbarAutoHideTime(builder *flatbuffers.Builder, scrollbarAutoHideTime float32) {
	builder.PrependFloat32Slot(17, scrollbarAutoHideTime, 0.2)
}
func ScrollViewOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}