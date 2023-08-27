package controllers

import (
	"context"
	"ecom/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var mockUser = []User{
	{
		Id:       "jkwnf(fjhbew)",
		Name:     "Ritik",
		Email:    "r@g.com",
		Password: "acb",
	},
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var jwtKey = []byte("Secrectkey")

type Claims struct {
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

const connectionString = "mongodb+srv://ritik:bY1MUy916TmqCu4o@golangdb.gqx5mtp.mongodb.net/?retryWrites=true&w=majority"

const dbName = "Ecom Apis"

// Most important

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	// connnect to mongoDb

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection succeeded")

	collection = client.Database(dbName).Collection(dbName)

	// collection instance

	fmt.Println("collection instance ready")
}

func insertOneUser(user models.User) {
	insert, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User add succesful with id  ", insert.InsertedID)
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
	var data LoginType
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var dbUser User
	// Process the data in 'data'
	for _, current := range mockUser {
		if current.Email == data.Email && current.Password == data.Password {
			fmt.Println("Match found:", current)
			dbUser = current
		}
	}

	// if no user found
	if dbUser == (User{}) {
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

	tokenString, _ := generateToken(dbUser.Name)
	response := map[string]interface{}{
		"token": tokenString,
		"user":  dbUser,
	}
	json.NewEncoder(w).Encode(response)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	fmt.Println("user")
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid User data ", http.StatusBadRequest)
	}

	insertOneUser(newUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// output
	response := Response{
		Status:  "success",
		Message: "user created succesfully",
	}

	for index, current := range mockUser {
		fmt.Println(index, " this is val", current)
	}
	val, _ := json.Marshal(response)
	w.Write(val)
}
