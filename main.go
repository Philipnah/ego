package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pterm/pterm"
)

func main() {
	area, err := pterm.DefaultArea.WithCenter().Start("Getting electricity price...")
	if err != nil {
		fmt.Print("pterm area could not be created")
		panic(err)
	}

	currentPrice, _ := getPrice()

	area.Update(pterm.Sprintf("Electricity price: %s", currentPrice))
}

func getPrice() (result string, err error) {
	datasetUrl := "https://api.energidataservice.dk/dataset/DatahubPricelist"
	apiArguments := "?end=now%2BP1D&sort=ValidFrom%20DESC"
	apiEndpoint := datasetUrl + apiArguments

	response, err := http.Get(apiEndpoint)

	if err != nil {
		fmt.Println("Error fetching data from API")
		return "", err
	}

	if response.StatusCode != 200 {
		fmt.Println("Error fetching data from API")
		fmt.Println("Status code: %i", response.StatusCode)
	}

	data, _ := io.ReadAll(response.Body)

	return string(data), nil
}
