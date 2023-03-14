package main

import (
	"Gee/src/main/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"nihao":  "sss",
			"wanjia": "cccc",
		})
	})
	r.Run(":9999")
}
