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

	emissionDK1Filter := "&filter=%7B%22PriceArea%22:[%22DK1%22]%7D"
	// emissionDK2Filter := "&filter=%7B%22PriceArea%22:[%22DK2%22]%7D"
	emissionApiArguments := "?start=" + date + "&end=now%2BP1D&sort=Minutes5DK%20ASC" + emissionDK1Filter
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
	spinner, _ := pterm.DefaultSpinner.Start("Fetching & computing 	: ...")
	spinner.ShowTimer = true

	// Get the json data from the API
	emissionsData := structures.Emissions{}
	error := getJson(endpoint, &emissionsData)
	if error != nil {
		panic(error)
	}

	// Get just the latest emissions number & time
	latestEmissions, latestTime, _ := currentEmissions(&emissionsData)
	formattedLatestTime := latestTime[11:16]

	spinner.UpdateText(pterm.Sprintf("CO2 emissions: %v g/kWh @ "+formattedLatestTime, latestEmissions))
	spinner.Success()

	// ask user if they want to see the graph
	if userWantsGraph() {
		bars := getBars(&emissionsData)
		pterm.DefaultBarChart.WithHorizontal().WithWidth(16).WithBars(bars).WithShowValue().Render()
	}

}

func userWantsGraph() bool {
	options := []string{"No", "Yes"}

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithDefaultText("Do you want emissions data for all of today?").WithOptions(options).Show()

	if selectedOption == "Yes" {
		return true
	} else {
		return false
	}
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

func currentEmissions(data *structures.Emissions) (result float64, time string, err error) {
	// Get the last record in the data to get the current emissions
	lastRecord := data.Records[len(data.Records)-1]
	return lastRecord.CO2Emission, lastRecord.Minutes5DK, nil
}

func previousEmissions(data *structures.Emissions) (emis []float64, time []string, err error) {
	emis = make([]float64, len(data.Records))
	time = make([]string, len(data.Records))

	// TODO: Combine emissions to hourly averages

	for i, record := range data.Records {
		emis[i] = record.CO2Emission
		time[i] = record.Minutes5DK
	}

	return emis, time, nil
}

func getBars(data *structures.Emissions) []pterm.Bar {
	prevEmis, prevTime, _ := previousEmissions(data)

	reducedPrevEmis, reducedPrevTime := reduceData(&prevEmis, &prevTime)

	bars := []pterm.Bar{}
	for i := 0; i < len(reducedPrevTime); i++ {
		bars = append(bars, pterm.Bar{Label: reducedPrevTime[i], Value: int(reducedPrevEmis[i])})
	}

	return bars
}

// Reduce the data to hourly averages
func reduceData(emis *[]float64, time *[]string) (reducedEmis []float64, reducedTime []string) {
	currentHour := timeToHour((*time)[0])
	currentEmis := 0.0
	count := 0
	for i := 0; i < len(*time); i++ {
		if currentHour == timeToHour((*time)[i]) {
			count++
			currentEmis += (*emis)[i]
		} else {
			reducedEmis = append(reducedEmis, currentEmis/float64(count))
			reducedTime = append(reducedTime, timeToHour((*time)[i-1]))
			currentHour = timeToHour((*time)[i])
			count = 1
			currentEmis = (*emis)[i]
		}
	}

	reducedEmis = append(reducedEmis, currentEmis/float64(count))
	reducedTime = append(reducedTime, timeToHour((*time)[len(*time)-1]))

	return reducedEmis, reducedTime
}

// takes a timestamp string, returns the hour of the timestamp with an "h" appended
func timeToHour(time string) (hour string) {
	return time[11:13] + "h"
}

func avgEmissions(data *structures.Emissions) (result float64, err error) {
	panic("Not implemented yet")
}

func lowHighEmissions(data *structures.Emissions) (low float64, high float64, err error) {
	panic("Not implemented yet")
}
