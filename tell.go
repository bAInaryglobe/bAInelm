package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Size() > 99*1024*1024 {
			fmt.Printf("%s: %.2f MB\n", path, float64(info.Size())/1024/1024)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
