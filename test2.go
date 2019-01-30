package main

import (
	"fmt"
	"os/exec"
)


func main() {
	out, err := exec.Command("/bin/bash", "-c", "xdg-open /home/josh/test.jpg").Output()
	fmt.Println(out, "\n", err)
}
