package main

import (
	"fmt"
	"log"
	"net/http" 
)

func main(){
	http.HandleFunc("/api/user",func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != ""{
			fmt.Fprintf(w,"Hello,%s!", name)
		} else{
			fmt.Fprintln(w,"name is empty!")

		}
	})
	if err := http.ListenAndServe(":9000", nil); err != nil{
		log.Println("Серверная ошибка:",err)
	}
}