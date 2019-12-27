package router

import (
	"BookSalon/booksalon-go/dbconn"
	"crypto/md5"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userSessions = make(map[string]int, 10)

// InitView will create an gin engine and init APIs
func InitView() *gin.Engine {
	r := gin.Default()
	r.POST("/login", handleLogin) // 登录

	userGroup := r.Group("/user", authRequired())
	{
		userGroup.GET("/", getAllUser)           // 获取所有用户
		userGroup.GET("/:userid", getUserInfo)   // 查询某个用户的信息
		userGroup.POST("/", createUser)          // 新建一个user
		userGroup.DELETE("/:userid", deleteUser) // 删除一个user

		userObj := userGroup.Group("/:userid", authExact())
		{
			userObj.GET("/teams", getAllUserTeam)              // 获取user用户的所有team
			userObj.POST("/team", createUserTeam)              // 新建一个隶属于user的team
			userObj.GET("/team/:teamid", getUserTeam)          // 获取user用户的某个team的信息
			userObj.GET("/team/:teamid/leader", getTeamLeader) // 获取user参加的team的leader
			userObj.GET("/team/:teamid/members", getAllMember) // 获取user参加的team的所有组员

			teamObj := userObj.Group("/team/:teamid", authTeamExact())
			{
				teamObj.PUT("/", updateUserTeam)                // 更新user下的team的信息
				teamObj.DELETE("/", deleteUserTeam)             // 删除user下的某个team
				teamObj.POST("/member", createTeamMember)       // 增加user下的某个team的组员
				teamObj.DELETE("/member/:id", deleteTeamMember) // 删除user下的某个team的某个组员
			}
		}
	}

	return r
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if usersession, err := c.Cookie("user"); err == nil { // 获取成功
			if val, ok := userSessions[usersession]; ok == true { // 并且存在 // TODO:添加为redis
				c.Set("userid", val)
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

func authExact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, _ := strconv.Atoi(c.Param("userid"))
		if valid, ok := c.Get("userid"); ok && valid == userid {
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You have no access",
		})
		c.Abort()
	}
}

func authTeamExact() gin.HandlerFunc {
	return func(c *gin.Context) {
		userid := c.Param("userid")
		teamid := c.Param("teamid")

		if _, _, err := dbconn.GetUserTeamObj(userid, teamid); err == nil { // TODO: 太低效了，重复查了两次
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You have no right to modify",
		})
		c.Abort()
	}
}

func handleLogin(c *gin.Context) {
	var loginInfo dbconn.UserAccount

	if err := c.ShouldBindJSON(&loginInfo); err != nil { // JSON
		c.JSON(http.StatusBadRequest, gin.H{"error": "format incomplete"})
		return
	}

	if user, err := dbconn.GetUserByPwd(&loginInfo); err == nil {
		// TODO: don't use cookies, use token to control the login status. now setting into encrip string
		md5Ctx := md5.New()
		md5Ctx.Write([]byte(user.Name))
		sessionID := md5Ctx.Sum(nil)
		userSessions[string(sessionID)] = int(user.ID)

		c.SetCookie("user", string(sessionID), 3600, "/", "localhost", false, false)
		c.JSON(http.StatusOK, gin.H{"loginState": "login success"})
		return
	}

	c.SetCookie("user", "", 0, "/", "localhost", false, false) // login faild. clear cookies
	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
