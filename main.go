package main

import (
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
	err = os.WriteFile(newFileName, data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
