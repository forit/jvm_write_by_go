package main

import (
	"fmt"
	"jvmgo/ch05/classfile"
	"jvmgo/ch05/instructions"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

func interpreter(memberInfo *classfile.MemberInfo) {
	codeAttribute := memberInfo.CodeAttribute()
	maxLocals := codeAttribute.MaxLocals()
	maxStack := codeAttribute.MaxStack()
	byteCodes := codeAttribute.ByteCodes()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, byteCodes)
}
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, byteCodes []byte) {
	frame := thread.PopFrame()

	reader := &base.BytecodeReader{}
	for {

		pc := frame.NextPC()
		thread.SetPc(pc)

		reader.Reset(byteCodes, pc)
		// read opCode
		opCode := reader.ReadUint8()
		inst := instructions.NewInstruction(opCode)
		inst.FetchOperands(reader)

		//calculate next pc
		frame.SetNextPC(reader.PC())

		//execute code
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)  // control opcode can calculate next pc
	}

}
