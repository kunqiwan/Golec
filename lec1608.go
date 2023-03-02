package main

import (
	"sort"
)

//func main() {
// nums :=[]int{5,7,3,2,6,1}
//}

func specialArray(nums []int) int {
	sort.Ints(nums)
	length := len(nums)

	for n:=length-1;n>=0;n--{
		mid := n
		le:=0
		re :=length -1
		for le<=re{
			midi :=le+(re-le)/2
			curr:=nums[midi]
			if curr >= mid {
				if length-midi >= mid {
					return mid
				}
				re = midi -1
			} else if curr <mid{
				le = midi+1
			}
		}
	}
	return -1
}