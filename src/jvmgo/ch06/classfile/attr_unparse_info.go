package classfile

type UnParseAttr struct {
	AttributeBaseInfo
	info []byte
}

func (self *UnParseAttr) readAttrInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.AttributeLength)
}

