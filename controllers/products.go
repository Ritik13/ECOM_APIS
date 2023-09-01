package controllers

import (
	"context"
	"ecom/database"
	"ecom/models"
	"encoding/json"
	"fmt"
	"net/http"

	"ecom/structs"

	"go.mongodb.org/mongo-driver/bson"
)

func setBasicHeader(w http.ResponseWriter) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func errorResponse(w http.ResponseWriter, err error, message string) {
	setBasicHeader(w)
	fmt.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	response := structs.Response{
		Status:  "error",
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var dbResult []models.Product

	cursor, err := database.ProductCollection.Find(ctx, bson.D{{}})

	// err occred
	if err != nil {
		errorResponse(w, err, "No such product")
		return

	}

	// if cursor found -> treat cursor as pointer , not address but return a record start

	cursorErr := cursor.All(ctx, &dbResult)

	if cursorErr != nil {
		errorResponse(w, err, "No record found")
		return

	} else {
		setBasicHeader(w)
		w.WriteHeader(http.StatusOK)
		if len(dbResult) == 0 {
			dbResult = []models.Product{} // Assign an empty slice here
		}
		response := structs.Response{
			Status:  "success",
			Message: dbResult,
		}
		defer cursor.Close(ctx)
		json.NewEncoder(w).Encode(response)

	}
	defer r.Body.Close()
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		errorResponse(w, err, "Error decoding product")
		return
	}

	updateFields := bson.M{
		"name":        product.Name,
		"price":       product.Price,
		"description": product.Description,
		"category":    product.Category,
		"stock":       product.Stock,
	}

	_, err = database.ProductCollection.UpdateOne(ctx, bson.M{"_id": product.Id}, bson.M{"$set": updateFields})

	if err != nil {
		errorResponse(w, err, "No record found")

	} else {
		setBasicHeader(w)
		w.WriteHeader(http.StatusOK)
		response := structs.Response{
			Status:  "success",
			Message: "Record successfully updated",
		}
		json.NewEncoder(w).Encode(response)

	}
	return
}

func AddNewProdect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		errorResponse(w, err, "Error decoding product")
		return
	}

	_, err = database.ProductCollection.InsertOne(ctx, product)

	if err != nil {
		errorResponse(w, err, "No record found")

	} else {
		setBasicHeader(w)
		w.WriteHeader(http.StatusOK)
		response := structs.Response{
			Status:  "success",
			Message: "Record successfully added",
		}
		json.NewEncoder(w).Encode(response)

	}
}
