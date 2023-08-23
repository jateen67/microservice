package db

import "database/sql"

type UserDBImpl struct {
	DB *sql.DB
}

func NewUserDBImpl(db *sql.DB) *UserDBImpl {
	return &UserDBImpl{DB: db}
}

func (u *UserDBImpl) GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, email, password FROM users WHERE email = $1"
	var user User
	err := u.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
