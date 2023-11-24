package controller

import (
	"errors"
	"os"
)

var (
	REVIEWS_SERVICE = os.Getenv("REVIEWS_SERVICE")
	RATINGS_SERVICE = os.Getenv("RATINGS_SERVICE")
)

func CheckEnv() error {
	//Check if REVIEWS_SERVICE is set
	if len(REVIEWS_SERVICE) < 0 {
		return errors.New("REVIEWS_SERVICE env variable could not be found")
	}

	//Check if RATINGS_SERVICE is set
	if len(RATINGS_SERVICE) < 0 {
		return errors.New("RATINGS_SERVICE env variable could not be found")
	}
	return nil
}
