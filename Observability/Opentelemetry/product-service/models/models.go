package models

// Product type

type Product struct {
	ID              int     `json:"id" uri:"id" gorm:"primaryKey"`
	Name            string  `json:product_name, name`
	Price           float64 `json:price`
	ProductCategory string  `json:product_category`
	Stocks          int     `json:stocks`
	Timestamp       string  `json:timestamp`
	ImageSource     string  `json:image_source`
}

// User

type User struct {
	ID            int    `json:"id" uri:"id" gorm:"primaryKey"`
	UUID          string `json:uuid`
	Username      string `json:username`
	Password      string `json:password`
	Name          string `json:name`
	Surname       string `json:surname`
	Email         string `json:email`
	ContactNumber string `json:contact_number`
	Address       string `json:address`
}

// Reviews

type Reviews struct {
	ID        int    `json:"id" uri:"id" gorm:"primaryKey"`
	Comment   string `json:Comments`
	ProductID int    `json:product_id`
}

//Ratings

type Ratings struct {
	ID        int    `json:"id" uri:"id" gorm:"primaryKey"`
	Rates     int    `json:rates`
	Category  string `json:category`
	ProductID int    `json:product_id`
}
