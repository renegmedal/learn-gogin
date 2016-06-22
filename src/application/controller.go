package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRoutes() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"TimesOff": timesOff,
			})
	})

	// added authentication
	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin", "test": "test",
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")

	return r
}
