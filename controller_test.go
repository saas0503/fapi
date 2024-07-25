package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

type UserController struct {
	HelloWord   Handler `GET:"/hello" pagination:"true" guard:"apiKey"`
	UpdateWorld Handler `PATCH:"/hello/:id" pagination:"true" guard:"apiKey"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "success",
		"message": "Welcome to Go standard library",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func TestController(t *testing.T) {
	userController := &UserController{
		HelloWord: HelloWorldHandler,
	}

	c := NewController("api")
	c.Registry(userController)
	fmt.Println(c.mux)
}
