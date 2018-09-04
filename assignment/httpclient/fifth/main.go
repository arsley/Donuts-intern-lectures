package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	prettyjson "github.com/hokaccha/go-prettyjson"
)

func isHTTPMethod(m string) bool {
	for _, v := range [5]string{"GET", "POST", "DELETE", "PATCH", "PUT"} {
		if m == v {
			return true
		}
	}
	return false
}

func main() {
	var (
		getURL    string
		reqMethod string
	)
	flag.StringVar(&getURL, "url", "http://mixch.tv/", "wanna *HTTP Method* URL")
	flag.StringVar(&reqMethod, "method", "GET", "wanna use HTTP method")
	flag.Parse()

	// assert HTTP method
	if isHTTPMethod(reqMethod) != true {
		fmt.Println("Error: Invalid HTTP method.")
		return
	}

	client := http.Client{}

	req, err := http.NewRequest(reqMethod, getURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	req = req.WithContext(ctx)
	defer cancel()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// prints header
	// for k, v := range res.Header {
	// 	fmt.Println(k, v)
	// }

	// body
	// fmt.Println(string(content))

	// wrong
	// modify in the future ...
	if res.Header["Content-Type"][0] == "application/json" {
		fmt.Println(prettyjson.Marshal(content))
	} else {
		fmt.Println(string(content))
	}
}
