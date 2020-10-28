package authmod

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller is for auth logic
type Controller struct{}

//Login is to process login request
func (auth *Controller) Auth(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	println(string(body))
	c.JSON(http.StatusOK, gin.H{"result": "ok"})

}
