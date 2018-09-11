package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Message : data structure of websocket
// label is not required if use same key; struct's member == json's key
type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

// Parse : JSON to GO's struct (single)
// {"username": "hello", "content": "json"} to Message struct
func Parse(data []byte) Message {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Fatalf("JSON parse error: %s", err)
	}
	return message
}

// ParseMany : JSON to GO's struct
// [
//   {"username": "hello", "content": "json"},
//   {"username": "hello", "content": "json"},
// ] to []Message
func ParseMany(data []byte) []Message {
	var message []Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Fatalf("JSON parse error: %s", err)
	}
	return message
}

// Stringify : Struct to JSON
func Stringify(username, content string) []byte {
	data := fmt.Sprintf("{\"username\": \"%s\", \"content\":\"%s\"}", username, content)
	return []byte(data)
}

// test code
// func main() {
// 	bytes, _ := ioutil.ReadFile("test.json")
// 	msg := Parse(bytes)
// 	fmt.Println(msg.Username, msg.Content)
// 	// msgs := ParseMany(bytes)
// 	// for _, v := range msgs {
// 	// 	fmt.Printf("Username: %s, Content: %s\n", v.Username, v.Content)
// 	// 	fmt.Println(string(Stringify(v.Username, v.Content)))
// 	// }
// }
