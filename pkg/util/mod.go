package util

import (
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func AppendQueryParam(originalURL, paramName, paramValue string) string {
    queryParams := url.Values{}
    queryParams.Set(paramName, paramValue)

    // Check if the original URL already has query parameters
    if strings.Contains(originalURL, "?") {
        return originalURL + "&" + queryParams.Encode()
    }

    return originalURL + "?" + queryParams.Encode()
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func IsDevMode() bool {
	return os.Getenv("ENV") == "dev"
}

func IfDevModeThen(value string) string {
	if os.Getenv("ENV") == "dev" {
		return value
	}
	return ""
}

func StringToInt(s string) (int, error) {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return i, nil
}

func InsteadOfEmptyString(potentialEmptyString string, replacementString string) string {
	if potentialEmptyString == "" {
		return replacementString
	}
	return potentialEmptyString
}

func RenderTemplate(w http.ResponseWriter, tmplFile string, data interface{}) {
	// Parse the template from the file
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing the data to it
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}