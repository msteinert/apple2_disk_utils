package main

import (
	"fmt"
	"os"

	"github.com/mikeakohn/apple2_disk_utils/apple2"
)

func main() {
	fmt.Println("init_dos33_sector: Apple II disk reader - Copyright 2015 Michael Kohn")

	if len(os.Args) != 4 {
		fmt.Println("Usage: " + os.Args[0] + " <binfile> <dos33.img> <hello_prog>")
		os.Exit(0)
	}

	apple2_disk := new(apple2.Apple2Disk)
	apple2_disk.Init()
	apple2_disk.AddDos(os.Args[2])
	apple2_disk.AddFile(os.Args[3], "HELLO", 0x9100)
	apple2_disk.Save(os.Args[1])
}
