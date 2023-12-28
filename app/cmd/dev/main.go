package main

import (
	"app/internal/connection"
	"app/internal/pages"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// Used to initialize the server
// func buildServer() {

// }

// Adds routes to the echo server
func addRoutes(e *echo.Echo) {
	pages.AddPagesRoutes(e)
}

// Add a function that checks for flags here
// func parseArgs() {

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())

	// Change this to use the env file and this doesn't work
	fmt.Println("connecting to database")
	db, err := connection.ConnectPgsql()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(db)
	// defer db.Close()

	fmt.Println("Adding routes")
	addRoutes(e)

	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start("0.0.0.0:8000"))
}
