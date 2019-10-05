package classfile

type LocalVariableTable struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

type LocalVariableTableAttribute struct {
	AttributeBaseInfo
	localVariableTableLength uint16
	localVariableTables []LocalVariableTable
}

func (self *LocalVariableTableAttribute) readAttrInfo(reader *ClassReader) {
	self.localVariableTableLength = reader.readUint16()
	self.localVariableTables = make([]LocalVariableTable, self.localVariableTableLength)
	for i := range self.localVariableTables {
		self.localVariableTables[i] = readLocalVariableTables(reader)
	}
}
func readLocalVariableTables(reader *ClassReader) LocalVariableTable {
	localVariableTable := LocalVariableTable{}
	localVariableTable.startPc = reader.readUint16()
	localVariableTable.length = reader.readUint16()
	localVariableTable.nameIndex = reader.readUint16()
	localVariableTable.descriptorIndex = reader.readUint16()
	localVariableTable.index = reader.readUint16()
	return localVariableTable
}
