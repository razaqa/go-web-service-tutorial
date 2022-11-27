package main

import (
	"fmt"
    "net/http"
    "time"
	"github.com/gin-gonic/gin"
)

var (
	nama = "Dibimbing"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./statics")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "hello.html", map[string]interface{}{
            "waktu": formatAsDate(time.Now()),
			"nama": nama,
        })
    })
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}

func formatAsDate(t time.Time) string {
    return fmt.Sprintf("%02d-%02d-%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
}