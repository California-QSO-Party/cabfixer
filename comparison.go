package main

import (
	"bytes"
	"fmt"
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
			if len(lines[i]) != 0 {
				headerParts := bytes.Split(lines[i], []byte(":"))
				headers[trim(headerParts[0])] = trim(headerParts[1])
			}
		}
	}
	return CabFile{headers, qsoLines}, nil
}

func CabEqual(output CabFile, answer CabFile) error {
	errQsoLines := qsoLinesEqual(output.qsoLines, answer.qsoLines)
	if errQsoLines != nil {
		return errQsoLines
	}
	errHeaderEqual := headersEqual(output.headers, answer.headers)
	if errHeaderEqual != nil {
		return errHeaderEqual
	}
	return nil
}

func trim(a []byte) string {
	return string(bytes.Trim(a, WhiteSpaces))
}

func headersEqual(a map[string]string, b map[string]string) error {
	if len(a) != len(b) {
		return fmt.Errorf("The number of headers is different")
	}
	for k, v := range a {
		if v != b[k] {
			return fmt.Errorf("Content of header %v doesn't match", k)
		}
	}
	return nil
}

func qsoLinesEqual(a []string, b []string) error {
	if len(a) != len(b) {
		return fmt.Errorf("Number of qso lines doesn't match")
	}
	for i := 0; i < len(a); i++ {
		if qsoLineEqual(a[i], b[i]) == false {
			return fmt.Errorf("QSO lines %v, %v don't match", a[i], b[i])
		}
	}
	return nil
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
		a[i] = strings.Trim(a[i], WhiteSpaces)
	}

}
