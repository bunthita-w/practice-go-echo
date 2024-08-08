package main

import (
	"log"
	"os"
	"practice-go-echo/middlewares"
	"practice-go-echo/routes"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var startServerTime time.Time
var store *sessions.CookieStore

func init() {
	startServerTime = time.Now()
}

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := []byte(os.Getenv("SESSION_KEY"))
	store = sessions.NewCookieStore(secret)
	store.Options = &sessions.Options{
		Path: "/",
	}

	e.Use(middlewares.CollectStartServer(store, startServerTime))
	e.Use(middlewares.CountVisitingTime(store))
	routes.Routes(e, store)

	e.Logger.Fatal(e.Start(":1323"))
}
