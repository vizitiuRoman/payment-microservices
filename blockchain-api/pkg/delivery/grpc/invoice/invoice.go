package invoice

import (
	"context"

	blockchain_use_cases "github.com/vizitiuRoman/blockchain-api/pkg/use_cases/blockchain"
	invoice_use_cases "github.com/vizitiuRoman/blockchain-api/pkg/use_cases/invoice"
)

type Invoice struct {
	UnimplementedInvoiceServiceServer

	BlockchainUseCases *blockchain_use_cases.UseCases
	InvoiceUseCases    *invoice_use_cases.UseCases
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (i *Invoice) GetInvoice(ctx context.Context, req *GetInvoiceInput) (*InvoiceOutput, error) {
	return nil, nil
}

func (i *Invoice) CreateInvoice(ctx context.Context, req *CreateInvoiceInput) (*InvoiceOutput, error) {
	return nil, nil
}
