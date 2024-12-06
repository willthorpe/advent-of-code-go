package day6

import "testing"

func TestDay6_Run(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution1 int
		wantSolution2 int
	}{
		{
			name: "Example data",
			input: []string{
				"....#.....",
				".........#",
				"..........",
				"..#.......",
				".......#..",
				"..........",
				".#..^.....",
				"........#.",
				"#.........",
				"......#...",
			},
			wantSolution1: 41,
			wantSolution2: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Day{
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
