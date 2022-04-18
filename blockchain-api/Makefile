generate: ## Create golang protos
	@echo "Generating protos"
	@echo " - Generating messages"
	@protoc --go_out=. proto/*.proto
	@echo " - Generating services"
	@protoc --go-grpc_out=. proto/*.proto