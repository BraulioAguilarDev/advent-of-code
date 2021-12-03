package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAimOk(t *testing.T) {

	rows := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	result := aim(rows)
	assert.Equal(t, 900, result)
}
