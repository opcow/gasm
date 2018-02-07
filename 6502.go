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
	0x69: opcode{code: 0x69, length: 2, mode: imm, mnemonic: "ADC"},
	0x65: opcode{code: 0x65, length: 2, mode: zpa, mnemonic: "ADC"},
	0x75: opcode{code: 0x75, length: 2, mode: zpx, mnemonic: "ADC"},
	0x6d: opcode{code: 0x6d, length: 3, mode: abs, mnemonic: "ADC"},
	0x7d: opcode{code: 0x7d, length: 3, mode: abx, mnemonic: "ADC"},
	0x79: opcode{code: 0x79, length: 3, mode: aby, mnemonic: "ADC"},
	0x61: opcode{code: 0x61, length: 2, mode: inx, mnemonic: "ADC"},
	0x71: opcode{code: 0x71, length: 2, mode: iny, mnemonic: "ADC"},
	0x29: opcode{code: 0x29, length: 2, mode: imm, mnemonic: "AND"},
	0x25: opcode{code: 0x25, length: 2, mode: zpa, mnemonic: "AND"},
	0x35: opcode{code: 0x35, length: 2, mode: zpx, mnemonic: "AND"},
	0x2d: opcode{code: 0x2d, length: 3, mode: abs, mnemonic: "AND"},
	0x3d: opcode{code: 0x3d, length: 3, mode: abx, mnemonic: "AND"},
	0x39: opcode{code: 0x39, length: 3, mode: aby, mnemonic: "AND"},
	0x21: opcode{code: 0x21, length: 2, mode: inx, mnemonic: "AND"},
	0x31: opcode{code: 0x31, length: 2, mode: iny, mnemonic: "AND"},
	0x0a: opcode{code: 0x0a, length: 1, mode: acc, mnemonic: "ASL"},
	0x06: opcode{code: 0x06, length: 2, mode: zpa, mnemonic: "ASL"},
	0x16: opcode{code: 0x16, length: 2, mode: zpx, mnemonic: "ASL"},
	0x0e: opcode{code: 0x0e, length: 3, mode: abs, mnemonic: "ASL"},
	0x1e: opcode{code: 0x1e, length: 3, mode: abx, mnemonic: "ASL"},
	0x24: opcode{code: 0x24, length: 2, mode: zpa, mnemonic: "BIT"},
	0x2c: opcode{code: 0x2c, length: 3, mode: abs, mnemonic: "BIT"},
	0x10: opcode{code: 0x10, length: 2, mode: rel, mnemonic: "BPL"},
	0x30: opcode{code: 0x30, length: 2, mode: rel, mnemonic: "BMI"},
	0x50: opcode{code: 0x50, length: 2, mode: rel, mnemonic: "BVC"},
	0x70: opcode{code: 0x70, length: 2, mode: rel, mnemonic: "BVS"},
	0x90: opcode{code: 0x90, length: 2, mode: rel, mnemonic: "BCC"},
	0xb0: opcode{code: 0xb0, length: 2, mode: rel, mnemonic: "BCS"},
	0xd0: opcode{code: 0xd0, length: 2, mode: rel, mnemonic: "BNE"},
	0xf0: opcode{code: 0xf0, length: 2, mode: rel, mnemonic: "BEQ"},
	0x00: opcode{code: 0x00, length: 1, mode: imp, mnemonic: "BRK"},
	0xc9: opcode{code: 0xc9, length: 2, mode: imm, mnemonic: "CMP"},
	0xc5: opcode{code: 0xc5, length: 2, mode: zpa, mnemonic: "CMP"},
	0xd5: opcode{code: 0xd5, length: 2, mode: zpx, mnemonic: "CMP"},
	0xcd: opcode{code: 0xcd, length: 3, mode: abs, mnemonic: "CMP"},
	0xdd: opcode{code: 0xdd, length: 3, mode: abx, mnemonic: "CMP"},
	0xd9: opcode{code: 0xd9, length: 3, mode: aby, mnemonic: "CMP"},
	0xc1: opcode{code: 0xc1, length: 2, mode: inx, mnemonic: "CMP"},
	0xd1: opcode{code: 0xd1, length: 2, mode: iny, mnemonic: "CMP"},
	0xe0: opcode{code: 0xe0, length: 2, mode: imm, mnemonic: "CPX"},
	0xe4: opcode{code: 0xe4, length: 2, mode: zpa, mnemonic: "CPX"},
	0xec: opcode{code: 0xec, length: 3, mode: abs, mnemonic: "CPX"},
	0xc0: opcode{code: 0xc0, length: 2, mode: imm, mnemonic: "CPY"},
	0xc4: opcode{code: 0xc4, length: 2, mode: zpa, mnemonic: "CPY"},
	0xcc: opcode{code: 0xcc, length: 3, mode: abs, mnemonic: "CPY"},
	0xc6: opcode{code: 0xc6, length: 2, mode: zpa, mnemonic: "DEC"},
	0xd6: opcode{code: 0xd6, length: 2, mode: zpx, mnemonic: "DEC"},
	0xce: opcode{code: 0xce, length: 3, mode: abs, mnemonic: "DEC"},
	0xde: opcode{code: 0xde, length: 3, mode: abx, mnemonic: "DEC"},
	0x49: opcode{code: 0x49, length: 2, mode: imm, mnemonic: "EOR"},
	0x45: opcode{code: 0x45, length: 2, mode: zpa, mnemonic: "EOR"},
	0x55: opcode{code: 0x55, length: 2, mode: zpx, mnemonic: "EOR"},
	0x4d: opcode{code: 0x4d, length: 3, mode: abs, mnemonic: "EOR"},
	0x5d: opcode{code: 0x5d, length: 3, mode: abx, mnemonic: "EOR"},
	0x59: opcode{code: 0x59, length: 3, mode: aby, mnemonic: "EOR"},
	0x41: opcode{code: 0x41, length: 2, mode: inx, mnemonic: "EOR"},
	0x51: opcode{code: 0x51, length: 2, mode: iny, mnemonic: "EOR"},
	0x18: opcode{code: 0x18, length: 1, mode: imp, mnemonic: "CLC"},
	0x38: opcode{code: 0x38, length: 1, mode: imp, mnemonic: "SEC"},
	0x58: opcode{code: 0x58, length: 1, mode: imp, mnemonic: "CLI"},
	0x78: opcode{code: 0x78, length: 1, mode: imp, mnemonic: "SEI"},
	0xb8: opcode{code: 0xb8, length: 1, mode: imp, mnemonic: "CLV"},
	0xd8: opcode{code: 0xd8, length: 1, mode: imp, mnemonic: "CLD"},
	0xf8: opcode{code: 0xf8, length: 1, mode: imp, mnemonic: "SED"},
	0xe6: opcode{code: 0xe6, length: 2, mode: zpa, mnemonic: "INC"},
	0xf6: opcode{code: 0xf6, length: 2, mode: zpx, mnemonic: "INC"},
	0xee: opcode{code: 0xee, length: 3, mode: abs, mnemonic: "INC"},
	0xfe: opcode{code: 0xfe, length: 3, mode: abx, mnemonic: "INC"},
	0x4c: opcode{code: 0x4c, length: 3, mode: abs, mnemonic: "JMP"},
	0x6c: opcode{code: 0x6c, length: 3, mode: ind, mnemonic: "JMP"},
	0x20: opcode{code: 0x20, length: 3, mode: abs, mnemonic: "JSR"},
	0xa9: opcode{code: 0xa9, length: 2, mode: imm, mnemonic: "LDA"},
	0xa5: opcode{code: 0xa5, length: 2, mode: zpa, mnemonic: "LDA"},
	0xb5: opcode{code: 0xb5, length: 2, mode: zpx, mnemonic: "LDA"},
	0xad: opcode{code: 0xad, length: 3, mode: abs, mnemonic: "LDA"},
	0xbd: opcode{code: 0xbd, length: 3, mode: abx, mnemonic: "LDA"},
	0xb9: opcode{code: 0xb9, length: 3, mode: aby, mnemonic: "LDA"},
	0xa1: opcode{code: 0xa1, length: 2, mode: inx, mnemonic: "LDA"},
	0xb1: opcode{code: 0xb1, length: 2, mode: iny, mnemonic: "LDA"},
	0xa2: opcode{code: 0xa2, length: 2, mode: imm, mnemonic: "LDX"},
	0xa6: opcode{code: 0xa6, length: 2, mode: zpa, mnemonic: "LDX"},
	0xb6: opcode{code: 0xb6, length: 2, mode: zpy, mnemonic: "LDX"},
	0xae: opcode{code: 0xae, length: 3, mode: abs, mnemonic: "LDX"},
	0xbe: opcode{code: 0xbe, length: 3, mode: aby, mnemonic: "LDX"},
	0xa0: opcode{code: 0xa0, length: 2, mode: imm, mnemonic: "LDY"},
	0xa4: opcode{code: 0xa4, length: 2, mode: zpa, mnemonic: "LDY"},
	0xb4: opcode{code: 0xb4, length: 2, mode: zpx, mnemonic: "LDY"},
	0xac: opcode{code: 0xac, length: 3, mode: abs, mnemonic: "LDY"},
	0xbc: opcode{code: 0xbc, length: 3, mode: abx, mnemonic: "LDY"},
	0x4a: opcode{code: 0x4a, length: 1, mode: acc, mnemonic: "LSR"},
	0x46: opcode{code: 0x46, length: 2, mode: zpa, mnemonic: "LSR"},
	0x56: opcode{code: 0x56, length: 2, mode: zpx, mnemonic: "LSR"},
	0x4e: opcode{code: 0x4e, length: 3, mode: abs, mnemonic: "LSR"},
	0x5e: opcode{code: 0x5e, length: 3, mode: abx, mnemonic: "LSR"},
	0xea: opcode{code: 0xea, length: 1, mode: imp, mnemonic: "NOP"},
	0x09: opcode{code: 0x09, length: 2, mode: imm, mnemonic: "ORA"},
	0x05: opcode{code: 0x05, length: 2, mode: zpa, mnemonic: "ORA"},
	0x15: opcode{code: 0x15, length: 2, mode: zpx, mnemonic: "ORA"},
	0x0d: opcode{code: 0x0d, length: 3, mode: abs, mnemonic: "ORA"},
	0x1d: opcode{code: 0x1d, length: 3, mode: abx, mnemonic: "ORA"},
	0x19: opcode{code: 0x19, length: 3, mode: aby, mnemonic: "ORA"},
	0x01: opcode{code: 0x01, length: 2, mode: inx, mnemonic: "ORA"},
	0x11: opcode{code: 0x11, length: 2, mode: iny, mnemonic: "ORA"},
	0xaa: opcode{code: 0xaa, length: 1, mode: imp, mnemonic: "TAX"},
	0x8a: opcode{code: 0x8a, length: 1, mode: imp, mnemonic: "TXA"},
	0xca: opcode{code: 0xca, length: 1, mode: imp, mnemonic: "DEX"},
	0xe8: opcode{code: 0xe8, length: 1, mode: imp, mnemonic: "INX"},
	0xa8: opcode{code: 0xa8, length: 1, mode: imp, mnemonic: "TAY"},
	0x98: opcode{code: 0x98, length: 1, mode: imp, mnemonic: "TYA"},
	0x88: opcode{code: 0x88, length: 1, mode: imp, mnemonic: "DEY"},
	0xc8: opcode{code: 0xc8, length: 1, mode: imp, mnemonic: "INY"},
	0x2a: opcode{code: 0x2a, length: 1, mode: acc, mnemonic: "ROL"},
	0x26: opcode{code: 0x26, length: 2, mode: zpa, mnemonic: "ROL"},
	0x36: opcode{code: 0x36, length: 2, mode: zpx, mnemonic: "ROL"},
	0x2e: opcode{code: 0x2e, length: 3, mode: abs, mnemonic: "ROL"},
	0x3e: opcode{code: 0x3e, length: 3, mode: abx, mnemonic: "ROL"},
	0x6a: opcode{code: 0x6a, length: 1, mode: acc, mnemonic: "ROR"},
	0x66: opcode{code: 0x66, length: 2, mode: zpa, mnemonic: "ROR"},
	0x76: opcode{code: 0x76, length: 2, mode: zpx, mnemonic: "ROR"},
	0x6e: opcode{code: 0x6e, length: 3, mode: abs, mnemonic: "ROR"},
	0x7e: opcode{code: 0x7e, length: 3, mode: abx, mnemonic: "ROR"},
	0x40: opcode{code: 0x40, length: 1, mode: imp, mnemonic: "RTI"},
	0x60: opcode{code: 0x60, length: 1, mode: imp, mnemonic: "RTS"},
	0xe9: opcode{code: 0xe9, length: 2, mode: imm, mnemonic: "SBC"},
	0xe5: opcode{code: 0xe5, length: 2, mode: zpa, mnemonic: "SBC"},
	0xf5: opcode{code: 0xf5, length: 2, mode: zpx, mnemonic: "SBC"},
	0xed: opcode{code: 0xed, length: 3, mode: abs, mnemonic: "SBC"},
	0xfd: opcode{code: 0xfd, length: 3, mode: abx, mnemonic: "SBC"},
	0xf9: opcode{code: 0xf9, length: 3, mode: aby, mnemonic: "SBC"},
	0xe1: opcode{code: 0xe1, length: 2, mode: inx, mnemonic: "SBC"},
	0xf1: opcode{code: 0xf1, length: 2, mode: iny, mnemonic: "SBC"},
	0x85: opcode{code: 0x85, length: 2, mode: zpa, mnemonic: "STA"},
	0x95: opcode{code: 0x95, length: 2, mode: zpx, mnemonic: "STA"},
	0x8d: opcode{code: 0x8d, length: 3, mode: abs, mnemonic: "STA"},
	0x9d: opcode{code: 0x9d, length: 3, mode: abx, mnemonic: "STA"},
	0x99: opcode{code: 0x99, length: 3, mode: aby, mnemonic: "STA"},
	0x81: opcode{code: 0x81, length: 2, mode: inx, mnemonic: "STA"},
	0x91: opcode{code: 0x91, length: 2, mode: iny, mnemonic: "STA"},
	0x9a: opcode{code: 0x9a, length: 1, mode: imp, mnemonic: "TXS"},
	0xba: opcode{code: 0xba, length: 1, mode: imp, mnemonic: "TSX"},
	0x48: opcode{code: 0x48, length: 1, mode: imp, mnemonic: "PHA"},
	0x68: opcode{code: 0x68, length: 1, mode: imp, mnemonic: "PLA"},
	0x08: opcode{code: 0x08, length: 1, mode: imp, mnemonic: "PHP"},
	0x28: opcode{code: 0x28, length: 1, mode: imp, mnemonic: "PLP"},
	0x86: opcode{code: 0x86, length: 2, mode: zpa, mnemonic: "STX"},
	0x96: opcode{code: 0x96, length: 2, mode: zpy, mnemonic: "STX"},
	0x8e: opcode{code: 0x8e, length: 3, mode: abs, mnemonic: "STX"},
	0x84: opcode{code: 0x84, length: 2, mode: zpa, mnemonic: "STY"},
	0x94: opcode{code: 0x94, length: 2, mode: zpx, mnemonic: "STY"},
	0x8c: opcode{code: 0x8c, length: 3, mode: abs, mnemonic: "STY"},
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
