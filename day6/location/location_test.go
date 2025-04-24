package location_test

import (
	"testing"

	"github.com/brainboxweb/advent-2024/day6/location"
	"github.com/brainboxweb/advent-2024/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../../testdata/"

func TestMove(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"day6_test.txt",
			41,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			theMap := location.NewMap(dataSet)
			result := theMap.Move()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsEndlessLoop(t *testing.T) {
	tests := []struct {
		dataFile string
		expected bool
	}{
		{
			"day6_test.txt",
			false,
		},
		{
			"day6_test2.txt",
			true,
		},
		{
			"day6_test3.txt",
			true,
		},
		{
			"day6_test4.txt",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.dataFile, func(t *testing.T) {
			data := helpers.GetDataString(testDataPath + tt.dataFile)
			dataSet := helpers.ToXY(data)
			theMap := location.NewMap(dataSet)
			result := theMap.IsEndlessLoop()
			assert.Equal(t, tt.expected, result)
		})
	}
}
