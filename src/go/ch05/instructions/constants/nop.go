package constants

import (
	"go/ch05/instructions/base"
	"go/ch05/rtda"
)

// Do nothing
type NOP struct { base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// Do nothing
}
