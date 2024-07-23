package structures

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
