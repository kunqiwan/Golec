package main

func main() {

}
func arithmeticTriplets(nums []int, diff int) int {
save :=map[int]struct{}{}
count :=0
for n:=0;n<len(nums);n++{
save[nums[n]] = struct{}{}
}
for m:=0;m<len(nums);m++{
curr := nums[m]
_,has := save[curr+diff]
if has{
	_,have := save[curr+diff+diff]
	if have{
		count++
	}
}

}
return count
}