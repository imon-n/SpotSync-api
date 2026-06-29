package middlewares

import (
	"goproject/internal/httpresponse"
	"net/http"

	"github.com/labstack/echo/v5"
)

func AdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c *echo.Context) error {

			role, ok := c.Get("user_role").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, httpresponse.Error{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				})
			}

			if role != "admin" {
				return c.JSON(http.StatusForbidden, httpresponse.Error{
					Code:    http.StatusForbidden,
					Message: "Admin access required",
				})
			}

			return next(c)
		}
	}
}