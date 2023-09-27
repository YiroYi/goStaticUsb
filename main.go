package main

// #cgo CFLAGS:-I/opt/homebrew/Cellar/libusb/1.0.26/include/libusb-1.0
// #cgo LDFLAGS: -L. -lusb-1.0
// #include "/opt/homebrew/Cellar/libusb/1.0.26/include/libusb-1.0/libusb.h"
import "C"

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}
