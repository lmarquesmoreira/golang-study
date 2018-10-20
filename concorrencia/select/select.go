package main

import (
	"fmt"
	"time"

	"github.com/lmarquesmoreira/html"
)

func theFast(url1, url2, url3 string) string {
	c1 := html.GetTitle(url1)
	c2 := html.GetTitle(url2)
	c3 := html.GetTitle(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1200 * time.Millisecond):
		return "all loses"
		// default:
		// 	return "no response yet"
	}
}

func main() {
	champion := theFast(
		"https://facebook.com",
		"https://bing.com",
		"https://youtube.com",
	)
	fmt.Println(champion)
}
