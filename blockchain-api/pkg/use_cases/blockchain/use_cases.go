package blockchain

import (
	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain"
	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain/bitcoin"
	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain/ethereum"
	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain/litecoin"
	"github.com/vizitiuRoman/blockchain-api/pkg/helpers/use_cases"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
	"github.com/vizitiuRoman/blockchain-api/pkg/use_cases/blockchain/create_wallet"
)

type UseCases struct {
	CreateWallet use_cases.UseCase[create_wallet.Input, create_wallet.Output]
}

func NewUseCases(repos *repositories.Repository) *UseCases {
	clients := map[string]blockchain.Blockchain{
		blockchain.BITCOIN: bitcoin.NewClient(&bitcoin.Config{
			BaseURL:  "",
			Password: "",
			Username: "",
		}),
		blockchain.LITECOIN: litecoin.NewClient(),
		blockchain.ETHEREUM: ethereum.NewClient(),
	}

	return &UseCases{
		CreateWallet: create_wallet.NewUseCase(
			repos.Wallet,
			clients,
		),
	}
}
