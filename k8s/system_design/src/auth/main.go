package main

import (
	configApp "auth-server/config/app"
	configLogger "auth-server/config/logger"
	"auth-server/database"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func loginHandler(c echo.Context) error {
	logger := configLogger.GetLogger()

	auth := c.Request().Header.Get("Authorization")
	credentials := strings.Split(auth, ":")
	email := credentials[0]
	password := credentials[1]

	if auth == "" {
		logger.Error("User trying to access without credentials")
		return c.JSON(http.StatusUnauthorized, "missing credentials")
	}

	db := database.Start()

	type userTable struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user userTable

	row := db.QueryRow("SELECT email, password FROM users WHERE email = $1", email)

	err := row.Scan(&user.Email, &user.Password)

	if err != nil {
		logger.Error("Error trying to get the user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something happened trying to get the user")
	}

	if user.Email != email || user.Password != password {
		logger.Info("Invalid user or password", zap.String("email", email), zap.String("password", password))
		return c.JSON(http.StatusUnauthorized, "incorrect user or password")
	}

	token, err := generateToken(user.Email, true)

	if err != nil {
		logger.Error("Error generating the token", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	return c.JSON(http.StatusOK, token)
}

func validateHandler(c echo.Context, claims ClaimsJWT) error {
	type claimResp struct {
		Email string `json:"email"`
		Admin bool   `json:"admin"`
	}

	return c.JSON(http.StatusOK, claimResp{
		Email: claims.Email,
		Admin: claims.Admin,
	})
}

func main() {
	logger := configLogger.GetLogger()
	cfgApp := configApp.GetAppConfig()

	e := echo.New()

	e.POST("/login", loginHandler)
	e.POST("/validate", middlewareAuth(validateHandler))

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf("%s:%s", cfgApp.Host, cfgApp.Port)); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Unexpected error trying to shutdown gracefully", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer logger.Sync()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
