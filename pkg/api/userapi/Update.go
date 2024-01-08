package userapi

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    
    previousURL := r.URL.Query().Get("previousURL")
    redirectBadLocationNumberURL := util.AppendQueryParam(previousURL, "UpdateUserFormErr", "must user valid location number")


    r.ParseMultipartForm(10 << 20)
    photo, _, err := r.FormFile("photo")
    var photoBytes []byte

    // Check if a new photo is provided
    if err == nil {
        defer photo.Close()
        photoBytes, err = io.ReadAll(photo)
        if err != nil {
            http.Error(w, "Failed to read the photo file", http.StatusInternalServerError)
            return
        }
    }

    idStr := r.PostForm.Get("id")

    // Convert the ID string to an integer
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    email := r.PostForm.Get("email")
    firstName := r.PostForm.Get("firstName")
    lastName := r.PostForm.Get("lastName")
    family := r.PostForm.Get("family")
    hobbies := r.PostForm.Get("hobbies")
    dreams := r.PostForm.Get("dreams")
    locationNumber, err := util.StringToInt(r.PostForm.Get("locationNumber"))
    if err != nil {
        http.Redirect(w, r, redirectBadLocationNumberURL, http.StatusSeeOther)
        return
    }

    // Fetch the user from the database to get the existing photo
    u := &usermod.Model{
        DB: db,
        ID: id,
    }
    err = database.DbGetById(u)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get user details: %s", err), http.StatusInternalServerError)
        return
    }

    // Check if a new photo is provided; if not, keep the existing one
    if len(photoBytes) == 0 {
        photoBytes = u.Photo
    }

    // Create a new User object with the updated information
    updatedUser := usermod.Model {
        DB: db,
        ID:        id,
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
        Family: family,
        Hobbies: hobbies,
        Dreams: dreams,
        LocationNumber: locationNumber,
        Photo:     photoBytes,
    }

    // Use DbUpdate to update the user in the database
    err = database.DbUpdate(&updatedUser)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update user: %s", err), http.StatusInternalServerError)
        return
    }

    // Redirect to the previous URL
    http.Redirect(w, r, previousURL, http.StatusSeeOther)
}