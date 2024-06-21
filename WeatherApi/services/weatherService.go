package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test-WeatherApi/database"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test-WeatherApi/models"
)

const (
	baseURL = "http://api.weatherapi.com/v1/history.json"
	apiKey  = ""
)

func FetchAndStoreWeatherData(userID primitive.ObjectID, lat, lon float64, signUpDate time.Time) error {
	currentDate := time.Now()
	startDate := signUpDate

	// Iterate from signUpDate to currentDate, month by month
	for currentDate.After(startDate) || currentDate.Equal(startDate) {
		// Calculate the start and end date for the current month
		endOfMonth := time.Date(currentDate.Year(), currentDate.Month(), 1, 0, 0, 0, 0, time.UTC).AddDate(0, 1, -1)
		startOfMonth := time.Date(endOfMonth.Year(), endOfMonth.Month(), 1, 0, 0, 0, 0, time.UTC)

		// Construct the API request URL dynamically
		url := fmt.Sprintf("%s?key=%s&q=%f,%f&dt=%s&end_dt=%s", baseURL, apiKey, lat, lon, startOfMonth.Format("2006-01-02"), endOfMonth.Format("2006-01-02"))

		// Make API request to fetch weather data
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("error fetching weather data:", err)
			return err
		}
		defer resp.Body.Close()

		// Parse the response and extract forecast data
		var weatherResponse struct {
			Location struct {
				Name    string  `json:"name"`
				Region  string  `json:"region"`
				Country string  `json:"country"`
				Lat     float64 `json:"lat"`
				Lon     float64 `json:"lon"`
			} `json:"location"`
			Forecast struct {
				ForecastDay []struct {
					Date string `json:"date"`
					Day  struct {
						AvgTempC float64 `json:"avgtemp_c"`
					} `json:"day"`
				} `json:"forecastday"`
			} `json:"forecast"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
			fmt.Println("error decoding weather response:", err)
			return err
		}

		// Create a WeatherRecord for the current month
		weatherRecord := models.WeatherRecord{
			UserID: userID,
			Location: models.Location{
				Name:    weatherResponse.Location.Name,
				Region:  weatherResponse.Location.Region,
				Country: weatherResponse.Location.Country,
				Lat:     weatherResponse.Location.Lat,
				Lon:     weatherResponse.Location.Lon,
			},
			Month: models.Forecast{
				ForecastDay: []models.Day{}, // Initialize empty slice
			},
			Years: models.Forecast{
				ForecastDay: []models.Day{}, // Initialize empty slice
			},
			CreatedAt: time.Now(),
		}

		// Extract and populate forecast data for the current month
		for _, dayForecast := range weatherResponse.Forecast.ForecastDay {
			day := models.Day{
				Date:     dayForecast.Date,
				AvgTempC: dayForecast.Day.AvgTempC,
			}
			weatherRecord.Month.ForecastDay = append(weatherRecord.Month.ForecastDay, day)
			weatherRecord.Years.ForecastDay = append(weatherRecord.Years.ForecastDay, day)
		}

		// Store the weather record in MongoDB
		weatherCollection, err := database.GetCollection("weather_records")
		if err != nil {
			fmt.Println("error getting weather collection:", err)
			return err
		}

		// Check if a weather record already exists for the current month
		filter := bson.M{
			"user_id":                userID,
			"month.forecastday.date": bson.M{"$in": []string{weatherRecord.Month.ForecastDay[0].Date}}, // Check by the first day of the month
		}

		var existingRecord models.WeatherRecord
		err = weatherCollection.FindOne(context.TODO(), filter).Decode(&existingRecord)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				// No existing record found, insert the new record
				_, err := weatherCollection.InsertOne(context.TODO(), weatherRecord)
				if err != nil {
					fmt.Println("error inserting weather record:", err)
					return err
				}
			} else {
				fmt.Println("error finding weather record:", err)
				return err
			}
		} else {
			// Update the existing record
			update := bson.M{
				"$push": bson.M{
					"month.forecastday": bson.M{"$each": weatherRecord.Month.ForecastDay},
					"years.forecastday": bson.M{"$each": weatherRecord.Years.ForecastDay},
				},
			}

			_, err := weatherCollection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				fmt.Println("error updating weather record:", err)
				return err
			}
		}

		// Move to the previous month
		currentDate = startOfMonth.AddDate(0, 0, -1)
	}

	return nil

}
