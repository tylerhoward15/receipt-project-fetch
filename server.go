package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"math"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var Receipts = make(map[string]Receipt)

type Id struct {
	Id string `json:"id"`
}

type Points struct {
	Points int `json:"points"`
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

func ruleOne(retailer *string) int {
	sum := 0
	// One pount for every alphanumeric character in the retailer name
	for _, char := range *retailer {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			sum++
		}
	}

	return sum
}

func ruleTwo(total *float64) int {
	// 50 points if the total is a round dollar amount with no cents
	if math.Mod(*total, 0.0) == 0 {
		return 50
	}

	return 0
}

func ruleThree(total *float64) int {
	// 25 points if the total is a multiple of .25
	if math.Mod(*total, 0.25) == 0 {
		return 25
	}

	return 0
}

func ruleFour(items *[]Item) int {
	// 5 points for every two items on the receipt
	return len(*items) / 2
}

func ruleFive(items *[]Item) int {
	// If the trimmed length of the item description is a multiple of 3, multiply the price by .2 and round up to the nearest integer. The result is the number of points earned
	sum := 0
	for item := range *items {
		if math.Mod(float64(len((*items)[item].ShortDescription)), 3) == 0 {
			price, err := strconv.ParseFloat((*items)[item].Price, 64)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			sum += int(math.Ceil(price * 0.2))
		}
	}
	return sum
}

func ruleSix(purchaseDay *int) int {
	// 6 points if the day in the purchase date is odd
	if *purchaseDay%2 == 1 {
		return 6
	}

	return 0
}

func ruleSeven(purchaseHour int) int {
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if purchaseHour > 14 && purchaseHour < 16 {
		return 10
	}

	return 0
}

func getPoints(c echo.Context) error {
	id := c.Param("id")
	if _, ok := Receipts[id]; !ok {
		return c.String(http.StatusNotFound, "Receipt not found")
	}

	sum := 0
	receipt := Receipts[id]

	date, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error parsing date")
	}
	_, _, day := date.Date()

	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error parsing total")
	}

	purchaseHour, err := strconv.Atoi(receipt.PurchaseTime[0:2])
	if err != nil {
		return c.String(http.StatusBadRequest, "Error parsing purchase time")
	}

	// rules
	sum += ruleOne(&receipt.Retailer)
	sum += ruleTwo(&total)
	sum += ruleThree(&total)
	sum += ruleFour(&receipt.Items)
	sum += ruleFive(&receipt.Items)
	sum += ruleSix(&day)
	sum += ruleSeven(purchaseHour)

	// this feels redundant and could be more elegant, but not a priority
	pointsResponse := new(Points)
	pointsResponse.Points = sum

	return c.JSON(http.StatusOK, pointsResponse)
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
