package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/database"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/models"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/services"
	"github.com/ahmedkhalaf1996/Go-React-MongoDB-WitherdAPI-Test/utils"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	user.Password = string(hashedPassword)
	user.ID = primitive.NewObjectID()

	userCollection, err := database.GetCollection("users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to database"})
		return
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	go func() {
		if err := services.FetchAndStoreWeatherData(user.ID, user.Lat, user.Lon, time.Now()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating wather data."})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Var(user.Username, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	if err := validate.Var(user.Password, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
		return
	}

	userCollection, err := database.GetCollection("users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while connecting to database"})
		return
	}

	var foundUser models.User
	err = userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateJWT(foundUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
		return
	}

	go func() {
		if err := services.FetchAndStoreWeatherData(foundUser.ID, foundUser.Lat, foundUser.Lon, time.Now()); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating wather data."})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{"token": token})
}
