//-----------------------------------------------------------------------------
/*

RISC-V ISA Definition

*/
//-----------------------------------------------------------------------------

package rvda

import (
	"fmt"
	"strings"
)

//-----------------------------------------------------------------------------

// daFunc is an instruction disassembly function
type daFunc func(name string, pc, ins uint) string

// insDefn is the definition of an instruction
type insDefn struct {
	defn string // instruction definition string (from the standard)
	da   daFunc // disassembly function
}

// ISAModule is a set/module of RISC-V instructions.
type ISAModule struct {
	ext  uint      // ISA extension bits per CSR misa
	ilen int       // instruction length
	defn []insDefn // instruction definitions
}

//-----------------------------------------------------------------------------
// RV32 instructions

// ISArv32i integer instructions.
var ISArv32i = ISAModule{
	ext:  misaExtI,
	ilen: 32,
	defn: []insDefn{
		{"imm[31:12] rd 0110111 LUI", daTypeUa},                         // U
		{"imm[31:12] rd 0010111 AUIPC", daTypeUa},                       // U
		{"imm[20|10:1|11|19:12] rd 1101111 JAL", daTypeJa},              // J
		{"imm[11:0] rs1 000 rd 1100111 JALR", daTypeIe},                 // I
		{"imm[12|10:5] rs2 rs1 000 imm[4:1|11] 1100011 BEQ", daTypeBa},  // B
		{"imm[12|10:5] rs2 rs1 001 imm[4:1|11] 1100011 BNE", daTypeBa},  // B
		{"imm[12|10:5] rs2 rs1 100 imm[4:1|11] 1100011 BLT", daTypeBa},  // B
		{"imm[12|10:5] rs2 rs1 101 imm[4:1|11] 1100011 BGE", daTypeBa},  // B
		{"imm[12|10:5] rs2 rs1 110 imm[4:1|11] 1100011 BLTU", daTypeBa}, // B
		{"imm[12|10:5] rs2 rs1 111 imm[4:1|11] 1100011 BGEU", daTypeBa}, // B
		{"imm[11:0] rs1 000 rd 0000011 LB", daTypeIc},                   // I
		{"imm[11:0] rs1 001 rd 0000011 LH", daTypeIc},                   // I
		{"imm[11:0] rs1 010 rd 0000011 LW", daTypeIc},                   // I
		{"imm[11:0] rs1 100 rd 0000011 LBU", daTypeIc},                  // I
		{"imm[11:0] rs1 101 rd 0000011 LHU", daTypeIc},                  // I
		{"imm[11:5] rs2 rs1 000 imm[4:0] 0100011 SB", daTypeSa},         // S
		{"imm[11:5] rs2 rs1 001 imm[4:0] 0100011 SH", daTypeSa},         // S
		{"imm[11:5] rs2 rs1 010 imm[4:0] 0100011 SW", daTypeSa},         // S
		{"imm[11:0] rs1 000 rd 0010011 ADDI", daTypeIb},                 // I
		{"imm[11:0] rs1 010 rd 0010011 SLTI", daTypeIa},                 // I
		{"imm[11:0] rs1 011 rd 0010011 SLTIU", daTypeIa},                // I
		{"imm[11:0] rs1 100 rd 0010011 XORI", daTypeIf},                 // I
		{"imm[11:0] rs1 110 rd 0010011 ORI", daTypeIa},                  // I
		{"imm[11:0] rs1 111 rd 0010011 ANDI", daTypeIa},                 // I
		{"000000 shamt6 rs1 001 rd 0010011 SLLI", daTypeId},             // I
		{"000000 shamt6 rs1 101 rd 0010011 SRLI", daTypeId},             // I
		{"010000 shamt6 rs1 101 rd 0010011 SRAI", daTypeId},             // I
		{"0000000 rs2 rs1 000 rd 0110011 ADD", daTypeRa},                // R
		{"0100000 rs2 rs1 000 rd 0110011 SUB", daTypeRa},                // R
		{"0000000 rs2 rs1 001 rd 0110011 SLL", daTypeRa},                // R
		{"0000000 rs2 rs1 010 rd 0110011 SLT", daTypeRa},                // R
		{"0000000 rs2 rs1 011 rd 0110011 SLTU", daTypeRa},               // R
		{"0000000 rs2 rs1 100 rd 0110011 XOR", daTypeRa},                // R
		{"0000000 rs2 rs1 101 rd 0110011 SRL", daTypeRa},                // R
		{"0100000 rs2 rs1 101 rd 0110011 SRA", daTypeRa},                // R
		{"0000000 rs2 rs1 110 rd 0110011 OR", daTypeRa},                 // R
		{"0000000 rs2 rs1 111 rd 0110011 AND", daTypeRa},                // R
		{"0000 pred succ 00000 000 00000 0001111 FENCE", daTypeIi},      // I
		{"0000 0000 0000 00000 001 00000 0001111 FENCE.I", daTypeIi},    // I
		{"0000000 00000 00000 000 00000 1110011 ECALL", daTypeIi},       // I
		{"0000000 00001 00000 000 00000 1110011 EBREAK", daTypeIi},      // I
		{"0000000 00010 00000 000 00000 1110011 URET", daTypeIi},        // I
		{"0001000 00010 00000 000 00000 1110011 SRET", daTypeIi},        // I
		{"0011000 00010 00000 000 00000 1110011 MRET", daTypeIi},        // I
		{"0001000 00101 00000 000 00000 1110011 WFI", daTypeIi},         // I
		{"0001001 rs2 rs1 000 00000 1110011 SFENCE.VMA", daTypeIk},      // I
		{"0010001 rs2 rs1 000 00000 1110011 HFENCE.BVMA", daTypeIk},     // I
		{"1010001 rs2 rs1 000 00000 1110011 HFENCE.GVMA", daTypeIk},     // I
		{"csr rs1 001 rd 1110011 CSRRW", daTypeIh},                      // I
		{"csr rs1 010 rd 1110011 CSRRS", daTypeIh},                      // I
		{"csr rs1 011 rd 1110011 CSRRC", daTypeIh},                      // I
		{"csr zimm 101 rd 1110011 CSRRWI", daTypeIj},                    // I
		{"csr zimm 110 rd 1110011 CSRRSI", daTypeIj},                    // I
		{"csr zimm 111 rd 1110011 CSRRCI", daTypeIj},                    // I
	},
}

// ISArv32m integer multiplication/division instructions.
var ISArv32m = ISAModule{
	ext:  misaExtM,
	ilen: 32,
	defn: []insDefn{
		{"0000001 rs2 rs1 000 rd 0110011 MUL", daTypeRa},    // R
		{"0000001 rs2 rs1 001 rd 0110011 MULH", daTypeRa},   // R
		{"0000001 rs2 rs1 010 rd 0110011 MULHSU", daTypeRa}, // R
		{"0000001 rs2 rs1 011 rd 0110011 MULHU", daTypeRa},  // R
		{"0000001 rs2 rs1 100 rd 0110011 DIV", daTypeRa},    // R
		{"0000001 rs2 rs1 101 rd 0110011 DIVU", daTypeRa},   // R
		{"0000001 rs2 rs1 110 rd 0110011 REM", daTypeRa},    // R
		{"0000001 rs2 rs1 111 rd 0110011 REMU", daTypeRa},   // R
	},
}

// ISArv32a atomic operation instructions.
var ISArv32a = ISAModule{
	ext:  misaExtA,
	ilen: 32,
	defn: []insDefn{
		{"00010 aq rl 00000 rs1 010 rd 0101111 LR.W", daTypeRb},    // R
		{"00011 aq rl rs2 rs1 010 rd 0101111 SC.W", daTypeRb},      // R
		{"00001 aq rl rs2 rs1 010 rd 0101111 AMOSWAP.W", daTypeRb}, // R
		{"00000 aq rl rs2 rs1 010 rd 0101111 AMOADD.W", daTypeRb},  // R
		{"00100 aq rl rs2 rs1 010 rd 0101111 AMOXOR.W", daTypeRb},  // R
		{"01100 aq rl rs2 rs1 010 rd 0101111 AMOAND.W", daTypeRb},  // R
		{"01000 aq rl rs2 rs1 010 rd 0101111 AMOOR.W", daTypeRb},   // R
		{"10000 aq rl rs2 rs1 010 rd 0101111 AMOMIN.W", daTypeRb},  // R
		{"10100 aq rl rs2 rs1 010 rd 0101111 AMOMAX.W", daTypeRb},  // R
		{"11000 aq rl rs2 rs1 010 rd 0101111 AMOMINU.W", daTypeRb}, // R
		{"11100 aq rl rs2 rs1 010 rd 0101111 AMOMAXU.W", daTypeRb}, // R
	},
}

// ISArv32f 32-bit floating point instructions.
var ISArv32f = ISAModule{
	ext:  misaExtF,
	ilen: 32,
	defn: []insDefn{
		{"imm[11:0] rs1 010 rd 0000111 FLW", daTypeIg},           // I
		{"imm[11:5] rs2 rs1 010 imm[4:0] 0100111 FSW", daTypeSb}, // S
		{"rs3 00 rs2 rs1 rm rd 1000011 FMADD.S", daTypeR4a},      // R4
		{"rs3 00 rs2 rs1 rm rd 1000111 FMSUB.S", daTypeR4a},      // R4
		{"rs3 00 rs2 rs1 rm rd 1001011 FNMSUB.S", daTypeR4a},     // R4
		{"rs3 00 rs2 rs1 rm rd 1001111 FNMADD.S", daTypeR4a},     // R4
		{"0000000 rs2 rs1 rm rd 1010011 FADD.S", daTypeRc},       // R
		{"0000100 rs2 rs1 rm rd 1010011 FSUB.S", daTypeRc},       // R
		{"0001000 rs2 rs1 rm rd 1010011 FMUL.S", daTypeRc},       // R
		{"0001100 rs2 rs1 rm rd 1010011 FDIV.S", daTypeRc},       // R
		{"0101100 00000 rs1 rm rd 1010011 FSQRT.S", daTypeRh},    // R
		{"0010000 rs2 rs1 000 rd 1010011 FSGNJ.S", daTypeRc},     // R
		{"0010000 rs2 rs1 001 rd 1010011 FSGNJN.S", daTypeRc},    // R
		{"0010000 rs2 rs1 010 rd 1010011 FSGNJX.S", daTypeRc},    // R
		{"0010100 rs2 rs1 000 rd 1010011 FMIN.S", daTypeRc},      // R
		{"0010100 rs2 rs1 001 rd 1010011 FMAX.S", daTypeRc},      // R
		{"1100000 00000 rs1 rm rd 1010011 FCVT.W.S", daTypeRk},   // R
		{"1100000 00001 rs1 rm rd 1010011 FCVT.WU.S", daTypeRk},  // R
		{"1110000 00000 rs1 000 rd 1010011 FMV.X.W", daTypeRd},   // R
		{"1010000 rs2 rs1 010 rd 1010011 FEQ.S", daTypeRf},       // R
		{"1010000 rs2 rs1 001 rd 1010011 FLT.S", daTypeRf},       // R
		{"1010000 rs2 rs1 000 rd 1010011 FLE.S", daTypeRf},       // R
		{"1110000 00000 rs1 001 rd 1010011 FCLASS.S", daTypeRd},  // R
		{"1101000 00000 rs1 rm rd 1010011 FCVT.S.W", daTypeRj},   // R
		{"1101000 00001 rs1 rm rd 1010011 FCVT.S.WU", daTypeRj},  // R
		{"1111000 00000 rs1 000 rd 1010011 FMV.W.X", daTypeRe},   // R
	},
}

// ISArv32d 64-bit floating point instructions.
var ISArv32d = ISAModule{
	ext:  misaExtD,
	ilen: 32,
	defn: []insDefn{
		{"imm[11:0] rs1 011 rd 0000111 FLD", daTypeIg},           // I
		{"imm[11:5] rs2 rs1 011 imm[4:0] 0100111 FSD", daTypeSb}, // S
		{"rs3 01 rs2 rs1 rm rd 1000011 FMADD.D", daTypeR4a},      // R4
		{"rs3 01 rs2 rs1 rm rd 1000111 FMSUB.D", daTypeR4a},      // R4
		{"rs3 01 rs2 rs1 rm rd 1001011 FNMSUB.D", daTypeR4a},     // R4
		{"rs3 01 rs2 rs1 rm rd 1001111 FNMADD.D", daTypeR4a},     // R4
		{"0000001 rs2 rs1 rm rd 1010011 FADD.D", daTypeRc},       // R
		{"0000101 rs2 rs1 rm rd 1010011 FSUB.D", daTypeRc},       // R
		{"0001001 rs2 rs1 rm rd 1010011 FMUL.D", daTypeRc},       // R
		{"0001101 rs2 rs1 rm rd 1010011 FDIV.D", daTypeRc},       // R
		{"0101101 00000 rs1 rm rd 1010011 FSQRT.D", daTypeRh},    // R
		{"0010001 rs2 rs1 000 rd 1010011 FSGNJ.D", daTypeRc},     // R
		{"0010001 rs2 rs1 001 rd 1010011 FSGNJN.D", daTypeRc},    // R
		{"0010001 rs2 rs1 010 rd 1010011 FSGNJX.D", daTypeRc},    // R
		{"0010101 rs2 rs1 000 rd 1010011 FMIN.D", daTypeRc},      // R
		{"0010101 rs2 rs1 001 rd 1010011 FMAX.D", daTypeRc},      // R
		{"0100000 00001 rs1 rm rd 1010011 FCVT.S.D", daTypeRi},   // R
		{"0100001 00000 rs1 rm rd 1010011 FCVT.D.S", daTypeRi},   // R
		{"1010001 rs2 rs1 010 rd 1010011 FEQ.D", daTypeRf},       // R
		{"1010001 rs2 rs1 001 rd 1010011 FLT.D", daTypeRf},       // R
		{"1010001 rs2 rs1 000 rd 1010011 FLE.D", daTypeRf},       // R
		{"1110001 00000 rs1 001 rd 1010011 FCLASS.D", daTypeRd},  // R
		{"1100001 00000 rs1 rm rd 1010011 FCVT.W.D", daTypeRk},   // R
		{"1100001 00001 rs1 rm rd 1010011 FCVT.WU.D", daTypeRk},  // R
		{"1101001 00000 rs1 rm rd 1010011 FCVT.D.W", daTypeRj},   // R
		{"1101001 00001 rs1 rm rd 1010011 FCVT.D.WU", daTypeRj},  // R
	},
}

// ISArv32c compressed instructions (subset of RV64C).
var ISArv32c = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		{"000 00000000 000 00 C.ILLEGAL", daTypeCIWa},                    // CIW (Quadrant 0)
		{"000 nzuimm[5:4|9:6|2|3] rd0 00 C.ADDI4SPN", daTypeCIWb},        // CIW
		{"010 uimm[5:3] rs10 uimm[2|6] rd0 00 C.LW", daTypeCSa},          // CL
		{"110 uimm[5:3] rs10 uimm[2|6] rs20 00 C.SW", daTypeCSa},         // CS
		{"000 nzimm[5] 00000 nzimm[4:0] 01 C.NOP", daNop},                // CI (Quadrant 1)
		{"000 nzimm[5] rs1/rd!=0 nzimm[4:0] 01 C.ADDI", daTypeCIc},       // CI
		{"010 imm[5] rd!=0 imm[4:0] 01 C.LI", daTypeCIa},                 // CI
		{"011 nzimm[9] 00010 nzimm[4|6|8:7|5] 01 C.ADDI16SP", daTypeCIb}, // CI
		{"011 nzimm[17] rd!={0,2} nzimm[16:12] 01 C.LUI", daTypeCIg},     // CI
		{"100 nzuimm[5] 00 rs10/rd0 nzuimm[4:0] 01 C.SRLI", daTypeCId},   // CI
		{"100 nzuimm[5] 01 rs10/rd0 nzuimm[4:0] 01 C.SRAI", daTypeCId},   // CI
		{"100 imm[5] 10 rs10/rd0 imm[4:0] 01 C.ANDI", daTypeCIf},         // CI
		{"100 0 11 rs10/rd0 00 rs20 01 C.SUB", daTypeCRc},                // CR
		{"100 0 11 rs10/rd0 01 rs20 01 C.XOR", daTypeCRc},                // CR
		{"100 0 11 rs10/rd0 10 rs20 01 C.OR", daTypeCRc},                 // CR
		{"100 0 11 rs10/rd0 11 rs20 01 C.AND", daTypeCRc},                // CR
		{"101 imm[11|4|9:8|10|6|7|3:1|5] 01 C.J", daTypeCJb},             // CJ
		{"110 imm[8|4:3] rs10 imm[7:6|2:1|5] 01 C.BEQZ", daTypeCBa},      // CB
		{"111 imm[8|4:3] rs10 imm[7:6|2:1|5] 01 C.BNEZ", daTypeCBa},      // CB
		{"000 nzuimm[5] rs1/rd!=0 nzuimm[4:0] 10 C.SLLI", daTypeCIe},     // CI (Quadrant 2)
		{"000 0 rs1/rd!=0 00000 10 C.SLLI64", daNone},                    // CI
		{"010 uimm[5] rd!=0 uimm[4:2|7:6] 10 C.LWSP", daTypeCSSa},        // CSS
		{"100 0 rs1!=0 00000 10 C.JR", daTypeCRd},                        // CR
		{"100 0 rd!=0 rs2!=0 10 C.MV", daTypeCRa},                        // CR
		{"100 1 00000 00000 10 C.EBREAK", daNone},                        // CI
		{"100 1 rs1!=0 00000 10 C.JALR", daTypeCRe},                      // CR
		{"100 1 rs1/rd!=0 rs2!=0 10 C.ADD", daTypeCRb},                   // CR
		{"110 uimm[5:2|7:6] rs2 10 C.SWSP", daTypeCSSb},                  // CSS
	},
}

// ISArv32cOnly compressed instructions (not in RV64C).
var ISArv32cOnly = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		{"001 imm[11|4|9:8|10|6|7|3:1|5] 01 C.JAL", daTypeCJc}, // CJ
	},
}

// ISArv32fc compressed 32-bit floating point instructions.
var ISArv32fc = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		{"011 uimm[5:3] rs10 uimm[2|6] rd0 00 C.FLW", daTypeCSc},  // CL
		{"011 uimm[5] rd uimm[4:2|7:6] 10 C.FLWSP", daNone},       // CSS
		{"111 uimm[5:3] rs10 uimm[2|6] rs20 00 C.FSW", daTypeCSc}, // CS
		{"111 uimm[5:2|7:6] rs2 10 C.FSWSP", daNone},              // CSS
	},
}

// ISArv32dc compressed 64-bit floating point instructions.
var ISArv32dc = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		{"001 uimm[5:3] rs10 uimm[7:6] rd0 00 C.FLD", daTypeCSc},  // CL
		{"001 uimm[5] rd uimm[4:3|8:6] 10 C.FLDSP", daNone},       // CSS
		{"101 uimm[5:3] rs10 uimm[7:6] rs20 00 C.FSD", daTypeCSc}, // CS
		{"101 uimm[5:3|8:6] rs2 10 C.FSDSP", daNone},              // CSS
	},
}

//-----------------------------------------------------------------------------
// RV64 instructions (+ RV32)

// ISArv64i Integer
var ISArv64i = ISAModule{
	ext:  misaExtI,
	ilen: 32,
	defn: []insDefn{
		{"imm[11:0] rs1 110 rd 0000011 LWU", daTypeIc},          // I
		{"imm[11:0] rs1 011 rd 0000011 LD", daTypeIa},           // I
		{"imm[11:5] rs2 rs1 011 imm[4:0] 0100011 SD", daTypeSa}, // S
		{"000000 shamt6 rs1 001 rd 0010011 SLLI", daTypeId},     // I
		{"000000 shamt6 rs1 101 rd 0010011 SRLI", daTypeId},     // I
		{"010000 shamt6 rs1 101 rd 0010011 SRAI", daTypeId},     // I
		{"imm[11:0] rs1 000 rd 0011011 ADDIW", daTypeIa},        // I
		{"0000000 shamt5 rs1 001 rd 0011011 SLLIW", daTypeId},   // I
		{"0000000 shamt5 rs1 101 rd 0011011 SRLIW", daTypeId},   // I
		{"0100000 shamt5 rs1 101 rd 0011011 SRAIW", daTypeId},   // I
		{"0000000 rs2 rs1 000 rd 0111011 ADDW", daTypeRa},       // R
		{"0100000 rs2 rs1 000 rd 0111011 SUBW", daTypeRa},       // R
		{"0000000 rs2 rs1 001 rd 0111011 SLLW", daTypeRa},       // R
		{"0000000 rs2 rs1 101 rd 0111011 SRLW", daTypeRa},       // R
		{"0100000 rs2 rs1 101 rd 0111011 SRAW", daTypeRa},       // R
	},
}

// ISArv64m Integer Multiplication and Division
var ISArv64m = ISAModule{
	ext:  misaExtM,
	ilen: 32,
	defn: []insDefn{
		{"0000001 rs2 rs1 000 rd 0111011 MULW", daTypeRa},  // R
		{"0000001 rs2 rs1 100 rd 0111011 DIVW", daTypeRa},  // R
		{"0000001 rs2 rs1 101 rd 0111011 DIVUW", daTypeRa}, // R
		{"0000001 rs2 rs1 110 rd 0111011 REMW", daTypeRa},  // R
		{"0000001 rs2 rs1 111 rd 0111011 REMUW", daTypeRa}, // R
	},
}

// ISArv64a Atomics
var ISArv64a = ISAModule{
	ext:  misaExtA,
	ilen: 32,
	defn: []insDefn{
		{"00010 aq rl 00000 rs1 011 rd 0101111 LR.D", daTypeRb},    // R
		{"00011 aq rl rs2 rs1 011 rd 0101111 SC.D", daTypeRb},      // R
		{"00001 aq rl rs2 rs1 011 rd 0101111 AMOSWAP.D", daTypeRb}, // R
		{"00000 aq rl rs2 rs1 011 rd 0101111 AMOADD.D", daTypeRb},  // R
		{"00100 aq rl rs2 rs1 011 rd 0101111 AMOXOR.D", daTypeRb},  // R
		{"01100 aq rl rs2 rs1 011 rd 0101111 AMOAND.D", daTypeRb},  // R
		{"01000 aq rl rs2 rs1 011 rd 0101111 AMOOR.D", daTypeRb},   // R
		{"10000 aq rl rs2 rs1 011 rd 0101111 AMOMIN.D", daTypeRb},  // R
		{"10100 aq rl rs2 rs1 011 rd 0101111 AMOMAX.D", daTypeRb},  // R
		{"11000 aq rl rs2 rs1 011 rd 0101111 AMOMINU.D", daTypeRb}, // R
		{"11100 aq rl rs2 rs1 011 rd 0101111 AMOMAXU.D", daTypeRb}, // R
	},
}

// ISArv64f Single-Precision Floating-Point
var ISArv64f = ISAModule{
	ext:  misaExtF,
	ilen: 32,
	defn: []insDefn{
		{"1100000 00010 rs1 rm rd 1010011 FCVT.L.S", daTypeRk},  // R
		{"1100000 00011 rs1 rm rd 1010011 FCVT.LU.S", daTypeRk}, // R
		{"1101000 00010 rs1 rm rd 1010011 FCVT.S.L", daTypeRj},  // R
		{"1101000 00011 rs1 rm rd 1010011 FCVT.S.LU", daTypeRj}, // R
	},
}

// ISArv64d Double-Precision Floating-Point
var ISArv64d = ISAModule{
	ext:  misaExtD,
	ilen: 32,
	defn: []insDefn{
		{"1100001 00010 rs1 rm rd 1010011 FCVT.L.D", daTypeRk},  // R
		{"1100001 00011 rs1 rm rd 1010011 FCVT.LU.D", daTypeRk}, // R
		{"1110001 00000 rs1 000 rd 1010011 FMV.X.D", daTypeRd},  // R
		{"1101001 00010 rs1 rm rd 1010011 FCVT.D.L", daTypeRj},  // R
		{"1101001 00011 rs1 rm rd 1010011 FCVT.D.LU", daTypeRj}, // R
		{"1111001 00000 rs1 000 rd 1010011 FMV.D.X", daTypeRe},  // R
	},
}

// ISArv64c Compressed
var ISArv64c = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		{"001 imm[5] rd!=0 imm[4:0] 01 C.ADDIW", daTypeCIc},      // CI
		{"011 uimm[5] rd uimm[4:3|8:6] 10 C.LDSP", daTypeCIh},    // CI
		{"011 uimm[5:3] rs10 uimm[7:6] rd0 00 C.LD", daTypeCSb},  // CL
		{"100 1 11 rs10/rd0 00 rs20 01 C.SUBW", daTypeCRc},       // CR
		{"100 1 11 rs10/rd0 01 rs20 01 C.ADDW", daTypeCRc},       // CR
		{"111 uimm[5:3] rs10 uimm[7:6] rs20 00 C.SD", daTypeCSb}, // CS
		{"111 uimm[5:3|8:6] rs2 10 C.SDSP", daTypeCSSc},          // CSS
	},
}

//-----------------------------------------------------------------------------

// ISArv128c Compressed
var ISArv128c = ISAModule{
	ext:  misaExtC,
	ilen: 16,
	defn: []insDefn{
		// C.SQ
		// C.LQ
		// C.LQSP
	},
}

//-----------------------------------------------------------------------------
// pre-canned ISA module sets

// ISArv32g = RV32imafd
var ISArv32g = []ISAModule{
	ISArv32i, ISArv32m, ISArv32a, ISArv32f, ISArv32d,
}

// ISArv32gc = RV32imafdc
var ISArv32gc = []ISAModule{
	ISArv32i, ISArv32m, ISArv32a, ISArv32f, ISArv32d,
	ISArv32c, ISArv32cOnly, ISArv32fc, ISArv32dc,
}

// ISArv64g = RV64imafd
var ISArv64g = []ISAModule{
	ISArv32i, ISArv32m, ISArv32a, ISArv32f, ISArv32d,
	ISArv64i, ISArv64m, ISArv64a, ISArv64f, ISArv64d,
}

// ISArv64gc = RV64imafdc
var ISArv64gc = []ISAModule{
	ISArv32i, ISArv32m, ISArv32a, ISArv32f, ISArv32d,
	ISArv32c, ISArv32dc,
	ISArv64i, ISArv64m, ISArv64a, ISArv64f, ISArv64d,
	ISArv64c,
}

//-----------------------------------------------------------------------------

// insMeta is instruction meta-data determined at runtime
type insMeta struct {
	defn      *insDefn   // the instruction definition
	name      string     // instruction mneumonic
	n         int        // instruction bit length
	val, mask uint       // value and mask of fixed bits in the instruction
	dt        decodeType // decode type
}

// decodeConstant returns go code for decoding constants for this instruction.
func (im *insMeta) decodeConstant() string {
	s := []string{}
	name := strings.ReplaceAll(im.name, ".", "_")
	name = strings.ToUpper(name)
	s = append(s, fmt.Sprintf("opcode%s = 0x", name))
	if im.n == 16 {
		s = append(s, fmt.Sprintf("%04x", im.val))
	} else {
		s = append(s, fmt.Sprintf("%08x", im.val))
	}
	s = append(s, fmt.Sprintf(" // %s", im.name))
	return strings.Join(s, "")
}

//-----------------------------------------------------------------------------

// ISA is an instruction set
type ISA struct {
	ext   uint       // ISA extension bits matching misa CSR
	ins16 []*insMeta // the set of 16-bit instructions in the ISA
	ins32 []*insMeta // the set of 32-bit instructions in the ISA
}

// NewISA creates an empty instruction set.
func NewISA(ext uint) *ISA {
	return &ISA{
		ext:   ext,
		ins16: make([]*insMeta, 0),
		ins32: make([]*insMeta, 0),
	}
}

// Add a ISA sub-module to the ISA.
func (isa *ISA) Add(module []ISAModule) error {
	for i := range module {
		isa.ext |= module[i].ext
		for j := range module[i].defn {
			im, err := parseDefn(&module[i].defn[j], module[i].ilen)
			if err != nil {
				return err
			}
			if im.n == 16 {
				isa.ins16 = append(isa.ins16, im)
			} else {
				isa.ins32 = append(isa.ins32, im)
			}
		}
	}
	return nil
}

// lookup returns the instruction meta information for an instruction.
func (isa *ISA) lookup(ins uint) *insMeta {
	if ins&3 == 3 {
		// 32-bit instruction
		for _, im := range isa.ins32 {
			if ins&im.mask == im.val {
				return im
			}
		}
	} else {
		// 16-bit instruction
		for _, im := range isa.ins16 {
			if ins&im.mask == im.val {
				return im
			}
		}
	}
	return nil
}

// GetExtensions returns the ISA extension bits.
func (isa *ISA) GetExtensions() uint {
	return isa.ext
}

//-----------------------------------------------------------------------------
