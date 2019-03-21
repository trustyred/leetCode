package main

import (
	"fmt"
	"math"
)

type Interval struct {
	Start int
	End int
}

func main() {

	l := []Interval{
		{Start:1,End:5},
	}
	result := insert(l, Interval{Start:2,End:3})
	fmt.Print(result)

}
func insert(intervals []Interval, newInterval Interval) []Interval {

	if len(intervals) == 0{
		return []Interval{newInterval}
	}
	start := intervals[0].Start
	end := intervals[0].End
	var result []Interval
	length := len(intervals)
	noEnter := true
	for i:=0;i<length;i++{
		if noEnter && newInterval.Start < intervals[i].Start{

			intervals = append(intervals,newInterval)
			for j:=len(intervals)-1;j>i;j--{
				intervals[j] = intervals[j-1]
			}
			intervals[i] = newInterval
			length = len(intervals)
			noEnter = false
			if i == 0{
				start = intervals[0].Start
				end = intervals[0].End
			}
			fmt.Println(intervals)
		}else if noEnter && i==length-1 && newInterval.Start>= intervals[i].Start{
			intervals = append(intervals, newInterval)
			length = len(intervals)
			noEnter = false
		}
		if intervals[i].Start <= end{
			end = int(math.Max(float64(end),float64(intervals[i].End)))
		}else{
			result = append(result, Interval{Start:start,End:end})
			start = intervals[i].Start
			end = intervals[i].End
		}
	}
	result = append(result, Interval{Start:start,End:end})
	return result
}