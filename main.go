package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	curDirFS := os.DirFS("./")
	for i := 1; i < len(os.Args); i++ {
		list, _ := fs.Glob(curDirFS, os.Args[i])
		for i := 0; i < len(list); i++ {
			ProcessFile(list[i])

		}
	}

}

func ProcessFile(fileName string) {
	ext := filepath.Ext(fileName)
	newFileName := fileName[0 : len(fileName)-len(ext)]
	newFileName = newFileName + ".xcbr"
	fmt.Printf("%v - %v\n", fileName, newFileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(data, []byte("\n"))
	qsoLines := make([][]byte, 0)
	for i := 0; i < len(lines); i++ {
		if bytes.HasPrefix(lines[i], []byte("QSO:")) {
			qsoLines = append(qsoLines, lines[i])
		}
	}

	outputData := bytes.Join(qsoLines, []byte("\n"))
	err = os.WriteFile(newFileName, outputData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func longestLine(lines [][]byte) int {
	max := 0
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > max {
			max = len(lines[i])
		}
	}
	return max
}

func identifyTableColumns(lines [][]byte) {
	longestLine := longestLine(lines)
	charCount := make([]int, longestLine)
	for j := 0; j < len(lines); j++ {
		count := 0
		for i := 0; i < longestLine; i++ {
			if lines[i][j] != ' ' {
				count++
			}
		}
		charCount[j] = count
	}
}
