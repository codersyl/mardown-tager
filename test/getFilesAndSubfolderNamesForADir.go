package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	dir, err := os.Getwd()
	errCheck(err)
	dir += "/"
	entries, err := os.ReadDir(dir)
	errCheck(err)
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		errCheck(err)
		infos = append(infos, info)
		printFileName(info)
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printFileName(f fs.FileInfo) {
	s := f.Name()
	if f.IsDir() {
		s += "/"
	}
	fmt.Println(s)
}
