package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initView(r *gin.Engine) {
	r.LoadHTMLGlob("../web/templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": "none",
		})
	})

	r.POST("/login", handleLogin)

	// 获取所有用户
	r.GET("/users", getAllUser)
	// 查询某个用户的信息
	r.GET("/user/:userid", getUserInfo)
	// 新建一个user
	r.POST("/user", createUser)
	// 删除一个user
	r.DELETE("/user/:userid", deleteUser)

	userGroup := r.Group("/user")
	{
		// 获取user用户的所有team
		userGroup.GET("/user/:userid/teams", getAllUserTeam)
		// 新建一个隶属于user的team
		userGroup.POST("/user/:userid/team", createUserTeam)
		// 获取user用户的某个team的信息
		userGroup.GET("/user/:userid/team/:teamid", getUserTeam)
		// 更新user下的team的信息
		userGroup.PUT("/user/:userid/team/:teamid", updateUserTeam)
		// 删除user下的某个team
		userGroup.DELETE("/user/:userid/team/:teamid", deleteUserTeam)

		// 获取user参加的team的leader
		userGroup.GET("/user/:userid/team/:teamid/leader", getTeamLeader)
		// 获取user参加的team的所有组员
		userGroup.GET("/user/:userid/team/:teamid/members", getAllMember)
		// 增加user下的某个team的组员
		userGroup.POST("/user/:userid/team/:teamid/member", createTeamMember)
		// 删除user下的某个team的某个组员
		userGroup.DELETE("/user/:userid/team/:teamid/member/:id", deleteTeamMember)
	}

}

func handleLogin(c *gin.Context) {
	var loginInfo UserAccount

	if err := c.ShouldBindJSON(&loginInfo); err != nil { // JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": "format incomplete"})
		return
	}

	if db.First(&loginInfo).RecordNotFound() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// TODO: account relationship with user need to modify
	var user User
	if db.First(&user, "name = ?", loginInfo.Account).RecordNotFound() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "account search error"})
		return
	}

	// TODO: don't use cookies, use token to control the login status
	c.SetCookie("user", string(user.ID), 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"loginState": "login success"})

}
