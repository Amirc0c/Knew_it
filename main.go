package main

import (
	"finalKnewIT/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/products", router.GetTennisThingType).Methods("GET")
	r.HandleFunc("/products/create", router.CreateTennisThingType).Methods("POST")
	r.HandleFunc("/racket/create", router.CreateRacket).Methods("POST")
	r.HandleFunc("/shoes/create", router.CreateShoes).Methods("POST")
	r.HandleFunc("/accessories/create", router.CreateAccessories).Methods("POST")
	r.HandleFunc("/balls/create", router.CreateBalls).Methods("POST")
	r.HandleFunc("/racket/get", router.GetRackets).Methods("GET")
	r.HandleFunc("/shoes/get", router.GetShoes).Methods("GET")
	r.HandleFunc("/accessories/get", router.GetAccessories).Methods("GET")
	r.HandleFunc("/balls/get", router.GetBalls).Methods("GET")
	r.HandleFunc("/racket/{id}", router.GetRacket).Methods("GET")
	r.HandleFunc("/shoes/{id}", router.GetAccess).Methods("GET")
	r.HandleFunc("/accessories/{id}", router.GetShoe).Methods("GET")
	r.HandleFunc("/balls/{id}", router.GetBall).Methods("GET")
	r.HandleFunc("/racket/{id}", router.DeleteRacket).Methods("DELETE")
	r.HandleFunc("/shoes/{id}", router.DeleteAccess).Methods("DELETE")
	r.HandleFunc("/accessories/{id}", router.DeleteShoe).Methods("DELETE")
	r.HandleFunc("/balls/{id}", router.DeleteBall).Methods("DELETE")
	r.HandleFunc("/racket/purchase", router.MakePurchaseForRacket).Methods("POST")
	r.HandleFunc("/shoes/purchase", router.MakePurchaseForShoes).Methods("POST")
	r.HandleFunc("/balls/purchase", router.MakePurchaseForBalls).Methods("POST")
	r.HandleFunc("/accessories/purchase", router.MakePurchaseForAccessories).Methods("POST")
	r.HandleFunc("/reviews/create", router.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{id}", router.GetReviewsByProductID).Methods("GET")
	// r.HandleFunc("/save-animal-types-csv", router.SaveAnimalTypesCSVHandler).Methods("GET")
	// r.HandleFunc("/save-animals-csv", router.SaveAnimalsCSVHandler).Methods("GET")
	// r.HandleFunc("/save-food-types-csv", router.SaveFoodTypesCSVHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
func main() {
	router.InitialMigration()
	initializeRouter()

}
