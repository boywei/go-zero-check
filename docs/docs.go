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
        "/convert": {
            "post": {
                "tags": [
                    "转换方法"
                ],
                "summary": "前端传过来的JSON转Go对象",
                "parameters": [
                    {
                        "type": "string",
                        "description": "json file path",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/convert-lustre": {
            "post": {
                "tags": [
                    "Lustre方法"
                ],
                "summary": "Lustre转Uppaal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "file path",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-current-status": {
            "get": {
                "tags": [
                    "模拟器"
                ],
                "summary": "获取当前状态",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-global": {
            "get": {
                "tags": [
                    "模拟器"
                ],
                "summary": "获取各自动机的局部变量(对应模拟器中间的局部变量)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-next": {
            "get": {
                "tags": [
                    "模拟器"
                ],
                "summary": "获取使能迁移中所有可能的下一步",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-sync": {
            "get": {
                "tags": [
                    "模拟器"
                ],
                "summary": "获取同步情况(对应模拟器右下角的同步图)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/get-trace": {
            "get": {
                "tags": [
                    "模拟器"
                ],
                "summary": "获取模拟Trace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/next": {
            "post": {
                "tags": [
                    "模拟器"
                ],
                "summary": "使能迁移下一步",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/open-trace": {
            "post": {
                "tags": [
                    "模拟器"
                ],
                "summary": "打开模拟Trace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "trace path",
                        "name": "file",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/random-trace": {
            "post": {
                "tags": [
                    "模拟器"
                ],
                "summary": "随机模拟Trace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/reset": {
            "post": {
                "tags": [
                    "模拟器"
                ],
                "summary": "重置使能迁移, 复位",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/save-trace": {
            "post": {
                "tags": [
                    "模拟器"
                ],
                "summary": "保存模拟Trace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "model id",
                        "name": "id",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/verify": {
            "post": {
                "tags": [
                    "验证器"
                ],
                "summary": "验证某条性质是否成立",
                "parameters": [
                    {
                        "type": "string",
                        "description": "property",
                        "name": "property",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":\"200\",\"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
