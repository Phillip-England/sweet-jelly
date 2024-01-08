package adminview

import (
	"cfasuite/pkg/database"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/util"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type UserData struct {
    Title string
    UpdateUserFormErr string
    User usermod.Model
    PreviousURL string
}

func User(w http.ResponseWriter, r *http.Request) {
    userID := r.URL.Path[len("/admin/user/"):]
    userIDInt, err := strconv.Atoi(userID)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    if !ok {
        http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
        return
    }
    user := usermod.NewUserModel(db)
    user.ID = userIDInt
    err = database.DbGetById(user)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get user details: %s", err), http.StatusInternalServerError)
        return
    }
    util.RenderTemplate(w, "./pkg/view/adminview/User.html", UserData{
        Title: "CFA Suite - User",
        UpdateUserFormErr: r.URL.Query().Get("UpdateUserFormErr"),
        User: *user,
        PreviousURL: r.URL.Path,
    })

    // b := util.PageBuilder{
    //     Title: "CFA Suite - User Details",
    // }

    // components := []string{
	// 	comp.Header("Admin User Page"),
    //     comp.AdminNav(),
    //     comp.UserDetails(user),
	// 	comp.UpdateUserForm(user, r, updateUserFormErr),
	// 	comp.UserPhoto(user),
    // }
    // b.AddComponents(components)
    // w.Write(b.HtmlBytes())
}