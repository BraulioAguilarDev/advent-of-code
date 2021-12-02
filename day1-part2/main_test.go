package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasurement_part2(t *testing.T) {

	expected := 5

	rows := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	result := summedMeasurements(rows)
	assert.Equal(t, expected, result)
}
