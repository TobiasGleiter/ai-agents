package weather

import (
	"fmt"
)

type Parameters struct {
	Properties map[string]interface{} `json:"properties"`
}

type FunctionCallResponse interface {
	Process() string
}

type WeatherResponse struct {
	Location string
	Format   string
	Temperature string
}

func (wr WeatherResponse) Process() string {
	return fmt.Sprintf("\nCurrent weather for %s: Temperature is %s in %s format", wr.Location, wr.Temperature, wr.Format)
}

func getCurrentWeather(params Parameters) FunctionCallResponse {
	format, _ := params.Properties["format"].(string)
	location, _ := params.Properties["location"].(string)
	temperature := "68"

	return WeatherResponse{Location: location, Format: format, Temperature: temperature}
}