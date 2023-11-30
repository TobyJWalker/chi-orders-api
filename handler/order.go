package handler

import (
	"fmt"
	"net/http"
)

// create an order struct
type Order struct {}

// create method
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
}

// list method
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list order")
}

// get by id method
func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get order by id")
}

// update by id method
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update order by id")
}

// delete method
func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete order")
}