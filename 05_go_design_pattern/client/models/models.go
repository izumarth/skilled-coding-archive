package models

import (
	"database/sql"
	"time"
)

var repo Repository

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMySQLRepository(conn)
	} else {
		repo = newTestRepository()
	}

	return &Models{
		DogBreed: DogBreed{},
	}
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"life_span"`
	Details          string `json:"details"`
	AlternateNames   string `json:"altenate_names"`
	GeographicOrigin string `json:"geographic_origin"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

type CatBreed struct {
	ID               int    `json:"id" xml:"id"`
	Breed            string `json:"breed" xml:"breed"`
	WeightLowLbs     int    `json:"weight_low_lbs" xml:"weight_low_lbs"`
	WeightHighLbs    int    `json:"weight_high_lbs" xml:"weight_high_lbs"`
	AverageWeight    int    `json:"average_weight" xml:"average_weight"`
	Lifespan         int    `json:"life_span" xml:"life_span"`
	Details          string `json:"details" xml:"details"`
	AlternateNames   string `json:"altenate_names" xml:"altenate_names"`
	GeographicOrigin string `json:"geographic_origin" xml:"geographic_origin"`
}

type Dog struct {
	ID               int       `json:"id"`
	DogName          string    `json:"dog_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Description      string    `json:"descriptiong"`
	Weight           int       `json:"weight"`
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Cat struct {
	ID               int       `json:"id"`
	CatName          string    `json:"cat_name"`
	BreedID          int       `json:"breed_id"`
	BreederID        int       `json:"breeder_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	SpayedOrNeutered bool      `json:"spayed_or_neutered"`
	Description      string    `json:"descriptiong"`
	Weight           int       `json:"weight"`
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
}

type Breeder struct {
	ID          int         `json:"id"`
	BreederName string      `json:"breeder_name"`
	Address     string      `json:"address"`
	City        string      `json:"city"`
	ProvState   string      `json:"prov_state"`
	Country     string      `json:"country"`
	Zip         string      `json:"zip"`
	Phone       string      `json:"phone"`
	Email       string      `json:"email"`
	Active      int         `json:"active"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
}

type Pet struct {
	Species     string `json:"species"`
	Breed       string `json:"breed"`
	MinWeight   int    `json:"min_weight"`
	MaxWeight   int    `json:"max_weight"`
	Description string `json:"description"`
	Lifespan    int    `json:"lifespan"`
}