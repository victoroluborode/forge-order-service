// 1. Setup and structure
package order

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db : db}
}


// 2. Create an Order (Create)
func (r *Repository) Create(ctx context.Context, o *Order) error {
	query := `
		INSERT INTO orders (order_id, user_id, reference_id, currency, status, total, created_at, updated_at)
		VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7)
		RETURNING order_id`

	return r.db.QueryRowContext(ctx, query, o.UserID, o.ReferenceID, o.Currency, o.Status, o.Total, o.CreatedAt, o.UpdatedAt,).Scan(&o.OrderID)
}


// 3. Fetching an Order (GetByID)
func (r *Repository) GetById(ctx context.Context, id string) (*Order, error) {
	query := `
		SELECT order_id, user_id, reference_id, currency, status, total, created_at, updated_at FROM orders WHERE order_id = $1`

	var o Order
	err := r.db.QueryRowContext(ctx, query, id).Scan(&o.OrderID, &o.UserID, &o.ReferenceID, &o.Currency, &o.Status, &o.Total, &o.CreatedAt, &o.UpdatedAt,)

	if err != nil {
		return nil, err
	}
	return &o, nil
}


// 4. Deleting an order (DeleteByID) 

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM orders WHERE order_id = $1`

	result, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}