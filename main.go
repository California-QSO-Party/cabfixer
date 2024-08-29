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

const endOfLog = "END-OF-LOG:"

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
	headerLines := make([][]byte, 0)
	hasXqso := false
	for i := 0; i < len(lines); i++ {
		if bytes.HasPrefix(lines[i], []byte("QSO:")) {
			qsoLines = append(qsoLines, append([]byte("  "), lines[i]...))
		} else if bytes.HasPrefix(lines[i], []byte("X-QSO:")) {
			hasXqso = true
			qsoLines = append(qsoLines, lines[i])
		} else {
			if len(bytes.Trim(lines[i], " \r\n\t")) != 0 {
				headerLines = append(headerLines, bytes.Trim(lines[i], "\r"))
			}
		}
	}
	if hasXqso == false {
		stripLeadingSpaces(qsoLines)
	}

	footerLines := make([][]byte, 0)
	if string(bytes.Trim(headerLines[len(headerLines)-1], " \r\n")) == endOfLog {
		headerLines = headerLines[0 : len(headerLines)-1]
		footerLines = append(footerLines, []byte(endOfLog))
	}
	headerLines = append(headerLines, []byte("X-CBR: 0.1"))
	columnPos := identifyTableColumns(qsoLines)
	markedUpQsoLines := markUpQSOLines(qsoLines, columnPos)
	outputDataLines := make([][]byte, 0)
	outputDataLines = append(outputDataLines, headerLines...)
	outputDataLines = append(outputDataLines, markedUpQsoLines...)
	outputDataLines = append(outputDataLines, footerLines...)

	outputData := bytes.Join(outputDataLines, []byte(eol))
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
		if charCount[i] > threshold {
			var j int
			for j = i; j >= 0; j-- {
				if j == 0 || charCount[j-1] == 0 {
					break
				}
			}
			if len(columnPos) == 0 || columnPos[len(columnPos)-1] < j {
				columnPos = append(columnPos, j)
			}
		}
		/*
			if charCount[i] > threshold && charCount[i-1] == 0 {
				columnPos = append(columnPos, i)
			}
		*/

	}
	columnPos = append(columnPos, longestLine)
	return columnPos

}

func markUpQSOLines(qsoLines [][]byte, columnPos []int) [][]byte {
	markedUpQsoLines := make([][]byte, len(qsoLines))
	for i := 0; i < len(qsoLines); i++ {
		for j := 0; j < len(columnPos)-1; j++ {
			l := min(columnPos[j], len(qsoLines[i]))
			r := min(columnPos[j+1], len(qsoLines[i]))
			e := bytes.Trim(qsoLines[i][l:r], "\r\n")
			markedUpQsoLines[i] = append(markedUpQsoLines[i], e...)
			if j < len(columnPos)-2 {
				markedUpQsoLines[i] = append(markedUpQsoLines[i], '|')
			}
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

func stripLeadingSpaces(lines [][]byte) {
	for i := 0; i < len(lines); i++ {
		lines[i] = bytes.TrimLeft(lines[i], " ")
	}
}
