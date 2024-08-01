package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	structures "github.com/Philipnah/ego/structures"

	"github.com/pterm/pterm"
)

func main() {
	// Emission stuff
	emissionUrl := "https://api.energidataservice.dk/dataset/CO2Emis"

	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")

	emissionApiArguments := "?start=" + date + "&end=now%2BP1D&sort=Minutes5DK%20ASC"
	emissionEndpoint := emissionUrl + emissionApiArguments
	// end of emission stuff
	// Price stuff
	// priceUrl := "https://api.energidataservice.dk/dataset/DatahubPricelist"
	// priceApiArguments := "?end=now%2BP1D&sort=ValidFrom%20DESC"
	// priceEndpoint := priceUrl + priceApiArguments
	// end of price stuff

	pterm.DefaultSection.Println("CO2 Emissions")

	loadEmissions(emissionEndpoint)
}

func loadEmissions(endpoint string) {
	spinner, _ := pterm.DefaultSpinner.Start("Current CO2 emissions: ...")
	spinner.ShowTimer = true

	// Get the json data from the API
	emissionsData := structures.Emissions{}
	error := getJson(endpoint, &emissionsData)
	if error != nil {
		panic(error)
	}

	// Get just the emissions number right now
	currentEmissions, _ := currentEmissions(&emissionsData)

	spinner.UpdateText(pterm.Sprintf("Current CO2 emissions: %v g/kWh", currentEmissions))
	spinner.Success()

	bars := getBars(&emissionsData)

	pterm.DefaultBarChart.WithHorizontal().WithWidth(15).WithBars(bars).WithShowValue().Render()

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

func currentEmissions(data *structures.Emissions) (result float64, err error) {
	// Get the last record in the data to get the current emissions
	return data.Records[len(data.Records)-1].CO2Emission, nil
}

func previousEmissions(data *structures.Emissions) (emis []float64, time []string, err error) {
	emis = make([]float64, len(data.Records))
	time = make([]string, len(data.Records))

	// TODO: Get only DK1 data
	// TODO: Combine emissions to hourly averages

	for i, record := range data.Records {
		emis[i] = record.CO2Emission
		time[i] = record.Minutes5DK
	}

	return emis, time, nil
}

func getBars(data *structures.Emissions) []pterm.Bar {
	prevEmis, prevTime, _ := previousEmissions(data)

	bars := []pterm.Bar{}
	for i := 0; i < len(prevEmis); i++ {
		bars = append(bars, pterm.Bar{Label: prevTime[i], Value: int(prevEmis[i])})
	}

	return bars
}

func avgEmissions(data *structures.Emissions) (result float64, err error) {
	panic("Not implemented yet")
}

func lowHighEmissions(data *structures.Emissions) (low float64, high float64, err error) {
	panic("Not implemented yet")
}
