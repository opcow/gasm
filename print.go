package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/opcow/cpu65"
)

var printInst = [cpu65.Modes]func(*cpu65.CPU){
	cpu65.Imm: printImm,
	cpu65.Zpa: printZpa,
	cpu65.Zpx: printZpx,
	cpu65.Abs: printAbs,
	cpu65.Abx: printAbx,
	cpu65.Aby: printAby,
	cpu65.Ind: printInd,
	cpu65.Inx: printInx,
	cpu65.Iny: printIny,
	cpu65.Acc: printAcc,
	cpu65.Rel: printRel,
	cpu65.Imp: printImp,
}

func printImm(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s #$%02x\n", p.Mnemonic(), p.Op(0)))
}

func printZpa(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s $%02x\n", p.Mnemonic(), p.Op(0)))
}

func printZpx(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s $%02x,X\n", p.Mnemonic(), p.Op(0)))
}

func printAbs(p *cpu65.CPU) {
	if p.Opcode() == 0x20 { // JSR
		if memmap[p.AbsJumpAddr()]&(1<<2) != 0 {
			printToFile(of, fmt.Sprintf("%s sub_%02x\n", p.Mnemonic(), p.AbsJumpAddr()))
		}
	} else {
		printToFile(of, fmt.Sprintf("%s $%02x%02x\n", p.Mnemonic(), p.Op(1), p.Op(0)))
	}
}

func printAbx(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s $%02x%02x,X\n", p.Mnemonic(), p.Op(1), p.Op(0)))
}

func printAby(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s $%02x%02x,Y\n", p.Mnemonic(), p.Op(1), p.Op(0)))
}

func printInd(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s ($%02x%02x)\n", p.Mnemonic(), p.Op(1), p.Op(0)))
}

func printInx(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s ($%02x,X)\n", p.Mnemonic(), p.Op(0)))
}

func printIny(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s ($%02x),Y\n", p.Mnemonic(), p.Op(0)))
}

func printAcc(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s \n", p.Mnemonic()))
}

func printRel(p *cpu65.CPU) {
	if memmap[p.BranchAddr()]&(1<<1) != 0 {
		printToFile(of, fmt.Sprintf("%s loc_%02x\n", p.Mnemonic(), p.BranchAddr()))
	} else {
		printToFile(of, fmt.Sprintf("%s $%02x\n", p.Mnemonic(), p.Op(0)))
	}
}

func printImp(p *cpu65.CPU) {
	printToFile(of, fmt.Sprintf("%s\n", p.Mnemonic()))
}

func printToFile(f *os.File, s string) {
	var b bytes.Buffer
	b.Write([]byte(s))
	b.WriteTo(f)
}

func printDataBlock(addr, end int) int {
	if *prAddr {
		printToFile(of, fmt.Sprintf("%04X            ", addr))
	}
	printToFile(of, fmt.Sprintf("    .byte $%02x", memory[addr]))
	i := 1
	for ; i < 8 && addr+i < end && memmap[addr+i] == 0; i++ {
		printToFile(of, fmt.Sprintf(", $%02x", memory[addr+i]))
	}
	printToFile(of, fmt.Sprint("\n"))
	return i
}

func printLine() {
	var buffer bytes.Buffer

	s := fmt.Sprintf("%04X %02X", cpu.PC(), cpu.Opcode())
	if cpu.InsLen() == 2 {
		s = fmt.Sprintf("%s %02X ", s, cpu.Op(0))
	} else if cpu.InsLen() == 3 {
		s = fmt.Sprintf("%s %02X %02X ", s, cpu.Op(0), cpu.Op(1))
	}
	buffer.Write([]byte(fmt.Sprintf("%-16s", s)))
	buffer.WriteTo(of)
}
