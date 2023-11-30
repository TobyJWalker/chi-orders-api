package handler

import (
	"chi-orders-api/model"
	"chi-orders-api/repository/order"
	"strconv"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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
	
	// get id from url
	idParam := chi.URLParam(r, "id")

	// convert to uint64
	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get order
	order, err := o.Repo.FindByID(r.Context(), orderID)
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
	w.WriteHeader(http.StatusOK)

}

// update by id method
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	
	// capture body structure
	var body struct {
		Status string `json:"status"`
	}

	// json decode to check for bad request
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get id from url
	idParam := chi.URLParam(r, "id")

	// convert to uint64
	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get order
	order, err := o.Repo.FindByID(r.Context(), orderID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// init constants
	const completedStatus = "completed"
	const shippedStatus = "shipped"
	
	// get current time
	now := time.Now().UTC()

	// check status
	switch body.Status {
	case shippedStatus:

		// check if order is already shipped
		if order.ShippedAt != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		order.ShippedAt = &now // set shipped at time
	case completedStatus:

		// check if order is already completed or not shipped
		if order.CompletedAt != nil || order.ShippedAt == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		order.CompletedAt = &now // set completed at time
	
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// update order
	err = o.Repo.Update(r.Context(), order)
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
	w.WriteHeader(http.StatusOK)

}

// delete method
func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	
	// get id from url
	idParam := chi.URLParam(r, "id")

	// convert to uint64
	const base = 10
	const bitSize = 64
	orderID, err := strconv.ParseUint(idParam, base, bitSize)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// delete order
	err = o.Repo.DeleteByID(r.Context(), orderID)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	w.WriteHeader(http.StatusOK)

}