package classfile


type CpRefInfo struct {
	CpBaseInfo
	classNameIndex   uint16
	nameAndTypeIndex uint16
}

type CpFieldrefInfo struct {
	CpRefInfo
}

type CpMethodrefInfo struct {
	CpRefInfo
}

type CpInterfaceMethodrefInfo struct {
	CpRefInfo
}

func (self *CpRefInfo) readInfo(reader *ClassReader) {
	self.classNameIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *CpRefInfo) getRefName() string {
	return self.Cp.getConstantInfo(self.nameAndTypeIndex).(*CpNameAndTypeInfo).getName()
}

func (self *CpRefInfo) getRefDescriptor() string {
	return self.Cp.getConstantInfo(self.nameAndTypeIndex).(*CpNameAndTypeInfo).getDescriptor()
}
func (self *CpRefInfo) ClassName() string {
	cpClassInfo := self.Cp.getConstantInfo(self.classNameIndex).(*CpClassInfo)
	return cpClassInfo.getClassName()
}
func (self *CpRefInfo) NameAndDescriptor() (string, string) {
	return self.getRefName(), self.getRefDescriptor()
}