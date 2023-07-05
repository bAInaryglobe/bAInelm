package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	input := "This is a sample input string"
	words := strings.Fields(input)
	if len(words) < 50 {
		cmd := exec.Command("python", "python_script.py", input)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		cmd := exec.Command("go", "run", "go_script.go", input)
		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
		}
	}
}