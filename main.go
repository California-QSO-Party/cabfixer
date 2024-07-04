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
	identifyTableColumns(qsoLines)

	outputData := bytes.Join(qsoLines, []byte("\n"))
	err = os.WriteFile(newFileName, outputData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func longestLine(qsoLines [][]byte) int {
	max := 0
	for i := 0; i < len(qsoLines); i++ {
		fmt.Printf("longestLine index: %d\n", i)
		if len(qsoLines[i]) > max {
			max = len(qsoLines[i])
		}
	}
	return max
}

func identifyTableColumns(lines [][]byte) []int {

	longestLine := longestLine(lines)
	charCount := make([]int, longestLine)
	for j := 0; j < longestLine; j++ {

		count := 0
		fmt.Printf("identifyTableColumns index j: %d\n", j)
		for i := 0; i < len(lines); i++ {
			if len(lines[i]) <= j {
				continue
			}
			fmt.Printf("identifyTableColumns index i: %d\n", i)
			fmt.Printf("identifyTableColumns content i: %s\n", lines[i])
			if lines[i][j] != ' ' {
				count++
			}
		}
		charCount[j] = count
	}
	fmt.Printf("%v\n", charCount)
	// Count positions of where each column starts.
	columnPos := []int{0}
	threshold := len(lines)/2 + 1
	for i := 1; i < len(charCount); i++ {
		if charCount[i] > threshold && charCount[i-1] == 0 {
			columnPos = append(columnPos, i)
		}
	}
	fmt.Printf("Start of columns ind: %v\n", columnPos)
	return columnPos

}
