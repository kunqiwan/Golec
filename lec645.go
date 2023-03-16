package main

func main() {

}

func findErrorNums(nums []int) []int {
	res :=make([]int,2)
occ :=make(map[int]int)
for _,value :=range nums{
	occ[value]++
}
for i:=1;i<=len(nums);i++{
	if occ[i] ==0{
		res[1]=i
	}else if occ[i]>1{
		res[0]=i
	}
}
return res
}
