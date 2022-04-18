package ethereum

import "github.com/vizitiuRoman/blockchain-api/pkg/blockchain"

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) CreateWallet() (*blockchain.Account, error) {
	return &blockchain.Account{}, nil
}
