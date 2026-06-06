package lunar

import (
	"testing"
	"time"
)

func TestGetPhase(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want Phase
	}{
		{
			name: "New Moon",
			date: time.Date(2024, 1, 11, 11, 57, 0, 0, time.UTC), // New Moon
			want: New,
		},
		{
			name: "Full Moon",
			date: time.Date(2024, 1, 25, 17, 54, 0, 0, time.UTC), // Full Moon
			want: Full,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPhase(tt.date); got != tt.want {
				t.Errorf("GetPhase() = %v, want %v", got, tt.want)
			}
		})
	}
}
