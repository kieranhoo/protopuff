package utils

import (
	"protopuff/internal/config"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenPassword(pass string) ([]byte, error) {
	cost, _ := strconv.Atoi(config.JwtCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return hash, err
}

func ComparePassword(hashPass string, pass string) error {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return errCompare
}
