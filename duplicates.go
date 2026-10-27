package main

import "slices"

// Задача 4. Удаление дубликатов из слайса без использования мапы
func removeDuplicates(values []string) []string {
	var unique []string
	for _, value := range values {
		if !slices.Contains(unique, value) {
			unique = append(unique, value)
		}
	}
	return unique
}
