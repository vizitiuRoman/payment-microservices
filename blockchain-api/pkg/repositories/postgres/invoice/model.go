package invoice

type Invoice struct {
	ID       string `db:"id"`
	WalletID string `db:"wallet_id"`
	Amount   string `db:"amount"`
	Currency string `db:"currency"`
}
