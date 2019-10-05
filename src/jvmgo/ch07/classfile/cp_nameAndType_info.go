package classfile

type CpNameAndTypeInfo struct {
	CpBaseInfo
	nameIndex uint16
	descriptorIndex uint16
}

func (self *CpNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}

func (self *CpNameAndTypeInfo) getName() string {
	return self.getUtf8(self.nameIndex)
}

func (self *CpNameAndTypeInfo) getDescriptor() string {
	return self.getUtf8(self.descriptorIndex)
}


