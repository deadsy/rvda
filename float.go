//-----------------------------------------------------------------------------
/*

RISC-V Floating Point

*/
//-----------------------------------------------------------------------------

package rvda

//-----------------------------------------------------------------------------

// Rounding modes.
const (
	frmRNE = 0 // Round to Nearest, ties to Even
	frmRTZ = 1 // Round towards Zero
	frmRDN = 2 // Round Down (towards -inf)
	frmRUP = 3 // Round Up (towards +inf)
	frmRRM = 4 // Round to Nearest, ties to Max Magnitude
	frmDYN = 7 // Use the value in the FRM csr
)

// Rounding mode names.
var rmName = [8]string{
	"rne", "rtz", "rdn", "rup", "rrm", "rm5", "rm6", "dyn",
}

//-----------------------------------------------------------------------------
