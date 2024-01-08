package appview

import (
	"net/http"
)

func Peers(w http.ResponseWriter, r *http.Request) {
    // db, ok := r.Context().Value(mw.DbKey).(*sql.DB)
    // if !ok {
    //     http.Error(w, "Failed to get database connection", http.StatusInternalServerError)
    //     return
    // }
	// user, ok := r.Context().Value(mw.AuthKey).(*usermod.Model)
	// if !ok {
	// 	http.Error(w, "Failed to get user", http.StatusInternalServerError)
	// 	return
	// }
    // b := util.PageBuilder{
    //     Title: "CFA Suite - User Details",
    // }
	// userRepo := usermod.NewUserRepo(db)
	// err := userRepo.GetAllByLocationNumber(user.LocationNumber)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("failed to get users by location: %s", err.Error()), http.StatusInternalServerError)
	// 	return
	// }

	// // Remove the current user from the list
	// var filteredUsers []usermod.Model
	// for _, u := range userRepo.Users {
	// 	if u.ID != user.ID {
	// 		filteredUsers = append(filteredUsers, u)
	// 	}
	// }

    // components := []string{
	// 	comp.Header("App Peers Page"),
	// 	comp.TeamNav(),
	// 	comp.UserDetails(user),
	// 	comp.PeerList(filteredUsers), // Use the new PeerList component with the filtered user list
    // }
    // b.AddComponents(components)
    // w.Write(b.HtmlBytes())
}
