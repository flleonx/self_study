package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"gateway-server/database"

	configApp "gateway-server/config/app"
	configLogger "gateway-server/config/logger"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.uber.org/zap"
)

func loginHandler(c echo.Context) error {
	logger := configLogger.GetLogger()
	cfgApp := configApp.GetAppConfig()

	type Credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var credentials Credentials

	if err := c.Bind(&credentials); err != nil {
		logger.Error("Error parsing user credentials", zap.Error(err))
		return c.JSON(http.StatusUnauthorized, "invalid credentials")
	}

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("http://%s/login", cfgApp.AuthServiceAddress)
	body := []byte(`{}`)
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		logger.Error("Error building login request", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	r.Header.Add("Content-Type", "application/json")
	logger.Debug("user credentials", zap.String("email", credentials.Email), zap.String("password", credentials.Password))
	r.Header.Add("Authorization", fmt.Sprintf("%s:%s", credentials.Email, credentials.Password))
	resp, err := httpClient.Do(r)
	if err != nil {
		logger.Error("Error trying user authentication", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error trying to get the body from auth request", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	var token string

	err = json.Unmarshal(data, &token)

	if err != nil {
		logger.Error("Error trying to unmarshal the token", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	return c.JSON(http.StatusOK, token)
}

func uploadHandler(c echo.Context, claims claimResp) error {
	logger := configLogger.GetLogger()

	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("Error trying to get the file", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something wrong happened")
	}

	err = upload(file, claims)

	if err != nil {
		logger.Error("Error uploading the file", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "error uploading the file")
	}

	return c.JSON(http.StatusOK, "uploaded")
}

func downloadHandler(c echo.Context, claims claimResp) error {
	logger := configLogger.GetLogger()
	fidString := c.QueryParam("fid")

	if fidString == "" {
		return c.JSON(http.StatusBadRequest, "fid is missing")
	}

	mongoClient := database.Start()
	db := mongoClient.Database("mp3s")
	bucket, err := gridfs.NewBucket(db)

	fid, err := primitive.ObjectIDFromHex(fidString)
	if err != nil {
		logger.Error("Error trying to parse the hex to object id instance", zap.Error(err))
		return c.JSON(http.StatusBadRequest, "invalid fid")
	}

	fileBuffer := bytes.NewBuffer(nil)

	if _, err := bucket.DownloadToStream(fid, fileBuffer); err != nil {
		logger.Error("Error loading the file into a buffer", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something unexpected happened")
	}

	fileName := fmt.Sprintf("%s.mp3", fidString)

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.Response().Header().Set("Content-Type", "application/octet-stream")
	c.Response().Header().Set("Content-Length", strconv.Itoa(fileBuffer.Len()))

	// Write the file buffer to the response
	_, err = c.Response().Write(fileBuffer.Bytes())
	if err != nil {
		logger.Error("Error writing the file buffer in the response", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, "something unexpected happened")
	}

	return nil
}

func main() {
	logger := configLogger.GetLogger()
	cfgApp := configApp.GetAppConfig()

	e := echo.New()

	e.POST("/login", loginHandler)
	e.POST("/upload", middlewareAuth(uploadHandler))
	e.GET("/download", middlewareAuth(downloadHandler))

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
