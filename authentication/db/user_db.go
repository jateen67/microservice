package db

type UserDB interface {
	GetUserByEmail(email string) (*User, error)
	PasswordCheck(hashedPassword, plainPassword string) error
}
