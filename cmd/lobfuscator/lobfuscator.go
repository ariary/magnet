package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

func main() {

	// deobfuscate list
	var reverse bool
	flag.BoolVar(&reverse, "d", false, "")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("Please provide key: lobfuscator [KEY]")
		os.Exit(92)
	}
	key := flag.Args()[0]
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		var encTargetFile string
		if reverse {
			//base64 decode
			data, err := encryption.Decode(scanner.Text())
			if err != nil {
				panic(err)
			}
			encTargetFile = encryption.Xor(string(data), key)
		} else {
			encTargetFile = encryption.Encode([]byte(encryption.Xor(scanner.Text(), key)))
		}

		fmt.Println(encTargetFile)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
