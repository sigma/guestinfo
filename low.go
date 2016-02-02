package guestinfo

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
