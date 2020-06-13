package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	get()

}

func get() {
	fmt.Println(" Entrando get a http")
	fmt.Println("Ingresa un número de post")
	var numberPost string
	fmt.Scanln(&numberPost)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + string(numberPost))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	//Convertir la respuesta a cadena
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String" + bodyString)

	//Convertimos la respuesta a todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Println("LA RESPUESTA DEL API COMO ESTRUCTURA", todoStruct)

}
