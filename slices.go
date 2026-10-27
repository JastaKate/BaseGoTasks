package main

//Суммирование и фильтрация слайсов

// sum возвращает сумму элементов
func sum(nums []int) int {
	numsSum := 0
	for _, num := range nums {
		numsSum += num
	}
	return numsSum
}

// filterByRating оставляет рейтинги, не меньшие min
func filterByRating(ratings []float64, min float64) []float64 {
	filteredRatings := make([]float64, 0, len(ratings))
	for _, rating := range ratings {
		if rating >= min {
			filteredRatings = append(filteredRatings, rating)
		}
	}
	return filteredRatings
}
