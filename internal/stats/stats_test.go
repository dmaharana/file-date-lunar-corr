package stats

import (
	"testing"
)

func TestPerformChiSquareTest(t *testing.T) {
	tests := []struct {
		name     string
		observed []int
		wantLowP bool
	}{
		{
			name:     "Uniform Distribution",
			observed: []int{10, 10, 10, 10, 10, 10, 10, 10},
			wantLowP: false,
		},
		{
			name:     "Non-Uniform Distribution",
			observed: []int{20, 0, 0, 0, 20, 0, 0, 0},
			wantLowP: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PerformChiSquareTest(tt.observed)
			if tt.wantLowP && got.PValue > 0.05 {
				t.Errorf("PerformChiSquareTest() PValue = %v, want <= 0.05", got.PValue)
			}
			if !tt.wantLowP && got.PValue <= 0.05 {
				t.Errorf("PerformChiSquareTest() PValue = %v, want > 0.05", got.PValue)
			}
		})
	}
}
