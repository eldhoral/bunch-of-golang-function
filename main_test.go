package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectedVolume float64 = 250
var expectedLuas float64 = 250

func TesHasilVolume(t *testing.T) {
	hasilVolume := hitungVolume(10, 5, 5)
	assert.Equal(t, hasilVolume, expectedVolume, "PASS!")
	fmt.Println("Sudah Jalan")
}

func TesHasilLuas(t *testing.T) {
	hasilLuas := hitungLuas(10, 5, 5)
	assert.Equal(t, hasilLuas, expectedLuas, "PASS!")
	fmt.Println("Sudah Jalan")
}
