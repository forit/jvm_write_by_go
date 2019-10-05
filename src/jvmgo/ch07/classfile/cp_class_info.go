package classfile

type CpClassInfo struct {
	CpBaseInfo
	classNameIndex uint16
}

func (self *CpClassInfo) readInfo(reader *ClassReader) {
	self.classNameIndex = reader.readUint16()
}

func (self *CpClassInfo) getClassName() string {
	return self.getUtf8(self.classNameIndex)
}
func (self *CpClassInfo) Name() string {
	return self.getUtf8(self.classNameIndex)
}

