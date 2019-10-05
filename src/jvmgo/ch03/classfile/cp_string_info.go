package classfile

type CpStringInfo struct {
	CpBaseInfo
	stringIndex uint16
}


func (self *CpStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *CpStringInfo) getString() string {
	return self.getUtf8(self.stringIndex)
}


