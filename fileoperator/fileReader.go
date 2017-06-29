package fileoperator

import (
	"os"
)

func ReadAll (filePath string ) (readData []byte,err error) {
	filePointer, err := os.OpenFile(filePath, os.O_RDONLY, 444)
	if err != nil {
		return
	}

	defer filePointer.Close()

	byteSize, err := GetFileSize(filePointer)

	readData = make([]byte, byteSize)
	_, err = filePointer.Read(readData)
	return
}

// 这个方法似乎是有问题的
func FileExists(filePath string) bool  {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}

	return false
}
