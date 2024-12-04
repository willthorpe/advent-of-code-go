package day1

import "testing"

func TestDay1_Run(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution1 int
		wantSolution2 int
	}{
		{
			name: "Example data",
			input: []string{
				"3   4",
				"4   3",
				"2   5",
				"1   3",
				"3   9",
				"3   3",
			},
			wantSolution1: 11,
			wantSolution2: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Day1{
				data: tt.input,
			}
			gotSolution1, gotSolution2 := d.Run()
			if gotSolution1 != tt.wantSolution1 {
				t.Errorf("Run() gotSolution1 = %v, want %v", gotSolution1, tt.wantSolution1)
			}
			if gotSolution2 != tt.wantSolution2 {
				t.Errorf("Run() gotSolution2 = %v, want %v", gotSolution2, tt.wantSolution2)
			}
		})
	}
}
