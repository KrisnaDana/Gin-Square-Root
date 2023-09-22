package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
	router := gin.Default()
	router.POST("/square-root", func(c *gin.Context) {
		input, _ := strconv.ParseFloat(c.PostForm("input"), 64)
        var hasil float64 = input/2
        var tebak float64 = 0
        var toleransi float64 = 0.000000001

		for abs(hasil-tebak) > toleransi {
			tebak = hasil
			hasil = 0.5 * (hasil + (input / hasil))
		}

		c.JSON(200, gin.H{
			"hasil": hasil,
		})
	})
	router.Run(":8080")
}