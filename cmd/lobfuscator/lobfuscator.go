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
	flag.BoolVar(&reverse, "d", false, "decode and decrypt string previously encrypted with lobfuscator")
	var keyfile string
	flag.StringVar(&keyfile, "kf", "", "specify file containig keys in the right order to decrypt")
	flag.Parse()

	if keyfile != "" && !reverse {
		fmt.Println("Use -kf with -d")
		os.Exit(92)
	}
	var key string
	if len(flag.Args()) < 1 && reverse && keyfile == "" {
		fmt.Println("Please provide key: lobfuscator -d [KEY] or lobfuscator -d -kf [KEYS_FILE]")
		os.Exit(92)
	}
	if len(flag.Args()) >= 1 {
		key = flag.Args()[0]
	}

	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)

	if reverse {
		if keyfile != "" {
			scanAndDecryptWithMultipleKeys(scanner, keyfile)
		} else {
			scanAndDecrypt(scanner, key)
		}

	} else if key != "" {
		scanAndEncryptWithKey(scanner, key)
	} else {
		scanAndEncryptByGenerateKey(scanner)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//browse all the input and decrypt them
func scanAndDecrypt(scanner *bufio.Scanner, key string) {
	for scanner.Scan() {
		var encTargetFile string
		//base64 decode
		data := encryption.Decode(scanner.Text())
		encTargetFile = encryption.Xor(string(data), key)

		fmt.Println(encTargetFile)
	}
}

//browse all the input and decrypt them using different keys
func scanAndDecryptWithMultipleKeys(scanner *bufio.Scanner, keyfile string) {
	file, err := os.Open(keyfile)
	if err != nil {
		log.Fatal("Failed opening keys file:", err)
	}
	defer file.Close()

	keyScanner := bufio.NewScanner(file)

	for scanner.Scan() {
		keyScanner.Scan()
		key := keyScanner.Text()
		var encTargetFile string
		//base64 decode
		data := encryption.Decode(scanner.Text())
		encTargetFile = encryption.Xor(string(data), key)

		fmt.Println(encTargetFile)
	}

	if err := keyScanner.Err(); err != nil {
		fmt.Println(err)
	}
}

//browse all the input and encrypt them with the key provided
func scanAndEncryptWithKey(scanner *bufio.Scanner, key string) {
	for scanner.Scan() {
		encTargetFile := encryption.Encode([]byte(encryption.Xor(scanner.Text(), key)))
		fmt.Println(encTargetFile)
	}
}

//browse all the input and encrypt them with the key provided
func scanAndEncryptByGenerateKey(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		key := encryption.GenerateRandomStringWithLength(len(text))
		encTargetFile := encryption.Encode([]byte(encryption.Xor(text, key)))
		fmt.Fprintf(os.Stderr, "key:%s\n", key)
		fmt.Println(encTargetFile)
	}
}
