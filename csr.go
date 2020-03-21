//-----------------------------------------------------------------------------
/*

RISC-V CSR Definition

*/
//-----------------------------------------------------------------------------

package rvda

import "fmt"

//-----------------------------------------------------------------------------
// machine isa register

// ISA Extension Bitmap
const (
	misaExtA = (1 << iota) // Atomic extension
	misaExtB               // Tentatively reserved for Bit-Manipulation extension
	misaExtC               // Compressed extension
	misaExtD               // Double-precision floating-point extension
	misaExtE               // RV32E base ISA
	misaExtF               // Single-precision floating-point extension
	misaExtG               // Additional standard extensions present
	misaExtH               // Hypervisor extension
	misaExtI               // RV32I/64I/128I base ISA
	misaExtJ               // Tentatively reserved for Dynamically Translated Languages extension
	misaExtK               // Reserved
	misaExtL               // Tentatively reserved for Decimal Floating-Point extension
	misaExtM               // Integer Multiply/Divide extension
	misaExtN               // User-level interrupts supported
	misaExtO               // Reserved
	misaExtP               // Tentatively reserved for Packed-SIMD extension
	misaExtQ               // Quad-precision floating-point extension
	misaExtR               // Reserved
	misaExtS               // Supervisor mode implemented
	misaExtT               // Tentatively reserved for Transactional Memory extension
	misaExtU               // User mode implemented
	misaExtV               // Tentatively reserved for Vector extension
	misaExtW               // Reserved
	misaExtX               // Non-standard extensions present
	misaExtY               // Reserved
	misaExtZ               // Reserved
)

//-----------------------------------------------------------------------------

// Register numbers for specific CSRs.
const (
	csrFFLAGS = 0x001
	csrFRM    = 0x002
	csrFCSR   = 0x003
)

//-----------------------------------------------------------------------------

var csrLookup = map[uint]string{
	// User CSRs 0x000 - 0x0ff (read/write)
	0x000: "ustatus",
	0x001: "fflags",
	0x002: "frm",
	0x003: "fcsr",
	0x004: "uie",
	0x005: "utvec",
	0x040: "uscratch",
	0x041: "uepc",
	0x042: "ucause",
	0x043: "utval",
	0x044: "uip",
	// User CSRs 0xc00 - 0xc7f (read only)
	0xc00: "cycle",
	0xc01: "time",
	0xc02: "instret",
	0xc03: "hpmcounter3",
	0xc04: "hpmcounter4",
	0xc05: "hpmcounter5",
	0xc06: "hpmcounter6",
	0xc07: "hpmcounter7",
	0xc08: "hpmcounter8",
	0xc09: "hpmcounter9",
	0xc0a: "hpmcounter10",
	0xc0b: "hpmcounter11",
	0xc0c: "hpmcounter12",
	0xc0d: "hpmcounter13",
	0xc0e: "hpmcounter14",
	0xc0f: "hpmcounter15",
	0xc10: "hpmcounter16",
	0xc11: "hpmcounter17",
	0xc12: "hpmcounter18",
	0xc13: "hpmcounter19",
	0xc14: "hpmcounter20",
	0xc15: "hpmcounter21",
	0xc16: "hpmcounter22",
	0xc17: "hpmcounter23",
	0xc18: "hpmcounter24",
	0xc19: "hpmcounter25",
	0xc1a: "hpmcounter26",
	0xc1b: "hpmcounter27",
	0xc1c: "hpmcounter28",
	0xc1d: "hpmcounter29",
	0xc1e: "hpmcounter30",
	0xc1f: "hpmcounter31",
	// User CSRs 0xc80 - 0xcbf (read only)
	0xc80: "cycleh",
	0xc81: "timeh",
	0xc82: "instreth",
	0xc83: "hpmcounter3h",
	0xc84: "hpmcounter4h",
	0xc85: "hpmcounter5h",
	0xc86: "hpmcounter6h",
	0xc87: "hpmcounter7h",
	0xc88: "hpmcounter8h",
	0xc89: "hpmcounter9h",
	0xc8a: "hpmcounter10h",
	0xc8b: "hpmcounter11h",
	0xc8c: "hpmcounter12h",
	0xc8d: "hpmcounter13h",
	0xc8e: "hpmcounter14h",
	0xc8f: "hpmcounter15h",
	0xc90: "hpmcounter16h",
	0xc91: "hpmcounter17h",
	0xc92: "hpmcounter18h",
	0xc93: "hpmcounter19h",
	0xc94: "hpmcounter20h",
	0xc95: "hpmcounter21h",
	0xc96: "hpmcounter22h",
	0xc97: "hpmcounter23h",
	0xc98: "hpmcounter24h",
	0xc99: "hpmcounter25h",
	0xc9a: "hpmcounter26h",
	0xc9b: "hpmcounter27h",
	0xc9c: "hpmcounter28h",
	0xc9d: "hpmcounter29h",
	0xc9e: "hpmcounter30h",
	0xc9f: "hpmcounter31h",
	// Supervisor CSRs 0x100 - 0x1ff (read/write)
	0x100: "sstatus",
	0x102: "sedeleg",
	0x103: "sideleg",
	0x104: "sie",
	0x105: "stvec",
	0x106: "scounteren",
	0x140: "sscratch",
	0x141: "sepc",
	0x142: "scause",
	0x143: "stval",
	0x144: "sip",
	0x180: "satp",
	// Machine CSRs 0xf00 - 0xf7f (read only)
	0xf11: "mvendorid",
	0xf12: "marchid",
	0xf13: "mimpid",
	0xf14: "mhartid",
	// Machine CSRs 0x300 - 0x3ff (read/write)
	0x300: "mstatus",
	0x301: "misa",
	0x302: "medeleg",
	0x303: "mideleg",
	0x304: "mie",
	0x305: "mtvec",
	0x306: "mcounteren",
	0x320: "mucounteren",
	0x321: "mscounteren",
	0x322: "mhcounteren",
	0x323: "mhpmevent3",
	0x324: "mhpmevent4",
	0x325: "mhpmevent5",
	0x326: "mhpmevent6",
	0x327: "mhpmevent7",
	0x328: "mhpmevent8",
	0x329: "mhpmevent9",
	0x32a: "mhpmevent10",
	0x32b: "mhpmevent11",
	0x32c: "mhpmevent12",
	0x32d: "mhpmevent13",
	0x32e: "mhpmevent14",
	0x32f: "mhpmevent15",
	0x330: "mhpmevent16",
	0x331: "mhpmevent17",
	0x332: "mhpmevent18",
	0x333: "mhpmevent19",
	0x334: "mhpmevent20",
	0x335: "mhpmevent21",
	0x336: "mhpmevent22",
	0x337: "mhpmevent23",
	0x338: "mhpmevent24",
	0x339: "mhpmevent25",
	0x33a: "mhpmevent26",
	0x33b: "mhpmevent27",
	0x33c: "mhpmevent28",
	0x33d: "mhpmevent29",
	0x33e: "mhpmevent30",
	0x33f: "mhpmevent31",
	0x340: "mscratch",
	0x341: "mepc",
	0x342: "mcause",
	0x343: "mtval",
	0x344: "mip",
	0x380: "mbase",
	0x381: "mbound",
	0x382: "mibase",
	0x383: "mibound",
	0x384: "mdbase",
	0x385: "mdbound",
	0x3a0: "pmpcfg0",
	0x3a1: "pmpcfg1",
	0x3a2: "pmpcfg2",
	0x3a3: "pmpcfg3",
	0x3b0: "pmpaddr0",
	0x3b1: "pmpaddr1",
	0x3b2: "pmpaddr2",
	0x3b3: "pmpaddr3",
	0x3b4: "pmpaddr4",
	0x3b5: "pmpaddr5",
	0x3b6: "pmpaddr6",
	0x3b7: "pmpaddr7",
	0x3b8: "pmpaddr8",
	0x3b9: "pmpaddr9",
	0x3ba: "pmpaddr10",
	0x3bb: "pmpaddr11",
	0x3bc: "pmpaddr12",
	0x3bd: "pmpaddr13",
	0x3be: "pmpaddr14",
	0x3bf: "pmpaddr15",
	// Machine CSRs 0xb00 - 0xb7f (read/write)
	0xb00: "mcycle",
	0xb02: "minstret",
	0xb03: "mhpmcounter3",
	0xb04: "mhpmcounter4",
	0xb05: "mhpmcounter5",
	0xb06: "mhpmcounter6",
	0xb07: "mhpmcounter7",
	0xb08: "mhpmcounter8",
	0xb09: "mhpmcounter9",
	0xb0a: "mhpmcounter10",
	0xb0b: "mhpmcounter11",
	0xb0c: "mhpmcounter12",
	0xb0d: "mhpmcounter13",
	0xb0e: "mhpmcounter14",
	0xb0f: "mhpmcounter15",
	0xb10: "mhpmcounter16",
	0xb11: "mhpmcounter17",
	0xb12: "mhpmcounter18",
	0xb13: "mhpmcounter19",
	0xb14: "mhpmcounter20",
	0xb15: "mhpmcounter21",
	0xb16: "mhpmcounter22",
	0xb17: "mhpmcounter23",
	0xb18: "mhpmcounter24",
	0xb19: "mhpmcounter25",
	0xb1a: "mhpmcounter26",
	0xb1b: "mhpmcounter27",
	0xb1c: "mhpmcounter28",
	0xb1d: "mhpmcounter29",
	0xb1e: "mhpmcounter30",
	0xb1f: "mhpmcounter31",
	// Machine CSRs 0xb80 - 0xbbf (read/write)
	0xb80: "mcycleh",
	0xb82: "minstreth",
	0xb83: "mhpmcounter3h",
	0xb84: "mhpmcounter4h",
	0xb85: "mhpmcounter5h",
	0xb86: "mhpmcounter6h",
	0xb87: "mhpmcounter7h",
	0xb88: "mhpmcounter8h",
	0xb89: "mhpmcounter9h",
	0xb8a: "mhpmcounter10h",
	0xb8b: "mhpmcounter11h",
	0xb8c: "mhpmcounter12h",
	0xb8d: "mhpmcounter13h",
	0xb8e: "mhpmcounter14h",
	0xb8f: "mhpmcounter15h",
	0xb90: "mhpmcounter16h",
	0xb91: "mhpmcounter17h",
	0xb92: "mhpmcounter18h",
	0xb93: "mhpmcounter19h",
	0xb94: "mhpmcounter20h",
	0xb95: "mhpmcounter21h",
	0xb96: "mhpmcounter22h",
	0xb97: "mhpmcounter23h",
	0xb98: "mhpmcounter24h",
	0xb99: "mhpmcounter25h",
	0xb9a: "mhpmcounter26h",
	0xb9b: "mhpmcounter27h",
	0xb9c: "mhpmcounter28h",
	0xb9d: "mhpmcounter29h",
	0xb9e: "mhpmcounter30h",
	0xb9f: "mhpmcounter31h",
	// Machine Debug CSRs 0x7a0 - 0x7af (read/write)
	0x7a0: "tselect",
	0x7a1: "tdata1",
	0x7a2: "tdata2",
	0x7a3: "tdata3",
	// Machine Debug Mode Only CSRs 0x7b0 - 0x7bf (read/write)
	0x7b0: "dcsr",
	0x7b1: "dpc",
	0x7b2: "dscratch",
	// Hypervisor CSRs 0x200 - 0x2ff (read/write)
	0x200: "hstatus",
	0x202: "hedeleg",
	0x203: "hideleg",
	0x204: "hie",
	0x205: "htvec",
	0x240: "hscratch",
	0x241: "hepc",
	0x242: "hcause",
	0x243: "hbadaddr",
	0x244: "hip",
}

// csrName returns the name of a given CSR.
func csrName(reg uint) string {
	if name, ok := csrLookup[reg]; ok {
		return name
	}
	return fmt.Sprintf("0x%03x", reg)
}

//-----------------------------------------------------------------------------
