package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1]
	count := len(strings.Fields(str))
	fmt.Println(count)
}
