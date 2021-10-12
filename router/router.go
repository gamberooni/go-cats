package router

import (
	"github.com/gamberooni/go-cats/handler"
	"github.com/gamberooni/go-cats/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *echo.Echo {
	e := echo.New()

	// middlewares
	// Common Log Format - 127.0.0.1 PostmanRuntime/7.28.0 - [12/Oct/2021 17:34:41 +0800] "GET /api/cats/2 HTTP/1.1" 500 3
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${remote_ip} ${user_agent} - [${time_custom}] \"${method} ${path} ${protocol}\" ${status} ${bytes_out}" + "\n",
		CustomTimeFormat: "02/Jan/2006 15:04:05 -0700",
	}))
	e.Use(middleware.Recover())

	cs := store.NewCatStore(db)
	ch := handler.NewCatHandler(*cs)

	// create /api group
	apiGroup := e.Group("/api")

	apiGroup.GET("/cats", ch.GetAllCats)
	apiGroup.POST("/cats", ch.AddCat)
	apiGroup.GET("/cats/:id", ch.GetCatById)
	apiGroup.PUT("/cats/:id", ch.UpdateCatById)
	apiGroup.DELETE("/cats/:id", ch.DeleteCatById)

	return e
}
