package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vaibhav-sd/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("server is getting stated...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
