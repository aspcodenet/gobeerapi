package main

// NÄR VI GÖR  MANUELLA TESTER PÅ DET HÄR SÄTTET
// så AAA
// Arrange
// Act
// Assert - verifikation

// Arrange
// Act
// Assert - verifikation

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// calculate?location=KROGEN&age=20&promille=0.5

func calculateBeer(c *gin.Context) {
	location := c.Query("location")
	age, _ := strconv.Atoi(c.Query("age"))
	promille, _ := strconv.ParseFloat(c.Query("promille"), 32)

	canBuy, err := CanBuyBeer(location, age, float32(promille))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"canBuy": canBuy})
	}

}

func start(c *gin.Context) {
	computerName, _ := os.Hostname()
	c.String(http.StatusOK, "Tjena "+computerName+" this is cool")
}

func main() {
	router := gin.Default()
	router.GET("/", start)
	router.GET("/calculate", calculateBeer)

	router.Run(":8080")

}
