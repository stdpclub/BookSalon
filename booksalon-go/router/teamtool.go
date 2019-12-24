package router

import (
	"BookSalon/booksalon-go/dbconn"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUserTeam(c *gin.Context) {
	userid := c.Param("userid")

	if teams, err := dbconn.GetUserTeams(userid); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"teams": teams,
		})
	}
}

func createUserTeam(c *gin.Context) {
	userid := c.Param("userid")

	var team dbconn.Team

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	if _, err := dbconn.CreateTeam(userid, &team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server create team error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func getUserTeam(c *gin.Context) {
	userid := c.Param("userid")
	teamid := c.Param("teamid")

	if _, team, err := dbconn.GetUserTeamObj(userid, teamid); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"team": team,
		})
	}
	return
}

func updateUserTeam(c *gin.Context) {
	var team dbconn.Team
	var err error

	userid := c.Param("userid")
	teamid := c.Param("teamid")

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "request format error",
		})
		return
	}
	if _, err = dbconn.UpdateTeam(userid, teamid, &team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "update error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": team,
	})
}

func deleteUserTeam(c *gin.Context) {
	userid := c.Param("userid")
	teamid := c.Param("teamid")

	if err := dbconn.DelTeam(userid, teamid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "server delete error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"team": nil,
	})
}
