// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ComponentOptions struct {
	_tab flatbuffers.Table
}

func GetRootAsComponentOptions(buf []byte, offset flatbuffers.UOffsetT) *ComponentOptions {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ComponentOptions{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ComponentOptions) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ComponentOptions) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ComponentOptions) NodeOptions(obj *WidgetOptions) *WidgetOptions {
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

func (rcv *ComponentOptions) ComType() string {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.String(o + rcv._tab.Pos)
	}
	return ""
}

func (rcv *ComponentOptions) ComAudioOptions(obj *ComAudioOptions) *ComAudioOptions {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(ComAudioOptions)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func ComponentOptionsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func ComponentOptionsAddNodeOptions(builder *flatbuffers.Builder, nodeOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(nodeOptions), 0)
}
func ComponentOptionsAddComType(builder *flatbuffers.Builder, comType flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(comType), 0)
}
func ComponentOptionsAddComAudioOptions(builder *flatbuffers.Builder, comAudioOptions flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(comAudioOptions), 0)
}
func ComponentOptionsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
