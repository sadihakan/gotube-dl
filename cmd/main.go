package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/node-to-go/internal"
)

var key string = ""
var maxResults string = ""
var url string = ""

func main() {
	flag.StringVar(&key, "key", "AIzaSyC0-9QZFnXu6SpFmpkaWkktWvyIDhqXR1M", "-key")
	flag.StringVar(&maxResults, "maxResults", "1", "-maxResults")
	flag.StringVar(&url, "url", "a", "-url")
	flag.Parse()

	dl := internal.NewDL(key, url, maxResults)
	url, err := dl.Search()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("URL: ", url)
}
