package main

import (
	"errors"
	"fmt"
	"time"
)

type ValidationError struct {
	Field string
	Value string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid date format: %s, field: %s", e.Value, e.Field)
}

var ErrUserNotFound = errors.New("user not found")

func validateUserData(userId int, userDate string) (string, error) {
	usersMap := map[int]string{
		1: "Kate",
		2: "John",
		3: "Alex",
	}

	user, ok := usersMap[userId]
	if !ok {
		return "", ErrUserNotFound
	}

	_, err := time.Parse("2006-01-02", userDate)
	if err != nil {
		return "", &ValidationError{
			Field: "userDate",
			Value: userDate,
		}
	}

	return user, err

}
