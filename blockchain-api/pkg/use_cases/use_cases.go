package use_cases

import (
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
	blockchain_use_cases "github.com/vizitiuRoman/blockchain-api/pkg/use_cases/blockchain"
)

type UseCases struct {
	Blockchain *blockchain_use_cases.UseCases
}

type Dependencies struct {
	Repos *repositories.Repository
}

func NewUseCases(deps *Dependencies) *UseCases {
	return &UseCases{
		Blockchain: blockchain_use_cases.NewUseCases(deps.Repos),
	}
}
