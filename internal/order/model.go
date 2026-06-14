package order

import "time"

type Status string

const (
	StatusPending Status = "pending"
	StatusAwaitingPayment Status = "awaiting_payment"
	StatusPaid Status = "paid"
	StatusFailed Status = "failed"
	StatusCancelled Status = "cancelled"
	StatusFulfilled Status = "fulfilled"
)



type Order struct {
	OrderID string
	UserID string
	ReferenceID string
	Currency string
	Status Status
	Total int64
	CreatedAt time.Time
	UpdatedAt time.Time
}


func CanTransition(from, to Status) bool {
	allowedTransitions := map[Status][]Status {
		StatusAwaitingPayment: {StatusPaid, StatusFailed, StatusCancelled, StatusAwaitingPayment},
		StatusCancelled: {},
		StatusFailed: {StatusAwaitingPayment, StatusCancelled},
		StatusFulfilled: {},
		StatusPaid: {StatusFulfilled},
		StatusPending: {StatusAwaitingPayment, StatusCancelled},
	}

	for _, allowedStatus := range allowedTransitions[from] {
		if allowedStatus == to {
			return true
		}
	}
	return false
}