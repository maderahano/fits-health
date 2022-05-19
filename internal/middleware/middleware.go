package middleware

import echo "github.com/labstack/echo/v4"

func APIKEYMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["X-Api-Key"] == nil {
			return echo.ErrUnauthorized
		}

		apiKey := c.Request().Header["X-Api-Key"][0]
		if apiKey != "ABC" {
			return echo.ErrUnauthorized
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
