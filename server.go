package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

type Item struct {
	Price            string `json:"price"`
	ShortDescription string `json:"shortDescription"`
}

func main() {
	e := echo.New()

	e.POST("/receipts/process", processReceipt)
	e.GET("/receipts/:id/points", getPoints)

	e.Logger.Fatal(e.Start(":1323"))
}

func getPoints(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Getting points for receipt with id: "+id)
}

func processReceipt(c echo.Context) error {
	receipt := new(Receipt)
	if err := c.Bind(receipt); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, receipt)
	// or
	// return c.XML(http.StatusCreated, u)
}
