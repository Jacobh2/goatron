package goatron

import (
    "github.com/gin-gonic/gin"
    "github.com/jacobh2/goatron/routes"
    "github.com/jacobh2/goatron/database"
    "github.com/jacobh2/goatron/datatypes"
    "os"
)

func Service()  {

    //Migrate the schema
    db, err := database.Database()

    if(err != nil){
        println("No connection to MySQL found!")
        return
    }

    defer db.Close()
    db.AutoMigrate(&datatypes.Todo{})
    
    router := gin.Default()

    baseURL := os.Getenv("BASE_URL")

    if(baseURL == ""){
        baseURL = "/api/v1/todos"
    }

    v1 := router.Group(baseURL)
    {
        v1.POST("/", routes.CreateTodo)
        v1.GET("/", routes.FetchAllTodo)
        v1.GET("/:id", routes.FetchSingleTodo)
        v1.PUT("/:id", routes.UpdateTodo)
        v1.DELETE("/:id", routes.DeleteTodo)
        v1.DELETE("/", routes.DeleteAllTodo)
    }

    ADDR := os.Getenv("ADDRESS")
    
    if(ADDR == ""){
        ADDR = "127.0.0.1"
    }

    PORT := os.Getenv("PORT")

    if(PORT == ""){
        PORT = "8080"
    }

    println("Will start service on port " + PORT)
    println("With address " + ADDR)

    router.Run(ADDR + ":" + PORT)
    
}