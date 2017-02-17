package main

import "net/http"

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}


func foo(w http.ResponseWriter, req *http.Request){
	//w.Header().Set("Server", "A GO web Server")
	//w.WriteHeader(200)
	w.Write([]byte("OK"))
}