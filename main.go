package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
)

func main() {
	r := iris.New()
	{
		users := r.Party("/users")

		users.Get("/", func(ctx iris.Context) {
			id := ctx.URLParam("id")
			fmt.Println(id)
			ctx.JSON(iris.Map{"code": 200, "result": iris.Map{"id": id}})
		})

		users.Get("/{id:int}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			fmt.Println(id)
			ctx.JSON(iris.Map{"code": 200, "result": iris.Map{"id": id}})
		})
	}

	log.Panic(r.Run(iris.Addr(":8080")))

}
