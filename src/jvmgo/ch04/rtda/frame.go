package rtda

type Frame struct {
	lower	*Frame
	localVars LocalVars 		//jvm stack frame local vars table
	operandStack *OperandStack 	//jvm stack frame operand Stack
}

func NewFrame(maxLocals uint16, maxStack uint16) *Frame {
	return &Frame{
		localVars: NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}