package models

import "laporan-keuangan/config"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(user *User) (int, error) {
	query := `INSERT INTO users (username, email, password) VALUES($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	return user.ID, err
}

func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, username,email, password FROM users WHERE email = $1`
	err := config.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUser(user *User) (*User, error) {
	users := &User{}
	query := `SELECT * FROM users`
	err := config.DB.QueryRow(query, user).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return users, nil
}
