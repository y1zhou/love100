package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupRouter ...
func SetupRouter(cookieSecret string) *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte(cookieSecret))
	r.Use(sessions.Sessions("userSession", store))

	rUsers := r.Group("/users")
	{
		rUsers.POST("/signup", CreateUser)
		rUsers.POST("/signin", LoginUser)
		rUsers.GET("/", FetchAllUsers)
	}

	rUser := r.Group("/user")
	rUser.GET("/", CurrentUser)
	rUser.Use(AuthUser())
	{
		rUser.POST("/signout", LogoutUser)
		rUser.PUT("/update", UpdateUser)
		rUser.DELETE("/delete", DeleteUser)
	}

	rContents := r.Group("/contents")
	{
		rContents.GET("/", FetchAllContents)
	}

	rContent := r.Group("/content")
	rContent.Use(AuthUser())
	{
		rContent.POST("/", CreateContent)
		rContent.PUT("/", UpdateContentTitle)
		rContent.PUT("/status", UpdateContentStatus)
		rContent.DELETE("/", DeleteContent)
	}

	return r
}

// AuthUser Check if user session is logged in
func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := sessions.Default(c).Get("userID")
		if userID == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Need to login first.",
				"err": "Invalid session token.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
