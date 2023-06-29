//-----------------------------------------------------------------------------
/*

Example code for the rvda package.

*/
//-----------------------------------------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/deadsy/rvda"
)

//-----------------------------------------------------------------------------

func disassemble() error {
	isa, err := rvda.New(32, rvda.RV32gc)
	if err != nil {
		return err
	}
	// dump the ISA descriptor
	fmt.Printf("%s\n", isa)
	// disassemble random instructions
	for pc := 0; pc < 64; pc++ {
		ins := uint(rand.Uint32())
		da := isa.Disassemble(uint(pc), ins)
		fmt.Printf("%s\n", da)
	}
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
