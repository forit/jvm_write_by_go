package classfile

import "fmt"

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type CpInfo interface {
	readInfo(reader *ClassReader)
}

type CpBaseInfo struct {
	Tag uint8
	Cp *ConstantPool
}

func (self *CpBaseInfo) getUtf8(index uint16) string {
	return self.Cp.getUtf8(index)
}

func (self *CpBaseInfo) getTag() uint8 {
	return self.Tag
}

func readCpInfo(reader *ClassReader, constantPool *ConstantPool) CpInfo {
	tag := reader.readUint8()
	c := newCpInfo(tag, constantPool)
	c.readInfo(reader)
	return c
}

func newCpInfo(tag uint8, constantPool *ConstantPool) CpInfo {
	cpBaseInfo := CpBaseInfo{Tag: tag, Cp:constantPool}
	switch tag {
	case CONSTANT_Class :
		return &CpClassInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Fieldref :
			return &CpFieldrefInfo{CpRefInfo{CpBaseInfo: cpBaseInfo}}
	case CONSTANT_Methodref :
			return &CpMethodrefInfo{CpRefInfo{CpBaseInfo: cpBaseInfo}}
	case CONSTANT_InterfaceMethodref :
			return &CpInterfaceMethodrefInfo{CpRefInfo{CpBaseInfo: cpBaseInfo}}
	case CONSTANT_String :
			return &CpStringInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Integer :
			return &CpIntegerInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Float :
			return &CpFloatInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Long :
			return &CpLongInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Double :
			return &CpDoubleInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_NameAndType :
			return &CpNameAndTypeInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_Utf8 :
			return &CpUtf8Info{CpBaseInfo: cpBaseInfo}
	case CONSTANT_MethodHandle :
			return &CpMethdoHandleInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_MethodType :
			return &CpMethodTypeInfo{CpBaseInfo: cpBaseInfo}
	case CONSTANT_InvokeDynamic :
			return &CpInvokeDynamicInfo{CpBaseInfo: cpBaseInfo}
	default:
		panic(fmt.Sprintf("contant pool tag wrong, tag:%d" , tag))
	}
}

