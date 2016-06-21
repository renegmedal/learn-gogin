package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	admin := r.Group("/admin")

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")

	// or use this
	// r.StaticFS("/public", http.Dir("./public"))

	r.Run(":3000")

}
