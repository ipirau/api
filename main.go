package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"log"
	"os"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("aiciiiii")
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// beers = append(beers, Beer{ID: "1", Name: "Corona", Volume: "0.33", Value: "10", Unit: "centiliters"})
	// beers = append(beers, Beer{ID: "2", Name: "Carlsberg", Volume: "0.5", Value: "8", Unit: "centiliters"})
	// beers = append(beers, Beer{ID: "3", Name: "Skol", Volume: "0.5", Value: "7.5", Unit: "centiliters"})

	// router.HandleFunc("/api/beers", getBeers).Methods("GET")
	// router.HandleFunc("/api/beers/{id}", getBeer).Methods("GET")
	// router.HandleFunc("/api/beers", createBeer).Methods("POST")
	// router.HandleFunc("/api/beers/{id}", updateBeer).Methods("PUT")
	// router.HandleFunc("/api/beers/{id}", deleteBeer).Methods("DELETE")

	http.ListenAndServe(":8081", router)
}
