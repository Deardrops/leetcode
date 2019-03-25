package main

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	res := []Interval{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := &res[len(res)-1]
		curr := intervals[i]
		if curr.Start <= last.End {
			if curr.End > last.End {
				last.End = curr.End
			}
		} else {
			res = append(res, curr)
		}
	}
	return res
}
