package main

import (
	"cfasuite/pkg/api/locationapi"
	"cfasuite/pkg/api/userapi"
	"cfasuite/pkg/database"
	"cfasuite/pkg/mw"
	"cfasuite/pkg/view/adminview"
	"cfasuite/pkg/view/appview"
	"cfasuite/pkg/view/guestview"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func HandlerPublicFiles(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/public/"):]
	fullPath := filepath.Join(".", "public", filePath)
	http.ServeFile(w, r, fullPath)
}

func main() {

	// tailwindCommand := exec.Command("tailwind", "-i", "./public/css/input.css", "-o", "./public/css/output.css")
	// _, _ = tailwindCommand.CombinedOutput()

	formatCode := exec.Command("go", "fmt", "./")
	err := formatCode.Run()
	if err != nil {
		panic("failed to format code")
	}

	tidyCode := exec.Command("go", "mod", "tidy")
	err = tidyCode.Run()
	if err != nil {
		panic("failed to tidy code")
	}

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := database.Init(os.Getenv("RESET_DB"))
	if err != nil {
		fmt.Println(err)
		panic("database failed to connect")
	}
	defer db.Close()

	// serving public files
	http.HandleFunc("/public/", HandlerPublicFiles)

	// guest views
	http.HandleFunc("/", guestview.Login)

	// admin views
	http.HandleFunc("/admin", adminview.Home)
	http.HandleFunc("/admin/users", mw.MwDb(db, adminview.Users))
	http.HandleFunc("/admin/user/", mw.MwDb(db, adminview.User))
	http.HandleFunc("/admin/locations", mw.MwDb(db, adminview.Locations))
	http.HandleFunc("/admin/location/", mw.MwDb(db, adminview.Location))

	// app views
	http.HandleFunc("/app", mw.MwDb(db, mw.Auth(appview.Home)))
	http.HandleFunc("/app/bio", mw.MwDb(db, mw.Auth(appview.Bio)))
	http.HandleFunc("/app/peers", mw.MwDb(db, mw.Auth(appview.Peers)))
	http.HandleFunc("/app/peer/", mw.MwDb(db, mw.Auth(appview.Peer)))

	// user api
	http.HandleFunc("/api/user/login", mw.MwDb(db, userapi.Login))
	http.HandleFunc("/api/user/bio", mw.MwDb(db, mw.Auth(userapi.Bio)))
	http.HandleFunc("/api/user/logout", mw.MwDb(db, userapi.Logout))
	http.HandleFunc("/api/user/register", mw.MwDb(db, userapi.Register))
	http.HandleFunc("/api/user/update", mw.MwDb(db, userapi.Update))

	// location api
	http.HandleFunc("/api/location/create", mw.MwDb(db, locationapi.Create))

	// running
	http.ListenAndServe(":8080", nil)

}
