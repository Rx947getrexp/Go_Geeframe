package main

import (
	"Gee/src/main/gee"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})

	r.LoadHtmlGlob("templates/*")
	r.Static("/asstes", "./static")

	stu1 := &student{Name: "keepfeelin", Age: 23}
	stu2 := &student{Name: "WZJ", Age: 22}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.hmpl ", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.hmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.Run(":9999")
}
