// Code generated by swaggo/swag. DO NOT EDIT
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
        "/Register": {
            "post": {
                "description": "用户注册",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_v1_user.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "仅限邮箱登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_handler_v1_user.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/:id": {
            "get": {
                "description": "Get an user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "通过用户id获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/follow": {
            "post": {
                "description": "Get an user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "通过用户id关注用户",
                "parameters": [
                    {
                        "description": "用户id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "仅限手机登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "phone",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PhoneLoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/unfollow": {
            "post": {
                "description": "Get an user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "通过用户id取消关注用户",
                "parameters": [
                    {
                        "description": "用户id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "Update a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "Update a user info by the user identifier",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserBaseModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}/followers": {
            "get": {
                "description": "Get an user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "通过用户id关注用户",
                "parameters": [
                    {
                        "description": "用户id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/users/{id}/following": {
            "get": {
                "description": "Get an user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "正在关注的用户列表",
                "parameters": [
                    {
                        "description": "用户id",
                        "name": "user_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用户信息",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/vcode": {
            "get": {
                "description": "Get an user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "根据手机号获取校验码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "区域码，比如86",
                        "name": "area_code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "details": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_handler_v1_user.LoginCredentials": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "internal_handler_v1_user.RegisterRequest": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserBaseModel": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string",
                    "maxLength": 128,
                    "minLength": 5
                },
                "phone": {
                    "type": "integer"
                },
                "sex": {
                    "type": "integer"
                },
                "username": {
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 1
                }
            }
        },
        "model.UserFollow": {
            "type": "object",
            "properties": {
                "fans_num": {
                    "description": "粉丝数",
                    "type": "integer"
                },
                "follow_num": {
                    "description": "关注数",
                    "type": "integer"
                },
                "is_fans": {
                    "description": "是否是粉丝 1:是 0:否",
                    "type": "integer"
                },
                "is_follow": {
                    "description": "是否关注 1:是 0:否",
                    "type": "integer"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "sex": {
                    "type": "integer"
                },
                "user_follow": {
                    "$ref": "#/definitions/model.UserFollow"
                },
                "username": {
                    "type": "string",
                    "example": "张三"
                }
            }
        },
        "user.PhoneLoginCredentials": {
            "type": "object",
            "required": [
                "phone",
                "verify_code"
            ],
            "properties": {
                "phone": {
                    "type": "integer",
                    "example": 13010002000
                },
                "verify_code": {
                    "type": "integer",
                    "example": 120110
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "eagle docs api",
	Description:      "eagle demo",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}