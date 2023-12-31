package userapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/sessionmod"
	"cfasuite/pkg/mw"
	"database/sql"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
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
	session := &sessionmod.Model{
		DB: db,
		Token: sessionToken,
	}
	err = database.DbDelete(session)
	if err != nil {
		http.Error(w, "failed to delete session", http.StatusInternalServerError)
		return
	}
    cookie := &http.Cookie{
        Name:     "session-token",
        Value:    "",
        HttpOnly: true,
        Path:     "/",
		MaxAge: -1,
    }
    http.SetCookie(w, cookie)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}