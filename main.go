package main

import (
	"strawjackal.org/yertle/webfinger"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET(".well-known/webfinger", webfinger.Handler)
	server.Run()

}
