package day3

import "testing"

func TestDay3_Run_Solution1(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution1 int
	}{
		{
			name: "Example data",
			input: []string{
				"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			wantSolution1: 161,
		},
		{
			name: "Example data split over lines",
			input: []string{
				"xmul(2,4)%&mul[3,7]!@^do_not_mul(5",
				",5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			wantSolution1: 161,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Day{
				data: tt.input,
			}
			gotSolution1, _ := d.Run()
			if gotSolution1 != tt.wantSolution1 {
				t.Errorf("Run() gotSolution1 = %v, want %v", gotSolution1, tt.wantSolution1)
			}
		})
	}
}

func TestDay3_Run_Solution2(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution2 int
	}{
		{
			name: "Example data",
			input: []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			wantSolution2: 48,
		},
		{
			name: "Example data split over lines",
			input: []string{
				"xmul(2,4)&mul[3,7]!^don't()_mul(5,",
				"5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			},
			wantSolution2: 48,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Day{
				data: tt.input,
			}
			_, gotSolution2 := d.Run()
			if gotSolution2 != tt.wantSolution2 {
				t.Errorf("Run() gotSolution2 = %v, want %v", gotSolution2, tt.wantSolution2)
			}
		})
	}
}
