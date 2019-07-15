package main

import (
	"net/http"

	gintemplater "github.com/firmanmm/gin-templater"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	config := gintemplater.NewConfig()
	templater := gintemplater.NewTemplater(engine, config)
	templater.Run()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.html", gin.H{
			"message": "Hello!",
			"who":     "A Message",
		})
	})

	engine.GET("/deep", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "home.deep.html", gin.H{
			"message": "Hello!",
			"who":     "A Deeper Message",
		})
	})

	engine.Run()
}
