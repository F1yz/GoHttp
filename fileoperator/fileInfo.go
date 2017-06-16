package fileoperator

import "os"

func GetFileInfo (file *os.File) (fileInfo os.FileInfo, err error) {
	fileInfo, err = file.Stat()
	return
}

func GetFileSize (file *os.File) (size int64, err error) {
	fileInfo, err := GetFileInfo(file)
	size = fileInfo.Size()
	return
}

