package controllers

import (
	"context"
	"ecom/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"ecom/database"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var jwtKey = []byte("Secrectkey")

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

// contexts

func insertOneUser(user models.User) (string, error) {
	ctx := context.Background()
	_, err := database.UserCollection.InsertOne(ctx, user)
	return "User Added successfully", err
}

// docs followed
// https://www.sohamkamani.com/golang/jwt-authentication/
// bY1MUy916TmqCu4o

func setJSONresponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func generateToken(userName string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func deadToken(w http.ResponseWriter, r *http.Request) {
	// make token invalid
}

func authProtector(token string) bool {
	claims := &Claims{}

	tokenString, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false
		}
	}

	if !tokenString.Valid {
		return false
	}

	return true
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	filterBy := bson.M{"email": data.Email}
	var result models.User
	fmt.Println(result, data.Email)
	database.UserCollection.FindOne(context.TODO(), filterBy).Decode(&result)

	// if no user found
	if result == (models.User{}) {
		w.WriteHeader(http.StatusBadRequest)
		response := Response{
			Status:  "eror",
			Message: "user Not Found",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// jwt token generate

	setJSONresponseHeader(w)
	w.WriteHeader(http.StatusOK)

	// generateToken

	tokenString, _ := generateToken(result.Name)
	response := map[string]interface{}{
		"token": tokenString,
		"user":  result,
	}
	json.NewEncoder(w).Encode(response)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	// convert body in json
	jsonError := json.NewDecoder(r.Body).Decode(&newUser)

	if jsonError != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var response Response
	w.Header().Set("Content-Type", "application/json")

	_, dbOperationError := insertOneUser(newUser)
	if dbOperationError != nil {
		fmt.Println(dbOperationError)
		w.WriteHeader(http.StatusBadRequest)
		response = Response{
			Status:  "error",
			Message: "Can not register user!",
		}
	} else {
		w.WriteHeader(http.StatusOK)
		response = Response{
			Status:  "success",
			Message: "user created succesfully",
		}
	}
	val, _ := json.Marshal(response)
	w.Write(val)
}
