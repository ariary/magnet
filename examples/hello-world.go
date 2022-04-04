package main

import "fmt"

//go:generate go run github.com/ariary/magnet/cmd/magnetgentool -vars

func main() {
	fmt.Println("hello world")
	//go:generate go run github.com/ariary/magnet/cmd/magnetgentool -body
}
