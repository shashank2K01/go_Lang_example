package main

import (
	"io"
	"net/http"
	"os"

	"github.com/LopsidedPlace/ginexample/controller"
	"github.com/LopsidedPlace/ginexample/middlewares"
	"github.com/LopsidedPlace/ginexample/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "hello",
	// 	})
	// })

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())

		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {

			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"message": "video input valid"})
			}

		})

	}

	viewRoutes := server.Group("/view")
	{

		viewRoutes.GET("/videos", videoController.ShowAll)

	}
	/*
	   port := os.Getenv("PORT")
	   if port==""{
	   	port ="5000"
	   }
	   server.Run(":"+port)
	*/
	server.Run(":8085")
}
