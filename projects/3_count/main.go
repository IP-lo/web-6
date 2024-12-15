package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main(){
	counter := 0
	http.HandleFunc("/count",func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET"{
			fmt.Fprintf(w,"полученное значение: %d",counter)
		} else if r.Method == "POST"{
			r.ParseForm()
			count, err := strconv.Atoi(r.FormValue("count")) 
			if err != nil{
				http.Error(w,"IT'S NOT NUMBER: ",http.StatusBadRequest)
			}else{
				counter += count
				fmt.Fprintf(w,"введённое значение: %d",count)
			}
			
		}
	})
	if err := http.ListenAndServe(":3333", nil); err != nil{
		log.Println("Серверная ошибка:",err)
	}
}