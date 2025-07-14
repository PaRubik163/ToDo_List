package main

import (
	data "todoapp/internal/database"
	hand "todoapp/internal/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
    err := data.InitSqlite()
	if err != nil{
		logrus.Fatal(err)
	}

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
