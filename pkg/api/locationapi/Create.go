package locationapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/mw"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    r.ParseForm()
    name := r.PostForm.Get("name")
    numberStr := r.PostForm.Get("number")
    number, err := strconv.Atoi(numberStr)
    if err != nil {
        http.Error(w, "Invalid location number", http.StatusBadRequest)
        return
    }
    location := &locationmod.Model{
        DB: db,
        Name:   name,
        Number: number,
    }
    err = database.DbInsert(location)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to create location: %s", err), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/admin/locations", http.StatusSeeOther)
}