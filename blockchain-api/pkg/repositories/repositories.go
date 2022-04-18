package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories/postgres"
)

type Wallet interface {
	Create() (string, error)
}

type Repository struct {
	Wallet
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Wallet: postgres.NewWalletRepository(db),
	}
}
