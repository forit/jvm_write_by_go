package classfile

import "fmt"

type AttributeInfo interface {
	readAttrInfo(reader *ClassReader)
}

type AttributeBaseInfo struct {
	Cp                 *ConstantPool
	AttributeNameIndex uint16
	AttributeLength    uint32
}

func (self *AttributeBaseInfo) getAttrName() string {
	return self.Cp.getUtf8(self.AttributeNameIndex)
}

func readAttributeInfo(reader *ClassReader, constantPool *ConstantPool) AttributeInfo {
	attributeNameIndex := reader.readUint16()
	attributeLength := reader.readUint32()
	attributeBaseInfo := AttributeBaseInfo{Cp:constantPool, AttributeNameIndex:attributeNameIndex, AttributeLength:attributeLength}
	attributeInfo := newAttributeInfo(attributeBaseInfo, reader)
	attributeInfo.readAttrInfo(reader)
	return attributeInfo
}

func newAttributeInfo(baseInfo AttributeBaseInfo, reader *ClassReader) AttributeInfo {
	switch baseInfo.getAttrName() {
	case "ConstantValue":
		return &ConstantValue{AttributeBaseInfo: baseInfo}
	case "Code":
		return &Code{AttributeBaseInfo:baseInfo}
	case "LineNumberTable":
		return &LineNumberTableAttribute{AttributeBaseInfo:baseInfo}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{AttributeBaseInfo:baseInfo}
	case "Exceptions":
		return &Exceptions{AttributeBaseInfo:baseInfo}
	case "SourceFile":
		return &SourceFile{AttributeBaseInfo:baseInfo}
	default:
		fmt.Printf("\nnot supprot this attributeInfo:%s\n", baseInfo.getAttrName())
		return &UnParseAttr{AttributeBaseInfo:baseInfo}
	}
}

