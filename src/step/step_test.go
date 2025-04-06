package step

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStep(t *testing.T) {
	tests := []struct {
		stepString string
		parsed     Step
	}{
		{"Cmaj@u2d4!120,30", Step{NotesString: []string{"Cmaj"}, Arpeggio: []string{"u2d4"}, Velocity: []int{120, 30}}},
		{"c,d,e#0.1,0.3", Step{NotesString: []string{"c", "d", "e"}, Gate: []float64{0.1, 0.3}}},
	}

	for _, test := range tests {
		t.Run(test.stepString, func(t *testing.T) {
			parsed := Step{Original: test.stepString}
			err := parsed.Parse()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assert.Equal(t, test.parsed.NotesString, parsed.NotesString, "NotesString should match")
			assert.Equal(t, test.parsed.Arpeggio, parsed.Arpeggio, "Arpeggio should match")
			assert.Equal(t, test.parsed.Velocity, parsed.Velocity, "Velocity should match")
			assert.Equal(t, test.parsed.Transpose, parsed.Transpose, "Transpose should match")
			assert.Equal(t, test.parsed.Probability, parsed.Probability, "Probability should match")
			assert.Equal(t, test.parsed.Gate, parsed.Gate, "Gate should match")
		})
	}

}
