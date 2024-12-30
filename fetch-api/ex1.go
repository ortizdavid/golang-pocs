package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `jon:"completed"`
}


func getTodo() {
	fmt.Println("1 - Performing HTTP Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Printf("API Response:\n%v", bodyString)
}


func postTodo() {
	fmt.Println("\n2 - Performing HTTP Post...")
	todo := Todo{
		UserId:    2,
		Id:        2,
		Title:     "New Todo",
		Completed: false,
	}

	jsonReq, _ := json.Marshal(todo)
	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}


func putTodo()  {
	fmt.Println("\n3 - Performing HTTP Put...")
	todo := Todo{
		UserId:    3,
		Id:        3,
		Title:     "",
		Completed: false,
	}

	jsonReq, _ := json.Marshal(todo)
	req, _ := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func main() {
	//getTodo()
	//postTodo()
	putTodo()
}
