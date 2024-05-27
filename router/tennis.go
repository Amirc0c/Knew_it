package router

import (

	// "html/template"

	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "os/user"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TennisThingType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Racket struct {
	ID                int             `json:"id"`
	Brand             string          `json:"brand"`
	TennisThingTypeID int             `json:"tennis_thing_type_id"`
	Type              TennisThingType `json:"type" gorm:"foreignKey:TennisThingTypeID"`
	Module            string          `json:"module"`
	Weight            string          `json:"weight"`
	HeadSize          int             `json:"head_size"`
	HandleSize        string          `json:"handle_size"`
	Price             float64         `json:"price"`
}

type Shoes struct {
	ID                int             `json:"id"`
	Brand             string          `json:"brand"`
	TennisThingTypeID int             `json:"tennis_thing_type_id"`
	Type              TennisThingType `json:"type" gorm:"foreignKey:TennisThingTypeID"`
	Module            string          `json:"module"`
	Size              int             `json:"size"`
	Cover             string          `json:"cover"`
	Price             float64         `json:"price"`
}

type Accessories struct {
	ID                int             `json:"id"`
	Brand             string          `json:"brand"`
	TennisThingTypeID int             `json:"tennis_thing_type_id"`
	Type              TennisThingType `json:"type" gorm:"foreignKey:TennisThingTypeID"`
	Price             float64         `json:"price"`
}

type Balls struct {
	ID                int             `json:"id"`
	Brand             string          `json:"brand"`
	TennisThingTypeID int             `json:"tennis_thing_type_id"`
	Type              TennisThingType `json:"type" gorm:"foreignKey:TennisThingTypeID"`
	Cover             string          `json:"cover"`
	Price             float64         `json:"price"`
}

type Users struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
}

type PurchaseForRacket struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	RacketID    int     `json:"racket_id"`
	Quantity    float64 `json:"quantity"`
	RacketPrice float64 `json:"racket_price"`
	TotalPrice  float64 `json:"total_price"`
}

type PurchaseForShoes struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	ShoesID    int     `json:"shoes_id"`
	Quantity   float64 `json:"quantity"`
	ShoesPrice float64 `json:"shoes_price"`
	TotalPrice float64 `json:"total_price"`
}

type PurchaseForAccessories struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	AccessID    int     `json:"access_id"`
	Quantity    float64 `json:"quantity"`
	AccessPrice float64 `json:"access_price"`
	TotalPrice  float64 `json:"total_price"`
}

type PurchaseForBalls struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	BallsID    int     `json:"balls_id"`
	Quantity   float64 `json:"quantity"`
	BallsPrice float64 `json:"balls_price"`
	TotalPrice float64 `json:"total_price"`
}
type Review struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Comment     string `json:"comment"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
}
type AllPurchases struct {
	RacketPurchases      []PurchaseForRacket      `json:"racket_purchases"`
	ShoesPurchases       []PurchaseForShoes       `json:"shoes_purchases"`
	AccessoriesPurchases []PurchaseForAccessories `json:"accessories_purchases"`
	BallsPurchases       []PurchaseForBalls       `json:"balls_purchases"`
}

type CacheData struct {
	Value      interface{}
	Expiration time.Time
}

type Cache struct {
	storage map[string]CacheData
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]CacheData),
	}
}

var db *gorm.DB
var err error
var racket Racket
var user Users
var CacheInstance *Cache

func InitialMigration() {
	dsn := "host=localhost user=postgres password=Amir2009 dbname=postgres port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	err = db.AutoMigrate(&TennisThingType{}, &Racket{}, &Shoes{}, &Accessories{}, &Balls{}, &Users{}, &PurchaseForRacket{}, &PurchaseForShoes{}, &PurchaseForAccessories{}, &PurchaseForBalls{}, &Review{})
	if err != nil {
		log.Fatalf("Ошибка при автомиграции таблиц: %v", err)
	}
}
func GetTennisThingType(w http.ResponseWriter, r *http.Request) {

	var TennisThingType []TennisThingType
	if err := db.Find(&TennisThingType).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TennisThingType)
}
func CreateTennisThingType(w http.ResponseWriter, r *http.Request) {

	var TennisThingType TennisThingType
	if err := json.NewDecoder(r.Body).Decode(&TennisThingType); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&TennisThingType)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("Created animal type:", TennisThingType)
}
func CreateUsers(w http.ResponseWriter, r *http.Request) {

	var users Users
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&users)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user Users
	if err := db.First(&user, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var user Users
	if err := db.Delete(&user, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []Users
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func CreateRacket(w http.ResponseWriter, r *http.Request) {
	var racket Racket
	if err := json.NewDecoder(r.Body).Decode(&racket); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&racket)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}
func GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reviews []Review
	db.Find(&reviews)
	json.NewEncoder(w).Encode(reviews)

}
func CreateShoes(w http.ResponseWriter, r *http.Request) {
	var shoes Shoes
	if err := json.NewDecoder(r.Body).Decode(&shoes); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&shoes)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}
func CreateAccessories(w http.ResponseWriter, r *http.Request) {
	var accessories Accessories
	if err := json.NewDecoder(r.Body).Decode(&accessories); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&accessories)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}
func CreateBalls(w http.ResponseWriter, r *http.Request) {
	var balls Balls
	if err := json.NewDecoder(r.Body).Decode(&balls); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&balls)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

}
func GetRackets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Rackets []Racket
	db.Find(&Rackets)
	json.NewEncoder(w).Encode(Rackets)
	fmt.Println("getrackets")
}
func GetShoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Shoes []Shoes
	db.Find(&Shoes)
	json.NewEncoder(w).Encode(Shoes)
	fmt.Println("getshoes")
}
func GetAccessories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Accessories []Accessories
	db.Find(&Accessories)
	json.NewEncoder(w).Encode(Accessories)
	fmt.Println("getAccessories")
}
func GetBalls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Balls []Balls
	db.Find(&Balls)
	json.NewEncoder(w).Encode(Balls)
	fmt.Println("getrackets")
}
func GetRacket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Racket Racket
	db.First(&Racket, params["id"])
	json.NewEncoder(w).Encode(Racket)
	fmt.Println("get animal")
}
func GetAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Access Accessories
	db.First(&Access, params["id"])
	json.NewEncoder(w).Encode(Access)
	fmt.Println("get Access")
}
func GetShoe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var shoe Shoes
	db.First(&shoe, params["id"])
	json.NewEncoder(w).Encode(shoe)
	fmt.Println("get shoe")
}
func GetBall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Ball Balls
	db.First(&Ball, params["id"])
	json.NewEncoder(w).Encode(Ball)
	fmt.Println("get animal")
}
func DeleteRacket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Racket Racket
	db.Delete(&Racket, params["id"])
	json.NewEncoder(w).Encode(" успешно")
	fmt.Println("delete")
}
func DeleteAccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var Access Accessories
	db.Delete(&Access, params["id"])
	json.NewEncoder(w).Encode(" успешно")
	fmt.Println("delete")
}
func DeleteShoe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var shoe Shoes
	db.Delete(&shoe, params["id"])
	json.NewEncoder(w).Encode(" успешно")
	fmt.Println("delete")
}
func DeleteBall(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var ball Balls
	db.Delete(&ball, params["id"])
	json.NewEncoder(w).Encode(" успешно")
	fmt.Println("delete")
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var review Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.Create(&review)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func GetReviewsByProductID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var reviews []Review
	db.Where("product_id = ?", params["id"]).Find(&reviews)
	json.NewEncoder(w).Encode(reviews)
}

func calculateTotalPriceAndUpdateBalance(userID int, quantity, price float64) (float64, error) {
	totalPrice := quantity * price

	var user Users
	if err := db.First(&user, userID).Error; err != nil {
		return 0, err
	}

	if user.Balance < totalPrice {
		return 0, fmt.Errorf("не хватает)")
	}

	user.Balance -= totalPrice
	if err := db.Save(&user).Error; err != nil {
		return 0, err
	}

	return totalPrice, nil
}

func MakePurchaseForRacket(w http.ResponseWriter, r *http.Request) {
	var purchase struct {
		UserID      int     `json:"user_id"`
		RacketID    int     `json:"racket_id"`
		Quantity    float64 `json:"quantity"`
		RacketPrice float64 `json:"racket_price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество ", http.StatusBadRequest)
		return
	}

	totalPrice, err := calculateTotalPriceAndUpdateBalance(purchase.UserID, purchase.Quantity, purchase.RacketPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseRecord := PurchaseForRacket{
		RacketID:    purchase.RacketID,
		Quantity:    purchase.Quantity,
		RacketPrice: totalPrice,
	}

	if err := db.Create(&purchaseRecord).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchaseRecord)
}
func MakePurchaseForShoes(w http.ResponseWriter, r *http.Request) {
	var purchase struct {
		UserID     int     `json:"user_id"`
		ShoesID    int     `json:"shoes_id"`
		Quantity   float64 `json:"quantity"`
		ShoesPrice float64 `json:"shoes_price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество ", http.StatusBadRequest)
		return
	}

	totalPrice, err := calculateTotalPriceAndUpdateBalance(purchase.UserID, purchase.Quantity, purchase.ShoesPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseRecord := PurchaseForShoes{
		UserID:     purchase.UserID,
		ShoesID:    purchase.ShoesID,
		Quantity:   purchase.Quantity,
		ShoesPrice: purchase.ShoesPrice,
		TotalPrice: totalPrice,
	}

	if err := db.Create(&purchaseRecord).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchaseRecord)
}

func MakePurchaseForAccessories(w http.ResponseWriter, r *http.Request) {
	var purchase struct {
		UserID      int     `json:"user_id"`
		AccessID    int     `json:"access_id"`
		Quantity    float64 `json:"quantity"`
		AccessPrice float64 `json:"access_price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество ", http.StatusBadRequest)
		return
	}

	totalPrice, err := calculateTotalPriceAndUpdateBalance(purchase.UserID, purchase.Quantity, purchase.AccessPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseRecord := PurchaseForAccessories{
		UserID:      purchase.UserID,
		AccessID:    purchase.AccessID,
		Quantity:    purchase.Quantity,
		AccessPrice: purchase.AccessPrice,
		TotalPrice:  totalPrice,
	}

	if err := db.Create(&purchaseRecord).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchaseRecord)
}

func MakePurchaseForBalls(w http.ResponseWriter, r *http.Request) {
	var purchase struct {
		UserID     int     `json:"user_id"`
		BallsID    int     `json:"balls_id"`
		Quantity   float64 `json:"quantity"`
		BallsPrice float64 `json:"balls_price"`
	}
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество ", http.StatusBadRequest)
		return
	}

	totalPrice, err := calculateTotalPriceAndUpdateBalance(purchase.UserID, purchase.Quantity, purchase.BallsPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	purchaseRecord := PurchaseForBalls{
		UserID:     purchase.UserID,
		BallsID:    purchase.BallsID,
		Quantity:   purchase.Quantity,
		BallsPrice: purchase.BallsPrice,
		TotalPrice: totalPrice,
	}

	if err := db.Create(&purchaseRecord).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchaseRecord)
}

func GetAllPurchases(w http.ResponseWriter, r *http.Request) {
	var racketPurchases []PurchaseForRacket
	var shoesPurchases []PurchaseForShoes
	var accessoriesPurchases []PurchaseForAccessories
	var ballsPurchases []PurchaseForBalls

	db.Find(&racketPurchases)
	db.Find(&shoesPurchases)
	db.Find(&accessoriesPurchases)
	db.Find(&ballsPurchases)

	allPurchases := AllPurchases{
		RacketPurchases:      racketPurchases,
		ShoesPurchases:       shoesPurchases,
		AccessoriesPurchases: accessoriesPurchases,
		BallsPurchases:       ballsPurchases,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allPurchases)
}
