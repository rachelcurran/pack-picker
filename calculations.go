package main

import (
	"github.com/ernestosuarez/itertools"

	"sort"
)

// Not used in current flow
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

func calculatedAllPacksForCombinations(packSizes []int, numberOfItems int) []int{
	allPackSizes := []int{}

	// For the given pack sizes get all required to reach the target, including multiples
	for i, pack := range packSizes {
		hasMultiple := false
		duplicatesCount := 0
		for j := (i + 1); j < (len(packSizes)); j++ {
			nextValue := packSizes[j]
			usingOrderValue := false

			if numberOfItems < packSizes[j] {
				usingOrderValue = true
				nextValue = numberOfItems
			}

			// If pack sizes are a factor of another then only get as many as needed to reach the next
			if nextValue % pack == 0 {
				hasMultiple = true
				if usingOrderValue {
					duplicatesCount = (nextValue / pack)
				} else {
					duplicatesCount = (nextValue / pack) - 1 // -1 because we need one less than the muliple
				}

				break
			}
		}

		if !hasMultiple {
			duplicatesCount = (numberOfItems / pack) + 1
		}

		for j := 0; j < duplicatesCount; j++{
			allPackSizes = append(allPackSizes, pack)
		}
	}

	return allPackSizes
}

func calculatePackCombinations(setPackSizes, packSizes []int, numberOfItems int) []PackCombination{
	count := len(packSizes) - 1 
	packCombinations := []PackCombination{}

	// Get all possible values that we can make given the set packs
	for count > 0 {
		for v := range itertools.CombinationsInt(packSizes, count) {
			sum := sum(v)
			if sum < numberOfItems + setPackSizes [len(setPackSizes) - 1]{
				// If this total was already found, keep the lowest number of packs
				if contains(packCombinations, sum){
					if isNewDuplicatePreferred(packCombinations,  sum, len(v)){
						packCombinations = removeElement(packCombinations, sum)
						packCombinations = append(packCombinations, PackCombination{
							Sum: sum,
							Packs: v,
						})
					}
				}else{
					packCombinations = append(packCombinations, PackCombination{
						Sum: sum,
						Packs: v,
					})
				}
			}
			}
		count--;
	} 

	return packCombinations
}

func calculateRequiredPacks(packCombinations []PackCombination, numberOfItems int) []int {
	for _, combination := range packCombinations {
		if numberOfItems <= combination.Sum {
			sort.Sort(sort.Reverse(sort.IntSlice(combination.Packs)))
			return combination.Packs
		}
	}

	return []int { 0 } // something has gone wrong somewhere
}
