package rtda

import (
	"jvmgo/ch07/rtda/heap"
	"math"
)

type OperandStack struct {
	top uint16			//top < attr_code_info maxStack
	slots []Slot
}

func NewOperandStack(maxStack uint16) *OperandStack {
	return &OperandStack{
		slots: make([]Slot, maxStack),
	}
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.top].num = val
	self.top++
}
func (self *OperandStack) PopInt() int32 {
	self.top--
	return self.slots[self.top].num
}

func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.top].num = int32(bits)
	self.top++
}
func (self *OperandStack) PopFloat() float32 {
	self.top--
	bits := uint32(self.slots[self.top].num)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.top].num = int32(val)
	self.slots[self.top+1].num = int32(val >> 32)
	self.top += 2
}
func (self *OperandStack) PopLong() int64 {
	self.top -= 2
	low := uint32(self.slots[self.top].num)
	high := uint32(self.slots[self.top+1].num)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(ref *heap.Object) {
	self.slots[self.top].ref = ref
	self.top++
}
func (self *OperandStack) PopRef() *heap.Object {
	self.top--
	ref := self.slots[self.top].ref
	self.slots[self.top].ref = nil
	return ref
}

func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.top] = slot
	self.top++
}
func (self *OperandStack) PopSlot() Slot {
	self.top--
	return self.slots[self.top]
}

func (self *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return self.slots[uint(self.top)-1-n].ref
}





