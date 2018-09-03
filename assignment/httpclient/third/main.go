package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
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

	fmt.Println(string(content))
}
