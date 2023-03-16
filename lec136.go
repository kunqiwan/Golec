package main

func main() {

}
func singleNumber(nums []int) int {
	length :=len(nums)
occ :=make(map[int]int,length)
for _,i :=range nums{
	occ[i]++
}
for key,element :=range occ{
	if element ==1{
		return key
	}
}
return nums[0]
}

// method two - bit operation
//func singleNumber(nums []int) int {
//	single := 0
//	for _, num := range nums {
//		single ^= num
//	}
//	return single
//}

