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

type PurchaseForRacket struct {
	ID       int     `json:"id"`
	RacketID int     `json:"racket_id"`
	Quantity float64 `json:"quantity"`

	RacketPrice float64 `json:"racket_price"`
}
type PurchaseForShoes struct {
	ID       int     `json:"id"`
	ShoesID  int     `json:"shoes_id"`
	Quantity float64 `json:"quantity"`

	ShoesPrice float64 `json:"shoes_price"`
}
type PurchaseForAccessories struct {
	ID       int     `json:"id"`
	AccessID int     `json:"access_id"`
	Quantity float64 `json:"quantity"`

	AccessPrice float64 `json:"access_price"`
}
type PurchaseForBalls struct {
	ID       int     `json:"id"`
	BallsID  int     `json:"Balls_id"`
	Quantity float64 `json:"quantity"`

	BallsPrice float64 `json:"balls_price"`
}

type Review struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Comment     string `json:"comment"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
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

var CacheInstance *Cache

func InitialMigration() {
	dsn := "host=localhost user=postgres password=Amir2009 dbname=postgres port=5432 sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	err = db.AutoMigrate(&TennisThingType{}, &Racket{}, &Shoes{}, &Accessories{}, &Balls{}, &PurchaseForRacket{}, &PurchaseForShoes{}, &PurchaseForAccessories{}, &PurchaseForBalls{}, &Review{})
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

	// tmpl := template.Must(template.ParseFiles("templates/GetAnimalTypes.html"))
	// if err := tmpl.Execute(w,TennisThingType); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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

	// tmpl := template.Must(template.ParseFiles("templates/CreateAnimalType.html"))
	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusCreated)
	fmt.Println("Created animal type:", TennisThingType)
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

	// tmpl := template.Must(template.ParseFiles("templates/CreateAnimal.html"))
	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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

	// tmpl := template.Must(template.ParseFiles("templates/CreateAnimal.html"))
	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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

	// tmpl := template.Must(template.ParseFiles("templates/CreateAnimal.html"))
	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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

	// tmpl := template.Must(template.ParseFiles("templates/CreateAnimal.html"))
	// if err := tmpl.Execute(w, nil); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
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

func MakePurchaseForRacket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var purchase PurchaseForRacket
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество", http.StatusBadRequest)
		return
	}

	total := calculateTotalPriceOfRacket(purchase)

	purchase.RacketPrice = total

	if err := db.Create(&purchase).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchase)
}

func calculateTotalPriceOfRacket(purchase PurchaseForRacket) float64 {
	totalPrice := (purchase.Quantity * purchase.RacketPrice)
	return totalPrice
}

func MakePurchaseForShoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var purchase PurchaseForShoes
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество", http.StatusBadRequest)
		return
	}

	total := calculateTotalPriceOfShoes(purchase)

	purchase.ShoesPrice = total

	if err := db.Create(&purchase).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchase)
}

func calculateTotalPriceOfShoes(purchase PurchaseForShoes) float64 {
	totalPrice := (purchase.Quantity * purchase.ShoesPrice)
	return totalPrice
}
func MakePurchaseForAccessories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var purchase PurchaseForAccessories
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество", http.StatusBadRequest)
		return
	}

	total := calculateTotalPriceOfAccessories(purchase)

	purchase.AccessPrice = total

	if err := db.Create(&purchase).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchase)
}

func calculateTotalPriceOfAccessories(purchase PurchaseForAccessories) float64 {
	totalPrice := (purchase.Quantity * purchase.AccessPrice)
	return totalPrice
}
func MakePurchaseForBalls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var purchase PurchaseForBalls
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if purchase.Quantity <= 0 {
		http.Error(w, "не правильное количество", http.StatusBadRequest)
		return
	}

	total := calculateTotalPriceOfBalls(purchase)

	purchase.BallsPrice = total

	if err := db.Create(&purchase).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(purchase)
}

func calculateTotalPriceOfBalls(purchase PurchaseForBalls) float64 {
	totalPrice := (purchase.Quantity * purchase.BallsPrice)
	return totalPrice
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
	db.Where("id = ?", params["id"]).Find(&reviews)
	json.NewEncoder(w).Encode(reviews)
}
