package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	lie string
}

type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex> ex26 wordfilepath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(sord, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("-----------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", linInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("-----------")
		fmt.Println()
	}
}