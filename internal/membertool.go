package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTeamLeader(c *gin.Context) {
	if _, err := checkUserState(c); err != nil { return}

	var user User
	var team Team
	var leader User

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	if err := getUserObjByID(c, team.LeaderID, &leader); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"leader": leader,
	})
}

func getAllMember(c *gin.Context) {
	var user User
	var team Team
	var members []User

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	if err := db.Model(&team).Related(&members, "Users").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server get team's users error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"members": members,
	})
	return
}

func createTeamMember(c *gin.Context) {
	if _, err := checkUserState(c); err != nil { return}
	var user User
	var team Team
	var newUser User

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	if db.First(&newUser).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the new user not found in user list",
		})
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// TODO: didn't add into teams
	if err := tx.Model(&team).Association("Users").Append(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server append new user into team error",
		})
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"new_user": newUser,
	})
	tx.Commit()
	return
}

func deleteTeamMember(c *gin.Context) {
	if _, err := checkUserState(c); err != nil { return}
	var user User
	var team Team
	var delUser User

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	if err := c.ShouldBindJSON(&delUser); err != nil {
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
	if err := tx.Model(&team).Association("Users").Delete(&delUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server delete user error",
		})
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"delete User": User{},
	})
	tx.Commit()
	return
}
