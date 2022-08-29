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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "summary": "Retrieves paginated response of users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.PageableResponse-response_UserResponse"
                        }
                    },
                    "400": {
                        "description": "in case of an error",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "summary": "Creates a user using the prerequisites provided",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "if a new user was created",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "in case of a bad request",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "summary": "Retrieves a user by the provided ID.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.PageableResponse-response_UserResponse"
                        }
                    },
                    "400": {
                        "description": "in case of a bad request",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    },
                    "404": {
                        "description": "if an unknown ID is provided",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "summary": "Updates an existing user.",
                "responses": {
                    "200": {
                        "description": "the updated users",
                        "schema": {
                            "$ref": "#/definitions/response.UserResponse"
                        }
                    },
                    "400": {
                        "description": "in case of a bad request",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "in case of an internal error",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "summary": "Deletes a user by id.",
                "responses": {
                    "204": {
                        "description": "if the user is deleted successfully"
                    },
                    "400": {
                        "description": "in case of a bad request",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "in case of an internal error",
                        "schema": {
                            "$ref": "#/definitions/response.APIErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.APIError": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string",
                    "example": "id"
                },
                "message": {
                    "type": "string",
                    "example": "invalid id"
                },
                "value": {}
            }
        },
        "response.APIErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.APIError"
                    }
                }
            }
        },
        "response.PageableResponse-response_UserResponse": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.UserResponse"
                    }
                },
                "empty": {
                    "type": "boolean"
                },
                "number": {
                    "type": "integer"
                },
                "numberOfElements": {
                    "type": "integer"
                },
                "totalElements": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "users.CreateUserRequest": {
            "type": "object",
            "required": [
                "age",
                "birthday",
                "email",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 50
                },
                "birthday": {
                    "type": "string",
                    "example": "2022-06-10"
                },
                "email": {
                    "type": "string",
                    "example": "Name@name.com"
                },
                "name": {
                    "description": "TODO, apply validation",
                    "type": "string",
                    "example": "Name"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "GO template API",
	Description:      "An template for implementing a hexagonal application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
