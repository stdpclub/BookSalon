package router

import (
	"BookSalon/booksalon-go/dbconn"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberIndex struct {
	MemberID string `json:"memberid" binding:"required"`
}

func getTeamLeader(c *gin.Context) {
	teamid := c.Param("teamid")

	if team, err := dbconn.GetTeamObjByID(teamid); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
	} else if leader, err := dbconn.GetUserObjByID(team.LeaderID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "get leader error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"leader": leader})
	}
}

func getAllMember(c *gin.Context) {
	userid := c.Param("userid")
	teamid := c.Param("teamid")
	members, err := dbconn.GetTeamMember(userid, teamid)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"members": members,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "server get team's users error" + err.Error(),
	})
	return
}

func createTeamMember(c *gin.Context) {
	var addUser memberIndex

	if err := c.ShouldBindJSON(&addUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	userid := c.Param("userid")
	teamid := c.Param("teamid")

	if newUser, err := dbconn.AddTeamMember(userid, teamid, addUser.MemberID); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"new_user": newUser,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "the new user not found in user list",
	})
	return
}

func deleteTeamMember(c *gin.Context) {
	// var user dbconn.User
	// var team dbconn.Team
	var delUser memberIndex

	if err := c.ShouldBindJSON(&delUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	userid := c.Param("userid")
	teamid := c.Param("teamid")

	if err := dbconn.DelTeamMember(userid, teamid, delUser.MemberID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server delete user error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"delete User": dbconn.User{},
	})
	return
}
