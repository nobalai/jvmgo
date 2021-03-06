package comparisons

import (
	"go/ch05/instructions/base"
	"go/ch05/rtda"
)

// Branch if int comparison with zero succeeds
type IFEQ struct { base.BranchInstruction }
type IFNQ struct { base.BranchInstruction }
type IFLT struct { base.BranchInstruction }
type IFLE struct { base.BranchInstruction }
type IFGT struct { base.BranchInstruction }
type IFGE struct { base.BranchInstruction }

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}
