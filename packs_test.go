package main

import (
	"github.com/stretchr/testify/assert"

    "testing"
)

// ***** getPacks *****
// ***** suggested pack sizes *****
// 250, 500, 1000, 2000, 5000

func TestGetPacks_GivenSuggestedPackSizes_ReturnsCorrect_1(t *testing.T) {
	// arrange 
	packSizes := []int{ 250, 500, 1000, 2000, 5000 }
	orderSize := 1

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 250 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenSuggestedPackSizes_ReturnsCorrect_250(t *testing.T) {
	// arrange 
	packSizes := []int{ 250, 500, 1000, 2000, 5000 }
	orderSize := 250

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 250 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenSuggestedPackSizes_ReturnsCorrect_251(t *testing.T) {
	// arrange 
	packSizes := []int{ 250, 500, 1000, 2000, 5000 }
	orderSize := 251

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 500 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenSuggestedPackSizes_ReturnsCorrect_501(t *testing.T) {
	// arrange 
	packSizes := []int{ 250, 500, 1000, 2000, 5000 }
	orderSize := 501

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 500, 250 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenSuggestedPackSizes_ReturnsCorrect_12001(t *testing.T) {
	// arrange 
	packSizes := []int{ 250, 500, 1000, 2000, 5000 }
	orderSize := 12001

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 5000, 5000, 2000, 250 }

	assert.Equal(t, result, expectedResult)
}

// ***** getPacks *****
// ***** other pack sizes *****
// 300, 500, 1000

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_50(t *testing.T) {
	// arrange 
	packSizes := []int{ 300, 500, 1000 }
	orderSize := 50

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 300 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_400(t *testing.T) {
	// arrange 
	packSizes := []int{ 300, 500, 1000 }
	orderSize := 400

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 500 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_501(t *testing.T) {
	// arrange 
	packSizes := []int{ 300, 500, 1000 }
	orderSize := 501

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 300, 300 }

	assert.Equal(t, result, expectedResult)
}


// *** 34, 57, 130 ***

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_12(t *testing.T) {
	// arrange 
	packSizes := []int{ 34, 57, 130 }
	orderSize := 12

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 34 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_51(t *testing.T) {
	// arrange 
	packSizes := []int{ 44, 67, 130 }
	orderSize := 51

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 67 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_100(t *testing.T) {
	// arrange 
	packSizes := []int{ 44, 67, 130 }
	orderSize := 100

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 67, 44 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_250(t *testing.T) {
	// arrange 
	packSizes := []int{ 44, 67, 130 }
	orderSize := 250

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 130, 130 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenOtherPackSizes_ReturnsCorrect_270(t *testing.T) {
	// arrange 
	packSizes := []int{ 44, 67, 130 }
	orderSize := 270

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 130, 67, 44, 44 }

	assert.Equal(t, result, expectedResult)
}


// ***** getPacks *****
// ***** large pack sizes *****
// 250000, 500000, 1000000, 2000000, 5000000

func TestGetPacks_GivenLargePackSizes_ReturnsCorrect_400(t *testing.T) {
	// arrange 
	packSizes := []int{ 250000, 500000, 1000000, 2000000, 5000000 }
	orderSize := 400

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 250000 }

	assert.Equal(t, result, expectedResult)
}

func TestGetPacks_GivenLargePackSizes_ReturnsCorrect_12000001(t *testing.T) {
	// arrange 
	packSizes := []int{ 250000, 500000, 1000000, 2000000, 5000000 }
	orderSize := 12000001

	// act 
	result := getPacks(orderSize, packSizes)

	// assert
    expectedResult := []int{ 5000000, 5000000, 2000000, 250000 }

	assert.Equal(t, result, expectedResult)
}