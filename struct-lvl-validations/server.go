package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// User contains user information.
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For example, this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't not just do this outside of validator. Doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
func UserStructLevelValidation(sl validator.StructLevel) {
	//user := structLevel.CurrentStruct.Interface().(User)
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}

	// plus can to more, even with different tag than "fnameorlname"
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}

	route.POST("/user", validateUser)
	route.Run(":8085")
}

func validateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed!",
			"error":   err.Error(),
		})
	}
}
