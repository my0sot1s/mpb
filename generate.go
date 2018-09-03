package main

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	command = ""
)

func runBash(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func main() {
	fmt.Println(string(runBash("sh proto/xshop/product/build.sh", true)))
}
