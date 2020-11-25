package models

// Weather stores information about the weather
type Weather struct {
	ID   int    `json:"id"`
	Main string `json:"main"`
	Desc string `json:"description"`
	Icon string `json:"icon"`
}

// MainResp is the main response of the API
type MainResp struct {
	Temp     float32 `json:"temp"`
	FLike    float32 `json:"feels_like"`
	TMin     float32 `json:"temp_min"`
	TMax     float32 `json:"temp_max"`
	Pressure float32 `json:"pressure"`
}

// Misc holds the Country
type Misc struct {
	Country string `json:"country"`
}

// Response is the whole API response
type Response struct {
	Weather []Weather `json:"weather"`
	Main    MainResp  `json:"main"`
	Misc    Misc      `json:"sys"`
}

// Preferences stores the preffered city of the user
type Preferences struct {
	City string `json:"city"`
}