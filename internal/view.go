package main

import (
	"log"
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
	r.DELETE("/user/:userid", func(c *gin.Context) {})

	// 获取user用户的所有team
	r.GET("/user/:userid/teams", func(c *gin.Context) {})
	// 新建一个隶属于user的team
	r.POST("/user/:userid/team", func(c *gin.Context) {})
	// 获取user用户的某个team的信息
	r.GET("/user/:userid/team/:teamid", func(c *gin.Context) {})
	// 更新user下的team的信息
	r.PUT("/user/:userid/team/:teamid", func(c *gin.Context) {})
	// 删除user下的某个team
	r.DELETE("/user/:userid/team/:teamid", func(c *gin.Context) {})

	// 获取user参加的team的leader
	r.GET("/user/:userid/team/:teamid/leader", func(c *gin.Context) {})
	// 获取user参加的team的所有组员
	r.GET("/user/:userid/team/:teamid/members", func(c *gin.Context) {})
	// 增加user下的某个team的组员
	r.POST("/user/:userid/team/:teamid/member", func(c *gin.Context) {})
	// 删除user下的某个team的某个组员
	r.DELETE("/user/:userid/team/:teamid/member/:id", func(c *gin.Context) {})

}

func handleLogin(c *gin.Context) {
	var loginInfo UserAccount
	log.Println(c.PostForm("account"), c.PostForm("password"))
	if err := c.ShouldBindJSON(&loginInfo); err != nil { // JSON
		if err = c.ShouldBind(&loginInfo); err != nil { // FORM
			c.JSON(http.StatusBadRequest, gin.H{"loginState": "incomplete"})
			return
		}
	}

	if !dbSearchUser(loginInfo) { // return manage page
		c.JSON(http.StatusUnauthorized, gin.H{"loginState": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"loginState": "login success"})
}
