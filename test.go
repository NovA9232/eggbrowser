package main

import (
	"strings"
	"fmt"
)

func main() {
	f := strings.Split("/home/josh/egg/", "/")
	fmt.Println(strings.Join(f[:len(f)-2], "/")+"/")
}
