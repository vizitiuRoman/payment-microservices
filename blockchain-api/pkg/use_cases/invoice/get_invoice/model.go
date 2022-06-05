package get_invoice

type Input struct {
	ID string
}

type Output struct {
	ID       string
	WalletID string
	Currency string
	Amount   string
}
