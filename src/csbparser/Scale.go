// automatically generated by the FlatBuffers compiler, do not modify

package csbparser

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Scale struct {
	_tab flatbuffers.Struct
}

func (rcv *Scale) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Scale) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Scale) ScaleX() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Scale) MutateScaleX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Scale) ScaleY() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Scale) MutateScaleY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateScale(builder *flatbuffers.Builder, scaleX float32, scaleY float32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependFloat32(scaleY)
	builder.PrependFloat32(scaleX)
	return builder.Offset()
}
