package controllers

import (
	"MyGram/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) GetUsers(ctx *gin.Context) {
	var (
		users  []models.Users
		result gin.H
	)

	c.masterDB.Find(&users)
	if len(users) <= 0 {
		result = gin.H{
			"data": nil,
		}
	} else {
		result = gin.H{
			"data": users,
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *Controllers) CreateUsers(ctx *gin.Context) {
	var (
		Users  models.Users
		result gin.H
	)

	if err := ctx.ShouldBindJSON(&Users); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.masterDB.Debug().Create(&Users).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_key\"") {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Email already exists",
			})

			return
		}

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_username_key\"") {
			ctx.JSON(http.StatusConflict, gin.H{
				"error":   "Conflict",
				"message": "Username already exists",
			})

			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})

		return
	}

	result = gin.H{
		"age":      Users.Age,
		"email":    Users.Email,
		"id":       Users.ID,
		"username": Users.Username,
	}
	ctx.JSON(http.StatusCreated, result)
}
