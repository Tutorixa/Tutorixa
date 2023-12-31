package main

import (
	"app/internal/app"
	"app/internal/storage/postgres"
	"app/internal/utils"
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	fmt.Println("Current Working Directory:")
	fmt.Println(os.Getwd())
	if utils.GetEnv() == "prod" || utils.GetEnv() == "production" {
		fmt.Println("WARNING: Running in production mode")
	}
	db, err := app.ConnectPgsql()

	if err != nil {
		fmt.Println(err)
	}
	postgresStorage := storage.NewPostgresStorage(db)

	build := flag.Bool("build", false, "ReBuild the database")
	if *build {
		fmt.Println("Building database")
		postgresStorage.BuildDevDB()
		return;
	}

	postgresStorage.BuildDevDB()
	server := app.NewApp("0.0.0.0:8000",postgresStorage) 
	err = server.Start(e)

	e.Logger.Fatal(err)
}
