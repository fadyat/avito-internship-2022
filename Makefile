lint:
	golangci-lint run

swagger:
	swag init --generalInfo cmd/balance/main.go --output docs
	swag fmt

balance-local:
	go run ./cmd/balance/main.go

balance-container:
	docker compose --project-directory ./build/balance up --build balance