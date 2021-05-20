package main

type samplePostRequest struct {
	Id int `json: "id"`
	UserId int `json:"userId"`
	Body string `json:"body"`
	Title string `json:"title"`
}
