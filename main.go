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
	columnPos := identifyTableColumns(qsoLines)
	markedUpQsoLines := markUpQSOLines(qsoLines, columnPos)

	outputData := bytes.Join(markedUpQsoLines, []byte("\n"))
	err = os.WriteFile(newFileName, outputData, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func longestLine(qsoLines [][]byte) int {
	max := 0
	for i := 0; i < len(qsoLines); i++ {

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

		for i := 0; i < len(lines); i++ {
			if len(lines[i]) <= j {
				continue
			}

			if lines[i][j] != ' ' {
				count++
			}
		}
		charCount[j] = count
	}
	// Count positions of where each column starts.
	columnPos := []int{0}
	threshold := len(lines)/2 + 1
	for i := 1; i < len(charCount); i++ {
		if charCount[i] > threshold && charCount[i-1] == 0 {
			columnPos = append(columnPos, i)
		}
	}
	columnPos = append(columnPos, longestLine)
	return columnPos

}

func markUpQSOLines(qsoLines [][]byte, columnPos []int) [][]byte {
	markedUpQsoLines := make([][]byte, len(qsoLines))
	for i := 0; i < len(qsoLines); i++ {
		markedUpQsoLines[i] = append(markedUpQsoLines[i], '|')
		for j := 0; j < len(columnPos)-1; j++ {
			l := min(columnPos[j], len(qsoLines[i]))
			r := min(columnPos[j+1], len(qsoLines[i]))
			e := bytes.Trim(qsoLines[i][l:r], " \t\r\n")
			markedUpQsoLines[i] = append(markedUpQsoLines[i], e...)
			markedUpQsoLines[i] = append(markedUpQsoLines[i], '|')
		}
	}
	return markedUpQsoLines
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
