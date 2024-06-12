package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Price struct {
	Data struct {
		Amount string `json:"amount"`
	} `json:"data"`
}

var apiKey = "87c9fcc8-60bd-44ca-9b49-427c1c7c871c"

func main() {
	upThreshold := 70000.0
	downThreshold := 68000.0

	price, err := getCurrentPrice()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("Current Bitcoin price: %.2f\n", price)
	fmt.Println(checkBitcoinPrice(price, upThreshold, downThreshold))
}

func getCurrentPrice() (float64, error) {
	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/prices/spot?currency=USD", nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var price Price
	err = json.Unmarshal(body, &price)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(price.Data.Amount, 64)
}

func checkBitcoinPrice(price, upThreshold, downThreshold float64) string {
	if price > upThreshold {
		return "Bitcoin price is above the up threshold!"
	} else if price < downThreshold {
		return "Bitcoin price is below the down threshold!"
	} else {
		return "Bitcoin price is within the threshold range."
	}
}