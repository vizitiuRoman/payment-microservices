package create_wallet

import (
	"errors"

	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain"
	"github.com/vizitiuRoman/blockchain-api/pkg/helpers/use_cases"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
)

type UseCase[In, Out any] struct {
	repo    repositories.Wallet
	clients map[string]blockchain.Blockchain
}

func NewUseCase(repo repositories.Wallet, clients map[string]blockchain.Blockchain) use_cases.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		repo:    repo,
		clients: clients,
	}
}

func (uc *UseCase[In, Out]) Execute(input *Input) (*Output, error) {
	client, ok := uc.clients[input.Currency]
	if !ok {
		return nil, errors.New("invalid currency")
	}

	account, err := client.CreateWallet()
	if err != nil {
		return nil, errors.New("cannot create wallet into blockchain")
	}

	_, err = uc.repo.Create()
	if err != nil {
		return nil, errors.New("cannot create wallet")
	}

	return &Output{
		Address: account.Address,
	}, nil
}
