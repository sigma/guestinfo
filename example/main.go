package main

import (
	"fmt"

	"github.com/sigma/guestinfo"
)

func main() {
	if guestinfo.IsVirtualWorld() {
		fmt.Println("ok")
	} else {
		fmt.Println("ko")
	}
}
