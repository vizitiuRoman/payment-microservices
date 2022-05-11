package main

import (
	"errors"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jmoiron/sqlx"
	"github.com/vizitiuRoman/blockchain-api/pkg/delivery/grpc/invoice"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories"
	"github.com/vizitiuRoman/blockchain-api/pkg/repositories/postgres"
	"github.com/vizitiuRoman/blockchain-api/pkg/use_cases"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresDB(&postgres.Config{
		DBName:   "db",
		Username: "db",
		Password: "db",
		Host:     "db",
		Port:     "5432",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("fail when connect to database")
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("fail when listen to tcp")
	}

	initServices(grpcServer, db)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Println(errors.New("Fail to serve gRPC"))
			signalChan <- os.Interrupt
		}
	}()

	log.Println("Server started and listening")
	<-signalChan
	gracefulStop(grpcServer, db)
}

func initServices(grpcServer *grpc.Server, db *sqlx.DB) {
	repos := repositories.NewRepository(db)
	_ = use_cases.NewUseCases(&use_cases.Dependencies{Repos: repos})

	invoice.RegisterInvoiceServiceServer(grpcServer, invoice.NewInvoice())
}

func gracefulStop(grpcServer *grpc.Server, db *sqlx.DB) {
	log.Println("Closing server...")
	defer log.Println("Server closed!")

	log.Println("Closing gRPC connections")
	grpcServer.GracefulStop()

	log.Println("Closing PostgreSQL connections")
	if err := db.Close(); err != nil {
		log.Println("fail when closing database")
	}
}
