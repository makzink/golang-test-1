package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

    
    r := mux.NewRouter()
    s := NewServer()

    r.HandleFunc("/employees", s.listEmployees).Methods("GET")
    r.HandleFunc("/employees", s.createEmployee).Methods("POST")
    r.HandleFunc("/employees/{id}", s.getEmployeeByID).Methods("GET")
    r.HandleFunc("/employees/{id}", s.updateEmployee).Methods("PUT")
    r.HandleFunc("/employees/{id}", s.deleteEmployee).Methods("DELETE")

    port := "8092"
    fmt.Println("Server up and running on port "+port)
    http.ListenAndServe(":"+port, r)
}
