package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.CpIntegerInfo:			//数值常量
			intInfo := cpInfo.(*classfile.CpIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.CpFloatInfo:
			floatInfo := cpInfo.(*classfile.CpFloatInfo)
			consts[i] = floatInfo.Value()
		case *classfile.CpLongInfo:
			longInfo := cpInfo.(*classfile.CpLongInfo)
			consts[i] = longInfo.Value()
			i++	 //Long double占用两个表项
		case *classfile.CpDoubleInfo:
			doubleInfo := cpInfo.(*classfile.CpDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.CpStringInfo:			//字符串
			stringInfo := cpInfo.(*classfile.CpStringInfo)
			consts[i] = stringInfo.String()
		case *classfile.CpClassInfo:		//符号引用
			classInfo := cpInfo.(*classfile.CpClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.CpFieldrefInfo:
			fieldrefInfo := cpInfo.(*classfile.CpFieldrefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.CpMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.CpMethodrefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.CpInterfaceMethodrefInfo:
			methodrefInfo := cpInfo.(*classfile.CpInterfaceMethodrefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, methodrefInfo)
		default:	//动态调用相关的三个常量池表项暂时不管
			// todo
		}
	}

	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}

