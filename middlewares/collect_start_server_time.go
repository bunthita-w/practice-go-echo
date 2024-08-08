package middlewares

import (
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func CollectStartServer(store *sessions.CookieStore, startServerTime time.Time) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			session, _ := store.Get(c.Request(), "stat")
			if session.Values["startServerTime"] == nil {
				session.Values["startServerTime"] = startServerTime.Format(time.RFC3339)
				err := session.Save(c.Request(), c.Response())
				if err != nil {
					return err
				}
			}
			return next(c)
		}
	}
}
