package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"productservice/crypt"
	"productservice/db"
	"productservice/models"
)

type Login struct {
	Username string `json:username`
	Password string `json:password`
}

// GET /v1/user

func GetUsers(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	users, err := sql.GetUsers()
	if err != nil {
		ServerError(c)
	}

	c.JSON(http.StatusOK, users)
}

// GET /v1/user/<int:id>

func GetUserById(c *gin.Context) {
	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var userID models.User

	if err := c.ShouldBindUri(&userID); err == nil {
		user, err := sql.GetUserById(userID.ID)
		if err != nil {
			ServerError(c)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// DELETE /v1/user/<int:id>

func DeleteUser(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var user models.User

	if c.ShouldBindJSON(&user) == nil {
		if sql.DeleteUser(&user) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}

// PUT /v1/user/<int:id>

func UpdateUser(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var user models.User

	if c.ShouldBindJSON(&user) == nil {
		if sql.UpdateUser(&user) == nil {
			SuccessCreated(c)
		} else {
			ServerError(c)
		}

	} else {
		ServerError(c)
	}

}

// POST /v1/user/create

func CreateUser(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var user models.User

	if c.ShouldBindJSON(&user) == nil {
		uuid := uuid.NewString()
		user.UUID = uuid
		hashpasswd, err := crypt.HashPassword(user.Password)
		if err == nil {
			if crypt.ValidatePasswordHash(user.Password, hashpasswd) {

				user.Password = hashpasswd

				if sql.CreateUser(&user) == nil {
					SuccessCreated(c)
				} else {
					ServerError(c)
				}
			}
		}

	} else {
		ServerError(c)
	}

}

func UserLogin(c *gin.Context) {

	dbInstance, err := db.SetupDatabase()

	if err != nil {
		ServerError(c)
	}

	sql := db.MySQLDB{DBhandler: dbInstance}

	var user models.User
	var loginUser Login

	if c.ShouldBindJSON(&loginUser) == nil {
		hashpasswd, err := crypt.HashPassword(loginUser.Password)
		if err == nil {
			if crypt.ValidatePasswordHash(loginUser.Password, hashpasswd) {
				user.Password = hashpasswd
				user.Username = loginUser.Username
				if sql.ValidateHashPassword(user.Password) == nil {
					Success(c)
				}
			}
		}

	} else {
		ServerError(c)
	}

}
