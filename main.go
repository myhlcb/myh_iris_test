package main

import (
	"github.com/kataras/iris/v12"
)

type Person struct {
	Name string `json:"username"`
}

func main() {
	app := iris.New()
	// page get post
	app.Get("/", func(c iris.Context) {
		c.HTML("<h1>hi iris</h1>")
	})
	app.Get("/hi", func(c iris.Context) {
		c.JSON(&Person{Name: "zhangs"})
	})
	app.Run(iris.Addr(":8080"))
}
