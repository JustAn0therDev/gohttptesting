package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"bytes"
)

func Post() {
	samplePostStruct := new(samplePostRequest)

	samplePostStruct.UserId = 2
	samplePostStruct.Body = "Write something here"
	samplePostStruct.Title = "A title here"

	jsonStringInByteArray, marshalingError := json.Marshal(samplePostStruct)

	checkError(marshalingError)

	resp, err := http.Post("http://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonStringInByteArray))

	checkError(err)

	var sampleStructFromResponse sampleResponse

	defer resp.Body.Close()

	decodeError := json.NewDecoder(resp.Body).Decode(&sampleStructFromResponse)

	checkError(decodeError)

	fmt.Println(sampleStructFromResponse)
}
