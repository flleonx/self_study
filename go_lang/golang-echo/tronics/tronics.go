package tronics

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

// func changeURI(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		c.Request().URL.Path = "/changed"
// 		fmt.Printf("%+v\n", c.Request())
// 		return next(c)
// 	}
// }

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Inside middleware")
		return next(c)
	}
}

func Start() {
	// port := os.Getenv("MY_APP_PORT")

	// if port == "" {
	// 	port = "8080"
	// }

	// e.Pre(changeURI)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(serverMessage)

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct, middleware.BodyLimit("1K"))
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.GET("/changed", uriChanged)

	// e.Logger.Printf(fmt.Sprintf("Listening on port %s", port))
	// e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

	e.Logger.Printf(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
