package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello World")

	cmd := exec.Command("python", "main.py")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(output))
}
