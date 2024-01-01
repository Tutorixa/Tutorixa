package app

import (
	"app/internal/handlers/account"
	"app/internal/handlers/scheduler"
	"app/internal/pages"
	"app/internal/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	listenAddr string
	store storage.Storage
}

func NewApp(listenAddr string, store storage.Storage) *App {
	return &App{
		listenAddr: listenAddr,
		store: store,
	}
}

func (a *App) Start(e *echo.Echo ) error {
	pages.AddPagesRoutes(e)
	
	accountFunctions := accounthandler.New(a.store)
	schedulerFunctions := schedulerhandler.New(a.store)
	
	e.POST("/sign-in/verify", accountFunctions.Verification)
	e.GET("/schedule/date", schedulerFunctions.Schedule)
	
	e.Use(middleware.CORS())
	return e.Start(a.listenAddr)
}