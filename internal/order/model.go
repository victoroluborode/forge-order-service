package order

import "time"

type Status string
type PaymentStatus string

const (
	StatusPending Status = "pending"
	StatusAwaitingPayment Status = "awaiting_payment"
	StatusPaid Status = "paid"
	StatusFailed Status = "failed"
	StatusCancelled Status = "cancelled"
	StatusFulfilled Status = "fulfilled"
	PaymentStatusUnpaid PaymentStatus = "unpaid"
	PaymentStatusAuthorized PaymentStatus = "authorized"
	PaymentStatusPaid PaymentStatus = "paid"
	PaymentStatusRefunded PaymentStatus = "refunded"
)



type Order struct {
	OrderID string
	UserID string
	ReferenceID string
	Currency string
	Status Status
	PaymentStatus 
	Total int64
	CreatedAt time.Time
	UpdatedAt time.Time
}