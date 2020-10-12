package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"graphsProject/cmd/graphs/apis"
	"graphsProject/cmd/graphs/config"
	"log"
	"net/http"
)

var router *gin.Engine

func main() {

	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			})
	})

	server, err := apis.GraphApi()

	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}

	go server.Serve()
	defer server.Close()

	router.GET("/socket.io/", gin.WrapH(server))
	router.POST("/socket.io/", gin.WrapH(server))

	router.Run()

}
