package main

import (
	"bytes"
	"os"
	"strings"
)

const WhiteSpaces = " \t\n\r"

type CabFile struct {
	headers  map[string]string
	qsoLines []string
}

func CabRead(filename string) (CabFile, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return CabFile{}, err
	}
	lines := bytes.Split(content, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		lines[i] = bytes.Trim(lines[i], WhiteSpaces)
	}
	headers := make(map[string]string)
	qsoLines := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		if bytes.HasPrefix(lines[i], []byte("QSO:")) {
			qsoLines = append(qsoLines, string(lines[i]))
		} else {
			headerParts := bytes.Split(lines[i], []byte(":"))
			headers[trim(headerParts[0])] = trim(headerParts[1])
		}
	}
	return CabFile{headers, qsoLines}, nil
}

func CabEqual(output CabFile, answer CabFile) bool {

}

func trim(a []byte) string {
	return string(bytes.Trim(a, WhiteSpaces))
}

func headersEqual(a map[string]string, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

func qsoLinesEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if qsoLineEqual(a[i], b[i]) == false {
			return false
		}
	}
	return true
}

func qsoLineEqual(a string, b string) bool {
	aCells := strings.Split(a, "|")
	bCells := strings.Split(b, "|")
	trimCellSpaces(aCells)
	trimCellSpaces(bCells)
	if len(aCells) != len(bCells) {
		return false
	}
	for k, v := range aCells {
		if v != bCells[k] {
			return false
		}
	}
	return true
}

func trimCellSpaces(a []string) {
	for i := 0; i < len(a); i++ {
		a[i] = strings.Trim(a, WhiteSpaces)
	}

}
