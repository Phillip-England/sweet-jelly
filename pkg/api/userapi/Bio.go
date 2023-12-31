package userapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"database/sql"
	"net/http"
)

func Bio(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "failed to get database connection", http.StatusInternalServerError)
        return
    }
	user, ok := r.Context().Value(mw.AuthKey).(*usermod.Model)
	if !ok {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

    r.ParseForm()

    family := r.PostForm.Get("family")
    hobbies := r.PostForm.Get("hobbies")
    dreams := r.PostForm.Get("dreams")

	user.DB = db
	user.Family = family
	user.Hobbies = hobbies
	user.Dreams = dreams

	err := database.DbUpdate(user)
	if err != nil {
		http.Redirect(w, r, "/app/bio?BioFormErr=db operation failed", http.StatusSeeOther)
		return
	}

    // Redirect to the previous URL
    previousURL := r.URL.Query().Get("previousURL")
    http.Redirect(w, r, previousURL, http.StatusSeeOther)
}