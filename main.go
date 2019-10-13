package main
import (
      "github.com/gin-gonic/gin"
      "db"


     )

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    db.Init()
    r.Run()
    db.Close()
}
