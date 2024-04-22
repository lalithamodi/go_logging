package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	f, _ := os.Create("ginLogging.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router.GET("/getData", getData)
	router.Run()

}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hii its getdata method",
	})

}

//https://youtu.be/6tTDCKF1sB0?si=NMAfhFoUxdOIM83a
