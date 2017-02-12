package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
    "github.com/jacobh2/goatron/datatypes"
    "github.com/jacobh2/goatron/database"
)

func UpdateTodo(c *gin.Context) {
    var todo datatypes.Todo
    todoId := c.Param("id")
    db, _ := database.Database()
    db.First(&todo, todoId)

    if (todo.ID == 0) {
        c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
        return
    }

    db.Model(&todo).Update("title", c.PostForm("title"))
    // Convert string to int
    completed, _ := strconv.Atoi(c.PostForm("completed"))
    db.Model(&todo).Update("completed", completed)
    db.Close()
    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo updated successfully!"})
}