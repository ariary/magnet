package magnet

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type MagnetSender interface {
	//SendFiles: send file to target
	SendFiles(fileList []string, cEndpoint string, debug bool)
	//Name: return MagnetSender name
	Name() string
}

// InitMagnetSender: return MagnetSender struct regarding the method parameter
func InitMagnetSender(method string) (sender MagnetSender) {
	switch {
	case method == "http":
		sender = MagnetHTTPSender{}
	case method == "tcp":
		sender = MagnetTCPSender{}
	}
	return sender
}

// MagnetHTTPSender: Sender that uses HTTP to send the files
type MagnetHTTPSender struct {
	Endpoint string
	Key      string
}

// Sendfiles: send all files of a list using HTTP multipart-form data request to a specified endpoint
func (sender MagnetHTTPSender) SendFiles(fileList []string, cEndpoint string, debug bool) {
	client := &http.Client{}
	for i := 0; i < len(fileList); i++ {
		err := SendFileHTTP(client, cEndpoint, fileList[i])
		if err != nil && debug {
			fmt.Println(fileList[i])
			fmt.Println("error:", err)
		}
	}
}

func (sender MagnetHTTPSender) Name() string {
	return "HTTP magnet sender"
}

// SendFileHTTP: use HTTP multipart form data request to upload a specific file on specified endpoint
func SendFileHTTP(client *http.Client, endpoint string, filename string) (err error) {
	// file exists?
	data, err := os.Open(os.ExpandEnv(filename))
	if err != nil {
		//surely file does not exist -> stop
		return err
	}

	var content io.Reader
	// file is dir?
	if isDir, err := isDirectory(data); err != nil {
		return err
	} else if isDir {
		// change data to be a tar archive
		var buf bytes.Buffer
		if err = DirToTar(data.Name(), &buf); err != nil {
			return err
		}
		content = &buf
		filename += ".tar.gz"
	} else { // is file
		content = data
	}

	// send file
	values := map[string]io.Reader{
		"file": content,
	}

	err = UploadHTTP(client, endpoint, values, filename)
	if err != nil {
		return err
	}

	return err
}

// UploadHTTP: construct upload request and submit it. The filename is used for the directory case when we upload an tar gz archive
// (As it is passed as a buffer we can't access the filename)
func UploadHTTP(client *http.Client, url string, values map[string]io.Reader, filename string) (err error) {
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
		} else { //surely a tar archive (pass as bytes.Buffer)
			if fw, err = w.CreateFormFile(key, strconv.Itoa(time.Now().Nanosecond())+"-"+filename); err != nil { //weird name to avoid collision
				return err
			}
			// Add other fields
			// if fw, err = w.CreateFormField(key); err != nil {
			// 	return err
			// }
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

// MagnetTCPSender: Sender that uses raw tcp to send the files
type MagnetTCPSender struct {
}

// Sendfiles: send all files of a list using TCP raw socket
func (sender MagnetTCPSender) SendFiles(fileList []string, cEndpoint string, debug bool) {
	for i := 0; i < len(fileList); i++ {
		err := SendFileTCP(cEndpoint, fileList[i])
		if err != nil && debug {
			fmt.Println(fileList[i])
			fmt.Println("error:", err)
		}
		time.Sleep(2 * time.Second)
	}
}

func (sender MagnetTCPSender) Name() string {
	return "TCP raw magnet sender"
}

func SendFileTCP(endpoint string, filename string) (err error) {
	file, err := os.Open(os.ExpandEnv(filename))
	if err != nil {
		return err
	}
	defer file.Close() // make sure to close the file even if we panic.
	connection, err := net.Dial("tcp", endpoint)
	if err != nil {
		return err
	}
	defer connection.Close()
	_, err = io.Copy(connection, file)

	return err
}
