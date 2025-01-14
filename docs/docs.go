// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Frolov Vladislav",
            "url": "https://hh.ru/resume/7b5e19efff0c43b3390039ed1f4e5a635a4558",
            "email": "vfrolov2004@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Удаляет песню из базы данных",
                "tags": [
                    "Songs"
                ],
                "summary": "Удалить песню",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное удаление",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Возвращает все песни с фильтрацией",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Список всех песен",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Название группы",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Начальная дата (yyyy-mm-dd)",
                        "name": "release_date_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Конечная дата (yyyy-mm-dd)",
                        "name": "release_date_end",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Номер страницы пагинации",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Название песни",
                        "name": "song",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.SongRequest"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка сервера",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет песню в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Создать новую песню",
                "parameters": [
                    {
                        "description": "Информация о песне",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SongRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Успешное создание",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Ошибка запроса",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/songs/{id}": {
            "get": {
                "description": "Возвращает песню по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Получить песню по ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SongRequest"
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Обновляет данные о песне в базе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Songs"
                ],
                "summary": "Обновить песню",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID песни",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Обновленная информация о песне",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SongRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное обновление",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Песня не найдена",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SongRequest": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Songs API",
	Description:      "API для управления песнями.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
