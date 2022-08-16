package controllers

import (
	"net/http"
	"strconv"

	"backend.api/config"
	"backend.api/models"
	"github.com/gin-gonic/gin"
)

type todoRequest struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func ListTodo(c *gin.Context) {
	var todoList []models.Todo
	config.DB.Find(&todoList)
	c.JSON(http.StatusOK, todoList)
}

func CreateTodo(c *gin.Context) {
	var payload todoRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{Author: payload.Author, Title: payload.Title, Description: payload.Description}
	config.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var payload todoRequest

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error format")
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := config.DB.Where("id = ?", todoId).First(&todo)

	if todoById.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.Author = payload.Author
	todo.Description = payload.Description
	todo.Title = payload.Title

	result := config.DB.Save(&todo)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Can not update todo"})
		return
	}

	var response todoResponse

	response.Author = todo.Author
	response.Description = todo.Description
	response.ID = todo.ID
	response.Title = todo.Title

	c.JSON(http.StatusOK, response)
}

func DeleteTodo(c *gin.Context) {
	todo := models.Todo{}

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "error format")
		return
	}

	remove := config.DB.Where("id = ?", todoId).Unscoped().Delete(&todo)

	if remove.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, remove.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "message": "Success to delete todo", "todoID": todoId})
}
