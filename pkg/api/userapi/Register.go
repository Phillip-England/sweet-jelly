package userapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
	if !ok {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")
	locationNumber := r.PostForm.Get("locationNumber")
	locationNumberInt, err := util.StringToInt(locationNumber)
	if err != nil {
		http.Redirect(w, r, "/admin/users?RegisterUserFormErr=location number must be a number", http.StatusSeeOther)
		return
	}

	user := &usermod.Model{
		DB: db,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		LocationNumber: locationNumberInt,
		Password:  password,
	}
	err = database.DbExists(user)
	if err != nil {
		http.Redirect(w, r, "/admin/users?RegisterUserFormErr=email already taken", http.StatusSeeOther)
		return
	}
	err = database.DbInsert(user)
	if err != nil {
		http.Redirect(w, r, "/admin/users?RegisterUserFormErr=database connection failed", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/admin/user/%d", user.ID), http.StatusSeeOther)
}