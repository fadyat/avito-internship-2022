lint:
	golangci-lint run

swagger:
	swag init --generalInfo cmd/balance/main.go --output docs
	swag fmt

local:
	go run ./cmd/balance/main.go

docker:
	docker compose --project-directory ./build/balance up psql --detach
	sleep 3
	docker compose --project-directory ./build/balance up balance

docker-build:
	docker compose --project-directory ./build/balance up psql --detach
	sleep 3
	docker compose --project-directory ./build/balance up --build balance