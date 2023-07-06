package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Posts struct {
	XMLName xml.Name `xml:"posts"`
	Rows    []Row    `xml:"row"`
}

type Comments struct {
	XMLName xml.Name `xml:"comments"`
	Rows    []Row    `xml:"row"`
}

type Row struct {
	Id     string `xml:"Id,attr"`
	PostId string `xml:"PostId,attr"`
	Title  string `xml:"Title,attr"`
	Body   string `xml:"Body,attr"`
	Text   string `xml:"Text,attr"`
}

func main() {
	err := filepath.Walk("STACKOVERFLOW", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Base(path) != "Posts.xml" {
			return nil
		}

		dir := filepath.Dir(path)
		subfolderName := filepath.Base(dir)

		commentsPath := filepath.Join(dir, "Comments.xml")
		if _, err := os.Stat(commentsPath); os.IsNotExist(err) {
			return nil
		}

		processFiles(path, commentsPath, subfolderName)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func processFiles(postsPath, commentsPath, subfolderName string) {
	f, err := os.Create(filepath.Join(filepath.Dir(postsPath), subfolderName+".txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	postsFile, err := ioutil.ReadFile(postsPath)
	if err != nil {
		log.Fatal(err)
	}

	var posts Posts
	err = xml.Unmarshal(postsFile, &posts)
	if err != nil {
		log.Fatal(err)
	}

	commentsFile, err := ioutil.ReadFile(commentsPath)
	if err != nil {
		log.Fatal(err)
	}

	var comments Comments
	err = xml.Unmarshal(commentsFile, &comments)
	if err != nil {
		log.Fatal(err)
	}

	for _, postRow := range posts.Rows {
		var texts []string
		for _, commentRow := range comments.Rows {
			if postRow.Id == commentRow.PostId {
				texts = append(texts, commentRow.Text)
			}
		}

		if len(texts) > 0 {
			body := removeTags(postRow.Body)
			title := removeTags(postRow.Title)

			for i, text := range texts {
				texts[i] = removeTags(text)
			}

			line := strings.Join([]string{body, title}, "b2A1I9n14a1r18y25_g7l12o15b2e5")
			line += "b2A1I9n14a1r18y25_g7l12o15b2e5b2A1I9n14a1r18y25_g7l12o15b2e5" + strings.Join(texts, "b2A1I9n14a1r18y25_g7l12o15b2e5") + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n"
			_, err = f.WriteString(line)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func removeTags(s string) string {
	re := regexp.MustCompile(`<a[^>]*href\s*=\s*['"]([^'"]+)['"][^>]*>`)
	s = re.ReplaceAllString(s, "$1")

	re = regexp.MustCompile(`<[^>]*>`)
	s = re.ReplaceAllString(s, "")

	s = html.UnescapeString(s)

	return s
}
