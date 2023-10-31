package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	// "github.com/labstack/echo"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	port := os.Getenv("MY_APP_PORT")

	if port == "" {
		port = "8000"
	}

	e := echo.New()
	products := []map[int]string{{1: "Mobile"}, {2: "Laptop"}, {3: "tv"}, {4: "airpod"}}
	e.GET("/product/:id", func(c echo.Context) error {
		var product map[int]string
		for _, p := range products {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}
		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)

	})
	v := validator.New() //create variable of validator  
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name" validate:"required,min=4"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err
		}
		product := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, product)
		return c.JSON(http.StatusCreated, product)

	})
	//e.Logger.Printf(fmt.Sprintf("Listnig on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

}
