package pkg

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HandlerApiLogin(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(dbKey).(*sql.DB)
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
		sessionToken := os.Getenv("ADMIN_SESSION_TOKEN")
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
    user := &User{}
    err := user.GetByEmail(db, email)
    if err != nil {
        http.Redirect(w, r, "/?LoginFormErr=invalid credentials", http.StatusSeeOther)
        return        
    }
    err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/?LoginFormErr=invalid credentials", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandlerApiLogout(w http.ResponseWriter, r *http.Request) {
    cookie := &http.Cookie{
        Name:     "session_token",
        Value:    "",
        HttpOnly: true,
        Path:     "/",
    }
    http.SetCookie(w, cookie)
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandlerApiRegisterUser(w http.ResponseWriter, r *http.Request) {
	db, ok := r.Context().Value(dbKey).(*sql.DB)
	if !ok {
		http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
		return
	}
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")
	user := &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
	}
	err := DbExists(user, db)
	if err != nil {
		http.Redirect(w, r, "/admin/users?RegisterUserFormErr=email already taken", http.StatusSeeOther)
		return
	}
	err = DbInsert(user, db)
	if err != nil {
		http.Redirect(w, r, "/admin/users?RegisterUserFormErr=database connection failed", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/admin/user/%d", user.id), http.StatusSeeOther)
}

func HandlerApiUpdateUser(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(dbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }

    // Parse the form data, including the file upload
    r.ParseMultipartForm(10 << 20) // 10 MB limit for the photo file
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

    // Fetch the user from the database to get the existing photo
    user := &User{
        id: id,
    }
    err = DbGetById(user, db, id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get user details: %s", err), http.StatusInternalServerError)
        return
    }

    // Check if a new photo is provided; if not, keep the existing one
    if len(photoBytes) == 0 {
        photoBytes = user.photo
    }

    // Create a new User object with the updated information
    updatedUser := User{
        id:        id,
        firstName: firstName,
        lastName:  lastName,
        email:     email,
        photo:     photoBytes,
    }

    // Use DbUpdate to update the user in the database
    err = DbUpdate(&updatedUser, db, id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to update user: %s", err), http.StatusInternalServerError)
        return
    }

    // Redirect to the previous URL
    previousURL := r.URL.Query().Get("previousURL")
    http.Redirect(w, r, previousURL, http.StatusSeeOther)
}