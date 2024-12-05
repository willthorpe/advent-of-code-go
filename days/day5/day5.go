package day5

import (
	"advent-of-code-2024/days/common"
	"advent-of-code-2024/input"
	"slices"
	"strings"
)

type Day struct {
	data      []string
	solution1 int
	solution2 int
}

func NewDay() *Day {
	i := input.NewInput("5", "2024")

	data := i.GetData()

	return &Day{
		data:      data,
		solution1: 0,
		solution2: 0,
	}
}

func (d *Day) Run() (int, int) {
	var rules [][]string
	var pages []string

	for _, line := range d.data {
		if strings.Contains(line, "|") {
			rules = append(rules, strings.Split(line, "|"))
		}

		if strings.Contains(line, ",") {
			pages = append(pages, line)
		}
	}

	for _, page := range pages {
		pass := true
		var rulesThatApply [][]string
		for _, rule := range rules {
			if !strings.Contains(page, rule[0]) || !strings.Contains(page, rule[1]) {
				continue
			}

			rulesThatApply = append(rulesThatApply, rule)

			if strings.Index(page, rule[1]) < strings.Index(page, rule[0]) {
				pass = false
			}
		}

		if pass {
			sp := strings.Split(page, ",")
			d.solution1 += common.ConvertStringToInt(sp[len(sp)/2])
		}

		if !pass {
			sp := strings.Split(page, ",")
			var pageLeftInRuleCount []int

			for _, p := range sp {
				countInFirstPosition := 0
				for _, rule := range rulesThatApply {
					if rule[0] == p {
						countInFirstPosition++
					}
				}

				pageLeftInRuleCount = append(pageLeftInRuleCount, countInFirstPosition)
			}

			medianIndex := slices.Index(pageLeftInRuleCount, slices.Min(pageLeftInRuleCount)+slices.Max(pageLeftInRuleCount)/2)

			d.solution2 += common.ConvertStringToInt(sp[medianIndex])
		}
	}

	return d.solution1, d.solution2
}
