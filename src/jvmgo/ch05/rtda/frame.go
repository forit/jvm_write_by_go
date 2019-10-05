package rtda

type Frame struct {
	lower	*Frame
	localVars LocalVars 		//jvm stack frame local vars table
	operandStack *OperandStack 	//jvm stack frame operand Stack
	thread *Thread
	Pc int
}

func NewFrame(thread *Thread , maxLocals uint16, maxStack uint16) *Frame {
	return &Frame{
		localVars: NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
		thread:thread,
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