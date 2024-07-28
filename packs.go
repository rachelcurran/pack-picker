package main 

import(
	"sort"
)

type PackCombination struct {
    Sum int 
    Packs []int
}

var _packSizes = []int {}

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

	// Check if the packs would return excess items
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
// This will occur with either a very large number of items in order or small pack sizes


func getPacks(numberOfItems int, packSizes []int) ([]int){

	sort.Sort(sort.IntSlice(packSizes))

	// Get all packs that could be needed to reach every combination
	allPackSizes := calculateAllPacksForCombinations(packSizes, numberOfItems)

	// Get all pack combinations that can be reached with given pack sizes
	packCombinations := calculatePackCombinations(packSizes, allPackSizes, numberOfItems)

	sort.Slice(packCombinations, func(i, j int) bool{
		return packCombinations[i].Sum < packCombinations[j].Sum
	})

	// Calculate the appropriate pack combination for number of items
	return calculateRequiredPacks(packCombinations, numberOfItems) 
}

