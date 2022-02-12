// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/": {
            "get": {
                "description": "Decrypt card by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Decrypt card by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.PlainCardResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update card",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Update card",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "List cards",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "List cards",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.CardResponse"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete card by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Delete card by id",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/confirm/{code}": {
            "get": {
                "description": "Confirmation for user accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Confirmation for user accounts",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/register": {
            "post": {
                "description": "Register using email and password",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Register using email and password",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/resend-confirmation": {
            "post": {
                "description": "Confirmation for user accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Confirmation for user accounts",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/reset-password": {
            "post": {
                "description": "Reset password using email",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Reset password using email",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/accounts/reset-password/{code}": {
            "post": {
                "description": "Reset password for a given reset code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Reset password for a given reset code",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/jwks": {
            "get": {
                "description": "Returns jwks details",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Returns jwks details",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Refresh access token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.TokenResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/shared.HTTPError"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "delete": {
                "description": "Logout user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Logout user",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/new": {
            "post": {
                "description": "Generate card preview",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cards"
                ],
                "summary": "Generate card preview",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.NewCardResponse"
                        }
                    }
                }
            }
        },
        "/system/health": {
            "get": {
                "description": "Get system health status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Get system health status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.HealthResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "Update user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    }
                }
            }
        },
        "/user/{user_id}": {
            "get": {
                "description": "Get user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/shared.HTTPError"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/email": {
            "put": {
                "description": "Update email for user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update email for user",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/password": {
            "put": {
                "description": "Update password for user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update password for user",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/schema.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.CardResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "encryptedData": {
                    "type": "string"
                },
                "encryptedKey": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "schema.HealthResponse": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "schema.NewCardResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "schema.PlainCardResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "schema.TokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "expiresIn": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "tokenType": {
                    "type": "string"
                }
            }
        },
        "schema.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "isConfirmed": {
                    "type": "boolean"
                },
                "provider": {
                    "type": "string"
                },
                "settings": {
                    "type": "object"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "shared.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "details": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
