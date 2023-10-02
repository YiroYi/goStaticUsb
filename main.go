package main

// #cgo CFLAGS:-I./libusb-1.0
// #cgo LDFLAGS: -L./lib -lusb-1.0
// #include "./libusb-1.0/libusb.h"

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

func main() {
	ctx := gousb.NewContext()
	// A Context manages all resources necessary for communicating with USB devices.
	//Through the Context users can iterate over available USB devices.
	defer ctx.Close()
	// The "defer" keyword will allow to call Close() when the main() is closed with this
	// we will avoid memory leaks if the connection is closed incorrectly or interrupted

	// Find the Zebra printer by its USB Vendor ID and Product ID
	dev, err := ctx.OpenDeviceWithVIDPID(gousb.ID(0x1504), gousb.ID(0x0037))
	fmt.Println(dev.Desc)
	if err != nil {
		log.Fatalf("Could not open device: %v", err)
	}
	defer dev.Close() // this method what will be do is to close the connection when the main function is ready

	// Set up the default configuration
	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("Could not get configuration: %v", err)
	}
	defer cfg.Close()

	// Claim an interface for communication
	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("Could not claim interface: %v", err)
	}
	defer intf.Close()

	// Send ZPL commands to the printer
	// zplCommand := "~Hi"
	zplCommand := "^XA^FO100,50^ADN,36,20^FDHello, Zebra Printer!^FS^XZ"
	outEndpoint, err := intf.OutEndpoint(2)

	if err != nil {
		log.Fatalf("Failed to send ZPL command: %v", err)
	}

	_, err = outEndpoint.Write([]byte(zplCommand))

	if err != nil {
		fmt.Println("Fatal Error")
	}

	fmt.Println("ZPL command sent successfully.")
}

/*
	What we are doing here is to tell go that he libusb library is here, then
	when we compile in the binary file we are including all this packages inside
	the executable, the binary run and find the library, however in other computer
	it is not working because by default it is looking the library in the device
  we need to run:
  export DYLD_LIBRARY_PATH=/Users/yiroyi/Desktop/goStaticUsb/lib
	In this way we are telling to find the library DYLD_Library in the above path
*/
