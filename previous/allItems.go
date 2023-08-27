package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

var mockData = []Item{
	{
		ID:          1,
		Title:       "Pumpkin",
		Price:       "70",
		Description: "Oth adverse food reactions, NEC, sequela",
	},
	{
		ID:          2,
		Title:       "Melon - Watermelon, Seedless",
		Price:       "34334",
		Description: "Nondisp apophyseal fx r femur, init for opn fx type 3A/B/C",
	},
	{
		ID:          3,
		Title:       "Potatoes - Yukon Gold, 80 Ct",
		Price:       "588",
		Description: "Dislocation of distal radioulnar joint of right wrist, init",
	},
	{
		ID:          4,
		Title:       "Beans - French",
		Price:       "7",
		Description: "Laceration with foreign body of right wrist, sequela",
	},
	{
		ID:          5,
		Title:       "Nori Sea Weed - Gold Label",
		Price:       "3313",
		Description: "Toxic effect of smoke, accidental (unintentional)",
	},
	{
		ID:          6,
		Title:       "Vodka - Smirnoff",
		Price:       "42",
		Description: "Toxic effect of venom of centipede/millipede, slf-hrm, sqla",
	},
	{
		ID:          7,
		Title:       "Goulash Seasoning",
		Price:       "18",
		Description: "Milt op w combat using blunt/pierc object, milt, subs",
	},
	{
		ID:          8,
		Title:       "Crackers - Trio",
		Price:       "81",
		Description: "Placentitis",
	},
	{
		ID:          9,
		Title:       "Stock - Beef, White",
		Price:       "8",
		Description: "Unspecified focal traumatic brain injury",
	},
	{
		ID:          10,
		Title:       "Table Cloth 53x69 White",
		Price:       "2910",
		Description: "Undrdose of unsp agents prim acting on the resp sys, subs",
	},
}

func getAllItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In getAllItem")

	queries := r.URL.Query()
	maxPrice := queries.Get("maxPrice")

	var outPutdata []Item
	if _, exists := queries["maxPrice"]; exists {

		for _, item := range mockData {
			currentItemPrice, _ := strconv.Atoi(item.Price)
			qPrice, _ := strconv.Atoi(maxPrice)
			if currentItemPrice < qPrice {
				outPutdata = append(outPutdata, item)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json, err := json.Marshal(outPutdata)
		if err != nil {
			panic(err)
			return
		}

		w.Write(json)

	}

	// for i, item := range mockData {
	// 	fmt.Printf("Index: %d , value => %v\n", i, item.Title)
	// }

	// jsonData, err := json.Marshal(mockData)

	// if err != nil {
	// 	panic(err)
	// 	return
	// }

	// w.Write(jsonData)
}

func getSelectedItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In getSelecteItem")
	params := mux.Vars(r)
	val, _ := strconv.Atoi(params["id"])
	var op Item
	for i := 0; i < len(mockData); i++ {
		if mockData[i].ID == val {
			op = mockData[i]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	send, err := json.Marshal(op)
	if err != nil {
		panic(err)
	}
	w.Write(send)
}
