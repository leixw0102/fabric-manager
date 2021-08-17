package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	name := c.Param("username")
	c.String(http.StatusOK, "%s wants to login", name)
}
func Singup(c *gin.Context) {
	name := c.Param("username")
	c.String(http.StatusOK, "%s wants to signup", name)
}
