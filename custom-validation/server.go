package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10" // 使用这个版本，Booing的验证关键字是binding，不可以使用validate，会导致bug
	//"gopkg.in/go-playground/validator.v9" //使用这个版本，Booking的验证关键字是validate，不要使用binding
	"net/http"
	"time"
)

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 对于Booking的验证，要求check_in和check_out必须大于当前日期
// 若当前时间为2023-02-27，则check_in=2023-02-28，check_out=2023-03-01
var bookabledate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	r := gin.Default()

	//validate := validator.New()
	//validate.RegisterValidation("bookabledate", bookabledate2)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookabledate2)
	}

	//r.GET("/bookable", func(c *gin.Context) {
	//	var book Booking
	//	if err := c.ShouldBind(&book); err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"error": err.Error(),
	//		})
	//		c.Abort()
	//		return
	//	}
	//	if err := validate.Struct(book); err != nil {
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"error": err.Error(),
	//		})
	//		c.Abort()
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "OK",
	//		"booking": book,
	//	})
	//})

	r.GET("/bookable", getBookable)
	r.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "booking": b})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func bookabledate2(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}

	return false
}
