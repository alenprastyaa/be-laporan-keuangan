package controllers

import (
	"laporan-keuangan/models"
	"laporan-keuangan/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetBudgetReport(c *fiber.Ctx) error {
	userVal := c.Locals("user_id")
	if userVal == nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "User ID tidak valid")
	}
	userID := int(userVal.(float64))
	month := c.QueryInt("month", 0)
	year := c.QueryInt("year", 0)
	summary, err := models.GetBudgetSummary(userID, month, year)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal mengambil data: "+err.Error())
	}
	return utils.ResponseSuccess(c, summary, "Laporan anggaran berhasil diambil")
}
func AddBudgetEntry(c *fiber.Ctx) error {
	var entry models.BudgetEntry
	if err := c.BodyParser(&entry); err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Format JSON salah. Cek tanda koma atau kurung.")
	}
	userVal := c.Locals("user_id")
	if userVal == nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid")
	}
	entry.UserID = int(userVal.(float64))
	entry.Jenis = strings.ToLower(entry.Jenis)

	if entry.Jenis != "pemasukan" && entry.Jenis != "pengeluaran" {
		return utils.ResponseError(c, fiber.StatusBadRequest, "Jenis harus 'pemasukan' atau 'pengeluaran'")
	}
	if err := models.CreateBudgetEntry(&entry); err != nil {
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal simpan ke DB: "+err.Error())
	}
	return utils.ResponseCreated(c, entry, "Item berhasil ditambahkan")
}

func DeleteBudgetEntry(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ResponseError(c, fiber.StatusBadRequest, "ID tidak valid")
	}
	userVal := c.Locals("user_id")
	if userVal == nil {
		return utils.ResponseError(c, fiber.StatusUnauthorized, "Token tidak valid")
	}
	userID := int(userVal.(float64))

	err = models.DeleteBudgetEntry(id, userID)
	if err != nil {
		if err.Error() == "data tidak ditemukan atau bukan milik anda" {
			return utils.ResponseError(c, fiber.StatusNotFound, err.Error())
		}
		return utils.ResponseError(c, fiber.StatusInternalServerError, "Gagal menghapus data")
	}

	return utils.ResponseSuccess(c, nil, "Item berhasil dihapus")
}
