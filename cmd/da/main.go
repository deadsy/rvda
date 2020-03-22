//-----------------------------------------------------------------------------
/*

Example code for the rvda package.

Open a RISC-V ELF file and disassemble the code.

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"

	"github.com/deadsy/rvda"
)

//-----------------------------------------------------------------------------

func disassemble() error {
	isa, err := rvda.New(32, rvda.RV32gc)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", isa)
	return nil
}

//-----------------------------------------------------------------------------

func main() {
	err := disassemble()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

//-----------------------------------------------------------------------------
