package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var (
		getURL string
	)
	flag.StringVar(&getURL, "url", "http://mixch.tv/", "GET URL")
	flag.Parse()

	res, err := http.Get(getURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(content))
}
