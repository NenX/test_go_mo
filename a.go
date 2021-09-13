package my_test_mod

import (
	"fmt"
	"os"
	"strings"
)

func PrintSomeMessage() {
	a := os.Environ()
	for _, v := range a {
		flag := strings.Contains(v, "GOPATH")
		if flag {
			fmt.Printf("your go path:%#v\n", v)

		}
	}
}
