package main

import (
	"os"
	"net/http"
)

func main()  {
	dir, _:=os.Getwd()
	http.ListenAndServe(":3000", http.FileServer(http.Dir(dir)))
}