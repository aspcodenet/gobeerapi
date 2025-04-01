package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Det finns X antal case vi skulle testa om vi gjorde manuellt
// 1. KROGEN, 18, 0.5 -> true
// 2. KROGEN, 18, 1.5 -> false
// 3. KROGEN, 17, 0.5 -> false
// 4. SYSTEMET, 20, 0.5 -> true
// 5. SYSTEMET, 20, 1.5 -> false
// 6. SYSTEMET, 19, 0.5 -> false

// Arrange, - arrangera världen så att den ser ut som vi vill ha den
//
//	logga in som en anv som handlat för minst 30000
//
// Act, 	- agera (buy)
// Assert   - kontrollera att det blev som vi ville ha det
//   - säkerställ att den blev 5% rabatt
func TestWhen_18_And_Not_Too_Drunk_And_On_Krogen_Then_Should_BeAllowed(t *testing.T) {
	// ARRANGE
	location := "K"
	age := 18
	promille := float32(0.5)

	// ACT
	canBuy, err := CanBuyBeer(location, age, promille)

	// ASSERT
	assert.True(t, canBuy)
	assert.Nil(t, err)
}

func TestWhen_17_And_Not_Too_Drunk_And_On_Krogen_Then_Should_NotBeAllowed(t *testing.T) {
	// ARRANGE
	location := "K"
	age := 17
	promille := float32(0.0)

	// ACT
	canBuy, err := CanBuyBeer(location, age, promille)

	// ASSERT
	assert.False(t, canBuy)
	assert.Nil(t, err)
}

func TestWhen_20_And_Not_Too_Drunk_And_On_Systemet_Then_Should_NotBeAllowed(t *testing.T) {
	// ARRANGE
	location := "S"
	age := 20
	promille := float32(0.0)

	// ACT
	canBuy, err := CanBuyBeer(location, age, promille)

	// ASSERT
	assert.True(t, canBuy)
	assert.Nil(t, err)
}
