package main

import "fmt"

func main() {
 a :=[]int{0,1,2,2}
	fmt.Println(totalFruit(a))
}
func totalFruit(fruits []int) int {
	if len(fruits) == 0{
		return 0
	}
	if len(fruits) == 1{
		return 1
	}
left :=0
right :=1
count :=0
max :=0
first := fruits[0]
pos := make(map[int]int, 2)
pos[first] = 0
second := -9999999
for right <len(fruits){
	fmt.Println("次数是",right-1)
	fmt.Println("map里有",pos)
	curr :=fruits[right]
	fmt.Println("curr是",curr)
	fmt.Println("first是",first)
	fmt.Println("second是",second)
	_,ok := pos[curr]
	right++
	if ok {
		pos[curr] = right-1
	} else {
		fmt.Println("second值是",second)
		//如果是设定值，说明是第一次，应该给second赋值
		if second == -9999999{
			second =curr
			pos[second] = right -1
		}  else {
			//second有值，应该挪动左滑窗
			if fruits[right-2]==first{
				temp := first
				first = second
				second = temp
			}
			left = pos[first]+1
			delete(pos, first)
			first = second
			second =curr
			pos[second] = right-1
			}

		}
	fmt.Println("right这时是",right)
	fmt.Println("left这时是",left)
	count = right - left
	fmt.Println("count为",count)
	if count>max{
		max = count
	}
	}
	return max
}


