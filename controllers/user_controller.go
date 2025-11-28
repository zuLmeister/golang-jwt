package controllers

import (
	"net/http"

	"practice-golang/config"
	"practice-golang/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	output := make([]map[string]interface{}, 0, len(users))
	for _, u := range users {
		output = append(output, map[string]interface{}{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": output})
}
