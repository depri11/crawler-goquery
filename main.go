package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	blogTitles, err := GetLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Blog Titles: ")
	fmt.Println(blogTitles)
}

func GetLatestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.body)
	if err != nil {
			
	}
}