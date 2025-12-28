package models

import "laporan-keuangan/config"

type Book struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Tahun   int    `json:"tahun"`
}

func CreateBook(book *Book) error {
	query := `INSERT INTO books (user_id, judul, penulis, tahun) VALUES($1, $2, $3, $4) RETURNING id`
	return config.DB.QueryRow(query, book.UserID, book.Judul, book.Penulis, book.Tahun).Scan(&book.ID)

}

func GetAllBook() ([]Book, error) {
	rows, err := config.DB.Query(`SELECT id, user_id, judul, penulis, tahun FROM books`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.UserID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
			continue
		}
		books = append(books, b)
	}
	return books, nil
}

func GetBooksByUserID(userID int) ([]Book, error) {
	query := `SELECT id, user_id, judul, penulis, tahun FROM books WHERE user_id = $1`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.UserID, &b.Judul, &b.Penulis, &b.Tahun); err != nil {
			continue
		}
		books = append(books, b)
	}
	return books, nil
}
