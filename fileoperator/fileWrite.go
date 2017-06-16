package fileoperator

import (
	"os"
	"fmt"
)

func WriteIn(filePath string, writeStr string) (err error) {
	filePointer, err := os.Create(filePath)
	if err != nil {
		return
	}

	defer filePointer.Close()
	_, err = filePointer.WriteString(writeStr)
	if err != nil {
		return
	}

	fmt.Println(writeStr)
	return
}