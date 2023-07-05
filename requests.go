package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			userInput := fmt.Sprintf("User input %d", i)
			f, err := os.OpenFile("requests.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			if _, err := f.WriteString(userInput + " Eshioshi Favowrite Bynary Globe\n"); err != nil {
				fmt.Println(err)
			}
		}(i)
	}
	wg.Wait()
}
