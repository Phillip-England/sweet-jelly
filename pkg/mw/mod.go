package mw

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/sessionmod"
	"cfasuite/pkg/model/usermod"
	"context"
	"database/sql"
	"net/http"
	"os"
)

type contextKey string
const DbKey contextKey = "db"
const AuthKey contextKey = "auth"


// runs auth and ensures an admin is using the associate endpoint
func AdminAuth(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		db, ok := r.Context().Value(DbKey).(*sql.DB)
		if !ok {
			http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
			return
		}
		sessionTokenCookie, err := r.Cookie("session-token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		sessionToken := sessionTokenCookie.Value
		if sessionToken != os.Getenv("ADMIN_SESSION_TOKEN") {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), DbKey, db)
		next(w, r.WithContext(ctx))
	}
}

// injects the db into the http context
func MwDb(db *sql.DB, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r.WithContext(context.WithValue(r.Context(), DbKey, db)))
	}
}

// runs auth and attaches the current user to the context
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		db, ok := r.Context().Value(DbKey).(*sql.DB)
		if !ok {
			http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
			return
		}
		sessionTokenCookie, err := r.Cookie("session-token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		sessionToken := sessionTokenCookie.Value
		session := sessionmod.NewSessionModel(db)
		session.Token = sessionToken
		err = session.GetByToken()
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		user := usermod.NewUserModel(db)
		user.ID = session.UserID
		err = database.DbGetById(user)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		ctx := context.WithValue(r.Context(), DbKey, db)
		ctx = context.WithValue(ctx, AuthKey, user)
		next(w, r.WithContext(ctx))
	}
}