package invoice

import (
	"github.com/vizitiuRoman/blockchain-api/pkg/helpers/use_cases"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
	"github.com/vizitiuRoman/blockchain-api/pkg/use_cases/invoice/get_invoice"
)

type UseCases struct {
	GetInvoice use_cases.UseCase[get_invoice.Input, get_invoice.Output]
}

func NewUseCases(repos *repositories.Repository) *UseCases {
	return &UseCases{
		GetInvoice: get_invoice.NewUseCase(repos.Invoice),
	}
}
