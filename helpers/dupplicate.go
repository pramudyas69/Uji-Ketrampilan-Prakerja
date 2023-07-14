package helpers

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsDuplicateError digunakan untuk memeriksa apakah error terkait duplikat data
func IsDuplicateError(err error) bool {
	var pgError *pgconn.PgError
	if errors.As(err, &pgError) {
		return pgError.Code == "23505" // Kode error PostgreSQL untuk kesalahan duplikat data
	}
	return false
}
