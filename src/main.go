package main

import (
	"fmt"
	"guide/helper"
	"guide/routes"
	"net/http"
)

func main() {

	helper.Load()
	
	serveRouter := routes.MainRouter()
	
	fmt.Println("Serveur lancé sur : http://localhost:8080")
	http.ListenAndServe("localhost:8080", serveRouter)
}
