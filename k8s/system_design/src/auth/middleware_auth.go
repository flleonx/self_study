package main

import (
	configApp "auth-server/config/app"
	configLogger "auth-server/config/logger"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type ClaimsJWT struct {
	Email string `json:"email"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func generateToken(email string, authz bool) (string, error) {
	cfgApp := configApp.GetAppConfig()

	expirationTime := time.Now().Add(2 * time.Hour)

	claims := &ClaimsJWT{
		Email: email,
		Admin: authz,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfgApp.JwtKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

type authedHandler func(echo.Context, ClaimsJWT) error

func middlewareAuth(handler authedHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		logger := configLogger.GetLogger()
		cfgApp := configApp.GetAppConfig()

		encodedJwt := c.Request().Header.Get("Authorization")

		if encodedJwt == "" {
			return c.JSON(http.StatusUnauthorized, "missing credentials")
		}

		encodedJwt = strings.Split(encodedJwt, " ")[1]

		claims := &ClaimsJWT{}
		decoded, err := jwt.ParseWithClaims(encodedJwt, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfgApp.JwtKey), nil
		})

		if err != nil {
			logger.Error("Error decoding the jwt token", zap.Error(err))
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, "unauthorized")
			}
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		if !decoded.Valid {
			logger.Error("Invalid token", zap.String("token", encodedJwt))
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		return handler(c, *claims)
	}
}
