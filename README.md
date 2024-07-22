# API Factory for Saas project

Factory API requires Go version 1.22 or higher to run.

## ⚙️ Installation

To start setting up your project. Create a new directory for your project and navigate into it. Then, initialize your project with Go modules by executing the following command in your terminal:

```shell
go mod init github.com/your/repo
```

To learn more about Go modules and how they work, you can check out the [Using Go Modules](https://go.dev/blog/using-go-modules) blog post.

After setting up your project, you can install Factory API with the `go get` command:

```shell
go get -u github.com/saas0503/factory-api
```

This command fetches the Factory Api package and adds it to your project's dependencies, allowing you to start building your web applications.



## ⚡️ Quickstart

```go
package main

import (
	api "github.com/saas0503/factory-api"
	"net/http"
)

func registry(r *api.App) {
	userRouter := r.AddGroup("user").Registry()
	userRouter.Get("", func(w http.ResponseWriter, r *http.Request) {
		api.JSON(w, api.ResponseOptions{
			Data: "users",
		})
	})
}

func main() {
	app := api.NewApp("api")

	registry(app)
	app.Listen(3000)
}
```

This simple server is easy to set up and run. It introduces the core concepts of API Factory: app initialization, route definition, and starting the server. Just run this Go program, and visit http://localhost:3000/api/user/ in your browser to see the message.