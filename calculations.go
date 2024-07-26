package main

func calculatePacks(packSizes []int, numberOfItems int) ([]int){

	packs := []int{}
	itemsRemaining := numberOfItems

	for itemsRemaining > 0 {
		
		// smaller than the smallest
		smallestPackSize := packSizes[len(packSizes)-1]
		if itemsRemaining <= smallestPackSize {
			packs = append(packs, smallestPackSize)
			itemsRemaining -= smallestPackSize
			break
		}

		for i := 0; i < len(packSizes); i++{
			if itemsRemaining >= packSizes[i]{
				packs = append(packs, packSizes[i])
				itemsRemaining -= packSizes[i]
				break
			}
		}
	}
	
	return packs
}

func sum(values []int) (int) {
	
	total := 0

	for _, v := range values {
		total += v
	} 

	return total
}
