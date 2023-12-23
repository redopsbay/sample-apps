package db

import (
	"fmt"
	"gorm.io/gorm"
	"reviews/logging"
	"reviews/models"
	"context"
)

//MySQLDB holds the necessary methods and database instance to apply *gorm.DBhandler (SQL) operations

type MySQLDB struct {
	DBhandler *gorm.DB
}

// Create applies the given struct or model to be applied on a database table by utilizing *gorm.DBhandler
func (sql *MySQLDB) AddReviews(ctx context.Context, reviews *models.Reviews) error {
	res := sql.DBhandler.WithContext(ctx).Create(reviews)
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

func (sql *MySQLDB) UpdateReviews(ctx context.Context, reviews *models.Reviews) error {

	//create models.Reviews
	var reviewsUpdate models.Reviews

	sql.DBhandler.WithContext(ctx).First(&reviewsUpdate)
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

func (sql *MySQLDB) DeleteReviews(ctx context.Context, reviews *models.Reviews) error {
	res := sql.DBhandler.WithContext(ctx).Delete(reviews)

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

func (sql *MySQLDB) GetReviews(ctx context.Context) ([]models.Reviews, error) {
	var reviews []models.Reviews
	res := sql.DBhandler.WithContext(ctx).Find(&reviews)
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

func (sql *MySQLDB) GetReviewsById(ctx context.Context, id int) (models.Reviews, error) {
	var retrievedreviews models.Reviews
	res := sql.DBhandler.WithContext(ctx).Where(models.Reviews{ID: id}).First(&retrievedreviews)
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

func (sql *MySQLDB) GetReviewsByProductId(ctx context.Context, product_id int) (models.Reviews, error) {
	var retrievedreviews models.Reviews
	res := sql.DBhandler.WithContext(ctx).Where(models.Reviews{ProductID: product_id}).First(&retrievedreviews)
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
