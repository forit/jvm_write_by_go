package classfile

type ConstantValue struct {
	AttributeBaseInfo
	constantValueIndex uint16
}

func (self *ConstantValue) readAttrInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

// cpInfo:number string
func (self *ConstantValue) getCpInfo() CpInfo {
	return self.Cp.getConstantInfo(self.constantValueIndex)
}
func (self *ConstantValue) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}

