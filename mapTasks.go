package main

// посчитать количество символов в строке
func countChars(s string) map[rune]int {
	result := make(map[rune]int)
	for _, runes := range s {
		//можно еще отдельно
		//_, ok := result[runes];
		if _, ok := result[runes]; !ok {
			result[runes] = 1
		} else {
			result[runes]++
		}
	}
	return result
}

// есть ли дубликаты в списке?
func hasDuplicates(nums []int) bool {
	numsMap := make(map[int]int)
	for _, num := range nums {
		_, ok := numsMap[num]
		if !ok {
			numsMap[num] = 1
		} else {
			return true
		}
	}
	return false
}

// добавить фильм каталог, если его еще там нет
func addFilm(catalog map[string]float64, filmName string, filmRating float64) {
	if _, ok := catalog[filmName]; !ok {
		catalog[filmName] = filmRating
	}
}

// получить топовые фильмы
func topFilms(catalog map[string]float64, min float64) []string {
	result := make([]string, 0)
	for filmName, filmRating := range catalog {
		if filmRating >= min {
			result = append(result, filmName)
		}
	}
	return result
}

// получить средний рейтинг каталога
func avgRating(catalog map[string]float64) float64 {
	var summ float64
	for _, filmRating := range catalog {
		summ += filmRating
	}
	return summ / float64(len(catalog))
}
