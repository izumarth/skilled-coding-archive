package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type mySqlRepository struct {
	DB *sql.DB
}

func newMySQLRepository(conn *sql.DB) Repository {
	return &mySqlRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository() Repository {
	return &testRepository{
		DB: nil,
	}
}
