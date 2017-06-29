package main

import (
	"httpserver"
	"fmt"
	"os"
)

func main() {
	header := httpserver.Header{}

	header.Add("ha", "1")
	header.Add("ha", "3")
	header.Add("ab", "33")

	for key, value := range header {
		fmt.Print(key)
		fmt.Print(value)
		fmt.Println()
	}

	fp, err := os.Open("G:\\code\\go\\src\\github.com\\leecode\\servefile\\serverfile.go")


	fileInfo, err := fp.Stat()

	if err != nil {

	}

	fmt.Println(fileInfo.ModTime().Unix())
}