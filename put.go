package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

// everything that is not a get, head or post request in go
// needs to be treated with creating a client struct instance and
// passing in a request object to its "Do" function.
func Put() {
	// using another way of creating a struct just for testing
	samplePostStruct := samplePostRequest{}

	samplePostStruct.Id = 1
	samplePostStruct.UserId = 2
	samplePostStruct.Title = "Updated title. Again!"
	samplePostStruct.Body = "Updated write something here. Again."

	// I don't need the value to be passed by reference here.
	// passing by reference just saves unused memory that can be used later
	// by the same device in the same or other processes.
	jsonByteArray, marshalingError := json.Marshal(&samplePostStruct)

	checkError(marshalingError)

	client := http.Client{}

	request, reqErr := http.NewRequest(http.MethodPut, "http://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonByteArray))

	checkError(reqErr)

	request.Header.Add("content-type", "application/json")

	resp, err := client.Do(request)

	checkError(err)

	var sampleStructFromResponse sampleResponse

	decodeError := json.NewDecoder(resp.Body).Decode(&sampleStructFromResponse)

	checkError(decodeError)

	fmt.Println(fmt.Sprintf("Status Code: %v", resp.StatusCode))
	fmt.Println(sampleStructFromResponse)
}
