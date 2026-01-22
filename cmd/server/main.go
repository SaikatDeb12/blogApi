package main

import (
	"fmt"
	"net/http"

	"blog/internal/router"
)

func main() {
	router := router.SetUpRouter()
	fmt.Println("Server started at port: 8080")
	http.ListenAndServe(":8080", router)
}
