package main

import (
	"github.com/thinkong/gnassign"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	blackList := gnassign.CheckList{}
	err := blackList.ReadConf("blacklist.conf")

	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.GET("/download_image", func(c *gin.Context) {
		requrl := c.Query("url")
		matched := blackList.SearchMatch(requrl)
		if matched {
			c.Status(http.StatusForbidden)
		} else {
			c.Status(http.StatusOK)
		}
	})
	r.Run(":8080")
}