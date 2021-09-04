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
	s := os.Args[1]
	var l = len(strings.Split(s, ""))
	fmt.Print(l)
}
