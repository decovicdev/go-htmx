package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

type Todo struct {
	Title string
	Done  bool
}

func main() {

	// todos := []Todo{
	// 	{Title: "Write presentation", Done: false},
	// 	{Title: "Plan trip to London", Done: true},
	// }

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "public")

	e.Renderer = newTemplates()

	e.GET("/css/output.css", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache, no-store, must-revalidate")
		return c.File("./public/css/output.css")
	})

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.POST("/", func(c echo.Context) error {

		fmt.Println(c.FormValue("title"))

		return c.Render(http.StatusOK, "todo", Todo{
			Title: c.FormValue("title"),
		})
	})

	e.Logger.Fatal(e.Start(":3000"))

	fmt.Println("Server started on port 3000")

}
