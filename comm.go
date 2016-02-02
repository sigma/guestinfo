// +build !debug

package guestinfo

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
