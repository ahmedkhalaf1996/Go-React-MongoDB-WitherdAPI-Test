package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/database"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/models"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.Claims)
	username := claims.Username

	// Retrieve user profile
	userCollection, err := database.GetCollection("users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to database"})
		return
	}

	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user data"})
		}
		return
	}

	// Retrieve weather record associated with the user
	weatherCollection, err := database.GetCollection("weather_records")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to database"})
		return
	}

	var weatherRecord models.WeatherRecord
	err = weatherCollection.FindOne(context.TODO(), bson.M{"user_id": user.ID}).Decode(&weatherRecord)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching weather record"})
			return
		}
		weatherRecord = models.WeatherRecord{
			UserID:    user.ID,
			Month:     models.Forecast{},
			Years:     models.Forecast{},
			CreatedAt: time.Now(),
		}
	}

	// Sanitize user data (remove sensitive fields)
	publicUser := user.SanitizeUser()

	// Return profile data including weather record
	c.JSON(http.StatusOK, gin.H{
		"user":         publicUser,
		"weather_data": weatherRecord,
	})
}
