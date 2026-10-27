package main

import "fmt"

// Задача 3. Генератор расписания кинотеатра GoFlix
func printSchedule() {
	films := []string{"Дюна", "Матрица", "Оппенгеймер", "Начало"}
	timeSlots := []string{"10:00", "13:00", "16:00", "19:00"}
	const hallsCount = 3

	for hall := 1; hall <= hallsCount; hall++ {
		fmt.Printf("Зал %d:\n", hall)
		for index, filmName := range films {
			fmt.Printf("%s — %s\n", timeSlots[index], filmName)
		}
		fmt.Println()
	}

search:
	for hall := 1; hall <= hallsCount; hall++ {
		for index, filmName := range films {
			if filmName == "Начало" {
				fmt.Printf("Найден фильм «%s»: зал %d, время %s\n", filmName, hall, timeSlots[index])
				break search
			}
		}
	}

	for _, filmName := range films {
		runesCount := 0
		for range filmName {
			runesCount++
		}
		fmt.Printf("«%s»: %d рун\n", filmName, runesCount)
	}
}
