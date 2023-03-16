package main

func main() {

}
//method 1
//func missingNumber(nums []int) int {
//	sort.Ints(nums)
//	length :=len(nums)
//	for i:=0;i<length;i++{
//		if i!=nums[i]{
//			return i
//		}
//
//	}
//	return length
//}
//method 2
func missingNumber(nums []int) int {
	occ :=make(map[int]bool)
	for _,v :=range nums{
		occ[v] = true
	}
	for i := 0; ; i++ {
		if !occ[i] {
			return i
		}
	}

}