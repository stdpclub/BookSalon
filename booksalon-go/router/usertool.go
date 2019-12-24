package router

import (
	"BookSalon/booksalon-go/dbconn"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUser(c *gin.Context) {
	if users, err := dbconn.GetUsers(); err == nil {

		// TODO: how to change the style of output. which is define in db?
		type userShow struct {
			ID   uint
			Name string
		}

		var retUserShow []userShow
		for _, user := range users {
			retUserShow = append(retUserShow, userShow{user.ID, user.Name})
		}

		c.JSON(http.StatusOK, gin.H{
			"users": retUserShow,
		})
	}
}

func getUserInfo(c *gin.Context) {
	userid := c.Param("userid")

	if user, err := dbconn.GetUserObjByID(userid); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "search error! not found",
	})
}

func createUser(c *gin.Context) {
	var user dbconn.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad requset format",
		})
		return
	}

	if _, err := dbconn.CreateUser(&user); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "server create user error",
	})
}

func deleteUser(c *gin.Context) {

	userid := c.Param("userid")

	if _, err := dbconn.DelUser(userid); err == nil {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "user not found",
	})
}
