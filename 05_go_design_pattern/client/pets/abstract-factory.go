package pets

import (
	"errors"
	"fmt"
	"go-breeders/models"
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
}

type DogAbstractFactroy struct{}

func (df *DogAbstractFactroy) NewPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactroy struct{}

func (cf *CatAbstractFactroy) NewPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
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