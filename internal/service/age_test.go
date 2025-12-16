package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name string
		dob  time.Time
		want int
	}{
		{
			name: "Birthday passed this year",
			dob:  now.AddDate(-20, -1, 0), // 20 years and 1 month ago
			want: 20,
		},
		{
			name: "Birthday is today",
			dob:  now.AddDate(-20, 0, 0), // Exactly 20 years ago
			want: 20,
		},
		{
			name: "Birthday not yet passed this year",
			dob:  now.AddDate(-20, 1, 0), // 19 years and 11 months ago (approx)
			want: 19,
		},
		{
			name: "Born today",
			dob:  now,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateAge(tt.dob); got != tt.want {
				t.Errorf("CalculateAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
