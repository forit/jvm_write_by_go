package rtda

type Stack struct {
	maxSize uint   // max size of jvm stack frame
	size    uint   // current size of jvm stack frame
	_top    *Frame // jvm stack frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{maxSize:maxSize}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("jvm stack over flow")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack empty")
	}

	self.size--

	frame := self._top
	self._top = self._top.lower

	frame.lower = nil
	return frame
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack empty")
	}
	return self._top
}
func (self *Stack) isEmpty() bool {
	return self._top == nil
}





