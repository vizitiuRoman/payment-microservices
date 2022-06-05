package get_invoice

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vizitiuRoman/blockchain-api/pkg/blockchain"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories/mocks"
	invoicerepo "github.com/vizitiuRoman/blockchain-api/pkg/repositories/postgres/invoice"
)

func TestGetInvoice(t *testing.T) {
	t.Run("should be found invoice by id", func(t *testing.T) {
		repo := mocks.NewInvoice(t)

		repo.On("Get", "1").Return(&invoicerepo.Invoice{
			ID:       "1",
			WalletID: "1",
			Amount:   "1",
			Currency: blockchain.BITCOIN,
		}, nil)

		useCase := NewUseCase(repo)

		input := &Input{ID: "1"}
		invoice, err := useCase.Execute(input)

		assert.Equal(t, &Output{ID: "1", Amount: "1", Currency: blockchain.BITCOIN, WalletID: "1"}, invoice)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("should not be found invoice by id", func(t *testing.T) {
		repo := mocks.NewInvoice(t)
		repo.On("Get", "1").Return(nil, errors.New("error"))

		useCase := NewUseCase(repo)

		input := &Input{ID: "1"}
		invoice, err := useCase.Execute(input)

		assert.Equal(t, "invoice not found", err.Error())
		assert.Nil(t, invoice)
		repo.AssertExpectations(t)
	})
}
