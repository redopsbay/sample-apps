package db

import (
	"fmt"
	"context"
	"gorm.io/gorm"
	"ratings/logging"
	"ratings/models"
)

//MySQLDB holds the necessary methods and database instance to apply *gorm.DBhandler (SQL) operations

type MySQLDB struct {
	DBhandler *gorm.DB
}

// Create applies the given struct or model to be applied on a database table by utilizing *gorm.DBhandler
func (sql *MySQLDB) AddRatings(ctx context.Context, ratings *models.Ratings) error {
	res := sql.DBhandler.WithContext(ctx).Create(ratings)
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

func (sql *MySQLDB) UpdateRatings(ctx context.Context, ratings *models.Ratings) error {

	//create models.Ratings
	var ratingsUpdate models.Ratings

	sql.DBhandler.WithContext(ctx).First(&ratingsUpdate)
	//Point it to ratingsUpdate
	res := sql.DBhandler.Save(ratings)

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

func (sql *MySQLDB) DeleteRating(ctx context.Context, ratings *models.Ratings) error {
	res := sql.DBhandler.WithContext(ctx).Delete(ratings)

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

func (sql *MySQLDB) GetRatings(ctx context.Context) ([]models.Ratings, error) {
	var ratings []models.Ratings
	res := sql.DBhandler.WithContext(ctx).Find(&ratings)
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

	return ratings, nil
}

func (sql *MySQLDB) GetRatingsById(ctx context.Context, id int) (models.Ratings, error) {
	var retrievedRatings models.Ratings
	res := sql.DBhandler.WithContext(ctx).Where(models.Ratings{ID: id}).First(&retrievedRatings)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.Ratings{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.Ratings{}, err
	}

	defer sqlDB.Close()

	return retrievedRatings, nil
}

func (sql *MySQLDB) GetRatingsByProductId(ctx context.Context, product_id int) (models.Ratings, error) {
	var retrievedRatings models.Ratings
	res := sql.DBhandler.WithContext(ctx).Where(models.Ratings{ProductID: product_id}).First(&retrievedRatings)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.Ratings{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.Ratings{}, err
	}

	defer sqlDB.Close()

	return retrievedRatings, nil
}
