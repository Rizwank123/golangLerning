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
	e.GET("/products/:id", func(c echo.Context) error {
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
	e.PUT("/products/:id", func(c echo.Context) error {
		var product map[int]string
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, "ID parameter is missing")
		}
		pID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error converting parameter to integer:", err)
			return err
		}
		fmt.Println("i am here ")
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
		type body struct {
			Name string `json:"product_name" "validate: required,min=4"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		if err := v.Struct(reqBody); err != nil {
			return err

		}
		product[pID] = reqBody.Name
		return c.JSON(http.StatusAccepted, product)
	})
	e.DELETE("/products/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, "ID parameter is missing")
		}
		pID, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error converting parameter to integer:", err)
			return err
		}
		var toDelete int
		for i, p := range products {
			for k := range p {
				if pID == k {
					toDelete = i
					break
				}
			}
		}
		if toDelete == 0 {
			return c.JSON(http.StatusNotFound, "id Not found")
		}
		pro := fmt.Sprintf("%v", products[toDelete])
		products = append(products[:toDelete], products[toDelete+1:]...)
		return c.JSON(http.StatusOK, pro)

	})
	//e.Logger.Printf(fmt.Sprintf("Listnig on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

}
   