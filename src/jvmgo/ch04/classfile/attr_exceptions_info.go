package classfile

type Exceptions struct {
	AttributeBaseInfo
	numberOfExceptions uint16
	exceptionIndexs []uint16
}

func (self *Exceptions) readAttrInfo(reader *ClassReader) {
	self.numberOfExceptions = reader.readUint16()
	self.exceptionIndexs = make([]uint16, self.numberOfExceptions)
	for i := range self.exceptionIndexs {
		self.exceptionIndexs[i] = reader.readUint16()
	}
}

