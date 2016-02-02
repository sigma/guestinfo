package guestinfo

import "github.com/sigma/bdoor"

var (
	virtualWorld = false
)

const (
	backdoorPort  = 0x5658
	backdoorMagic = 0x564D5868

	backdoorCmdGetVersion = 10
)

func init() {
	virtualWorld = bdoor.HypervisorPortCheck()
}

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return virtualWorld
}
