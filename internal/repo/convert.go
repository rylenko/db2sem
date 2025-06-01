package repo

import (
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func convertToPgBool(value *bool) pgtype.Bool {
	var pgBool pgtype.Bool

	if value != nil {
		pgBool.Bool = *value
		pgBool.Valid = true
	}

	return pgBool
}

func convertToPgText(value *string) pgtype.Text {
	var pgText pgtype.Text

	if value != nil {
		pgText.String = *value
		pgText.Valid = true
	}

	return pgText
}

func convertFromPgDate(date pgtype.Date) (time.Time, error) {
	if !date.Valid {
		return time.Time{}, errors.New("invalid value")
	}

	return date.Time, nil
}

func convertFromPgText(text pgtype.Text) *string {
	if !text.Valid {
		return nil
	}

	return &text.String
}

func convertFromPgFloat8(float8 pgtype.Float8) (float64, error) {
	if !float8.Valid {
		return 0, errors.New("invalid value")
	}

	return float8.Float64, nil
}

func convertFromPgInt2(int2 pgtype.Int2) *int16 {
	var value *int16

	if int2.Valid {
		value = &int2.Int16
	}

	return value
}

func convertFromPgNumeric(numeric pgtype.Numeric) (float64, error) {
	float8, err := numeric.Float64Value()
	if err != nil {
		return 0, fmt.Errorf("get float8 value: %w", err)
	}

	return convertFromPgFloat8(float8)
}

func convertFromPgTimestamptz(tz pgtype.Timestamptz) (time.Time, error) {
	if !tz.Valid {
		return time.Time{}, errors.New("invalid tz")
	}

	return tz.Time, nil
}

func convertToPgInt2(value *int16) pgtype.Int2 {
	var int2 pgtype.Int2

	if value != nil {
		int2.Int16 = *value
		int2.Valid = true
	}

	return int2
}

func convertToPgInt8(value *int64) pgtype.Int8 {
	var pgInt pgtype.Int8

	if value != nil {
		pgInt.Int64 = *value
		pgInt.Valid = true
	}

	return pgInt
}

func convertToPgNumeric(number float64) (pgtype.Numeric, error) {
	var numeric pgtype.Numeric
	if err := numeric.Scan(fmt.Sprintf("%f", number)); err != nil {
		return numeric, fmt.Errorf("scan: %w", err)
	}

	return numeric, nil
}

func convertToPgDate(date time.Time) pgtype.Date {
	return pgtype.Date{
		Valid: true,
		Time:  date,
	}
}

func convertToPgTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:  t,
		Valid: true,
	}
}
