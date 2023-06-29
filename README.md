[![Go Report Card](https://goreportcard.com/badge/github.com/deadsy/rvda)](https://goreportcard.com/report/github.com/deadsy/rvda)
[![GoDoc](https://godoc.org/github.com/deadsy/rvda?status.svg)](https://godoc.org/github.com/deadsy/rvda)

# rvda
RISC-V Disassembler

## usage

### code

```
isa, _ := rvda.New(32, rvda.RV32gc)
fmt.Printf("isa: %s\n", isa)

addr := uint(0xdeadbeef)
ins := uint(0x483f8297)
da := isa.Disassemble(addr, ins)

fmt.Printf("decode: %#v\n", da)
fmt.Printf("string: %s\n", da)
```

### output

```
isa: RV32 ext "acdfim"
decode: &rvda.Disassembly{Addr:0xdeadbeef, AddrLength:0x20, Ins:0x483f8297, InsLength:0x4, Assembly:"auipc t0,0x483f8"}
string: deadbeef: 483f8297      auipc t0,0x483f8
```


