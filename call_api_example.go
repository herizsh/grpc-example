package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TodoC struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func mainC() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Create an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", todo)
}
