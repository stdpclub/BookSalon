package main

import (
	"crypto/md5"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initView(r *gin.Engine) {
	// r.LoadHTMLGlob("../web/templates/*")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"username": "none",
	// 	})
	// })

	r.POST("/login", handleLogin)

	userGroup := r.Group("/user", authRequired())
	{
		// 获取所有用户
		userGroup.GET("/", getAllUser)
		// 查询某个用户的信息
		userGroup.GET("/:userid", getUserInfo)
		// 新建一个user
		userGroup.POST("/", createUser)
		// 删除一个user
		userGroup.DELETE("/:userid", deleteUser)

		// 获取user用户的所有team
		userGroup.GET("/:userid/teams", getAllUserTeam)
		// 新建一个隶属于user的team
		userGroup.POST("/:userid/team", createUserTeam)
		// 获取user用户的某个team的信息
		userGroup.GET("/:userid/team/:teamid", getUserTeam)
		// 更新user下的team的信息
		userGroup.PUT("/:userid/team/:teamid", updateUserTeam)
		// 删除user下的某个team
		userGroup.DELETE("/:userid/team/:teamid", deleteUserTeam)

		// 获取user参加的team的leader
		userGroup.GET("/:userid/team/:teamid/leader", getTeamLeader)
		// 获取user参加的team的所有组员
		userGroup.GET("/:userid/team/:teamid/members", getAllMember)
		// 增加user下的某个team的组员
		userGroup.POST("/:userid/team/:teamid/member", createTeamMember)
		// 删除user下的某个team的某个组员
		userGroup.DELETE("/:userid/team/:teamid/member/:id", deleteTeamMember)
	}

}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if usersession, err := c.Cookie("user"); err == nil { // 获取成功
			if val, ok := userSessions[usersession]; ok == true { // 并且存在 // TODO:添加为redis
				c.Set("userName", val)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "you haven't login",
		})
		c.Abort()
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

	// TODO: don't use cookies, use token to control the login status. now setting into encrip string

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(user.Name))
	sessionID := md5Ctx.Sum(nil)
	userSessions[string(sessionID)] = user.Name

	c.SetCookie("user", string(sessionID), 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"loginState": "login success"})

}
