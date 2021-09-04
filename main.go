package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// var l = -1
	// for i := 0; i < len(os.Args); i++ {
	// fmt.Println(os.Args[i])
	// if len(os.Args[i]) > 0 {
	// l++
	// }
	// }
	var l int = 0
	s := os.Args[1]
	var arr = strings.Split(s, " ")
	for _, w := range arr {
		if w != "" {
			l++
		}
	}
	fmt.Println(arr)
	// var l = len(arr)
	fmt.Print(l)
}
