package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

type Comment struct {
	Text string `xml:"Text,attr"`
}

type Comments struct {
	XMLName xml.Name  `xml:"comments"`
	Rows    []Comment `xml:"row"`
}

func main() {
	xmlFile, err := os.Open("Comments.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var comments Comments
	xml.Unmarshal(byteValue, &comments)

	file, err := os.Create("overview.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for _, row := range comments.Rows {
		text := row.Text
		text = html.UnescapeString(text)
		text = extractLinks(text)
		file.WriteString(text + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n" + "\n")
	}
}

func extractLinks(text string) string {
	re := regexp.MustCompile(`<a[^>]*href="([^"]*)"[^>]*>(.*?)</a>`)
	text = re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		href := submatches[1]
		linkText := submatches[2]
		u, err := url.Parse(href)
		if err != nil || u.Scheme == "" || u.Host == "" {
			return linkText
		}
		return href
	})
	text = stripTags(text)
	return text
}

func stripTags(text string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	text = re.ReplaceAllString(text, "")
	return text
}
