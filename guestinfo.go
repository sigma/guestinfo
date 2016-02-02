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

type getVersionRequest struct {
	registers
}

func newGetVersionRequest() getVersionRequest {
	return getVersionRequest{
		registers{
			_ax: backdoorMagic,
			_bx: 0xFFFFFFFF,
			_cx: backdoorCmdGetVersion,
			_dx: backdoorPort,
		},
	}
}

type getVersionResponse struct {
	hvMessage
}

func (r getVersionResponse) isVMware() bool {
	return r.hvMessage.bx() == backdoorMagic
}

func init() {
	virtualWorld = bdoor.HypervisorPortCheck()
}

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return virtualWorld
}
