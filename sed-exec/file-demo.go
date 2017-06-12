package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("sed", "-i", "--", "s/s/w/g", "file").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The output is %s\n", out)
}
