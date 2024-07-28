package main

import (
	"github.com/stretchr/testify/assert"

    "testing"
)

// ***** calculateAllPacksForCombinations *****

func TestCalculateAllPacksForCombinations_GivenPackSizesThatAreFactors_ReturnsAllPacks(t *testing.T) {
	// arrange 
	packSizes := []int{ 10, 50, 100 }
	orderSize := 140

	// act 
	result := calculateAllPacksForCombinations(packSizes, orderSize)

	// assert
    expectedResult := []int{ 10, 10, 10, 10, 50, 100, 100 }

	assert.Equal(t, result, expectedResult)
}

func TestCalculateAllPacksForCombinations_GivenPackSizesThatAreNotFactors_ReturnsAllPacks(t *testing.T) {
	// arrange 
	packSizes := []int{ 20, 30 }
	orderSize := 110

	// act 
	result := calculateAllPacksForCombinations(packSizes, orderSize)

	// assert
    expectedResult := []int{ 20, 20, 20, 20, 20, 20, 30, 30, 30, 30 }

	assert.Equal(t, result, expectedResult)
}

// ***** calculatePackCombinations *****

func TestCalculatePackCombinations_GivenListOfValues_ReturnsAllPossibleSums(t *testing.T) {
	// arrange 
	packSizes := []int{ 10, 50 }
	packs := []int{ 10, 10, 10, 10, 50 }
	orderSize := 60

	// act 
	result := calculatePackCombinations(packSizes, packs, orderSize)

	// assert
    expectedResult := []PackCombination{
		{ Sum : 10, Packs: []int{10} },
		{ Sum : 20, Packs: []int{10, 10} },
		{ Sum : 30, Packs: []int{10,10,10} },
		{ Sum : 40, Packs: []int{10,10,10,10} },
		{ Sum : 50, Packs: []int{50} },
		{ Sum : 60, Packs: []int{50,10} },
		{ Sum : 70, Packs: []int{50,10,10} },
		{ Sum : 80, Packs: []int{50,10,10,10} },
	}

	assert.Equal(t, len(result), len(expectedResult))
}


// ***** calculateRequiredPacks *****

func TestCalculateRequiredPacks_GivenListOfValuesAndTarget_ReturnsRequiredPacks(t *testing.T) {
	// arrange 

	packCombinations := []PackCombination{
		{ Sum : 50, Packs: []int{50} },
		{ Sum : 100, Packs: []int{50, 50} },
		{ Sum : 150, Packs: []int{150} },
	}

	orderSize := 90

	// act 
	result := calculateRequiredPacks(packCombinations, orderSize)

	// assert
    expectedResult := []int{50,50}

	assert.Equal(t, result, expectedResult)
}