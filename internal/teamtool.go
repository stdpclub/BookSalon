package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUserTeam(c *gin.Context) {
	checkUserState(c)

	var teams []Team
	user := User{}
	userid := c.Param("userid")
	db.First(&user, "id = ?", userid)
	db.Model(&user).Related(&teams, "Teams")

	c.JSON(http.StatusOK, gin.H{
		"teams": teams,
	})
}

func createUserTeam(c *gin.Context) {
	checkUserState(c)

	userid := c.Param("userid")

	var team Team

	if err := c.ShouldBindJSON(&team); err != nil {
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

	if err := tx.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server create team error",
		})
		tx.Rollback()
		return
	}

	var user User
	if tx.First(&user, "id = ?", userid).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		tx.Rollback()
	}
	if err := tx.Model(&user).Association("Teams").Append(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "server create link error",
		})
		tx.Rollback()
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
	tx.Commit()
}

func getUserTeam(c *gin.Context) {
	checkUserState(c)

}

func updateUserTeam(c *gin.Context) { checkUserState(c) }

func deleteUserTeam(c *gin.Context) { checkUserState(c) }
