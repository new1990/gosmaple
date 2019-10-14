package main
import (
      "github.com/new1990/gosmaple/server"
      "github.com/new1990/gosmaple/db"


     )

func main() {
    // r := gin.Default()
    // r.GET("/ping", func(c *gin.Context) {
    //     c.JSON(200, gin.H{
    //         "message": "pong",
    //     })
    // })
    db.Init()
    // r.Run()
    server.Init()
    db.Close()
}
