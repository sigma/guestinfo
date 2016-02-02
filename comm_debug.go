// +build debug

package guestinfo

import "fmt"

func hvCommunicate(hvInput hvMessage) hvMessage {
	iax := hvInput.ax()
	ibx := hvInput.bx()
	icx := hvInput.cx()
	idx := hvInput.dx()
	idi := hvInput.di()
	isi := hvInput.si()

	fmt.Println("hvCommunicate")
	fmt.Println("Before:")
	fmt.Println("  ax =", iax)
	fmt.Println("  bx =", ibx)
	fmt.Println("  cx =", icx)
	fmt.Println("  dx =", idx)
	fmt.Println("  di =", idi)
	fmt.Println("  si =", isi)

	oax, obx, ocx, odx, odi, osi := asmCommunicate(iax, ibx, icx, idx, idi, isi)

	fmt.Println("After:")
	fmt.Println("  ax =", oax)
	fmt.Println("  bx =", obx)
	fmt.Println("  cx =", ocx)
	fmt.Println("  dx =", odx)
	fmt.Println("  di =", odi)
	fmt.Println("  si =", osi)

	return registers{_ax: oax, _bx: obx, _cx: ocx, _dx: odx, _di: odi, _si: osi}
}
