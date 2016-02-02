package guestinfo

import "github.com/sigma/bdoor"

var (
	virtualWorld = false
)

func init() {
	virtualWorld = bdoor.HypervisorPortCheck()
}

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return virtualWorld
}
