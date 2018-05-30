package comparisons

import (
	"go/ch05/instructions/base"
	"go/ch05/rtda"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct { base.BranchInstruction }
type IF_ICMPNE struct { base.BranchInstruction }
type IF_ICMPLT struct { base.BranchInstruction }
type IF_ICMPLE struct { base.BranchInstruction }
type IF_ICMPGT struct { base.BranchInstruction }
type IF_ICMPGE struct { base.BranchInstruction }

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, self.Offset)
	}
}
