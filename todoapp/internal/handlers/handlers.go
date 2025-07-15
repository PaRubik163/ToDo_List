package handlers

import (
	"net/http"
	"strconv"
	data "todoapp/internal/database"
	"todoapp/internal/models"

	"github.com/gin-gonic/gin"
)
//of course i might write an auth but it won`t` be soon as fuck and i`m too lazy

//This is all about tasks 
func ShowTasksHTML(c *gin.Context){
	c.HTML(http.StatusOK, "tasks.html", nil)
}

func GetTask(c *gin.Context) {
	tasks, err := data.GetAll()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newtask models.Todo

	if err := c.Bind(&newtask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid input"})
		return 
	}

	err := data.InsertTasks(newtask)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newtask)
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := data.DeleteTaskFromDB(id)

	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Complete": "Task deleted"})
}


func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := data.UpdateDone(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Complete": "task update"})
}
