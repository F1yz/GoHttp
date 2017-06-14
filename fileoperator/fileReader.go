package fileoperator

import (
	"os"
	"io"
)

func ReadAll (filePath string ) (readData []byte,err error) {
	filePointer, err := os.OpenFile(filePath, os.O_RDONLY, 444)
	if err != nil {
		return
	}

	defer filePointer.Close()
	readBytes := make([]byte, 512)
	for  {
		len, err := filePointer.Read(readBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		readData = append(readData, readBytes[:len]...)
	}

	return
}
