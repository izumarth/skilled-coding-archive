package main

import (
	"fmt"
	"go-breeders/models"
	"go-breeders/pets"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.render(w, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	log.Println(breed)
	log.Println(breed.ID)

	dom, _ := app.App.Models.Dog.GetDogOfMonthById(1)
	log.Println(dom)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedID:          breed.ID,
			Color:            "black",
			DateOfBirth:      dob,
			SpayedOrNeutered: false,
			Description:      "hoge",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	data := make(map[string]any)
	data["dog"] = dog

	app.render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("Dog"))

}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("Cat"))

}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactroy("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactroy("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}
	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	p, err := pets.NewPetBuilder().
		SetSepecies("dog").
		SetBreed("izumart").
		SetWeight(15).
		SetDescription("hoge Descriptoin").
		SetColor("Black and White").
		SetAge(3).
		SetAgeEstimated(true).
		Build()

	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools

	catBreeds, err := app.App.CatService.GetAllCatBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	// Setup toolbox
	var t toolbox.Tools

	// GEt Species form URL itself.
	species := chi.URLParam(r, "species")

	// Get breed from the URL
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	// Create a pet from abstract factory
	pet, err := pets.NewPetWithBreedFromAbstractFactroy(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	// Write the result as JSON
	_ = t.WriteJSON(w, http.StatusOK, pet)
}
