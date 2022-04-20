package get_invoice

import (
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
		return nil, err
	}

	return &Output{
		Amount:   invoice,
		Currency: invoice,
	}, nil
}
