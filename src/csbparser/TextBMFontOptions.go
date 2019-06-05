// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TextBMFontOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsTextBMFontOptions(buf []byte, offset flatbuffers.UOffsetT) *TextBMFontOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TextBMFontOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TextBMFontOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TextBMFontOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TextBMFontOptions) WidgetOptions(obj *WidgetOptions) *WidgetOptions {
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

func (rcv *TextBMFontOptions) FileNameData(obj *ResourceData) *ResourceData {
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

func (rcv *TextBMFontOptions) Text() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func TextBMFontOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TextBMFontOptionsAddWidgetOptions(builder *flatbuffers.Builder, widgetOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(widgetOptions), 0)
}
func TextBMFontOptionsAddFileNameData(builder *flatbuffers.Builder, fileNameData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fileNameData), 0)
}
func TextBMFontOptionsAddText(builder *flatbuffers.Builder, text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(text), 0)
}
func TextBMFontOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
