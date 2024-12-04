package day2

import "testing"

func TestDay2_Run(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution1 int
		wantSolution2 int
	}{
		{
			name: "Example data",
			input: []string{
				"7 6 4 2 1",
				"1 2 7 8 9",
				"9 7 6 2 1",
				"1 3 2 4 5",
				"8 6 4 4 1",
				"1 3 6 7 9",
			},
			wantSolution1: 2,
			wantSolution2: 4,
		},
		{
			name:          "Ascending data",
			input:         []string{"1 2 3 4 5"},
			wantSolution1: 1,
			wantSolution2: 1,
		},
		{
			name:          "Descending data",
			input:         []string{"5 4 3 2 1"},
			wantSolution1: 1,
			wantSolution2: 1,
		},
		{
			name:          "Data increasing by 2",
			input:         []string{"1 3 5 7 9"},
			wantSolution1: 1,
			wantSolution2: 1,
		},
		{
			name:          "Data increasing by 3",
			input:         []string{"1 4 7 10 13"},
			wantSolution1: 1,
			wantSolution2: 1,
		},
		{
			name:          "Data increasing by 4",
			input:         []string{"1 5 9 13 17"},
			wantSolution1: 0,
			wantSolution2: 0,
		},
		{
			name:          "Ascending data with descending at the end",
			input:         []string{"3 2 3 4 5"},
			wantSolution1: 0,
			wantSolution2: 1,
		},
		{
			name:          "Descending data with ascending at the start",
			input:         []string{"3 4 3 2 1"},
			wantSolution1: 0,
			wantSolution2: 1,
		},
		{
			name:          "Duplicate data at the start",
			input:         []string{"1 1 2 3 4"},
			wantSolution1: 0,
			wantSolution2: 1,
		},
		{
			name:          "Duplicate data at the end",
			input:         []string{"1 2 3 4 4"},
			wantSolution1: 0,
			wantSolution2: 1,
		},
		{
			name:          "Duplicate data within the data",
			input:         []string{"1 2 2 3 4"},
			wantSolution1: 0,
			wantSolution2: 1,
		},
		{
			name:          "Triple data at the start of the data",
			input:         []string{"2 2 2 3 4"},
			wantSolution1: 0,
			wantSolution2: 0,
		},
		{
			name:          "Triple data at the end of the data",
			input:         []string{"2 3 4 4 4"},
			wantSolution1: 0,
			wantSolution2: 0,
		},
		{
			name:          "Triple data within the data",
			input:         []string{"1 2 2 2 3 4"},
			wantSolution1: 0,
			wantSolution2: 0,
		},
		{
			name:          "All same value",
			input:         []string{"2 2 2 2"},
			wantSolution1: 0,
			wantSolution2: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Day2{
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
