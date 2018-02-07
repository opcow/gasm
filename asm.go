package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/golang-collections/collections/stack"
)

const (
	maxMem  = 0x10000
	maxSegs = 64
)

type segCounter struct {
	count int
	seg   [maxSegs]struct {
		start int
		end   int
	}
}

var memory [maxMem]byte
var memmap [maxMem]byte

var segments segCounter
var jumps *stack.Stack

var printInst [imp + 1]func(instr *instruction)

func initPrintFuncs() {
	printInst[imm] = printImm
	printInst[zpa] = printZp
	printInst[zpx] = printZpx
	printInst[abs] = printAbs
	printInst[abx] = printAbsx
	printInst[aby] = printAbsy
	printInst[ind] = printInd
	printInst[inx] = printIndx
	printInst[iny] = printIndy
	printInst[acc] = printAcc
	printInst[rel] = printRel
	printInst[imp] = printImp
}

func printImm(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s #$%02x\n", instr.mnemonic, instr.ops[0]))
}

func printZp(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s $%02x\n", instr.mnemonic, instr.ops[0]))
}

func printZpx(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s $%02x,X\n", instr.mnemonic, instr.ops[0]))
}

func printAbs(instr *instruction) {
	if instr.code == 0x20 { // JSR
		if memmap[instr.branch]&(1<<2) != 0 {
			printToFile(of, fmt.Sprintf("%s sub_%02x\n", instr.mnemonic, instr.branch))
		}
	} else {
		printToFile(of, fmt.Sprintf("%s $%02x%02x\n", instr.mnemonic, instr.ops[1], instr.ops[0]))
	}
}

func printAbsx(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s $%02x%02x,X\n", instr.mnemonic, instr.ops[1], instr.ops[0]))
}

func printAbsy(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s $%02x%02x,Y\n", instr.mnemonic, instr.ops[1], instr.ops[0]))
}

func printInd(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s ($%02x%02x)\n", instr.mnemonic, instr.ops[1], instr.ops[0]))
}

func printIndx(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s ($%02x,X)\n", instr.mnemonic, instr.ops[0]))
}

func printIndy(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s ($%02x),Y\n", instr.mnemonic, instr.ops[0]))
}

func printAcc(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s \n", instr.mnemonic))
}

func printRel(instr *instruction) {
	if memmap[instr.branch]&(1<<1) != 0 {
		printToFile(of, fmt.Sprintf("%s loc_%02x\n", instr.mnemonic, instr.branch))
	} else {
		printToFile(of, fmt.Sprintf("%s $%02x\n", instr.mnemonic, instr.ops[0]))
	}
}

func printImp(instr *instruction) {
	printToFile(of, fmt.Sprintf("%s\n", instr.mnemonic))
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

func disAsm() {

	for i := 0; i < segments.count; i++ {
		disAsmSegP1(segments.seg[i].start, segments.seg[i].end)
	}
	for i := 0; i < segments.count; i++ {
		printToFile(of, fmt.Sprintf("\n    .ORG $%04x ; seg %02d | %d bytes\n\n", segments.seg[i].start, i,
			segments.seg[i].end-segments.seg[i].start))
		disAsmSegP2(segments.seg[i].start, segments.seg[i].end)
	}
}

// disassemble pass 2
func disAsmSegP2(addr, end int) {
	var instr instruction

	for addr < end {
		if memmap[addr]&(1<<1) != 0 {
			printToFile(of, fmt.Sprintf("loc_%04x\n", addr))
		}
		if memmap[addr]&(1<<2) != 0 {
			printToFile(of, fmt.Sprintf("sub_%04x\n", addr))
		}
		if memmap[addr] == 0 {
			addr += printDataBlock(addr, end)
			continue
		}
		fetchInstr(addr, &instr)
		if *prAddr {
			printLine(addr, &instr)
		}
		printToFile(of, fmt.Sprint("    "))
		printInst[instr.mode](&instr)
		addr = instr.next
	}
}

// disassemble pass 1 follows and records jumps/branches
// so that labels and data blocks can be created
func disAsmSegP1(addr, end int) {
	var instr instruction
	jumps = stack.New()

	for addr < end {
		fetchInstr(addr, &instr)
		for i := addr; i < instr.next; i++ {
			memmap[i] |= (1 << 0)
		}
		if instr.code == 0x20 {
			memmap[instr.branch] |= (1 << 2)
			if isAddressInSeg(instr.branch) >= 0 {
				jumps.Push(instr.branch)
			}
		} else if instr.mode == rel {
			memmap[instr.branch] |= (1 << 1)
		} else if instr.code == 0x60 {
			if jumps.Len() == 0 {
				return
			}
			addr = jumps.Pop().(int)
			continue
		}
		addr = instr.next
	}
}

func printLine(addr int, in *instruction) {
	var buffer bytes.Buffer

	buffer.Write([]byte(fmt.Sprintf("%04X ", addr)))
	buffer.Write([]byte(fmt.Sprintf("%02X ", in.code)))
	if in.length == 2 {
		buffer.Write([]byte(fmt.Sprintf("%02X ", in.ops[0])))
	} else if in.length == 3 {
		buffer.Write([]byte(fmt.Sprintf("%02X %02X ", in.ops[0], in.ops[1])))
	}
	buffer.Write([]byte(fmt.Sprintf("%-16s", buffer.String())))
	buffer.WriteTo(of)
}

func isAddressInSeg(addr int) int {

	for i := 0; i < segments.count; i++ {
		if addr >= segments.seg[i].start && addr < segments.seg[i].end {
			return i
		}
	}
	return -1
}
