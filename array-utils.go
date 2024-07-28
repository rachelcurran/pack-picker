package main

func sum (values []int) (int) {
	
	total := 0

	for _, v := range values {
		total += v
	} 

	return total
}

func contains (values []PackCombination, x int) bool {
	for _, value := range values {
		if value.Sum == x{
			return true
		}
	}
	return false
}

func isNewDuplicatePreferred (values []PackCombination, sum int, packCount int) bool {
	for _, value := range values {
		if value.Sum == sum {
			if len(value.Packs) > packCount {
				return true
			}
			return false
		}
	}
	return false
}

func removeElement (values []PackCombination, sum int) ([]PackCombination) { 

	newValues := []PackCombination{}
	for i, value := range values {
		if value.Sum == sum {
			newValues = append(values[:i], values[i + 1:]...)
			break
		}
	}

	return newValues
}