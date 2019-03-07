package source

type Response struct {
	HourlyForecast  []HourlyForecast `json:"hourly_forecast"`
	WeatherResponse WeatherResponse  `json:"response"`
}

type WeatherResponse struct {
	Error Error `json:"error"`
}

type HourlyForecast struct {
	FCTtime     FCTtime     `json:"FCTTIME"`
	Temperature Temperature `json:"temp"`
}

type FCTtime struct {
	Civil string `json:"civil"`
}

// temperature holds all the temperature data
type Temperature struct {
	Fahrenheit string `json:"english"`
}

type Error struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
