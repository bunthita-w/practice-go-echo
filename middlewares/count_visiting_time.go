package middlewares

import (
	"strings"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func CountVisitingTime(store *sessions.CookieStore) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiPath := strings.TrimSpace(c.Path())

			session, _ := store.Get(c.Request(), "stat")
			count, ok := session.Values[apiPath].(int)
			if !ok {
				count = 0
			}

			count++

			session.Values[apiPath] = count

			// fmt.Println("API Path:", apiPath)
			// fmt.Println("Session Count:", session.Values[apiPath])
			// fmt.Println("------------------------------------****")
			// count3, ok := session.Values["/calculator"].(int)
			// if !ok {
			// 	count3 = 0
			// }
			// count2, ok := session.Values["/"].(int)
			// if !ok {
			// 	count2 = 0
			// }

			// fmt.Println("Count2:", count2)
			// fmt.Println("Count3:", count3)
			// fmt.Println("------------------------------------****")

			err := session.Save(c.Request(), c.Response())
			if err != nil {
				return err
			}
			return next(c)
		}
	}
}
