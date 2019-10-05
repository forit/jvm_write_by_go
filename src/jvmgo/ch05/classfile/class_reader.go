package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	value := self.data[0]
	self.data = self.data[1:]
	return value
}

func (self *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return value
}

func (self *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return value
}

func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	data := make([]uint16, n)
	for i := range data {
		data[i] = self.readUint16()
	}
	return data
}


