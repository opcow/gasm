package main

const (
	unk = iota
	imm
	zpa
	zpx
	zpy
	abs
	abx
	aby
	ind
	inx
	iny
	acc
	rel
	imp
)

type opcode struct {
	code     int
	length   int
	mode     int
	mnemonic string
}

var opcodes = [256]opcode{
	0x69: opcode{0x69, 2, imm, "ADC"},
	0x65: opcode{0x65, 2, zpa, "ADC"},
	0x75: opcode{0x75, 2, zpx, "ADC"},
	0x6d: opcode{0x6d, 3, abs, "ADC"},
	0x7d: opcode{0x7d, 3, abx, "ADC"},
	0x79: opcode{0x79, 3, aby, "ADC"},
	0x61: opcode{0x61, 2, inx, "ADC"},
	0x71: opcode{0x71, 2, iny, "ADC"},
	0x29: opcode{0x29, 2, imm, "AND"},
	0x25: opcode{0x25, 2, zpa, "AND"},
	0x35: opcode{0x35, 2, zpx, "AND"},
	0x2d: opcode{0x2d, 3, abs, "AND"},
	0x3d: opcode{0x3d, 3, abx, "AND"},
	0x39: opcode{0x39, 3, aby, "AND"},
	0x21: opcode{0x21, 2, inx, "AND"},
	0x31: opcode{0x31, 2, iny, "AND"},
	0x0a: opcode{0x0a, 1, acc, "ASL"},
	0x06: opcode{0x06, 2, zpa, "ASL"},
	0x16: opcode{0x16, 2, zpx, "ASL"},
	0x0e: opcode{0x0e, 3, abs, "ASL"},
	0x1e: opcode{0x1e, 3, abx, "ASL"},
	0x24: opcode{0x24, 2, zpa, "BIT"},
	0x2c: opcode{0x2c, 3, abs, "BIT"},
	0x10: opcode{0x10, 2, rel, "BPL"},
	0x30: opcode{0x30, 2, rel, "BMI"},
	0x50: opcode{0x50, 2, rel, "BVC"},
	0x70: opcode{0x70, 2, rel, "BVS"},
	0x90: opcode{0x90, 2, rel, "BCC"},
	0xb0: opcode{0xb0, 2, rel, "BCS"},
	0xd0: opcode{0xd0, 2, rel, "BNE"},
	0xf0: opcode{0xf0, 2, rel, "BEQ"},
	0x00: opcode{0x00, 1, imp, "BRK"},
	0xc9: opcode{0xc9, 2, imm, "CMP"},
	0xc5: opcode{0xc5, 2, zpa, "CMP"},
	0xd5: opcode{0xd5, 2, zpx, "CMP"},
	0xcd: opcode{0xcd, 3, abs, "CMP"},
	0xdd: opcode{0xdd, 3, abx, "CMP"},
	0xd9: opcode{0xd9, 3, aby, "CMP"},
	0xc1: opcode{0xc1, 2, inx, "CMP"},
	0xd1: opcode{0xd1, 2, iny, "CMP"},
	0xe0: opcode{0xe0, 2, imm, "CPX"},
	0xe4: opcode{0xe4, 2, zpa, "CPX"},
	0xec: opcode{0xec, 3, abs, "CPX"},
	0xc0: opcode{0xc0, 2, imm, "CPY"},
	0xc4: opcode{0xc4, 2, zpa, "CPY"},
	0xcc: opcode{0xcc, 3, abs, "CPY"},
	0xc6: opcode{0xc6, 2, zpa, "DEC"},
	0xd6: opcode{0xd6, 2, zpx, "DEC"},
	0xce: opcode{0xce, 3, abs, "DEC"},
	0xde: opcode{0xde, 3, abx, "DEC"},
	0x49: opcode{0x49, 2, imm, "EOR"},
	0x45: opcode{0x45, 2, zpa, "EOR"},
	0x55: opcode{0x55, 2, zpx, "EOR"},
	0x4d: opcode{0x4d, 3, abs, "EOR"},
	0x5d: opcode{0x5d, 3, abx, "EOR"},
	0x59: opcode{0x59, 3, aby, "EOR"},
	0x41: opcode{0x41, 2, inx, "EOR"},
	0x51: opcode{0x51, 2, iny, "EOR"},
	0x18: opcode{0x18, 1, imp, "CLC"},
	0x38: opcode{0x38, 1, imp, "SEC"},
	0x58: opcode{0x58, 1, imp, "CLI"},
	0x78: opcode{0x78, 1, imp, "SEI"},
	0xb8: opcode{0xb8, 1, imp, "CLV"},
	0xd8: opcode{0xd8, 1, imp, "CLD"},
	0xf8: opcode{0xf8, 1, imp, "SED"},
	0xe6: opcode{0xe6, 2, zpa, "INC"},
	0xf6: opcode{0xf6, 2, zpx, "INC"},
	0xee: opcode{0xee, 3, abs, "INC"},
	0xfe: opcode{0xfe, 3, abx, "INC"},
	0x4c: opcode{0x4c, 3, abs, "JMP"},
	0x6c: opcode{0x6c, 3, ind, "JMP"},
	0x20: opcode{0x20, 3, abs, "JSR"},
	0xa9: opcode{0xa9, 2, imm, "LDA"},
	0xa5: opcode{0xa5, 2, zpa, "LDA"},
	0xb5: opcode{0xb5, 2, zpx, "LDA"},
	0xad: opcode{0xad, 3, abs, "LDA"},
	0xbd: opcode{0xbd, 3, abx, "LDA"},
	0xb9: opcode{0xb9, 3, aby, "LDA"},
	0xa1: opcode{0xa1, 2, inx, "LDA"},
	0xb1: opcode{0xb1, 2, iny, "LDA"},
	0xa2: opcode{0xa2, 2, imm, "LDX"},
	0xa6: opcode{0xa6, 2, zpa, "LDX"},
	0xb6: opcode{0xb6, 2, zpy, "LDX"},
	0xae: opcode{0xae, 3, abs, "LDX"},
	0xbe: opcode{0xbe, 3, aby, "LDX"},
	0xa0: opcode{0xa0, 2, imm, "LDY"},
	0xa4: opcode{0xa4, 2, zpa, "LDY"},
	0xb4: opcode{0xb4, 2, zpx, "LDY"},
	0xac: opcode{0xac, 3, abs, "LDY"},
	0xbc: opcode{0xbc, 3, abx, "LDY"},
	0x4a: opcode{0x4a, 1, acc, "LSR"},
	0x46: opcode{0x46, 2, zpa, "LSR"},
	0x56: opcode{0x56, 2, zpx, "LSR"},
	0x4e: opcode{0x4e, 3, abs, "LSR"},
	0x5e: opcode{0x5e, 3, abx, "LSR"},
	0xea: opcode{0xea, 1, imp, "NOP"},
	0x09: opcode{0x09, 2, imm, "ORA"},
	0x05: opcode{0x05, 2, zpa, "ORA"},
	0x15: opcode{0x15, 2, zpx, "ORA"},
	0x0d: opcode{0x0d, 3, abs, "ORA"},
	0x1d: opcode{0x1d, 3, abx, "ORA"},
	0x19: opcode{0x19, 3, aby, "ORA"},
	0x01: opcode{0x01, 2, inx, "ORA"},
	0x11: opcode{0x11, 2, iny, "ORA"},
	0xaa: opcode{0xaa, 1, imp, "TAX"},
	0x8a: opcode{0x8a, 1, imp, "TXA"},
	0xca: opcode{0xca, 1, imp, "DEX"},
	0xe8: opcode{0xe8, 1, imp, "INX"},
	0xa8: opcode{0xa8, 1, imp, "TAY"},
	0x98: opcode{0x98, 1, imp, "TYA"},
	0x88: opcode{0x88, 1, imp, "DEY"},
	0xc8: opcode{0xc8, 1, imp, "INY"},
	0x2a: opcode{0x2a, 1, acc, "ROL"},
	0x26: opcode{0x26, 2, zpa, "ROL"},
	0x36: opcode{0x36, 2, zpx, "ROL"},
	0x2e: opcode{0x2e, 3, abs, "ROL"},
	0x3e: opcode{0x3e, 3, abx, "ROL"},
	0x6a: opcode{0x6a, 1, acc, "ROR"},
	0x66: opcode{0x66, 2, zpa, "ROR"},
	0x76: opcode{0x76, 2, zpx, "ROR"},
	0x6e: opcode{0x6e, 3, abs, "ROR"},
	0x7e: opcode{0x7e, 3, abx, "ROR"},
	0x40: opcode{0x40, 1, imp, "RTI"},
	0x60: opcode{0x60, 1, imp, "RTS"},
	0xe9: opcode{0xe9, 2, imm, "SBC"},
	0xe5: opcode{0xe5, 2, zpa, "SBC"},
	0xf5: opcode{0xf5, 2, zpx, "SBC"},
	0xed: opcode{0xed, 3, abs, "SBC"},
	0xfd: opcode{0xfd, 3, abx, "SBC"},
	0xf9: opcode{0xf9, 3, aby, "SBC"},
	0xe1: opcode{0xe1, 2, inx, "SBC"},
	0xf1: opcode{0xf1, 2, iny, "SBC"},
	0x85: opcode{0x85, 2, zpa, "STA"},
	0x95: opcode{0x95, 2, zpx, "STA"},
	0x8d: opcode{0x8d, 3, abs, "STA"},
	0x9d: opcode{0x9d, 3, abx, "STA"},
	0x99: opcode{0x99, 3, aby, "STA"},
	0x81: opcode{0x81, 2, inx, "STA"},
	0x91: opcode{0x91, 2, iny, "STA"},
	0x9a: opcode{0x9a, 1, imp, "TXS"},
	0xba: opcode{0xba, 1, imp, "TSX"},
	0x48: opcode{0x48, 1, imp, "PHA"},
	0x68: opcode{0x68, 1, imp, "PLA"},
	0x08: opcode{0x08, 1, imp, "PHP"},
	0x28: opcode{0x28, 1, imp, "PLP"},
	0x86: opcode{0x86, 2, zpa, "STX"},
	0x96: opcode{0x96, 2, zpy, "STX"},
	0x8e: opcode{0x8e, 3, abs, "STX"},
	0x84: opcode{0x84, 2, zpa, "STY"},
	0x94: opcode{0x94, 2, zpx, "STY"},
	0x8c: opcode{0x8c, 3, abs, "STY"},
}

type instruction struct {
	code     int
	addr     int
	length   int
	mode     int
	next     int
	branch   int
	ops      [2]byte
	mnemonic string
}

func fetchInstr(addr int, in *instruction) {

	in.code = int(memory[addr])
	in.mnemonic = opcodes[in.code].mnemonic
	in.addr = addr
	in.length = opcodes[in.code].length

	if in.length > 1 {
		in.ops[0] = memory[addr+1]
	}
	if in.length == 3 {
		in.ops[1] = memory[addr+2]
	}

	in.next = addr + in.length
	in.mode = opcodes[in.code].mode

	if in.code == 0x20 || in.code == 0x4c {
		in.branch = int(memory[addr+1]) | int(memory[addr+2])<<8
	} else if in.mode == rel {
		in.branch = in.next + int(int8(memory[addr+1]))
	} else {
		in.branch = -1
	}
}
