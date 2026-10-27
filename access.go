package main

import "fmt"

// Задача 2. Контроль доступа к GoFlix
func verifyAccount(userAge int, subscriptionType string, isAccountBanned bool, ageRate int) bool {
	if isAccountBanned {
		return false
	}
	if !isAllowed(userAge, ageRate) {
		return false
	}

	switch subscriptionType {
	case "free":
		fmt.Println("Доступно с рекламой")
	case "premium":
		fmt.Println("Приятного просмотра!")
	default:
		return false
	}
	return true
}
