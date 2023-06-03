package tronics

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Alternative using Echo
type ProductValidator struct {
	validator *validator.Validate
}

func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

var products = []map[int]string{{1: "mobiles"}, {2: "tv"}, {3: "laptops"}}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	var product map[int]string
	pID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	for _, p := range products {
		for k := range p {

			if pID == k {
				product = p
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
		// Vendor          string `json:"vendor" validate:"min=5,max=10"`
		// Email           string `json:"email" validate:"required_with=Vendor,email"`
		// Website         string `json:"website" validate:"url"`
		// Country         string `json:"country" validate:"len=2"`
		// DefaultDeviceIP string `json:"default_device_ip" validate:"ip"`
	}

	var reqBody body

	// e.Validator = &ProductValidator{validator: v}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}

	if err := v.Struct(reqBody); err != nil {
		return err
	}

	// if err := c.Validate(reqBody); err != nil {
	// 	return err
	// }

	product := map[int]string{
		len(products) + 1: reqBody.Name,
	}

	products = append(products, product)

	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	pID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	type body struct {
		Name string `json:"product_name" validate:"required,min=4"`
	}

	var reqBody body
	if err = c.Bind(&reqBody); err != nil {
		return err
	}

	if err = v.Struct(&reqBody); err != nil {
		return err
	}

	var updated bool

	for _, p := range products {
		if _, present := p[pID]; present {
			p[pID] = reqBody.Name
			updated = true
		}
	}

	if updated {
		return c.JSON(http.StatusOK, products)
	}

	return c.JSON(http.StatusNotFound, fmt.Sprintf("Product with id: %d does not exist", pID))
}

func deleteProduct(c echo.Context) error {
	var product map[int]string
	var index int

	pID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	for i, p := range products {
		for k := range p {

			if pID == k {
				product = p
				index = i
			}
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, "product not found")
	}

	splice := func(s []map[int]string, index int) []map[int]string {
		return append(s[:index], s[index+1:]...)
	}

	products = splice(products, index)

	return c.JSON(http.StatusOK, product)
}

func uriChanged(c echo.Context) error {
	fmt.Println("URI changed successfully")
	return c.JSON(http.StatusOK, fmt.Sprintf("URI changed to: %s", c.Request().URL))
}
