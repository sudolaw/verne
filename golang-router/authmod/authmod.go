package authmod

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller is for auth logic
type Controller struct{}

//Login is to process login request
func (auth *Controller) Auth(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"datdflkngkla": "hello world"})

}
