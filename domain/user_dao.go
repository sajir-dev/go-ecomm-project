package domain

import (
	"errors"
	"fmt"
	"time"

	"../db/config"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) CreateUser(username string, password string, dob time.Time, phone string) (*User, error) {
	// validate the username availability
	var user *User
	row := config.DB.QueryRow(`SELECT * FROM users WHERE username= $1`, username)
	row.Scan(&user.UserID, &user.Username, &user.Password, &user.Dob, &user.Phone)

	if user.Username == username {
		return nil, errors.New("username not available")
	}
	// Store a user in the config.DB

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err = config.DB.Exec(`INSERT INTO users (username, password, dob, phone) VALUES ($1, $2, $3, $4)`, username, pw, dob, phone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	row = config.DB.QueryRow(`SELECT * FROM users WHERE username=$1`, username)
	err = row.Scan(&user.UserID, &user.Username, &user.Password, &user.Dob, &user.Phone)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
