package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"Hello, web!")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Println("Серверная ошибка:",err)
	}
}