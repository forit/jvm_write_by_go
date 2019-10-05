package rtda

import "jvmgo/ch06/rtda/heap"

type Frame struct {
	lower	*Frame
	localVars LocalVars 		//jvm stack frame local vars table
	operandStack *OperandStack 	//jvm stack frame operand Stack
	thread *Thread
	Pc int
	method *heap.Method
}

func NewFrame(thread *Thread, method *heap.Method, maxLocals uint16, maxStack uint16) *Frame {
	return &Frame{
		localVars: NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
		thread:thread,
		method:method,
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) SetNextPC(pc int) {
	self.Pc = pc
}

func (self *Frame) NextPC() int {
	return self.Pc
}

func (self *Frame) Method() *heap.Method {
	return self.method
}