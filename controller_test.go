package fapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

type UserController struct {
	Base        BaseController
	HelloWord   Handler `GET:"/hello" pagination:"true" guard:"apiKey"`
	UpdateWorld Handler `PATCH:"/hello/:id" pagination:"true" guard:"apiKey"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) error {
	response := map[string]string{
		"status":  "success",
		"message": "Welcome to Go standard library",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}
	return nil
}

func TestController(t *testing.T) {
	base := BaseController{
		Prefix:      "abc",
		Middlewares: []middleware{},
	}
	userController := &UserController{
		Base:      base,
		HelloWord: HelloWorldHandler,
	}

	abc := reflect.ValueOf(userController).Elem().FieldByName("Base").Interface().(BaseController)
	fmt.Println(&abc)
}
