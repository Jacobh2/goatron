package goatron

import (
    "github.com/gin-gonic/gin"
    "github.com/jacobh2/goatron/routes"
    "github.com/jacobh2/goatron/database"
    "github.com/jacobh2/goatron/datatypes"
)

func Service()  {

    //Migrate the schema
    db, _ := database.Database()
    defer db.Close()
    db.AutoMigrate(&datatypes.Todo{})
    
    router := gin.Default()

    v1 := router.Group("/api/v1/todos")
    {
        v1.POST("/", routes.CreateTodo)
        v1.GET("/", routes.FetchAllTodo)
        v1.GET("/:id", routes.FetchSingleTodo)
        v1.PUT("/:id", routes.UpdateTodo)
        v1.DELETE("/:id", routes.DeleteTodo)
        v1.DELETE("/", routes.DeleteAllTodo)
    }
    router.Run()
    
}