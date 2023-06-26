//-----------------------------------------------------------------------------
/*

RISC-V Disassembler

*/
//-----------------------------------------------------------------------------

package rvda

import "fmt"

//-----------------------------------------------------------------------------

var abiXName = [32]string{
	"zero", "ra", "sp", "gp", "tp", "t0", "t1", "t2",
	"s0", "s1", "a0", "a1", "a2", "a3", "a4", "a5",
	"a6", "a7", "s2", "s3", "s4", "s5", "s6", "s7",
	"s8", "s9", "s10", "s11", "t3", "t4", "t5", "t6",
}

var abiFName = [32]string{
	"ft0", "ft1", "ft2", "ft3", "ft4", "ft5", "ft6", "ft7",
	"fs0", "fs1", "fa0", "fa1", "fa2", "fa3", "fa4", "fa5",
	"fa6", "fa7", "fs2", "fs3", "fs4", "fs5", "fs6", "fs7",
	"fs8", "fs9", "fs10", "fs11", "ft8", "ft9", "ft10", "ft11",
}

//-----------------------------------------------------------------------------
// default decode

func daNone(name string, pc uint, ins uint) string {
	return fmt.Sprintf("%s TODO", name)
}

//-----------------------------------------------------------------------------
// Type I Decodes

func daTypeIa(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], abiXName[rs1], imm)
}

func daTypeIb(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	if rd == 0 && rs1 == 0 && imm == 0 {
		return fmt.Sprintf("nop")
	}
	if rs1 == 0 {
		return fmt.Sprintf("li %s,%d", abiXName[rd], imm)
	}
	if imm == 0 {
		return fmt.Sprintf("mv %s,%s", abiXName[rd], abiXName[rs1])
	}
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], abiXName[rs1], imm)
}

func daTypeIc(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiXName[rd], imm, abiXName[rs1])
}

func daTypeId(name string, pc uint, ins uint) string {
	shamt, rs1, rd := decodeIc(ins)
	return fmt.Sprintf("%s %s,%s,0x%x", name, abiXName[rd], abiXName[rs1], shamt)
}

func daTypeIe(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	if imm == 0 && rd == 0 && rs1 == 1 {
		return "ret"
	}
	if rd == 1 {
		if imm == 0 {
			return fmt.Sprintf("%s %s", name, abiXName[rs1])
		}
		return fmt.Sprintf("%s %d(%s)", name, imm, abiXName[rs1])
	}
	if imm == 0 {
		return fmt.Sprintf("%s %s,%s", name, abiXName[rd], abiXName[rs1])
	}
	return fmt.Sprintf("%s %s,%d(%s)", name, abiXName[rd], imm, abiXName[rs1])
}

func daTypeIf(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	if imm == -1 {
		return fmt.Sprintf("not %s,%s", abiXName[rd], abiXName[rs1])
	}
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], abiXName[rs1], imm)
}

func daTypeIg(name string, pc uint, ins uint) string {
	imm, rs1, rd := decodeIa(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiFName[rd], imm, abiXName[rs1])
}

func daTypeIh(name string, pc uint, ins uint) string {
	csrReg, rs1, rd := decodeIb(ins)

	if csrReg == csrFCSR {
		if rd == 0 {
			return fmt.Sprintf("fscsr %s", abiXName[rs1])
		}
		if rs1 == 0 {
			return fmt.Sprintf("frcsr %s", abiXName[rd])
		}
		return fmt.Sprintf("fscsr %s,%s", abiXName[rd], abiXName[rs1])
	}

	if csrReg == csrFFLAGS {
		if rs1 == 0 && name == "csrrs" {
			return fmt.Sprintf("frflags %s", abiXName[rd])
		}
		return fmt.Sprintf("fsflags %s,%s", abiXName[rd], abiXName[rs1])
	}

	if rd == 0 {
		return fmt.Sprintf("%s %s,%s", csrRemap1(name), csrName(csrReg), abiXName[rs1])
	}

	if rs1 == 0 && name == "csrrs" {
		return fmt.Sprintf("%s %s,%s", csrRemap2(name), abiXName[rd], csrName(csrReg))
	}

	return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], csrName(csrReg), abiXName[rs1])
}

func daTypeIi(name string, pc uint, ins uint) string {
	return fmt.Sprintf("%s", name)
}

func daTypeIj(name string, pc uint, ins uint) string {
	csrReg, uimm, rd := decodeIb(ins)
	if csrReg == csrFRM {
		return fmt.Sprintf("fsrmi %s,%d", abiXName[rd], uimm)
	}
	if rd == 0 {
		return fmt.Sprintf("%s %s,%d", csrRemap1(name), csrName(csrReg), uimm)
	}
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], csrName(csrReg), uimm)
}

func daTypeIk(name string, pc uint, ins uint) string {
	rs2, rs1 := decodeId(ins)
	if rs2 == 0 && rs1 == 0 {
		return fmt.Sprintf("%s", name)
	}
	return fmt.Sprintf("%s %s,%s", name, abiXName[rs2], abiXName[rs1])
}

//-----------------------------------------------------------------------------
// Type U Decodes

func daTypeUa(name string, pc uint, ins uint) string {
	imm, rd := decodeU(ins)
	return fmt.Sprintf("%s %s,0x%x", name, abiXName[rd], uint(imm)&0xfffff)
}

//-----------------------------------------------------------------------------
// Type S Decodes

func daTypeSa(name string, pc uint, ins uint) string {
	imm, rs2, rs1 := decodeS(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiXName[rs2], imm, abiXName[rs1])
}

func daTypeSb(name string, pc uint, ins uint) string {
	imm, rs2, rs1 := decodeS(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiFName[rs2], imm, abiXName[rs1])
}

//-----------------------------------------------------------------------------
// Type R Decodes

func daTypeRa(name string, pc uint, ins uint) string {
	rs2, rs1, _, rd := decodeR(ins)
	if name == "sub" && rs1 == 0 {
		return fmt.Sprintf("neg %s,%s", abiXName[rd], abiXName[rs2])
	}
	return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], abiXName[rs1], abiXName[rs2])
}

func daTypeRb(name string, pc uint, ins uint) string {
	rs2, rs1, _, rd := decodeR(ins)
	if rs2 == 0 {
		return fmt.Sprintf("%s %s,(%s)", name, abiXName[rd], abiXName[rs1])
	}
	return fmt.Sprintf("%s %s,%s,(%s)", name, abiXName[rd], abiXName[rs2], abiXName[rs1])
}

func daTypeRc(name string, pc uint, ins uint) string {
	rs2, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s,%s", name, abiFName[rd], abiFName[rs1], abiFName[rs2])
}

func daTypeRd(name string, pc uint, ins uint) string {
	_, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiXName[rd], abiFName[rs1])
}

func daTypeRe(name string, pc uint, ins uint) string {
	_, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiFName[rd], abiXName[rs1])
}

func daTypeRf(name string, pc uint, ins uint) string {
	rs2, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], abiFName[rs1], abiFName[rs2])
}

func daTypeRh(name string, pc uint, ins uint) string {
	_, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiFName[rd], abiFName[rs1])
}

// fcvt rd = float, rs1 = float
// fcvt to {s,d} from {d,s}
func daTypeRi(name string, pc uint, ins uint) string {
	_, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiFName[rd], abiFName[rs1])
}

// fcvt rd = float, rs1 = int
// fcvt to {s,d} from {l,lu,w,wu}
func daTypeRj(name string, pc uint, ins uint) string {
	_, rs1, _, rd := decodeR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiFName[rd], abiXName[rs1])
}

// fcvt rd = int, rs1 = float
// fcvt to {l,lu,w,wu} from {d,s}
func daTypeRk(name string, pc uint, ins uint) string {
	_, rs1, rm, rd := decodeR(ins)
	if rm != frmDYN {
		return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], abiFName[rs1], rmName[rm])
	}
	return fmt.Sprintf("%s %s,%s", name, abiXName[rd], abiFName[rs1])
}

//-----------------------------------------------------------------------------
// Type R4 Decodes

func daTypeR4a(name string, pc uint, ins uint) string {
	rs3, rs2, rs1, _, rd := decodeR4(ins)
	return fmt.Sprintf("%s %s,%s,%s,%s", name, abiFName[rd], abiFName[rs1], abiFName[rs2], abiFName[rs3])
}

//-----------------------------------------------------------------------------
// Type B Decodes

func daTypeBa(name string, pc uint, ins uint) string {
	imm, rs2, rs1 := decodeB(ins)
	adr := int(pc) + imm

	if rs2 == 0 {
		switch name {
		case "bge", "beq", "bne", "blt":
			return fmt.Sprintf("%sz %s,%x", name, abiXName[rs1], adr)
		}
	}

	return fmt.Sprintf("%s %s,%s,%x", name, abiXName[rs1], abiXName[rs2], adr)
}

//-----------------------------------------------------------------------------
// Type J Decodes

func daTypeJa(name string, pc uint, ins uint) string {
	imm, rd := decodeJ(ins)
	if rd == 0 {
		return fmt.Sprintf("j %x", int(pc)+imm)
	}
	return fmt.Sprintf("%s %s,%x", name, abiXName[rd], int(pc)+imm)
}

//-----------------------------------------------------------------------------
// Type CI Decodes

func daNop(name string, pc uint, ins uint) string {
	return "nop"
}

func daTypeCIa(name string, pc uint, ins uint) string {
	imm, rd := decodeCIa(ins)
	return fmt.Sprintf("%s %s,%d", name, abiXName[rd], imm)
}

func daTypeCIb(name string, pc uint, ins uint) string {
	imm := decodeCIb(ins)
	return fmt.Sprintf("%s sp,sp,%d", name, imm)
}

func daTypeCIc(name string, pc uint, ins uint) string {
	imm, rd := decodeCIa(ins)
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], abiXName[rd], imm)
}

func daTypeCId(name string, pc uint, ins uint) string {
	uimm, rd := decodeCIc(ins)
	return fmt.Sprintf("%s %s,%s,0x%x", name, abiXName[rd], abiXName[rd], uimm)
}

func daTypeCIe(name string, pc uint, ins uint) string {
	uimm, rd := decodeCId(ins)
	return fmt.Sprintf("%s %s,%s,0x%x", name, abiXName[rd], abiXName[rd], uimm)
}

func daTypeCIf(name string, pc uint, ins uint) string {
	imm, rd := decodeCIe(ins)
	return fmt.Sprintf("%s %s,%s,%d", name, abiXName[rd], abiXName[rd], imm)
}

func daTypeCIg(name string, pc uint, ins uint) string {
	imm, rd := decodeCIf(ins)
	return fmt.Sprintf("%s %s,0x%x", name, abiXName[rd], imm)
}

func daTypeCIh(name string, pc uint, ins uint) string {
	uimm, rd := decodeCIg(ins)
	return fmt.Sprintf("%s %s,%d(sp)", name, abiXName[rd], uimm)
}

//-----------------------------------------------------------------------------
// Type CIW Decodes

func daTypeCIWa(name string, pc uint, ins uint) string {
	return "illegal"
}

func daTypeCIWb(name string, pc uint, ins uint) string {
	uimm, rd := decodeCIW(ins)
	return fmt.Sprintf("%s %s,sp,%d", name, abiXName[rd], uimm)
}

//-----------------------------------------------------------------------------
// Type CJ Decodes

func daTypeCJb(name string, pc uint, ins uint) string {
	imm := decodeCJ(ins)
	return fmt.Sprintf("%s %x", name, int(pc)+imm)
}

func daTypeCJc(name string, pc uint, ins uint) string {
	imm := decodeCJ(ins)
	return fmt.Sprintf("%s ra,%x", name, int(pc)+imm)
}

//-----------------------------------------------------------------------------
// Type CR Decodes

func daTypeCRa(name string, pc uint, ins uint) string {
	rd, rs := decodeCR(ins)
	return fmt.Sprintf("%s %s,%s", name, abiXName[rd], abiXName[rs])
}

func daTypeCRb(name string, pc uint, ins uint) string {
	rd, rs := decodeCR(ins)
	return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], abiXName[rd], abiXName[rs])
}

func daTypeCRc(name string, pc uint, ins uint) string {
	rd, rs := decodeCRa(ins)
	return fmt.Sprintf("%s %s,%s,%s", name, abiXName[rd], abiXName[rd], abiXName[rs])
}

func daTypeCRd(name string, pc uint, ins uint) string {
	rs1, _ := decodeCR(ins)
	if rs1 == 1 {
		return "ret"
	}
	return fmt.Sprintf("%s %s", name, abiXName[rs1])
}

func daTypeCRe(name string, pc uint, ins uint) string {
	rs1, _ := decodeCR(ins)
	return fmt.Sprintf("%s %s", name, abiXName[rs1])
}

//-----------------------------------------------------------------------------
// Type CS/CL Decodes

func daTypeCSa(name string, pc uint, ins uint) string {
	uimm, rs1, rs2 := decodeCS(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiXName[rs2], uimm, abiXName[rs1])
}

func daTypeCSb(name string, pc uint, ins uint) string {
	uimm, rs1, rs2 := decodeCSa(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiXName[rs2], uimm, abiXName[rs1])
}

func daTypeCSc(name string, pc uint, ins uint) string {
	uimm, rs1, rs2 := decodeCS(ins)
	return fmt.Sprintf("%s %s,%d(%s)", name, abiFName[rs2], uimm, abiXName[rs1])
}

//-----------------------------------------------------------------------------
// Type CSS Decodes

func daTypeCSSa(name string, pc uint, ins uint) string {
	uimm, rd := decodeCSSa(ins)
	return fmt.Sprintf("%s %s,%d(sp)", name, abiXName[rd], uimm)
}

func daTypeCSSb(name string, pc uint, ins uint) string {
	imm, rs2 := decodeCSSb(ins)
	return fmt.Sprintf("%s %s,%d(sp)", name, abiXName[rs2], imm)
}

func daTypeCSSc(name string, pc uint, ins uint) string {
	uimm, rs2 := decodeCSSc(ins)
	return fmt.Sprintf("%s %s,%d(sp)", name, abiXName[rs2], uimm)
}

//-----------------------------------------------------------------------------
// Type CB Decodes

func daTypeCBa(name string, pc uint, ins uint) string {
	imm, rs := decodeCB(ins)
	return fmt.Sprintf("%s %s,%x", name, abiXName[rs], int(pc)+imm)
}

//-----------------------------------------------------------------------------

// daInstruction returns the disassembly for a 16/32-bit instruction.
func (isa *ISA) daInstruction(pc uint, ins uint) string {
	im := isa.lookup(ins)
	if im != nil {
		return im.defn.da(im.name, pc, ins)
	}
	return "illegal"
}

//-----------------------------------------------------------------------------

// Disassembly returns the result of the disassembler call.
type Disassembly struct {
	Addr       uint // address
	AddrLength uint // address length in bits
	Ins        uint // instruction
	InsLength  uint // instruction length in bytes
	Assembly   string
}

func (da *Disassembly) String() string {
	addrFmt := fmt.Sprintf("%%0%dx", da.AddrLength>>2)
	addrStr := fmt.Sprintf(addrFmt, da.Addr)
	if da.InsLength == 2 {
		return fmt.Sprintf("%s: %04x     \t%s", addrStr, da.Ins, da.Assembly)
	}
	return fmt.Sprintf("%s: %08x \t%s", addrStr, da.Ins, da.Assembly)
}

// Disassemble a RISC-V instruction at the address.
func (isa *ISA) Disassemble(addr, ins uint) *Disassembly {
	da := &Disassembly{
		Addr:       addr,
		AddrLength: isa.mxlen,
	}
	if ins&3 == 3 {
		da.Ins = uint(uint32(ins))
		da.InsLength = 4
	} else {
		da.Ins = uint(uint16(ins))
		da.InsLength = 2
	}
	da.Assembly = isa.daInstruction(addr, ins)
	return da
}

//-----------------------------------------------------------------------------
