package classfile

type LineNumberTable struct {
	startPc uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	AttributeBaseInfo
	lineNumberTableLength uint16
	lineNumberTables []LineNumberTable
}

func (self *LineNumberTableAttribute) readAttrInfo(reader *ClassReader) {
	self.lineNumberTableLength = reader.readUint16();
	self.lineNumberTables = make([]LineNumberTable, self.lineNumberTableLength)
	for i := range self.lineNumberTables {
		self.lineNumberTables[i] = readLineNumberTable(reader)
	}
}
func readLineNumberTable(reader *ClassReader) LineNumberTable {
	lineNumberTable := LineNumberTable{}
	lineNumberTable.startPc = reader.readUint16()
	lineNumberTable.lineNumber = reader.readUint16()
	return lineNumberTable
}

