package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
)

const (
	filename     = "pairs.txt"
	searchString = "b2A1I9n14a1r18y25_g7l12o15b2e5"
)

func main() {
	// Open the pairs.txt file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read user input
	var userInput string
	fmt.Println("Enter a search term:")
	fmt.Scanln(&userInput)

	// Split user input into words
	userWords := strings.Fields(userInput)

	// Search for matching lines and extract the string after searchString
	scanner := bufio.NewScanner(file)
	type Line struct {
		Score  int
		Answer string
	}
	var lines []Line
	for scanner.Scan() {
		line := scanner.Text()
		lineWords := strings.Fields(line)
		score := 0
		for _, word := range lineWords {
			for _, userWord := range userWords {
				if word == userWord {
					score++
				}
			}
		}
		startIndex := strings.Index(line, searchString)
		if startIndex != -1 {
			answer := line[startIndex+len(searchString):]
			lines = append(lines, Line{Score: score, Answer: answer})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Score > lines[j].Score
	})

	if len(lines) > 0 {
		n := rand.Intn(9) + 2
		if n > len(lines) {
			n = len(lines)
		}
		fmt.Println("Answers:")
		for i := 0; i < n; i++ {
			fmt.Println(lines[i].Answer)
		}
		return
	}

	// No matching line found
	connected, err := ioutil.ReadFile("connected.txt")
	if err != nil {
		log.Fatal(err)
	}
	if string(connected) == "1" {
		fmt.Println("I'm sorry, I don't have an answer to your question right now.")
	} else {
		files, err := ioutil.ReadDir("CONNECTED")
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			if !file.IsDir() && file.Size() > 0 {
				content, err := ioutil.ReadFile("CONNECTED/" + file.Name())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(content))
				return
			}
		}
	}
}
