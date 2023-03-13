package main

func main() {

}

func unequalTriplets(nums []int) int {
right :=0
distinct :=0
count:=0
left:=0
occ :=make(map[int]int)
for right<len(nums){
	curr:=nums[right]
	ok,_ :=occ[curr]
	if ok>0{
		occ[curr]++
	} else {
		occ[curr]++
		distinct++
	}
	right++
	for distinct==3{
		old :=nums[left]
		occ[old]--
		if occ[old] ==0{
			distinct--
		}
	left++
	count++
	}
}
return count
}
