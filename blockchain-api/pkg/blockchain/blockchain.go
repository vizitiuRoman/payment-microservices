package blockchain

type Blockchain interface {
	CreateWallet() (*Account, error)
}

type Account struct {
	Address    string
	PrivateKey string
}
