package dto

type WeatherDto struct {
	Tempotary   float64 `json:"tempotary"`
	Description string  `json:"description"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float32 `json:"windSpeed"`
	Main        string  `json:"main"`
}
