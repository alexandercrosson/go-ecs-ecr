package main

import(
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello You Too!")
	})
}

func main(){
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":80", nil)
}
