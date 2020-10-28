package authmod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Controller is for auth logic
type Controller struct{}

type Reqe struct {
	PeerAddr     string
	PeerPort     int
	Mountpoint   string
	Client       string
	Username     string
	Password     string
	CleanSession bool
}

//Login is to process login request
func (auth *Controller) Auth(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	j := string(body)
	var t Reqe
	json.Unmarshal([]byte(j), &t)
	fmt.Print(t.Username)
	c.JSON(http.StatusOK, gin.H{"result": "ok"})

}
