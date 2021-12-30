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
        "/api/archive/categories": {
            "get": {
                "description": "Archive categories paginator/filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Archive categories paginator/filter",
                "operationId": "archive-categories",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "active",
                        "name": "active",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Pagination"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/api/archive/categories/restore/:id": {
            "put": {
                "description": "restore category from archive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Restore category",
                "operationId": "restore-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resources.CategoryResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/api/categories": {
            "get": {
                "description": "get all categories with pagination data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get all categories paginator",
                "operationId": "get-all-category-pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "active",
                        "name": "active",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Pagination"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            },
            "post": {
                "description": "create categories of books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Create category",
                "operationId": "create-category",
                "parameters": [
                    {
                        "description": "category info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/resources.CategoryResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/api/categories/:id": {
            "get": {
                "description": "get one category by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get one category",
                "operationId": "get-one-category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resources.CategoryResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            },
            "put": {
                "description": "update category of books",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Update category",
                "operationId": "update-category",
                "parameters": [
                    {
                        "description": "category info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.UpdateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resources.CategoryResource"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete category of books (soft)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Delete category (soft)",
                "operationId": "delete-category",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        },
        "/api/categories/list": {
            "get": {
                "description": "get all categories list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "Get all categories list",
                "operationId": "get-all-category-list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.getAllListsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/http.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "rows": {},
                "sort": {
                    "type": "string"
                },
                "totalPages": {
                    "type": "integer"
                },
                "totalRows": {
                    "type": "integer"
                }
            }
        },
        "http.getAllListsResponse": {
            "type": "object",
            "properties": {
                "rows": {}
            }
        },
        "http.response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "input.CreateCategory": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "desc": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "input.UpdateCategory": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "desc": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 256
                }
            }
        },
        "resources.CategoryResource": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AdminAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "UsersAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Version:     "1.0",
	Host:        "localhost:8060",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Library App API",
	Description: "API Server for Library project",
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
	swag.Register("swagger", &s{})
}
