package postgres

import "github.com/jmoiron/sqlx"

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (w *WalletRepository) Create() (string, error) {
	return "", nil
}
