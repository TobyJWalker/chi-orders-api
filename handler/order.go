package handler

import (
	"chi-orders-api/model"
	"chi-orders-api/repository/order"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// create an order struct for the repository
type Order struct {
	Repo *order.PostgresRepo
}

// create method
func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	
	// struct for expected post data
	var body struct {
		CustomerID uuid.UUID `json:"customer_id"`
		LineItems []model.LineItem `json:"line_items"`
	}

	// json decode to check for bad request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// creation time
	now := time.Now().UTC()

	// create order
	order := model.Order{
		CustomerID: body.CustomerID,
		LineItems: body.LineItems,
		CreatedAt: &now,
	}

	// insert order
	err := o.Repo.Insert(r.Context(), order)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// marshall order to json
	res, err := json.Marshal(order)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	w.Write(res)
	w.WriteHeader(http.StatusCreated)

}

// list method
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	
	// get orders
	orders, err := o.Repo.FindAll(r.Context())
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// marshall orders to json
	res, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	w.Write(res)
	w.WriteHeader(http.StatusOK)

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