package userapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/sessionmod"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"database/sql"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
	if !ok {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if (email == adminEmail && password == adminPassword) {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
    user := &usermod.Model{
		DB: db,
	}
    err := user.GetByEmail(email)
    if err != nil {
        http.Redirect(w, r, "/?LoginFormErr=invalid credentials", http.StatusSeeOther)
        return        
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/?LoginFormErr=invalid credentials", http.StatusSeeOther)
		return
	}
	session := &sessionmod.Model{
		DB: db,
		UserID: user.ID,
	}
	err = database.DbInsert(session)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{
		Name: "session-token",
		Value: session.Token,
		HttpOnly: true,
		Path: "/",
		// make this one expire in 1 hour
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/app", http.StatusSeeOther)
}