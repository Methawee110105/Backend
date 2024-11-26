package employee

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee GET Method!",
    })
}

func POST(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee POST Method!",
    })
}

func PUT(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee PUT Method!",
    })
}

func DELETE(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
    "message": "Employee DELETE Method!",
    })
}
func GetEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func POSTEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func PUTEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
func DELETEEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
    "message": id,
    })
}
