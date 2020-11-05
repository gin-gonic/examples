package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// version 1
	apiV1 := router.Group("/v1")

	apiV1.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V1 Users")
	})

	// User only can be added by authorized person
	authV1 := apiV1.Group("/", AuthMiddleWare())

	authV1.POST("users/add", AddV1User)

	//version 2
	apiV2 := router.Group("/v2")

	apiV2.GET("users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List Of V2 Users")
	})

	// User only can be added by authorized person
	authV2 := apiV2.Group("/", AuthMiddleWare())

	authV2.POST("users/add", AddV2User)

	_ = router.Run(":8081")

}

func AddV1User(c *gin.Context) {

	// AddUser

	c.JSON(http.StatusOK, "V1 User added")
}

func AddV2User(c *gin.Context) {

	// AddUser

	c.JSON(http.StatusOK, "V2 User added")
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {

		// here you can add your authentication method to authorize users.
		username := c.PostForm("user")
		password := c.PostForm("password")

		if username == "foo" && password == "bar" {
			return
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
