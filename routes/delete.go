package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/jacobh2/goatron/datatypes"
    "github.com/jacobh2/goatron/database"
)

func DeleteTodo(c *gin.Context) {
    var todo datatypes.Todo
    todoId := c.Param("id")
    db, _ := database.Database()
    db.First(&todo, todoId)

    if (todo.ID == 0) {
        c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
        return
    }

    db.Delete(&todo)
    db.Close()
    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo deleted successfully!"})
}

func DeleteAllTodo(c *gin.Context) {
    db, _ := database.Database()
    db.Where("").Delete(&datatypes.Todo{})
    db.Close()
    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "All todos deleted successfully!"})
}