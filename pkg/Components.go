package pkg

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func LoginForm(formErr string) string {
	return fmt.Sprintf(`
		<form action="/api/user/login" method="POST">
			<h2>Log In</h2>
			<p>%s</p>
			<label for='email'>email</label>
			<input name='email' type='text' />
			<label for='password'>password</label>
			<input type='password' name='password' />
			<input type='submit' />
		</form>
	`, formErr)
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
					<a href='/api/user/logout'>Logout</a>
				</li>
			</ul>
		</nav>
	`
}

func UserList(users []User) string {
	var userListHTML string
	for _, user := range users {
		userHTML := fmt.Sprintf(`
			<li>
				<a href="/admin/user/%d">%d %s %s %s</a>
			</li>
		`, user.id, user.id, user.firstName, user.lastName, user.email)
		userListHTML += userHTML
	}

	return fmt.Sprintf(`
		<ul>
			%s
		</ul>
	`, userListHTML)
}

func UserDetails(user *User) string {
    return fmt.Sprintf(`
		<div>
			<h1>%s %s</h1>
		</div>
    `, user.firstName, user.lastName)
}

func UserPhoto(user *User) string {
    if len(user.photo) > 0 {
        return fmt.Sprintf(`<img src="data:image/jpeg;base64,%s" alt="%s %s">`, base64.StdEncoding.EncodeToString(user.photo), user.firstName, user.lastName)
    }
    return `<img src="/public/img/person.jpg" alt="Default Image">`
}

func UpdateUserForm(user *User, r *http.Request, formErr string) string {
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
		<label for='photo'>photo</label>
		<input type='file' name='photo' value='%b' accept='image/*' />
		<input type='submit' />
	</form>`, r.URL.Path, formErr, user.id, user.firstName, user.lastName, user.email, user.photo)
}
