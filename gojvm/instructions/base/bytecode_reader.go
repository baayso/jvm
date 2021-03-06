package base

type BytecodeReader struct {
	code []byte // 存放字节码
	pc   int    // 记录读取到了哪个字节
}

func (b *BytecodeReader) Reset(code []byte, pc int) {
	b.code = code
	b.pc = pc
}

func (b *BytecodeReader) ReadUint8() uint8 {
	i := b.code[b.pc]

	b.pc++

	return i
}

func (b *BytecodeReader) ReadInt8() int8 {
	return int8(b.ReadUint8())
}

func (b *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(b.ReadUint8())
	byte2 := uint16(b.ReadUint8())

	return (byte1 << 8) | byte2
}

func (b *BytecodeReader) ReadInt16() int16 {
	return int16(b.ReadUint16())
}

func (b *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(b.ReadUint8())
	byte2 := int32(b.ReadUint8())
	byte3 := int32(b.ReadUint8())
	byte4 := int32(b.ReadUint8())

	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
