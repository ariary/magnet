package magnet

import (
	"fmt"
	"os"
)

// isDirectory determines if a file represented
// by the File pointer is a dire or not
func isDirectory(data *os.File) (bool, error) {
	fileInfo, err := data.Stat()
	if err != nil {
		return false, err
	}
	if fileInfo.IsDir() {
		fmt.Println("Directoryyyyyyyyy")
	}

	return fileInfo.IsDir(), err
}
