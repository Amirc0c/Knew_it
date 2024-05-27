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
	r.HandleFunc("/products", router.CreateTennisThingType).Methods("POST")
	r.HandleFunc("/racket", router.CreateRacket).Methods("POST")
	r.HandleFunc("/shoes", router.CreateShoes).Methods("POST")
	r.HandleFunc("/accessories", router.CreateAccessories).Methods("POST")
	r.HandleFunc("/balls", router.CreateBalls).Methods("POST")
	r.HandleFunc("/racket", router.GetRackets).Methods("GET")
	r.HandleFunc("/shoes", router.GetShoes).Methods("GET")
	r.HandleFunc("/accessories", router.GetAccessories).Methods("GET")
	r.HandleFunc("/balls", router.GetBalls).Methods("GET")
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
	r.HandleFunc("/reviews", router.CreateReview).Methods("POST")
	r.HandleFunc("/reviews/{id}", router.GetReviewsByProductID).Methods("GET")
	r.HandleFunc("/reviews", router.GetReviews).Methods("GET")
	r.HandleFunc("/users", router.CreateUsers).Methods("POST")
	r.HandleFunc("/users", router.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", router.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", router.DeleteUser).Methods("DELETE")
	r.HandleFunc("/purchase", router.GetAllPurchases).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
func main() {
	router.InitialMigration()
	initializeRouter()

}
