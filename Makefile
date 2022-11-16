# Running a linter on the code
# He, probably, will not be able to find all the errors, but he will help to find the most obvious ones
lint:
	golangci-lint run

# Creating a swagger file
# Based on annotations in the code, it will create a file with the description of the API
swagger:
	swag init --generalInfo cmd/balance/main.go --output docs
	swag fmt

local:
	docker compose --project-directory ./build/balance up psql-local --detach
	go run ./cmd/balance/main.go

docker:
	docker compose --project-directory ./build/balance up psql-local --detach
	sleep 5
	docker compose --project-directory ./build/balance up balance

# Run service in the container using docker-compose and docker
# Service will be recreated with every call of this command
# Also, database will expose here 5432 port to localhost:5432
docker-build:
	docker compose --project-directory ./build/balance up psql-local --detach
	sleep 5
	docker compose --project-directory ./build/balance up --build balance

# If you want to run migrations, you need to run docker-compose up psql-local first
# Also, if you don't have a migrations package, you can need to download it first
#
# 		go get -u github.com/golang-migrate/migrate/v4/cmd/migrate
#
migrations-up:
	migrate -path ./internal/migrations/postgres \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		up

migrations-down:
	migrate -path ./internal/migrations/postgres \
		-database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" \
		down

test:
	 go test ./... -cover
