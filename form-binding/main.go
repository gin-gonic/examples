package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Booking struct {
	Name     string     `form:"name" binding:"required"`
	CheckIn  *time.Time `form:"check_in" time_format:"2006-01-02" binding:"required"`
	CheckOut *time.Time `form:"check_out" time_format:"2006-01-02"`
}

func main() {
	router := gin.Default()
	router.POST("/book", bookingHandler)
	_ = router.Run(":8080")
}

func bookingHandler(c *gin.Context) {
	var booking Booking
	if err := c.ShouldBind(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := gin.H{
		"message":  "Booking successful!",
		"name":     booking.Name,
		"check_in": booking.CheckIn,
	}

	if booking.CheckOut != nil {
		resp["check_out"] = booking.CheckOut
	}

	c.JSON(http.StatusOK, resp)
}
