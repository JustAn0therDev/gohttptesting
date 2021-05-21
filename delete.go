package main

import (
	"fmt"
	"net/http"
	"bytes"
)

// nice. To be honest, they are a little "redundant" but pretty straight-forward to work with when
// using requests from a Go program. Now to the Web API creation!
func Delete() {
	reqUrl := "https://jsonplaceholder.typicode.com/posts/1"

	client := http.Client{}

	req, err := http.NewRequest(http.MethodDelete, reqUrl, bytes.NewBuffer(make([]byte, 0)))

	checkError(err)

	resp, respErr := client.Do(req)

	checkError(respErr)

	fmt.Println(resp.StatusCode)
}
