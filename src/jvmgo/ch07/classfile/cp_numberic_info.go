package classfile

import (
	"encoding/binary"
	"math"
)

type CpIntegerInfo struct {
	CpBaseInfo
	bytes []byte
}

func (self *CpIntegerInfo) readInfo(reader *ClassReader) {
	self.bytes = reader.readBytes(4)
}

func (self *CpIntegerInfo) Value() int32 {
	value := binary.BigEndian.Uint32(self.bytes)
	return int32(value)
}

type CpFloatInfo struct {
	CpBaseInfo
	bytes []byte
}

func (self *CpFloatInfo) readInfo(reader *ClassReader) {
	self.bytes = reader.readBytes(4)
}

func (self *CpFloatInfo) Value() float32 {
	value := binary.BigEndian.Uint32(self.bytes)
	return float32(math.Float32frombits(value))
}

type CpLongInfo struct {
	CpBaseInfo
	highBytes	[]byte
	lowBytes	[]byte
}

func (self *CpLongInfo) readInfo(reader *ClassReader) {
	self.highBytes = reader.readBytes(4)
	self.lowBytes = reader.readBytes(4)
}

func (self *CpLongInfo) Value() int64 {
	high := binary.BigEndian.Uint32(self.highBytes)
	low := binary.BigEndian.Uint32(self.lowBytes)
	value := uint64(high)<<32 | uint64(low)
	return int64(value)
}

type CpDoubleInfo struct {
	CpBaseInfo
	highBytes	[]byte
	lowBytes	[]byte
}

func (self *CpDoubleInfo) readInfo(reader *ClassReader) {
	self.highBytes = reader.readBytes(4)
	self.lowBytes = reader.readBytes(4)
}

func (self *CpDoubleInfo) Value() float64 {
	high := binary.BigEndian.Uint32(self.highBytes)
	low := binary.BigEndian.Uint32(self.lowBytes)
	value := uint64(high)<<32 | uint64(low)
	return float64(math.Float64frombits(value))
}