package main

import (
	"encoding/json"
	"net/http"
	"time"
	"database/sql"
	_ "github.com/jack/pgx/v5/stdlib"

	"github.com/victoroluborode/forge-order-service/internal/order"
)

type CreateOrderRequest struct {
	UserID string `json:"user_id"`
	Total int64 `json:"total"`
	Currency string `json:"currency"`
}

func createOrderHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Total <= 0 || req.Currency == "" || req.UserID == "" {
		http.Error(w,"invalid request body", http.StatusBadRequest)
		return
	}

	newOrder := order.Order{
		OrderID: "temp-id-123",
		UserID: req.UserID,
		Currency: req.Currency,
		Status: order.StatusPending,
		Total: req.Total,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}


db, err := sql.Open("pgx", "postgres://victoroluborode:IAMdamilare170%23@localhost:5432/forge_orders?sslmode=disable")

if err != nil {
	log.Fatal(err)
}

defer db.Close()

repo := order.NewRepository(db)

func main() {
	http.HandleFunc("/orders", createOrderHandler)
	http.ListenAndServe(":8080", nil)
}