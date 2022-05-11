clean:
	@go clean ./...
	@go mod tidy

build:
	cd cmd && go build -trimpath -gcflags=-trimpath=%CD% -asmflags=-trimpath=%CD% -ldflags "-s -w" main.go

test:
	@go test -race -v -p 2 -coverpkg=./... -covermode=atomic -coverprofile cover.out.tmp ./...

generate: ## Create golang protos
	@echo "Generating protos"
	@echo " - Generating messages"
	@protoc --go_out=. proto/*.proto
	@echo " - Generating services"
	@protoc --go-grpc_out=. proto/*.proto