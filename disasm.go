// Package gasm a 6502 disassembler/assembler
package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
	"github.com/opcow/cpu65"
)

const (
	maxSegs = 64
)

type segCounter struct {
	count int
	seg   [maxSegs]struct {
		start int
		end   int
	}
}

var memory [cpu65.MaxMem]byte // segment will load into 64k memory
var memmap [cpu65.MaxMem]byte // used for tracing control flow, etc

var segments segCounter

var cpu cpu65.CPU

func disAsm() {

	for i := 0; i < segments.count; i++ {
		disAsmSegP1(segments.seg[i].start, segments.seg[i].end)
	}
	for i := 0; i < segments.count; i++ {
		printToFile(of, fmt.Sprintf("\n    .ORG $%04x ; seg %02d | %d bytes\n\n", segments.seg[i].start, i,
			segments.seg[i].end-segments.seg[i].start))
		disAsmSegP2(segments.seg[i].start, segments.seg[i].end)
	}
	graphStart()
	for i := 0; i < segments.count; i++ {
		printLabels(segments.seg[i].start, segments.seg[i].end)
	}
	//	for i := 0; i < segments.count; i++ {
	printNodes(segments.seg[0].start)
	//	}
	for i := 0; i < segments.count; i++ {
		printBranches(segments.seg[i].start, segments.seg[i].end)
	}
	graphEnd()
}

func printLabels(addr, end int) {

	cpu.SetPC(addr)
	for i := addr; i < end; {
		if memmap[i] == 0 {
			i++
			continue
		}
		cpu.FetchInstr()
		graphLabel(&cpu)
		i = cpu.Next()
	}
}

func printNodes(addr int) {

	returns := stack.New()
	cpu.SetPC(addr)
	firstNode = true
	for {
		if memmap[addr] == 0 {
			break
		}
		cpu.FetchInstr()
		graphNode(cpu.PC())
		if cpu.Opcode() == 0x20 {
			returns.Push(addr + cpu.InsLen())
			j := cpu.AbsJumpAddr()
			addr = j
			cpu.SetPC(j)
		} else if cpu.Opcode() == 0x60 {
			if returns.Len() == 0 {
				break
			}
			addr = returns.Pop().(int)
			cpu.SetPC(addr)
		} else {
			addr = cpu.Next()
		}
	}
	fmt.Println(";")
}

func printBranches(addr, end int) {

	cpu.SetPC(addr)
	firstNode = true
	for i := addr; i < end; {
		if memmap[i] == 0 {
			i++
			continue
		}
		cpu.FetchInstr()
		if cpu.Mode() == cpu65.Rel {
			graphBranch(i, cpu.BranchAddr())
		}
		//graphNode(&cpu)
		i = cpu.Next()
	}
}

// disassemble pass 2
func disAsmSegP2(addr, end int) {
	cpu.SetPC(addr)

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
		cpu.FetchInstr()
		if *prAddr {
			printLine()
		}
		printToFile(of, fmt.Sprint("    "))
		printInst[cpu.Mode()](&cpu)
		addr = cpu.Next()
	}
}

// disassemble pass 1 follows and records jumps/branches
// so that labels and data blocks can be created
func disAsmSegP1(addr, end int) {
	jumps := stack.New()
	cpu.AttachMem(&memory)
	cpu.SetPC(addr)
	var brAddr int

FETCH:
	for {
		cpu.FetchInstr()
		// leave breadcrumbs where we've been
		for i := addr; i < addr+cpu.InsLen(); i++ {
			memmap[i] |= (1 << 0)
		}
		if cpu.Opcode() == 0x20 {
			// if JSR to somewhere mark that location
			// and push it onto the stack
			brAddr = cpu.AbsJumpAddr()
			memmap[brAddr] |= (1 << 2)
			if isAddressInSeg(brAddr) >= 0 {
				jumps.Push(brAddr)
			}
		} else if cpu.Opcode() == 0x4c {
			// absolute jump
			addr = cpu.AbsJumpAddr()
			cpu.SetPC(addr)
			continue
		} else if cpu.Mode() == cpu65.Rel {
			// if it's a branch just mark the location
			// for label printing
			memmap[cpu.BranchAddr()] |= (1 << 1)
		} else if cpu.Opcode() == 0x60 || cpu.Opcode() == 0x6c {
			// if RTS then pop any jumps and follow
			// also stop on indirect jump
			for jumps.Len() != 0 {
				addr = jumps.Pop().(int)
				// don't follow if location has been visited
				if memmap[addr]&(1<<0) != 1 {
					cpu.SetPC(addr)
					continue FETCH
				}
			}
			return
		} else if cpu.Opcode() == 0 {
			return
		}
		addr = cpu.Next()
	}
}

func isAddressInSeg(addr int) int {

	for i := 0; i < segments.count; i++ {
		if addr >= segments.seg[i].start && addr < segments.seg[i].end {
			return i
		}
	}
	return -1
}
