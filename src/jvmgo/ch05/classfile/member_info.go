package classfile

type FieldInfo MemberInfo
type MethodInfo MemberInfo

type MemberInfo struct {
	Cp *ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	attributeInfoCount uint16
	attributeInfos []AttributeInfo
}

func (self MemberInfo) Name() string {
	return self.Cp.getUtf8(self.nameIndex)
}

func (self MemberInfo) Descriptor() string {
	return self.Cp.getUtf8(self.descriptorIndex)
}

func (self MemberInfo) CodeAttribute() *Code {
	for _, attr := range self.attributeInfos {
		_, ok := attr.(*Code)
		if ok {
			return attr.(*Code)
		}
	}
	return nil
}



func readMemberInfo(reader *ClassReader, cp *ConstantPool) *MemberInfo {
	memberInfo := &MemberInfo{Cp:cp}
	memberInfo.accessFlags = reader.readUint16()
	memberInfo.nameIndex = reader.readUint16()
	memberInfo.descriptorIndex = reader.readUint16()
	memberInfo.attributeInfoCount = reader.readUint16()
	memberInfo.attributeInfos = make([]AttributeInfo, memberInfo.attributeInfoCount)
	for i := range memberInfo.attributeInfos {
		memberInfo.attributeInfos[i] = readAttributeInfo(reader, cp)
	}
	return memberInfo
}







