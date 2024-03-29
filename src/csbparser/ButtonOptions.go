// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ButtonOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsButtonOptions(buf []byte, offset flatbuffers.UOffsetT) *ButtonOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ButtonOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ButtonOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ButtonOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ButtonOptions) WidgetOptions(obj *WidgetOptions) *WidgetOptions {
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

func (rcv *ButtonOptions) NormalData(obj *ResourceData) *ResourceData {
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

func (rcv *ButtonOptions) PressedData(obj *ResourceData) *ResourceData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
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

func (rcv *ButtonOptions) DisabledData(obj *ResourceData) *ResourceData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
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

func (rcv *ButtonOptions) FontResource(obj *ResourceData) *ResourceData {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
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

func (rcv *ButtonOptions) Text() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *ButtonOptions) FontName() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *ButtonOptions) FontSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ButtonOptions) MutateFontSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(18, n)
}

func (rcv *ButtonOptions) TextColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
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

func (rcv *ButtonOptions) CapInsets(obj *CapInsets) *CapInsets {
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

func (rcv *ButtonOptions) Scale9Size(obj *FlatSize) *FlatSize {
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

func (rcv *ButtonOptions) Scale9Enabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ButtonOptions) MutateScale9Enabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(26, n)
}

func (rcv *ButtonOptions) Displaystate() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return true
}

func (rcv *ButtonOptions) MutateDisplaystate(n bool) bool {
	return rcv._tab.MutateBoolSlot(28, n)
}

func (rcv *ButtonOptions) OutlineEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ButtonOptions) MutateOutlineEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(30, n)
}

func (rcv *ButtonOptions) OutlineColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
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

func (rcv *ButtonOptions) OutlineSize() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 1
}

func (rcv *ButtonOptions) MutateOutlineSize(n int32) bool {
	return rcv._tab.MutateInt32Slot(34, n)
}

func (rcv *ButtonOptions) ShadowEnabled() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *ButtonOptions) MutateShadowEnabled(n bool) bool {
	return rcv._tab.MutateBoolSlot(36, n)
}

func (rcv *ButtonOptions) ShadowColor(obj *Color) *Color {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
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

func (rcv *ButtonOptions) ShadowOffsetX() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 2.0
}

func (rcv *ButtonOptions) MutateShadowOffsetX(n float32) bool {
	return rcv._tab.MutateFloat32Slot(40, n)
}

func (rcv *ButtonOptions) ShadowOffsetY() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return -2.0
}

func (rcv *ButtonOptions) MutateShadowOffsetY(n float32) bool {
	return rcv._tab.MutateFloat32Slot(42, n)
}

func (rcv *ButtonOptions) ShadowBlurRadius() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ButtonOptions) MutateShadowBlurRadius(n int32) bool {
	return rcv._tab.MutateInt32Slot(44, n)
}

func ButtonOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(21)
}
func ButtonOptionsAddWidgetOptions(builder *flatbuffers.Builder, widgetOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(widgetOptions), 0)
}
func ButtonOptionsAddNormalData(builder *flatbuffers.Builder, normalData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(normalData), 0)
}
func ButtonOptionsAddPressedData(builder *flatbuffers.Builder, pressedData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(pressedData), 0)
}
func ButtonOptionsAddDisabledData(builder *flatbuffers.Builder, disabledData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(disabledData), 0)
}
func ButtonOptionsAddFontResource(builder *flatbuffers.Builder, fontResource flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(fontResource), 0)
}
func ButtonOptionsAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(text), 0)
}
func ButtonOptionsAddFontName(builder *flatbuffers.Builder, fontName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(fontName), 0)
}
func ButtonOptionsAddFontSize(builder *flatbuffers.Builder, fontSize int32) {
	builder.PrependInt32Slot(7, fontSize, 0)
}
func ButtonOptionsAddTextColor(builder *flatbuffers.Builder, textColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(8, flatbuffers.UOffsetT(textColor), 0)
}
func ButtonOptionsAddCapInsets(builder *flatbuffers.Builder, capInsets flatbuffers.UOffsetT) {
	builder.PrependStructSlot(9, flatbuffers.UOffsetT(capInsets), 0)
}
func ButtonOptionsAddScale9Size(builder *flatbuffers.Builder, scale9Size flatbuffers.UOffsetT) {
	builder.PrependStructSlot(10, flatbuffers.UOffsetT(scale9Size), 0)
}
func ButtonOptionsAddScale9Enabled(builder *flatbuffers.Builder, scale9Enabled bool) {
	builder.PrependBoolSlot(11, scale9Enabled, false)
}
func ButtonOptionsAddDisplaystate(builder *flatbuffers.Builder, displaystate bool) {
	builder.PrependBoolSlot(12, displaystate, true)
}
func ButtonOptionsAddOutlineEnabled(builder *flatbuffers.Builder, outlineEnabled bool) {
	builder.PrependBoolSlot(13, outlineEnabled, false)
}
func ButtonOptionsAddOutlineColor(builder *flatbuffers.Builder, outlineColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(14, flatbuffers.UOffsetT(outlineColor), 0)
}
func ButtonOptionsAddOutlineSize(builder *flatbuffers.Builder, outlineSize int32) {
	builder.PrependInt32Slot(15, outlineSize, 1)
}
func ButtonOptionsAddShadowEnabled(builder *flatbuffers.Builder, shadowEnabled bool) {
	builder.PrependBoolSlot(16, shadowEnabled, false)
}
func ButtonOptionsAddShadowColor(builder *flatbuffers.Builder, shadowColor flatbuffers.UOffsetT) {
	builder.PrependStructSlot(17, flatbuffers.UOffsetT(shadowColor), 0)
}
func ButtonOptionsAddShadowOffsetX(builder *flatbuffers.Builder, shadowOffsetX float32) {
	builder.PrependFloat32Slot(18, shadowOffsetX, 2.0)
}
func ButtonOptionsAddShadowOffsetY(builder *flatbuffers.Builder, shadowOffsetY float32) {
	builder.PrependFloat32Slot(19, shadowOffsetY, -2.0)
}
func ButtonOptionsAddShadowBlurRadius(builder *flatbuffers.Builder, shadowBlurRadius int32) {
	builder.PrependInt32Slot(20, shadowBlurRadius, 0)
}
func ButtonOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
