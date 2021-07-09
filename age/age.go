package age

func FilterAges(firstAge int, secondAge int, ages []int) []int {
	filteredAges := []int{}

	for i := 0; i < len(ages); i++ {
		age := ages[i]
		if age >= firstAge && age <= secondAge {
			filteredAges = append(filteredAges, age)
		}
	}

	return filteredAges
}
