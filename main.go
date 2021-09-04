package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var l = len(strings.Split(strings.Trim(os.Args[1], " "), " "))
	fmt.Println(l)
}
