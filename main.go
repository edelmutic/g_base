package main

import (
	"io"
	"net/http"
	"strconv"
)


func main() {
	http.HandleFunc("/fib", handleFib)
	http.ListenAndServe(":8080", nil)
	
}

func handleFib(w http.ResponseWriter, r *http.Request)  {
	 num := r.FormValue("num")
	 n, err := strconv.Atoi(num)
	 if err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
	 }
	io.WriteString(w, strconv.Itoa()) 
}