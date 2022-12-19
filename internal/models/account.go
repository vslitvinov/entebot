package models

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	chars   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

type Accounty struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ISArchive  bool      `json:"archive"`
	ISVerified bool      `json:"verified"`
}

func (a *Accounty) GeneratePasswordHash() error {
	b, err := bcrypt.GenerateFromPassword([]byte(a.Password), 14)
	if err != nil {
		return fmt.Errorf("Accounty.GenerateFromPassword: %w", err)
	}

	a.Password = string(b)

	return nil
}

func (a *Accounty) CompareHashAndPassword(pd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pd)); err != nil {
		return fmt.Errorf("Accounty.CompareHashAndPassword: %w", err)
	}

	return nil
}

func (a *Accounty) RandomPassword() {
	bytes := make([]byte, 16)

	c := chars + special
	for i := range bytes {
		bytes[i] = c[rand.Intn(len(chars))]
	}

	a.Password = string(bytes)
}
