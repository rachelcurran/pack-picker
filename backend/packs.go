package main 

import(
	"github.com/ernestosuarez/itertools"
	"sort"
)

type PackCombination struct {
    Sum int 
    Packs []int
}

var _packSizes = []int {}
// do we want fallbackPackSizes incase nothing is added first 
// can add message in the front to explain this

func setPackSizes(packSizes []int) ([]int){
	_packSizes = packSizes

	return _packSizes
}

func getPackSizes()([]int){
	return _packSizes
}

func getPacksOld(numberOfItems int) ([]int){
	return getPacksFromSizes(numberOfItems, _packSizes)
}

func getPacksFromSizes(numberOfItems int, packSizes []int) ([]int){

	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	packs := calculatePacks(packSizes, numberOfItems) 
	packsSum := sum(packs)

	// Check if the packs would return exess items
    // If so, check for optimum solution with fewer packs
    if packsSum != numberOfItems {
		packs = calculatePacks(packSizes, packsSum)
	}

	return packs
}

// The initial approach was calculating a list of packs that always resulted in the omptimum number of items, 
// however in certain cases with excess items, this was not always the optimum number of packs. 
// In this scenario, re-running the logic to calculate the packs but with a new target results in corret number of 
// packs and items.

// It was found after testing that the above method only works when pack sizes fit within each other 
// and was not always accurate when more obscure pack sizes are used

// The below method is more acurate and will return the smallest number of items in the smallest
// number of packs.
// However when using obsure pack sizes and number of items in order, this can can take a long time to process
// so please use carefully 
// This will occur with either a very large number of items in order or small pack sizes, when pack sizes don't fit in others


func getPacks(numberOfItems int, packSizes []int) ([]int){

	sort.Sort(sort.IntSlice(packSizes))
	allPackSizes := []int{}

	// For the given pack sizes get all required to reach the target, including multiples
	for i, pack := range packSizes {
		hasMultiple := false
		duplicatesCount := 0
		for j := (i + 1); j < (len(packSizes)); j++ {
			nextValue := packSizes[j]
			if numberOfItems < packSizes[j] {
				nextValue = numberOfItems
			}

			// If pack sizes are a factor of another then only get as many as needed to reach the next
			if nextValue % pack == 0 {
				hasMultiple = true
				duplicatesCount = (nextValue / pack) - 1 // -1 because we need one less than the muliple
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

	packCombinations := []PackCombination{}
	count := len(allPackSizes) - 1 

	// Get all possible values that we can make given the set packs
	for count > 0 {
		for v := range itertools.CombinationsInt(allPackSizes, count) {
			sum := sum(v)
			if sum < numberOfItems + packSizes[len(packSizes) - 1]{
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

	sort.Slice(packCombinations, func(i, j int) bool{
		return packCombinations[i].Sum < packCombinations[j].Sum
	})

	return calculateRequiredPacks(packCombinations, numberOfItems) 
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
