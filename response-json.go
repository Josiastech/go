package main

import(
	"encoding/json"
	"net/http"
)

type Profile struct {
	Name string
	Hobies []string
}

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	profile := Profile{"Alex", []string{"snowborading", "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(js)
}