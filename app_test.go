package main_test

import (
	main "bounded-presentation-clock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Kind_functions(t *testing.T) {
	tcs := []struct {
		name     string
		input    int64
		expected main.SegmentedDuration
	}{
		{
			name:  "sec",
			input: 3670,
			expected: main.SegmentedDuration{
				Hour: 1,
				Min:  1,
				Sec:  10,
			},
		},
		{
			name:  "1 min 10 sec",
			input: 70,
			expected: main.SegmentedDuration{
				Min: 1,
				Sec: 10,
			},
		},
		{
			name:  "20 sec",
			input: 20,
			expected: main.SegmentedDuration{
				Sec: 20,
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := main.NewSegmentedDuration(tc.input)

			assert.Equal(t, tc.expected.Sec, actual.Sec, "sec")
			assert.Equal(t, tc.expected.Min, actual.Min, "min")
			assert.Equal(t, tc.expected.Hour, actual.Hour, "hour")
		})
	}
}
