/**
 * should handle file upload & download (gzip & range download support)
 * 
 **/

package http

import (
	"fmt"
	"io"
	"os"
	"mime"
	"net/http" // @TODO: replace with our own http package later.
	"path"
)

const defaultMimeTypeForDownload = "application/octet-stream"
const maxBuffferSize = 4096	// @TODO: read from yaml

func min(x, y int64) int64 {
	if x > y {
		return y
	}

	return x
}

func ServeFile(filepath string, response http.ResponseWriter, req *http.Request) {
	fmt.Println(fmt.Sprintf("filePath : %s", filepath))
	f, err := os.Open(filepath)
	if err != nil {
		// 404 response
		return
	}

	defer f.Close()

	// get file info
	fileInfo, err := f.Stat()
	if err != nil {
		// 505 error?
		return
	}

	response.Header().Set("Last-Modified", fileInfo.ModTime().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	// parse mimeType
	if mimeType := mime.TypeByExtension(path.Ext(filepath)); mimeType != "" {
		// set response header.
		response.Header().Set("Content-Type", mimeType)
	} else {
		response.Header().Set("Content-Type", defaultMimeTypeForDownload)
	}

	outputWriter := response.(io.Writer)

	//@TODO: gzip handling? content-range header parse & partial contents read & response.


	// resposne data.
	dataBuf := make([]byte, min(maxBuffferSize, fileInfo.Size()))

	// read file data to dataBuf
	for {
		n, err := f.Read(dataBuf)

		if err != nil {
			break
		}

		outputWriter.Write(dataBuf[0:n])
	}
}
