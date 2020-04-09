package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Beer struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Volume string `json:"volume"`
	Value  string `json:"value"`
	Unit   string `json:"unit"`
}

var beers []Beer

func test() {
	fmt.Println("It works!")
}

func getBeers(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beers)
}

func getBeer(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range beers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}

	fmt.Println(params)
}

func createBeer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var beer Beer

	_ = json.NewDecoder(req.Body).Decode(&beer)
	beer.ID = strconv.Itoa(12345)
	beers = append(beers, beer)
	json.NewEncoder(w).Encode(beer)
}

func updateBeer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range beers {
		if item.ID == params["id"] {
			beers = append(beers[:index], beers[index+1:]...)
			var book Beer
			fmt.Println(req.Body)
			_ = json.NewDecoder(req.Body).Decode(&book)
			book.ID = params["id"]
			beers = append(beers, book)
			json.NewEncoder(w).Encode(fmt.Sprintf("The product with the id %s has been updated successfully", params["id"]))
			return
		}
	}
}

func deleteBeer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range beers {
		if item.ID == params["id"] {
			beers = append(beers[:index], beers[:index+1]...)
			break
		}
	}

	json.NewEncoder(w).Encode(fmt.Sprintf("The product with the id %s has been deleted successfully", params["id"]))
}
