package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Gets information from the url: "https://jsonplaceholder.typicode.com/todos/2"
func Get() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	checkError(err)

	defer resp.Body.Close()

	var sampleResp sampleResponse

	decodeError := json.NewDecoder(resp.Body).Decode(&sampleResp)

	checkError(decodeError)

	concatUserId := fmt.Sprintf("UserId:  %v", sampleResp.UserId)
	concatId := fmt.Sprintf("Id: %v", sampleResp.Id)
	concatTitle := fmt.Sprintf("Title: %v", sampleResp.Title)
	concatCompleted := fmt.Sprintf("Completed: %v", sampleResp.Completed)

	fmt.Println(concatUserId)
	fmt.Println(concatId)
	fmt.Println(concatTitle)
	fmt.Println(concatCompleted)
}
