package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pterm/pterm"
)

type Emissions struct {
	Total   int
	Limit   int
	Dataset string
	Records []struct {
		// Minutes5UTC string
		Minutes5DK  string
		PriceArea   string
		CO2Emission float64
	}
}

func main() {
	// Emission stuff
	emissionUrl := "https://api.energidataservice.dk/dataset/CO2Emis"
	emissionApiArguments := "?end=now%2BP1D"
	emissionEndpoint := emissionUrl + emissionApiArguments
	// end of emission stuff
	// Price stuff
	priceUrl := "https://api.energidataservice.dk/dataset/DatahubPricelist"
	priceApiArguments := "?end=now%2BP1D&sort=ValidFrom%20DESC"
	priceEndpoint := priceUrl + priceApiArguments
	fmt.Print(priceEndpoint)
	// end of price stuff

	area, err := pterm.DefaultArea.WithCenter().Start("Getting electricity price...")
	if err != nil {
		fmt.Print("pterm area could not be created")
		panic(err)
	}

	emissionsData := Emissions{}
	error := getJson(emissionEndpoint, &emissionsData)
	if error != nil {
		panic(error)
	}

	area.Update(pterm.Sprintf("Electricity price: %v", emissionsData.Records[1].CO2Emission))
}

func getJson(endpoint string, target interface{}) (err error) {
	response, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error fetching data from API")
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("Error fetching data from API\nStatus code: %v", response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)

	return json.Unmarshal(responseBody, target)
}
