package main

import (
	"github.com/stretchr/testify/assert"

    "testing"
)

// ***** sum *****

func TestSum_GivenNumbers_ReturnsSum(t *testing.T) {
	// arrange 
	numbersToSum := []int{5, 4, 6}

	// act 
	result := sum(numbersToSum)

	// assert
    expectedResult := 15

	assert.Equal(t, result, expectedResult)
}


// ***** contains *****

func TestContains_GivenNumbersWithContainingValue_ReturnsTrue(t *testing.T) {
	// arrange 
	packCombinations := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	// act 
	result := contains(packCombinations, 15)

	// assert
    expectedResult := true

	assert.Equal(t, result, expectedResult)
}

func TestContains_GivenNumbersWithoutContainingValue_ReturnsFalse(t *testing.T) {
	// arrange 
	packCombinations := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	// act 
	result := contains(packCombinations, 50)

	// assert
    expectedResult := false

	assert.Equal(t, result, expectedResult)
}


// ***** isNewDuplicatePreferred *****

func TestIsNewDuplicatePreferred_GivenCombinationWithSumExistsInLessPacks_ReturnsTrue(t *testing.T) {
	// arrange 
	packCombinations := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	sum := 15
	packCount := 2

	// act 
	result := isNewDuplicatePreferred(packCombinations, sum, packCount)

	// assert
    expectedResult := true

	assert.Equal(t, result, expectedResult)
}

func TestIsNewDuplicatePreferred_GivenCombinationWithSumExistsInMorePacks_ReturnsTrue(t *testing.T) {
	// arrange 
	packCombinations := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	sum := 15
	packCount := 5

	// act 
	result := isNewDuplicatePreferred(packCombinations, sum, packCount)

	// assert
    expectedResult := false

	assert.Equal(t, result, expectedResult)
}


// ***** removeElement *****

func TestRemoveElement_GivenArrayAndValueToRemove_ReturnsArrayWithElement(t *testing.T) {
	// arrange 
	sumToRemove := 34

	packCombinations := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : sumToRemove, Packs: []int{2,4,6} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	// act 
	result := removeElement(packCombinations, sumToRemove)

	// assert
    expectedResult := []PackCombination{ 
		{ Sum : 15, Packs: []int{1,2,3} },
		{ Sum : 20, Packs: []int{5,6,7} },
	}

	assert.Equal(t, result, expectedResult)
}

