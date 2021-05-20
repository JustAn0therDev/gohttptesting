package main

type sampleResponse struct {
	UserId int `json: "userId"`
	Id int `json: "id"`
	Title string `json: "title"`
	Completed bool `json: "completed"`
}
