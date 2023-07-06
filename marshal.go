package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	folder := "triviaqa-unfiltered"

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			jsonFile, err := os.Open(filepath.Join(folder, file.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer jsonFile.Close()

			var data interface{}
			if err := json.NewDecoder(jsonFile).Decode(&data); err != nil {
				fmt.Println(err)
				return
			}

			xmlFile, err := os.Create(filepath.Join(folder, strings.TrimSuffix(file.Name(), ".json")+".xml"))
			if err != nil {
				fmt.Println(err)
				return
			}
			defer xmlFile.Close()

			if _, err := xmlFile.Write([]byte(xml.Header)); err != nil {
				fmt.Println(err)
				return
			}

			if err := xml.NewEncoder(xmlFile).Encode(data); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
