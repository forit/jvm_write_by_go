package rtda

// java thread runtime data structure
type Thread struct {
	pc	int			//jvm program counter. int :32bit machine 32,64bit machine 64
	stack *Stack	//jvm stack
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}

func (self *Thread) getPc() int {
	return self.pc
}

func (self *Thread) setPc(pc int) {
	self.pc = pc
}

func (self *Thread) pushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) popFrame(frame *Frame) *Frame {
	return self.stack.pop()
}

func (self *Thread) currentFrame() *Frame{
	return self.stack.top()
}

