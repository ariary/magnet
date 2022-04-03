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
	if len(flag.Args()) < 2 {
		fmt.Println("Please provide key and filename: lobfuscator [KEY] [FILENAME]")
		os.Exit(92)
	}
	key := flag.Args()[0]
	filename := flag.Args()[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
