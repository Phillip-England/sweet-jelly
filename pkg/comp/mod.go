package comp

import (
	"cfasuite/pkg/model/locationmod"
	"cfasuite/pkg/model/usermod"
	"cfasuite/pkg/util"
	"encoding/base64"
	"fmt"
	"net/http"
)

func Header(text string) string {
	return fmt.Sprintf(`<h1>%s</h1>`, text)
}

func LoginForm(formErr string) string {
	var emailValue string
	var passwordValue string
	if util.IsDevMode() {
		fmt.Println("dev mode!")
		emailValue = "test@gmail.com"
		passwordValue = "aspoaspo"
	}
	return fmt.Sprintf(`
		<form action="/api/user/login" method="POST">
			<h2>Log In</h2>
			<p>%s</p>
			<label for='email'>email</label>
			<input name='email' type='text' value="%s" />
			<label for='password'>password</label>
			<input type='password' name='password' value="%s" />
			<input type='submit' />
		</form>
	`, formErr, emailValue, passwordValue)
}

func RegisterUserForm(formErr string) string {
	return fmt.Sprintf(`
		<form class='flex' action="/api/user/register" method="POST">
			<h2>Register Users</h2>
			<p>%s</p>
			<label for='firstName'>first name</label>
			<input name='firstName' type='text' />
			<label for='lastName'>last name</label>
			<input name='lastName' type='text' />
			<label for='email'>email</label>
			<input name='email' type='text' />
			<label for='locationNumber'>location number</label>
			<input type='text' name='locationNumber' />
			<label for='password'>password</label>
			<input type='password' name='password' />
			<input type='submit' />
		</form>
	`, formErr)
}

func AdminNav() string {
	return `
		<nav>
			<ul>
				<li>
					<a href='/admin'>Home</a>
				</li>
				<li>
					<a href='/admin/users'>Users</a>
				</li>
				<li>
					<a href='/admin/locations'>Locations</a>
				</li>
				<li>
					<a href='/api/user/logout'>Logout</a>
				</li>
			</ul>
		</nav>
	`
}

func TeamNav() string {
	return `
		<nav>
			<ul>
				<li>
					<a href='/app'>Home</a>
				</li>
				<li>
					<a href='/app/bio'>Update Bio</a>
				</li>
				<li>
					<a href='/app/peers'>Peers</a>
				</li>
				<li>
					<a href='/api/user/logout'>Logout</a>
				</li>
			</ul>
		</nav>
	`
}

func UserList(users []usermod.Model) string {
	var userListHTML string
	for _, user := range users {
		userHTML := fmt.Sprintf(`
			<li>
				<a href="/admin/user/%d">%d %s %s %s</a>
			</li>
		`, user.ID, user.ID, user.FirstName, user.LastName, user.Email)
		userListHTML += userHTML
	}

	return fmt.Sprintf(`
		<ul>
			%s
		</ul>
	`, userListHTML)
}

func UserDetails(user *usermod.Model) string {
    return fmt.Sprintf(`
		<div>
			<h1>%s %s</h1>
		</div>
    `, user.FirstName, user.LastName)
}

func UserPhoto(user *usermod.Model) string {
    if len(user.Photo) > 0 {
        return fmt.Sprintf(`<img src="data:image/jpeg;base64,%s" alt="%s %s">`, base64.StdEncoding.EncodeToString(user.Photo), user.FirstName, user.LastName)
    }
    return `<img src="/public/img/person.jpg" alt="Default Image">`
}

func UpdateUserForm(user *usermod.Model, r *http.Request, formErr string) string {
	return fmt.Sprintf(`
	<form action='/api/user/update?previousURL=%s' method='POST' enctype='multipart/form-data'>
		<h1>Update User Details</h1>
		<p>%s</p>
		<input type='hidden' name='id' value='%d' />
		<label for='firstName'>first name</label>
		<input name='firstName' type='text' value='%s' />
		<label for='lastName'>last name</label>
		<input name='lastName' type='text' value='%s' />
		<label for='email'>email</label>
		<input name='email' type='text' value='%s' />
		<label for='locationNumber'>location number</label>
		<input name='locationNumber' type='text' value='%d' />
		<label for='family'>Family</label>
		<textarea name='family'>%s</textarea>
		<label for='hobbies'>Hobbies</label>
		<textarea name='hobbies'>%s</textarea>
		<label for='dreams'>Dreams</label>
		<textarea name='dreams'>%s</textarea>
		<label for='photo'>photo</label>
		<input type='file' name='photo' value='%b' accept='image/*' />
		<input type='submit' />
	</form>`, r.URL.Path, formErr, user.ID, user.FirstName, user.LastName, user.Email, user.LocationNumber, user.Family, user.Hobbies, user.Dreams, user.Photo)
}

func UserBioForm(user *usermod.Model, r *http.Request, formErr string) string {
	return fmt.Sprintf(`
		<form action='/api/user/bio?previousURL=%s' method='POST'>
			<h1>User Bio</h1>
			<p>%s</p>
			<label for='family'>Friends and Family</label>
			<textarea name='family'>%s</textarea>
			<label for='hobbies'>Hobbies and Intrests</label>
			<textarea name='hobbies'>%s</textarea>
			<label for='dreams'>Dreams and Aspirations</label>
			<textarea name='dreams'>%s</textarea>
			<input type='submit' />
		</form>
	`, r.URL.Path, formErr, user.Family, user.Hobbies, user.Dreams)
}

func CreateLocationForm(formErr string) string {
	return fmt.Sprintf(`
		<form action='/api/location/create' method='POST'>
			<h2>Create Location</h2>
			<p>%s</p>
			<label for='name'>Name</label>
			<input name='name' type='text' />
			<label for='number'>Number</label>
			<input name='number' type='number' />
			<input type='submit' />
		</form>
	`, formErr)
}

func LocationList(locations []locationmod.Model) string {
    var locationListHTML string
    for _, location := range locations {
        locationHTML := fmt.Sprintf(`
            <li>
                <a href="/admin/location/%d">%s - %d</a>
            </li>
        `, location.ID, location.Name, location.Number)
        locationListHTML += locationHTML
    }

    return fmt.Sprintf(`
        <ul>
            %s
        </ul>
    `, locationListHTML)
}

func LocationDetails(location *locationmod.Model) string {
    return fmt.Sprintf(`
        <div>
            <h1>%s</h1>
            <p>Location ID: %d</p>
        </div>
    `, location.Name, location.ID)
}

func Todo(items []string) string {
	var todoListHTML string
	for _, item := range items {
		todoItemHTML := fmt.Sprintf("<li>%s</li>", item)
		todoListHTML += todoItemHTML
	}

	return fmt.Sprintf("<ul>%s</ul>", todoListHTML)
}

func PeerList(users []usermod.Model) string {
	var userListHTML string
	for _, user := range users {
		userHTML := fmt.Sprintf(`
			<li>
				<a href="/app/peer/%d">%d %s %s %s</a>
			</li>
		`, user.ID, user.ID, user.FirstName, user.LastName, user.Email)
		userListHTML += userHTML
	}

	return fmt.Sprintf(`
		<ul>
			%s
		</ul>
	`, userListHTML)
}


func UserBio(peer *usermod.Model) string {
	return fmt.Sprintf(`
		<div>
			<p>Family Details:</p>
			<p>%s</p>
			<p>Hobbies:</p>
			<p>%s</p>
			<p>Dreams:</p>
			<p>%s</p>
		</div>
	`, util.InsteadOfEmptyString(peer.Family, "not set"), util.InsteadOfEmptyString(peer.Hobbies, "not set"), util.InsteadOfEmptyString(peer.Dreams, "not set"))
}