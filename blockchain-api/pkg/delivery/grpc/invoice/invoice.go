package invoice

import "context"

type Invoice struct {
	UnimplementedInvoiceServiceServer
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (i *Invoice) GetInvoice(ctx context.Context, req *InvoiceInput) (*InvoiceOutput, error) {
	return nil, nil
}
