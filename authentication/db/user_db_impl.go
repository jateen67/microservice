package db

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserDBImpl struct {
	DB *sql.DB
}

func NewUserDBImpl(db *sql.DB) *UserDBImpl {
	return &UserDBImpl{DB: db}
}

func (u *UserDBImpl) GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, email, password, first_name, last_name, created_at FROM users WHERE email = $1"
	var user User
	err := u.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.CreatedAt)
	if err != nil {
		log.Println("error getting user by email:", err)
		return nil, err
	}
	return &user, nil
}

func (u *UserDBImpl) PasswordCheck(hashedPassword string, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println("error validating user:", err)
		return err
	}
	return nil
}
