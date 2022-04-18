package grpc

import "github.com/vizitiuRoman/blockchain-api/pkg/delivery/grpc/invoice"

type Grpc struct {
	Invoice *invoice.Invoice
}

func NewGrpc() *Grpc {
	return &Grpc{
		Invoice: invoice.NewInvoice(),
	}
}
