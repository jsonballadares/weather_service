package model

// API Documentation Referenced: https://www.weather.gov/documentation/services-web-api

// Created these structs to model the metadata we care to extract from the NWS API to accomplish our goal of
// returning the short forecast for an area for today

// Client is supplying lat / long which will be used to query the /points endpoint to get metadata we need to accomplish
// our goal defined above. There are many other endpoints we can use but for the sake of this assignment /points seems
// to be perfect for what we are trying to achieve

// PointsResponse /points/{lat},{lon} can be made into a much large struct housing more of what the payload returns
// but for the sake of this projects objective we will only capture the forecast endpoint so we can provide the shortForecast
// forecast is really the only property needed for the objective of this assignment as we can then make a query
// to the given URL to provide a shortForecast "Partly Cloudy", "Cloudy with a chance of meatballs", "Clear"
type PointsResponse struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

// ForecastResponse the meat and potatoes of this project as we need the shortForecast as well as the temperature of the
// first Period to Characterize it as hot, cold, moderate. The one that's named "Today" based on docs should be the first
// period returned and if not we can always do some search / sort to guarantee that but will be assuming it is
// the ForecastResponse.Properties.Periods[0] for the sake of simplicity
type ForecastResponse struct {
	Properties struct {
		Periods []struct {
			Name            string  `json:"name"`
			Temperature     float64 `json:"temperature"`
			TemperatureUnit string  `json:"temperatureUnit"`
			ShortForecast   string  `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

// WeatherResponse
type WeatherResponse struct {
	Forecast string `json:"forecast"`
	Category string `json:"temperature_category"`
}
