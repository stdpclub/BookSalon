package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserObj(c *gin.Context, user *User) error {
	userid := c.Param("userid")
	return getUserObjByID(c, userid, user)
}

func getUserObjByID(c *gin.Context, userid string, user *User) error {
	if db.First(user, "id = ?", userid).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return db.Error
	}
	return nil
}

func checkUserState(c *gin.Context) (username string, err error) {
	if username, err = c.Cookie("user"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you haven't login",
		})
		return "", err
	}

	return username, nil
}

func getAllUser(c *gin.Context) {
	if _, err := checkUserState(c); err != nil {
		return
	}

	var users []User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func getUserInfo(c *gin.Context) {
	if _, err := checkUserState(c); err != nil {
		return
	}

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
	if _, err := checkUserState(c); err != nil {
		return
	}

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
	if _, err := checkUserState(c); err != nil {
		return
	}

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
