// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},"swagger":"2.0","info":{"description":"{{escape .Description}}","title":"{{.Title}}","contact":{},"version":"{{.Version}}"},"host":"{{.Host}}","basePath":"{{.BasePath}}","paths":{"/":{"get":{"description":"get the welcome message","tags":["Base"],"responses":{"200":{"description":"OK","schema":{"type":"string"}}}}},"/todo":{"get":{"description":"get all todos","produces":["application/json"],"tags":["Todo"],"parameters":[{"type":"boolean","description":"Filter Undone Todos optionally with ` + "`" + `true` + "`" + `","name":"undone","in":"query"}],"responses":{"200":{"description":"OK","schema":{"$ref":"#/definitions/models.Todo"}}}},"post":{"description":"create a new todo","consumes":["application/json"],"produces":["application/json"],"tags":["Todo"],"parameters":[{"minLength":3,"type":"string","description":"Todo Title","name":"title","in":"query","required":true}],"responses":{"200":{"description":"OK","schema":{"type":"string"}},"400":{"description":"Bad Request","schema":{"type":"string"}}}}},"/todo/:id":{"get":{"description":"get todo with id","produces":["application/json"],"tags":["Todo"],"parameters":[{"type":"string","description":"Todo Id","name":"id","in":"path","required":true}],"responses":{"200":{"description":"OK","schema":{"$ref":"#/definitions/models.Todo"}},"400":{"description":"Bad Request","schema":{"type":"string"}},"404":{"description":"Not Found","schema":{"type":"string"}}}},"delete":{"description":"delete todo with id","produces":["application/json"],"tags":["Todo"],"parameters":[{"type":"string","description":"Todo Id","name":"id","in":"path","required":true}],"responses":{"204":{"description":"No Content"},"400":{"description":"Bad Request","schema":{"type":"string"}},"404":{"description":"Not Found","schema":{"type":"string"}}}}},"/todo/:id/done":{"patch":{"description":"mark todo done","produces":["application/json"],"tags":["Todo"],"parameters":[{"type":"string","description":"Todo Id","name":"id","in":"path","required":true}],"responses":{"200":{"description":"OK","schema":{"$ref":"#/definitions/models.Todo"}},"400":{"description":"Bad Request","schema":{"type":"string"}},"404":{"description":"Not Found","schema":{"type":"string"}}}}},"/version":{"get":{"description":"get the version","tags":["Base"],"responses":{"200":{"description":"OK","schema":{"type":"string"}}}}}},"definitions":{"models.Todo":{"type":"object","properties":{"createdAt":{"type":"string"},"done":{"type":"boolean"},"id":{"type":"string"},"title":{"type":"string"}}}}}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Simple Go Fiber App",
	Description:      "Simplest API written in go",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
