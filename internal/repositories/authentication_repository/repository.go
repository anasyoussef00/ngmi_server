package authentication_repository

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	userModel "ngmi_server/internal/models/user_model"
	"ngmi_server/pkg/db"
	"ngmi_server/pkg/log"
)

type Repository interface {
	Register(userRegisterReq userModel.RegisterReq) (sql.Result, error)
	Login(userLoginReq userModel.LoginReq) error
}

type repository struct {
	db     *db.DB
	logger log.Logger
}

func NewRepository(db *db.DB, logger log.Logger) Repository {
	return repository{db: db, logger: logger}
}

func (r repository) Register(userRegisterReq userModel.RegisterReq) (sql.Result, error) {
	//var username string
	//selectQuery := "SELECT username FROM users WHERE username=?"
	//if err := r.db.DB().Get(&username, selectQuery, userRegisterReq.Username); err != nil {
	//	if username != "" {
	//		return nil, fmt.Errorf("%s already exists, please use a different username and try again", username)
	//	}
	//
	//
	//}
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterReq.Password), bcrypt.DefaultCost); err != nil {
		return nil, fmt.Errorf("an error has occurred while trying to hash your password: %v", err)
	} else {
		insertQuery := `INSERT INTO users (first_name, last_name, username, email_address, password, birth_date, gender) VALUES (?, ?, ?, ?, ?, ?, ?)`
		stringHashedPassword := string(hashedPassword)
		return r.db.DB().Exec(insertQuery, userRegisterReq.FirstName, userRegisterReq.LastName, userRegisterReq.Username, userRegisterReq.EmailAddress, stringHashedPassword, userRegisterReq.BirthDate, userRegisterReq.Gender)
	}
}

func (r repository) Login(userLoginReq userModel.LoginReq) error {
	var u userModel.User
	selectQuery := `SELECT id, first_name, last_name, username, email_address, password, birth_date, gender FROM users WHERE username=?`
	if err := r.db.DB().Get(&u, selectQuery, userLoginReq.Username); err != nil {
		return fmt.Errorf("an error has occurred while trying to find the user with the given username: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userLoginReq.Password), []byte(u.Password)); err != nil {
		return errors.New("password does not match, please try again")
	}

	return nil
}
