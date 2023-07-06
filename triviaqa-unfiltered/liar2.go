package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const maxChunkSize int64 = 49 * 1024 * 1024

func main() {
	// Get the current directory.
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate over all the files in the directory.
	files, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		// Check if the file is not empty.
		stat, err := os.Stat(file)
		if err != nil || stat.Size() == 0 {
			continue
		}

		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		chunkNum := 1
		chunkSize := int64(0)
		emptyLineCount := 0
		buffer := []string{}

		newFile, w := createNewFile(file, chunkNum)
		defer newFile.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			chunkSize += int64(len(line))

			if line == "" {
				emptyLineCount++
			} else {
				emptyLineCount = 0
			}

			buffer = append(buffer, line)

			if chunkSize >= maxChunkSize && emptyLineCount >= 5 {
				for _, l := range buffer[:len(buffer)-emptyLineCount] {
					fmt.Fprintln(w, l)
				}
				newFile.Close()
				chunkNum++
				newFile, w = createNewFile(file, chunkNum)
				defer newFile.Close()
				chunkSize = int64(0)
				buffer = buffer[len(buffer)-emptyLineCount:]
			}

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
				return
			}
		}

		for _, l := range buffer {
			fmt.Fprintln(w, l)
		}
	}
}

func createNewFile(file string, chunkNum int) (*os.File, io.Writer) {
	newFileName := fmt.Sprintf("%s.%d.json", file, chunkNum)
	newFile, err := os.Create(newFileName)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	w := bufio.NewWriter(newFile)
	return newFile, w
}
