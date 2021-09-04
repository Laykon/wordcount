package main

import (
	"fmt"
	"os"
)

func main() {
	var l = len(os.Args) - 1
	fmt.Println(l)
}
