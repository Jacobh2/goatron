package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/jacobh2/goatron/datatypes"
    "github.com/jacobh2/goatron/database"
)

func FetchAllTodo(c *gin.Context) {
    var todos []datatypes.Todo
    var _todos []datatypes.TransformedTodo

    db, _ := database.Database()

    // Get all todos, and put result in todos array
    db.Find(&todos)
    // Directly close db connection
    db.Close()

    if (len(todos) == 0) {
        c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
        return
    }

    //transforms the todos for building a good response
    for _, item := range todos {
            completed := item.Completed != 0
            _todos = append(_todos, datatypes.TransformedTodo{ID: item.ID, Title:item.Title, Completed: completed})
    }
    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _todos})
}

func FetchSingleTodo(c *gin.Context) {
    var todo datatypes.Todo
    todoId := c.Param("id")

    db, _ := database.Database()
    db.First(&todo, todoId)
    db.Close()

    if (todo.ID == 0) {
        c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
        return
    }

    completed := todo.Completed != 0

    _todo := datatypes.TransformedTodo{ID: todo.ID, Title:todo.Title, Completed: completed}

    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _todo})
}