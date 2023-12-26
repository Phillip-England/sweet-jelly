package pkg

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
)

func HandlerPublicFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/public/"):]
	fullPath := filepath.Join(".", "public", filePath)
	http.ServeFile(w, r, fullPath)
}

// generates html for the login page and handles 404s
func HandlerViewLogin(w http.ResponseWriter, r *http.Request) {
	// handling all 404s
	if r.URL.Path != "/" {
		b := PageBuilder{
			Title: "CFA Suite - 404 Not Found",
		}
		components := []string{
			"<h2>404 Not Found</h2>",
		}
		b.AddComponents(components)
		w.Write(b.HtmlBytes())
		return
	}
	// handling login page / index page
	b := PageBuilder{
		Title: "CFA Suite - Login",
	}
	components := []string{
		"<h1>Login Page</h1>",
		LoginForm(r.URL.Query().Get("LoginFormErr")),
	}
	b.AddComponents(components)
	w.Write(b.HtmlBytes())
}

// generates html for the admin home page
func HandlerViewAdminHome(w http.ResponseWriter, r *http.Request) {
	b := PageBuilder{
		Title: "CFA Suite - Home",
	}
	components := []string{
		"<h1>Admin Home Page</h1>",
		AdminNav(),
	}
	b.AddComponents(components)
	w.Write(b.HtmlBytes())
}

// generates html to view all users and create users
func HandlerViewAdminUsers(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(dbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    userRepo := &UserRepository{}
    err := DbGetAll(userRepo, db)
    if err != nil {
        http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
        return
    }

    // Sort users alphabetically by first name
    sort.Slice(userRepo.users, func(i, j int) bool {
        return userRepo.users[i].firstName < userRepo.users[j].firstName
    })

    b := PageBuilder{
        Title: "CFA Suite - Users",
    }
    components := []string{
        "<h1>Admin Users Page</h1>",
        AdminNav(),
        RegisterUserForm(r.URL.Query().Get("RegisterUserFormErr")),
        UserList(userRepo.users),
    }

    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}


// generates html to view a single user
func HandlerViewAdminUser(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Path[len("/admin/user/"):]
    userIDInt, err := strconv.Atoi(userID)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
	updateUserFormErr := r.URL.Query().Get("UpdateUserFormErr")
    db, ok := r.Context().Value(dbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    user := &User{
        id: userIDInt,
    }
    err = DbGetById(user, db, userIDInt)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get user details: %s", err), http.StatusInternalServerError)
        return
    }
    b := PageBuilder{
        Title: "CFA Suite - User Details",
    }

    components := []string{
        AdminNav(),
        UserDetails(user),
		UpdateUserForm(user, r, updateUserFormErr),
		UserPhoto(user),
    }
    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}