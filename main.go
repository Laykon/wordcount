package main

import (
	"fmt"
	"os"
)

func main() {
	var l = -1
	for i := 0; i < len(os.Args); i++ {
		if len(os.Args[i]) > 0 {
			l++
		}
	}
	fmt.Print(l)
}
