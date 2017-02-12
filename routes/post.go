package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "github.com/jacobh2/goatron/datatypes"
    "github.com/jacobh2/goatron/database"
)

func CreateTodo(c *gin.Context) {
    // Convert string to int
    completed, _ := strconv.Atoi(c.PostForm("completed"))
    // Create new Todo struct with provided values
    todo := datatypes.Todo{Title: c.PostForm("title"), Completed: completed};
    //Get db connection
    db, _ := database.Database()
    // Make sure to close db connection on finish
    defer db.Close()
    // Save the todo
    db.Save(&todo)
    // Respond to client with http status
    c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Todo item created successfully!", "resourceId": todo.ID})
}