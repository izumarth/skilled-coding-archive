package pets

import (
	"errors"
	"fmt"
	"go-breeders/configuration"
	"go-breeders/models"
	"log"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
	NewPetWithBreed(breed string) AnimalInterface
}

type DogAbstractFactroy struct{}

func (df *DogAbstractFactroy) NewPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

func (df *DogAbstractFactroy) NewPetWithBreed(b string) AnimalInterface {
	app := configuration.GetInstance()
	breed, _ := app.Models.DogBreed.GetBreedByName(b)
	return &DogFromFactory{
		Pet: &models.Dog{
			Breed: *breed,
		},
	}
}

type CatAbstractFactroy struct{}

func (cf *CatAbstractFactroy) NewPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func (cf *CatAbstractFactroy) NewPetWithBreed(b string) AnimalInterface {
	app := configuration.GetInstance()
	breed, err := app.CatService.Remote.GetCatBreedsByName(b)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &CatFromFactory{
		Pet: &models.Cat{
			Breed: *breed,
		},
	}
}

func NewPetFromAbstractFactroy(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var dogFactory DogAbstractFactroy
		dog := dogFactory.NewPet()
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactroy
		cat := catFactory.NewPet()
		return cat, nil
	default:
		return nil, errors.New("invalid species")
	}
}

func NewPetWithBreedFromAbstractFactroy(species string, breed string) (AnimalInterface, error) {
	log.Println(species)
	switch species {
	case "dog":
		var dogFactory DogAbstractFactroy
		dog := dogFactory.NewPetWithBreed(breed)
		return dog, nil
	case "cat":
		var catFactory CatAbstractFactroy
		cat := catFactory.NewPetWithBreed(breed)
		return cat, nil
	default:
		return nil, errors.New("invalid species")
	}
}
