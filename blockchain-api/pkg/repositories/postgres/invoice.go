package postgres

import "github.com/jmoiron/sqlx"

type InvoiceRepository struct {
	db *sqlx.DB
}

func NewInvoiceRepository(db *sqlx.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (w *InvoiceRepository) Create() (string, error) {
	return "", nil
}

func (w *InvoiceRepository) Get(id string) (string, error) {
	return "", nil
}
