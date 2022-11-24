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
        "/v1/users": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "upsert wallet address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "upsert new user",
                "parameters": [
                    {
                        "description": "upsert new user request requires username, password, and role",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_VinCPR_backend_api.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/github.com_VinCPR_backend_api.userResponse"
                        }
                    }
                }
            }
        },
        "/v1/users/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "upsert wallet address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "upsert new user",
                "parameters": [
                    {
                        "description": "login request requires username, password",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_VinCPR_backend_api.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/github.com_VinCPR_backend_api.loginUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.createUserRequest": {
            "type": "object",
            "required": [
                "password",
                "role_name",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 10
                },
                "role_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.loginUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expires_at": {
                    "type": "string"
                },
                "session_id": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "role_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "github.com_VinCPR_backend_api.createUserRequest": {
            "type": "object",
            "required": [
                "password",
                "role_name",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 10
                },
                "role_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "github.com_VinCPR_backend_api.loginUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "github.com_VinCPR_backend_api.loginUserResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "access_token_expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "refresh_token_expires_at": {
                    "type": "string"
                },
                "session_id": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/github.com_VinCPR_backend_api.userResponse"
                }
            }
        },
        "github.com_VinCPR_backend_api.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "role_name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
