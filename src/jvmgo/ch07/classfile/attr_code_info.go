package classfile

type Code struct {
	AttributeBaseInfo
	maxStack uint16
	maxLocals uint16
	codeLength uint32
	codes []byte
	exceptionTableLength uint16
	exceptionTables []ExceptionTable
	attributeCount uint16
	attributes      []AttributeInfo
}

type ExceptionTable struct {
	startPc uint16
	endPc	uint16
	handlerPc uint16
	catchType uint16
}

func (self *Code) readAttrInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	self.codeLength = reader.readUint32()
	self.codes = reader.readBytes(self.codeLength)
	self.exceptionTableLength = reader.readUint16()
	self.exceptionTables = make([]ExceptionTable, self.exceptionTableLength)
	for j := range self.exceptionTables {
		self.exceptionTables[j] = readExceptionTable(reader)
	}
	self.attributeCount = reader.readUint16()
	self.attributes = make([]AttributeInfo, self.attributeCount)
	for k := range self.attributes {
		self.attributes[k] = readAttributeInfo(reader, self.Cp)
	}
}
func readExceptionTable(reader *ClassReader) ExceptionTable {
	exceptionTable := ExceptionTable{}
	exceptionTable.startPc = reader.readUint16()
	exceptionTable.endPc = reader.readUint16()
	exceptionTable.handlerPc = reader.readUint16()
	exceptionTable.catchType = reader.readUint16()
	return exceptionTable
}

func (self *Code) MaxLocals() uint16{
	return self.maxLocals
}

func (self *Code) MaxStack() uint16{
	return self.maxStack
}

func (self *Code) ByteCodes() []byte{
	return self.codes
}
func (self *Code) Code() []byte {
	return self.codes
}





