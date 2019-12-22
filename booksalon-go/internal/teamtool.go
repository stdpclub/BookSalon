package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUserTeamObj(c *gin.Context, user *User, team *Team) error {
	var teams []Team
	var err error

	if err = getUserObj(c, user); err != nil {
		return db.Error
	}

	// TODO: why return db.Error will be nil
	teamid := c.Param("teamid")
	if db.Model(user).Related(&teams, "Teams").First(team, "id = ?", teamid).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "team not found",
		})
	}

	return nil
}

func getAllUserTeam(c *gin.Context) {
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
		return
	}
	if err := tx.Model(&user).Association("Teams").Append(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "server create link error",
		})
		tx.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
	tx.Commit()
}

func getUserTeam(c *gin.Context) {
	var user User
	var team Team
	// var teams []Team

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func updateUserTeam(c *gin.Context) {
	var user User
	var team Team
	// var teams []Team
	// userid := c.Param("userid")
	// teamid := c.Param("teamid")

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request format error",
		})
		return
	}

	if err := getUserObj(c, &user); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	// TODO:there is a error. teamid useless, didn't change old team.
	if err := tx.Model(&user).Association("Teams").Replace(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update error",
		})
		tx.Rollback()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
	tx.Commit()
}

func deleteUserTeam(c *gin.Context) {
	var user User
	var team Team
	// var teams []Team

	if err := getUserTeamObj(c, &user, &team); err != nil {
		return
	}

	tx := db.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Model(&user).Association("Teams").Delete(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server delete error",
		})
		tx.Rollback()
		return
	}

	tx.Commit()
}
