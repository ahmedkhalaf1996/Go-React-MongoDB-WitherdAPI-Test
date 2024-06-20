package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Day struct {
	Date     string  `bson:"date" json:"date" validate:"required"`
	AvgTempC float64 `bson:"avgtemp_c" json:"avgtemp_c" validate:"required"`
}

type Forecast struct {
	ForecastDay []Day `bson:"forecastday" json:"forecastday" validate:"required"`
}

type Location struct {
	Name    string  `bson:"name" json:"name"`
	Region  string  `bson:"region" json:"region"`
	Country string  `bson:"country" json:"country"`
	Lat     float64 `bson:"lat" json:"lat"`
	Lon     float64 `bson:"lon" json:"lon"`
}

type WeatherRecord struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id" validate:"required"`
	Location  Location           `bson:"location" json:"location" validate:"required"`
	Month     Forecast           `bson:"month" json:"month" validate:"required"`
	Years     Forecast           `bson:"years" json:"years" validate:"required"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
