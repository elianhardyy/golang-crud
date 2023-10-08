package controllers

import (
	"server/models"
	"server/service"

	"github.com/gin-gonic/gin"
)

var err error

func GetAll(c *gin.Context) {
	var users []models.User
	err = service.Get(&users)
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, users)
	}
}

func Add(c *gin.Context) {
	var users models.User
	c.BindJSON(&users)
	err = service.Create(&users)
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, users)
	}
}

func Edit(c *gin.Context) {
	var users models.User
	id := c.Param("id")
	c.BindJSON(&users)
	err = service.Update(&users, id)
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, users)
	}
}

func Delete(c *gin.Context) {
	var users models.User
	id := c.Param("id")
	err := service.Destroy(&users, id)
	if err != nil {
		c.JSON(404, err)
	} else {
		c.JSON(200, users)
	}
}
