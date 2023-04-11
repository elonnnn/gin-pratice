package v1

import "github.com/gin-gonic/gin"

func AddMember(c *gin.Context) {
	name := c.Query("name")
	price := c.DefaultQuery("price", "100")
	c.JSON(200, gin.H{
		"v1":    "AddMember",
		"name":  name,
		"price": price,
	})

}