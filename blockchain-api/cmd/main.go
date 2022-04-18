package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
	"github.com/vizitiuRoman/blockchain-api/pkg/use_cases"
	"github.com/vizitiuRoman/blockchain-api/pkg/use_cases/blockchain/create_wallet"
)

func main() {
	repos := repositories.NewRepository(&sqlx.DB{})

	useCases := use_cases.NewUseCases(&use_cases.Dependencies{Repos: repos})

	wallet, err := useCases.Blockchain.CreateWallet.Execute(&create_wallet.Input{
		Currency: "BTQC",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wallet.Address)
}
