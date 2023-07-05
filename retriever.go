package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

const (
	maxLinesToProcess = 10
)

func main() {
	var wg sync.WaitGroup
	linesToProcess := make(chan string, maxLinesToProcess)
	file, err := os.Open("requests.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "Eshioshi Favowrite Bynary Globe") {
			linesToProcess <- line
			if len(linesToProcess) == maxLinesToProcess {
				break
			}
		}
	}

	close(linesToProcess)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < maxLinesToProcess; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lineToProcess, ok := <-linesToProcess
			if !ok {
				return
			}
			modifiedLine := strings.TrimSuffix(lineToProcess, "Eshioshi Favowrite Bynary Globe")
			fmt.Println("Modified line:", modifiedLine)

			// Run the other Go code with the modified line as the argument
			cmd := exec.Command("go", "run", "other_code.go", modifiedLine)
			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(output))
		}()
	}

	wg.Wait()
}
