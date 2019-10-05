package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic           uint32
	minorVersion    uint16
	majorVersion    uint16
	constantPool    ConstantPool
	accessFlags     uint16
	thisClassIndex  uint16
	superClassIndex uint16
	interfaces      []uint16
	fields          []*MemberInfo
	methods         []*MemberInfo
	attributes      []AttributeInfo
}

func (self *ClassFile) read(reader *ClassReader) {
	//magic
	self.readAndCheckMagic(reader)
	//minor_version major_version
	self.readAndCheckVersion(reader)
	//constantPoll
	self.readConstantPool(reader)
	//access flags 按位表示flags
	self.accessFlags = reader.readUint16()
	//this class 常量池中constant_class_info
	self.thisClassIndex = reader.readUint16()
	//super class 常量池中constant_class_info
	self.superClassIndex = reader.readUint16()
	//interfaces 可能是index数组 index常量池中constant_class_info
	self.interfaces = reader.readUint16s()
	//fields
	self.readFields(reader)
	//methods
	self.readMethods(reader)
	//classfile attrs
	self.readAttrs(reader)

}
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0XCAFEBABE {
		panic(fmt.Sprintf("magic wrong, expect %s but %s",  "0XCAFEBABE", string(self.magic)))
	}
}
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:			//minorVersion could not zero
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic(fmt.Sprintf("unsupprot Versioh %d.%d", self.majorVersion, self.minorVersion))
}
func (self *ClassFile) readConstantPool(reader *ClassReader) {
	count := reader.readUint16()
	self.constantPool = readConstantPool(int(count), reader)
}

func (self *ClassFile) readFields(reader *ClassReader) {
	fieldsCount := reader.readUint16()
	self.fields = make([]*MemberInfo, fieldsCount)
	for i := range self.fields {
		self.fields[i] = readMemberInfo(reader, &self.constantPool)
	}
}
func (self *ClassFile) readMethods(reader *ClassReader) {
	methodCount := reader.readUint16()
	self.methods = make([]*MemberInfo, methodCount)
	for i := range self.methods {
		self.methods[i] = readMemberInfo(reader, &self.constantPool)
	}
}
func (self *ClassFile) readAttrs(reader *ClassReader) {
	attrCount := reader.readUint16()
	self.attributes = make([]AttributeInfo, attrCount)
	for i := range self.attributes {
		self.attributes[i] = readAttributeInfo(reader, &self.constantPool)
	}
}
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}
func (self *ClassFile) ClassName() string {
	return self.constantPool.getConstantInfo(self.thisClassIndex).(*CpClassInfo).getClassName()
}
func (self *ClassFile) SuperClassName() string {
	if self.superClassIndex == 0 {
		return ""
	}
	return self.constantPool.getConstantInfo(self.superClassIndex).(*CpClassInfo).getClassName()
}
func (self *ClassFile) InterfaceNames() []string {
	s := make([]string, len(self.interfaces))
	for i, index := range self.interfaces {
		s[i] = self.constantPool.getConstantInfo(index).(*CpClassInfo).getClassName()
	}
	return s
}
func (self *ClassFile) FieldCount() interface{} {
	return len(self.fields)
}
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}


func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf := &ClassFile{}
	cf.read(cr)
	return cf, nil
}

