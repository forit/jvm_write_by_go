package main

import (
	"fmt"
	"jvmgo/ch07/instructions"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
	"jvmgo/ch07/rtda/heap"
)

func interpreter(method *heap.Method, verboseInstFlag bool) {

	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, verboseInstFlag)
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, verboseInstFlag bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()

		pc := frame.NextPC()
		thread.SetPc(pc)

		reader.Reset(frame.Method().Code(), pc)
		// read opCode
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)

		//calculate next pc
		frame.SetNextPC(reader.PC())

		if verboseInstFlag {
			fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		}

		//execute code
		inst.Execute(frame)  // control opcode can calculate next pc

		if thread.IsEmptyStack() {
			break
		}
	}

}
