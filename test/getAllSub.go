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
	helper(dir, 0)
}

func helper(dir string, level int) {
	entries, err := os.ReadDir(dir)
	errCheck(err)
	// infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		errCheck(err)
		// infos = append(infos, info)
		preSpace(level)
		s := dir + getFileName(info)
		fmt.Println(s)
		if info.IsDir() {
			helper(dir+getFileName(info), level+1)
		}
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

func getFileName(f fs.FileInfo) string {
	s := f.Name()
	if f.IsDir() {
		s += "/"
	}
	return s
}

func preSpace(level int) {
	for level > 0 {
		fmt.Printf("    ")
		level--
	}
}
