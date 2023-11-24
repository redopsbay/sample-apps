package db

import (
	"fmt"
	"gorm.io/gorm"
	"reviews/logging"
	"reviews/models"
)

//MySQLDB holds the necessary methods and database instance to apply *gorm.DBhandler (SQL) operations

type MySQLDB struct {
	DBhandler *gorm.DB
}

// Create applies the given struct or model to be applied on a database table by utilizing *gorm.DBhandler
func (sql *MySQLDB) AddReviews(reviews *models.Reviews) error {
	res := sql.DBhandler.Create(reviews)
	if res.Error != nil {
		logging.Error(res.Error)
		return res.Error
	}

	logging.Info(res.RowsAffected)

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return err
	}

	defer sqlDB.Close()

	return nil
}

func (sql *MySQLDB) UpdateReviews(reviews *models.Reviews) error {

	//create models.Reviews
	var reviewsUpdate models.Reviews

	sql.DBhandler.First(&reviewsUpdate)
	//Point it to reviewsUpdate
	res := sql.DBhandler.Save(reviews)

	//Validate if no error encountered
	if res.Error != nil {
		logging.Error(res.Error)
		return res.Error
	}

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return err
	}

	defer sqlDB.Close()

	return nil
}

func (sql *MySQLDB) DeleteReviews(reviews *models.Reviews) error {
	res := sql.DBhandler.Delete(reviews)

	if res.Error != nil {
		logging.Error(res.Error)
		return res.Error
	}

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return err
	}

	defer sqlDB.Close()

	return nil
}

func (sql *MySQLDB) GetReviews() ([]models.Reviews, error) {
	var reviews []models.Reviews
	res := sql.DBhandler.Find(&reviews)
	if res.Error != nil {
		logging.Error(res.Error)
		return nil, res.Error
	}

	logging.Info(fmt.Sprintf("Rows affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return nil, err
	}

	defer sqlDB.Close()

	return reviews, nil
}

func (sql *MySQLDB) GetReviewsById(id int) (models.Reviews, error) {
	var retrievedreviews models.Reviews
	res := sql.DBhandler.Where(models.Reviews{ID: id}).First(&retrievedreviews)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.Reviews{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.Reviews{}, err
	}

	defer sqlDB.Close()

	return retrievedreviews, nil
}

func (sql *MySQLDB) GetReviewsByProductId(product_id int) (models.Reviews, error) {
	var retrievedreviews models.Reviews
	res := sql.DBhandler.Where(models.Reviews{ProductID: product_id}).First(&retrievedreviews)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.Reviews{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.Reviews{}, err
	}

	defer sqlDB.Close()

	return retrievedreviews, nil
}
