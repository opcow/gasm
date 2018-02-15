package main

import (
	"fmt"
	"os"

	"github.com/opcow/cpu65"
)

var firstNode bool

func graphStart() {
	fmt.Fprintln(os.Stdout, "digraph 001 {")
	fmt.Fprintln(os.Stdout, "graph [layout = dot]  node [shape = plaintext]")
}

func graphEnd() {
	fmt.Fprintln(os.Stdout, "\n}")
}

func graphLabel(c *cpu65.CPU) {
	fmt.Fprintf(os.Stdout, "%d [label=\"%s\"];\n", c.PC(), c.Mnemonic())
}

func graphNode(a int) {
	if firstNode {
		fmt.Fprintf(os.Stdout, "%d", a)
		firstNode = false
	} else {
		fmt.Fprintf(os.Stdout, "\n -> %d", a)
	}
}

func graphBranch(s, d int) {
	fmt.Fprintf(os.Stdout, "%d -> %d;\n", s, d)
}

func graphJump(s, d int) {
	fmt.Fprintf(os.Stdout, "%d -> %d;\n", s, d)
}

func graphReturn(s, d int) {
	fmt.Fprintf(os.Stdout, "%d -> %d;\n", s, d)
}
