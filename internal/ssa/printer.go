package ssa

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/ssa"
)

func PrintFunction(fn *ssa.Function) {
	fmt.Printf("Basic blocks: %d\n", len(fn.Blocks))

	for _, block := range fn.Blocks {
		fmt.Printf("\nBlock %d (%s)\n", block.Index, block.Comment)
		fmt.Printf("  Preds:\n    %s\n", blocksAsString(block.Preds))
		fmt.Printf("  Succs:\n    %s\n", blocksAsString(block.Succs))
		fmt.Printf("  Instructions (%d):\n", len(block.Instrs))
		for _, instr := range block.Instrs {
			fmt.Printf("    %T: %s\n", instr, instr)
		}
	}
}

func blocksAsString(blocks []*ssa.BasicBlock) string {
	if len(blocks) == 0 {
		return "-"
	}
	indexes := make([]string, len(blocks))
	for i, block := range blocks {
		indexes[i] = fmt.Sprintf("Block %d", block.Index)
	}
	return strings.Join(indexes, "\n    ")
}
