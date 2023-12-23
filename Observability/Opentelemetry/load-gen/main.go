package main

import (
	"context"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	API_ENDPOINT = os.Getenv("API_ENDPOINT")
	TIMEOUT      = os.Getenv("REQUEST_TIMEOUT")
)

type Reviews struct {
	ID        int
	Comment   string
	ProductID int
}

type Product struct {
	ID              int
	Name            string
	Price           float64
	ProductCategory string
	Stocks          int
	Timestamp       string
	ImageSource     string
}

type Ratings struct {
	ID        int
	Rates     int
	Category  string
	ProductID int
}

type User struct {
	ID            int
	UUID          string
	Username      string
	Password      string
	Name          string
	Surname       string
	Email         string
	ContactNumber string
	Address       string
}

func addRatings(channel chan<- int) {
	var ratings []Ratings
	ratingsJSON, err := os.Open("seed/ratings.json")

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	defer ratingsJSON.Close()

	byteRatings, err := ioutil.ReadAll(ratingsJSON)

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	json.Unmarshal(byteRatings, &ratings)

	for i := 0; i < len(ratings); i++ {

		json_data, err := json.Marshal(ratings[i])

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		resp, err := http.Post(fmt.Sprintf("%s/product/ratings/create", API_ENDPOINT),
			"application/json",
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		defer resp.Body.Close()
	}
}

func getRatings(channel chan<- int) {
	resp, err := http.Get(fmt.Sprintf("%s/product/ratings/", API_ENDPOINT))
	if err != nil {
		log.Println(err)
		channel <- 1
		return
	}

	defer resp.Body.Close()
}

func addReviews(channel chan<- int) {
	var reviews []Reviews
	reviewsJSON, err := os.Open("seed/reviews.json")

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	defer reviewsJSON.Close()

	byteReviews, err := ioutil.ReadAll(reviewsJSON)

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	json.Unmarshal(byteReviews, &reviews)

	for i := 0; i < len(reviews); i++ {

		json_data, err := json.Marshal(reviews[i])

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		resp, err := http.Post(fmt.Sprintf("%s/product/reviews/create", API_ENDPOINT),
			"application/json",
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		defer resp.Body.Close()
	}
}

func getReviews(channel chan<- int) {
	resp, err := http.Get(fmt.Sprintf("%s/product/reviews/", API_ENDPOINT))
	if err != nil {
		log.Println(err)
		channel <- 1
		return
	}

	defer resp.Body.Close()
}

func addProducts(channel chan<- int) {
	var product []Product
	productsJSON, err := os.Open("seed/products.json")

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	defer productsJSON.Close()

	byteProducts, err := ioutil.ReadAll(productsJSON)

	if err != nil {
		log.Println(err)
		channel <- 1
	}

	json.Unmarshal(byteProducts, &product)

	for i := 0; i < len(product); i++ {

		json_data, err := json.Marshal(product[i])

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		resp, err := http.Post(fmt.Sprintf("%s/product/create", API_ENDPOINT),
			"application/json",
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Println(err)
			channel <- 1
		}

		defer resp.Body.Close()
	}
}

func getProducts(channel chan<- int) {
	resp, err := http.Get(fmt.Sprintf("%s/product/", API_ENDPOINT))
	if err != nil {
		log.Println(err)
		channel <- 1
		return
	}

	defer resp.Body.Close()
}

func Seed(ctx context.Context) {

	channel := make(chan int)

	for {
		select {
		case <-ctx.Done():
			close(channel)
			return
		case result := <-channel:
			if result == 1 {
				close(channel)
				return
			}
		default:
			go addRatings(channel)
			go getRatings(channel)
			go addReviews(channel)
			go getReviews(channel)
			go addProducts(channel)
			go getProducts(channel)

			timeout, _ := strconv.Atoi(TIMEOUT)

			if timeout > 1 {
				time.Sleep(time.Duration(timeout) * time.Second)
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	Seed(ctx)
}
