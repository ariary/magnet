package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ariary/magnet/pkg/xor"
)

func main() {

	// deobfuscate list
	var reverse bool
	flag.BoolVar(&reverse, "d", false, "")
	flag.Parse()
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
			data, err := base64.StdEncoding.DecodeString(scanner.Text())
			if err != nil {
				panic(err)
			}
			encTargetFile = xor.EncryptDecrypt(string(data), key)
		} else {
			encTargetFile = base64.StdEncoding.EncodeToString([]byte(xor.EncryptDecrypt(scanner.Text(), key)))
		}

		fmt.Println(encTargetFile)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
