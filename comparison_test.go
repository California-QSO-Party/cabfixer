package main

import (
	"fmt"
	"testing"
)

func TestCabRead(t *testing.T) {

	got, _ := CabRead("testcases/AF6HO-20231017-191821-471.xcbr")

	fmt.Printf("headers:\n")
	for k, v := range got.headers {
		fmt.Printf("%v: %v\n", k, v)
	}
	fmt.Printf("QSO lines:\n")
	for _, v := range got.qsoLines {
		fmt.Printf("%v\n", v)
	}

}
