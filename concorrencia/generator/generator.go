package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan - canal somente leitura
func getTitle(urls ...string) <-chan string {
	c := make(chan string)
	regex, _ := regexp.Compile("<title>(.*?)<\\/title>")

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			c <- regex.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := getTitle("https://google.com", "https://bing.com")
	t2 := getTitle("https://amazon.com", "https://youtube.com")

	fmt.Println("1)", <-t1, "|", <-t2)
	fmt.Println("2):", <-t1, "|", <-t2)
}
