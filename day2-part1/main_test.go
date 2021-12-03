package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiveOk(t *testing.T) {

	expected := 150

	rows := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	h, v := Dive(rows)

	assert.Equal(t, expected, h*v)
}
