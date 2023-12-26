package pkg

import (
	"context"
	"database/sql"
	"net/http"
)

type contextKey string
const dbKey contextKey = "db"

// injects the db into the http context
func MwDb(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r.WithContext(context.WithValue(r.Context(), dbKey, db)))
	}
}