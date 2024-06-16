package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	curDirFS := os.DirFS("./")
	for i := 1; i < len(os.Args); i++ {
		list, _ := fs.Glob(curDirFS, os.Args[i])
		for i := 0; i < len(list); i++ {
			fmt.Printf("%v\n", list[i])
		}
	}

}
