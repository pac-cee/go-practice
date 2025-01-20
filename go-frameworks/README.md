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
    }
    c.JSON(http.StatusCreated, user)
}

func getUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"id": id})
}


2.Fiber framework 


package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Middleware
    app.Use(func(c *fiber.Ctx) error {
        c.Set("X-Custom-Header", "Hello, Fiber!")
        return c.Next()
    })

    // Routes
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "Hello, World!",
        })
    })

    // Route groups
    api := app.Group("/api")
    v1 := api.Group("/v1")

    v1.Post("/users", func(c *fiber.Ctx) error {
        var user User
        if err := c.BodyParser(&user); err != nil {
            return err
        }
        return c.Status(201).JSON(user)
    })

    app.Listen(":3000")
}



3.Buffalo framework


package actions

import (
    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/buffalo/middleware"
)

func App() *buffalo.App {
    app := buffalo.New(buffalo.Options{
        Env: "development",
    })

    app.Use(middleware.PopTransaction)
    app.Use(middleware.CSRF)

    app.GET("/", HomeHandler)

    api := app.Group("/api")
    api.GET("/users", UsersHandler)
    api.POST("/users", CreateUserHandler)

    return app
}

func HomeHandler(c buffalo.Context) error {
    return c.Render(200, r.JSON(map[string]string{
        "message": "Welcome to Buffalo!",
    }))
}


4.Echo framework


package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "net/http"
