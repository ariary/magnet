package main

import (
	"flag"
	"fmt"
	"os"

	jen "github.com/dave/jennifer/jen"
)

func main() {
	var body, vars bool
	flag.BoolVar(&body, "body", false, "generate body magnet function to include in main")
	flag.BoolVar(&vars, "vars", false, "generate vars magnet function and import to include outside the main")
	flag.Parse()
	if body == vars {
		fmt.Println("Use -body flag OR -vars flag (XOR)")
		os.Exit(1)
	}

	if vars {
		//import
		fmt.Println("import \"github.com/ariary/magnet/pkg/magnet\"")
		//global vars
		fmt.Printf("%#v\n", jen.Null().Var().Id("FileList").Id("string"))
		fmt.Printf("%#v\n", jen.Null().Var().Id("Endpoint").Id("string"))
		fmt.Printf("%#v\n", jen.Null().Var().Id("Key").Id("string"))
	}

	if body {
		//magnet.Magnet(FileList, Endpoint, Key, false)
		fmt.Printf("%#v\n", jen.Id("magnet").Dot("Magnet").Call(jen.Id("FileList"), jen.Id("Endpoint"), jen.Id("Key"), jen.Id("false")))

	}
}
