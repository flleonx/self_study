package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	configApp "gateway-server/config/app"
	configLogger "gateway-server/config/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type claimResp struct {
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}

type authedHandler func(echo.Context, claimResp) error

func middlewareAuth(handler authedHandler) func(c echo.Context) error {
	return func(c echo.Context) error {
		logger := configLogger.GetLogger()
		cfgApp := configApp.GetAppConfig()

		token := c.Request().Header.Get("Authorization")

		if token == "" {
			logger.Error("User trying to access without credentials")
			return c.JSON(http.StatusUnauthorized, "missing credentials")
		}

		httpClient := http.Client{
			Timeout: 10 * time.Second,
		}

		url := fmt.Sprintf("http://%s/validate", cfgApp.AuthServiceAddress)
		body := []byte("")
		r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			logger.Error("Error creating validation request", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, "something wrong happened")
		}

		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		resp, err := httpClient.Do(r)
		if err != nil {
			logger.Error("Error trying to auth the user", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, "something wrong happened")
		}

		defer resp.Body.Close()

		data, err := io.ReadAll(resp.Body)

		logger.Debug("auth service response body", zap.ByteString("claims body", data))

		if err != nil {
			logger.Error("Error trying to get the body from auth request", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, "something wrong happened")
		}

		var claims claimResp

		err = json.Unmarshal(data, &claims)

		if err != nil {
			logger.Error("Error trying to unmarshal the claims", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, "something wrong happened")
		}

		if !claims.Admin {
			return c.JSON(http.StatusUnauthorized, "you're not an admin")
		}

		return handler(c, claims)
	}
}
