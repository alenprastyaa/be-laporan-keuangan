package models

import (
	"errors"
	"fmt"
	"laporan-keuangan/config"
	"time"
)

type BudgetEntry struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Kategori  string    `json:"kategori"`
	NamaItem  string    `json:"nama_item"`
	Jumlah    float64   `json:"jumlah"`
	Jenis     string    `json:"jenis"`
	CreatedAt time.Time `json:"created_at"`
}

type BudgetSummary struct {
	TotalPemasukan   float64                  `json:"total_pemasukan"`
	TotalPengeluaran float64                  `json:"total_pengeluaran"`
	SisaAnggaran     float64                  `json:"sisa_anggaran"`
	Detail           map[string][]BudgetEntry `json:"detail"`
}

func CreateBudgetEntry(entry *BudgetEntry) error {
	query := `INSERT INTO budget_entries (user_id, kategori, nama_item, jumlah, jenis) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return config.DB.QueryRow(query, entry.UserID, entry.Kategori, entry.NamaItem, entry.Jumlah, entry.Jenis).Scan(&entry.ID)
}
func GetBudgetSummary(userID int, month int, year int) (BudgetSummary, error) {
	// Base Query
	query := `SELECT id, user_id, kategori, nama_item, jumlah, jenis, created_at FROM budget_entries WHERE user_id = $1`

	args := []interface{}{userID}
	paramIdx := 2

	// Jika ada filter bulan dan tahun
	if month > 0 && year > 0 {
		query += fmt.Sprintf(" AND EXTRACT(MONTH FROM created_at) = $%d AND EXTRACT(YEAR FROM created_at) = $%d", paramIdx, paramIdx+1)
		args = append(args, month, year)
	}

	query += ` ORDER BY created_at DESC`

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		return BudgetSummary{}, err
	}
	defer rows.Close()

	summary := BudgetSummary{
		Detail: make(map[string][]BudgetEntry),
	}

	for rows.Next() {
		var b BudgetEntry
		if err := rows.Scan(&b.ID, &b.UserID, &b.Kategori, &b.NamaItem, &b.Jumlah, &b.Jenis, &b.CreatedAt); err != nil {
			continue
		}

		summary.Detail[b.Kategori] = append(summary.Detail[b.Kategori], b)

		if b.Jenis == "pemasukan" {
			summary.TotalPemasukan += b.Jumlah
		} else {
			summary.TotalPengeluaran += b.Jumlah
		}
	}

	summary.SisaAnggaran = summary.TotalPemasukan - summary.TotalPengeluaran
	return summary, nil
}

func DeleteBudgetEntry(id int, userID int) error {
	query := `DELETE FROM budget_entries WHERE id = $1 AND user_id = $2`
	result, err := config.DB.Exec(query, id, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("data tidak ditemukan atau bukan milik anda")
	}

	return nil
}
