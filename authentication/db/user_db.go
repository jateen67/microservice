package db

type UserDB interface {
	GetUserByEmail(email string) (*User, error)
}
