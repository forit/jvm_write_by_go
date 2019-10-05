package rtda

import (
	"jvmgo/ch07/rtda/heap"
	"math"
)

type LocalVars []Slot  //length:attr_code_info maxLocals

func NewLocalVars(maxLocals uint16) LocalVars {
	return make([]Slot, maxLocals)
}

func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self.SetInt(index, int32(bits))
}
func (self LocalVars) GetFloat(index uint) float32 {
	bits := self.GetInt(index)
	return math.Float32frombits(uint32(bits))
}

// long consumes two slots
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}
func (self LocalVars) GetDouble(index uint) float64 {
	bits := self.GetLong(index)
	return math.Float64frombits(uint64(bits))
}

func (self LocalVars) SetRef(index uint, ref *heap.Object) {
	self[index].ref = ref
}
func (self LocalVars) GetRef(index uint) *heap.Object {
	return self[index].ref
}
func (self LocalVars) SetSlot(index uint, slot Slot) {
	self[index] = slot
}



