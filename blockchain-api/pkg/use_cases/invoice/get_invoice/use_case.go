package get_invoice

import (
	"errors"

	"github.com/vizitiuRoman/blockchain-api/pkg/helpers/use_cases"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
)

type UseCase[In, Out any] struct {
	repo repositories.Invoice
}

func NewUseCase(repo repositories.Invoice) use_cases.UseCase[Input, Output] {
	return &UseCase[Input, Output]{
		repo: repo,
	}
}

func (uc *UseCase[In, Out]) Execute(input *Input) (*Output, error) {
	invoice, err := uc.repo.Get(input.ID)
	if err != nil {
		return nil, errors.New("invoice not found")
	}

	return &Output{
		ID:       invoice.ID,
		WalletID: invoice.WalletID,
		Amount:   invoice.Amount,
		Currency: invoice.Currency,
	}, nil
}
