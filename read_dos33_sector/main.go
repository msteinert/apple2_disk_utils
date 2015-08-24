package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikeakohn/apple2_disk_utils/apple2"
)

func main() {
	fmt.Println("read_dos33_sector: Apple II disk reader - Copyright 2015 Michael Kohn")

	if len(os.Args) != 4 {
		fmt.Println("Usage: " + os.Args[0] + " <binfile> <track> <sector>")
		os.Exit(0)
	}

	apple2_disk := new(apple2.Apple2Disk)

	if apple2_disk.Load(os.Args[1]) == false {
		os.Exit(1)
	}

	track, _ := strconv.Atoi(os.Args[2])
	sector, _ := strconv.Atoi(os.Args[3])

	apple2_disk.DumpSector(track, sector)
	fmt.Print("Is Free: ")
	if apple2_disk.IsSectorFree(track, sector) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
