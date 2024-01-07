package adminview

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"net/http"
	"sort"
)

type UsersData struct {
    Title string
    RegisterUserErr string
    Users []usermod.Model
    NoPhoto bool
}

func Users(w http.ResponseWriter, r *http.Request) {
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    userRepo := usermod.NewUserRepo(db)
    err := database.DbGetAll(userRepo)
    if err != nil {
        http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
        return
    }
    sort.Slice(userRepo.Users, func(i, j int) bool {
        return userRepo.Users[i].FirstName < userRepo.Users[j].FirstName
    })
    util.RenderTemplate(w, "./pkg/view/adminview/Users.html", UsersData{
        Title: "",
        RegisterUserErr: r.URL.Query().Get("RegisterUserErr"),
        Users: userRepo.Users,
    })
}