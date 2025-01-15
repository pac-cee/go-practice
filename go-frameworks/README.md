1.GIN basics



.go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // REST API example
    r.POST("/users", createUser)
    r.GET("/users/:id", getUser)

    r.Run(":8080")
}

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
