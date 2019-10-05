package classfile

type SourceFile struct {
	AttributeBaseInfo
	sourceFileIndex uint16
}

func (self *SourceFile) readAttrInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}