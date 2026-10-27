package main

import "fmt"

// Задача 1. Расчёт стоимости билетов.

// ticketPrice добавляет к базовой цене надбавки за вечер и выходной день
func ticketPrice(basePrice int, evening bool, weekend bool) int {
	if evening {
		basePrice += 50
	}
	if weekend {
		basePrice += 100
	}
	return basePrice
}

// formatReceipt возвращает текст чека и общую стоимость билетов
func formatReceipt(filmName string, ticketsCount int, ticketPrice int) (string, int) {
	finalPrice := ticketPrice * ticketsCount
	receipt := fmt.Sprintf(
		"Фильм: %s, билетов: %d, итого: %d руб.",
		filmName, ticketsCount, finalPrice,
	)
	return receipt, finalPrice
}

// isAllowed проверяет, соответствует ли возраст зрителя рейтингу фильма
func isAllowed(age int, ageRate int) bool {
	return age >= ageRate
}
