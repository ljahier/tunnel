package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// function which return an available port
func getPort() int {
	// store in map or whatever all port defined and get the last port and increment to 1 and check if this port is available
	return 0
}

// func loadBalancer() {
// 	reverseProxy()
// }

func reverseProxy(ch chan string) {
	fmt.Println("hello1")
	fmt.Println("hello2")
	fmt.Println() // show new open port
	r := gin.Default()
	value := <-ch
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": value,
		})
	})
	r.Run(":9002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
