package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/me/gin-study/controller"
	"github.com/me/gin-study/middleware"
	"github.com/me/gin-study/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func init() {
	logWritterSetup()
}

func logWritterSetup() {
	f, _ := os.Create("gin-study.log")
	//gin.DefaultWriter = io.Writer(f)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()
	/*
		server := gin.Default() // will return an egine
		server.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "boom",
			})
		})
	*/
	// middlewares
	server.Use(gin.Recovery())
	// custom middleware
	server.Use(middleware.Logger(), middleware.BasicAuth())

	// debug using gindumb
	// server.Use(gindump.Dump())

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Video info saved",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		// rendering
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// server.Run(":8080")

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
