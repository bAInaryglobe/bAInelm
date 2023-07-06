package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("stack.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		link := scanner.Text()
		cmd := exec.Command("curl", "-L", "-O", "-J", "-C", "-", link)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
