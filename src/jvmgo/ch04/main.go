package main

import (
	"fmt"
	"jvmgo/ch04/classfile"
	"jvmgo/ch04/classpath"
	"jvmgo/ch04/rtda"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1 fake msg")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}
func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("calsspath:%s class:%s args:%s. jvm start!!!!!\n", cp, cmd.class, cmd.args)

	classname := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("\ncould not find or load main class %s\n", cmd.class)
		panic(err)
		return
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		fmt.Printf("\n%s\n", err)
		panic(err)
		return
	}

	printClassInfo(cf)

	testJvmStackFrame()
}

func printClassInfo(cf *classfile.ClassFile) {
	//fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	//fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	//fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	//fmt.Printf("this class: %v\n", cf.ClassName())
	//fmt.Printf("super class: %v\n", cf.SuperClassName())
	//fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	//fmt.Printf("fields count: %v\n", len(cf.Fields()))
	//for _, f := range cf.Fields() {
	//	fmt.Printf("  %s\n", f.Name())
	//}
	//fmt.Printf("methods count: %v\n", len(cf.Methods()))
	//for _, m := range cf.Methods() {
	//	fmt.Printf("  %s\n", m.Name())
	//}
}

func testJvmStackFrame() {
	frame := rtda.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 2997924580)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.1415926)
	vars.SetDouble(7, 2.71828182845)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
}

func testOperandStack(ops *rtda.OperandStack) {
	ops.PushInt(100)
	ops.PushInt(-100)
	ops.PushLong(2997924580)
	ops.PushLong(-2997924580)
	ops.PushFloat(3.1415926)
	ops.PushDouble(2.71828182845)
	ops.PushRef(nil)
	println(ops.PopRef())
	println(ops.PopDouble())
	println(ops.PopFloat())
	println(ops.PopLong())
	println(ops.PopLong())
	println(ops.PopInt())
	println(ops.PopInt())
}