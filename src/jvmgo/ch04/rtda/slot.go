package rtda

//slot for LocalVars or OperandStack
type Slot struct {
	num int32
	ref *Object
}

