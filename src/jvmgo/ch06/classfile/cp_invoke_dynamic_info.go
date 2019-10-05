package classfile

type CpInvokeDynamicInfo struct {
	CpBaseInfo
	bootStrapMethodAttrIndex uint16 //classfile bootStrapMethod 属性表中获取
	nameAndTypeIndex uint16
}

func (self *CpInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootStrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

type CpMethodTypeInfo struct {
	CpBaseInfo
	descriptor uint16
}

func (self *CpMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptor = reader.readUint16()
}

type CpMethdoHandleInfo struct {
	CpBaseInfo
	referenceKind uint8 //1-9
	referenceIndex uint16 //索引cp_ref_info中struct
}

func (self *CpMethdoHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}
