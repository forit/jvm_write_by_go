package rtda

import "jvmgo/ch07/rtda/heap"

//slot for LocalVars or OperandStack
type Slot struct {
	num int32
	ref *heap.Object
}

