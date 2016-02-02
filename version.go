package guestinfo

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
