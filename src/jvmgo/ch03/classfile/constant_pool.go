package classfile

import "fmt"

type ConstantPool []CpInfo

func (self ConstantPool) getConstantInfo(index uint16) CpInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("invalid constantPool index:%d", index))
}

func (self *ConstantPool) getUtf8(index uint16) string {
	constantUtf8Info := self.getConstantInfo(index).(*CpUtf8Info)
	return constantUtf8Info.getUtf8()
}

func readConstantPool(constantPoolNum int, reader *ClassReader) ConstantPool {
	constantPool := make(ConstantPool, constantPoolNum)
	for i := 1; i < constantPoolNum; i++ {
		constantPool[i] = readCpInfo(reader, &constantPool)

		//占用两个表成员项
		switch constantPool[i].(type) {
		case *CpLongInfo, *CpDoubleInfo:
			i++
		}
	}

	return constantPool
}

