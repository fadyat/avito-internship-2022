FROM golang:1.19-alpine AS builder

WORKDIR /app
COPY .. .
RUN go mod download -x && go mod verify
RUN cd cmd/balance && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /main main.go

FROM scratch
COPY --from=builder /main /main
