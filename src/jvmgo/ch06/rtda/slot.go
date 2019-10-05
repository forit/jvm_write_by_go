package rtda

import "jvmgo/ch06/rtda/heap"

//slot for LocalVars or OperandStack
type Slot struct {
	num int32
	ref *heap.Object
}

