package invoice

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(input *Invoice) (output *Invoice, error error) {
	query := `
		INSERT INTO invoices(wallet_id, currency, amount)
        VALUES(:wallet_id, :currency, :amount)
        returning *
	`
	row, err := r.db.NamedQuery(query, input)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&output)
	return output, err
}

func (r *Repository) Get(id string) (output *Invoice, error error) {
	query := `
		SELECT * FROM invoices
        WHERE id = $1
    `
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	err = row.Scan(&output)
	return output, err
}
