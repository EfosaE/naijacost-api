package util

import "github.com/jackc/pgx/v5/pgtype"

func ToFloat8(f float64) pgtype.Float8 {
	return pgtype.Float8{
		Float64: f,
		Valid:   true,
	}
}
