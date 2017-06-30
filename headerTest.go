package main

import (
	"httpserver"
	"fmt"
	"os"
	"strings"
)

func main() {

	ranges := []string {
		"bytes 100-1000, 1001-1010",
		"bytes 0-",
	}

	for _, s := range ranges {
		info := strings.Split(s, " ")

		startEndRange := info[1]

		rangeInfo := strings.TrimRight(startEndRange, ",")

		rangeArra := strings.Split(rangeInfo, "-")

		fmt.Println(rangeArra)
	}
	return

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