package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func dbSearchUser(user UserAccount) bool {
	var usert UserAccount
	if db.Where("account = ? AND password = ?", user.Account, user.Password).First(&usert).RecordNotFound() {
		return false
	}

	return true
}

func checkUserState(c *gin.Context) {

}

func getAllUser(c *gin.Context) {
	checkUserState(c)

	var users []User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func getUserInfo(c *gin.Context) {
	checkUserState(c)

	var user User
	userid := c.Param("userid")
	if err := db.Where("id = ?", userid).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "search error! not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func createUser(c *gin.Context) {
	checkUserState(c)

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server create user error",
		})
		tx.Rollback()
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func deleteUser(c *gin.Context) {
	checkUserState(c)

	var user User
	userid := c.Param("userid")
	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if db.Where("id = ?", userid).Unscoped().Delete(user).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		tx.Rollback()
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{})
}
