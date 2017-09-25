// Binary ztex_fpga interacts with FPGAs on ZTEX USB modules.
package main

import (
	"fmt"
	"log"

	"github.com/aljumi/ztex"
	"github.com/google/gousb"

	getopt "github.com/pborman/getopt/v2"
)

var (
	resetFlag  = getopt.BoolLong("reset", 'r', "reset FPGA")
	statusFlag = getopt.BoolLong("status", 's', "output FPGA status")

	helpFlag = getopt.BoolLong("help", 'h', "display this help and exit")
)

func reset(d *ztex.Device) error {
	if err := d.ResetFPGA(); err != nil {
		return fmt.Errorf("(*ztex.Device).ResetFPGA: %v", err)
	}
	return nil
}

func printStatus(d *ztex.Device) error {
	fst, err := d.FPGAStatus()
	if err != nil {
		return fmt.Errorf("(*ztex.Device).FPGAStatus: %v", err)
	}

	fmt.Printf("FPGA Status:\n")
	fmt.Printf("  Configured: %v\n", fst.FPGAConfigured)
	fmt.Printf("  Checksum: %v\n", fst.FPGAChecksum)
	fmt.Printf("  Transferred: %v\n", fst.FPGATransferred)
	fmt.Printf("  Init: %v\n", fst.FPGAInit)
	fmt.Printf("  Result: %v\n", fst.FPGAResult)
	fmt.Printf("  Swapped: %v\n", fst.FPGASwapped)

	return nil
}

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

	if *resetFlag {
		if err := reset(d); err != nil {
			log.Fatalf("reset: %v", err)
		}
	}

	if *statusFlag {
		if err := printStatus(d); err != nil {
			log.Fatalf("printStatus: %v", err)
		}
	}
}
