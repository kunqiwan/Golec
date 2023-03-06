package main


func targetIndices(nums []int, target int) []int {
	count :=0
	pos:= 0
	//result  := make([]int, 100)
	for n:=0;n<len(nums);n++{
		if nums[n]<target{
			pos++
		}
		if nums[n] == target{
			count++

		}
	}

	//fmt.Println("pos:",pos)
	//fmt.Println("count",count)
	m :=0
	reals := pos
	result := make([]int,count,100)
	if count ==0{
		return result
	}
	for m < count{
		result[m] = reals
		m++
		reals++
	}
	return result
}
