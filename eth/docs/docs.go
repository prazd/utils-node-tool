// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-02-04 19:25:16.939232826 +0300 MSK m=+0.064597131

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is BUTTON Node API responseModels documentation",
        "title": "Swagger BUTTON Node API",
        "contact": {
            "name": "API Support",
            "email": "nk"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "1.0"
    },
    "basePath": "/",
    "paths": {
        "/eth/balance/{address}": {
            "get": {
                "description": "return balance of account in ETH for specific node",
                "produces": [
                    "application/json"
                ],
                "summary": "ETH balance of account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.BalanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/eth/balances": {
            "post": {
                "description": "return balances of accounts in ETH",
                "produces": [
                    "application/json"
                ],
                "summary": "ETH balance of accounts by list",
                "parameters": [
                    {
                        "description": "addressesArray",
                        "name": "addressesArray",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.BalancesResponse"
                            }
                        }
                    }
                }
            }
        },
        "/eth/gasPrice": {
            "get": {
                "description": "return Amount of ETH that you need to send a transaction",
                "produces": [
                    "application/json"
                ],
                "summary": "return gas price of specific node",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.GasPriceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/eth/tokenBalance/{sc-address}/{address}": {
            "get": {
                "description": "return balance of specific token in ETH node",
                "produces": [
                    "application/json"
                ],
                "summary": "return balance of specific token in ETH node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "sc-address",
                        "name": "sc-address",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.BalanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/eth/tokenBalances": {
            "post": {
                "description": "return tokens balances of account",
                "produces": [
                    "application/json"
                ],
                "summary": "ETH ERC-20 tokens balance of account by list of smart contracts",
                "parameters": [
                    {
                        "description": "addressesArray",
                        "name": "addressesArray",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.BalancesResponse"
                            }
                        }
                    }
                }
            }
        },
        "/eth/transactionFee": {
            "get": {
                "description": "return Amount of ETH that you need to send a transaction",
                "produces": [
                    "application/json"
                ],
                "summary": "return Amount of ETH that you need to send a transaction",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.TransactionFeeResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.BalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                }
            }
        },
        "responses.BalancesResponse": {
            "type": "object",
            "properties": {
                "balances": {
                    "type": "array",
                    "items": {
                        "type": "\u0026{%!s(token.Pos=287) string float64}"
                    },
                    "example": [
                        "0"
                    ]
                }
            }
        },
        "responses.GasPriceResponse": {
            "type": "object",
            "properties": {
                "gasPrice": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "responses.TransactionFeeResponse": {
            "type": "object",
            "properties": {
                "fee": {
                    "type": "number",
                    "example": 0
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
