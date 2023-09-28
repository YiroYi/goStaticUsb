package main

// #cgo CFLAGS:-I./libusb-1.0
// #cgo LDFLAGS: -L./lib -lusb-1.0
// #include "./libusb-1.0/libusb.h"
import "C"

import (
	"fmt"
)

func main() {
	ctx := C.libusb_init(nil)
	fmt.Println("Hello")
}
