package main

import (
	"fmt"

	"github.com/sigma/guestinfo"
)

func main() {
	if guestinfo.IsVirtualWorld() {
		fmt.Println("Runnning in", guestinfo.VMwareProduct().Name())
	} else {
		fmt.Println("Not running in a VMware Virtual Machine")
	}
}
