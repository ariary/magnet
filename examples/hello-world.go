package main

import (
	"fmt"

	"github.com/ariary/magnet/pkg/magnet"
)

var FileList, Key, Endpoint string

//go:generate go run github.com/ariary/magnet/cmd/magnetgentool -vars

func main() {
	fmt.Println("hello world")

	magnet.Magnet(FileList, Endpoint, Key, false)

	//go:generate go run github.com/ariary/magnet/cmd/magnetgentool -body
}
