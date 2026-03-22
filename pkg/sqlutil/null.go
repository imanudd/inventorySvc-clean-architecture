package sqlutil

import (
	"database/sql"
	"time"
)

func NewNull[T any](value T) any {
	switch v := any(value).(type) {
	case int:
		return NullInt64FromInt(v)
	case int64:
		return NullInt64(v)
	case string:
		return NullString(v)
	case bool:
		return NullBool(v)
	case float64:
		return NullFloat64(v)
	case time.Time:
		return NullTime(v)
	}

	return nil
}

func NewNullPtr[T any](value *T) any {
	switch v := any(value).(type) {
	case *int:
		return NullInt64FromIntPtr(v)
	case *int64:
		return NullInt64(*v)
	case *string:
		return NullStringPtr(v)
	case *bool:
		return NullBoolPtr(v)
	case *float64:
		return NullFloat64Ptr(v)
	case *time.Time:
		return NullTimePtr(v)
	}

	return nil
}

func NullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: true}
}

func NullInt64FromInt(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: true}
}

func NullInt64FromIntPtr(i *int) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{Valid: false}
	}
	return sql.NullInt64{Int64: int64(*i), Valid: true}
}

func NullInt64Optional(i int64) sql.NullInt64 {
	return sql.NullInt64{Int64: i, Valid: i > 0}
}

func NullInt64FromIntOptional(i int) sql.NullInt64 {
	return sql.NullInt64{Int64: int64(i), Valid: i > 0}
}

func NullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: true}
}

func NullStringPtr(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: *s, Valid: *s != ""}
}

func NullStringOptional(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

func NullBool(b bool) sql.NullBool {
	return sql.NullBool{Bool: b, Valid: true}
}

func NullBoolPtr(b *bool) sql.NullBool {
	if b == nil {
		return sql.NullBool{Valid: false}
	}
	return sql.NullBool{Bool: *b, Valid: true}
}

func NullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: f, Valid: true}
}

func NullFloat64Ptr(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{Valid: false}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

func NullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func NullTimePtr(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: !t.IsZero()}
}

func NullTimeOptional(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: !t.IsZero()}
}

func Int64FromNullInt64(n sql.NullInt64) int64 {
	if !n.Valid {
		return 0
	}
	return n.Int64
}

func StringFromNullString(n sql.NullString) string {
	if !n.Valid {
		return ""
	}
	return n.String
}

func BoolFromNullBool(n sql.NullBool) bool {
	if !n.Valid {
		return false
	}
	return n.Bool
}

func Float64FromNullFloat64(n sql.NullFloat64) float64 {
	if !n.Valid {
		return 0
	}
	return n.Float64
}

func TimeFromNullTime(n sql.NullTime) time.Time {
	if !n.Valid {
		return time.Time{}
	}
	return n.Time
}
