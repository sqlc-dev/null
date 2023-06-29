package null_tests

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"github.com/sqlc-dev/null"
)

func scan[T comparable](t *testing.T, db *sql.DB, query string, expected T) {
	var actual null.Nullable[T]
	err := db.QueryRow(query).Scan(&actual)
	if err != nil {
		t.Error(err)
		return
	}
	if !actual.Valid {
		t.Error("expected nullable to be valid")
	}
	if actual.Val != expected {
		t.Errorf("want %#v; got %#v", expected, actual.Valid)
	}
}

func TestNull(t *testing.T) {
	source := "postgres://localhost?sslmode=disable"
	db, err := sql.Open("postgres", source)
	if err != nil {
		t.Fatal(err)
	}
	scan(t, db, `SELECT 1`, 1)
}
