package main

import (
	"fmt"
	"golang-sql/internal/auth"
	"golang-sql/internal/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authedHandler func(*CustomContext, database.User) error

func middlewareAuth(handler authedHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)

		apiKey, err := auth.GetAPIKey(c.Request().Header)

		if err != nil {
			return cc.JSON(http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
		}

		user, err := cc.DB.GetUserByAPIKey(c.Request().Context(), apiKey)

		if err != nil {
			return cc.JSON(http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
		}

		return handler(cc, user)
	}
}
