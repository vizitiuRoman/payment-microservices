package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories/postgres/invoice"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories/postgres/wallet"
)

//go:generate mockery --dir . --name Wallet --name Invoice --output ./mocks

type Wallet interface {
	Create() (string, error)
}

type Invoice interface {
	Create(*invoice.Invoice) (*invoice.Invoice, error)
	Get(id string) (*invoice.Invoice, error)
}

type Repository struct {
	Wallet
	Invoice
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet:  wallet.NewRepository(db),
		Invoice: invoice.NewRepository(db),
	}
}
