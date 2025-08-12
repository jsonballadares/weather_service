package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jasonballadares/weather_service/internal/env"
	"github.com/jasonballadares/weather_service/internal/model"
)

// ShortForecast is the main business logic that drives this application
// takes a latitude and longitude in the url parameters and returns a shortForecast and temperature characterization
// e.g., curl "http://localhost:8080/shortForecast?latitude=40.730610&longitude=-73.935242"
func ShortForecast(e *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := extractAndValidateRequest(w, r)

		if err != nil {
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}

		// First things first lets use the supplied lat & long to query /points to get the forecastURL
		pointsURL := fmt.Sprintf("%s/points/%s,%s", e.BaseUrlNWS, c.Latitude, c.Longitude)
		resp, err := http.Get(pointsURL)
		if err != nil {
			http.Error(w, "Error contacting National Weather Service API", http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		if resp.StatusCode != http.StatusOK {
			http.Error(w, fmt.Sprintf("National Weather Service API error: %s", resp.Status), resp.StatusCode)
			fmt.Println(fmt.Sprintf("ShortForecast: statusCode %d with status %s", resp.StatusCode, resp.Status))
			return
		}

		var points model.PointsResponse
		if err := json.NewDecoder(resp.Body).Decode(&points); err != nil {
			http.Error(w, "Error decoding points response", http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}

		// Now we have the forecast url so lets query it and get the data we need for today's forecast
		forecastResp, err := http.Get(points.Properties.Forecast)
		if err != nil {
			http.Error(w, "Error fetching forecast data", http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(forecastResp.Body)

		var forecast model.ForecastResponse
		if err := json.NewDecoder(forecastResp.Body).Decode(&forecast); err != nil {
			http.Error(w, "Error decoding forecast data", http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}

		// now we are almost done we have the metadata needed about the coordinates supplied by user
		var todayShortForecast string
		var temp float64

		// The NWS API returns either a Today or Tonight name depending upon when the request was made.
		// there is also no guarantee on the order of Name, so we must do a linear search to find the right data
		var matchedPeriod *model.Period
		for _, period := range forecast.Properties.Periods {
			if period.Name == "Today" || period.Name == "Tonight" {
				matchedPeriod = &period
				break // exit loop after finding the first match
			}
		}

		if matchedPeriod != nil {
			todayShortForecast = matchedPeriod.ShortForecast
			temp = matchedPeriod.Temperature
			name := matchedPeriod.Name
			tempUnit := matchedPeriod.TemperatureUnit
			fmt.Println(fmt.Sprintf("Name: %s\nTemperature: %f %v\nShort Forecast: %v", name, temp, tempUnit, todayShortForecast))
		} else {
			http.Error(w, "Error fetching forecast data. no Today/Tonight exists", http.StatusInternalServerError)
			return
		}

		// we have achieved our goal lets serve the response back to the client
		result := model.WeatherResponse{
			ShortForecast:               todayShortForecast,
			TemperatureCharacterization: temperatureCharacterization(temp),
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			http.Error(w, "Error returning forecast data", http.StatusInternalServerError)
			fmt.Println(fmt.Sprintf("ShortForecast: %v", err))
			return
		}
	}
}

// extractAndValidateRequest expects user to supply lat & lon like the below request
// curl "http://localhost:8080/shortForecast?latitude=40.730610&longitude=-73.935242"
func extractAndValidateRequest(w http.ResponseWriter, r *http.Request) (*model.Coordinates, error) {
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	// validate
	if latitude == "" {
		http.Error(w, "Please provide a latitude", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: client did not supply latitude")
	}

	if longitude == "" {
		http.Error(w, "Please provide a longitude", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: client did not supply longitude")
	}

	// Parse latitude
	latFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		http.Error(w, "Invalid latitude format", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: latitude is not a valid float")
	}
	if latFloat < -90 || latFloat > 90 {
		http.Error(w, "Latitude must be between -90 and 90", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: latitude out of range")
	}

	// Parse longitude
	lonFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		http.Error(w, "Invalid longitude format", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: longitude is not a valid float")
	}
	if lonFloat < -180 || lonFloat > 180 {
		http.Error(w, "Longitude must be between -180 and 180", http.StatusBadRequest)
		return nil, fmt.Errorf("validation error: longitude out of range")
	}

	// successfully validated
	coordinates := &model.Coordinates{
		Latitude:  latitude,
		Longitude: longitude,
	}

	return coordinates, nil
}

// temperatureCharacterization satisfies the assignment's requirement to "characterize" temperature of short forecast
func temperatureCharacterization(temp float64) string {
	// I live in Florida, so anything under 60 is cold anything above 90 is hot
	if temp <= 60 {
		return "cold"
	} else if temp >= 90 {
		return "hot"
	}
	return "moderate"
}
