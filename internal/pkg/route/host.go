package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Host(r *gin.Engine){
	r.GET("./", func(c *gin.Context){
		c.String(http.StatusOK, "hello world")
	})
}
