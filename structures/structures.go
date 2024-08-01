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

type Prices struct {
	Total   int
	Sort    string
	Limit   int
	Dataset string
	Records []struct {
		ChargeOwner          string  `json:"ChargeOwner"`
		GLNNumber            string  `json:"GLN_Number"`
		ChargeType           string  `json:"ChargeType"`
		ChargeTypeCode       string  `json:"ChargeTypeCode"`
		Note                 string  `json:"Note"`
		Description          string  `json:"Description"`
		ValidFrom            string  `json:"ValidFrom"`
		ValidTo              string  `json:"ValidTo"`
		VATClass             string  `json:"VATClass"`
		Price1               float64 `json:"Price1"`
		Price2               float64 `json:"Price2"`
		Price3               float64 `json:"Price3"`
		Price4               float64 `json:"Price4"`
		Price5               float64 `json:"Price5"`
		Price6               float64 `json:"Price6"`
		Price7               float64 `json:"Price7"`
		Price8               float64 `json:"Price8"`
		Price9               float64 `json:"Price9"`
		Price10              float64 `json:"Price10"`
		Price11              float64 `json:"Price11"`
		Price12              float64 `json:"Price12"`
		Price13              float64 `json:"Price13"`
		Price14              float64 `json:"Price14"`
		Price15              float64 `json:"Price15"`
		Price16              float64 `json:"Price16"`
		Price17              float64 `json:"Price17"`
		Price18              float64 `json:"Price18"`
		Price19              float64 `json:"Price19"`
		Price20              float64 `json:"Price20"`
		Price21              float64 `json:"Price21"`
		Price22              float64 `json:"Price22"`
		Price23              float64 `json:"Price23"`
		Price24              float64 `json:"Price24"`
		TransparentInvoicing int     `json:"TransparentInvoicing"`
		TaxIndicator         int     `json:"TaxIndicator"`
		ResolutionDuration   string  `json:"ResolutionDuration"`
	}
}
