package main

import (
	"fmt"
	"os"
	"path"
)

func countGoFiles(folder string) int {
	files, err := os.ReadDir(folder)
	if err != nil {
		return 0
	}

	count := 0
	for _, f := range files {
		if path.Ext(f.Name()) == ".go" {
			count++
		}
	}
	//fmt.Printf("%+v", f.Name())
	return count
}

func main() {
	folder := "/Volumes/ARCHIVE/PRIVATE/WS/github/lets-go/snippets"
	count := countGoFiles(folder)
	fmt.Printf("Count of files: %d\n", count)

}
