package main

import (
	"fmt"
	"go-mongo/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(os.Getenv("mongoURI"))
	fmt.Println("test Alo: " + os.Getenv("alo"))
	sm := mux.NewRouter()
	sm.HandleFunc("/p", handlers.CreateProduct).Methods(http.MethodPost)
	sm.HandleFunc("/ps", handlers.GetAllProduct).Methods(http.MethodGet)
	sm.HandleFunc("/p/{id}", handlers.UpdateProduct).Methods(http.MethodPut)
	sm.HandleFunc("/pByName", handlers.GetProductByName).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8000", sm))
}
