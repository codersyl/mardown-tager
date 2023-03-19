package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// func used in main
// len()
// os.Open()
// log.Fatal()
// fmt.Println()
// getTagsFromOneFile()

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tags := getTagsFromOneFile(f)
	fmt.Println(tags)
	return
}

// func used in getTagsFromOneFile()
// os.Read()
// len()
// append()
func getTagsFromOneFile(f *os.File) (tags []string) { // "os" "log" ""
	data := make([]byte, 2048)
	curTag := make([]byte, 0, 500)
	lastRemain := false
	for { // one read and analyse per loop
		// 读文件，并判断是否到尾
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF { // "io" used here
				log.Fatal(err)
			}
			if len(curTag) > 0 {
				// fmt.Println("end is", len(tags))
				tags = append(tags, string(curTag))
			}
			break
		}

		// 寻找 tag
		i := 0
		if lastRemain { // 上个 Buffer 有标签没结束
			for i < count && data[i] != ' ' && data[i] != '#' && data[i] != '\n' && data[i] != '\r' { // tag停止的字符 ' ' '#' '\n'
				i++
			}
			curTag = append(curTag, data[:i]...)
			if i != count {
				lastRemain = false
				tags = append(tags, string(curTag))
				curTag = curTag[:0]
			}
		}

		for i < count { // try to find one tag per loop
			if data[i] != '#' {
				i++
				continue
			}
			r := i + 1 // end flag for tag
			for r < count {
				if data[r] != ' ' && data[r] != '#' && data[r] != '\n' && data[r] != '\r' {
					r++
				} else {
					break
				}
			}

			if r == count { // 本次 Buffer 未能结束标签
				curTag = append(curTag, data[i+1:count]...)
				lastRemain = true
				i = count
				continue
			}
			if r == i+1 { // 长度为 0 的标签，不记录
				i++
				continue
			}
			// 一个完整标签
			tags = append(tags, string(data[i+1:r]))
			i = r
		}
	}
	return tags
}
