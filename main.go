package main

import (
	"cfasuite/pkg"
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	// tailwindCommand := exec.Command("tailwind", "-i", "./public/css/input.css", "-o", "./public/css/output.css")
	// _, _ = tailwindCommand.CombinedOutput()

	// dotenv
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// db connection
	db, err := pkg.ConnectDb()
	if err != nil {
		message := fmt.Sprintf("failed to connect to db because of this err: %s", err.Error())
		fmt.Println(message)
		return
	}
	defer db.Close()

	initDb := false

	if initDb {

		// deleting all tables
		err = pkg.DbDeleteAllTables(db)
		if err != nil {
			message := fmt.Sprintf(`failed to delete all tables in db due to: %s`, err.Error())
			fmt.Println(message)
			return
		}
	
		// creating user table
		err = pkg.DbCreateUserTable(db)
		if err != nil {
			message := fmt.Sprintf(`failed to create user table due to this error: %s`, err.Error())
			fmt.Println(message)
			return
		}

	}

	// serving public files
	http.HandleFunc("/public/", pkg.HandlerPublicFiles)

	// serving views
	http.HandleFunc("/", pkg.HandlerViewLogin)
	http.HandleFunc("/admin", pkg.HandlerViewAdminHome)
	http.HandleFunc("/admin/users", pkg.MwDb(db, pkg.HandlerViewAdminUsers))
	http.HandleFunc("/admin/user/", pkg.MwDb(db, pkg.HandlerViewAdminUser))


	// serving api
	http.HandleFunc("/api/user/register", pkg.MwDb(db, pkg.HandlerApiRegisterUser))
	http.HandleFunc("/api/user/update", pkg.MwDb(db, pkg.HandlerApiUpdateUser))
	http.HandleFunc("/api/user/login", pkg.MwDb(db, pkg.HandlerApiLogin))
	http.HandleFunc("/api/user/logout", pkg.HandlerApiLogout)

	// running
	http.ListenAndServe(":8080", nil)

	
}
