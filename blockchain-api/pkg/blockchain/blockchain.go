package blockchain

const (
	CONFIRMED = "CONFIRMED"
	PENDING   = "PENDING"
)

const (
	BITCOIN      = "BTC"
	LITECOIN     = "LTC"
	ETHEREUM     = "ETH"
	TETHER_TRC20 = "USDT-TRC20"
	TETHER_ERC20 = "USDT-ERC20"
)

type Blockchain interface {
	CreateWallet() (*Account, error)
}

type Account struct {
	Address    string
	PrivateKey string
}
