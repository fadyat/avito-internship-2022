// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Artyom Fadeyev",
            "url": "https://github.com/fadyat"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/health": {
            "get": {
                "description": "Healthcheck endpoint, that checks if the service is alive and database connection is working.",
                "tags": [
                    "health"
                ],
                "summary": "Healthcheck",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.HealthSuccess"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/service": {
            "get": {
                "description": "Get all outer services info in the system",
                "tags": [
                    "OuterService"
                ],
                "summary": "Get all services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Services"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new outer service info in the system",
                "tags": [
                    "OuterService"
                ],
                "summary": "New service",
                "parameters": [
                    {
                        "description": "Outer service short info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OuterService"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ServiceCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/service/{id}": {
            "get": {
                "description": "Get outer service info in the system by id",
                "tags": [
                    "OuterService"
                ],
                "summary": "Get service by id",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "service_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.OuterService"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/release": {
            "post": {
                "description": "Release of the user's balance to another service",
                "tags": [
                    "Transaction"
                ],
                "summary": "Release of the user's balance",
                "parameters": [
                    {
                        "description": "Release info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Reservation"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ReservationReleased"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/replenishment": {
            "post": {
                "description": "Transaction of the user's balance by a certain amount and creating a replenishment transaction",
                "tags": [
                    "Transaction"
                ],
                "summary": "Transaction of the user's balance",
                "parameters": [
                    {
                        "description": "Transaction info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Transaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.TransactionCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/reservation": {
            "post": {
                "description": "Cancel reservation of the user's balance from another service",
                "tags": [
                    "Transaction"
                ],
                "summary": "Cancel reservation of the user's balance",
                "parameters": [
                    {
                        "description": "Reservation info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Reservation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ReservationCancelled"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/reservation/report": {
            "get": {
                "description": "Get reservation report",
                "tags": [
                    "Transaction"
                ],
                "summary": "Get reservation report",
                "parameters": [
                    {
                        "description": "Reservation report time",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReportTime"
                        }
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/user/{id}": {
            "get": {
                "description": "Get all user transactions in paginated form",
                "tags": [
                    "Transaction"
                ],
                "summary": "Get user transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Page size",
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "created_at",
                            "amount"
                        ],
                        "type": "string",
                        "default": "created_at, amount",
                        "description": "Order by",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.TransactionPaginated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/transaction/withdrawal": {
            "post": {
                "description": "Transaction of the user's balance by a certain amount and creating a withdrawal transaction",
                "tags": [
                    "Transaction"
                ],
                "summary": "Transaction of the user's balance",
                "parameters": [
                    {
                        "description": "Transaction info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Transaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.TransactionCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/wallet": {
            "get": {
                "description": "Get wallets from the system",
                "tags": [
                    "UserWallet"
                ],
                "summary": "Get wallets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserWallets"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new wallet in the system",
                "tags": [
                    "UserWallet"
                ],
                "summary": "New wallet",
                "parameters": [
                    {
                        "description": "Wallet info",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserWallet"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.UserWalletCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "407": {
                        "description": "Proxy Authentication Required",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        },
        "/api/v1/wallet/{id}": {
            "get": {
                "description": "Get user wallet from the system by id",
                "tags": [
                    "UserWallet"
                ],
                "summary": "Get wallet",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "Wallet id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserWallet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.OuterService": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "description": "@description: Name is a name of the service.\n@example:     aboba-service",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "url": {
                    "description": "@description: URL is a link to the implementation of the service.\n@example:     https://aboba-service.com",
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "dto.ReportTime": {
            "type": "object",
            "required": [
                "month",
                "year"
            ],
            "properties": {
                "month": {
                    "description": "@description: Month is a month of the report.\n@example:     1",
                    "type": "integer"
                },
                "year": {
                    "description": "@description: Year is a year of the report.\n@example:     2021",
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "dto.Reservation": {
            "type": "object",
            "required": [
                "amount",
                "order_id",
                "service_id",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "description": "@description: Amount is the amount of money, that was transferred.\n@example:     100",
                    "type": "integer",
                    "maximum": 1000000,
                    "minimum": 1
                },
                "order_id": {
                    "description": "@description: OrderID is a unique identifier of the order, that belongs to the service.\n@example:     1",
                    "type": "integer",
                    "minimum": 1
                },
                "service_id": {
                    "description": "@description: ServiceID is a unique identifier of the service, that made this transaction.\n@example:     1",
                    "type": "integer",
                    "minimum": 1
                },
                "user_id": {
                    "description": "@description: UserID is a unique identifier of the user, that owns this transaction.\n@example:     1",
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "dto.Transaction": {
            "type": "object",
            "required": [
                "amount",
                "user_id"
            ],
            "properties": {
                "amount": {
                    "description": "@description: Amount is the amount of money, that was transferred.\n@example:     100",
                    "type": "integer",
                    "maximum": 1000000,
                    "minimum": 1
                },
                "user_id": {
                    "description": "@description: UserID is a unique identifier of the user, that owns this transaction.\n@example:     1",
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "dto.UserWallet": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "description": "@description: UserID is a unique identifier of the user, that owns this wallet.\n@example:     1",
                    "type": "integer",
                    "minimum": 1
                }
            }
        },
        "models.OuterService": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description: ID is given unique identifier of the service.\n@example:     1",
                    "type": "integer"
                },
                "name": {
                    "description": "@description: Name is a name of the service.\n@example:     aboba-service",
                    "type": "string"
                },
                "url": {
                    "description": "@description: URL is a link to the implementation of the service.\n@example:     https://aboba-service.com",
                    "type": "string"
                }
            }
        },
        "models.Transaction": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "@description: Amount is the amount of money, that was transferred.\n@example:     100",
                    "type": "integer"
                },
                "created_at": {
                    "description": "@description: CreatedAt is the time, when the transaction was created.\n@example:     2021-01-01T00:00:00Z",
                    "type": "string"
                },
                "id": {
                    "description": "@description: ID is a unique identifier of the transaction.\n@example:     1",
                    "type": "integer"
                },
                "type": {
                    "description": "@description: Type is the type of the transaction.\n@example:     Replenishment",
                    "type": "string"
                },
                "user_id": {
                    "description": "@description: UserID is a unique identifier of the user, that owns this transaction.\n@example:     1",
                    "type": "integer"
                }
            }
        },
        "models.UserWallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "description": "@description: Balance is a current balance of the wallet.\n@example:     100",
                    "type": "integer"
                },
                "user_id": {
                    "description": "@description: UserID is a unique identifier of the user, that owns this wallet.\n@example:     1",
                    "type": "integer"
                }
            }
        },
        "responses.ErrorResp": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "@description: Description is a description of the error.\n@example:     invalid request with id=1",
                    "type": "string"
                },
                "message": {
                    "description": "@description: Message is a message of the error.\n@example:     invalid request",
                    "type": "string"
                }
            }
        },
        "responses.HealthSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "@description: Message is a success message.\n@example:     OK",
                    "type": "string"
                }
            }
        },
        "responses.Pagination": {
            "type": "object",
            "properties": {
                "found": {
                    "description": "@description: Found is a number of found items.\n@example:     10",
                    "type": "integer"
                },
                "next_page": {
                    "description": "@description: NextPage is a number of the next page.\n@example:     3",
                    "type": "integer"
                },
                "page": {
                    "description": "@description: Page is a number of the current page.\n@example:     2",
                    "type": "integer"
                },
                "per_page": {
                    "description": "@description: PerPage is a number of items per page.\n@example:     10",
                    "type": "integer"
                },
                "prev_page": {
                    "description": "@description: PrevPage is a number of the previous page.\n@example:     1",
                    "type": "integer"
                },
                "total": {
                    "description": "@description: Total is a total number of items.\n@example:     100",
                    "type": "integer"
                }
            }
        },
        "responses.ReservationCancelled": {
            "description": "ReservationCancelled is a response for reservation cancellation",
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description ID is given unique identifier of the reservation\n@example     1",
                    "type": "integer"
                }
            }
        },
        "responses.ReservationReleased": {
            "description": "ReservationReleased is a response for reservation release",
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description ID is given unique identifier of the reservation\n@example     1",
                    "type": "integer"
                }
            }
        },
        "responses.ServiceCreated": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description: ID is given unique identifier of the service.\n@example:     1",
                    "type": "integer"
                }
            }
        },
        "responses.Services": {
            "type": "object",
            "properties": {
                "services": {
                    "description": "@description: Services is a list of services.\n@example:     [{\"id\":1,\"name\":\"aboba-service\",\"url\":\"https://aboba-service.com\"}]",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OuterService"
                    }
                }
            }
        },
        "responses.TransactionCreated": {
            "description": "TransactionCreated is a response for transaction creation",
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description ID is given unique identifier of the transaction\n@example     1",
                    "type": "integer"
                }
            }
        },
        "responses.TransactionPaginated": {
            "description": "TransactionPaginated is a response for paginated transactions",
            "type": "object",
            "properties": {
                "pagination": {
                    "description": "@description Pagination is a pagination object, which have info about pages\n@example     {\"prev_page\":1,\"page\":2,\"next_page\":3,\"found\":10,\"limit\":10,\"total\":100}",
                    "$ref": "#/definitions/responses.Pagination"
                },
                "transactions": {
                    "description": "@description Transactions is a list of transactions, which are paginated by page and perPage\n@example     [{\"id\":1,\"user_id\":1,\"service_id\":1,\"amount\":100,\"type\":\"replenishment\",\"created_at\":\"2021-10-01T00:00:00Z\"}]",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Transaction"
                    }
                }
            }
        },
        "responses.UserWalletCreated": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "@description: ID is given unique identifier of the wallet.\n@example:     1",
                    "type": "integer"
                }
            }
        },
        "responses.UserWallets": {
            "type": "object",
            "properties": {
                "wallets": {
                    "description": "@description: UserWallets is a list of user wallets.\n@example:     [{\"user_id\":1,\"balance\":100}]",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserWallet"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:80",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Avito Internship 2022 Balance API",
	Description:      "This is a sample server for a balance API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
