package day5

import "testing"

func TestDay5_Run(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		wantSolution1 int
		wantSolution2 int
	}{
		{
			name: "Example data",
			input: []string{
				"47|53",
				"97|13",
				"97|61",
				"97|47",
				"75|29",
				"61|13",
				"75|53",
				"29|13",
				"97|29",
				"53|29",
				"61|53",
				"97|53",
				"61|29",
				"47|13",
				"75|47",
				"97|75",
				"47|61",
				"75|61",
				"47|29",
				"75|13",
				"53|13",

				"75,47,61,53,29",
				"97,61,53,29,13",
				"75,29,13",
				"75,97,47,61,53",
				"61,13,29",
				"97,13,75,29,47",
			},
			wantSolution1: 143,
			wantSolution2: 123,
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
