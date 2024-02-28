package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var Receipts = make(map[string]Receipt)

type Id struct {
	Id string `json:"id"`
}

// Receipt struct)
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}

// Item struct

type Item struct {
	Price            string `json:"price"`
	ShortDescription string `json:"shortDescription"`
}

func main() {
	e := echo.New()

	e.POST("/receipts/process", processReceipt)
	e.GET("/receipts/:id/points", getPoints)

	// testing
	e.GET("/receipts", getReceipts)

	e.Logger.Fatal(e.Start(":1323"))
}

func getReceipts(c echo.Context) error {
	return c.JSON(http.StatusOK, Receipts)
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

	uuid := uuid.NewString()
	Receipts[uuid] = *receipt

	// this feels redundant and could be more elegant, but not a priority
	idResponse := new(Id)
	idResponse.Id = uuid

	return c.JSON(http.StatusOK, idResponse)
}
