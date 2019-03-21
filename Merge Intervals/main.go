package main

import (
	"fmt"
	"sort"
	"math"
)

type Interval struct {
	Start int
	End int
}

type intervalList []Interval

func (l intervalList)Len() int{
	return len(l)
}
func (l intervalList)Less(i,j int) bool{
	return l[i].Start < l[j].Start
}
func (l intervalList)Swap(i,j int){
	l[i],l[j] = l[j],l[i]
}


func main() {
	l := intervalList{
		{Start:1,End:2},
		{Start:3,End:4},
		{Start:5,End:6},
		{Start:7,End:8},

	}

	result := merge(l)
	fmt.Println(result)

}

func merge(intervals []Interval) []Interval {
	if len(intervals) <2{
		return intervals
	}
	l := intervalList(intervals)
	sort.Sort(l)
	var result intervalList
	start := l[0].Start
	end := l[0].End
	for i:=1;i<len(l);i++{
		if l[i].Start <= end{
			end = int(math.Max(float64(l[i].End),float64(end)))
		}else{
			result = append(result, Interval{Start:start,End:end})
			start = l[i].Start
			end = l[i].End
		}
	}
	result = append(result, Interval{Start:start,End:end})

	return result
}