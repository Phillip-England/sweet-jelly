package adminview

import (
	"cfasuite/pkg/comp"
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"net/http"
	"sort"
)

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

    // Sort users alphabetically by first name
    sort.Slice(userRepo.Users, func(i, j int) bool {
        return userRepo.Users[i].FirstName < userRepo.Users[j].FirstName
    })

    b := util.PageBuilder{
        Title: "CFA Suite - Users",
    }
    components := []string{
        comp.Header("Admin Users Page"),
        comp.AdminNav(),
        comp.RegisterUserForm(r.URL.Query().Get("RegisterUserFormErr")),
        comp.UserList(userRepo.Users),
    }

    b.AddComponents(components)
    w.Write(b.HtmlBytes())
}