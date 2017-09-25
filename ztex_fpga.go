// Binary ztex_fpga interacts with FPGAs on ZTEX USB modules.
package main

import (
	"log"

	"github.com/aljumi/ztex"
	"github.com/google/gousb"

	getopt "github.com/pborman/getopt/v2"
)

var (
	helpFlag = getopt.BoolLong("help", 'h', "display this help and exit")
)

func main() {
	getopt.Parse()
	if *helpFlag {
		getopt.Usage()
		return
	}

	ctx := gousb.NewContext()
	defer ctx.Close()

	d, err := ztex.OpenDevice(ctx)
	if err != nil {
		log.Fatalf("ztex.OpenDevice: %v", err)
	}
	defer d.Close()

	// TODO: Code this.
}
