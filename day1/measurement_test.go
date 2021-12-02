package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasurement(t *testing.T) {

	expected := 7

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

	result := measurement(rows)
	assert.Equal(t, expected, result)
}
