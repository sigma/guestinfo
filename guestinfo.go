package guestinfo

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	virtualWorld = false
)

const (
	backdoorPort  = 0x5658
	backdoorMagic = 0x564D5868

	backdoorCmdGetVersion = 10
)

func asmCommunicate(iax, ibx, icx, idx, idi, isi uint32) (oax, obx, ocx, odx, odi, osi uint32)

type registers struct {
	_ax, _bx, _cx, _dx, _di, _si uint32
}

func (r registers) ax() uint32 { return r._ax }

func (r registers) bx() uint32 { return r._bx }

func (r registers) cx() uint32 { return r._cx }

func (r registers) dx() uint32 { return r._dx }

func (r registers) di() uint32 { return r._di }

func (r registers) si() uint32 { return r._si }

type hvMessage interface {
	ax() uint32
	bx() uint32
	cx() uint32
	dx() uint32
	di() uint32
	si() uint32
}

func hvCommunicate(hvInput hvMessage) hvMessage {
	iax := hvInput.ax()
	ibx := hvInput.bx()
	icx := hvInput.cx()
	idx := hvInput.dx()
	idi := hvInput.di()
	isi := hvInput.si()

	oax, obx, ocx, odx, odi, osi := asmCommunicate(iax, ibx, icx, idx, idi, isi)
	return registers{_ax: oax, _bx: obx, _cx: ocx, _dx: odx, _di: odi, _si: osi}
}

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

func (r getVersionResponse) magic() uint32 {
	return r.hvMessage.bx()
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGSEGV, syscall.SIGILL)

	defer func() {
		signal.Reset(syscall.SIGSEGV, syscall.SIGILL)
	}()

	res := make(chan bool, 1)

	go func() {
		magic := getVersionResponse{hvCommunicate(newGetVersionRequest())}.magic()
		res <- (magic == backdoorMagic)
		c <- syscall.SIGILL
	}()

	<-c
	select {
	case r := <-res:
		virtualWorld = r
	default:
		virtualWorld = false
	}
}

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return virtualWorld
}
