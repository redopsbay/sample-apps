package db

import (
	"productservice/logging"
	"productservice/models"
	"fmt"
	"context"
	"gorm.io/gorm"
)

//MySQLDB holds the necessary methods and database instance to apply *gorm.DBhandler (SQL) operations

type MySQLDB struct {
	DBhandler *gorm.DB
}

// Create applies the given struct or model to be applied on a database table by utilizing *gorm.DBhandler
func (sql *MySQLDB) CreateProduct(ctx context.Context, product *models.Product) error {
	res := sql.DBhandler.WithContext(ctx).Create(product)
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

func (sql *MySQLDB) UpdateProduct(ctx context.Context, product *models.Product) error {

	//create models.Product
	var productUpdate models.Product

	sql.DBhandler.WithContext(ctx).First(&productUpdate)
	//Saved copied objects
	res := sql.DBhandler.Save(product)

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

func (sql *MySQLDB) DeleteProduct(ctx context.Context, product *models.Product) error {
	res := sql.DBhandler.WithContext(ctx).Delete(product)

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

func (sql *MySQLDB) GetProducts(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	res := sql.DBhandler.WithContext(ctx).Find(&products)
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

	return products, nil
}

func (sql *MySQLDB) GetProductById(ctx context.Context, id int) (models.Product, error) {
	var retrievedProduct models.Product
	res := sql.DBhandler.WithContext(ctx).Where(models.Product{ID: id}).First(&retrievedProduct)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.Product{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.Product{}, err
	}

	defer sqlDB.Close()

	return retrievedProduct, nil
}

func (sql *MySQLDB) GetUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	res := sql.DBhandler.WithContext(ctx).Find(&users)
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

	return users, nil
}

func (sql *MySQLDB) GetUserById(ctx context.Context, id int) (models.User, error) {
	var retrievedUser models.User
	res := sql.DBhandler.WithContext(ctx).Where(models.User{ID: id}).First(&retrievedUser)
	if res.Error != nil {
		logging.Error(res.Error)
		return models.User{}, res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return models.User{}, err
	}

	defer sqlDB.Close()

	return retrievedUser, nil
}

func (sql *MySQLDB) DeleteUser(ctx context.Context, user *models.User) error {
	res := sql.DBhandler.WithContext(ctx).Delete(user)

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

func (sql *MySQLDB) UpdateUser(ctx context.Context, user *models.User) error {

	//create models.Product
	var UserUpdate models.User

	sql.DBhandler.WithContext(ctx).First(&UserUpdate)
	*user = UserUpdate
	//Point it to productUpdate
	sql.DBhandler.Save(&UserUpdate)

	//Copy data from argument
	UserUpdate = *user

	//Saved copied objects
	res := sql.DBhandler.Save(UserUpdate)

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

func (sql *MySQLDB) CreateUser(ctx context.Context, user *models.User) error {
	res := sql.DBhandler.WithContext(ctx).Create(user)
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

func (sql *MySQLDB) ValidateHashPassword(ctx context.Context, hashpasswd string) error {
	var user models.User
	res := sql.DBhandler.WithContext(ctx).Where(models.User{Password: hashpasswd}).First(user)
	if res.Error != nil {
		logging.Error(res.Error)
		return res.Error
	}

	logging.Info(fmt.Sprintf("Rows Affected: %d", res.RowsAffected))

	sqlDB, err := sql.DBhandler.DB()

	if err != nil {
		logging.Error(err)
		return err
	}

	defer sqlDB.Close()

	return nil
}
