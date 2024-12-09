// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/todoapp/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token if credentials are correct",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "JWT token",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "401": {
                        "description": "Wrong email or password",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/signup": {
            "post": {
                "description": "Creates new user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "User data for signup",
                        "name": "userdata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "email \u0026 password should not be empty",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Unexpected error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/tasks": {
            "get": {
                "description": "Fetches all tasks associated with the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks of a user",
                "responses": {
                    "200": {
                        "description": "List of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "UserId not found in context or server error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new task for the authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Creates a new task",
                "parameters": [
                    {
                        "description": "Task input data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TaskInputRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "New task created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload or error while creating task",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/tasks/view": {
            "get": {
                "description": "Fetches tasks with optional filtering by due date, priority, category, and status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "View tasks with filter options",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter tasks by due date (format: YYYY-MM-DD)",
                        "name": "due_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter tasks by priority (e.g., Low, Medium, High)",
                        "name": "priority",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter tasks by category",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter tasks by status (e.g., Pending, Completed)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Filtered list of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/tasks/{id}": {
            "put": {
                "description": "Updates only the provided fields of a specific task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Updates an existing task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task input data",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TaskInputRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload or task ID missing",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a specific task by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Deletes a task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Task ID missing or deletion failed",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/user-update": {
            "patch": {
                "description": "Updates an user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "User update",
                "parameters": [
                    {
                        "description": "User data for update",
                        "name": "userdata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success update",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "400": {
                        "description": "email \u0026 password should not be empty/at least 1 field required to update profile",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        },
        "/todoapp/users/all": {
            "get": {
                "description": "Fetches all users from db",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "List of all users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/helpers.Response"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/helpers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Task": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "priority": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "task_id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.TaskInputRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "due_date": {
                    "type": "string"
                },
                "priority": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "helpers.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error_message": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
