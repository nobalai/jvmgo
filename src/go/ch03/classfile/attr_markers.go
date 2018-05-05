package classfile

type MarkerAttribute struct {}

func (self *MarkAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

type DeprecatedAttribute struct { MarkerAttribute }
type SyntheticAttribute struct { MarkerAttribute }
