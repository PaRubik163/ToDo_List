package main

import (
	hand "todoapp/internal/handlers"
	data "todoapp/internal/database"

	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
    data.InitSqlite()

	r := gin.Default()
	r.LoadHTMLGlob("templates/tasks.html")
	r.Static("/static", "./templates/static")
	
	r.GET("/favicon.ico", func (c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"ICON": "icon"})
	})
	//На этом пути отображается страница
	r.GET("/tasks", hand.ShowTasksHTML)
	
	//На этих путях JS общается с Go
	api := r.Group("/api")
	{
		api.GET("/tasks", hand.GetTask)
		api.POST("/tasks", hand.CreateTask)
		api.PUT("/tasks/:id", hand.UpdateTask)
		api.DELETE("/tasks/:id", hand.DeleteTask)
	}

	r.Run("localhost:8060")
}
