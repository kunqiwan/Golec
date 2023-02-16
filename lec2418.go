package main

import (
	"sort"
)

func main() {

}
func sortPeople(names []string, heights []int) []string {
length :=len(names)
equal := make(map[int]string)
for n:=0;n<length;n++{
	equal[heights[n]] = names[n]
}
sort.Ints(heights)
res :=make([]string,length)
	for m:=length;m<length;m--{
		res[m] = equal[heights[m]]
	}

return res
}