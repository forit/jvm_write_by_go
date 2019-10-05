package rtda

import "jvmgo/ch07/rtda/heap"

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

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPc(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame{
	return self.stack.top()
}

func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}
func (self *Thread) IsEmptyStack() bool {
	return self.stack.isEmpty()
}
