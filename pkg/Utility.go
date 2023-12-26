package pkg

import (
	"math/rand"
	"net/url"
	"strings"
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