// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "ライセンス(必須)",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/carts/": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "carts"
                ],
                "summary": "カートに商品を追加する",
                "parameters": [
                    {
                        "description": "カートの商品",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/cart.PostCartsParams"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/v1/health": {
            "get": {
                "description": "ヘルスチェック",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "ヘルスチェック",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health_handler.HealthResponse"
                        }
                    }
                }
            }
        },
        "/v1/orders": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "注文をする",
                "parameters": [
                    {
                        "description": "注文商品",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/order.OrderParams"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "int"
                        }
                    }
                }
            }
        },
        "/v1/products": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "商品一覧を取得する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/products.fetchProductResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "商品を保存する",
                "parameters": [
                    {
                        "description": "登録商品",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/products.PostProductsParams"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/products.postProductResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "ユーザーを取得する",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.getUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "cart.PostCartsParams": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "health_handler.HealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "order.OrderParams": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "productID": {
                    "type": "string"
                }
            }
        },
        "products.PostProductsParams": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "products.fetchProductResponse": {
            "type": "object",
            "properties": {
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/products.productsWithOwnerModel"
                    }
                }
            }
        },
        "products.postProductResponse": {
            "type": "object",
            "properties": {
                "product": {
                    "$ref": "#/definitions/products.productResponseModel"
                }
            }
        },
        "products.productResponseModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "products.productsWithOwnerModel": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner_id": {
                    "type": "string"
                },
                "owner_name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "user.getUserResponse": {
            "type": "object",
            "properties": {
                "users": {
                    "$ref": "#/definitions/user.userResponseModel"
                }
            }
        },
        "user.userResponseModel": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "バージョン(1.0)",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "アプリケーション名",
	Description:      "説明",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
